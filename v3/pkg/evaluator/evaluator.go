// pkg/evaluator/evaluator.go

package evaluator

import (
	"fmt"

	"github.com/devicemxl/nexusl/pkg/ast"
	"github.com/devicemxl/nexusl/pkg/evaluator/environment"
	"github.com/devicemxl/nexusl/pkg/evaluator/store"
	"github.com/devicemxl/nexusl/pkg/object"
	tk "github.com/devicemxl/nexusl/pkg/token"
)

// Eval evalúa un nodo del AST y devuelve un objeto de NexusL.
func Eval(node ast.Node,
	env *environment.Environment,
	knowledgeStore *store.KnowledgeStore) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env, knowledgeStore)

	case *ast.FlatTripletaStatement:
		return evalFlatTripletaStatement(node, env, knowledgeStore)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.StringLiteral:
		return &object.StringObject{Value: node.Value}

	case *ast.IntegerLiteral:
		return &object.IntegerObject{Value: node.Value}

	// Agrega casos para otros tipos de expresiones si los tienes (ej. CallExpression)
	// case *ast.CallExpression:
	// 	return evalCallExpression(node, env, knowledgeStore)

	default:
		return object.NewError("unknown node type for evaluation: %T", node)
	}
}

func evalProgram(program *ast.Program, env *environment.Environment, knowledgeStore *store.KnowledgeStore) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement, env, knowledgeStore)
		// Manejar errores si es necesario (ej. si result es un ErrorObject)
		if isError(result) {
			return result
		}
	}
	return result
}

func evalFlatTripletaStatement(fts *ast.FlatTripletaStatement, env *environment.Environment, knowledgeStore *store.KnowledgeStore) object.Object {
	// Evaluar sujeto y objeto primero. El verbo (predicado) es un Identifier puro.
	subjectObj := Eval(fts.Subject, env, knowledgeStore)
	if isError(subjectObj) {
		return subjectObj
	}

	objectVal := Eval(fts.Object, env, knowledgeStore) // El objeto puede ser NULL si no está presente
	if isError(objectVal) {
		return objectVal
	}

	verbName := fts.Verb.Value // El nombre literal del predicado/verbo

	switch fts.Token.Type {
	case tk.DEF:
		// Para `def`, queremos definir un símbolo o una estructura.
		// `def David is:symbol;` o `def agentX has:action (moveTo roomB);`
		// Aquí, `is:symbol` o `has:action` actuarían como "definidores" de tipo o capacidad.
		// El `objectVal` podría ser un StringObject("symbol") o una estructura de acción.

		// Asumiendo que `subjectObj` es un SymbolObject para `def`
		sObj, ok := subjectObj.(*object.SymbolObject)
		if !ok {
			return object.NewError("DEF expects a symbol as subject, got %s", subjectObj.Type())
		}

		// Ejemplo: `def David is:person;`
		// Esto se almacenaría como un hecho: `(David is:person null)` o `(David type person)`
		// Y potencialmente se añadiría a los miembros del símbolo David si `is:person` define una propiedad.
		// Por ahora, lo registramos como un hecho general.
		// Considera una forma de distinguir entre una "definición de tipo" y una "asignación de propiedad inicial".

		// Convertir el sujeto, predicado y objeto en Terms para la Tripleta.
		subjectTerm := object.NewSymbolTerm(sObj.Value) // Sujeto es siempre un símbolo para DEF
		predicateTerm := object.NewSymbolTerm(verbName) // Predicado es siempre un símbolo para DEF
		objectTerm := object.NewLiteralTerm(objectVal)  // Objeto es un literal (puede ser string, int, bool)

		tripleta := &object.TripletaObject{
			Subject:   subjectTerm,
			Predicate: predicateTerm,
			Object:    objectTerm,
		}
		knowledgeStore.AddTripleta(tripleta)
		fmt.Printf("DEFINED FACT: %s %s %s;\n", tripleta.Subject.String(), tripleta.Predicate.String(), tripleta.Object.String())
		return object.NULL

	case tk.FACT:
		// Para `fact`, queremos asentar un hecho en la base de conocimiento.
		// `fact David has:rightLeg "healthy";`
		// `fact motor1 experiences overheating;`

		// Convertir subjectObj a SymbolObject
		sObj, ok := subjectObj.(*object.SymbolObject)
		if !ok {
			return object.NewError("FACT expects a symbol as subject, got %s", subjectObj.Type())
		}

		// Convertir el sujeto, predicado y objeto en Terms para la Tripleta.
		subjectTerm := object.NewSymbolTerm(sObj.Value)
		predicateTerm := object.NewSymbolTerm(verbName)
		// El objeto puede ser un literal, un símbolo o una referencia.
		// Aquí asumimos que objectVal ya es un object.Object evaluado.
		var objectTerm *object.Term
		if objectVal == object.NULL {
			objectTerm = object.NewLiteralTerm(object.NULL) // O un Term específico para "nulo" si es significativo
		} else {
			objectTerm = object.NewLiteralTerm(objectVal) // Convertir el object.Object a un TermLiteral
		}

		tripleta := &object.TripletaObject{
			Subject:   subjectTerm,
			Predicate: predicateTerm,
			Object:    objectTerm,
		}
		knowledgeStore.AddTripleta(tripleta)
		fmt.Printf("FACT ADDED: %s %s %s;\n", tripleta.Subject.String(), tripleta.Predicate.String(), tripleta.Object.String())

		// Opcional: También podrías querer que los hechos actualicen los miembros del SymbolObject.
		// Esto crea una dualidad entre la base de conocimiento (verdad declarativa)
		// y el estado del objeto (representación programática/cache).
		// Decidir si mantener esta dualidad o si el acceso a propiedades siempre consulta el KS.
		// Por ahora, mantengamos la actualización de miembros si el predicado es una propiedad directa.
		sObj.SetMember(verbName, objectVal) // Asigna la propiedad al símbolo
		return object.NULL

	case tk.IDENTIFIER: // Esto cubre las invocaciones directas como `cli print "hello";`
		// En este caso, `fts.Token.Literal` es el sujeto (ej. "cli")
		// `fts.Subject` ya fue evaluado a `subjectObj` (ej. SymbolObject "cli")
		// `fts.Verb` es el nombre de la acción/método (ej. "print")
		// `fts.Object` es el argumento (ej. StringLiteral "hello")

		sObj, ok := subjectObj.(*object.SymbolObject)
		if !ok {
			return object.NewError("Expected symbol for invocation, got %s", subjectObj.Type())
		}

		// Intenta obtener el miembro (método/acción) del sujeto
		method := sObj.GetMember(verbName)
		if method == nil {
			return object.NewError("Symbol '%s' has no member or action '%s'", sObj.Value, verbName)
		}

		// Si el miembro es una función BuiltinObject, invócala.
		if builtinFn, isBuiltin := method.(*object.BuiltinObject); isBuiltin {
			// El `objectVal` es el argumento. Las built-ins pueden tomar varios, pero `print` toma uno.
			// Necesitamos manejar esto de forma más general para futuras funciones.
			// Por ahora, asumimos que `objectVal` es el único argumento.
			if objectVal == object.NULL { // Si no hay objeto explícito, pasar nulo o un array vacío
				return builtinFn.Fn() // Invocar sin argumentos
			}
			return builtinFn.Fn(objectVal) // Invocar con el objeto como argumento
		}

		// Si es otro tipo de objeto (no una función invocable), ¿qué significa?
		// Podría ser una "referencia" a una acción, pero no una invocación directa.
		return object.NewError("Member '%s' of '%s' is not an invokable action/function (type: %s)", verbName, sObj.Value, method.Type())

	default:
		// Esto debería ser un error del parser si llega aquí, ya que parseStatement debería filtrar.
		return object.NewError("unknown statement type for FlatTripletaStatement: %s", fts.Token.Literal)
	}
}

// evalIdentifier evalúa un identificador.
func evalIdentifier(node *ast.Identifier, env *environment.Environment) object.Object {
	// 1. Verificar si el identificador es una variable existente en el entorno.
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	// 2. Si no es una variable existente, asumir que es un nuevo símbolo global.
	// ¡CRUCIAL!: Usar NewSymbolObject para asegurar que el mapa 'Members' esté inicializado.
	sym := object.NewSymbolObject(node.Value)

	// 3. Registrar built-ins en el símbolo 'cli' si 'cli' es el identificador
	// Esta es una forma simple de inyectar built-ins en un símbolo específico.
	// En un sistema más maduro, podrías tener un espacio de nombres de built-ins
	// o cargar funciones built-in en el entorno global al inicio.
	if node.Value == "cli" {
		for name, builtin := range object.Builtins {
			sym.SetMember(name, builtin)
		}
	}

	env.Set(node.Value, sym) // Añadir el nuevo símbolo al entorno
	return sym
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
