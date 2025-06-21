package main

import "fmt"

type coreVerbType string

const (
	//
	// Core Verb Actions
	//----------------------
	// These keywords are used to define the core actions or operations
	// and to define the semantic roles of expressions in the language,
	// indicating their purpose or function.
	HAS coreVerbType = "HAS" // Semantic possession or attribute binding (e.g., x HAS value)
	DO  coreVerbType = "DO"  // Imperative or procedural action (e.g., DO { ... }) - Omitida en verbosDO se convirtio en palabra clave para facilitar el rec del programa
	IS  coreVerbType = "IS"  // Semantic identity or classification (e.g., x IS type)\
)

func (s coreVerbType) coreVerb() string {
	switch s {
	//
	// Core Verb Actions
	//----------------------
	// These keywords are used to define the core actions or operations
	// and to define the semantic roles of expressions in the language,
	// indicating their purpose or function.
	case HAS:
		return "HAS" // Semantic possession or attribute binding (e.g., x HAS value)
	case DO:
		return "DO" // Imperative or procedural action (e.g., DO { ... }) - Omitida en verbosDO se convirtio en palabra clave para facilitar el rec del programa
	case IS:
		return "IS" // Semantic identity or classification (e.g., x IS type)\
	default:
		return fmt.Sprintf("Unknown Core Verb: (%s)", string(s))
	}
}
