package main

import "fmt"

// expandCompositeTriplet
/*
Explicación
expandCompositeTriplet: Esta función toma un FlatTripletStatement compuesto y verifica si el campo Statement del objeto es un slice de verbos. Si es así, crea un triplet simple para cada verbo en el slice.

Iteración sobre Verbos: Para cada verbo en el slice, se crea un nuevo FlatTripletStatement con el mismo sujeto, ámbito y condición, pero con el verbo actual como el MainVerb.

Función MainVerb: En este ejemplo, la función MainVerb simplemente devuelve una cadena que indica que la entidad está realizando la acción del verbo. Puedes personalizar esta función según tus necesidades específicas.

Este enfoque te permite manejar triplets compuestos y expandirlos en triplets simples, lo cual es útil para representar y procesar múltiples acciones o funciones asociadas a un sujeto de manera eficiente.
*/
func expandCompositeTriplet(composite FlatTripletStatement) []FlatTripletStatement {
	var triplets []FlatTripletStatement

	// Asumimos que el objeto contiene un slice de verbos en el campo Statement
	verbs, ok := composite.Object.Statement.([]string)
	if !ok {
		// Manejar el caso en que el objeto no es un slice de verbos
		return []FlatTripletStatement{composite}
	}

	// Iterar sobre cada verbo y crear un triplet simple
	for _, verb := range verbs {
		triplet := FlatTripletStatement{
			Scope:   composite.Scope,
			Subject: composite.Subject,
			Verb: VerbStatement{
				ModalVerb: composite.Verb.ModalVerb,
				MainVerb: func(entity *Entity) (interface{}, error) {
					// Aquí puedes definir la función específica para cada verbo si es necesario
					return fmt.Sprintf("%s is %s", entity.PublicName, verb), nil
				},
			},
			Object: ObjectStatement{
				Condition: composite.Object.Condition,
				Statement: verb,
			},
		}
		triplets = append(triplets, triplet)
	}

	return triplets
}

// expandCompositeTripletWithAttributes
/*

expandCompositeTripletWithAttributes: Esta función toma un FlatTripletStatement compuesto y verifica si el campo Condition del objeto es un slice de atributos. Si es así, crea un triplet simple para cada atributo en el slice.

Iteración sobre Atributos: Para cada atributo en el slice, se crea un nuevo FlatTripletStatement con el mismo sujeto, ámbito, verbo y declaración, pero con el atributo actual como la Condition.

Función MainVerb: En este ejemplo, la función MainVerb devuelve una cadena que indica que la entidad está realizando la acción del verbo. Puedes personalizar esta función según tus necesidades específicas.

Este enfoque te permite manejar triplets compuestos con múltiples atributos y expandirlos en triplets simples, lo cual es útil para representar y procesar múltiples atributos asociados a una acción de manera eficiente.

*/

func expandCompositeTripletWithAttributes(composite FlatTripletStatement) []FlatTripletStatement {
	var triplets []FlatTripletStatement

	// Asumimos que el objeto contiene un slice de atributos en el campo Statement
	attributes, ok := composite.Object.Statement.([]string)
	if !ok {
		// Manejar el caso en que el objeto no es un slice de atributos
		return []FlatTripletStatement{composite}
	}

	// Iterar sobre cada atributo y crear un triplet simple
	for _, attribute := range attributes {
		triplet := FlatTripletStatement{
			Scope:   composite.Scope,
			Subject: composite.Subject, // Asegúrate de copiar el sujeto
			Verb: VerbStatement{
				ModalVerb: composite.Verb.ModalVerb,
				MainVerb:  composite.Verb.MainVerb,
			},
			Object: ObjectStatement{
				Condition: attribCondType(attribute), // Conversión segura de tipos
				Statement: composite.Object.Statement,
			},
		}
		triplets = append(triplets, triplet)
	}

	return triplets
}

/*

func main() {
    // Crear un triplet compuesto
    compositeTriplet := FlatTripletStatement{
        Scope: "def",
        Subject: 1, // EntityID para "david"
        Verb: VerbStatement{
            ModalVerb: "could",
        },
        Object: ObjectStatement{
            Condition: "how:fast",
            Statement: []string{"run", "eat", "drive"},
        },
    }

    // Expandir el triplet compuesto en triplets simples
    triplets := expandCompositeTriplet(compositeTriplet)

    // Imprimir los triplets simples resultantes
    for _, triplet := range triplets {
        fmt.Printf("Scope: %s, Subject: %d, Verb: %v, Object: %v\n",
            triplet.Scope, triplet.Subject, triplet.Verb, triplet.Object)
    }




    // Crear un triplet compuesto con atributos
    compositeTriplet := FlatTripletStatement{
        Scope: "def",
        Subject: 1, // EntityID para "david"
        Verb: VerbStatement{
            ModalVerb: "do",
            MainVerb: func(entity *Entity) (interface{}, error) {
                return fmt.Sprintf("%s is running", entity.PublicName), nil
            },
        },
        Object: ObjectStatement{
            Condition: "how", // Este campo podría no ser necesario si usas Statement para atributos
            Statement: []string{"fast", "silently", "happy"},
        },
    }

    // Expandir el triplet compuesto en triplets simples
    triplets := expandCompositeTripletWithAttributes(compositeTriplet)

    // Imprimir los triplets simples resultantes
    for _, triplet := range triplets {
        fmt.Printf("Scope: %s, Subject: %d, Verb: %v, Object: %v\n",
            triplet.Scope, triplet.Subject, triplet.Verb, triplet.Object)
    }

}
*/
