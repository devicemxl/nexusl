# **Verbos Auxiliares y Modales**

La lista de auxiliares y modales es fundamental para capturar la riqueza del lenguaje natural en la base de conocimiento. Nos permite ir más allá del simple sujeto-verbo-objeto para incluir **tiempo** y **modalidades específicas** (permiso, posibilidad, necesidad, etc.).

El enfoque de **reificación del evento** es perfecto para esto, ya que nos permite adjuntar múltiples predicados descriptivos a una única instancia de una acción.

Para el funcionamiento asumimos  que david tiene el ID de símbolo UUID_david. Cada acción generará un nuevo EventID único (ej. uuid_accion_1, uuid_accion_2, etc.).

### **Mapeo de Matices**

**Predicados Base para cualquier Evento:**

```prolog
event(EventID).  
action_type(EventID, ActionVerb)./* (ej. run, jump)  */
agent(EventID, SubjectID).
```

**Predicados de Tiempo (Tense):**

```prolog
tense(EventID, past).  
tense(EventID, present).  
tense(EventID, future).
```

**Predicados de Modalidad (Modality):**

```prolog
modality(EventID, Type). /* 
```

donde Type puede ser:  

* Habilidad
    * could
    * can
    * be_able_to
    * maybe
* permission
    * allowed_to  
* Posibilidad
    * might
    * may  
* Necesidad/Obligación 
    * had_to
    * must
    * will_have_to
    * need_to 
    * need
    * will_need  
* Sugerencia/Recomendación
    * should_have
    * should

**Predicado de Negación:**

```prolog
negated(EventID).
```

### **Ejemplos Modelados**

#### **1. david do run; (asumiendo 'do' implica habilidad/posibilidad en presente, como 'CAN')**

Aquí, "do run" se interpreta como "David tiene la habilidad/posibilidad de correr" en presente.

Code snippet

// 1. Instancia del evento de correr de David  
event(uuid_david_do_run_1).  
action_type(uuid_david_do_run_1, run).  
agent(uuid_david_do_run_1, UUID_david).

// 2. Matices (tiempo y modalidad)  
tense(uuid_david_do_run_1, present).  
modality(uuid_david_do_run_1, can). // O could ser 'allowed_to' si es el significado de 'do'

#### **2. david had_to run;**

Implica una obligación pasada.

Code snippet

// 1. Instancia del evento de correr de David  
event(uuid_david_had_to_run_1).  
action_type(uuid_david_had_to_run_1, run).  
agent(uuid_david_had_to_run_1, UUID_david).

// 2. Matices  
tense(uuid_david_had_to_run_1, past).  
modality(uuid_david_had_to_run_1, obligatory). // O podrías usar 'had_to' directamente como valor si te gusta más específico

#### **3. david must no:run;**

Implica una obligación presente de **no** correr.

Code snippet

// 1. Instancia del evento de correr de David  
event(uuid_david_must_not_run_1).  
action_type(uuid_david_must_not_run_1, run). // La acción sigue siendo 'run'  
agent(uuid_david_must_not_run_1, UUID_david).

// 2. Matices  
tense(uuid_david_must_not_run_1, present).  
modality(uuid_david_must_not_run_1, obligatory).  
negated(uuid_david_must_not_run_1).           // Indica que el evento en sí está negado

#### **4. david willHaveTo run;**

Implica una obligación futura.

Code snippet

// 1. Instancia del evento de correr de David  
event(uuid_david_will_have_to_run_1).  
action_type(uuid_david_will_have_to_run_1, run).  
agent(uuid_david_will_have_to_run_1, UUID_david).

// 2. Matices  
tense(uuid_david_will_have_to_run_1, future).  
modality(uuid_david_will_have_to_run_1, obligatory). // O podrías usar 'will_have_to' como valor

### **Cómo lo maneja el Transpilador NexusL:**

El transpilador de NexusL sería el componente clave que, al ver una sentencia como david had_to run;, haría lo siguiente:

1. **Identificar el Sujeto:** "david" -> busca su UUID_david.  
2. **Identificar el Verbo Principal:** "run".  
3. **Identificar los Auxiliares/Modales y su Negación:** "had_to", "no:".  
4. **Generar un Nuevo EventID (UUID) en Go.**  
5. **Crear un Conjunto de Triples en Go** que representen:  
   * Triple(EventID, "type", "run_action")  
   * Triple(EventID, "agent", UUID_david)  
   * Triple(EventID, "tense", "past")  
   * Triple(EventID, "modality", "obligatory")  
   * Y si fuera no:run;, también Triple(EventID, "negated", true) (o simplemente un negated(EventID) si la representación del Triple no maneja booleanos directamente, como en Datalog).  
6. **Guardar estos Triples en bbolt.**

Al inicio de la aplicación, estos Triples se cargarán y se convertirían en los hechos Datalog que markkurossi/datalog usaría para la inferencia.

Este modelo te da una gran flexibilidad para realizar consultas complejas sobre las acciones y sus matices. Por ejemplo:

* ?- agent(X, UUID_david), modality(X, obligatory), tense(X, future). (¿Qué obligaciones futuras tiene David?)  
* ?- action_type(X, run), negated(X). (¿Qué acciones de "correr" han sido negadas?)

En sí, este es el camino para construir la KB