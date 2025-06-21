# Entity

An Entity in nexusL is any independently existing object, concept, or digital structure that can be named, referenced, or related within the system. Entities form the core symbolic units of the language's representational layer and act as the primary subjects or complements within triplets. As a symbol in the lisp world.

Entities can be physical, abstract, collective, or virtual in nature. While their manifestations vary, all entities share the property of being referable, distinguishable, and persistent or evolvable within the symbolic framework of the language.

An entity is any identifiable symbolic construct that can occupy a position in a triplet:

| structured definition |
| ----- |
|  ∀x (Entity(x) ↔ CanBeSubjectOrComplement(x))  |

In triplet terms:

| nexusL |
| ----- |
|  person1 hasName "Maria"; cityX locatedIn regionY;  |

Entities are distinguished from values (which represent data) and relations (which connect other elements), and may change over time through defined events or transitions.

Examples:

| Category | Examples |
| ----- | ----- |
| Persons | Juan, Maria |
| Places | Madrid, New York |
| **Objects** | Book, Computer |
| **Concepts** | Love, Happiness |
| **Organizations** | CompanyX, Government |
| **Places** | Madrid, New York |
| **Virtual Entities** | VideoGameAvatar, SocialMediaProfile, OnlineBankAccount |

These examples illustrate that entities may refer to real-world instances, symbolic abstractions, or purely digital constructs.

## *Subtypes*

nexusL supports semantic categorization of entities into several well-defined subtypes to aid in reasoning, inference, and structure validation.

| Physical Entity |  |
| ----- | :---- |
| Axiom | ∀e (IsPhysicalEntity(e) → Entity(e)) |
| Description | Tangible, real-world objects with physical presence |
| Example | person, location, device |
| nexusL |  |

| Abstract Entity |  |
| ----- | :---- |
| Axiom | ∀e (IsAbstractEntity(e) → Entity(e)) |
| Description | Intangible concepts or ideas  |
| Example | truth, justice |
| nexusL |  |

| Collective Entity |  |
| ----- | :---- |
| Axiom | ∀e (IsCollectiveEntity(e) → Entity(e)) |
| Description | Groupings of other entities functioning as a whole |
| Example | team, organization |
| nexusL |  |

| Virtual Entity |  |
| ----- | :---- |
| Axiom | ∀e (IsVirtualEntity(e) → Entity(e)) |
| Description | Independently modeled digital constructs  |
| Example | truth, justice |
| nexusL | webAccount, avatar |

This classification supports structured reasoning and allows constraints to be enforced by type-aware systems.

### Temporal and Dynamic Evolution

Entities in nexusL are not necessarily static. They may evolve over time through observable or inferred transitions. These changes reflect real-world dynamics and support symbolic modeling of lifecycle states.

| Type | Description | Example Scenario |
| ----- | ----- | ----- |
| Evolution | Gradual change in attributes or state | A person ages, a company grows |
| Transformation | Radical change resulting in a new identity or structure | A person migrates, a product is repurposed |
| Extinction | Permanent removal or disappearance of an entity | A person dies, a business closes |

Forms of Change

### Formal Representation of Change

Let:

* Entidad(x) denote that x is an entity.  
* Estado(e) represent a state.  
* Tiempo(t) represent a time point.

Then an evolution may be expressed as:

| structured definition |
| ----- |
|  EvolvesTo(x, s2, t2) ∧ has:State(x, s1, t1) ∧ t1 < t2  |

This expresses that entity x transitioned from state s1 at time t1 to state s2 at time t2.

| nexusL |
| ----- |
| // Temporal Relations in Triplet Syntax person1 has:State "student" at t1; person1 has:State "engineer" at t2; person1 mutate:To "engineer" at t2; |

These constructs allow for temporal querying, state-based pattern recognition, and event-driven modeling.

### Integration with Events

Events can act upon entities, thereby producing state changes, creations, or deletions.

| Event Type | Effect on Entity |
| ----- | ----- |
| Creation Event | Introduces a new entity into the system |
| Modification Event | Alters the state or attributes of the entity |
| **Deletion Event** | Removes the entity from active representation |

| structured definition |
| ----- |
|  // Event-State Linkage Axiom ∀e, x (Event(e) ∧ Affects(e, x) ∧ Entity(x) → ∃s (StateChange(x, s)))  |

This enables a tightly coupled semantic model between dynamic occurrences (events) and static constructs (entities).

| nexusL |
| ----- |
|  // Flat Syntax organizationX type Company; avatar42 is:InstanceOf VirtualEntity;  |
|  // Structured Syntax def user123 has:(   (type Person)   (has:Name "Alice")   (is:PartOf groupAlpha) );  |

This nested form enables compositional and scoped definition of complex entity structures.

Open Questions for Extension

1. Identity and Persistence  
    What mechanisms define the continuity of an entity through transformation?

2. Entity Fusion and Splitting  
    Can one entity become many, or many become one? How should this be modeled?

3. Historical Querying  
    Should the system allow querying past states or only current representations?

4. Entity Versioning  
    Should entities have version identifiers, and how should these relate to state transitions?


## Formalizacion

Para "definir" o "formalizar" conceptualmente el Entity X en el contexto de ∃X (Propiedad(X)), especialmente en un sistema "sujeto-acción-atributo/objeto", podemos desglosarlo así:

### Overview:

Un "Entity" X en este contexto representa una **referencia abstracta a una entidad individual cuya existencia se afirma, pero cuya identidad o propiedades específicas no son (aún) completamente conocidas o irrelevantes en ese punto**.

Es el "algo" genérico que sabemos que está "ahí afuera" y que posee una característica específica, sin que tengamos que saber *cuál* es ese "algo" en particular. Es un **marcador de posición para un testigo**.

* **No es un valor:** No es 5 ni "rojo".
* **No es una variable en el sentido de programación:** Aunque lo usemos como variable X en lógica, conceptualmente no es solo un *nombre* de una variable; es el *referente* al que apunta esa variable si la existencia se verifica.
* **Es un "placeholder" o "token":** Es un token que se "llena" con una entidad específica si se demuestra su existencia.

**Analogía:** Imagina que estás investigando un crimen y dices: "Existe *un* testigo que vio el coche rojo." No sabes *quién* es ese testigo todavía, pero el "un testigo" es tu Entity X. Sabes que *existe* alguien que cumple la condición de "haber visto el coche rojo", y ese "alguien" es lo que el X representa en ese momento.

### Formalización dentro de tu Sistema (Sujeto-Acción-Atributo/Objeto):

Podríamos formalizarlo definiendo un nuevo tipo de entidad o un estado para una referencia.

1.  **Definición de un Tipo/Estado EntityicReference:**

    * **Clase/Tipo:** EntityicReference
    * **Propiedad intrínseca:** isExistentiallyQuantified: TRUE (indica que su existencia ha sido afirmada por un ∃ cuantificador).
    * **Identificador (opcional):** Un nombre interno único para ese Entity (e.g., _X1, _X2) si el sistema necesita diferenciarlos internamente antes de que sean instanciados.
    * **Relación inicial:** Este EntityicReference puede tener una relación inicial con la proposición que lo introduce.

    **Ejemplo de declaración en tu sistema:**
    
    Sujeto: X
    Acción: existAs
    Objeto/Atributo: EntityicReference
    Contexto: ProposiciónExistencial(P(X)) // P(X) es la propiedad que se le atribuye
    
    Esto diría: "X existe como una referencia simbólica en el contexto de la proposición existencial P(X)."

2.  **Transición de Estado del Entity:**

    Un Entity X podría tener estados:

    * **exist:** El estado inicial cuando se introduce con ∃X. Se sabe que existe un referente, pero no su identidad concreta.
    * **embodied(Y):** Una vez que el sistema encuentra un Y concreto que satisface las propiedades del X existencial, el Entity X pasa a ser "instanciado como Y". Aquí Y sería una entidad con propiedades conocidas (e.g., Sparky, LaLlantaDelanteraDerecha).

    **Formalización de la Transición:**
    Cuando el sistema demuestra (o encuentra) que Propiedad(Y) es verdad para una entidad Y *conocida* en el dominio, entonces el EntityicReference X se "resuelve" a Y.

    
    // Pseudocódigo de transición conceptual:
    SI (PROPOSICION (∃X)(Propiedad(X)) ES_VERDADERA) ENTONCES
        // Esto implica que hay un "testigo" concreto
        ENCONTRAR un Y en el DOMINIO_DE_ENTIDADES tal que (Propiedad(Y) ES_VERDADERA)
        SI Y FUE_ENCONTRADO ENTONCES
            EL_SIMBOLO_X.estado = INSTANTIATED_AS(Y)
            EL_SIMBOLO_X.referenciaConcreta = Y
        FIN_SI
    FIN_SI
    

### En resumen:

Tu idea de X existAs Entity es una excelente manera de conceptualizar el X en una cuantificación existencial (∃X) dentro de un sistema basado en declaraciones. Formalizarlo implicaría:

* Definir un tipo o estado EntityicReference para estas entidades abstractas.
* Establecer reglas para cómo estas referencias simbólicas pueden ser "instanciadas" o "resueltas" a entidades concretas una vez que se demuestra su existencia y se identifica un ejemplo específico.

Esta aproximación es muy robusta y se alinea bien con cómo los motores de inferencia en lógica y los lenguajes como Prolog manejan la satisfacción de consultas existenciales.
¡Qué gran pregunta de seguimiento! La definición conceptual y formal de un "Entity" en este contexto es crucial. Estás tocando la base de lo que significa la "referencia" en sistemas formales y computacionales.

### Definición Conceptual de un "Entity" (X)

Conceptualizamos un "Entity" (X) en este contexto como:

**"Una unidad atómica y distintiva de referencia dentro de un sistema, que está diseñada para denotar (o 'apuntar a') una entidad individual en un dominio de discurso, sin prejuzgar la naturaleza o las propiedades de esa entidad hasta que se le asignen o se infieran."**

En otras palabras:

* **Identificador Único:** Es una etiqueta que nos permite hablar de "algo" sin ambigüedad. Si tenemos X e Y, son cosas distintas a menos que se demuestre lo contrario (X = Y).
* **Referente Potencial:** No es la cosa en sí misma, sino el medio para referirse a ella. Como un puntero en programación o un pronombre en el lenguaje natural.
* **Arbitrario y Convencional:** Su forma (X, mi_carro, id123) es arbitraria y su significado se establece por convención dentro del sistema.
* **Sustituyente de una Instancia:** Cuando decimos ∃X, estamos afirmando que existe al menos una instancia concreta en el "mundo" (o en el modelo de nuestro sistema) que puede ser denotada por este Entity X y que cumple con ciertas propiedades.
* **Sin Semántica Inherente (a priori):** Por sí solo, X no significa "carro", "persona" o "número". Su significado (su semántica) se le confiere a través de las relaciones y predicados en los que participa (car(X), esRojo(X)).

### Formalización del "Entity" (X)

En la lógica formal, específicamente en la Lógica de Primer Orden (First-Order Logic), X se formaliza como una **variable**.

Aquí te detallo cómo se formalizaría, ligándolo a tu idea de "X existAs Entity":

1.  **Alfabeto de un Lenguaje Formal:**
    Un lenguaje de primer orden se define por un **alfabeto** (Entitys) que incluye:
    * **Entitys de Variables:** Un conjunto infinito de Entitys que usaremos como "marcadores de posición" para objetos individuales. Típicamente se usan letras minúsculas al final del alfabeto: x, y, z, x₁, x₂, ...
    * **Entitys de Constantes:** Entitys para objetos individuales específicos y fijos. Ej: sparky, el_sol, 5.
    * **Entitys de Predicados:** Para propiedades y relaciones. Ej: car( ), fourWheeler( ), member( , ).
    * **Entitys de Función (opcional):** Para funciones que mapean objetos a otros objetos. Ej: padre_de( ), suma( , ).
    * **Conectivos Lógicos:** ∧ (y), ∨ (o), ¬ (no), ⇒ (implica), ⇔ (si y solo si).
    * **Cuantificadores:** ∀ (para todo), ∃ (existe).
    * **Paréntesis y Comas:** Para estructura.

2.  **Términos:**
    Los términos son lo que denotan objetos en el universo de discurso. Las variables (como X) son un tipo de término. Las constantes y las aplicaciones de funciones también son términos.

3.  **Fórmulas Atómicas:**
    Estas son las unidades más básicas de verdad/falsedad. Consisten en un Entity de predicado aplicado a uno o más términos.
    * Ej: car(X), member(w, s). Aquí, X, w, s son tus "Entitys" que denotan entidades.

4.  **Cuantificación Existencial (∃X):**
    Cuando aplicas ∃X a una fórmula (ej. ∃X (car(X))), estás formalmente declarando que:
    **"Existe al menos una asignación para la variable X (dentro del dominio de discurso de nuestro modelo) tal que la fórmula car(X) se vuelve verdadera bajo esa asignación."**

    En el pseudocódigo, X existAs Entity es la representación de que X es un elemento de ese conjunto de "Entitys de Variables" y que su existencia se afirma al encontrar una asignación que satisface la condición.

### Conexión con tu sistema "sujeto-acción-atributo/objeto":

* X sería tu **sujeto** (el Entity que referencia al algo).
* existAs sería la **acción/predicado** que afirma la naturaleza existencial de X.
* Entity sería el **atributo/objeto** que describe qué tipo de "cosa" es X a un nivel fundamental: un referente abstracto.

Esta es la forma basica de nexusL para construir un sistema que pueda manejar la indeterminación de la existencia antes de que los detalles específicos de un objeto sean conocidos.

## Conclusion

The concept of Entity in nexusL is foundational to its semantic architecture. Entities represent the symbolic building blocks through which knowledge is constructed, operations are performed, and systems are modeled. By supporting both static categorization and dynamic evolution, nexusL enables robust, temporally-aware reasoning over physical, abstract, collective, and virtual constructs.

This multifaceted representation provides the groundwork for integrating ontological modeling, reactive systems, and symbolic logic within a unified language framework.

Entity es una unidad atómica y distintiva de referencia dentro de un sistema, que está diseñada para denotar (o 'apuntar a') una entidad individual en un dominio de discurso, sin prejuzgar la naturaleza o las propiedades de esa entidad hasta que se le asignen o se infieran.
