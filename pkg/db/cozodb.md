# Fase 2: Introducción de CozoDB como Backend de Conocimiento

Objetivo: Reemplazar el KnowledgeStore en memoria por una integración con CozoDB para el almacenamiento y la consulta de tripletas de conocimiento.

## pkg/db/cozodb.go (NUEVO PAQUETE)

Este paquete encapsulará toda la lógica de interacción con CozoDB.

Necesitarás usar el binding de Go para CozoDB (si existe uno oficial y estable, o crearlo si es necesario).

Funciones clave:
- Connect(dsn string) (*CozoDB, error): Establece la conexión.
- AddTripleta(subject, predicate, object string): Inserta una tripleta como un hecho en CozoDB. CozoDB usa un modelo de tabla, por lo que una tabla knowledge(subject, predicate, object) sería lo natural.
- Query(datalogQuery string) ([][]interface{}, error): Ejecuta consultas Datalog y devuelve resultados. Esto es crucial para la inferencia.
- DefineRule(datalogRule string) error: Permite añadir reglas Datalog a la base de datos (si CozoDB lo soporta dinámicamente).
- Modificación del Evaluador (pkg/evaluator/evaluator.go):

El evaluator dejará de usar pkg/evaluator/store/store.go directamente.

En su lugar, recibirá una instancia de *db.CozoDB.

Cuando procese una FlatTripletaStatement de tipo assert o cualquier otra forma de declarar un hecho, invocará cozodb.AddTripleta().

Cuando se necesite una consulta (futura sintaxis query X is a Y;), el evaluador formulará la consulta Datalog y la enviará a cozodb.Query().