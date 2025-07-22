// main.go
package main

import (
	"fmt"

	"github.com/devicemxl/nexusl/ds"                       // Importa tu paquete de Símbolos
	prologo "github.com/devicemxl/nexusl/internal/proloGo" // Importa tu paquete de Unificación
)

func main() {
	fmt.Println("INFO: main() function started.") // <-- AÑADE ESTO
	fmt.Println("--- Ejemplos de Unificación en ProloGo ---")

	// Ejemplo 1: Unificación de Constantes
	fmt.Println("\n--- Ejemplo 1: Constantes ---")
	constA := ds.NewConstantSymbol("a", "alpha")
	constB := ds.NewConstantSymbol("b", "beta")
	constA2 := ds.NewConstantSymbol("a", "alpha") // Otra constante con el mismo valor

	env1 := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", constA.PublicName, constA2.PublicName, prologo.Unify(constA, constA2, env1))
	fmt.Println("Ligaduras resultantes (Ej1a):", env1.Bindings) // Debería estar vacío si no hay variables

	env1b := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", constA.PublicName, constB.PublicName, prologo.Unify(constA, constB, env1b))
	fmt.Println("Ligaduras resultantes (Ej1b):", env1b.Bindings)

	// Ejemplo 2: Unificación de Variable con Constante
	fmt.Println("\n--- Ejemplo 2: Variable con Constante ---")
	variableX := ds.NewVariableSymbol("X")
	constValue := ds.NewConstantSymbol("valor_constante", 123)

	env2 := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", variableX.PublicName, constValue.PublicName, prologo.Unify(variableX, constValue, env2))
	fmt.Printf("X desreferenciado: %s\n", prologo.Deref(variableX, env2).String())
	fmt.Println("Ligaduras resultantes (Ej2):", env2.Bindings)

	// Ejemplo 3: Unificación de dos Variables
	fmt.Println("\n--- Ejemplo 3: Dos Variables ---")
	variableY := ds.NewVariableSymbol("Y")
	variableZ := ds.NewVariableSymbol("Z")

	env3 := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", variableY.PublicName, variableZ.PublicName, prologo.Unify(variableY, variableZ, env3))
	fmt.Printf("Y desreferenciado: %s\n", prologo.Deref(variableY, env3).String()) // Y debería apuntar a Z
	fmt.Printf("Z desreferenciado: %s\n", prologo.Deref(variableZ, env3).String()) // Z debería apuntar a Z
	fmt.Println("Ligaduras resultantes (Ej3):", env3.Bindings)

	// Ejemplo 4: Unificación de Listas
	fmt.Println("\n--- Ejemplo 4: Listas ---")
	// Lista 1: [a, X]
	listHeadA := ds.NewConstantSymbol("a", "a")
	listVarX := ds.NewVariableSymbol("X")
	list1 := ds.NewListSymbol(listHeadA, ds.NewListSymbol(listVarX, ds.NullSymbol))

	// Lista 2: [A, b]
	listVarA := ds.NewVariableSymbol("A")
	listHeadB := ds.NewConstantSymbol("b", "b")
	list2 := ds.NewListSymbol(listVarA, ds.NewListSymbol(listHeadB, ds.NullSymbol))

	env4 := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", list1.String(), list2.String(), prologo.Unify(list1, list2, env4))
	fmt.Printf("X desreferenciado: %s\n", prologo.Deref(listVarX, env4).String())
	fmt.Printf("A desreferenciado: %s\n", prologo.Deref(listVarA, env4).String())
	fmt.Println("Ligaduras resultantes (Ej4):", env4.Bindings)

	// Ejemplo 5: Unificación de Estructuras
	fmt.Println("\n--- Ejemplo 5: Estructuras ---")
	functorLikes := ds.NewConstantSymbol("likes", "likes")
	functorHates := ds.NewConstantSymbol("hates", "hates")
	personJohn := ds.NewConstantSymbol("john", "John")
	//personMary := ds.NewConstantSymbol("mary", "Mary")
	objectFood := ds.NewConstantSymbol("food", "Food")
	structVarP := ds.NewVariableSymbol("P")

	// Estructura 1: likes(john, X)
	struct1 := ds.NewStructureSymbol(functorLikes, []*ds.Symbol{personJohn, structVarP})

	// Estructura 2: likes(Y, food)
	structVarY := ds.NewVariableSymbol("Y")
	struct2 := ds.NewStructureSymbol(functorLikes, []*ds.Symbol{structVarY, objectFood})

	env5 := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", struct1.String(), struct2.String(), prologo.Unify(struct1, struct2, env5))
	fmt.Printf("P desreferenciado: %s\n", prologo.Deref(structVarP, env5).String())
	fmt.Printf("Y desreferenciado: %s\n", prologo.Deref(structVarY, env5).String())
	fmt.Println("Ligaduras resultantes (Ej5):", env5.Bindings)

	// Ejemplo 6: Fallo de Unificación (estructuras con diferente functor)
	fmt.Println("\n--- Ejemplo 6: Fallo por Functor ---")
	// Estructura 3: hates(john, X)
	struct3 := ds.NewStructureSymbol(functorHates, []*ds.Symbol{personJohn, structVarP}) // Reutiliza P

	env6 := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", struct1.String(), struct3.String(), prologo.Unify(struct1, struct3, env6))
	fmt.Println("Ligaduras resultantes (Ej6):", env6.Bindings) // Debería estar vacío

	// Ejemplo 7: Fallo de Unificación (constantes diferentes)
	fmt.Println("\n--- Ejemplo 7: Fallo por Constante ---")
	varA := ds.NewVariableSymbol("A")
	varB := ds.NewVariableSymbol("B")
	c1 := ds.NewConstantSymbol("c1", 10)
	c2 := ds.NewConstantSymbol("c2", 20)

	listConst1 := ds.NewListSymbol(varA, ds.NewListSymbol(c1, ds.NullSymbol))
	listConst2 := ds.NewListSymbol(c2, ds.NewListSymbol(varB, ds.NullSymbol))

	env7 := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", listConst1.String(), listConst2.String(), prologo.Unify(listConst1, listConst2, env7))
	fmt.Printf("A desreferenciado: %s\n", prologo.Deref(varA, env7).String()) // A debería ser 20
	fmt.Printf("B desreferenciado: %s\n", prologo.Deref(varB, env7).String()) // B debería ser 10
	fmt.Println("Ligaduras resultantes (Ej7):", env7.Bindings)                // No debería haber ligaduras para A y B si falla

	// Ejemplo 8: Unificación con variable anónima
	fmt.Println("\n--- Ejemplo 8: Variable Anónima ---")
	anonVar := ds.AnonymousSymbol
	someConst := ds.NewConstantSymbol("anything", "Hola Mundo")
	someVar := ds.NewVariableSymbol("Var")

	env8a := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", anonVar.String(), someConst.String(), prologo.Unify(anonVar, someConst, env8a))
	fmt.Println("Ligaduras resultantes (Ej8a):", env8a.Bindings) // Siempre true, sin ligaduras

	env8b := prologo.NewEnvironment()
	fmt.Printf("Unificar %s y %s: %t\n", anonVar.String(), someVar.String(), prologo.Unify(anonVar, someVar, env8b))
	fmt.Println("Ligaduras resultantes (Ej8b):", env8b.Bindings) // Siempre true, sin ligaduras para 'someVar'

}
