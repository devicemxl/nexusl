// Package environment proporciona un entorno para almacenar variables y sus valores.
package environment

import "github.com/devicemxl/nexusl/pkg/object"

// Environment almacena los enlaces de variables y los ámbitos.
type Environment struct {
	// store es un mapa de variables y sus valores.
	store map[string]object.Object
	// outer es un puntero al entorno exterior, para ámbitos anidados.
	outer *Environment
}

// NewEnvironment devuelve un nuevo entorno.
func NewEnvironment() *Environment {
	s := make(map[string]object.Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment devuelve un nuevo entorno anidado.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get devuelve el valor de una variable en el entorno.
func (e *Environment) Get(name string) (object.Object, bool) {
	// Buscar la variable en el entorno actual.
	obj, ok := e.store[name]
	// Si no se encuentra, buscar en el entorno exterior.
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set establece el valor de una variable en el entorno.
func (e *Environment) Set(name string, val object.Object) object.Object {
	e.store[name] = val
	return val
}
