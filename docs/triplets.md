ok, bueno si. Todo el esfuerso fue a nivel de bit. Pero bueno, ahora bien tengo diversos "tipos de triplets" (este ha sido el fallo recurrente) asi que quisiera comenzar ya definido el Entity con esta estructura. Los tipos son:

describe: Triplet "plana" de declaracion de existencia
como: 
    - detecta verbo auxiliar (is or do) como segundo elemento simple
    - entonces tercer elemento es verboName y es Entity
    - attributo no existe

``` nexus
david is Entity;
```
seria parseado como:

sujeto: david verboAux:is verboName:Entity attrbName:"" attribValue:""
y daria al final como resultado

```go
david := NewEntity()
david.InstantiateAs("david")
// se agrega a bbolt
nBucket.add(blalbla)// AUN POR DEFINIR
```

se agrega a datalog

```prolog

existAs(david, Entity).
```


describe: Triplet "plana" de declaracion de accion "simple"
como: 
    - detecta verbo auxiliar (is or do) como segundo elemento simple
    - entonces tercer elemento es verboName
    - attributo no existe

*** Actualemte solo esos dos verbos, pero se extenderia hacia (si es muy complejo):
	COULD  // Past
	CAN // present
	BE_ABLE_TO  // Future
	// Permission or possibility
	// These keywords indicate the permission or possibility of an action.
	MAYBE  // Past
	PERMISSION   // Present
	ALLOWED_TO   // Future
	// Possibility
	// These keywords indicate the possibility of an action.
	// They are used to express uncertainty or potential outcomes.
	MIGHT // Past
	MAY    // Present - future
	// Necessity or obligation
	// These keywords indicate the necessity or obligation of an action.
	// They are used to express requirements or recommendations.
	HAD_TO    // Past
	MUST    // Present
	WILL_HAVE_TO  // Future
	// Suggestion or recommendation
	// These keywords indicate a suggestion or recommendation for an action.
	// They are used to express advice or guidance.
	SHOULD_HAVE // Past
	SHOULD    // Present - future
	// Requirement or necessity
	// These keywords indicate a requirement or necessity for an action.
	// They are used to express obligations or needs.
	NEED_TO   // Past
	NEED    // Present - future
	WILL_NEED   //
``` nexus
david had_to run;
david must no:run;
david willHaveTo run;
```
seria parseado como:

sujeto: david verboAux:is verboName:Entity attrbName:"" attribValue:""
y daria al final como resultado

```go
// AUN POR DEFINIR
//
// se agrega a bbolt
nBucket.add(blalbla)// AUN POR DEFINIR
```

se agrega a datalog

```prolog

run(david).
```
