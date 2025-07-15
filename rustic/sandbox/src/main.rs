pub mod types;

use std::collections::HasMap;
use lazy_static::lazy_static;
use crate::types::typeX; 

//use uuid::Uuid;

/*
PRIMERO: HAY QUE RECORDAR QUE ESTO CONTITUYE LA FASE DEL EVALUADOR
        DEL LENGUAJE - LO QUE SIGNIFICA QUE YA PASO LA FASE DEL NEXUSMESH
        (SEA OF NODES) 
        AQUI HAY:
            UN TRIPLET A NIVEL DE PLANTILLA
            SIETE O MAS TRIPLETS A NIVEL DE OBJETOS CON SUS RESPECTIVAS FUNCIONES

Definir "bloques de contexto" o "secciones" en NexusL para diferentes tipos de construcciones
(functions, facts, assertions, algebra, etc.) aporta varios beneficios clave:

Beneficios de los Bloques de Contexto:
Organización y Legibilidad:

Hace el código de NexusL mucho más fácil de leer y entender, especialmente a medida que la base
de conocimiento crece. Un desarrollador puede ver rápidamente qué tipo de declaraciones se
encuentran en cada sección.
facts { ... } deja claro que esa sección contiene solo declaraciones de hechos base.
algebra { ... } indica que se van a definir expresiones matemáticas.

Claridad Semántica y Propósito:

Los bloques sirven como una señal explícita para el transpilador (y para el programador) sobre
el "propósito" de las declaraciones dentro de ellos.
Esto permite al transpilador aplicar diferentes reglas de procesamiento y generación de código
Go según el tipo de bloque.

Extensibilidad y Modularidad del Transpilador:

Es mucho más fácil añadir nuevos tipos de funcionalidad (como álgebra) si ya tienes un mecanismo
de bloques. Simplemente defines un nuevo tipo de bloque y le asignas su propio módulo de manejo
en el transpilador.

Tu parser.go puede tener una función parseFactsBlock(), otra parseAlgebraBlock(), etc.

Posibilidad de Scoping o Namespacing (a futuro):

Aunque Datalog en sí no tiene "scopes" en el sentido de lenguajes imperativos, podrías usar estos
bloques para implementar un tipo de "namespacing" o "módulos".

Por ejemplo, context MySystem { ... } podría implicar que todos los predicados o símbolos definidos
dentro de ese bloque llevan un prefijo implícito (MySystem_some_predicate) en el Datalog generado,
o que pertenecen a un "grafo con nombre" si tu motor Datalog lo soportara (aunque markkurossi/datalog
no lo hace de forma nativa).

*/

// =====================================================
// DEFINITIONS
// =====================================================
//
// TRIPLET SCOPE
// --------------
//
// Definir tipos para diferentes unidades de medida
define_numeric_type!(tripletScope, i8, integer);
const None:tripletScope    = tripletScope::new(0);/*
const Fact:tripletScope    = tripletScope::new(1);
const Def:tripletScope     = tripletScope::new(2);
const Proc:tripletScope    = tripletScope::new(3);
const Prog:tripletScope    = tripletScope::new(4);
const Macro:tripletScope   = tripletScope::new(5);
const Templ:tripletScope   = tripletScope::new(6);
const Expr:tripletScope    = tripletScope::new(7);*/
//
lazy_static!{
    static ref ScopeTable:HasMap <tripletScope, & 'static string > = {
        let mut thisHash = HashMap.new();
        thisHash.insert!( None, "None");
    }
}
//ScopeTable::lock().unwraph();
/*
ScopeTable::insert!( Fact,    "fact");
ScopeTable::insert!( Def,     "def");
ScopeTable::insert!( Proc,    "Proc");
ScopeTable::insert!( Prog,    "Script");
ScopeTable::insert!( Macro,   "Macro");
ScopeTable::insert!( Templ,   "Template");
ScopeTable::insert!( Expr,    "Expr");*/
/*
The Header struct from the triplet object contains the metadata for the
triplet. All this data is necessary for use at the KV value and other
persistent storages.
*/
pub struct Header {
    pub IdH: String, //change to Uuid
    pub TripletScope: tripletScope, // unsigned 8-bit integer - fact, rule, def, proc, prog, macro, template,...
    pub Content: String,  // triplet: fact, rule, def, proc, prog, macro, template,...
    pub Hash: String // tailHash - bit sign -> Security layer
}

fn main() {
    println!("Hello, world!");
}
