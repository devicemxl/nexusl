# Semántica en Acción: Predicados Modales, Embeddings y Búsqueda Lógica Extendida en NexusL

NexusL, como lenguaje diseñado para agentes inteligentes que operan en el mundo real, necesita comprender no solo los hechos (`David is human`), sino también las **modalidades** de esos hechos y acciones: ¿`puede` David correr? ¿`debe` el robot recargarse? ¿`intentó` el agente realizar una tarea?

Esta capacidad de razonar sobre la **posibilidad, permiso, intención, necesidad y otras modalidades** es fundamental. En NexusL, logramos esto no mediante una explosión de "keywords" en el corazón del lenguaje, sino a través de un sistema unificado de **`NliEntity`s (símbolos), embeddings de significado y un motor de búsqueda y razonamiento lógico extendido.**

-----

## El Desafío de las Modalidades: Más Allá de la Coincidencia Exacta

En un sistema declarativo tradicional, si una regla busca `(robot has_capacity "heavy_lifting")`, solo encontrará esa tripleta exacta. Si el conocimiento está expresado como `(robot can_lift "heavy_objects")` o `(robot is_capable_of "lifting_heavy_things")`, la consulta fallará a menos que se hayan escrito reglas explícitas para cada sinónimo o una jerarquía ontológica manual muy detallada.

Esto presenta varios problemas:

  * **Rigidez:** Requiere un conocimiento exacto de la terminología para las consultas.
  * **Mantenimiento:** Escalar el sistema para manejar nuevas expresiones o sinónimos es engorroso.
  * **Comprensión Limitada:** El sistema no "entiende" realmente que `has_capacity` y `can_lift` son conceptos relacionados.

NexusL supera esto integrando la **semántica profunda** directamente en su proceso de búsqueda y razonamiento.

-----

## La Solución de NexusL: Patrones Semánticos y Embeddings en el Corazón Lógico

En NexusL, los "verbos modales" y las "modalidades de agente" se conceptualizan como **patrones semánticos** que se expresan mediante la combinación de `NliEntity`s (símbolos) y la estructura de las tripletas. Su significado no se basa en una coincidencia léxica exacta, sino en la **similitud semántica** capturada por los embeddings.

### 1\. `NliEntity`s como Vehículos de Significado

Cada `NliEntity` en NexusL (`robot`, `move`, `has_ability`, `intends_to`, `past`, `future`) tiene un **embedding vectorial** asociado. Estos embeddings son generados por un mini-LLM (vía ONNX Runtime) y capturan el significado contextual de la entidad en un espacio de alta dimensión. Símbolos con significados similares (ej., `move` y `relocate`) tendrán embeddings "cercanos" en este espacio.

### 2\. Patrones de Tripletas para Expresar Modalidades

Las modalidades se construyen mediante la forma en que se combinan estas `NliEntity`s en tripletas:

  * **Habilidad/Capacidad:** `[Agente] has_ability [Acción/Habilidad]`.
      * Ej: `robot has_ability navigate_rough_terrain.`
  * **Permiso:** `[Agente] has_permission [Acción/Recurso]`.
      * Ej: `David has_permission access_server.`
  * **Intención:** `[Agente] intends_to [Acción/Meta]`.
      * Ej: `agent_rover intends_to recharge_battery.`
  * **Ejecución con Temporalidad:** `[Agente] do [Acción] at_time [Momento]`.
      * Ej: `robot do deliver_package at_time "2025-07-01T15:00:00Z".`

Estos **predicados modales** (`has_ability`, `has_permission`, `intends_to`, `do`, `at_time`, etc.) son `NliEntity`s comunes. Su rol "auxiliar" se deriva de cómo los módulos de NexusL y las reglas Datalog los interpretan en estos patrones.

### 3\. La Base de Conocimiento de "Conceptos Modales"

Para potenciar el motor de búsqueda extendido, NexusL mantiene una base de conocimiento interna (posiblemente dentro de CozoDB como un conjunto especial de tripletas) de "conceptos modales" predefinidos. Cada concepto incluye:

  * Un **nombre conceptual único** (ej., `WAS_EXECUTED_CONCEPT`, `CAN_PERFORM_CONCEPT`).
  * Una **descripción en lenguaje natural** (`// Indicates past execution...`).
  * **Ejemplos de tripletas/patrones de NexusL** que expresan ese concepto.
  * El **embedding del texto descriptivo y de los ejemplos**, lo que permite la búsqueda por similitud semántica.

-----

## Búsqueda Lógica Extendida: Inteligencia en la Consulta

Esta infraestructura permite que la parte "Prolog-like" de NexusL vaya mucho más allá de las coincidencias de patrones exactos.

### 1\. Consultas Semánticas de Alto Nivel:

Cuando un agente o un LLM necesita información, no tiene que adivinar los predicados exactos. Puede preguntar usando el lenguaje natural o conceptos abstractos:

  * **Proceso:**
    1.  **Entrada Semántica:** Una consulta llega al sistema (ej., "¿Qué acciones *ha completado* el robot?").
    2.  **Resolución de Concepto Modal:** El módulo de consulta de NexusL (usando RAG y embeddings) compara la intención "ha completado" con la base de conocimiento de "conceptos modales". Encuentra la mejor coincidencia (ej., `WAS_EXECUTED_CONCEPT`).
    3.  **Generación de Patrón Datalog Dinámico:** Basado en el `WAS_EXECUTED_CONCEPT` y su patrón de tripletas asociado, el sistema construye una consulta Datalog que busca tripletas que se ajusten a ese patrón semántico. Esto puede implicar buscar múltiples predicados o combinaciones:
        ```datalog
        ; Consulta generada dinámicamente por el motor de búsqueda extendido
        QUERY ?robot ?predicado_ejecucion ?accion ?tiempo :-
            ?robot ?predicado_ejecucion ?accion,
            ?sujeto do ?accion at_time ?tiempo, ; Busca el patrón de ejecución
            (?predicado_ejecucion semantically_similar_to "EXECUTION_CONCEPT"), ; Predicado que indica ejecución (ej: 'do', 'performs_action')
            (?tiempo temporal_relation "past"). ; Predicado/Función que verifica si el tiempo es pasado
        ```
  * **Beneficio:** La consulta no busca `robot WAS_EXECUTED delivery` (que no existiría), sino que infiere que `WAS_EXECUTED` se relaciona con `do` y un contexto de tiempo pasado, recuperando `robot do deliver_package at_time "2025-07-01T15:00:00Z"`.

### 2\. Razonamiento Modal Robusto:

Las reglas Datalog pueden operar sobre esta semántica extendida. Esto permite una lógica más general y menos dependiente de los nombres exactos de los predicados.

  * **Ejemplo de Regla para "Posibilidad Operacional":**
    ```datalog
    ; Un agente 'can_operate_on' un 'recurso' si tiene la 'habilidad' necesaria Y el 'permiso' para ello.
    ?agente can_operate_on ?recurso :-
        ?agente has_ability_for ?recurso,      ; Usa un predicado de habilidad
        ?agente has_permission_to_access ?recurso. ; Usa un predicado de permiso
    ```
    Aquí, `has_ability_for` y `has_permission_to_access` son tus `NliSymbol`s. Tu motor de búsqueda extendido puede mapear consultas como "¿Es posible para el robot operar la puerta?" a esta regla, buscando por similitud semántica con `can_operate_on`.

### 3\. Tolerancia a la Variación y Extensibilidad:

  * **Variación Lingüística:** El sistema puede manejar sinónimos y frases ligeramente diferentes sin requerir reglas Datalog adicionales por cada variante. "Robot *completó* la tarea" puede mapearse al mismo concepto que "Robot *ejecutó* la tarea".
  * **Definición Dinámica:** Puedes añadir nuevos conceptos modales a tu base de conocimiento sin recompilar NexusL. Simplemente defines el nuevo patrón de tripletas, su descripción y sus ejemplos.

-----

## Implicaciones y Valor Estratégico

La integración de los "modal verbs" como un motor de búsqueda extendido en la lógica de NexusL es un pilar fundamental para:

  * **Interacción Intuitiva:** Facilita que usuarios humanos y LLMs interactúen con el sistema de manera más natural, consultando intenciones, capacidades o estados pasados sin necesidad de conocer la sintaxis interna precisa de cada tripleta.
  * **Agentes más Inteligentes:** Permite a los agentes razonar de forma más profunda sobre el mundo, entendiendo no solo "qué es" sino "qué puede ser", "qué debe ser" o "qué fue".
  * **Explicabilidad:** Las consultas semánticas y las reglas de Datalog facilitan la trazabilidad de las inferencias y decisiones del sistema.
  * **Robustez:** El sistema es más resistente a la ambigüedad y la variabilidad inherente al lenguaje y al modelado de dominios complejos.

Al hacer que la semántica sea un componente activo de la búsqueda y el razonamiento, NexusL se convierte en un sistema verdaderamente inteligente, capaz de comprender y operar con las sutilezas del significado en el mundo real.
