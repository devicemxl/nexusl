use uuid::Uuid

use crate::types::define_numeric_type; 

//pub TripletScope: string


/*
The Header struct from the triplet object contains the metadata for the
triplet. All this data is necessary for use at the KV value and other
persistent storages.
*/
pub struct Header {
    pub IdH: Uuid,
    pub TripletScope: u8, // unsigned 8-bit integer - fact, rule, def, proc, prog, macro, template,...
    pub Content: Body,  // triplet: fact, rule, def, proc, prog, macro, template,...
    pub Hash: tailHash // bit sign -> Security layer
}
/*
The Body struct from the triplet object contains the triplet in self and the IdH.
*/
pub struct Body {
    pub IdH: Uuid,
    pub Subject: Symbol, // 
    pub Verb: ActiveWord,  // aka Verb
    pub Attr: Any,  // verb, attrb, aobject, symbol...
}
/*
The Tail struct from the triplet object contains the Exportable data and the IdH,
this adds an additional layer of security, because you can hack the function but
need to change this block too
*/
pub struct Tail {
    pub IdH: Uuid,
    pub Export: NDArray<T>,  // Only Apply for proc, nil for any other
}

fn TestFunc {
    
pub struct Persona {
    pub nombre: String,
    pub edad: u32, // unsigned 32-bit integer
    pub ciudad: String,
    }
}
// Bloque de implementación para la struct Persona
impl Persona {
    // Función asociada (constructor, no toma `self`)
    // En Rust, los "constructores" suelen ser funciones asociadas llamadas `new`
    pub fn new(nombre: String, edad: u32, ciudad: String) -> Self {
        Persona { nombre, edad, ciudad }
    }

    // Método de instancia (toma `&self` o `&mut self`)
    // Equivale al método `Saludar` de Go
    pub fn saludar(&self) -> String {
        format!("Hola, mi nombre es {} y tengo {} años.", self.nombre, self.edad)
    }

    // Método de instancia que modifica el struct (toma `&mut self`)
    // Equivale al método `CumplirAnos` de Go
    pub fn cumplir_anos(&mut self) {
        self.edad += 1;
    }

    // Método que consume el struct (toma `self`)
    pub fn despedirse(self) {
        println!("Adiós de parte de {}", self.nombre);
        // `self` es movido aquí, no se puede usar `self` después de esta llamada
    }
}

fn main() {
    // Crear una instancia de Persona
    let mut p1 = Persona::new(String::from("Alice"), 30, String::from("Nueva York"));
    println!("{}", p1.saludar()); // Imprime: Hola, mi nombre es Alice y tengo 30 años.

    p1.cumplir_anos();
    println!("Nueva edad: {}", p1.edad); // Imprime: Nueva edad: 31

    // p1.despedirse(); // Si descomentas esto, p1 ya no se puede usar después
    // println!("{}", p1.nombre); // Error: uso de valor movido
}