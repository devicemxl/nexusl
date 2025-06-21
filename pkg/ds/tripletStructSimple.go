package ds

type ModalVerb struct {
	Verb modalVerbType
}

type MainVerb struct {
	Verb coreVerbType
}

type VerbStatement struct {
	ModalVerb modalVerbType
	MainVerb  func(entity *Entity) (interface{}, error) // Función asociada a la entidad
}

type SimpleObjectStatement struct {
	Statement interface{}
}

type ObjectStatement struct {
	Condition attribCondType
	Statement interface{}
}

// FlatTripletStatement represents a flat triplet such as "subject verb object;"
// Or a statement like "def subject predicate object;" or "fact subject predicate object;"
// .
// def david could:run how:fast;
// def david could:{run, eat, drive} how:fast;
type FlatTripletStatement struct {
	Scope   tripletScope // The first token of the statement (DEF, FACT, IDENTIFIER)
	Subject EntityID     // Typically an Identifier (or a more complex expression if allowed)
	Verb    VerbStatement
	Object  ObjectStatement
}

// FlatTripletSymbolDeclaration represents a statement like
// "def David is entity;"
type FlatTripletSymbolDeclaration struct {
	Scope     tripletScope
	Subject   EntityID
	ModalVerb ModalVerb
	Object    ObjectStatement
}

// FlatTripletAttributeDeclaration represents a statement like "def David has color;"
type FlatTripletAttributeDeclaration struct {
	Scope     tripletScope
	Subject   EntityID
	ModalVerb ModalVerb
	Object    SimpleObjectStatement
}

// FlatTripletAttributeValueDeclaration represents a statement like "def David color is:red;"
type FlatTripletAttributeValueDeclaration struct {
	Scope   tripletScope
	Subject EntityID
	Object  SimpleObjectStatement
	Value   ObjectStatement
}

// FlatTripletActionDeclaration represents a statement like "def David do run;"
type FlatTripletActionDeclaration struct {
	Scope     tripletScope
	Subject   EntityID
	ModalVerb ModalVerb
	MainVerb  MainVerb
}

// FlatTripletActionAttributeDeclaration represents a statement like "def David run how:fast;"
type FlatTripletActionAttributeDeclaration struct {
	Scope    tripletScope
	Subject  EntityID
	MainVerb MainVerb
	Object   ObjectStatement
}

//

/*
func main() {
	// Example usage of FlatTripletSymbolDeclaration
	symbolDeclaration := FlatTripletSymbolDeclaration{
		Scope:     "def",
		Subject:   EntityID("David"),
		ModalVerb: ModalVerb{Verb: "IS"},
		Object:    ObjectStatement{Condition: "entity"},
	}

	fmt.Printf("Symbol Declaration: %+v\n", symbolDeclaration)

	// Example usage of FlatTripletActionAttributeDeclaration
	actionAttributeDeclaration := FlatTripletActionAttributeDeclaration{
		Scope:    "def",
		Subject:  EntityID("David"),
		MainVerb: MainVerb{Verb: "run"},
		Object:   ObjectStatement{Condition: "how:fast"},
	}

	fmt.Printf("Action Attribute Declaration: %+v\n", actionAttributeDeclaration)
}
*/
