# 🧠 NexusL

A Semantic Language for Agent Orchestration

## Overview

NexusL is not a new platform meant to replace existing systems. It doesn’t aim to dominate or reinvent the programming landscape. Instead, **NexusL is a shared language**—a symbolic structure for expressing, exchanging, and reasoning about knowledge between distributed, autonomous agents.

Rather than centralizing logic, NexusL provides a **semantic bridge**, enabling systems to communicate meaningfully.

## Why Go as the Core of NexusL?

While symbolic or logic-based languages like Lisp, Nim, or Clojure are often favored for DSLs, **Go was chosen as the foundation for NexusL for strategic reasons**:

- ✅ **Simplicity and stability** — Go is readable, maintainable, and battle-tested.
- ✅ **Excellent for concurrency and communication** — Perfect for agent networks.
- ✅ **Semantic neutrality** — Go doesn't enforce a symbolic paradigm, making it ideal for implementing a language that does.
- ✅ **Built for real-world systems** — Go is designed to power reliable servers and infrastructure.

> **“Go doesn't define the meaning—it transmits it reliably. NexusL defines the meaning.”**

## Modular and Distributed by Design

NexusL ≠ Platform

NexusL ≈ Semantic Protocol

NexusL’s power lies not in syntax but in **its model for symbolic communication between agents**. These agents can be implemented in any language (Python, Rust, R, C#, etc.), but they all **share a common symbolic layer** via NexusL.

This enables a **federated, distributed architecture** where intelligence **emerges from semantic interactions**, not central control.

```text

1. A Python agent observes an event → generates a NexusL triplet.

2. NexusL interprets and updates the semantic graph.

3. A C# agent reacts to the new inference.

4. trunKV stores and versions the resulting knowledge state.

```

Many symbolic AI projects fail because they try to be all-in-one platforms.  
**NexusL does not make that mistake.**

Its goal is not to replace Python, Go, or Prolog—**but to help them speak the same symbolic language**. Each agent retains autonomy while gaining **semantic interoperability**.

## More Than the Sum of Its Parts

NexusL is inspired by the belief that **real intelligence emerges from shared context and distributed reasoning**.

Think of it as a **semantic operating system**—not one that performs specific tasks, but one that **links, orchestrates, and adds meaning** to what other systems do.

> **“NexusL is neutral: it doesn’t enforce a paradigm—it enables orchestration.”**  
> That’s its greatest strength, and its core philosophy.

```text
nexusl/
├── cmd/
│   └── nexusl/
│       └── main.go           # Punto de entrada principal para la aplicación NexusL
│   └── dbsetup/
│       └── main.go           # Punto de entrada para la herramienta dbsetup
│
├── internal/                 # Código interno del proyecto, no para ser importado por otros repos
│   ├── argo/                 # Módulo para el C Lazy Loader (Argo)
│   │   ├── loader.go         # Lógica principal del cargador
│   │   ├── memory.go         # Gestión de área de memoria para C (evitar GC, punteros)
│   │   ├── c_interface.go    # Definiciones de llamadas FFI/CGo para C (si aplica)
│   │   └── utils.go          # Funciones utilitarias específicas de Argo
│   │
│   ├── proloGo/              # Módulo para el evaluador lógico (Prolog-like)
│   │   ├── engine.go         # Lógica principal del motor de inferencia
│   │   ├── rules.go          # Manejo y almacenamiento de reglas lógicas
│   │   ├── facts.go          # Gestión de hechos lógicos (podría interactuar con trunKV)
│   │   ├── unification.go    # Implementación de unificación
│   │   ├── backtracking.go   # Lógica de backtracking
│   │   └── builtin.go        # Predicados predefinidos
│   │
│   ├── gothic/               # Módulo principal de NexusL (anteriormente "Gothic")
│   │   ├── ast/              # Abstract Syntax Tree (Árbol de Sintaxis Abstracta)
│   │   ├── lexer/            # Analizador léxico
│   │   ├── parser/           # Analizador sintáctico
│   │   ├── token/            # Definiciones de tokens
│   │   ├── runtime/          # Componentes de tiempo de ejecución (ej. VM, evaluador)
│   │   ├── sema/             # Análisis semántico (manejo de tipos, resolución de símbolos, etc.)
│   │   ├── db/               # Lógica de base de datos específica de Gothic (si no es trunKV)
│   │   └── core/             # Funciones utilitarias o componentes centrales de Gothic no encajables en otro sitio
│   │
│   ├── trunKV/               # Módulo del motor cognitivo trunKV
│   │   ├── tripletstore/     # Implementación del almacén de nodos semánticos
│   │   │   ├── storage.go    # Interfaces y lógica de persistencia (ej. para LMDB, BadgerDB, etc.)
│   │   │   ├── models.go     # Definiciones de estructuras (Triplet, Embedding, Link, CommitSnapshot)
│   │   │   └── ops.go        # Implementación de storeTriplet, updateEmbedding, buildSemanticLinks
│   │   │
│   │   ├── versioning/       # Control de versiones y gestión de ramas
│   │   │   ├── branches.go   # Funciones de rama: create_branch, switch_branch, delete_branch
│   │   │   ├── commits.go    # Lógica de commit (commitSemanticState), inspect, reset_hard
│   │   │   └── diff.go       # Implementación de diff
│   │   │
│   │   ├── merge/            # Lógica de fusión semántica
│   │   │   ├── basic.go      # merge (fusión básica)
│   │   │   └── amem.go       # semanticMerge con A-MEM
│   │   │
│   │   ├── fns/              # Evaluación de Lógica Neutrosófica Fermateana
│   │   │   ├── evaluation.go # Lógica para FnEvaluation, FnBranchEvaluation
│   │   │   └── aggregation.go# Lógica para FnBrachSet (aggregateFNS)
│   │   │
│   │   └── cognitivemap/     # Módulo para el F-NCM
│   │       └── fncm.go       # Implementación del Fermatean Neutrosophic Cognitive Map
│   │
│   └── shared/               # Utilidades compartidas por `gothic` y `trunKV` (si las hay, ej. logging, errores)
│
├── pkg/                      # Código que se pretende que sea importado por otros proyectos (APIs públicas)
│   └── nexuslapi/            # Si NexusL expone una API Go para ser usada como librería
│       └── api.go            # Funciones públicas o interfaces del sistema NexusL
│   └── trunkvapi/            # Si trunKV expone una API Go para ser usada como librería
│       └── api.go            # Funciones públicas o interfaces de trunKV
│
├── tools/                    # Herramientas de desarrollo, scripts, utilidades auxiliares
│   └── dbsetup/              # Herramienta para setup de DB (si es compleja y separada de `cmd/dbsetup`)
│       └── dbsetup.go        # Lógica de la herramienta
│
├── docs/                     # Toda la documentación del proyecto
│   ├── overview/             # Documentos de alto nivel (lo que ahora tienes en `0_goal.md` y `Architecture.md`)
│   │   └── architecture.md
│   │   └── conceptual_overview.md
│   │   └── ...
│   │
│   ├── gothic/               # Documentación específica de Gothic
│   │   ├── language_spec.md
│   │   ├── runtime.md
│   │   └── ...
│   │
│   ├── trunkv/               # Documentación específica de trunKV
│   │   ├── 0_goal.md         # Podría ser `introduction.md` o similar ahora
│   │   ├── 0a_files.md       # (Este es más bien un artefacto de análisis, no para docs finales)
│   │   ├── 0b_funciones.md   # Podría ir a `api_spec.md` o similar
│   │   ├── 1_commit.md
│   │   ├── ...
│   │   └── fn_cognitive_map.md # 11_FnCognitiveMap.md
│   │   └── api_reference.md    # Un resumen de todas las funciones como una referencia de API
│   │   └── design_decisions.md # Para justificar elecciones de diseño
│   │
│   ├── formal_spec/          # Especificaciones formales (si las hay)
│   ├── logbook/              # Registro de decisiones, avances, etc.
│   ├── penx/                 # Documentación de PENx
│   └── user_guide/           # Guías de usuario, ejemplos, tutoriales
│       ├── getting_started.md
│       ├── examples/
│       └── ...
│
├── tests/                    # Pruebas de integración, rendimiento, etc.
│   ├── e2e/
│   └── integration/
│
├── vendor/                   # Dependencias de Go (si se anaden módulos en modo vendor)
├── go.mod                    # Módulos de Go (dependencias del proyecto)
├── go.sum
└── README.md                 # README principal del proyecto
```

## What's Next?

Our focus is to build a solid foundation:  

- ~~The parser~~
- ~~The semantic triplet core~~
- The versioned store (`trunKV`)  
- Agent integration via flexible I/O

**Full interop will come later**, but the symbolic engine is already designed for stability—leveraging Go’s robust backend to its fullest.

## Conclusion

NexusL is a commitment to **connected intelligence**, **distributed meaning**, and a new kind of language:  
Not one that replaces existing tools, but one that **integrates, expands, and reasons** across them.

> We're not building perfect agents.  
> **We're building the medium for agents to understand each other.**
