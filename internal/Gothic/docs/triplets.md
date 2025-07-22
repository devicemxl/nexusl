ok, bueno si. Todo el esfuerso fue a nivel de bit. Pero bueno, ahora bien tengo diversos "tipos de triplets" (este ha sido el fallo recurrente) asi que quisiera comenzar ya definido el Entity con esta estructura. Los tipos son:

describe: Triplet "plana" de declaracion de existencia
como: 
    - detecta verbo auxiliar (is or do) como segundo elemento simple
    - entonces tercer elemento es verboName y es Entity
    - attributo no existe

``` nexus
David is entity;
```
seria parseado como:

sujeto: david verboAux:is verboName:entity attrbName:"" attribValue:""
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

``` nexus
david need run;
```

seria parseado como:

sujeto: david verboAux:need verboName:run attrbName:"" attribValue:""
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
