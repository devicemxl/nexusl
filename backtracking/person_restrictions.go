package main

import (
	"fmt"
)

// DominioFD representa un dominio mutable con un conjunto de valores posibles.
type DominioFD struct {
	Valores []map[interface{}]interface{} // Cada valor es un mapa con atributos y sus valores
}

// Filtrar reduce el dominio manteniendo solo los valores que cumplen la condición.
func (d *DominioFD) Filtrar(filtro func(map[interface{}]interface{}) bool) {
	var nuevos []map[interface{}]interface{}
	for _, val := range d.Valores {
		if filtro(val) {
			nuevos = append(nuevos, val)
		}
	}
	d.Valores = nuevos
}

// Propagador es una función que filtra dominios y verifica la consistencia.
type Propagador func(asignacion map[interface{}]map[interface{}]interface{}, dominios map[interface{}]*DominioFD) bool

// CrearPropagadorIgualdadLugar genera un propagador que asegura igualdad en "lugar".
func CrearPropagadorIgualdad(v1, v2 interface{}, restrict interface{}) Propagador {
	return func(asig map[interface{}]map[interface{}]interface{}, doms map[interface{}]*DominioFD) bool {
		p1 := asig[v1]
		p2 := asig[v2]

		if p1 != nil && p2 != nil {
			return p1[restrict] == p2[restrict]
		} else if p1 != nil {
			val := p1[restrict]
			doms[v2].Filtrar(func(plan map[interface{}]interface{}) bool {
				return plan[restrict] == val
			})
		} else if p2 != nil {
			val := p2[restrict]
			doms[v1].Filtrar(func(plan map[interface{}]interface{}) bool {
				return plan[restrict] == val
			})
		}
		return true
	}
}

// BacktrackingPropagado realiza backtracking con propagación de restricciones.
func BacktrackingPropagado(
	variables []interface{},
	dominios map[interface{}]*DominioFD,
	asignacion map[interface{}]map[interface{}]interface{},
	propagadores []Propagador,
) []map[interface{}]map[interface{}]interface{} {

	var soluciones []map[interface{}]map[interface{}]interface{}

	// Verificar si la asignación es completa
	completa := true
	for _, nombre := range variables {
		if asignacion[nombre] == nil {
			completa = false
			break
		}
	}
	if completa {
		copia := make(map[interface{}]map[interface{}]interface{})
		for k, v := range asignacion {
			copia[k] = v
		}
		return []map[interface{}]map[interface{}]interface{}{copia}
	}

	// Seleccionar siguiente variable sin asignar
	var actual interface{}
	for _, nombre := range variables {
		if asignacion[nombre] == nil {
			actual = nombre
			break
		}
	}

	// Probar cada valor posible del dominio
	for _, valor := range dominios[actual].Valores {
		nuevaAsignacion := copiarAsignacion(asignacion)
		nuevaAsignacion[actual] = valor

		nuevosDominios := copiarDominios(dominios)

		// Propagar restricciones
		valido := true
		for _, p := range propagadores {
			if !p(nuevaAsignacion, nuevosDominios) {
				valido = false
				break
			}
		}

		if valido {
			subs := BacktrackingPropagado(variables, nuevosDominios, nuevaAsignacion, propagadores)
			soluciones = append(soluciones, subs...)
		}
	}

	return soluciones
}

// Funciones auxiliares para copiar estado (asignaciones y dominios)
func copiarAsignacion(original map[interface{}]map[interface{}]interface{}) map[interface{}]map[interface{}]interface{} {
	copia := make(map[interface{}]map[interface{}]interface{})
	for k, v := range original {
		copia[k] = v
	}
	return copia
}

func copiarDominios(orig map[interface{}]*DominioFD) map[interface{}]*DominioFD {
	nuevo := make(map[interface{}]*DominioFD)
	for k, dom := range orig {
		vals := make([]map[interface{}]interface{}, len(dom.Valores))
		copy(vals, dom.Valores)
		nuevo[k] = &DominioFD{Valores: vals}
	}
	return nuevo
}

func main() {
	// Definimos los dominios iniciales de cada persona
	dominios := map[interface{}]*DominioFD{
		"Alice": {
			Valores: []map[interface{}]interface{}{
				{"lugar": "parque", "horario": 1},
				{"lugar": "cafetería", "horario": 3},
				{"lugar": "cine", "horario": 2},
			},
		},
		"Bob": {
			Valores: []map[interface{}]interface{}{
				{"lugar": "parque", "horario": 1},
				{"lugar": "cafetería", "horario": 3},
				{"lugar": "cine", "horario": 4},
			},
		},
	}

	variables := []interface{}{"Alice", "Bob"}
	asignacion := make(map[interface{}]map[interface{}]interface{})

	// Propagadores para imponer igualdad en "lugar" entre Alice y Bob
	propagadores := []Propagador{
		CrearPropagadorIgualdad("Alice", "Bob", "horario"),
	}

	sols := BacktrackingPropagado(variables, dominios, asignacion, propagadores)

	for i, sol := range sols {
		fmt.Printf("\nSolución %d:\n", i+1)
		for k, v := range sol {
			fmt.Printf("%s va a %s en la %s\n", k, v["lugar"], v["horario"])
		}
	}
}
