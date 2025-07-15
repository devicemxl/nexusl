// . ___
// <|°_°|>
// .
package token

const (
	// ======================================================== #
	// Lexical Foundation
	// ======================================================== #
	// These tokens represent the most fundamental building blocks recognized by the lexer.
	// They include special markers for the start/end of streams, unrecognized symbols,
	// and basic formatting elements like whitespace. They are crucial for the initial
	// parsing phase of any nexusL program.
	//
	// Core Lexical Markers
	// ----------------------
	// Special tokens that denote structural boundaries or unhandled input.
	TOKEN   TokenClass = "TOKEN"
	ILLEGAL TokenClass = "ILLEGAL" // Purpose: Represents a token that is not recognized by the lexer.
	// Context: Indicates a syntax error or an unsupported character sequence.
	// Syntax/Example: Any sequence of characters that doesn't match a defined token pattern.
	//
	// Purpose: Marks the end of the input file or stream.
	// Context: Essential for the parser to know when to stop processing.
	// Syntax/Example: (Implicit at the end of every program file)
	EOF        TokenClass = "EOF"
	WHITESPACE TokenClass = "WHITESPACE" // Purpose: Represents any sequence of spaces, tabs, or newlines.
	// Context: Typically ignored by the parser but recognized by the lexer for formatting purposes.
	// Syntax/Example: " ", "\t", "\n"
	//
	// String Literal Delimiter
	// ----------------------
	// Specific token for marking the beginning of string literals.
	DOUBLE_QUOTE TokenClass = `"` // Purpose: Denotes the opening and closing delimiter for single-line string literals.
	SINGLE_QUOTE TokenClass = `'`
	BACKTICK     TokenClass = `'`
	// Context: Ensures that the characters enclosed are treated as a string value.
	// Syntax/Example: "hello", "robot_id"
	//
	// ======================================================== #
	// Core Data Types and Literals
	// ======================================================== #
	// These tokens represent the basic, atomic values directly expressible in the language.
	// They are fundamental for representing constant data, numbers, booleans, and text.
	// Unlike variables or identifiers, these are fixed values used directly in expressions or assignments.
	//
	// Primitive Numerical Types
	// ----------------------
	// Tokens for representing numerical constant values.]
	INTEGER TokenClass = "INT" // Purpose: Represents a whole number (integer literal).
	// Context: Used for discrete numerical values without fractional components.
	// Syntax/Example: 30, 100, -5
	COMPLEX TokenClass = "COMPLEX"
	FLOAT   TokenClass = "FLOAT" // Purpose: Represents a floating-point number (decimal literal).
	// Context: Used for numerical values with fractional components.
	// Syntax/Example: 3.14, 0.001, -9.87
	//
	// Primitive Textual and Boolean Types
	// ----------------------
	// Tokens for representing textual and logical constant values.
	TRUE    TokenClass = "TRUE"    // Si 'true' es una palabra clave
	FALSE   TokenClass = "FALSE"   // Si 'false' es una palabra clave
	BOOLEAN TokenClass = "BOOLEAN" // Purpose: Represents a boolean literal (true or false).
	// Context: Used for logical conditions, comparisons, and boolean operations.
	// Syntax/Example: true, false
	STRING TokenClass = "STRING" // "hello world"
	CHAR   TokenClass = "CHAR"   // 'a'
	// Context: A sequence of characters enclosed by STRING_QUOTE (`"`).
	// Syntax/Example: "hello world", "status: OK"
	MULTILINE_STRING TokenClass = "MULTILINE_STRING" // Purpose: Represents a string literal that can span multiple lines.
	// Context: Typically enclosed in triple quotes (e.g., `"""..."""`).
	// Syntax/Example: `"""This is a
	//                 multi-line string"""`
	//
	// Special Literal Values (Emptiness and Uncertainty)
	// ----------------------
	// Tokens representing concepts of absence, lack of information, or indeterminate states.
	//
	// Special class for lexer
	EMPTY TokenClass = "EMPTY"
	//
	NIL TokenClass = "NIL"
	// Purpose: Represents the absence of any object or value.
	// Context: Similar to 'null' in other languages, indicating no value is present or assigned.
	// Syntax/Example: NIL
	MAYBE TokenClass = "MAYBE"
	// Purpose: Represents a value or state that is not yet known but is determinable.
	// Context: Implies there is a possibility to determine or resolve the situation later.
	// Syntax/Example: // e.g., a sensor reading not yet available, but expected
	// 					thisKind COULD_HAVE value;
	//					sensor HAS:value IS:UNDETERMINED;
	UNKNOWN TokenClass = "UNKNOWN"
	// Purpose: Represents a value or state for which there is no knowledge.
	// Context: Suggests a complete lack of information, without an immediate expectation of resolution.
	// Syntax/Example: UNKNOWN (e.g., the exact cause of a complex system failure)
	//
	// ======================================================== #
	// Identifiers and Entities
	// ======================================================== #
	// These tokens define how program elements are named and referenced within nexusL.
	// They form the basis for creating and interacting with variables, functions,
	// and the core symbolic entities that represent subjects and objects in triplets.
	//
	// General Identifiers
	// ----------------------
	// A universal token for naming various program constructs.
	IDENTIFIER TokenClass = "IDENTIFIER" // Purpose: Represents a user-defined name for variables, functions, predicates, custom types, etc.
	// Context: Fundamental for referring to named elements in the language.
	// Syntax/Example: my_agent, calculate_path, is_active
	//
	// Core Symbolic Entity
	// ----------------------
	// The fundamental token for declaring and referencing symbolic entities in the triplet model.
	//
	// SYMBOL
	// Purpose: Represents the declaration of a symbolic entity or a reference to one.
	// Context: In nexusL's triplet model (subject predicate object), SYMBOL typically acts as the 'subject' or 'object'.
	// Syntax/Example: david is SYMBOL;
	SYMBOL TokenClass = "SYMBOL"
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
	TYPE TokenClass = "TYPE" // Purpose: Declares a new custom type or data structure.
	// Context: Used to define the blueprint for entities or complex data, specifying their properties (HAS).
	// Syntax/Example: (TYPE Robot (HAS (name STRING) (battery_level INT) (location LocationType)))
	STRUCT TokenClass = "STRUCT" // Purpose: Similar to TYPE, specifically for defining record-like data structures with named fields.
	// Context: Often interchangeable with TYPE, but can imply a simple data container without complex behaviors.
	// Syntax/Example: (STRUCT Point (x INT) (y INT))
	ENUM TokenClass = "ENUM" // Purpose: Declares an enumeration, a set of named symbolic constants.
	// Context: Useful for defining discrete states, categories, or fixed sets of options.
	// Syntax/Example: 	Point is {(x INT) (y INT)};
	// 					coord IS:STRUCT Point;
	//
	// Type Relationships and Constraints
	// ----------------------
	// Define relationships between types or enforce type-specific conditions.
	IMPLEMENT TokenClass = "IMPLEMENT" // Purpose: Indicates that a TYPE or ENTITY adheres to a specific INTERFACE.
	// Context: Used to enforce contracts, ensuring a type provides certain functions or predicates.
	// Syntax/Example: (TYPE Robot IMPLEMENTS Drivable)
	INTERFACE TokenClass = "INTERFACE" // Purpose: Defines a contract or a set of behaviors that a TYPE must provide, without specifying implementation.
	// Context: Specifies expected predicates or functions that implementing types must define.
	// Syntax/Example: (INTERFACE Drivable (HAS_FUNCTION move (TO LocationType)))
	//
	// Type Modifiers and Annotations
	// ----------------------
	// Provide additional semantic information or constraints on types, fields, or parameters.
	READ_ONLY TokenClass = "READ_ONLY" // Purpose: Marks a field or property as immutable after its initial assignment.
	// Context: Ensures data integrity by preventing subsequent modifications to a value.
	// Syntax/Example: (TYPE Robot (HAS (serial_number STRING READ_ONLY)))
	OPTIONAL TokenClass = "OPTIONAL" // Purpose: Indicates that a field or parameter might not have a value (e.g., can be NIL).
	// Context: Allows for fields that are not always present or required, preventing errors when a value is missing.
	// Syntax/Example: (TYPE User (HAS (email STRING OPTIONAL)))
	//GENERIC TokenClass = "GENERIC" // Purpose: Used to define type parameters for generic types or functions, allowing them to operate on various types.
	// Context: Enables the creation of reusable code structures (e.g., a List that can hold any type T).
	// Syntax/Example: (FUNC identity (param (value GENERIC ?T)) (RETURN value)), (TYPE List (OF GENERIC ?T))
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
	HAS TokenClass = "HAS" // Purpose: Expresses possession, attribute binding, or a component relationship.
	// Context: Used to assert that a subject has a certain property or possession.
	// Syntax/Example: (robot HAVE (battery_level 75)), (house HAVE (door_color blue))
	IS TokenClass = "IS" // Purpose: Expresses identity, classification, or a current state.
	// Context: Used to assert that a subject is of a certain type, in a specific state, or identical to an object.
	// Syntax/Example: (robot IS (type humanoid)), (light IS (state ON)), (Alice IS Bob)
	//
	// Imperative Action Predicate
	// ----------------------
	// A predicate to explicitly indicate an imperative or procedural action.
	// Note: While 'DO' was considered a core imperative, in a triplet-based system,
	// many actions might be represented as direct predicates (e.g., (robot MOVE (to kitchen))).
	// This token provides an explicit grouping for a sequence of imperative actions if needed.
	DO TokenClass = "DO" // Purpose: Initiates a sequence of imperative or procedural actions.
	// Context: Groups commands that are to be executed in order, often within an agent's behavior definition.
	// Syntax/Example: (DO { (robot move (to (room kitchen))); (robot report_status) })
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
	HOW TokenClass = "HOW" // Purpose: Specifies the manner or method by which an action is performed.
	// Context: Adds detail about *how* something is done.
	// Syntax/Example: robot DO:move TO:kitchen HOW:efficient;
	WHEN TokenClass = "WHEN" // Purpose: Specifies a temporal condition or time point for a fact or action.
	// Context: Used to assert *when* something occurs or is true.
	// Syntax/Example: robot HAS:STATUS HOW:ONLINE WHEN:TIME(10:30AM);
	WHERE TokenClass = "WHERE" // Purpose: Specifies a spatial condition or location for a fact or action.
	// Context: Used to assert *where* something is or where an action takes place.
	// Syntax/Example: robot IS:LOCATED WHERE:(kitchen IS:position(x 10 y 20));
	//
	// ======================================================== #
	// Logical Predicates
	// ======================================================== #
	//
	// Purpose: Defines a logical rule that can be applied to infer new knowledge or relationships.
	// Context: Used to create logical implications that can be applied to derive new facts or conclusions
	//          based on existing knowledge.
	// Syntax/Example:  Una regla: "Si un robot está en una habitación y esa habitación está activa, entonces el robot tiene una ubicación en esa habitación."
	// 					RULE {(robot HAS_LOCATION ?room) IF (robot IS_IN ?room) AND (room IS_ACTIVE)}
	RULE TokenClass = "RULE"
	// Purpose: Asserts a fact or relationship in the knowledge base, making it available for querying and reasoning.
	// Context: Used to declare known truths or relationships that the inference engine can use to derive further knowledge.
	// Syntax/Example: 	FACT robot HAS:location kitchen);
	// 					FACT sensor DO:detects motion;
	FACT TokenClass = "FACT"
	// RETRACT -- SEE PROLOG TOKENS
	GOAL TokenClass = "GOAL"
	//
	PLAN TokenClass = "PLAN"
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
	LET TokenClass = "LET" // Purpose: Declares a new immutable binding (similar to a constant in its scope).
	// Context: Creates a local, block-scoped binding that cannot be reassigned after initialization.
	// Syntax/Example: (LET my_constant (VALUE 10))
	VAR TokenClass = "VAR" // Purpose: Declares a new mutable variable.
	// Context: Creates a local, block-scoped variable whose value can be reassigned.
	// Syntax/Example: (VAR counter (VALUE 0)), (SET counter (+ counter 1))
	CONST TokenClass = "CONST" // Purpose: Declares a compile-time constant.
	// Context: Binds an identifier to a value that is known at "compile" or expansion time and cannot change.
	// Syntax/Example: (CONST MAX_RETRIES (VALUE 3))
	//
	// Scope Definition
	// ----------------------
	// Keyword for explicitly defining a new lexical scope.
	SCOPE TokenClass = "SCOPE" // Purpose: Defines a new lexical scope for variables and functions.
	// Context: Limits the visibility and lifetime of declared entities to within this block, promoting modularity.
	// Syntax/Example: (SCOPE my_local_block { (LET temp_val 10); (DO (process temp_val)) })
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
	PROGRAM TokenClass = "PROGRAM" // Purpose: Represents the root-level declaration of an nexusL program.
	// Context: Encapsulates the entire application, defining its inputs, outputs, and content.
	// Syntax/Example: (PROGRAM my_agent_app (INPUT ...) (HEADER ...) (CONTENT ...))
	HEADER TokenClass = "HEADER" // Purpose: Declares metadata or descriptive information for the program.
	// Context: Contains non-executable details like program name, version, author, etc.
	// Syntax/Example: (HEADER (name "RobotController") (version "1.0"))
	CONTENT TokenClass = "CONTENT" // Purpose: Designates the main logical or executable block of a program.
	// Context: Contains the primary functions, rules, and actions of the nexusL program.
	// Syntax/Example: (CONTENT { (robot IS (state IDLE)); (start_mission) })
	METADATA TokenClass = "METADATA" // Purpose: Declares additional, arbitrary metadata associated with a program or entity.
	// Context: Flexible key-value storage for extra descriptive data not covered by HEADER.
	// Syntax/Example: (METADATA (author "John Doe") (creation_date "2024-05-15"))
	//
	// Input and Output Declarations
	// ----------------------
	// Keywords for defining external interfaces of a program.
	INPUT TokenClass = "INPUT" // Purpose: Declares expected inputs to the program or a function.
	// Context: Specifies parameters or data sources that the program will consume.
	// Syntax/Example: (INPUT (sensor_data STREAM) (target_location LocationType))
	OUTPUT TokenClass = "OUTPUT" // Purpose: Declares expected outputs or results produced by the program or a function.
	// Context: Specifies the data or effects the program will generate externally.
	// Syntax/Example: (OUTPUT (robot_status STRING) (mission_report ReportType))
	//
	// Module and Dependency Management
	// ----------------------
	// Keywords for organizing code into reusable units and managing external dependencies.
	MODULE TokenClass = "MODULE" // Purpose: Declares a module, a compilation unit or namespace that encapsulates related code.
	// Context: Organizes a set of functions, facts, and rules into a reusable component, preventing naming clashes.
	// Syntax/Example: (MODULE RobotMovementLib { (FUNC move ...); (FACT has_wheels ...) })
	LIBRARY TokenClass = "LIBRARY" // Purpose: Declares a package, a collection or grouping of related modules.
	// Context: Provides a higher level of organization for larger projects, grouping modules together.
	// Syntax/Example: (PACKAGE RobotControlSystem (MODULE Navigation) (MODULE ArmControl))
	IMPORT TokenClass = "IMPORT" // Purpose: Includes external dependencies or modules into the current scope.
	// Context: Allows access to functions, types, and facts defined in other nexusL modules or external libraries.
	// Syntax/Example: (IMPORT (RobotMovementLib))
	EXPORT TokenClass = "EXPORT" // Purpose: Exports a module or package for use in other nexusL programs.
	// Context: Makes the defined module/package available for import by other nexusL programs.
	// Syntax/Example: (EXPORT_MODULE RobotMovementLib)
	//
	// Access Control
	// ----------------------
	// Keywords for defining visibility of entities within modules or packages.
	PRIVATE TokenClass = "PRIVATE" // Purpose: Declares an entity (function, variable, fact) as private to its enclosing module/package.
	// Context: Limits access to the declared entity only from within its defined scope.
	// Syntax/Example: (PRIVATE (FUNC internal_helper_func))
	PUBLIC TokenClass = "PUBLIC" // Purpose: Declares an entity as publicly accessible from outside its enclosing module/package.
	// Context: Makes the declared entity available for use by other modules that import it.
	// Syntax/Example: (PUBLIC (FUNC move_robot))
	PROTECTED TokenClass = "PROTECTED" // Purpose: Declares an entity as accessible only within its module and by derived modules.
	// Context: Allows access to the entity from subclasses or modules that extend the current module.
	// Syntax/Example: (PROTECTED (FUNC calculate_path))
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
	FUNC TokenClass = "FUNC" // Purpose: Declares a new named function or procedure.
	// Context: Defines a reusable block of code that can be invoked with arguments.
	// Syntax/Example: (FUNC my_function (PARAM (x INT)) (CODE { (PRINT x) }))
	//
	// Function Components
	// ----------------------
	// Keywords for specifying parameters and the body of a function.
	PARAM TokenClass = "PARAM" // Purpose: Delimits the declaration of parameters for a function.
	// Context: Defines the names and types of arguments a function expects.
	// Syntax/Example: (PARAM (x INT) (y FLOAT))
	CODE TokenClass = "CODE" // Purpose: Delimits the main executable body of a function.
	// Context: Contains the sequence of statements that are executed when the function is called.
	// Syntax/Example: (CODE { (SET result (+ x y)); (RETURN result) })
	RETURN TokenClass = "RETURN" // Purpose: Indicates a value to be returned from a function or a value to be exposed from a scope.
	// Context: Used within a function's CODE block to specify the return value. Can also signify a value that should be accessible outside a smaller scope.
	// Syntax/Example: (CODE { (LET temp (+ 1 2)); (EXPORT temp) })
	//
	// ======================================================== #
	// Metaprogramming and Reflection
	// ======================================================== #
	// These tokens empower nexusL to introspect, manipulate, and generate its own code
	// at various stages of execution. This capability is fundamental for creating
	// Domain-Specific Languages (DSLs), automating code patterns, and enabling
	// advanced self-modifying or adaptive behaviors in intelligent agents.
	//
	// Metaprogramming Constructs
	// ----------------------
	// Keywords for defining and interacting with code as data.
	MACRO TokenClass = "MACRO" // Purpose: Defines a macro, which is a piece of code that transforms or generates other code.
	// Context: Macros are expanded (processed) before the main evaluation phase, allowing for syntactic extension
	//          and compile-time code manipulation. They are central to creating custom DSLs in nexusL.
	// Syntax/Example: (MACRO log_debug (message) (QUOTE (PRINT "DEBUG: " (UNQUOTE message))))
	REFLECT TokenClass = "REFLECT" // Purpose: Provides introspection capabilities, allowing the program to examine its own structure or the metadata of entities at runtime.
	// Context: Used to query type definitions, entity properties, or program structure dynamically.
	//          Essential for dynamic behavior, serialization, or understanding the agent's internal state.
	// Syntax/Example: (REFLECT (GET_TYPE_INFO Robot)), (REFLECT (GET_PREDICATES my_entity))
	QUOTE TokenClass = "QUOTE" // Purpose: Prevents the immediate evaluation of an expression, treating it as a literal data structure (an S-expression).
	// Context: Crucial in metaprogramming to manipulate code as data without executing it. It allows expressions to be passed around and transformed.
	// Syntax/Example: (SET my_code (QUOTE (+ 1 2))), (MACRO_DEFINITION (QUOTE (body of code)))
	// Note: Consider supporting `'` as syntactic sugar for QUOTE, e.g., '(+ 1 2).
	UNQUOTE TokenClass = "UNQUOTE" // Purpose: Evaluates an expression within a `QUOTEd` context.
	// Context: Used inside quoted expressions (often within macros) to selectively execute a portion of the code
	//          or inject a variable's value into the generated code.
	// Syntax/Example: (MACRO_EXPANSION (QUOTE (PRINT "Value is: " (UNQUOTE my_var)))), `,value`
	// Note: Consider supporting `,` as syntactic sugar for UNQUOTE, e.g., `,my_variable.
	//
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
	STATE TokenClass = "STATE" // Purpose: Declares a reactive state variable or property.
	// Context: Used to define properties of entities that are expected to change and that other parts of the system might need to observe.
	// Syntax/Example: (robot HAS (STATE (battery_level 75))), (sensor HAS (STATE (status "active")))
	//
	// Asynchronous Operations (Promises)
	// ----------------------
	// Manages operations that complete over time, providing clear success or failure states
	// for handling non-blocking computations.
	PROMISE TokenClass = "PROMISE" // Purpose: Represents an asynchronous operation whose result will be available in the future.
	// Context: Used for tasks that don't return immediately (e.g., network requests, long computations) but will eventually resolve or reject.
	// Syntax/Example: (SET fetch_data_promise (ASYNC (GET_URL "http://api.data.com")))
	RESOLVE TokenClass = "RESOLVE" // Purpose: Signals the successful completion of a PROMISE with a resulting value.
	// Context: Used within the asynchronous operation's implementation to indicate success and provide its outcome.
	// Syntax/Example: (RESOLVE fetch_data_promise (data_object))
	REJECT TokenClass = "REJECT" // Purpose: Signals the failure of a PROMISE with an associated error.
	// Context: Used within the asynchronous operation's implementation to indicate failure and provide an error reason.
	// Syntax/Example: (REJECT fetch_data_promise (ERROR "Network unreachable"))
	AWAIT TokenClass = "AWAIT" // Purpose: Pauses the execution of the current code block until a PROMISE is RESOLVED or REJECTED.
	// Context: Used to synchronize with asynchronous operations, ensuring their completion before proceeding. Requires a procedural flow for the waiting context.
	// Syntax/Example: (SET result (AWAIT fetch_data_promise)), (IF (AWAIT (is_robot_ready)) (THEN ...))
	//
	// Reactive Observers and Events
	// ----------------------
	// Defines mechanisms for defining reactions to changes in state or data,
	// enabling a push-based communication model.
	ON_CHANGE TokenClass = "ON_CHANGE" // Purpose: Triggers a specified action or handler when a monitored state or triplet changes.
	// Context: Establishes a reactive dependency, executing a callback whenever the observed data updates.
	// Syntax/Example: (ON_CHANGE (robot battery_level) (CALL (alert_low_battery)))
	EMIT TokenClass = "EMIT" // Purpose: Explicitly broadcasts a state change or an event, triggering `ON_CHANGE` listeners.
	// Context: Used by components to signal that an observable state has been updated or an event has occurred.
	// Syntax/Example: (EMIT (robot battery_level) 15), (EMIT (sensor_event "motion_detected"))
	//
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
	IF TokenClass = "IF" // Purpose: Executes a block of code if a specified condition evaluates to 'true'.
	// Context: Fundamental for decision-making processes, allowing agents to react differently based on current states or observations.
	// Syntax/Example: (IF (robot HAS (low_battery)) { (robot recharge) })
	ELSE TokenClass = "ELSE" // Purpose: Provides an alternative block of code to execute if the preceding 'IF' condition (and any 'ELSE_IF's) evaluates to 'false'.
	// Context: Ensures a fallback action or default behavior when the primary condition is not met.
	// Syntax/Example: (IF (sensor DETECTS (obstacle)) { (robot stop) } ELSE { (robot continue_move) })
	//
	// Looping Constructs
	// ----------------------
	// Tokens for defining repetitive execution of code blocks, enabling iteration over data or continuous actions.
	WHILE TokenClass = "WHILE" // Purpose: Repeats a block of code as long as a specified condition remains 'true'.
	// Context: Suitable for continuous monitoring, repeated attempts, or processing until a state is achieved.
	// Syntax/Example: (WHILE (robot IS (state MOVING)) { (robot update_position_sensor) })
	FOR TokenClass = "FOR" // Purpose: Iterates over a sequence (e.g., a list, range, or collection), executing a block of code for each element.
	// Context: Ideal for processing collections of data, performing actions on multiple entities, or iterating a fixed number of times.
	// Syntax/Example: (FOR item IN (GET_ALL (sensors)) { (sensor_check item) })
	BREAK TokenClass = "BREAK" // Purpose: Immediately terminates the innermost enclosing loop or 'SWITCH' statement.
	// Context: Used to exit a loop prematurely when a specific condition is met, even if the loop's primary condition is still 'true'.
	// Syntax/Example: (WHILE true { (IF (robot IS (state EMERGENCY)) { (BREAK) }) })
	CONTINUE TokenClass = "CONTINUE" // Purpose: Skips the rest of the current iteration of the innermost enclosing loop and proceeds to the next iteration.
	// Context: Useful for skipping over specific elements or conditions within a loop without exiting the loop entirely.
	// Syntax/Example: (FOR task IN (pending_tasks) { (IF (task IS (state COMPLETED)) { (CONTINUE) }); (process_task task) })
	//
	// Function and Block Control
	// ----------------------
	// Tokens that manage the flow of execution specifically within functions or delimited code blocks.
	//
	// Multi-way Branching
	// ----------------------
	// Tokens for selecting one out of multiple code blocks to execute based on the value of an expression.
	SWITCH TokenClass = "SWITCH" // Purpose: Evaluates an expression and transfers control to the 'CASE' block whose value matches the expression.
	// Context: Provides a structured way to handle multiple distinct possibilities based on a single variable or expression's value.
	// Syntax/Example: (SWITCH (robot_status) { (CASE "IDLE" { ... }); (CASE "MOVING" { ... }) })
	CASE TokenClass = "CASE" // Purpose: Defines a specific branch within a 'SWITCH' statement that corresponds to a particular value of the evaluated expression.
	// Context: Each 'CASE' represents a potential outcome of the 'SWITCH' expression.
	// Syntax/Example: (SWITCH (command) { (CASE "MOVE" { (robot move) }) })
	DEFAULT TokenClass = "DEFAULT" // Purpose: Defines a fallback branch within a 'SWITCH' statement that is executed if no other 'CASE' matches the expression's value.
	// Context: Ensures that there is always a path of execution, handling unexpected or unhandled values.
	// Syntax/Example: (SWITCH (input_type) { (CASE "TEXT" { ... }) DEFAULT { (LOG "Unknown input type") } })
	//
	// Exception Handling
	// ----------------------
	// Tokens providing mechanisms to detect, respond to, and manage runtime errors or exceptional conditions,
	// enhancing the robustness and fault-tolerance of agent behaviors.
	TRY TokenClass = "TRY" // Purpose: Encloses a block of code where exceptions (runtime errors) might occur.
	// Context: Designates a section of code to be monitored for errors, allowing for graceful error recovery.
	// Syntax/Example: (TRY { (sensor_read_data); (process_data) })
	CATCH TokenClass = "CATCH" // Purpose: Specifies a block of code to execute if an exception is thrown within the corresponding 'TRY' block.
	// Context: Allows the program to handle specific types of errors, preventing crashes and implementing recovery logic.
	// Syntax/Example: (TRY { (network_request) } CATCH (error_obj) { (LOG "Request failed:" error_obj) })
	FINALLY TokenClass = "FINALLY" // Purpose: Defines a block of code that is always executed, regardless of whether an exception occurred or was handled.
	// Context: Ensures cleanup operations (e.g., closing resources) are performed in all execution paths.
	// Syntax/Example: (TRY { (open_file) } CATCH { (handle_error) } FINALLY { (close_file) })
	THROW TokenClass = "THROW" // Purpose: Explicitly raises an exception, interrupting the normal program flow and signaling an error condition.
	// Context: Used when a specific error condition is detected that prevents further meaningful execution within the current context.
	// Syntax/Example: (IF (sensor_value IS_INVALID) { (THROW (ERROR "Sensor data out of range")) })
	//
	// Debugging and Validation
	// ----------------------
	// Tokens that aid in verifying assumptions and debugging program behavior during development or runtime.
	ASSERT TokenClass = "ASSERT" // Purpose: Checks if a given condition is true; if false, it typically halts execution or signals a critical error.
	// Context: Used for defensive programming, ensuring that assumptions about the program's state or data are met. Crucial during development for early error detection.
	// Syntax/Example: (ASSERT (> battery_level 10) "Battery is too low to proceed")
	//
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
	AND_GATE TokenClass = "AND" // Purpose: Logical AND. Returns `true` if all connected conditions are `true`.
	// Context: Used to ensure multiple conditions are met simultaneously.
	// Syntax/Example: (IF (&& (> x 0) (< x 10)) ...)
	//
	// Purpose: Logical OR. Returns `true` if at least one connected condition is `true`.
	// Context: Used when any one of several conditions is sufficient.
	// Syntax/Example: (IF (|| (= status "ALERT") (= status "CRITICAL")) ...)
	OR_GATE TokenClass = "OR"
	// Purpose: Logical NOT. Inverts the boolean value of its operand.
	// Context: Used to express the negation of a condition.
	// Syntax/Example: (IF (! (IS_ACTIVE sensor)) ...)
	NOT_GATE TokenClass = "NOT"
	// Purpose: produces an output which is false only if all its inputs are true
	NAND_GATE TokenClass = "NAND"
	// Purpose: Logical implication. Expresses a "if P then Q" relationship.
	// Context: Forms the core of rules and conditional reasoning.
	// Syntax/Example: (low_battery ?Robot) IMPLIES (needs_recharge ?Robot);
	// Aunque la tabla de verdad es la misma, la implicación lógica en sí misma
	// es una relación declarativa entre dos proposiciones (un hecho implica
	// otro hecho), mientras que un if/then en programación es una estructura
	// de control de flujo imperativa que dicta cómo el programa debe actuar
	// basándose en una condición.
	//
	// La equivalencia entre IMPLY (implicación material) y un if/then es correcta en términos
	// de su comportamiento lógico. Para nLi, esto te brinda una poderosa herramienta para:
	//
	// Definir reglas de inferencia sobre tu base de conocimientos (tripletas de estado).
	// Especificar acciones que deben realizarse cuando se cumplen ciertas condiciones, manteniendo
	// una sintaxis coherente basada en tripletas.
	IMPLY_GATE TokenClass = "IMPLY"
	XOR_GATE   TokenClass = "XOR"
	NOR_GATE   TokenClass = "NOR"
	//
	// Logical Variables and Querying
	// ----------------------
	// Represents variables whose values are determined through a logical inference engine
	// or by matching patterns in the knowledge base.
	INFER TokenClass = "INFER" // Purpose: Defines a logical variable (often prefixed with '?') or initiates a query against the knowledge base to find matching facts.
	// Context: Used to represent placeholders in patterns that the inference engine will attempt to bind to concrete values. Also triggers a search.
	// Syntax/Example: (QUERY (robot location ?L)), (FOR_ALL (?X) (robot HAS (color ?X)))
	// Note: While 'VAR_LOGIC' could be a keyword, 'QUERY' is more idiomatic for declarative logic programming (e.g., Prolog).
	//
	// Domain Definition and Set Operations
	// ----------------------
	// Defines the set of possible values for logical variables or explicitly creates collections.
	DOMAIN TokenClass = "DOMAIN" // Purpose: Declares the initial set of possible values for one or more logical variables.
	// Context: Narrows the search space for the inference engine, optimizing constraint satisfaction.
	// Syntax/Example: (DOMAIN (?Color (red green blue))), (DOMAIN (?Age (RANGE 0 120)))
	// SET TokenClass = "SET" // Purpose: Used to explicitly define a collection of elements.
	// Context: Can be used within DOMAIN declarations or for general data manipulation to create unordered collections.
	// Syntax/Example: (SET my_colors (red green blue)), (DOMAIN (?Status (SET active paused error)))
	//
	// Constraint Declaration
	// ----------------------
	// Expresses declarative conditions that logical variables or relationships must satisfy.
	CONSTRAINT TokenClass = "CONSTRAINT" // Purpose: Defines a declarative rule or condition that must hold true for a solution to be valid.
	// Context: Limits the possible assignments of logical variables. The inference engine works to find assignments that satisfy all defined constraints.
	// Syntax/Example: (CONSTRAINT (EQUALITY (Alice.horario) (Bob.horario))), (CONSTRAINT (NOT (= ?X ?Y)))
	//
	// Inference Control and Unification
	// ----------------------
	// Commands to trigger the logical inference process and attempt to make expressions equivalent.
	SOLVE TokenClass = "SOLVE" // Purpose: Initiates the search for solutions that satisfy all declared constraints, domains, and facts in the knowledge base.
	// Context: Triggers the constraint satisfaction problem (CSP) solver or logical inference engine to find valid bindings for logical variables.
	// Syntax/Example: (SOLVE (QUERY (robot location ?L)))
	UNIFY TokenClass = "UNIFY" // Purpose: Attempts to unify two logical expressions or variables, making them equivalent by finding consistent substitutions for their variables.
	// Context: A core operation in logic programming, used to match patterns and bind variables. If unification succeeds, variables become consistent.
	// Syntax/Example: (UNIFY (?X 10) (?Y 10)) // Binds ?X to ?Y (or both to 10 if one is unbound)
	//
	// Reification
	// ----------------------
	// Commands to convert logical constructs into first-class data elements (variables)
	// that can be manipulated and reasoned about within the logic system itself.
	REIFY TokenClass = "REIFY" // Purpose: Converts a logical condition or statement into a variable, allowing it to be treated as a data point in the logic system.
	// Context: Enables meta-reasoning, where the system can reason about its own logical statements.
	// Syntax/Example: (REIFY (EXISTS (?X) (robot HAS (color ?X)))) // Converts this existence claim into a variable
	REIFY_CONSTRAINT TokenClass = "REIFY_CONSTRAINT" // Purpose: Converts a defined constraint into a variable or data structure.
	// Context: Allows constraints themselves to be dynamically added, removed, or reasoned about, rather than being static rules.
	// Syntax/Example: (SET my_constraint_var (REIFY_CONSTRAINT (> ?Age 18)))
	REIFY_DOMAIN TokenClass = "REIFY_DOMAIN" // Purpose: Converts a defined domain into a variable or data structure.
	// Context: Enables dynamic modification or reasoning about the possible value sets of variables.
	// Syntax/Example: (SET valid_colors (REIFY_DOMAIN (?C (red green blue))))
	//
	// Logical Quantifiers and Operators
	// ----------------------
	// Tokens for expressing universal, existential, and implication relationships in logic.
	FOR_ALL TokenClass = "FOR_ALL" // Purpose: Universal quantifier. Asserts that a condition holds true for every element in a specified domain or for all possible bindings of a variable.
	// Context: Used to define universally true statements or rules.
	// Syntax/Example: (FOR_ALL (?X) (IMPLIES (human ?X) (mortal ?X)))
	EXIST TokenClass = "EXIST" // Purpose: Existential quantifier. Asserts that there is at least one element for which a condition holds true.
	// Context: Used to state the existence of something without necessarily identifying it.
	// Syntax/Example: (EXISTS (?R) (robot HAS (active_task ?R)))
	//
	//
	// FACT SEE -- Logical Predicates --
	//
	// RULE SEE -- Logical Predicates --
	//
	//
	RETRACT TokenClass = "RETRACT" // Purpose: Removes a fact or rule from the knowledge base, preventing it from being used in future reasoning.
	// Context: Used to retract previously asserted facts or rules, allowing for dynamic updates to the knowledge base.
	// Syntax/Example: 	RETRACT robot HAS:location kitchen;
	// 					RETRACT sensor DO:detect motion;
	CUT TokenClass = "CUT" // Purpose: Un predicado que, cuando se satisface, impide que el motor de inferencia busque soluciones alternativas para las metas a su izquierda en la cláusula actual, y también impide que el motor intente otras cláusulas de la misma regla.
	// Context: Used to optimize reasoning by preventing backtracking once a certain condition is met, effectively pruning the search space.
	// Syntax/Example: 	Buscar la primera puerta disponible y no otras
	// 					RULE {find_door ?D IF (door ?D IS_AVAILABLE) AND CUT AND (enter_door ?D) }
	FAIL TokenClass = "FAIL" // Purpose: Un predicado que siempre falla, forzando al motor de inferencia a retroceder (backtrack) y buscar una solución alternativa.
	// Context: Used to indicate that a certain condition cannot be satisfied, prompting the inference engine to backtrack and try other possibilities.
	// Syntax/Example: 	Si una condición no se cumple, forzar el fallo de la regla actual
	//					validate_data IF (data_is_corrupt) AND FAIL;
	COLLECT_ALL TokenClass = "COLLECT_ALL" // Purpose: Encuentra todas las soluciones para un patrón de consulta dado y las recolecta en una lista o colección especificada.
	// Context: A diferencia de QUERY que puede encontrar la primera solución y permitir backtracking para otras, COLLECT_ALL está diseñado para obtener todas las posibles coincidencias para una consulta, útil para análisis completos o generación de informes.
	// Syntax/Example: 	Recolectar todos los sensores que están online en una lista
	// 					COLLECT_ALL ?online_sensors WHERE (sensor ?online_sensors IS_ONLINE);
	TRACE TokenClass = "TRACE" // Purpose: Activa el modo de trazado para la ejecución de reglas lógicas o consultas, mostrando los pasos del motor de inferencia.
	// Context: Herramienta de depuración esencial para entender cómo el motor de inferencia satisface las metas, realiza unificaciones y maneja el backtracking.
	// Syntax/Example:	TRACE IS on; // Activa el trazado
	// 					Habilitar el trazado para una consulta específica
	// 					TRACE (QUERY (robot HAS (location ?where)));
	// 					Activar el trazado para todas las reglas dentro de un bloque
	// 					RULE:{DO (TRACE ON) AND (run_diagnosis_rules) AND (TRACE OFF)};
	//
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
	EXPR TokenClass = "EXPR" // Purpose: Encapsulates an expression to be treated as a symbolic entity for manipulation, rather than immediate evaluation.
	// Context: Essential for symbolic algebra, calculus, and logical reasoning where the structure of the expression itself is the subject of computation.
	// Syntax/Example: (EXPR (+ ?x (* 2 ?y))), (EXPR (AND (> ?a 0) (< ?b 10)))
	//
	// Function Invocation / Application
	//----------------------
	// Tokens for explicitly invoking or applying functions, emphasizing the first-class nature
	// of functions in nexusL (i.e., functions can be passed as data).
	//
	// Explicit Function Invocation
	INVOKE TokenClass = "INVOKE" // Purpose: Explicitly calls a function with a given set of arguments.
	// Context: Used when the function to be called is known directly. It's a clear, direct way to execute a named function.
	// Syntax/Example: (INVOKE my_function arg1 arg2), (INVOKE (robot move) (to (room kitchen)))
	//
	// Function Application (Lisp-style)
	APPLY TokenClass = "APPLY" // Purpose: Applies a function (which is itself a value, often passed as an argument) to a list of arguments.
	// Context: Powerful for higher-order functions and dynamic dispatch, allowing the function to be determined at runtime.
	// Syntax/Example: (APPLY my_func_var (LIST arg1 arg2)), (APPLY (GET_STRATEGY current_state) (data_input))
	//
	// Algebraic and Transformation Operations
	// ----------------------
	// Functions that perform common operations on symbolic expressions, enabling automated reasoning.
	DERIVE TokenClass = "DERIVE" // Purpose: Computes the derivative of a symbolic expression with respect to a specified variable.
	// Context: Fundamental for calculus-based reasoning, optimization, and modeling dynamic systems.
	// Syntax/Example: (DERIVE (EXPR (* ?x ?x)) ?x) // Result: (EXPR (* 2 ?x))
	SIMPLIFY TokenClass = "SIMPLIFY" // Purpose: Simplifies a symbolic expression into a more concise or canonical form (e.g., combining like terms, reducing fractions).
	// Context: Reduces complexity of expressions, making them easier to understand or evaluate.
	// Syntax/Example: (SIMPLIFY (EXPR (+ ?x (- ?y ?x)))) // Result: (EXPR ?y)
	EXPAND TokenClass = "EXPAND" // Purpose: Expands a symbolic expression, often by applying distributive laws or algebraic identities (e.g., polynomial expansion).
	// Context: Converts expressions into a more detailed or summed-out form.
	// Syntax/Example: (EXPAND (EXPR (* (+ ?a ?b) ?c))) // Result: (EXPR (+ (* ?a ?c) (* ?b ?c)))
	EQ_SOLVE TokenClass = "EQ_SOLVE" // Purpose: Solves a symbolic equation for a specified variable.
	// Context: Finds the values of variables that satisfy an equation or system of equations. Crucial for constraint satisfaction and planning.
	// Syntax/Example: (SOLVE_EQ (EQUAL (EXPR (+ ?x 5)) (EXPR 10)) ?x) // Result: (SOLUTION (= ?x 5))
	SUBSTITUTE TokenClass = "SUBSTITUTE" // Purpose: Replaces all occurrences of a specific variable or sub-expression within a symbolic expression with a new value or expression.
	// Context: Used for evaluating expressions under specific conditions, specializing general formulas, or transforming expressions.
	// Syntax/Example: (SUBSTITUTE (EXPR (+ ?x ?y)) ?x 10) // Result: (EXPR (+ 10 ?y))
	//
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
	SAVE TokenClass = "SAVE" // Purpose: Persists the current state of the knowledge base, or specified portions thereof, to a persistent storage medium.
	// Context: Essential for agents to retain learned facts and states across different operational sessions or system restarts.
	// Syntax/Example: (SAVE (ALL)) // Saves the entire knowledge base.
	//                 (SAVE (FACTS (robot location))) // Saves only facts related to robot location.
	LOAD TokenClass = "LOAD" // Purpose: Loads a knowledge base or specific data from storage, merging it with the current state of the in-memory knowledge base.
	// Context: Allows agents to restore previous states or load predefined knowledge modules.
	// Syntax/Example: (LOAD "my_agent_knowledge.nlk") // Loads from a file.
	//                 (LOAD (MODULE "robot_capabilities")) // Loads a specific module.
	//
	// Querying the Knowledge Base
	// ----------------------
	// Tokens to define patterns for retrieving specific information from the knowledge base,
	// supporting declarative data access.
	FIND TokenClass = "FIND" // Purpose: Searches the knowledge base for triplets that match a specified pattern, binding logical variables.
	// Context: The primary mechanism for asking questions to the knowledge base and retrieving relevant facts.
	// Syntax/Example: (QUERY (robot HAS (color ?C))), (QUERY (robot location ?L) (robot IS (type humanoid)))
	// Note: If `QUERY` is used for both declaring logical variables (as in the "Logic and Constraint Programming" block) and for the search operation itself, ensure the parser can differentiate contextually. In a Lisp-like context, `(QUERY (pattern))` usually implies the search.
	BIND TokenClass = "BIND" // Purpose: A specialized form of `QUERY`, often implying unification against a pattern to extract specific bindings.
	// Context: Can be used as a more explicit keyword for pattern matching operations, potentially providing specific binding results or a single, first match. It emphasizes the extraction of variable bindings.
	// Syntax/Example: (MATCH (robot location ?L)) // Binds `?L` to the robot's location if a match is found.
	// Note: Consider if `MATCH` is truly distinct from `QUERY`'s behavior. If `QUERY`'s results implicitly provide bindings, `MATCH` might be redundant or could be a syntactic sugar for a specific `QUERY` mode (e.g., returning only the first match).
	//
	// Transactional Control (for Data Integrity)
	// ----------------------
	// Keywords to ensure data consistency and atomicity during complex updates to the knowledge base,
	// preventing partial or corrupted states.
	TRANSACTION TokenClass = "TRANSACTION" // Purpose: Initiates a block of operations that should be treated as an atomic unit.
	// Context: All changes made within a `TRANSACTION` block are provisional and are not permanently applied to the knowledge base until a `COMMIT` command is issued. If an error occurs or `ROLLBACK` is called, all provisional changes are discarded.
	// Syntax/Example: (TRANSACTION { (ADD_FACT (user Alice online_status true)); (REMOVE_FACT (user Bob online_status false)) })
	COMMIT TokenClass = "COMMIT" // Purpose: Finalizes and applies all provisional changes made within the current `TRANSACTION` block.
	// Context: Makes the changes permanent in the knowledge base, ensuring atomicity (all or nothing).
	// Syntax/Example: (COMMIT) // Must be called after a TRANSACTION block.
	ROLLBACK TokenClass = "ROLLBACK" // Purpose: Discards all changes made within the current `TRANSACTION` block, reverting the knowledge base to its state before the transaction began.
	// Context: Used to undo changes if an error occurs, a condition is not met, or a user explicitly cancels the operation.
	// Syntax/Example: (ROLLBACK) // Must be called within or after a TRANSACTION block if changes are to be discarded.
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
	GO TokenClass = "GO" // Spawns a new concurrent process (e.g., GO (robot scan_area)).
	//
	// Inter-Process Communication
	// ----------------------
	// Provides a safe mechanism for concurrent processes to send and receive data.
	CHANNEL TokenClass = "CHANNEL" // Represents a communication channel for data exchange between concurrent processes.
	//
	// Asynchronous Operation Management
	// ----------------------
	// Define reactions to the completion or failure of concurrent tasks or batches.
	ON_COMPLETE TokenClass = "ON_COMPLETE" // Defines a handler to execute when an asynchronous operation or batch finishes successfully.
	// Syntax: (ON_COMPLETE (some_task_id) (DO ...))
	ON_FAIL TokenClass = "ON_FAIL" // Defines a handler to execute if an asynchronous operation or batch encounters an error.
	// Syntax: (ON_FAIL (some_task_id) (HANDLE_ERROR ...))
	DELAY TokenClass = "DELAY" // Pauses execution until all specified asynchronous operations or tasks have completed.
	// Syntax: (DELAY (LIST task1_id task2_id))
	//
	// Batch Processing
	// ----------------------
	// Groups multiple concurrent operations to manage their collective completion.
	BATCH TokenClass = "BATCH" // Groups a set of concurrent tasks; its completion signifies all tasks within are done.
	// Syntax: (BATCH (GO (task1)) (GO (task2)))
	//
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
	FROM TokenClass = "FROM" // Purpose: Specifies the origin or source of an action, state, or transition.
	// Context: Used to indicate "from where" or "from what state" an action or change originates.
	// Syntax/Example: (robot MOVE FROM (room kitchen)), (transition_state FROM (state idle) TO (state active))
	TO TokenClass = "TO" // Purpose: Specifies the destination or target of an action, state, or transition.
	// Context: Used to indicate "to where" or "to what state" an action or change proceeds.
	// Syntax/Example: (robot MOVE TO (room bedroom)), (transition_state FROM (state idle) TO (state active))
	AT TokenClass = "AT" // Purpose: Specifies a current location, state, or point of focus.
	// Context: Used to assert that a subject is "at" a particular place or in a specific state.
	// Syntax/Example: (robot IS AT (location charging_station)), (task IS AT (status pending))
	VIA TokenClass = "VIA" // Purpose: means through which an action occurs.
	// Context: Used to indicate a necessary passage or medium for an action.
	// Syntax/Example: (robot MOVE TO (room laboratory) VIA (corridor main_hall))
	//
	// Temporal Relational Qualifiers
	// ----------------------
	// Predicates that define the temporal order or duration relative to other events or states.
	BEFORE TokenClass = "BEFORE" // Purpose: Specifies that an action or fact occurs prior to another specified event or time.
	// Context: Used to establish temporal precedence.
	// Syntax/Example: (activate_security_system BEFORE (time 06:00AM)), (cleanup_task BEFORE (guest_arrival))
	AFTER TokenClass = "AFTER" // Purpose: Specifies that an action or fact occurs subsequent to another specified event or time.
	// Context: Used to establish temporal succession.
	// Syntax/Example: (send_report AFTER (data_collection_complete)), (robot_returns AFTER (mission_accomplished))
	DURING TokenClass = "DURING" // Purpose: Specifies that an action or state holds true for the duration of another event or period.
	// Context: Used to establish temporal overlap.
	// Syntax/Example: (robot_status IS (charging) DURING (power_outage)), (monitor_sensor DURING (experiment_running))
	///
	// Causal, Consequential, and Behavioral Predicates
	// ----------------------
	// Predicates that express causality, required conditions, outcomes, or internal modes of operation.
	BECAUSE TokenClass = "BECAUSE" // Purpose: Specifies the cause or reason for a particular fact or action.
	// Context: Establishes a causal link in the knowledge base.
	// Syntax/Example: (alarm_triggered BECAUSE (smoke_detected)), (robot_stopped BECAUSE (battery_low))
	RESULT TokenClass = "RESULT" // Purpose: Specifies the direct outcome, consequence, or effect of an action or event.
	// Context: Links an action to its produced state or value.
	// Syntax/Example: (move_action RESULT (location updated)), (calculation RESULT (?sum))
	REQUIRE TokenClass = "REQUIRE" // Purpose: Specifies a necessary precondition or dependency for an action or state to occur.
	// Context: Used to define prerequisites that must be met.
	// Syntax/Example: (move_to_dock REQUIRES (battery_level >= 20)), (open_door REQUIRES (authentication_successful))
	ENABLE TokenClass = "ENABLE" // Purpose: Specifies that an action or state allows or makes possible another action or state.
	// Context: Defines an enabling relationship, where one event unlocks another.
	// Syntax/Example: (charge_complete ENABLES (robot_movement)), (door_unlocked ENABLES (entry_access))
	MODE TokenClass = "MODE" // Purpose: Specifies an operating mode or a qualitative state of a subject or system.
	// Context: Describes the specific behavioral configuration or state an agent is in.
	// Syntax/Example: (robot IS (MODE (SAFE))), (network_device IS (MODE (DIAGNOSTIC)))
	//
	// ======================================================== #
	// Funcional Predicates
	// ======================================================== #
	// These tokens enable functional programming paradigms within nexusL, allowing for
	// concise expression of behavior, higher-order functions, and flexible
	// inline definitions of executable logic, which is particularly useful for agents.
	//
	// Set Builder
	// ----------------------
	//
	LIST_BUILDER   TokenClass = "@("
	ARRAY_BUILDER  TokenClass = "@["
	SET_BUILDER    TokenClass = "@{"
	VECTOR_BUILDER TokenClass = "@<"
	//
	//
	//
	// REDUCE
	// ----------------------
	// as JS functional reduction
	// ej (accumulator, currentValue) => accumulator + currentValue
	REDUCE TokenClass = "REDUCE"
	//
	//
	// MONAD
	// ----------------------
	//
	MONAD TokenClass = "MONAD"
	//
	//
	// Monads Linker
	// ----------------------
	//
	LINKAGE TokenClass = "LINKAGE"
	//
	// Multi-Expression Lambda Definition
	// ----------------------
	// Introduces a lambda function that can contain a block of multiple expressions or statements.
	LAMBDA TokenClass = "LAMBDA" // Purpose: Defines an anonymous function (a lambda) that can contain a block of code with multiple expressions.
	// Context: Used when the function body requires more than a single expression, often involving sequencing of operations, local variable declarations, or control flow.
	// Syntax/Example: (LAMBDA (param1 param2) { (LET temp (+ param1 param2)); (RETURN (* temp 2)) })
	//
	// Concise Lambda Definition
	// ----------------------
	// An operator for defining anonymous functions with a single expression body, promoting brevity.

	// Purpose: Defines a concise anonymous function where the body is a single expression.
	//
	// Context: Used for simple, inline functions, often passed as arguments to higher-order functions (e.g., map, filter). The ARROW separates the parameter list from the expression body.
	//
	// Syntax/Example: (MAP (my_list) (x -> (+ x 1))), (FILTER (data) (item -> (< item 10)))
	ARROW TokenClass = "->"

	//
	// Pattern Matching
	MATCH TokenClass = "MATCH"
	//
	// 1. Currificamos func_a
	// curry func_a_curried = func_a;???
	// func_a_curried curry func_a;
	//
	// 2. Usamos el pipe para pasar los argumentos uno a uno, encadenando las funciones parciales
	//    Imaginemos que tenemos los valores para x, y, z: `q_val`, `y_val`, `z_val`
	//
	//    |> q_val  		 # Pasa `q_val` a `func_a_curried`. Devuelve `(y) => (z) => func_a(q_val, y, z)`
	//    |> y_val           # Pasa `y_val` a la función resultante. Devuelve `(z) => func_a(q_val, y_val, z)`
	//    |> z_val;          # Pasa `z_val` a la función resultante. Ahora se ejecuta `func_a(q_val, y_val, z_val)`
	//                       # y el resultado se asigna a `resultado_final`.
	//
	// El resultado sería equivalente a func_a(q_val, y_val, z_val)
	// print(func_a_curried);
	CURRY TokenClass = "CURRY"
	PIPE  TokenClass = "|>"
	//
	MONADIC_PIPE TokenClass = "|*"
	//
	// ======================================================== #
	// Comments
	// ======================================================== #
	// These tokens are used to embed human-readable notes and explanations directly within the code.
	// Comments are ignored by the nexusL parser and do not affect program execution,
	// serving purely for documentation and clarity.
	//
	// Comment Styles
	// ----------------------
	// Different delimiters for various commenting needs.
	//
	// SINGLE_LINE_COMMENT
	// Purpose: Marks the start of a single-line comment.
	// Context: All text from this token to the end of the line is a comment.
	// Syntax/Example: // This is a comment about the next line of code.
	SINGLE_LINE_COMMENT TokenClass = "//"
	// COMMENT_INLINE
	// Purpose: Marks the start of a single-line comment (alternative style, common in scripting).
	// Context: All text from this token to the end of the line is a comment.
	// Syntax/Example: (my_function arg) # This is an inline comment.
	COMMENT_INLINE TokenClass = "#"
	// COMMENT_MULTI_LINE
	// Purpose: Marks the start of a multi-line comment block.
	// Context: All text between `/*` and the matching `*/` is treated as a comment.
	// Syntax/Example:
	/* 				This is a
							multi-line
	   				comment block. */
	COMMENT_MULTI_LINE TokenClass = "/*"
	//
	// ======================================================== #
	// Operators and Delimiters
	// ======================================================== #
	// These tokens define the fundamental operators for various computations
	// and the structural delimiters that organize expressions and blocks in nexusL.
	// They are crucial for parsing and interpreting the syntax of the language.
	//
	// Delimiters and Grouping
	// ----------------------
	// Structural tokens used to organize expressions, blocks, and data structures.
	//
	// Purpose: Used as a separator in key-value pairs, type annotations, or specific syntax patterns.
	// Context: Common for `key:value,`.
	// Syntax/Example: {"address":"abc 53", "id":234}
	COLON TokenClass = ":"
	// Purpose: Path/object linker,
	// Context: Common for `param::Type`, or `modalVerb::mainVerb`.
	// Syntax/Example: (SET has::username "Alice")
	RESOLUTION TokenClass = "::"
	// Purpose: Separates distinct statements or expressions within a block.
	// Context: Indicates the end of a logical unit of code, allowing multiple statements on one line or within a block.
	// Syntax/Example: { (PRINT "Hello"); (CALL_FUNCTION) }
	SEMICOLON TokenClass = ";"
	// Purpose: Separates elements in lists, parameters, or arguments.
	// Context: Used within sequences of items (e.g., `(LIST 1, 2, 3)` or `(FUNC f (param1, param2))`).
	// Syntax/Example: (CREATE_POINT 10, 20)
	COMMA TokenClass = ","
	// Purpose: Member access operator, used to access properties or methods of an entity/object.
	// Context: Connects an object identifier to its attribute or sub-property.
	// Syntax/Example: robot.battery_level, sensor.readings.latest
	DOT TokenClass = "."
	//
	// Data Structures Delimiters
	// ----------------------
	// Specific delimiters for various collection types.

	/*
		Purpose: Marks the beginning of a list or S-expression.

		Context: Fundamental for nexusL's Lisp-inspired syntax, encapsulating function calls, special forms, and data lists.

		Syntax/Example:
					FACT (subject predicate object);
	*/

	/*
	 Purpose: Marks the end of a list or S-expression.

	 Context: Closes the corresponding LPAREN.

	 Syntax/Example:

	 		func_call( arg1 arg2)
	*/

	LPAREN TokenClass = "("
	// Purpose: Marks the beginning of an array or vector literal.
	// Context: Used for defining ordered collections of elements.
	// Syntax/Example: [1 2 3], ["apple" "banana"]
	RPAREN TokenClass = ")"
	// Purpose: Marks the end of an array or vector literal.
	// Context: Closes the corresponding LBRACKET.
	// Syntax/Example: [element1 element2]
	LBRACKET TokenClass = "["
	// Purpose: Marks the beginning of a code block or map/dictionary literal.
	// Context: Used for function bodies, control flow blocks (IF, WHILE), or defining unordered key-value collections.
	// Syntax/Example: { (statement1); (statement2) }, (MAP "key1" "val1" "key2" "val2")
	RBRACKET TokenClass = "]"
	// Purpose: Marks the end of a code block or map/dictionary literal.
	// Context: Closes the corresponding LCURLY.
	// Syntax/Example: { (statement1); (statement2) }
	LCURLY TokenClass = "{"
	RCURLY TokenClass = "}"
	//
	LVECT TokenClass = "<["
	RVECT TokenClass = "]>"
	//
	// ======================================================== #
	// Mathematical and Logical Operators
	// ======================================================== #
	// These tokens define operations for performing calculations, comparisons,
	// and logical evaluations within nexusL. They are fundamental for processing data,
	// evaluating conditions, and driving decision-making processes within intelligent agents.
	//
	// Assignment Operators
	// ----------------------
	// Operators used to assign values to variables or update their existing values.
	AT_EACH      TokenClass = "@"
	ASSIGN_EQUAL TokenClass = "=" // Purpose: Assigns a value to a variable or property.
	// Context: Used for direct assignment. Note: This differs from `EQUALITY` (==) which is for comparison.
	// Syntax/Example: (SET my_var = 10), (robot HAS (color = "red"))
	// Note: If you also have `:=` for assignment, consider if `=` should be for comparison or assignment.
	// Given `EQUALITY (==)`, this `=` is best kept strictly for assignment.
	//
	// Arithmetic Operators
	// ----------------------
	// Perform standard mathematical calculations on numerical values.
	PLUS TokenClass = "+" // Purpose: Performs addition.
	// Context: Sums two or more numerical values.
	// Syntax/Example: (+ 5 3), (+ x y z)
	MINUS TokenClass = "-" // Purpose: Performs subtraction. Also used for unary negation.
	// Context: Calculates the difference between numbers or negates a single number.
	// Syntax/Example: (- 10 4), (- my_value)
	MULTIPLY TokenClass = "*" // Purpose: Performs multiplication.
	// Context: Calculates the product of two or more numerical values.
	// Syntax/Example: (* 2 8), (* width height)
	DIVIDE TokenClass = "/" // Purpose: Performs division.
	// Context: Divides the first numerical value by the second.
	// Syntax/Example: (/ 20 5), (/ total_distance total_time)
	MODULO TokenClass = "%" // Purpose: Computes the remainder of a division operation.
	// Context: Useful for cyclical operations, checking divisibility, or generating patterns.
	// Syntax/Example: (% 10 3) // Result is 1
	POWER TokenClass = "**" // Purpose: Exponentiation operator.
	// Context: Raises a base number to a given power.
	// Syntax/Example: (** 2 3) // 2^3 = 8
	//
	// Comparison Operators
	// ----------------------
	// Used to compare two values, yielding a boolean (`true` or `false`) outcome.
	EQUALITY TokenClass = "==" // Purpose: Checks if two values are equal.
	// Context: Used in conditional statements to test for equivalence.
	// Syntax/Example: (IF (== sensor_reading 100) ...)
	NOT_EQUALS TokenClass = "!=" // Purpose: Checks if two values are not equal.
	// Context: Used to test for dissimilarity.
	// Syntax/Example: (IF (!= current_state "idle") ...)
	GREATER TokenClass = ">" // Purpose: Checks if the left-hand value is strictly greater than the right-hand value.
	// Context: Used for range checks or ordering comparisons.
	// Syntax/Example: (IF (> temperature 25) ...)
	LESS TokenClass = "<" // Purpose: Checks if the left-hand value is strictly less than the right-hand value.
	// Context: Used for range checks or ordering comparisons.
	// Syntax/Example: (IF (< battery_level 10) ...)
	GREATER_EQUALS TokenClass = ">=" // Purpose: Checks if the left-hand value is greater than or equal to the right-hand value.
	// Context: Inclusive range checks.
	// Syntax/Example: (IF (>= score 90) ...)
	LESS_EQUALS TokenClass = "<=" // Purpose: Checks if the left-hand value is less than or equal to the right-hand value.
	// Context: Inclusive range checks.
	// Syntax/Example: (IF (<= speed_limit 60) ...)
	//
	// Bitwise Operators
	// ----------------------
	// Perform operations directly on the binary representations (bits) of integer numbers.
	// Useful for low-level data manipulation, flags, or specialized algorithms.
	BIT_AND TokenClass = "&" // Purpose: Bitwise AND. Performs a bitwise AND operation on corresponding bits of two integers.
	// Context: Used for masking bits or checking if specific bits are set.
	// Syntax/Example: (& 0b1010 0b1100) // Result: 0b1000
	BIT_OR TokenClass = "|" // Purpose: Bitwise OR. Performs a bitwise OR operation on corresponding bits of two integers.
	// Context: Used for setting specific bits.
	// Syntax/Example: (| 0b1010 0b0011) // Result: 0b1011
	BIT_XOR TokenClass = "^" // Purpose: Bitwise XOR (Exclusive OR). Performs a bitwise XOR operation.
	// Context: Used for toggling bits or encryption.
	// Syntax/Example: (^ 0b1010 0b0110) // Result: 0b1100
	BIT_NOT TokenClass = "~" // Purpose: Bitwise NOT (Inversion). Inverts all bits of an integer.
	// Context: Flips all bits (0 becomes 1, 1 becomes 0).
	// Syntax/Example: (~ 0b1010) // (Depends on integer size, e.g., for 8-bit, 0b1010 becomes 0b0101 if unsigned)
	BIT_SHL TokenClass = "<<" // Purpose: Bitwise Shift Left. Shifts bits to the left by a specified number of positions.
	// Context: Equivalent to multiplying by powers of 2.
	// Syntax/Example: (<< 1 3) // Result: 8 (0b0001 << 3 = 0b1000)
	BIT_SHR TokenClass = ">>" // Purpose: Bitwise Shift Right. Shifts bits to the right by a specified number of positions.
	// Context: Equivalent to integer division by powers of 2.
	// Syntax/Example: (>> 8 2) // Result: 2 (0b1000 >> 2 = 0b0010)
	//
	// Compound Assignment Operators
	// Combine an arithmetic operation with an assignment, providing a concise way to update variables.
	//
	ASSIGN TokenClass = ":=" // Purpose: Direct assignment operator.
	// Context: Assigns a value to a variable (e.g., in a `VAR` declaration or `SET` command).
	// Syntax/Example: (VAR x := 10)
	// Note: The original `is:` could be mapped to this during parsing if `(LET x is: Value)` becomes `(LET x := Value)`
	ASSIGN_PLUS TokenClass = "+=" // Purpose: Adds the right-hand value to the left-hand variable and assigns the result.
	// Context: Shorthand for `x = x + value`.
	// Syntax/Example: (SET counter += 1)
	ASSIGN_MINUS TokenClass = "-=" // Purpose: Subtracts the right-hand value from the left-hand variable and assigns the result.
	// Context: Shorthand for `x = x - value`.
	// Syntax/Example: (SET health -= 10)
	ASSIGN_MULTIPLY TokenClass = "*=" // Purpose: Multiplies the left-hand variable by the right-hand value and assigns the result.
	// Context: Shorthand for `x = x * value`.
	// Syntax/Example: (SET scale *= 2)
	ASSIGN_DIVIDE TokenClass = "/=" // Purpose: Divides the left-hand variable by the right-hand value and assigns the result.
	// Context: Shorthand for `x = x / value`.
	// Syntax/Example: (SET total_score /= 2)
	ASSIGN_MODULO TokenClass = "%=" // Purpose: Computes the modulus of the left-hand variable by the right-hand value and assigns the result.
	// Context: Shorthand for `x = x % value`.
	// Syntax/Example: (SET index %= max_items)
	//
	// Increment and Decrement Operators
	// Increase or decrease the value of a variable by one. Typically used for counters.
	//
	// Purpose: Increases the value of a variable by 1.
	// Context: Shorthand for `x = x + 1`. Often used in loops or counters.
	// Syntax/Example: (SET loop_count ++)
	// INCREMENT TokenClass = "++"
	// Purpose: Decreases the value of a variable by 1.
	// Context: Shorthand for `x = x - 1`. Often used in loops or counters.
	// Syntax/Example: (SET remaining_attempts --)
	// DECREMENT TokenClass = "--"
	//
	//
	//
	// ES NECESARIO IMPLEMENTAR COMO KEYWORDS PARA LOS FACTS QUE GENERA EL LLM INTERNO
	// NO ESTAN DESTINADOS A SER INVOCADOS O CONSIMIDOS EN LA PARTE MACANICA O LOGICA
	// SI NO EN LA ETAPA DE RAZONAMIENTO BASICO -  SERAN LA PIEDRA ANGULAR DEL RAZONAMIENTO
	// COMPLEJO... EN UN FUTURO.
	/*
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
	*/
)
