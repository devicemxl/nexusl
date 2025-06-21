# Índice

1. Introducción: La Visión de NexusL
    1. Propósito y Filosofía: ¿Qué es NexusL y por qué existe?
    2. Problemas que Busca Resolver: Limitaciones de los lenguajes tradicionales.
    3. Rasgos Distintivos Clave: Homoiconicidad basada en grafos, programas como agentes.
2. Fundamentos de NexusL: El Grafo de Tripletas
    1. El Concepto de Tripletas: Sujeto, Verbo, Atributo.
    2. Nodos como Tripletas: Representación unificada de información y operaciones.
    3. La Noción de Entity: "Todo es un Entity".
        1. Tipos de Entities: Literales, referencias a tripletas, conceptos externos (CLI).
        2. Persistencia de Entities: El rol de la base de datos embebida.
    4. El Verbo: Acciones, funciones y puertas de interacción.
3. Especificación de Entidades Fundamentales de NexusL
    - Este capítulo detallará la estructura, comportamiento y semántica de las piezas constitutivas de NexusL. Para cada entidad principal (e.g., Entity, Triplet, Program/Agent), se seguirá el siguiente formato:
    1. Entidad: [Nombre de la Entidad, por ejemplo, "Entity" o "Triplet"]
        1. Propósito
            - Breve explicación de por qué existe esta entidad.
            - ¿Qué rol desempeña en el lenguaje o el runtime?
            - ¿Qué tipos de problemas o constructos soporta?
    2. Alcance y Contexto
        - ¿En qué contextos es válida o utilizada esta entidad?
        - ¿Es interna (runtime), de cara al usuario (sintaxis), o meta (por ejemplo, un tipo de reflexión del lenguaje)?
        - ¿Dónde aparece en el flujo de ejecución del lenguaje?
        1. Cross-References:
            - Usada por: [Lista de otras entidades o componentes que la utilizan, e.g., Acción, Evaluador, Agente]
            - Construida a partir de: [Lista de entidades más atómicas de las que se compone, e.g., Constante, Variable]
            - Sustituye/Extiende: - (Opcional, para indicar si reemplaza un concepto tradicional o lo mejora)
    3. Sintaxis Abstracta
        - Representación como tipo de dato algebraico, regla gramatical o estructura interna.

        ``` prolog
        Concept ::= Form1 | Form2
        Subtipos:
        ConceptLiteral ::= Entity
        ConceptPattern ::= Variable | Wildcard | ...
        ```

        - Explicación del Código: Describir la notación o el pseudo-código.
    4. Sintaxis Concreta (Opcional)
        - Si la entidad tiene una representación directa en el código que escribe el usuario.
        - Incluir ejemplos de snippets de código que ilustren su uso:

        ``` lisp
        (agentA does (light isOn true))
        (when (battery below 10%) do (agentA moveTo chargingStation))
        ```

        - Explicación del Código: Describir la lógica de los ejemplos.
    5. Modelo Semántico
        1. Rol Semántico
            - ¿Qué significa esta entidad? ¿Cómo se comporta o influye en la ejecución del programa?
            - ¿Es un valor, una función, una transformación, una regla?
        2. Reglas de Evaluación
            - Usar semántica operacional, reglas de inferencia o lenguaje natural.
            ``` lisp
            (EVAL-ACTION)
            ⟨act, σ⟩ ⇓ σ′
            where act: Acción, σ: Estado
            ```

            o

            ```prolog

            aplicar((X is Y), Estado, Estado′) :-
                Estado′ := Estado ∪ {(X, is, Y)}.
            ```

            - Explicación del Código: Detallar la notación o la regla.

        3. Comportamiento Composicional
            - ¿Cómo se compone esta entidad con otras? Secuencial, condicional, reactivo, etc.
    6. Subtipos o Especializaciones
        - Listar y describir subtipos si existen.

        | Subtipo | Descripción | Diferencias |
        | :-- | :--- | :--- |
        | Constante | Entity fijo | No puede ser unificado |
        | Variable | Marcador de posición | Unificable en runtime |
        | TripletPattern| Forma de coincidencia parcial | Puede incluir wildcards |

        - Incluir jerarquía, restricciones o tipos paramétricos si es necesario.

    7. Interacciones y Dependencias
        - Explicar cómo esta entidad interactúa con otras, por ejemplo:
        - Cómo es creada por o modifica otra (ej. Acción → Estado).
        - Cómo es consumida en la coincidencia de patrones o la inferencia (ej. Evaluador).
        - Cómo los agentes la manipulan o razonan sobre ella.
    8. Notas de Implementación
        - GO Representation: type, estructura, restricciones.

        ```go
        type Triplet struct {
            subject Entity
            predicate Entity
            obj Entity
        }
        ```

        - Prolog (like) Representation (o inspiración Datalog/CozoDB): Forma de predicado, estrategia de unificación.

        ```prolog
        triplet(subject, predicate, object)
        ```

        - Explicación del Código: - Describir la equivalencia o cómo se mapea.
        - Consideraciones de Rendimiento: - Indexabilidad, impacto en memoria.
    9. Ejemplos: Proporcionar ejemplos simples y anotados para mostrar: Creación, transformación, interacción, coincidencia de patrones.

    ```lisp

    (agentX believes (agentY wants (robot1 at room3)))
    (agentX knows (battery of robot1 below 20))
    ```

    - Explicación del Código: Describir la lógica y el resultado esperado.

    10. Extensiones o Meta-Mecánicas
    - ¿Puede esta entidad ser reflejada o manipulada a nivel meta?
    - ¿Puede el usuario definir nuevas instancias o variantes?
    - ¿Existe un sistema de macros o una fase de compilador que actúe sobre ella?
4. La Arquitectura Interna: Inspiración en "Sea of Nodes"
    1. Unificación de Flujo de Control y Dependencia de Datos.
    2. Nodos Clave para Estructuras de Control:
        1. Nodos de Región y Flujo de Control (REGION, IF, LOOP).
        2. Nodos Phi y la forma SSA (Static Single Assignment).
        3. Representación de Llamadas a Accion (CALL).
    3. Implicaciones para la Optimización y el Análisis.
5. NexusL como Sistema de Inteligencia Distribuida
    1. Programas como Agentes: Mundos Aislados y Persistentes.
        1. Base de Datos Exclusiva por Programa/Agente.
        2. El Agente como Entidad con Memoria y Estado Propio.
    2. Comunicación Inter-Agente: Flujo de Información.
        1. Paso de Mensajes de Entities y Tripletas Serializados.
        2. Analogías con Arquitecturas de Microservicios y Sistemas Reactivos.
    3. Interacción con Sistemas Externos:
        1. Verbos como Puertas de Enlace a APIs (Ej. chat_gpt, mistral, claude, gemini,... ).
        2. El CLI como Agente Inicial.
6. Persistencia y Gestión del Entorno
    1. Selección de Bases de Datos Embebidas: CozoDB (Nim) y bbolt (Go).
        1. Ventajas de cada opción para NexusL.
        2. Modelado del grafo sobre la base de datos.
    2. Ciclo de Vida de los Programas/Agentes:
        1. Inicialización y Generación de Bases de Datos.
        2. Carga y Persistencia de Estado.
    3. Gestión de Dependencias y "Librerías": Poblar la DB del Agente.
7. Interacciones y Dependencias; explicar cómo esta entidad interactúa con otras, por ejemplo:
    - Cómo es creada por o modifica otra (Acción → Estado).
    - Cómo es consumida en la coincidencia de patrones o la inferencia (Evaluador).
    - Cómo los agentes la manipulan o razonan sobre ella.
    1. Exportación a RDF para Interoperabilidad Semántica;
        - Beneficios de la exportación a RDF (integración con Triple Stores, SPARQL, herramientas de ontología).
        - Mapeo del modelo de tripletas de NexusL al modelo RDF (Sujeto, Predicado, Objeto).
        - Ejemplo de serialización de tripletas de NexusL a formatos RDF (Turtle, JSON-LD).
        - Implicaciones para el razonamiento y la gestión de conocimiento en agentes.
8. Desafíos y Oportunidades Futuras
    1. Diseñar una Sintaxis Coherente: Cómo escribir el grafo.
    2. Estrategias de Serialización y Deserialización Eficientes.
    3. Manejo de Errores y Resiliencia en un Flujo Distribuido.
    4. Desarrollo de Herramientas de Depuración y Visualización del Grafo.
    5. Potencial para Metaprogramación Avanzada y Auto-Modificación.
9. Conclusión
    1. Recapitulación de la Visión y el Potencial de NexusL.
    2. Próximos Pasos en el Desarrollo del Proyecto.
