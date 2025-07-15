## **Curry con |\>**

El **currying** y el **operador pipe (|\>)** son dos pilares de la programación funcional que, cuando se combinan, desbloquean un estilo de codificación excepcionalmente **legible, modular y expresivo** en tu Domain Specific Language (DSL). Más allá de ser meras "azúcares sintácticas", estas herramientas transforman fundamentalmente cómo interactúas con las funciones y cómo fluyen los datos a través de tu sistema.

---

### **Entendiendo el Currying**

En su esencia, el **currying** es una técnica que toma una función que espera múltiples argumentos y la transforma en una secuencia de funciones, cada una de las cuales acepta un solo argumento.

**Considera una función clásica:** sumar(a, b, c)

Sin currying, la llamarías: sumar(1, 2, 3\)

Con currying, la función sumar (una vez currificada) se convierte en: sumar\_curry(a)(b)(c)

Esto significa que cuando le das el primer argumento (a), no obtienes el resultado final, sino una **nueva función** que "recuerda" a y espera b. Cuando le das b, obtienes otra **nueva función** que espera c. Solo cuando se le proporciona el último argumento (c) se ejecuta la función original y se obtiene el resultado final.

---

### **El Poder del Operador Pipe (|\>)**

El **operador pipe (|\>)** es una construcción sintáctica que mejora la legibilidad de las secuencias de operaciones. Toma el **resultado de la expresión de su izquierda** y lo pasa como el **primer argumento a la función de su derecha**.

**Sin el operador pipe, las llamadas anidadas lucen así:**

resultado \= funcion\_c(funcion\_b(funcion\_a(dato\_inicial)))

Esta sintaxis se lee de adentro hacia afuera, lo que puede ser contraintuitivo y difícil de seguir en cadenas largas.

**Con el operador pipe, la lectura fluye de izquierda a derecha:**

resultado \= dato\_inicial |\> funcion\_a |\> funcion\_b |\> funcion\_c

Esto refleja el **flujo natural de los datos** a través de una serie de transformaciones, haciendo que el código sea mucho más fácil de leer y entender.

---

### **La Sinergia: Currying y Pipe Juntos**

La verdadera magia ocurre cuando combinas el currying con el operador pipe. El currying, al devolver una función parcial en cada paso, se integra perfectamente con la forma en que el pipe encadena las operaciones.

**Imagina una función func\_a(x, y, z) currificada como funcionAlCuryy.**

// 1\. Currificamos func\_a  
const funcionAlCuryy \= curry(func\_a); // funcionAlCuryy ahora es (x) \=\> (y) \=\> (z) \=\> func\_a(x, y, z)

// 2\. Encadenamos los argumentos usando el pipe  
//    Suponiendo que 'valor\_x', 'valor\_y', y 'valor\_z' son tus datos  
const resultado\_final \= valor\_x  
    |\> funcionAlCuryy             // Pasa valor\_x a funcionAlCuryy. Retorna una función que espera 'y' y 'z'.  
    |\> valor\_y                  // Pasa valor\_y a la función resultante. Retorna una función que espera 'z'.  
    |\> valor\_z;                 // Pasa valor\_z a la última función. Finalmente, func\_a(valor\_x, valor\_y, valor\_z) se ejecuta.

Este patrón es especialmente útil para: Transformaciones de Datos Secuenciales: 

Imagina que tienes una **estructura de datos** para la temperatura en Kelvin.

struct Temperatura {  
  valor: f64,  
  unidad: String // Para este ejemplo, asumimos "Kelvin"  
}

Ahora, quieres definir funciones que la conviertan a Celsius o Fahrenheit. Estas funciones tomarían una Temperatura y devolverían una *nueva* Temperatura en la unidad deseada.

**1\. Funciones de Conversión:**

Necesitaríamos funciones puras que realicen las conversiones.

* **Kelvin a Celsius:** C=K−273.15  
* **Celsius a Fahrenheit:** F=Ctimes1.8+32

Vamos a definirlas conceptualmente (en pseudocódigo que simula la forma de una función funcional):

// Función: kelvin\_a\_celsius  
// Toma una 'Temperatura' en Kelvin y devuelve una 'Temperatura' en Celsius  
fn kelvin\_a\_celsius(temp\_kelvin: Temperatura) \-\> Temperatura {  
  new\_valor \= temp\_kelvin.valor \- 273.15;  
  return Temperatura { valor: new\_valor, unidad: "Celsius" };  
}

// Función: celsius\_a\_fahrenheit  
// Toma una 'Temperatura' en Celsius y devuelve una 'Temperatura' en Fahrenheit  
fn celsius\_a\_fahrenheit(temp\_celsius: Temperatura) \-\> Temperatura {  
  new\_valor \= (temp\_celsius.valor \* 1.8) \+ 32.0;  
  return Temperatura { valor: new\_valor, unidad: "Fahrenheit" };  
}

---

**2\. Usando el Operador Pipe (|\>) para Encadenar Conversiones:**

Aquí es donde el operador pipe brilla. Puedes tomar tu valor inicial en Kelvin y pasarlo a través de una "tubería" de transformaciones:

// Supongamos que tienes una temperatura inicial en Kelvin  
let temp\_inicial\_kelvin \= Temperatura { valor: 300.15, unidad: "Kelvin" };

// Opción 1: Convertir de Kelvin a Celsius  
let temp\_en\_celsius \= temp\_inicial\_kelvin  
    |\> kelvin\_a\_celsius; // Resultado: Temperatura { valor: 27.0, unidad: "Celsius" }

// Opción 2: Convertir de Kelvin a Fahrenheit (pasando por Celsius)  
let temp\_en\_fahrenheit \= temp\_inicial\_kelvin  
    |\> kelvin\_a\_celsius       // Primero, a Celsius  
    |\> celsius\_a\_fahrenheit;  // Luego, de Celsius a Fahrenheit

// Resultado de temp\_en\_fahrenheit: Temperatura { valor: 80.6, unidad: "Fahrenheit" }

### **¿Por qué esto es poderoso?**

1. **Legibilidad:** La secuencia de operaciones es clara y fluye de izquierda a derecha, como lees una frase. Puedes ver inmediatamente el camino que sigue el dato.  
2. **Composición:** Estás componiendo funciones. celsius\_a\_fahrenheit está operando sobre el resultado de kelvin\_a\_celsius sin anidar llamadas de función (celsius\_a\_fahrenheit(kelvin\_a\_celsius(temp))).  
3. **Modularidad:** Cada función de conversión es independiente y realiza una única tarea bien definida.  
4. **Reutilización:** Puedes usar kelvin\_a\_celsius de forma aislada, o como parte de una cadena más larga.

Como vimos con el ejemplo de la temperatura, puedes definir una temperatura en Kelvin y luego transformarla a Celsius y luego a Fahrenheit de forma fluida y asi permitir la **Creación de Funciones Especializadas sobre la Marcha.** Permite aplicar argumentos de forma incremental a una función general, construyendo operaciones más específicas que se adaptan al flujo de tus datos; ademas facilita la construcción de nuevas funciones a partir de la composición de funciones currificadas existentes, sin necesidad de anidar explícitamente las llamadas.

---

### **¿Por qué es Crucial para nexusL?**

Integrar el currying y el operador pipe en nexusL aporta beneficios significativos:

1. **Claridad Mejorada:** El código se lee como una serie de pasos lógicos y secuenciales, lo que reduce la carga cognitiva y facilita el mantenimiento.  
2. **Modularidad Intrínseca:** Fomenta la creación de funciones pequeñas, puras y reutilizables que pueden combinarse de innumerables maneras.  
3. **Expresividad Elevada:** Permite a los usuarios de nexusL expresar lógica compleja de transformación de datos o de flujo de control de una manera concisa e intuitiva.  
4. **Menos errores:** Al reducir la anidación y hacer el flujo de datos explícito, se disminuyen las posibilidades de cometer errores de encadenamiento.

Al proporcionar estas herramientas en nexusL, no solo se está ofreciendo "azúcar sintáctica", sino que estamos habilitando un paradigma de programación que promueve la **composición funcional, la inmutabilidad y la legibilidad**, haciendo que el lenguaje sea más potente y agradable de usar.