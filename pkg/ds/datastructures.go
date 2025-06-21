package ds

import "fmt"

// LIST
// LA IDEA ES IMPLEMENTAR UNA LISTA DE FUNCIONAMIENTO PYTHONICO
// EL INDICE PERMITE OPERACIONES POR INDICE
// SE GUARDA EL NOMBRE DEL OBJETO PARA LLAMARLO SI SE USA
// EL VALOR NO SE UTILIZA -
//
// # RECORDAR ESO PARA EL MECANISMO DE BORRADO DE VARIABLES Y ENTITY
//
/*

Ordered: The items in a list maintain their insertion order.
Changeable (Mutable): You can modify, add, or remove elements from a list after it has been created.
Indexed: Each item in a list has an index, starting from 0 for the first element, allowing you to access specific items directly. Negative indices can be used to access items from the end of the list (e.g., -1 for the last item).
Allow Duplicates: Lists can contain multiple occurrences of the same value.
Heterogeneous: Lists can store items of different data types within the same list (e.g., integers, strings, other lists).

Common List Operations:
Accessing Elements:
Use indexing (e.g., my_list[0]) or slicing (e.g., my_list[1:3]).
Modifying Elements:
Assign a new value to an element at a specific index (e.g., my_list[0] = 'New Value').
Adding Elements:
Use methods like append() to add to the end, insert() to add at a specific position, or extend() to add elements from another iterable.
Removing Elements:
Use methods like remove() to remove a specific value, pop() to remove by index, or the del keyword to remove by index or slice.
Other Operations:
Includes sorting (sort()), reversing (reverse()), checking for membership (in operator), and finding the length (len()).

*/

type nxsList struct {
	values []interface{} // Usamos un slice para almacenar los valores
}

func xList() *nxsList {
	return &nxsList{
		values: make([]interface{}, 0),
	}
}

func (nl *nxsList) Append(value interface{}) {
	nl.values = append(nl.values, value)
}

func (nl *nxsList) Remove(index int) error {
	if index < 0 || index >= len(nl.values) {
		return fmt.Errorf("index out of range")
	}
	// Eliminar el elemento en el índice dado
	nl.values = append(nl.values[:index], nl.values[index+1:]...)
	return nil
}

func (nl *nxsList) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(nl.values) {
		return nil, fmt.Errorf("index out of range")
	}
	return nl.values[index], nil
}

func (nl *nxsList) Print() {
	for i, val := range nl.values {
		fmt.Printf("Index: %d, Value: %v\n", i, val)
	}
}

// DICT
// LA IDEA ES IMPLEMENTAR UNA DICT DE FUNCIONAMIENTO PYTHONICO
// EL DICT PERMITE OPERACIONES POR KEY
// SE GUARDA EL NOMBRE DEL OBJETO PARA LLAMARLO SI SE USA
// EL VALOR SI SE UTILIZA -
//
/*

Key-Value Pairs:
Each element in a dictionary consists of a unique key associated with a specific value.
Unordered (in older Python versions):
While modern Python versions (3.7+) maintain insertion order, dictionaries were historically unordered, meaning the order of items was not guaranteed.
Changeable (Mutable):
You can add, remove, or modify key-value pairs within a dictionary after it's created.
Keys must be immutable and unique:
Keys within a dictionary must be immutable data types (like strings, numbers, or tuples) and must be unique within that dictionary.
Efficient Lookups:
Dictionaries are implemented using hash tables, which allows for very fast retrieval of values by their corresponding keys.

Accessing Dictionary Items:
Values are accessed using their corresponding keys:

Modifying Dictionaries:
Adding/Updating: Assign a value to a new or existing key:

Removing: Use the del keyword or the pop() method:

*/
type nxsPairAtom struct {
	itemName string
	value    interface{} // Usamos un slice para almacenar los valores
}

func xDict(atoms ...nxsPairAtom) map[string]interface{} {
	// Declare and initialize a map
	names := make(map[string]any)
	for _, atom := range atoms {
		// Add a new key-value pair
		names[atom.itemName] = atom.value
	}
	return names
}

// SET
// LA IDEA ES IMPLEMENTAR UNA SET DE FUNCIONAMIENTO PYTHONICO
// EL SET PERMITE OPERACIONES POR VALUE
// SE GUARDA EL NOMBRE DEL OBJETO PARA LLAMARLO
// EL VALOR SI SE UTILIZA -
//
//
/*

Mathematical Set Operations:
Sets support various mathematical operations, such as:
Union: Combining elements from two or more sets.
Intersection: Finding common elements between sets.
Difference: Finding elements present in one set but not in another.
Symmetric Difference: Finding elements unique to each of two sets.


Key characteristics of Python sets:
Unordered:
Elements within a set do not have a defined order, meaning you cannot access them by index.
Unique Elements:
Sets automatically handle duplicate entries, ensuring that each element appears only once. If you try to add an existing element, the set will not change.
Mutable (but elements must be immutable):
You can add or remove elements from a set after it's created. However, the individual elements stored within a set must be immutable data types (e.g., numbers, strings, tuples). You cannot store mutable objects like lists or dictionaries directly within a set.

*/

// Función para crear un conjunto
func nxsSet(elements ...interface{}) map[interface{}]struct{} {
	set := make(map[interface{}]struct{})
	for _, element := range elements {
		set[element] = struct{}{}
	}
	return set
}

// Función para añadir un elemento al conjunto
func addToSet(set map[interface{}]struct{}, element interface{}) {
	set[element] = struct{}{}
}

// Función para eliminar un elemento del conjunto
func removeFromSet(set map[interface{}]struct{}, element interface{}) {
	delete(set, element)
}

// Función para verificar si un elemento está en el conjunto
func contains(set map[interface{}]struct{}, element interface{}) bool {
	_, exists := set[element]
	return exists
}

// Función para realizar la unión de dos conjuntos
func union(set1, set2 map[interface{}]struct{}) map[interface{}]struct{} {
	result := make(map[interface{}]struct{})
	for element := range set1 {
		result[element] = struct{}{}
	}
	for element := range set2 {
		result[element] = struct{}{}
	}
	return result
}

// Función para realizar la intersección de dos conjuntos
func intersection(set1, set2 map[interface{}]struct{}) map[interface{}]struct{} {
	result := make(map[interface{}]struct{})
	for element := range set1 {
		if contains(set2, element) {
			result[element] = struct{}{}
		}
	}
	return result
}

// Función para realizar la diferencia de dos conjuntos
func difference(set1, set2 map[interface{}]struct{}) map[interface{}]struct{} {
	result := make(map[interface{}]struct{})
	for element := range set1 {
		if !contains(set2, element) {
			result[element] = struct{}{}
		}
	}
	return result
}

// TREE
//
// https://blog.devgenius.io/trees-in-go-9b6ff346dcfc
// https://certik.github.io/scipy-2013-tutorial/html/tutorial/manipulation.html
// .
// Estructura para un nodo del árbol
type nxsmaryNode struct {
	data  int
	left  *nxsmaryNode
	right *nxsmaryNode
}

// Estructura para el árbol
type nxsTree struct {
	root *nxsmaryNode
}

// Función para insertar un nodo en el árbol
func (t *nxsTree) insert(data int) {
	if t.root == nil {
		t.root = &nxsmaryNode{data: data}
	} else {
		t.root.insert(data)
	}
}

// Método para insertar un nodo (recursivo)
func (node *nxsmaryNode) insert(data int) {
	if data < node.data {
		if node.left == nil {
			node.left = &nxsmaryNode{data: data}
		} else {
			node.left.insert(data)
		}
	} else if data > node.data {
		if node.right == nil {
			node.right = &nxsmaryNode{data: data}
		} else {
			node.right.insert(data)
		}
	}
	// Si data == node.data, no hacemos nada (no permitimos duplicados)
}

// Función para recorrer el árbol en orden (in-order traversal)
func (t *nxsTree) traverseInOrder() {
	if t.root != nil {
		t.root.traverseInOrder()
	}
}

// Método para recorrer el árbol en orden (recursivo)
func (node *nxsmaryNode) traverseInOrder() {
	if node.left != nil {
		node.left.traverseInOrder()
	}
	fmt.Print(node.data, " ")
	if node.right != nil {
		node.right.traverseInOrder()
	}
}

// m-ary tree
/*
La estructura de datos propuesta para el Árbol de Sintaxis Abstracta (AST) con múltiples hijos es un tipo de árbol general, también conocido como árbol n-ario. En este tipo de árbol, cada nodo puede tener cero, uno, o más hijos, lo cual lo hace adecuado para representar estructuras jerárquicas complejas como las que se encuentran en un AST.

Características del Árbol N-ario
Múltiples Hijos: A diferencia de un árbol binario, donde cada nodo tiene como máximo dos hijos, en un árbol n-ario, cada nodo puede tener un número arbitrario de hijos. Esto es útil para representar construcciones sintácticas que pueden tener múltiples sub-expresiones o sub-declaraciones.

Flexibilidad: Esta estructura es muy flexible y puede adaptarse a una amplia variedad de lenguajes y gramáticas. Puedes tener nodos que representan diferentes tipos de construcciones sintácticas, como expresiones, declaraciones, literales, operadores, etc.

Recorrido: Los árboles n-arios pueden ser recorridos utilizando algoritmos de recorrido en profundidad (depth-first) o en anchura (breadth-first). El recorrido en profundidad es comúnmente utilizado para evaluar o transformar un AST.

Representación de Jerarquías: Los árboles n-arios son excelentes para representar jerarquías y relaciones anidadas, lo cual es esencial para un AST, donde las construcciones del lenguaje pueden estar anidadas unas dentro de otras.

Ejemplo de Uso en un AST
En el contexto de un AST, un árbol n-ario puede ser utilizado para representar la estructura sintáctica de un programa. Por ejemplo:

Nodo Raíz: Representa el programa completo.
Nodos Intermedios: Representan construcciones como funciones, bucles, condicionales, etc.
Nodos Hoja: Representan literales, identificadores, operadores, etc.

*/
// Tipo de nodo del M-ary tree
type NodeType string

// Estructura para un nodo del AST
type maryNode struct {
	NodeType NodeType
	Value    string      // Valor o dato asociado con el nodo
	Children []*maryNode // Hijos del nodo
}

// Estructura para el árbol
type maryTree struct {
	Root *maryNode
}

// Función para agregar un nodo hijo a un nodo padre específico
func (tree *maryTree) AddNode(parentValue interface{}, newNode *maryNode) bool {
	parent := tree.findNode(tree.Root, parentValue)
	if parent != nil {
		parent.Children = append(parent.Children, newNode)
		return true
	}
	return false
}

// Función para encontrar un nodo con un valor específico (recursivo)
func (tree *maryTree) findNode(node *maryNode, value interface{}) *maryNode {
	if node == nil {
		return nil
	}
	if node.Value == value {
		return node
	}
	for _, child := range node.Children {
		foundNode := tree.findNode(child, value)
		if foundNode != nil {
			return foundNode
		}
	}
	return nil
}

// Función para eliminar un nodo con un valor específico
func (tree *maryTree) RemoveNode(value interface{}) bool {
	if tree.Root == nil {
		return false
	}
	if tree.Root.Value == value {
		tree.Root = nil
		return true
	}
	return tree.removeNode(nil, tree.Root, value)
}

// Función auxiliar para eliminar un nodo (recursivo)
func (tree *maryTree) removeNode(parent *maryNode, node *maryNode, value interface{}) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		if parent != nil {
			for i, child := range parent.Children {
				if child == node {
					parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
					return true
				}
			}
		}
		return false
	}
	for _, child := range node.Children {
		if tree.removeNode(node, child, value) {
			return true
		}
	}
	return false
}

// Función para seleccionar un nodo con un valor específico
func (tree *maryTree) SelectNode(value interface{}) *maryNode {
	return tree.findNode(tree.Root, value)
}

// Función para recorrer el árbol (recorrido en profundidad)
func (node *maryNode) Traverse(depth int) {
	fmt.Printf("%sNode: %v (Depth: %d)\n", string(node.NodeType), node.Value, depth)
	for _, child := range node.Children {
		child.Traverse(depth + 1)
	}
}

/*
func main() {
	list := NewnxsList()
	list.Append("valor1")
	list.Append("valor2")
	list.Append("valor3")

	fmt.Println("Lista inicial:")
	list.Print()

	err := list.Remove(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nLista después de eliminar el índice 1:")
	list.Print()
}

    // Crear un diccionario utilizando la función xDict
    myDict := xDict(
        nxsPairAtom{itemName: "nombre", value: "Juan"},
        nxsPairAtom{itemName: "edad", value: 30},
        nxsPairAtom{itemName: "ciudad", value: "Madrid"},
    )

    // Imprimir el diccionario
    fmt.Println("Diccionario:", myDict)

    // Acceder a un valor en el diccionario
    if nombre, ok := myDict["nombre"].(string); ok {
        fmt.Println("Nombre:", nombre)
    }


    // Crear conjuntos
    set1 := nexuslSet(1, 2, 3, 4)
    set2 := nexuslSet(3, 4, 5, 6)

    // Imprimir conjuntos
    fmt.Println("Set 1:", set1)
    fmt.Println("Set 2:", set2)

    // Añadir un elemento al conjunto
    addToSet(set1, 5)
    fmt.Println("Set 1 después de añadir 5:", set1)

    // Eliminar un elemento del conjunto
    removeFromSet(set1, 1)
    fmt.Println("Set 1 después de eliminar 1:", set1)

    // Verificar si un elemento está en el conjunto
    fmt.Println("Set 1 contiene 2:", contains(set1, 2))
    fmt.Println("Set 1 contiene 1:", contains(set1, 1))

    // Realizar operaciones de conjuntos
    fmt.Println("Unión de Set 1 y Set 2:", union(set1, set2))
    fmt.Println("Intersección de Set 1 y Set 2:", intersection(set1, set2))
    fmt.Println("Diferencia de Set 1 y Set 2:", difference(set1, set2))



	// Crear un nuevo árbol
	tree := &nxsTree{}

	// Insertar nodos en el árbol
	tree.insert(5)
	tree.insert(3)
	tree.insert(8)
	tree.insert(1)
	tree.insert(4)
	tree.insert(7)
	tree.insert(9)

	// Recorrer el árbol en orden
	fmt.Println("Recorrido en orden del árbol:")
	tree.traverseInOrder()






    // Crear un nuevo árbol
    tree := &maryTree{
        Root: &TreeNode{
            NodeType: RootNode,
            Value:    "Root",
        },
    }

    // Añadir nodos al árbol
    tree.AddNode("Root", &TreeNode{Value: "Child1"})
    tree.AddNode("Root", &TreeNode{Value: "Child2"})
    tree.AddNode("Child1", &TreeNode{Value: "Child1.1"})

    // Recorrer el árbol
    fmt.Println("Tree Traversal:")
    tree.Root.Traverse(0)

    // Seleccionar un nodo
    node := tree.SelectNode("Child1.1")
    if node != nil {
        fmt.Printf("\nSelected Node: %v\n", node.Value)
    }

    // Eliminar un nodo
    tree.RemoveNode("Child1")
    fmt.Println("\nTree Traversal after removing Child1:")
    tree.Root.Traverse(0)
}
*/
