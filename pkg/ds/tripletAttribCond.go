package ds

import "fmt"

type attribCondType string

const (
	// Core attribute conditions
	HOW   attribCondType = "HOW"   // Semantic manner or method (e.g., x HOW action)
	WHEN  attribCondType = "WHEN"  // Semantic temporal condition (e.g., x WHEN condition)
	WHERE attribCondType = "WHERE" // Semantic spatial condition (e.g., x WHERE location)
)

func (s attribCondType) attribCond() string {
	switch s {
	//
	// Core attribute conditions
	case HOW:
		return "HOW" // Semantic manner or method (e.g., x HOW action)
	case WHEN:
		return "WHEN" // Semantic temporal condition (e.g., x WHEN condition)
	case WHERE:
		return "WHERE" // Semantic spatial condition (e.g., x WHERE location)
	default:
		return fmt.Sprintf("Unknown Attribute Condition: (%s)", string(s))
	}
}
