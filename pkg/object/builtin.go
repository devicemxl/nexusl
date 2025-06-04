// pkg/object/builtin.go

package object

import "fmt"

// BuiltinFunction representa una función nativa del sistema.
// Es la signatura para todas las funciones built-in de NexusL.
type BuiltinFunction func(args ...Object) Object

// BuiltinObject representa una función built-in en el sistema de tipos de NexusL.
type BuiltinObject struct {
	Fn BuiltinFunction
}

func (b *BuiltinObject) Type() ObjectType { return BUILTIN_OBJ }
func (b *BuiltinObject) String() string   { return "builtin function" }

// =============================================================
// Aquí es donde se definen las funciones built-in concretas.
// =============================================================

// Builtins es un mapa que almacena todas las funciones built-in disponibles.
// Se usará para inicializar el entorno global o los SymbolObjects como 'cli'.
var Builtins = map[string]*BuiltinObject{
	"print": {Fn: BuiltinPrint}, // Ejemplo: 'print'
	// "add": {Fn: BuiltinAdd},     // Ejemplo: 'add' para suma, si la expones como builtin
	// "len": {Fn: BuiltinLen},     // Ejemplo: 'len' para longitud de algo
	// "typeof": {Fn: BuiltinTypeof}, // Ejemplo: 'typeof' para obtener el tipo de un objeto
	// ... agrega más built-ins aquí
}

// BuiltinPrint es una función built-in que imprime el argumento.
func BuiltinPrint(args ...Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}
	fmt.Println(args[0].String()) // Imprime la representación string del objeto
	return NULL                   // 'print' no devuelve un valor útil, así que devuelve NULL
}

// Ejemplo de otro built-in (descomenta si lo necesitas)
/*
func BuiltinAdd(args ...Object) Object {
    if len(args) != 2 {
        return NewError("wrong number of arguments. got=%d, want=2", len(args))
    }
    num1, ok1 := args[0].(*IntegerObject)
    num2, ok2 := args[1].(*IntegerObject)

    if !ok1 || !ok2 {
        return NewError("arguments to 'add' must be numbers, got %T and %T", args[0], args[1])
    }
    return &IntegerObject{Value: num1.Value + num2.Value}
}
*/
