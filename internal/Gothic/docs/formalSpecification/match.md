### **Pattern Matching**

En esencia, los patrones de coincidencia te permiten **comparar una expresión (o un valor) con una serie de "patrones"** definidos. Cuando la expresión coincide con un patrón, se ejecuta el bloque de código asociado a ese patrón. Va más allá de una simple igualdad; puede descomponer estructuras de datos, vincular variables a partes de los datos, e incluso evaluar condiciones complejas.

Es como una estructura switch/case súper cargada.

---

### **¿Cómo Funciona?**

Un patrón de coincidencia típicamente implica:

1. **Una expresión a comparar:** El valor o dato que quieres analizar.  
2. **Múltiples "brazos" o "cláusulas":** Cada uno con un patrón y un bloque de código.  
3. **Evaluación secuencial (generalmente):** Los patrones se evalúan en orden. El primer patrón que coincide es el que se ejecuta.  
4. **Descomposición de datos:** Permite extraer valores de estructuras complejas (tuplas, objetos, uniones, etc.) directamente en variables.  
5. **Patrones de guardia (opcional):** Condiciones adicionales que deben ser verdaderas para que un patrón coincida.  
6. **Patrón comodín:** Un patrón que coincide con cualquier valor (similar a un default en un switch).

---

### **Ejemplos Conceptuales**

Veamos cómo se vería en pseudocódigo o en un lenguaje con esta característica (como Rust, Scala, Elixir, Haskell o incluso JavaScript con algunas propuestas):

#### **1\. Coincidencia por Valor Simple**

// Imagina una función que clasifica un número  
fn clasificar\_numero(num):  
  match num:  
    case 0:  
      print("Es cero")  
    case 1:  
      print("Es uno")  
    case \_: // Patrón comodín: coincide con cualquier otro valor  
      print("Es otro número")

clasificar\_numero(0) // Imprime "Es cero"  
clasificar\_numero(5) // Imprime "Es otro número"

---

#### **2\. Coincidencia por Estructura (Descomposición de Datos)**

Aquí es donde los patrones de coincidencia se vuelven realmente potentes. Puedes extraer los componentes de un objeto o una tupla directamente en variables.

// Imagina una estructura de "Punto"  
struct Punto { x: Int, y: Int }

// Una función para describir la posición de un punto  
fn describir\_punto(p: Punto):  
  match p:  
    case Punto { x: 0, y: 0 }:  
      print("El punto está en el origen.")  
    case Punto { x: val\_x, y: 0 }: // Si 'y' es 0, extrae 'x' en 'val\_x'  
      print("El punto está en el eje X en", val\_x)  
    case Punto { x: 0, y: val\_y }: // Si 'x' es 0, extrae 'y' en 'val\_y'  
      print("El punto está en el eje Y en", val\_y)  
    case Punto { x: val\_x, y: val\_y }: // Extrae ambos valores  
      print("El punto está en (", val\_x, ",", val\_y, ")")

describir\_punto(Punto { x: 0, y: 0 })   // Imprime "El punto está en el origen."  
describir\_punto(Punto { x: 5, y: 0 })   // Imprime "El punto está en el eje X en 5"  
describir\_punto(Punto { x: 3, y: 7 })   // Imprime "El punto está en (3, 7)"

---

#### **3\. Coincidencia con Patrones de Guardia (Guard Clauses)**

Puedes añadir condiciones if a un patrón para una coincidencia más específica.

fn clasificar\_edad(edad: Int):  
  match edad:  
    case val if val \< 0:  
      print("Edad inválida (negativa).")  
    case val if val \>= 0 and val \<= 12:  
      print("Es un niño.")  
    case val if val \> 12 and val \<= 18:  
      print("Es un adolescente.")  
    case \_:  
      print("Es un adulto.")

clasificar\_edad(-5)  // Imprime "Edad inválida (negativa)."  
clasificar\_edad(15)  // Imprime "Es un adolescente."

---

#### **4\. Coincidencia con Tipos de Unión (Enums o ADTs \- Algebraic Data Types)**

Esto es fundamental en lenguajes como Rust o Haskell, donde puedes tener un tipo que puede ser uno de varios valores posibles.

// Definición de un tipo 'Mensaje' que puede ser de diferentes variantes  
enum Mensaje {  
  Quit,  
  Move { x: Int, y: Int },  
  Write(String),  
  ChangeColor(Int, Int, Int),  
}

fn procesar\_mensaje(msg: Mensaje):  
  match msg:  
    case Mensaje::Quit:  
      print("El programa se cerrará.")  
    case Mensaje::Move { x, y }: // Descompone directamente x e y  
      print("Moviendo a (", x, ",", y, ")")  
    case Mensaje::Write(texto): // Extrae el String interno  
      print("Escribiendo texto:", texto)  
    case Mensaje::ChangeColor(r, g, b): // Extrae los tres Int  
      print("Cambiando color a RGB(", r, ",", g, ",", b, ")")

procesar\_mensaje(Mensaje::Quit)  
procesar\_mensaje(Mensaje::Move { x: 10, y: 20 })  
procesar\_mensaje(Mensaje::Write("Hola mundo".to\_string()))

---

### **¿Por qué son tan Importantes en la Programación Funcional?**

1. **Exhaustividad:** Muchos lenguajes con patrones de coincidencia pueden verificar si has cubierto *todos los casos posibles*. Si te olvidas de un patrón, el compilador puede advertirte, lo que reduce los errores en tiempo de ejecución.  
2. **Claridad y Legibilidad:** Reemplazan anidados if/else y switch gigantes con una estructura clara y concisa que muestra las posibles formas de los datos y cómo se maneja cada una.  
3. **Descomposición Implícita:** Simplifican el acceso a los datos dentro de estructuras complejas, eliminando la necesidad de múltiples líneas de extracción manual.  
4. **Inmutabilidad:** Fomentan un estilo de programación inmutable, ya que no se modifican los datos, sino que se reacciona a su forma.  
5. **Facilitan el Manejo de ADTs:** Son la forma natural de trabajar con Tipos de Datos Algebraicos (como los enums de Rust o data de Haskell), que son fundamentales en la programación funcional para modelar dominios complejos.

---

En resumen, los patrones de coincidencia son una herramienta poderosa que mejora la robustez, la legibilidad y la expresividad del código, especialmente cuando se trabaja con datos estructurados y se necesita manejar diferentes casos de manera elegante.

¿Hay algún aspecto particular de los patrones de coincidencia que te interese explorar más a fondo?