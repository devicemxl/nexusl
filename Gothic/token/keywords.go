package token

// Mapeo de palabras clave para una búsqueda eficiente.
var Keywords = map[string]TokenClass{
	//
	// Emptiness and Uncertainty
	// ----------------------
	// Keywords representing concepts of absence, lack of information, or indeterminate states.
	"mil":     NIL,
	"maybe":   MAYBE,
	"unknown": UNKNOWN,
	//
	// ======================================================== #
	// Identifiers and Entities
	// ======================================================== #
	// These keywords define how program elements are named and referenced within nexusL.
	// They form the basis for creating and interacting with variables, functions,
	// and the core symbolic entities that represent subjects and objects in triplets.
	//
	"symbol": SYMBOL,
	//
	// ======================================================== #
	// Advanced Typing and Structure Definition
	// ======================================================== #
	// These keywords allow for the definition of custom data structures,
	// complex types, and structured relationships, enhancing type safety,
	// code clarity, and the ability to reason about the properties of entities.
	//
	// Custom Type Declarations
	// ----------------------
	// Define new composite types to group related data and behaviors.
	"type":   TYPE,
	"struct": STRUCT,
	"enum":   ENUM,
	//
	// Type Modifiers and Annotations
	// ----------------------
	// Provide additional semantic information or constraints on types, fields, or parameters.
	"read_only": READ_ONLY,
	"optional":  OPTIONAL,
	//
	// ======================================================== #
	// Core Semantic Predicates
	// ======================================================== #
	// These tokens define the fundamental semantic roles and relationships
	// within nexusL's triplet-based knowledge representation. They form the core
	// 'verbs' or 'predicates' that describe facts, states, and simple actions.
	//
	// Basic Relational Predicates
	// ----------------------
	// Core predicates establishing fundamental relationships between subjects and objects.
	"has": HAS,
	"is":  IS,
	"do":  DO,
	//
	// ======================================================== #
	// Contextual and Conditional Predicates
	// ======================================================== #
	// These tokens provide semantic context for triplets and actions,
	// specifying conditions related to manner, time, or location.
	// They enhance the expressiveness of nexusL by allowing for more nuanced
	// descriptions of events and states.
	//
	// Contextual Modifiers
	// ----------------------
	// Predicates that define the manner, temporal, or spatial conditions for an action or fact.
	"how":   HOW,
	"where": WHERE,
	"when":  WHEN,
	// ======================================================== #
	// Logical Predicates
	// ======================================================== #
	//
	// Purpose: Defines a logical rule that can be applied to infer new knowledge or relationships.
	// Context: Used to create logical implications that can be applied to derive new facts or conclusions
	//          based on existing knowledge.
	// Syntax/Example:  Una regla: "Si un robot está en una habitación y esa habitación está activa, entonces el robot tiene una ubicación en esa habitación."
	// 					RULE {(robot HAS_LOCATION ?room) IF (robot IS_IN ?room) AND (room IS_ACTIVE)}
	"rule": RULE,
	"fact": FACT,
	//
	// ======================================================== #
	// Variable and Scope Declarations
	// ======================================================== #
	// These tokens define how variables and constants are declared and
	// how their visibility and lifetime (scope) are managed within nexusL programs.
	// They are fundamental for managing data flow and preventing naming conflicts.
	//
	// Variable and Constant Declarations
	// ----------------------
	// Keywords for binding identifiers to values or expressions.
	"let":   LET,
	"var":   VAR,
	"const": CONST,
	//
	// Scope Definition
	// ----------------------
	// Keyword for explicitly defining a new lexical scope.
	"scope": SCOPE,
	//
	// ======================================================== #
	// Program Structure and Organization
	// ======================================================== #
	// These tokens define the top-level organization of nexusL code,
	// facilitating modularity, dependency management, and metadata declaration.
	// They are crucial for building complex, well-structured agent programs.
	//
	// Top-Level Program Structure
	// ----------------------
	// Keywords for defining the overall program and its main components.
	"program":  PROGRAM,
	"header":   HEADER,
	"content":  CONTENT,
	"metadata": METADATA,
	//
	// Input and Output Declarations
	// ----------------------
	// Keywords for defining external interfaces of a program.
	"input":  INPUT,
	"output": OUTPUT,
	//
	// Module and Dependency Management
	// ----------------------
	// Keywords for organizing code into reusable units and managing external dependencies.
	"module":  MODULE,
	"library": LIBRARY,
	"import":  IMPORT,
	"export":  EXPORT,
	//
	// Access Control
	// ----------------------
	// Keywords for defining visibility of entities within modules or packages.
	"private":   PRIVATE,
	"public":    PUBLIC,
	"protected": PROTECTED,
	//
	// ======================================================== #
	// Function Definitions
	// ======================================================== #
	// These tokens are dedicated to defining callable units of behavior (functions)
	// within nexusL. They specify how functions are declared, their parameters,
	// and the executable code they contain.
	//
	// Function Declaration
	// ----------------------
	// Keyword for defining a new named function.
	"func": FUNC,
	//
	// Function Components
	// ----------------------
	// Keywords for specifying parameters and the body of a function.
	"param":  PARAM,
	"code":   CODE,
	"return": RETURN,
	//
	// ======================================================== #
	// Meta-programming and Reflection
	// ======================================================== #
	// These tokens empower nexusL to introspect, manipulate, and generate its own code
	// at various stages of execution. This capability is fundamental for creating
	// Domain-Specific Languages (DSLs), automating code patterns, and enabling
	// advanced self-modifying or adaptive behaviors in intelligent agents.
	//
	// Meta-programming Constructs
	// ----------------------
	// Keywords for defining and interacting with code as data.
	"macro":   MACRO,
	"reflect": REFLECT,
	"quote":   QUOTE,
	"unquote": UNQUOTE,
	// ======================================================== #
	// State and Reactive Programming
	// ======================================================== #
	// These keywords facilitate the management of mutable state and reactive behaviors,
	// allowing nexusL agents and systems to respond dynamically to changes and handle
	// asynchronous operations gracefully. They are essential for building intelligent
	// systems that interact with dynamic environments.
	//
	// State Declaration
	// ----------------------
	// Defines a mutable or reactive property within an entity or system,
	// indicating that its value can change over time and trigger reactions.
	"state": STATE,
	// Asynchronous Operations (Promises)
	// ----------------------
	// Manages operations that complete over time, providing clear success or failure states
	// for handling non-blocking computations.
	"promise": PROMISE,
	"resolve": RESOLVE,
	"reject":  REJECT,
	"await":   AWAIT,
	// Reactive Observers and Events
	// ----------------------
	// Defines mechanisms for defining reactions to changes in state or data,
	// enabling a push-based communication model.
	"on_change": ON_CHANGE,
	"emit":      EMIT,
	// ======================================================== #
	// Flow Control
	// ======================================================== #
	// These keywords manage the execution order of operations within nexusL programs.
	// They enable conditional logic, repetitive execution (looping), function returns,
	// and robust error handling, all of which are crucial for defining complex behaviors,
	// decision-making processes, and reactive responses in intelligent agents.
	//
	// Conditional Execution
	// ----------------------
	// Tokens used to define branches in the execution path based on boolean conditions.
	"if":   IF,
	"else": ELSE,
	// Looping Constructs
	// ----------------------
	// Tokens for defining repetitive execution of code blocks, enabling iteration over data or continuous actions.
	"while":    WHILE,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	// Multi-way Branching
	// ----------------------
	// Tokens for selecting one out of multiple code blocks to execute based on the value of an expression.
	"switch":  SWITCH,
	"case":    CASE,
	"default": DEFAULT,
	// Exception Handling
	// ----------------------
	// Tokens providing mechanisms to detect, respond to, and manage runtime errors or exceptional conditions,
	// enhancing the robustness and fault-tolerance of agent behaviors.
	"try":     TRY,
	"catch":   CATCH,
	"finally": FINALLY,
	"throw":   THROW,
	// Debugging and Validation
	// ----------------------
	// Tokens that aid in verifying assumptions and debugging program behavior during development or runtime.
	"assert": ASSERT,
	// ======================================================== #
	// Logic and Constraint Programming
	// ======================================================== #
	// These keywords are essential for defining logical relationships, variable
	// domains, and constraints within nexusL. They enable declarative problem-solving,
	// symbolic reasoning, and the implementation of sophisticated planning and
	// decision-making capabilities for intelligent agents.
	//
	// Logical Connectives (Boolean Logic)
	// ----------------------
	// Combine or modify boolean expressions, returning a boolean result. Essential for complex conditions.
	"and":   AND_GATE,
	"or":    OR_GATE,
	"not":   NOT_GATE,
	"nand":  NAND_GATE,
	"imply": IMPLY_GATE,
	"xor":   XOR_GATE,
	"nor":   NOR_GATE,
	// Logical Variables and Querying
	// ----------------------
	// Represents variables whose values are determined through a logical inference engine
	// or by matching patterns in the knowledge base.
	//"find": FIND,
	// Domain Definition and Set Operations
	// ----------------------
	// Defines the set of possible values for logical variables or explicitly creates collections.
	"domain": DOMAIN,
	"infer":  INFER,
	// Constraint Declaration
	// ----------------------
	// Expresses declarative conditions that logical variables or relationships must satisfy.
	"constraint": CONSTRAINT,
	// Inference Control and Unification
	// ----------------------
	// Commands to trigger the logical inference process and attempt to make expressions equivalent.
	"solve": SOLVE,
	"unify": UNIFY,
	// Reification
	// ----------------------
	// Commands to convert logical constructs into first-class data elements (variables)
	// that can be manipulated and reasoned about within the logic system itself.
	"reify":           REIFY,
	"reify_constrain": REIFY_CONSTRAINT,
	"reify_domain":    REIFY_DOMAIN,
	// Logical Quantifiers and Operators
	// ----------------------
	// Tokens for expressing universal, existential, and implication relationships in logic.
	"for_all": FOR_ALL,
	"exist":   EXIST,
	// FACT SEE -- Logical Predicates --
	// RULE SEE -- Logical Predicates --
	"retract":     RETRACT,
	"cut":         CUT,
	"fail":        FAIL,
	"collect_all": COLLECT_ALL,
	"trace":       TRACE,
	// ======================================================== #
	// Symbolic Computation
	// ======================================================== #
	// These keywords empower nexusL to manipulate and reason about mathematical or
	// logical expressions as abstract symbols, rather than just their numerical values.
	// This opens the door for advanced AI capabilities such as automated
	// differentiation, algebraic simplification, equation solving, and logical
	// theorem proving, which are critical for complex agent reasoning and planning.
	//
	// Symbolic Expression Declaration
	// ----------------------
	// Designates a sequence of operations and variables to be treated as a symbolic expression
	// that can be analyzed and transformed, not just immediately evaluated.
	"expr": EXPR,
	// Function Invocation / Application
	//----------------------
	// Tokens for explicitly invoking or applying functions, emphasizing the first-class nature
	// of functions in nexusL (i.e., functions can be passed as data).
	//
	// Explicit Function Invocation
	"invoke": INVOKE,
	"apply":  APPLY, // Algebraic and Transformation Operations
	// ----------------------
	// Functions that perform common operations on symbolic expressions, enabling automated reasoning.
	"derive":     DERIVE,
	"simplify":   SIMPLIFY,
	"expand":     EXPAND,
	"eq_solve":   EQ_SOLVE,
	"substitute": SUBSTITUTE,
	// ======================================================== #
	// Persistence and Querying
	// ======================================================== #
	// These keywords manage the storage, retrieval, and integrity of the agent's
	// knowledge base in nexusL. They enable agents to remember facts across sessions,
	// efficiently query their understanding of the world, and ensure data consistency,
	// forming the foundation of a dynamic and intelligent memory system.
	//
	// Knowledge Base Management
	// ----------------------
	// Tokens to explicitly control the saving and loading of the knowledge base or specific data portions.
	"save": SAVE,
	"load": LOAD,
	// Querying the Knowledge Base
	// ----------------------
	// Tokens to define patterns for retrieving specific information from the knowledge base,
	// supporting declarative data access.
	"find": FIND,
	"bind": BIND,
	// Transactional Control (for Data Integrity)
	// ----------------------
	// Keywords to ensure data consistency and atomicity during complex updates to the knowledge base,
	// preventing partial or corrupted states.
	"transaction": TRANSACTION,
	"commit":      COMMIT,
	"rollback":    ROLLBACK,
	//
	// ======================================================== #
	// Concurrency and Parallelism
	// ======================================================== #
	// These keywords facilitate the execution of multiple operations simultaneously,
	// enabling agents to handle complex environments, interact with other agents,
	// and perform tasks without blocking the main execution flow.
	//
	// Asynchronous Process Execution
	// ----------------------
	// Initiates a new, independent thread of execution.
	"go": GO,
	//
	// Inter-Process Communication
	// ----------------------
	// Provides a safe mechanism for concurrent processes to send and receive data.
	"channel": CHANNEL,
	//
	// Asynchronous Operation Management
	// ----------------------
	// Define reactions to the completion or failure of concurrent tasks or batches.
	"on_complete": ON_COMPLETE,
	"on_fail":     ON_FAIL,
	"delay":       DELAY,
	// Batch Processing
	// ----------------------
	// Groups multiple concurrent operations to manage their collective completion.
	"batch": BATCH,
	// ======================================================== #
	// Relational and Flow Control Predicates
	// ======================================================== #
	// These tokens specify precise relationships, temporal sequencing,
	// or conditions that influence the flow and state changes of agent actions
	// and knowledge representation. They provide a structured way to express
	// causality, prerequisites, outcomes, and directed state transitions.
	//
	// Directional/Relational Qualifiers
	// ----------------------
	// Predicates that define relationships of origin, destination, or location,
	// applicable not only to physical points but also to states or conceptual entities.
	"from": FROM,
	"to":   TO,
	"at":   AT,
	"via":  VIA,
	// Temporal Relational Qualifiers
	// ----------------------
	// Predicates that define the temporal order or duration relative to other events or states.
	"before": BEFORE,
	"after":  AFTER,
	"during": DURING,
	// Causal, Consequential, and Behavioral Predicates
	// ----------------------
	// Predicates that express causality, required conditions, outcomes, or internal modes of operation.
	"because": BECAUSE,
	"result":  RESULT,
	"require": REQUIRE,
	"enable":  ENABLE,
	"mode":    MODE,
	// ======================================================== #
	// Funcional Predicates
	// ======================================================== #
	// These tokens enable functional programming paradigms within nexusL, allowing for
	// concise expression of behavior, higher-order functions, and flexible
	// inline definitions of executable logic, which is particularly useful for agents.
	"reduce":  REDUCE,
	"monad":   MONAD,
	"linkage": LINKAGE,
	//"builder": BUILDER,
	"lambda": LAMBDA,
	"match":  MATCH,
	"curry":  CURRY,
	// ES NECESARIO IMPLEMENTAR COMO KEYWORDS PARA LOS FACTS QUE GENERA EL LLM INTERNO
	// NO ESTAN DESTINADOS A SER INVOCADOS O CONSIMIDOS EN LA PARTE MACANICA O LOGICA
	// SI NO EN LA ETAPA DE RAZONAMIENTO BASICO -  SERAN LA PIEDRA ANGULAR DEL RAZONAMIENTO
	// COMPLEJO... EN UN FUTURO.
	/*
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
		"was_able":     WAS_ABLE,
		"is_able":      IS_ABLE,
		"will_be_able": WILL_BE_ABLE,
		//
		// Capacity and Potential Modalities
		// ----------------------
		// Indicate the potential (rather than actual) capability to develop a skill or perform an action.
		"had_capacity":       HAD_CAPACITY,
		"has_capacity":       HAS_CAPACITY,
		"will_have_capacity": WILL_HAVE_CAPACITY,
		//
		// Execution and Performance Modalities
		// ----------------------
		// Indicate the actual execution or performance status of an action.
		"was_executed":     WAS_EXECUTED,
		"is_executed":      IS_EXECUTED,
		"will_be_executed": WILL_BE_EXECUTED,
		//
		// Permission and Allowance Modalities
		// ----------------------
		// Indicate whether an action is permitted or allowed.
		"was_allowed":     WAS_ALLOWED,
		"is_allowed":      IS_ALLOWED,
		"will_be_allowed": WILL_BE_ALLOWED,
		//
		// Possibility and Likelihood Modalities
		// ----------------------
		// Express uncertainty, potential outcomes, or likelihood of an action/event.
		"might_have":  MIGHT_HAVE,
		"may":         MAY,
		"will_likely": WILL_LIKELY,
		//
		// Intention and Plan Modalities
		// ----------------------
		// Indicate a deliberate course of action or a stated plan of an agent.
		"was_intending": WAS_INTENDING,
		"is_intending":  IS_INTENDING,
		"will_intend":   WILL_INTEND,
		//
		// Necessity and Obligation Modalities
		// ----------------------
		// Express requirements, duties, or conditions that must be met.
		"had_to_have":  HAD_TO_HAVE,
		"must_have":    MUST_HAVE,
		"will_have_to": WILL_HAVE_TO,
		//
		// Suggestion Modalities
		// ----------------------
		// Express advice, recommendations, or a preferred course of action.
		"should_have": SHOULD_HAVE,
		"should":      SHOULD,
		"will_should": WILL_SHOULD,
		//
		// Expectation Modalities
		// ----------------------
		// Express anticipation or what is expected to happen.
		"was_expecting": WAS_EXPECTING,
		"is_expecting":  IS_EXPECTING,
		"will_expect":   WILL_EXPECT,
		//
		// Expectation Modalities
		// ----------------------
		// Express anticipation or what is expected to happen.
		"was_needed": WAS_NEEDED,
		"is_needed":  IS_NEEDED,
		"will_need":  WILL_NEED,
	*/
	"true":  BOOLEAN, // Usamos BOOLEAN como TokenClass para "true"/"false"
	"false": BOOLEAN, // Usamos BOOLEAN como TokenClass para "true"/"false"
}
