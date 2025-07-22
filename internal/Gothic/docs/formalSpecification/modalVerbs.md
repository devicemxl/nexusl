# Modal Verbs
-----

Cómo se traduce esto a consultas en CozoDB? Para entenderlo es clave para ver el proceso completo de vinculación.

Vamos a usar el ejemplo del `WAS_EXECUTED` y cómo lo consultaríamos, integrando la semántica y los procedimientos.

-----

```go

// WAS_EXECUTED
// definition: Indicate the past execution or performance status of an action.
// definition: carry out or put into effect (a plan, order, or course of action) in past time
// Purpose: Expresses that an action was completed in the past.
// Context: Confirms the historical execution of an action.

// WAS_EXECUTED
// NaturalLang Example: cli WAS EXECUTED at 12:00 pm
// Example:
//				 def A is:(Cli do:run WHEN:time(12:00 (HAS:meridian_block PM)));
//				 if A.run.time < Sys time now():
// 					 fact A was_executed is:true;

```

## Consultando Hechos y Modalidades en NexusL (CozoDB / Datalog)

En NexusL, cuando hablamos de "consultas en Prolog o Datalog", nos referimos a interactuar con CozoDB. CozoDB es tu motor Datalog que almacena las tripletas y te permite hacer inferencias y búsquedas.

El proceso se vería así:

### 1\. Los "Hechos" en CozoDB

Primero, tus tripletas (los hechos y las definiciones) deben estar almacenadas en CozoDB. Esto incluye:

  * **Hechos de la acción:**

    ```datalog
    ; Esto se generaría cuando se define la acción 'A' y se le asigna el 'Proc'
    :create a {id: "A", public_name: "CliAction", thing: "Action", proc_id: "proc_run_cli"}

    ; Esto se generaría cuando el sistema registra la ejecución de la acción
    :create a_execution_event {id: "exec_A_1", type: "execution_event", action_id: "A", agent_id: "Cli", time: "12:00PM"}

    ; Y, crucialmente, el hecho no ejecutable que indica que A fue ejecutada
    :create fact {subject_id: "A", predicate_id: "was_executed", object_value: true}
    ```

    *(Nota: `proc_id` sería una referencia al `Proc` real, que está en Go/Nim y no directamente en CozoDB como una tripleta ejecutable, pero la referencia existe).*

  * **Definiciones de Conceptos Modales (para el RAG y búsqueda semántica):**
    Estos serían hechos especiales en CozoDB, en una tabla dedicada a las definiciones de "conceptos".

    ```datalog
    :create concept_definition {
        concept_id: "WAS_EXECUTED_CONCEPT",
        text_definition: "Indicate the past execution or performance status of an action. Carry out or put into effect (a plan, order, or course of action) in past time.",
        purpose: "Expresses that an action was completed in the past.",
        context_info: "Confirms the historical execution of an action.",
        nl_example: "cli WAS_EXECUTED at 12:00 pm",
        code_example: "def A is:(Cli do:run WHEN:time(12:00 (HAS:meridian_block PM))); if A.run.time < Sys time now(): fact A was_executed is:true;"
    }
    ```

    Cada uno de estos campos (`text_definition`, `nl_example`, `code_example`) tendría su propio **embedding vectorial** asociado, almacenado junto con la definición en CozoDB o en un índice vectorial separado.

### 2\. El Proceso de Consulta

Ahora, imagina que quieres saber "¿Qué acciones fueron ejecutadas?" o "Muéstrame las acciones que se completaron".

**Paso 1: Consulta Inicial (Lenguaje Natural o Concepto)**

  * **Entrada del Usuario/LLM:** "Show me the actions that were completed."
  * **Módulo de Procesamiento de Lenguaje Natural (PLN) / Integración LLM:** NexusL toma esta frase y la convierte en una "intención de consulta" interna.

**Paso 2: Búsqueda Semántica Extendida (Usando RAG y Embeddings)**

  * El módulo de PLN/LLM toma la frase "actions that were completed" y genera un **embedding de consulta**.
  * Este embedding se usa para buscar en tu base de datos de `concept_definition` (usando CozoDB con capacidades de búsqueda vectorial o un índice vectorial asociado) las `concept_id` que tienen los embeddings más similares a la consulta.
  * **Resultado del RAG:** Identifica que `WAS_EXECUTED_CONCEPT` es altamente relevante. Recupera todas las tripletas asociadas a `WAS_EXECUTED_CONCEPT` (definiciones, ejemplos, etc.).

**Paso 3: Formulación Dinámica de la Consulta Datalog**

  * El módulo de consulta de NexusL, sabiendo que el usuario busca algo relacionado con `WAS_EXECUTED_CONCEPT`, y utilizando los `code_example` y `definition` recuperados por el RAG, formula la consulta Datalog adecuada.

  * En este caso, sabe que `WAS_EXECUTED_CONCEPT` se relaciona con el predicado `was_executed` en los hechos.

    ```datalog
    ?[action_name], ?[execution_status]
    :project action_name, execution_status
    :from fact
    :where
        fact.predicate_id = "was_executed" and
        fact.object_value = true and
        action_name = fact.subject_id ; Renombrar para claridad en la salida
    ```

**Paso 4: Ejecución de la Consulta en CozoDB**

  * NexusL envía esta consulta Datalog a CozoDB.
  * CozoDB busca en sus tablas de `fact` (o `a`, `a_execution_event`, etc., dependiendo del esquema exacto) y encuentra las coincidencias.
  * **Resultado de CozoDB:**
    ```
    | action_name | execution_status |
    |-------------|------------------|
    | "A"         | true             |
    | "Task_B"    | true             |
    ```

**Paso 5: Presentación del Resultado**

  * NexusL toma el resultado de CozoDB.
  * Si la consulta original fue de un LLM, puede usar este resultado estructurado para generar una respuesta en lenguaje natural: "Sí, la acción 'CliAction' (A) fue ejecutada."
  * Si fue una consulta de un desarrollador, muestra los datos directamente.

### El "Prolog-like" y la Vinculación del Hecho

La "vinculación del hecho" en CozoDB ocurre en el momento de la consulta:

1.  Cuando NexusL le pide a CozoDB que busque `fact.predicate_id = "was_executed"`, CozoDB está buscando las tripletas donde el "texto hueco" `was_executed` aparece como predicado.

2.  La "fuerza" de ese `was_executed` no viene de ser un token hardcodeado, sino del hecho de que tu sistema RAG lo ha asociado semánticamente a un concepto de alto nivel (`WAS_EXECUTED_CONCEPT`) y, más importante, que **tú has definido reglas o procedimientos en Nim/Go que reaccionan a la presencia de `was_executed` en una tripleta.**

      * Por ejemplo, tu módulo de auditoría podría tener una regla:
        ```nim
        # En tu código Nim, que se ejecuta si una condición Datalog se cumple
        if cozoDB.query("?[task_id] :from fact :where fact.predicate_id = 'was_executed' and fact.subject_id = ?task_id and ?task_id.has_property('critical')").is_some():
            log.info("Tarea crítica ejecutada: " & taskId)
            send_alert("Critical task " & taskId & " was completed.")
        ```

    Aquí, el "texto hueco" `was_executed` es el **punto de unión entre tu base de conocimiento declarativa y tus procedimientos imperativos en Nim/Go**.

Este proceso muestra cómo NexusL va más allá de un Prolog simple: no solo buscas patrones de símbolos, sino que utilizas la semántica profunda para construir esos patrones, y esos patrones activan comportamientos predefinidos y controlados en tu runtime de Nim/Go.

---

COMO SE VE EN LA PRIMER PARTE DE LA DOCUMENTACION ESTOS SON EMBEDDINGS
QUE INTERPRETA EL LLM (PEQUENO) PARA PODER GENERAR CONTEXTOS DESDE EL
CODIGO Y EJECUTAR CODIGO DE MANERA SEGURA - SEGUN LAS REGLAS COMPILADAS

```go
	//
	// ======================================================== #
	// Modal Verbs and Agent Modalities
	// ======================================================== #
	// These tokens represent modal verbs that express the mood or modality of an action
	// or state. They are crucial for modeling agent intentions, capabilities, permissions,
	// necessities, and expectations, providing a rich semantic layer for agent reasoning.
	//
	// Ability and Permission Modalities
	// ----------------------
	// Indicate whether an agent (or entity) possesses the actual skill or permission to perform an action.
	WAS_ABLE TokenClass = "WAS_ABLE" // Purpose: Expresses past ability or permission.
	// Context: "The agent had the skill/permission to perform the action."
	// Syntax/Example: robot WAS_ABLE:move heavy_object;
	IS_ABLE TokenClass = "IS_ABLE" // Purpose: Expresses present ability or permission.
	// Context: "The agent currently has the skill/permission."
	// Syntax/Example: robot IS_ABLE:navigate WHERE:complex_terrain;
	WILL_BE_ABLE TokenClass = "WILL_BE_ABLE" // Purpose: Expresses future ability or permission.
	// Context: "The agent will acquire the skill/permission in the future."
	// Syntax/Example: robot WILL_BE_ABLE:fly WHEN:upgrade_installed;
	//
	// Capacity and Potential Modalities
	// ----------------------
	// Indicate the potential (rather than actual) capability to develop a skill or perform an action.
	HAD_CAPACITY TokenClass = "HAD_CAPACITY" // Purpose: Expresses past potential or latent capability.
	// Context: "The agent possessed the inherent potential."
	// Syntax/Example: brain HAD_CAPACITY complex_calculations;
	HAS_CAPACITY TokenClass = "HAS_CAPACITY" // Purpose: Expresses present potential or latent capability.
	// Context: "The agent currently possesses the inherent potential."
	// Syntax/Example: storage_unit HAS_CAPACITY:(1000 (HAS:unit GB));
	WILL_HAVE_CAPACITY TokenClass = "WILL_HAVE_CAPACITY" // Purpose: Expresses future potential or latent capability.
	// Context: "The agent will possess the inherent potential in the future."
	// Syntax/Example: new_chip WILL_HAVE_CAPACITY AI_processing;
	//
	// Execution and Performance Modalities
	// ----------------------
	// Indicate the actual execution or performance status of an action.
	WAS_EXECUTED TokenClass = "WAS_EXECUTED" // Purpose: Expresses that an action was completed in the past.
	// Context: Confirms the historical execution of an action.
	// Syntax/Example: command WAS_EXECUTED WHEN:time(12:00 (HAS:meridian_block PM));
	IS_EXECUTED TokenClass = "IS_EXECUTED" // Purpose: Expresses that an action is currently being performed or has just completed.
	// Context: Indicates ongoing or recent completion of an action.
	// Syntax/Example: movement_sequence IS_EXECUTED WHEN:now;
	WILL_BE_EXECUTED TokenClass = "WILL_BE_EXECUTED" // Purpose: Expresses that an action will be performed in the future.
	// Context: Indicates a planned or guaranteed future execution.
	// Syntax/Example: cleanup_protocol WILL_BE_EXECUTED AFTER:event_finished;
	//
	// Permission and Allowance Modalities
	// ----------------------
	// Indicate whether an action is permitted or allowed.
	WAS_ALLOWED TokenClass = "WAS_ALLOWED" // Purpose: Expresses that an action was permitted in the past.
	// Context: Refers to past permissions or rules.
	// Syntax/Example: guest_user WAS_ALLOWED:stay WHERE:access_to_area;
	IS_ALLOWED TokenClass = "IS_ALLOWED" // Purpose: Expresses that an action is currently permitted.
	// Context: Refers to present permissions or rules.
	// Syntax/Example: robot IS_ALLOWED:move WHERE:zone_green;
	WILL_BE_ALLOWED TokenClass = "WILL_BE_ALLOWED" // Purpose: Expresses that an action will be permitted in the future.
	// Context: Refers to future permissions or rule changes.
	// Syntax/Example: software_update WILL_BE_ALLOWED AFTER:security_patch;
	//
	// Possibility and Likelihood Modalities
	// ----------------------
	// Express uncertainty, potential outcomes, or likelihood of an action/event.
	MIGHT_HAVE TokenClass = "MIGHT_HAVE" // Purpose: Expresses past possibility or potential outcome.
	// Context: "It was possible that something happened, but not certain."
	// Syntax/Example: (agent MIGHT_HAVE (detected_anomaly))
	MAY TokenClass = "MAY" // Purpose: Expresses present possibility or permission.
	// Context: "It is possible for something to happen, or permission is granted."
	// Syntax/Example: (sensor MAY (report_false_positive))
	WILL_LIKELY TokenClass = "WILL_LIKELY" // Purpose: Expresses future high probability or likelihood.
	// Context: "It is probable that something will happen."
	// Syntax/Example: (system WILL_LIKELY (experience_load_spike))
	//
	// Intention and Plan Modalities
	// ----------------------
	// Indicate a deliberate course of action or a stated plan of an agent.
	WAS_INTENDING TokenClass = "WAS_INTENDING" // Purpose: Expresses a past intention or plan.
	// Context: "The agent had a plan or aim to do something."
	// Syntax/Example: (robot WAS_INTENDING (to (recharge)))
	IS_INTENDING TokenClass = "IS_INTENDING" // Purpose: Expresses a present intention or plan.
	// Context: "The agent currently holds this plan or intention."
	// Syntax/Example: (agent IS_INTENDING (to (verify_data)))
	WILL_INTEND TokenClass = "WILL_INTEND" // Purpose: Expresses a future intention or plan.
	// Context: "The agent will form this intention or plan."
	// Syntax/Example: (new_protocol WILL_INTEND (to (optimize_energy_use)))
	//
	// Necessity and Obligation Modalities
	// ----------------------
	// Express requirements, duties, or conditions that must be met.
	HAD_TO_HAVE TokenClass = "HAD_TO_HAVE" // Purpose: Expresses past necessity or obligation.
	// Context: "It was required or obligatory for something to happen."
	// Syntax/Example: (robot HAD_TO_HAVE (completed_calibration))
	MUST_HAVE TokenClass = "MUST_HAVE" // Purpose: Expresses present necessity or strong obligation.
	// Context: "It is a current requirement or duty."
	// Syntax/Example: (system MUST_HAVE (secure_connection))
	WILL_HAVE_TO TokenClass = "WILL_HAVE_TO" // Purpose: Expresses future necessity or obligation.
	// Context: "It will be required or obligatory for something to happen."
	// Syntax/Example: (agent WILL_HAVE_TO (report_status_daily))
	//
	// Suggestion Modalities
	// ----------------------
	// Express advice, recommendations, or a preferred course of action.
	SHOULD_HAVE TokenClass = "SHOULD_HAVE" // Purpose: Expresses a past suggestion or an unfulfilled expectation.
	// Context: "It would have been advisable, or something was expected but didn't happen."
	// Syntax/Example: (robot SHOULD_HAVE (checked_sensors_first))
	SHOULD TokenClass = "SHOULD" // Purpose: Expresses a present suggestion, recommendation, or advisable action.
	// Context: "It is advisable to do this."
	// Syntax/Example: (agent SHOULD (verify_checksum))
	WILL_SHOULD TokenClass = "WILL_SHOULD" // Purpose: Expresses a future suggestion or recommendation.
	// Context: "It will be advisable to do this in the future."
	// Syntax/Example: (new_policy WILL_SHOULD (prioritize_safety))
	//
	// Expectation Modalities
	// ----------------------
	// Express anticipation or what is expected to happen.
	WAS_EXPECTING TokenClass = "WAS_EXPECTING" // Purpose: Expresses a past expectation or anticipation.
	// Context: "Something was anticipated to happen in the past."
	// Syntax/Example: (system WAS_EXPECTING (data_upload))
	IS_EXPECTING TokenClass = "IS_EXPECTING" // Purpose: Expresses a present expectation or anticipation.
	// Context: "Something is currently anticipated to happen."
	// Syntax/Example: (agent IS_EXPECTING (response_from_server))
	WILL_EXPECT TokenClass = "WILL_EXPECT" // Purpose: Expresses a future expectation or anticipation.
	// Context: "Something will be anticipated to happen in the future."
	// Syntax/Example: (monitoring_module WILL_EXPECT (hourly_reports))
	//
	// Requirement and Necessity Modalities (Alternative phrasing)
	// ----------------------
	// Indicate a strong need or prerequisite for an action or state.
	WAS_NEEDED TokenClass = "WAS_NEEDED" // Purpose: Expresses a past requirement or necessity.
	// Context: "Something was a prerequisite or indispensable."
	// Syntax/Example: (authorization WAS_NEEDED (for (access_to_data)))
	IS_NEEDED TokenClass = "IS_NEEDED" // Purpose: Expresses a present requirement or necessity.
	// Context: "Something is currently a prerequisite or indispensable."
	// Syntax/Example: (calibration IS_NEEDED (before (operation)))
	WILL_NEED TokenClass = "WILL_NEED" // Purpose: Expresses a future requirement or necessity.
	// Context: "Something will be a prerequisite or indispensable in the future."
	// Syntax/Example: (new_feature WILL_NEED (additional_resources))
	//
```