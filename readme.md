<center>

# NexusL 

<strong>*Code is knowledge and knowledge is code*.</strong>

</center>


Este repositorio contiene el diseño y la implementación inicial de NexusL, un Lenguaje de Dominio Específico para Conocimiento e Inferencia, bajo el paradigma de representación en tripletas (sujeto, predicado, objeto).

## Objetivos de Diseño de NexusL

* **Todo se expresa como tripletas:** 
    - Cada unidad de conocimiento es una estructura `(sujeto predicado objeto)`. El objeto puede ser un literal (número, string, booleano), un símbolo, una acción invocable, o una referencia a otra tripleta/estructura.
* **Sintaxis basada en S-expressions (estilo Lisp):** 
    - Aprovechando la homoiconicidad para una representación de código como datos; *Code is knowledge and knowledge is code*.
* **Enfoque declarativo:** La base del conocimiento es dinámica y evoluciona con el tiempo.
* **Capacidad de actuar sobre el mundo:** Permitir la ejecución de acciones además de la mera representación de hechos.
* **Homoiconicidad:** El código es dato, facilitando la metaprogramación y la manipulación del lenguaje.
* **Inferencia lógica tipo Prolog:** Unificación, reglas, y encadenamiento (backward y forward).
* **Expresividad matemática:** Soporte para aritmética básica y extendida.
* **Distinción entre hechos, acciones y funciones:** Claridad en el modelo de conocimiento.
* **Modelo generalista:** No limitado a ontologías o dominios fijos, buscando ser extensible.

## Estrategia de Desarrollo y Herramientas

La implementación de NexusL se realizará en Go, aprovechando sus capacidades para el desarrollo de compiladores e intérpretes. El proceso se dividirá en fases incrementales para asegurar una base sólida y una integración modular de funcionalidades avanzadas.

### Fase 1: El Núcleo Homoicónico y la Representación de Tripletas (Go)

En esta fase inicial, nos centraremos en la creación de la infraestructura fundamental del lenguaje, incluyendo el análisis léxico, sintáctico, la representación de datos en memoria y un evaluador básico.

* **Lenguaje Base:** **Go**
* **Componentes Clave:**
    * **Definición de la `Tripleta`:** La unidad fundamental de conocimiento.
    * **Lexer (Analizador Léxico):** Convertirá el código fuente (S-expressions) en tokens.
    * **Parser (Analizador Sintáctico):** Construirá el Árbol de Sintaxis Abstracta (AST) a partir de los tokens, reflejando la homoiconicidad.
    * **Evaluador/Intérprete Básico:** Procesará el AST para almacenar hechos simples y ejecutar funciones nativas (built-ins).
    * **Sistema de Almacenamiento:** Una estructura de datos en memoria para la base de conocimiento inicial.

### Fase 2: Inferencia Lógica y Variables (Integración de Prolog/Datalog)

Una vez que el núcleo del lenguaje pueda manejar tripletas y S-expressions, se introducirá la capacidad de inferencia lógica.

* **Motor de Inferencia:** Se priorizará una implementación propia de los algoritmos de unificación y resolución (SLD Resolution) en Go para un control total y una comprensión profunda.
* **Manejo de Variables:** Introducción de tipos de términos para variables lógicas.
* **Reglas Lógicas:** Definición y procesamiento de reglas para la inferencia.
* **Sintaxis para Reglas y Consultas:** Extensión del lenguaje para definir y consultar reglas.

### Fase 3: Expresividad Matemática y Acciones/Reactivity

Esta fase se enfocará en enriquecer la expresividad del lenguaje y su capacidad para interactuar con el "mundo".

* **Expresividad Matemática:** Implementación de operaciones aritméticas y funciones matemáticas como built-ins.
* **Acciones Invocables y Efectos Secundarios:** Mecanismos claros para distinguir y ejecutar acciones con efectos sobre el entorno.
* **Reactivity:** Introducción de patrones para la reactividad, como observadores, encadenamiento hacia adelante (forward chaining) y manejo de eventos.

### Consideraciones Adicionales

* **Namespaces:** Un sistema para organizar y evitar colisiones de nombres.
* **Metadatos:** Posibilidad de asociar metadatos a las tripletas (ej. fuente, confianza).
* **Error Handling y Debugging:** Mecanismos robustos para el manejo de errores y la depuración del código.

## Estructura de Directorios Propuesta Inicialmente

La siguiente estructura facilita la modularidad y el desarrollo organizado:

```
nexusl/
├── cmd/
│   └── nexusl/
│       └── main.go       // Punto de entrada del intérprete/CLI (REPL)
├── pkg/
│   ├── ast/
│   │   └── ast.go        // Definiciones del Árbol de Sintaxis Abstracta (AST)
│   │   └── expressions.go
│   │   └── literals.go 
│   │   └── statlements.go 
│   ├── compiler/         // Empty
│   ├── db/               // Empty
│   ├── evaluator/
│   │   └── evaluator.go  // Intérprete/Evaluador principal del AST
│   │   └── store
│   │   |   └── store.go  // Almacén de conocimiento en memoria (hechos)
│   │   └── environment
│   │       └── environment.go // Contexto de ejecución (variables, namespaces)
│   ├── lexer/
│   │   └── lexer.go      // Analizador Léxico (scanner)
│   ├── object/           // Representación de los tipos de datos de NexusL en tiempo de ejecución
│   │   └── builtin.go    // Implementación de funciones nativas (built-ins)
│   │   └── literals.go   // 
│   │   └── object.go     // Definición de Term, Tripleta, etc. como objetos de runtime
│   │   └── symbols.go    // 
│   │   └── triplet.go    // 
│   ├── parser/
│   │   └── parser.go     // Analizador Sintáctico (parser)
│   ├── inference/        // Componentes para la inferencia lógica
│   ├── compiler/         // (Opcional, para fases posteriores: compilación a bytecode)
│   └── util/             // Utilidades generales
├── tests/                // Directorio para pruebas unitarias e integración
│   └── ...
├── go.mod                // Módulo Go (gestionado por `go mod`)
├── go.sum                // Sumas de verificación de módulos
└── README.md             // Este archivo
```

## Próximos Pasos para Empezar

Para iniciar el desarrollo, se recomienda seguir estos pasos:

1.  **Crear la estructura de directorios** descrita arriba.
2.  **Inicializar el módulo Go:** Ejecutar `go mod init github.com/tu_usuario/nexusl` (reemplaza `tu_usuario` con tu nombre de usuario o el path del repositorio).
3.  **Definir las estructuras principales:**
    * En `pkg/object/object.go`: Define la interfaz `Object` y las structs concretas como `NumberObject`, `StringObject`, `BooleanObject`, `SymbolObject`, `TripletaObject`, y `VariableObject`. Estas son las representaciones en tiempo de ejecución de los datos de NexusL.
    * En `pkg/parser/ast.go`: Define la interfaz `Node` y `Expression`, y las structs para los nodos del Árbol de Sintaxis Abstracta como `Program`, `Tripleta`, `Literal`, `Symbol`, `Variable`, `CallExpression`, etc.
4.  **Implementar el Léxer (`pkg/lexer/lexer.go`):** Capaz de tokenizar S-expressions básicas (paréntesis, símbolos, números, strings).
5.  **Implementar el Parser inicial (`pkg/parser/parser.go`):** Que pueda construir un AST para una única tripleta `(sujeto predicado objeto)` y expresiones anidadas simples.
6.  **Crear un `main.go` simple en `cmd/nexusl/`:** Para leer una línea de entrada, pasarla por el léxer y el parser, e imprimir el AST resultante. Esto servirá para validar las primeras etapas.
7.  **Empezar a escribir pruebas unitarias** para el léxer y el parser desde el principio, ya que son componentes críticos.

¡Este camino proporcionará una base sólida para construir NexusL de manera robusta y extensible!
