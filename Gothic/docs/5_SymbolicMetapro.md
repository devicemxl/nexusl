# Cálculo Simbólico y Metaprogramación en NexusL

NexusL no es solo un lenguaje para la manipulación de tripletas y el control de agentes; su diseño incorpora capacidades avanzadas de **Cálculo Simbólico** y **Metaprogramación**, lo que le permite a los agentes y desarrolladores razonar sobre el código y las expresiones de maneras muy potentes y flexibles. Estas características abren las puertas a comportamientos de agente más complejos, adaptables y auto-modificables.

---

## Cálculo Simbólico: Razonando con Expresiones Abstractas

El **cálculo simbólico** en NexusL permite que las expresiones matemáticas y lógicas sean manipuladas como **símbolos abstractos**, no solo como valores numéricos o booleanos inmediatos. A diferencia de un cálculo aritmético estándar donde `+ 2 3` se evalúa instantáneamente a `5`, el cálculo simbólico trata expresiones como `+ ?x * 2 ?y` como una **fórmula** ($x + 2y$) que puede ser inspeccionada, transformada o simplificada.

Para lograr esto, NexusL introduce el token **`EXPR`**. Una expresión envuelta en `EXPR ...` le indica al evaluador que debe mantener la estructura de la expresión en su Árbol de Sintaxis Abstracta (AST), permitiendo que operaciones especializadas trabajen directamente sobre esa forma.

**Tokens clave para el Cálculo Simbólico:**

* **`EXPR`**: Encapsula una expresión para que sea tratada simbólicamente.
* **`DERIVE`**: Para calcular derivadas de expresiones simbólicas.
* **`SIMPLIFY`**: Para reducir expresiones a su forma más sencilla.
* **`EXPAND`**: Para desarrollar expresiones (ej., productos, potencias).
* **`SOLVE_EQ`**: Para resolver ecuaciones simbólicamente.
* **`SUBSTITUTE`**: Para reemplazar variables o subexpresiones dentro de una expresión simbólica.

**Beneficios para NexusL:**

* **Agentes que Razonan Matemáticamente:** Permite a los agentes realizar cálculos abstractos, optimizaciones o análisis de funciones sin valores concretos.
* **Flexibilidad en la Resolución de Problemas:** Las ecuaciones y fórmulas pueden ser manipuladas y transformadas antes de la instanciación o ejecución numérica.
* **Representación de Conocimiento Rica:** Facilita la representación de leyes físicas, relaciones funcionales o modelos dinámicos como expresiones manipulables dentro del grafo de conocimiento.

---

## Metaprogramación: El Código como Dato y la Auto-Transformación

La **metaprogramación** es la capacidad de un programa para **inspeccionar, manipular o generar otro programa** (o a sí mismo) en tiempo de ejecución o en una fase de "compilación" previa a la ejecución. En NexusL, esto se inspira fuertemente en la tradición Lisp de tratar el código como **datos**, lo que permite una extensibilidad sintáctica y una abstracción sin precedentes.

Los tokens de metaprogramación le dan a NexusL la capacidad de:

1.  **Manipular código como datos:** Usar el propio lenguaje para escribir funciones que reciben código NexusL como entrada y producen código NexusL como salida.
2.  **Extender el lenguaje:** Definir nuevas construcciones sintácticas (`macros`) que se expanden a código NexusL estándar, permitiendo a los desarrolladores crear sus propios DSLs (Domain-Specific Languages) dentro de NexusL.
3.  **Introspección:** Permitir que un programa examine su propia estructura interna (`reflection`) o el estado de sus componentes en tiempo de ejecución.

**Tokens clave para la Metaprogramación y Reflexión:**

* **`MACRO`**: Para definir funciones que operan en el código de NexusL (AST) antes de la evaluación. Las macros se expanden, reemplazando su invocación por el código que generan.
* **`QUOTE`**: Un operador fundamental que le indica al evaluador que la expresión que le sigue debe ser tratada como **datos literales** (una estructura de tripletas/AST), no como código para ser evaluado inmediatamente. (Considera el azúcar sintáctico `'` para `QUOTE` por concisión: `' + 1 2`).
* **`UNQUOTE`**: El complemento de `QUOTE`. Permite que una parte de una expresión `QUOTEd` sea **evaluada** dentro de ese contexto de datos. Esto es crucial dentro de las macros para "inyectar" resultados de cálculos o variables en el código que se está generando. (Considera el azúcar sintáctico `,` para `UNQUOTE`: `,my_variable`).
* **`REFLECT`**: Para la introspección. Permite que el programa examine su propia estructura (tipos, entidades, definiciones) o el estado de sus componentes en tiempo de ejecución.

**Observación clave sobre `EXPR` vs. `QUOTE`:**

Aunque ambos se refieren a la manipulación de "expresiones", **`QUOTE`** se centra en tratar el **código** de NexusL como datos para la **metaprogramación (generación/transformación de sintaxis)**. **`EXPR`**, por otro lado, se enfoca en tratar **fórmulas matemáticas/lógicas** como datos para el **cálculo simbólico (manipulación algebraica/cálculo)**. Son herramientas complementarias para diferentes tipos de abstracción, cada una sirviendo a un propósito específico dentro de la potencia de NexusL.

---

Esta visión establece a NexusL como un lenguaje con capacidades declarativas, reactivas y de razonamiento simbólico de vanguardia, ideal para sistemas de agentes inteligentes y aplicaciones que requieren alta flexibilidad y adaptabilidad.
