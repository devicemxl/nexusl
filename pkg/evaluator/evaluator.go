// pkg/evaluator/evaluator.go

package evaluator

import (
	"fmt"

	"github.com/devicemxl/nexusl/pkg/ast"
	"github.com/devicemxl/nexusl/pkg/evaluator/environment" // Importar el environment
	"github.com/devicemxl/nexusl/pkg/evaluator/store"       // Importar el store
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

	case *ast.IntegerLiteral: // <--- NUEVO: Para literales enteros
		return &object.IntegerObject{Value: node.Value}

		// ... otros tipos de nodos AST (si los hubiera en el futuro)
	}
	return nil // O un objeto de error
}

func evalProgram(program *ast.Program, env *environment.Environment, knowledgeStore *store.KnowledgeStore) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement, env, knowledgeStore)
		// Si quieres manejar el resultado de cada sentencia, hazlo aquí.
		// Por ejemplo, detener la ejecución si hay un error.
		if _, ok := result.(*object.ErrorObject); ok {
			return result
		}
	}
	return result
}
func evalFlatTripletaStatement(fts *ast.FlatTripletaStatement, env *environment.Environment, knowledgeStore *store.KnowledgeStore) object.Object {
	// 1. Evaluar el Objeto (complemento) de la tripleta. Es un valor que se usará.
	objectVal := Eval(fts.Object, env, knowledgeStore)
	if _, ok := objectVal.(*object.ErrorObject); ok {
		return objectVal
	}

	// 2. Evaluar el Sujeto de la tripleta.
	// Esto es crucial: evalIdentifier ahora devuelve un *SymbolObject* que puede contener miembros.
	subjectVal := Eval(fts.Subject, env, knowledgeStore)
	if _, ok := subjectVal.(*object.ErrorObject); ok {
		return subjectVal
	}
	sObj, ok := subjectVal.(*object.SymbolObject) // Asegurarse de que el sujeto es un SymbolObject.
	if !ok {
		return object.NewError(
			"subject of statement must be a symbol, got %T (%s)",
			subjectVal,
			subjectVal.Type())
	}

	// 3. Obtener el nombre del Verbo (predicado/método).
	verbName := fts.Verb.Value // Todavía es un string literal del AST.

	// === EL CAMBIO FUNDAMENTAL: Decidir la acción basada en fts.Token.Type ===
	switch fts.Token.Type { // NOTA: Esto es `fts.Token.Type`, NO `sObj.Value` (sujeto de la tripleta)
	case tk.DEF:
		// Caso: `def David is:symbol;` o `def David has:rightLeg;`
		// Esto es una *declaración de runtime* o de estructura del lenguaje.

		if verbName == "is:symbol" {
			// Cuando `def David is:symbol;`, `sObj` ya es un SymbolObject porque `evalIdentifier` lo crea.
			// Esta declaración podría usarse para validación o metadatos futuros.
			// Por ahora, solo confirmamos su declaración.
			fmt.Printf("DEF: Declared symbol '%s'\n", sObj.Value)
			return object.NULL
		}
		// `def David has:rightLeg;` -- Como discutimos, esto es para declarar la *existencia* de una propiedad.
		// Podría ser un hecho en CozoDB, o una validación. No crea el slot en el SymbolObject aquí.
		// La creación del slot ocurre en la asignación `David rightLeg "healthy";`
		// Si se quiere que `def` cree un slot vacío, la lógica iría aquí.
		fmt.Printf("DEF: Declared property/meta for '%s': %s %s\n", sObj.Value, verbName, objectVal.String())
		return object.NULL
		// ... (otros tipos de 'def' se añadirían aquí, ej. para definir funciones)

	case tk.FACT:
		// Caso: `fact David is:dummy;` o `fact David walk:"to kitchen";`
		// Esto es una *aserción de conocimiento* para el Knowledge Store.
		// Los componentes de la tripleta (`sObj`, `verbName`, `objectVal`) se convierten a `Term`s.
		subjectTerm := object.NewSymbolTerm(sObj.Value)
		predicateTerm := object.NewSymbolTerm(verbName)
		objTerm := object.NewLiteralTerm(objectVal) // `objectVal` ya es un `NexusL Object` evaluado

		knowledgeTripleta := &object.TripletaObject{
			Subject:   subjectTerm,
			Predicate: predicateTerm,
			Object:    objTerm,
		}
		knowledgeStore.AddTripleta(knowledgeTripleta)
		fmt.Printf("FACT: Asserted %s\n", knowledgeTripleta.String())
		return object.NULL

	case tk.IDENTIFIER:
		// Caso: `cli print "hello";` o `David walk to:"kitchen";` o `David rightLeg "healthy";`
		// Esto es una *invocación de método* o *asignación de atributo*.

		// 1. Intentar como invocación de método: Buscar `verbName` (ej. "print", "walk") en los miembros de `sObj`.
		member, found := sObj.GetMember(verbName)
		if found {
			if builtinFn, isBuiltin := member.(*object.BuiltinObject); isBuiltin {
				// Si es una función builtin (como `cli.print`), invocarla.
				return builtinFn.Fn(objectVal) // `objectVal` es el argumento para la función.
			} else {
				// Si el miembro existe pero NO es una función (ej. `David.rightLeg` ya existe como atributo)
				// Se interpreta como una RE-ASIGNACIÓN de atributo.
				sObj.SetMember(verbName, objectVal)
				fmt.Printf("ASSIGN: %s.%s = %s\n", sObj.Value, verbName, objectVal.String())
				return object.NULL
			}
		} else {
			// 2. Si no es un método existente (ni un miembro existente), se trata como ASIGNACIÓN de un NUEVO atributo.
			sObj.SetMember(verbName, objectVal) // Esto crea el slot 'rightLeg' y le asigna 'healthy'
			fmt.Printf("ASSIGN: %s.%s = %s (new attribute)\n", sObj.Value, verbName, objectVal.String())
			return object.NULL
		}

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

	// Añadir el nuevo símbolo al entorno global para que futuras referencias
	// a este mismo nombre (ej. "cli") obtengan la misma instancia de SymbolObject.
	env.Set(node.Value, sym)

	return sym
}
