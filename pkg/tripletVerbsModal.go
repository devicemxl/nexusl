package main

import "fmt"

type modalVerbType string

const (
	//
	// modal verbs
	//----------------------
	// (future after prolog implementation)
	//
	// Ability or permission
	// These keywords indicate the ability or permission to perform an action.
	// They are used to express capability or potential.
	COULD      modalVerbType = "COULD"      // Past
	CAN        modalVerbType = "CAN"        // present
	BE_ABLE_TO modalVerbType = "BE_ABLE_TO" // Future
	// Permission or possibility
	// These keywords indicate the permission or possibility of an action.
	COULD_HAVE modalVerbType = "COULD_HAVE" // Past
	PERMISSION modalVerbType = "PERMISSION" // Present
	ALLOWED_TO modalVerbType = "ALLOWED_TO" // Future
	// Possibility
	// These keywords indicate the possibility of an action.
	// They are used to express uncertainty or potential outcomes.
	MIGHT modalVerbType = "MIGHT" // Past
	MAY   modalVerbType = "MAY"   // Present - future
	// Necessity or obligation
	// These keywords indicate the necessity or obligation of an action.
	// They are used to express requirements or recommendations.
	HAD_TO       modalVerbType = "HAD_TO"       // Past
	MUST         modalVerbType = "MUST"         // Present
	WILL_HAVE_TO modalVerbType = "WILL_HAVE_TO" // Future
	// Suggestion or recommendation
	// These keywords indicate a suggestion or recommendation for an action.
	// They are used to express advice or guidance.
	SHOULD_HAVE modalVerbType = "SHOULD_HAVE" // Past
	SHOULD      modalVerbType = "SHOULD"      // Present - future
	// Requirement or necessity
	// These keywords indicate a requirement or necessity for an action.
	// They are used to express obligations or needs.
	NEED_TO   modalVerbType = "NEED_TO"   // Past
	NEED      modalVerbType = "NEED"      // Present - future
	WILL_NEED modalVerbType = "WILL_NEED" //
)

func (s modalVerbType) modalVerb() string {
	switch s {
	//
	// modal verbs
	//----------------------
	// (future after prolog implementation)
	//
	// Ability or permission
	// These keywords indicate the ability or permission to perform an action.
	// They are used to express capability or potential.
	case COULD:
		return "COULD" // Past
	case CAN:
		return "CAN" // present
	case BE_ABLE_TO:
		return "BE_ABLE_TO" // Future
	// Permission or possibility
	// These keywords indicate the permission or possibility of an action.
	case COULD_HAVE:
		return "COULD_HAVE" // Past
	case PERMISSION:
		return "PERMISSION" // Present
	case ALLOWED_TO:
		return "ALLOWED_TO" // Future
	// Possibility
	// These keywords indicate the possibility of an action.
	// They are used to express uncertainty or potential outcomes.
	case MIGHT:
		return "MIGHT" // Past
	case MAY:
		return "MAY" // Present - future
	// Necessity or obligation
	// These keywords indicate the necessity or obligation of an action.
	// They are used to express requirements or recommendations.
	case HAD_TO:
		return "HAD_TO" // Past
	case MUST:
		return "MUST" // Present
	case WILL_HAVE_TO:
		return "WILL_HAVE_TO" // Future
	// Suggestion or recommendation
	// These keywords indicate a suggestion or recommendation for an action.
	// They are used to express advice or guidance.
	case SHOULD_HAVE:
		return "SHOULD_HAVE" // Past
	case SHOULD:
		return "SHOULD" // Present - future
	// Requirement or necessity
	// These keywords indicate a requirement or necessity for an action.
	// They are used to express obligations or needs.
	case NEED_TO:
		return "NEED_TO" // Past
	case NEED:
		return "NEED" // Present - future
	case WILL_NEED:
		return "WILL_NEED" //
	default:
		return fmt.Sprintf("Unknown Modal Verb: (%s)", string(s))
	}
}
