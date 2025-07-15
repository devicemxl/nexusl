/// Macro para generar tipos numéricos newtype con operadores completos
/// 
/// Esta macro genera automáticamente todo el boilerplate necesario para crear
/// un tipo newtype que se comporta como un tipo numérico primitivo, pero con
/// seguridad de tipos adicional.
/// 
/// # Características generadas:
/// - Estructura newtype con el tipo subyacente
/// - Operadores aritméticos básicos (+, -, *, /, %, -)
/// - Operadores de asignación (+=, -=, *=, /=, %=)
/// - Comparaciones (==, !=, <, <=, >, >=)
/// - Traits estándar (Debug, Clone, Copy, Display)
/// - Conversiones (From, Into)
/// - Métodos de acceso (new, value, as_ref)
/// 
/// # Parámetros:
/// - `$name`: Nombre del nuevo tipo a crear
/// - `$inner_type`: Tipo primitivo subyacente (i32, f32, etc.)
/// - `$variant`: Especifica si es "int" (entero) o "float" (flotante)
/// 
/// # Diferencias entre variantes:
/// - **int**: Implementa `Eq` y `Ord` completos (comparación total)
/// - **float**: Solo implementa `PartialEq` y `PartialOrd` (comparación parcial)
/// 
/// # Ejemplo:
/// ```rust
/// // Definir un tipo para representar kilómetros
/// define_numeric_type!(Kilometers, f32, float);
/// 
/// // Definir un tipo para contar elementos
/// define_numeric_type!(Count, i32, int);
/// 
/// let distance = Kilometers::new(5.5);
/// let count = Count::new(10);
/// ```
macro_rules! define_numeric_type {
    // Variante para tipos enteros con signo (implementa Eq y Ord)
    ($name:ident, $inner_type:ty, signed) => {
        /// Tipo newtype que encapsula un valor numérico entero con signo
        /// 
        /// Este tipo proporciona seguridad de tipos mientras mantiene
        /// todas las operaciones numéricas del tipo subyacente.
        #[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash)]
        pub struct $name($inner_type);
        
        impl $name {
            /// Crea una nueva instancia del tipo
            /// 
            /// # Argumentos
            /// * `value` - El valor del tipo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let count = Count::new(42);
            /// ```
            pub fn new(value: $inner_type) -> Self {
                $name(value)
            }
            
            /// Extrae el valor subyacente por valor
            /// 
            /// # Retorna
            /// El valor del tipo primitivo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let count = Count::new(42);
            /// let inner: i32 = count.value();
            /// ```
            pub fn value(self) -> $inner_type {
                self.0
            }
            
            /// Obtiene una referencia al valor subyacente
            /// 
            /// # Retorna
            /// Una referencia al valor del tipo primitivo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let count = Count::new(42);
            /// let inner_ref: &i32 = count.as_ref();
            /// ```
            pub fn as_ref(&self) -> &$inner_type {
                &self.0
            }
            
            /// Obtiene una referencia mutable al valor subyacente
            /// 
            /// # Retorna
            /// Una referencia mutable al valor del tipo primitivo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let mut count = Count::new(42);
            /// *count.as_mut() = 100;
            /// ```
            pub fn as_mut(&mut self) -> &mut $inner_type {
                &mut self.0
            }
            
            /// Convierte el tipo a su valor absoluto
            /// 
            /// # Retorna
            /// Un nuevo valor con el valor absoluto
            /// 
            /// # Ejemplo
            /// ```rust
            /// let negative = Count::new(-42);
            /// let positive = negative.abs();
            /// ```
            pub fn abs(self) -> Self {
                $name(self.0.abs())
            }
        }
        
        // Generar implementaciones comunes (incluye negación)
        define_numeric_type!(@impl_common $name, $inner_type, with_neg);
    };
    
    // Variante para tipos enteros sin signo (implementa Eq y Ord)
    ($name:ident, $inner_type:ty, unsigned) => {
        /// Tipo newtype que encapsula un valor numérico entero sin signo
        /// 
        /// Este tipo proporciona seguridad de tipos mientras mantiene
        /// todas las operaciones numéricas del tipo subyacente.
        /// 
        /// Nota: Los tipos sin signo no soportan negación unaria ni abs()
        /// ya que siempre son positivos.
        #[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash)]
        pub struct $name($inner_type);
        
        impl $name {
            /// Crea una nueva instancia del tipo
            /// 
            /// # Argumentos
            /// * `value` - El valor del tipo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let id = UserId::new(42);
            /// ```
            pub fn new(value: $inner_type) -> Self {
                $name(value)
            }
            
            /// Extrae el valor subyacente por valor
            /// 
            /// # Retorna
            /// El valor del tipo primitivo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let id = UserId::new(42);
            /// let inner: u32 = id.value();
            /// ```
            pub fn value(self) -> $inner_type {
                self.0
            }
            
            /// Obtiene una referencia al valor subyacente
            /// 
            /// # Retorna
            /// Una referencia al valor del tipo primitivo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let id = UserId::new(42);
            /// let inner_ref: &u32 = id.as_ref();
            /// ```
            pub fn as_ref(&self) -> &$inner_type {
                &self.0
            }
            
            /// Obtiene una referencia mutable al valor subyacente
            /// 
            /// # Retorna
            /// Una referencia mutable al valor del tipo primitivo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let mut id = UserId::new(42);
            /// *id.as_mut() = 100;
            /// ```
            pub fn as_mut(&mut self) -> &mut $inner_type {
                &mut self.0
            }
            
            /// Verifica si el valor es par
            /// 
            /// # Retorna
            /// `true` si el valor es par, `false` si es impar
            /// 
            /// # Ejemplo
            /// ```rust
            /// let id = UserId::new(42);
            /// assert!(id.is_even());
            /// ```
            pub fn is_even(self) -> bool {
                self.0 % 2 == 0
            }
            
            /// Verifica si el valor es impar
            /// 
            /// # Retorna
            /// `true` si el valor es impar, `false` si es par
            /// 
            /// # Ejemplo
            /// ```rust
            /// let id = UserId::new(43);
            /// assert!(id.is_odd());
            /// ```
            pub fn is_odd(self) -> bool {
                self.0 % 2 != 0
            }
            
            /// Obtiene el siguiente valor (incrementa en 1)
            /// 
            /// # Retorna
            /// Un nuevo valor incrementado en 1
            /// 
            /// # Ejemplo
            /// ```rust
            /// let id = UserId::new(42);
            /// let next = id.next();
            /// ```
            pub fn next(self) -> Self {
                $name(self.0 + 1)
            }
            
            /// Obtiene el valor anterior (decrementa en 1)
            /// 
            /// # Retorna
            /// Un nuevo valor decrementado en 1
            /// 
            /// # Panics
            /// Si el valor actual es 0, causará un panic por overflow
            /// 
            /// # Ejemplo
            /// ```rust
            /// let id = UserId::new(42);
            /// let prev = id.prev();
            /// ```
            pub fn prev(self) -> Self {
                $name(self.0 - 1)
            }
            
            /// Obtiene el valor anterior de forma segura
            /// 
            /// # Retorna
            /// `Some(nuevo_valor)` si no hay overflow, `None` si el valor actual es 0
            /// 
            /// # Ejemplo
            /// ```rust
            /// let id = UserId::new(42);
            /// let prev = id.checked_prev().unwrap();
            /// 
            /// let zero = UserId::new(0);
            /// assert_eq!(zero.checked_prev(), None);
            /// ```
            pub fn checked_prev(self) -> Option<Self> {
                self.0.checked_sub(1).map(|v| $name(v))
            }
        }
        
        // Generar implementaciones comunes (sin negación)
        define_numeric_type!(@impl_common $name, $inner_type, no_neg);
    };
    
    // Variante para tipos flotantes (solo PartialEq y PartialOrd)
    ($name:ident, $inner_type:ty, float) => {
        /// Tipo newtype que encapsula un valor numérico flotante
        /// 
        /// Este tipo proporciona seguridad de tipos mientras mantiene
        /// todas las operaciones numéricas del tipo subyacente.
        /// 
        /// Nota: Los tipos flotantes no implementan `Eq` ni `Ord` debido
        /// a la presencia de valores especiales como NaN.
        #[derive(Debug, Clone, Copy, PartialEq, PartialOrd)]
        pub struct $name($inner_type);
        
        impl $name {
            /// Crea una nueva instancia del tipo
            /// 
            /// # Argumentos
            /// * `value` - El valor del tipo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let distance = Kilometers::new(5.5);
            /// ```
            pub fn new(value: $inner_type) -> Self {
                $name(value)
            }
            
            /// Extrae el valor subyacente por valor
            /// 
            /// # Retorna
            /// El valor del tipo primitivo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let distance = Kilometers::new(5.5);
            /// let inner: f32 = distance.value();
            /// ```
            pub fn value(self) -> $inner_type {
                self.0
            }
            
            /// Obtiene una referencia al valor subyacente
            /// 
            /// # Retorna
            /// Una referencia al valor del tipo primitivo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let distance = Kilometers::new(5.5);
            /// let inner_ref: &f32 = distance.as_ref();
            /// ```
            pub fn as_ref(&self) -> &$inner_type {
                &self.0
            }
            
            /// Obtiene una referencia mutable al valor subyacente
            /// 
            /// # Retorna
            /// Una referencia mutable al valor del tipo primitivo subyacente
            /// 
            /// # Ejemplo
            /// ```rust
            /// let mut distance = Kilometers::new(5.5);
            /// *distance.as_mut() = 10.0;
            /// ```
            pub fn as_mut(&mut self) -> &mut $inner_type {
                &mut self.0
            }
            
            /// Convierte el tipo a su valor absoluto
            /// 
            /// # Retorna
            /// Un nuevo valor con el valor absoluto
            /// 
            /// # Ejemplo
            /// ```rust
            /// let negative = Kilometers::new(-5.5);
            /// let positive = negative.abs();
            /// ```
            pub fn abs(self) -> Self {
                $name(self.0.abs())
            }
            
            /// Verifica si el valor es finito (no infinito ni NaN)
            /// 
            /// # Retorna
            /// `true` si el valor es finito, `false` en caso contrario
            /// 
            /// # Ejemplo
            /// ```rust
            /// let distance = Kilometers::new(5.5);
            /// assert!(distance.is_finite());
            /// ```
            pub fn is_finite(self) -> bool {
                self.0.is_finite()
            }
            
            /// Verifica si el valor es NaN (Not a Number)
            /// 
            /// # Retorna
            /// `true` si el valor es NaN, `false` en caso contrario
            /// 
            /// # Ejemplo
            /// ```rust
            /// let distance = Kilometers::new(f32::NAN);
            /// assert!(distance.is_nan());
            /// ```
            pub fn is_nan(self) -> bool {
                self.0.is_nan()
            }
            
            /// Redondea el valor al entero más cercano
            /// 
            /// # Retorna
            /// Un nuevo valor con el valor redondeado
            /// 
            /// # Ejemplo
            /// ```rust
            /// let distance = Kilometers::new(5.7);
            /// let rounded = distance.round();
            /// ```
            pub fn round(self) -> Self {
                $name(self.0.round())
            }
            
            /// Redondea hacia abajo (floor)
            /// 
            /// # Retorna
            /// Un nuevo valor con el valor redondeado hacia abajo
            /// 
            /// # Ejemplo
            /// ```rust
            /// let distance = Kilometers::new(5.7);
            /// let floored = distance.floor();
            /// ```
            pub fn floor(self) -> Self {
                $name(self.0.floor())
            }
            
            /// Redondea hacia arriba (ceil)
            /// 
            /// # Retorna
            /// Un nuevo valor con el valor redondeado hacia arriba
            /// 
            /// # Ejemplo
            /// ```rust
            /// let distance = Kilometers::new(5.3);
            /// let ceiled = distance.ceil();
            /// ```
            pub fn ceil(self) -> Self {
                $name(self.0.ceil())
            }
        }
        
        // Generar implementaciones comunes (con negación para flotantes)
        define_numeric_type!(@impl_common $name, $inner_type, with_neg);
    };
    
    // Regla interna para implementaciones comunes CON negación
    (@impl_common $name:ident, $inner_type:ty, with_neg) => {
        // =====================================================
        // OPERADORES ARITMÉTICOS BÁSICOS
        // =====================================================
        
        /// Implementa la suma entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(5);
        /// let b = MyType::new(3);
        /// let result = a + b; // MyType(8)
        /// ```
        impl std::ops::Add for $name {
            type Output = Self;
            
            fn add(self, other: Self) -> Self {
                $name(self.0 + other.0)
            }
        }
        
        /// Implementa la resta entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(5);
        /// let b = MyType::new(3);
        /// let result = a - b; // MyType(2)
        /// ```
        impl std::ops::Sub for $name {
            type Output = Self;
            
            fn sub(self, other: Self) -> Self {
                $name(self.0 - other.0)
            }
        }
        
        /// Implementa la multiplicación entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(5);
        /// let b = MyType::new(3);
        /// let result = a * b; // MyType(15)
        /// ```
        impl std::ops::Mul for $name {
            type Output = Self;
            
            fn mul(self, other: Self) -> Self {
                $name(self.0 * other.0)
            }
        }
        
        /// Implementa la división entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(10);
        /// let b = MyType::new(2);
        /// let result = a / b; // MyType(5)
        /// ```
        impl std::ops::Div for $name {
            type Output = Self;
            
            fn div(self, other: Self) -> Self {
                $name(self.0 / other.0)
            }
        }
        
        /// Implementa el módulo (resto) entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(10);
        /// let b = MyType::new(3);
        /// let result = a % b; // MyType(1)
        /// ```
        impl std::ops::Rem for $name {
            type Output = Self;
            
            fn rem(self, other: Self) -> Self {
                $name(self.0 % other.0)
            }
        }
        
        /// Implementa la negación unaria (cambio de signo)
        /// 
        /// # Ejemplo
        /// ```rust
        /// let positive = MyType::new(5);
        /// let negative = -positive; // MyType(-5)
        /// ```
        impl std::ops::Neg for $name {
            type Output = Self;
            
            fn neg(self) -> Self {
                $name(-self.0)
            }
        }
        
        // Generar el resto de implementaciones comunes
        define_numeric_type!(@impl_operators $name, $inner_type);
    };
    
    // Regla interna para implementaciones comunes SIN negación
    (@impl_common $name:ident, $inner_type:ty, no_neg) => {
        // =====================================================
        // OPERADORES ARITMÉTICOS BÁSICOS (SIN NEGACIÓN)
        // =====================================================
        
        /// Implementa la suma entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(5);
        /// let b = MyType::new(3);
        /// let result = a + b; // MyType(8)
        /// ```
        impl std::ops::Add for $name {
            type Output = Self;
            
            fn add(self, other: Self) -> Self {
                $name(self.0 + other.0)
            }
        }
        
        /// Implementa la resta entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(5);
        /// let b = MyType::new(3);
        /// let result = a - b; // MyType(2)
        /// ```
        impl std::ops::Sub for $name {
            type Output = Self;
            
            fn sub(self, other: Self) -> Self {
                $name(self.0 - other.0)
            }
        }
        
        /// Implementa la multiplicación entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(5);
        /// let b = MyType::new(3);
        /// let result = a * b; // MyType(15)
        /// ```
        impl std::ops::Mul for $name {
            type Output = Self;
            
            fn mul(self, other: Self) -> Self {
                $name(self.0 * other.0)
            }
        }
        
        /// Implementa la división entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(10);
        /// let b = MyType::new(2);
        /// let result = a / b; // MyType(5)
        /// ```
        impl std::ops::Div for $name {
            type Output = Self;
            
            fn div(self, other: Self) -> Self {
                $name(self.0 / other.0)
            }
        }
        
        /// Implementa el módulo (resto) entre dos valores del mismo tipo
        /// 
        /// # Ejemplo
        /// ```rust
        /// let a = MyType::new(10);
        /// let b = MyType::new(3);
        /// let result = a % b; // MyType(1)
        /// ```
        impl std::ops::Rem for $name {
            type Output = Self;
            
            fn rem(self, other: Self) -> Self {
                $name(self.0 % other.0)
            }
        }
        
        // Generar el resto de implementaciones comunes
        define_numeric_type!(@impl_operators $name, $inner_type);
    };
    
    // Regla interna para operadores de asignación y otras implementaciones comunes
    (@impl_operators $name:ident, $inner_type:ty) => {
        // =====================================================
        // OPERADORES DE ASIGNACIÓN COMPUESTOS
        // =====================================================
        
        /// Implementa la suma con asignación (+=)
        /// 
        /// # Ejemplo
        /// ```rust
        /// let mut a = MyType::new(5);
        /// a += MyType::new(3); // a es ahora MyType(8)
        /// ```
        impl std::ops::AddAssign for $name {
            fn add_assign(&mut self, other: Self) {
                self.0 += other.0;
            }
        }
        
        /// Implementa la resta con asignación (-=)
        /// 
        /// # Ejemplo
        /// ```rust
        /// let mut a = MyType::new(5);
        /// a -= MyType::new(3); // a es ahora MyType(2)
        /// ```
        impl std::ops::SubAssign for $name {
            fn sub_assign(&mut self, other: Self) {
                self.0 -= other.0;
            }
        }
        
        /// Implementa la multiplicación con asignación (*=)
        /// 
        /// # Ejemplo
        /// ```rust
        /// let mut a = MyType::new(5);
        /// a *= MyType::new(3); // a es ahora MyType(15)
        /// ```
        impl std::ops::MulAssign for $name {
            fn mul_assign(&mut self, other: Self) {
                self.0 *= other.0;
            }
        }
        
        /// Implementa la división con asignación (/=)
        /// 
        /// # Ejemplo
        /// ```rust
        /// let mut a = MyType::new(10);
        /// a /= MyType::new(2); // a es ahora MyType(5)
        /// ```
        impl std::ops::DivAssign for $name {
            fn div_assign(&mut self, other: Self) {
                self.0 /= other.0;
            }
        }
        
        /// Implementa el módulo con asignación (%=)
        /// 
        /// # Ejemplo
        /// ```rust
        /// let mut a = MyType::new(10);
        /// a %= MyType::new(3); // a es ahora MyType(1)
        /// ```
        impl std::ops::RemAssign for $name {
            fn rem_assign(&mut self, other: Self) {
                self.0 %= other.0;
            }
        }
        
        // =====================================================
        // CONVERSIONES Y TRAITS DE UTILIDAD
        // =====================================================
        
        /// Permite convertir desde el tipo primitivo subyacente
        /// 
        /// # Ejemplo
        /// ```rust
        /// let value: MyType = 42.into();
        /// let value = MyType::from(42);
        /// ```
        impl From<$inner_type> for $name {
            fn from(value: $inner_type) -> Self {
                $name(value)
            }
        }
        
        /// Permite convertir hacia el tipo primitivo subyacente
        /// 
        /// # Ejemplo
        /// ```rust
        /// let my_type = MyType::new(42);
        /// let primitive: $inner_type = my_type.into();
        /// ```
        impl Into<$inner_type> for $name {
            fn into(self) -> $inner_type {
                self.0
            }
        }
        
        /// Permite usar el tipo en contextos donde se espera una referencia
        /// al tipo primitivo subyacente
        /// 
        /// # Ejemplo
        /// ```rust
        /// let my_type = MyType::new(42);
        /// let reference: &$inner_type = my_type.as_ref();
        /// ```
        impl AsRef<$inner_type> for $name {
            fn as_ref(&self) -> &$inner_type {
                &self.0
            }
        }
        
        /// Permite usar el tipo en contextos donde se espera una referencia
        /// mutable al tipo primitivo subyacente
        /// 
        /// # Ejemplo
        /// ```rust
        /// let mut my_type = MyType::new(42);
        /// let reference: &mut $inner_type = my_type.as_mut();
        /// ```
        impl AsMut<$inner_type> for $name {
            fn as_mut(&mut self) -> &mut $inner_type {
                &mut self.0
            }
        }
        
        /// Implementa el trait Display para permitir imprimir el tipo
        /// de manera legible
        /// 
        /// # Ejemplo
        /// ```rust
        /// let my_type = MyType::new(42);
        /// println!("{}", my_type); // Imprime: 42
        /// ```
        impl std::fmt::Display for $name {
            fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(f, "{}", self.0)
            }
        }
        
        /// Implementa Default si el tipo subyacente lo implementa
        /// 
        /// # Ejemplo
        /// ```rust
        /// let default_value = MyType::default();
        /// ```
        impl Default for $name {
            fn default() -> Self {
                $name(Default::default())
            }
        }
    };
}

// =====================================================
// EJEMPLOS DE USO
// =====================================================

// Definir tipos para diferentes unidades de medida
define_numeric_type!(Meters, f32, float);
define_numeric_type!(Kilometers, f32, float);
define_numeric_type!(Seconds, f32, float);

// Definir tipos para contadores y IDs
define_numeric_type!(UserId, u32, unsigned);
define_numeric_type!(Count, i32, signed);
define_numeric_type!(Score, i32, signed);

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_basic_operations() {
        let a = Meters::new(10.0);
        let b = Meters::new(5.0);
        
        // Operaciones aritméticas
        assert_eq!(a + b, Meters::new(15.0));
        assert_eq!(a - b, Meters::new(5.0));
        assert_eq!(a * b, Meters::new(50.0));
        assert_eq!(a / b, Meters::new(2.0));
        
        // Operaciones de asignación
        let mut c = a;
        c += b;
        assert_eq!(c, Meters::new(15.0));
        
        // Comparaciones
        assert!(a > b);
        assert!(b < a);
        assert_eq!(a, Meters::new(10.0));
    }
    
    #[test]
    fn test_conversions() {
        let meters = Meters::new(100.0);
        let value: f32 = meters.into();
        assert_eq!(value, 100.0);
        
        let meters2: Meters = 200.0.into();
        assert_eq!(meters2, Meters::new(200.0));
    }
    
    #[test]
    fn test_integer_types() {
        // Test para tipos con signo
        let score1 = Score::new(10);
        let score2 = Score::new(-5);
        
        assert_eq!(score1 + score2, Score::new(5));
        assert_eq!(score2.abs(), Score::new(5));
        assert_eq!(-score1, Score::new(-10));
        
        // Test para tipos sin signo
        let id1 = UserId::new(1);
        let id2 = UserId::new(2);
        
        assert!(id2 > id1);
        assert_eq!(id1 + id2, UserId::new(3));
        assert!(id2.is_even());
        assert!(id1.is_odd());
        assert_eq!(id2.next(), UserId::new(3));
        assert_eq!(id2.prev(), UserId::new(1));
        
        // Test para métodos específicos de unsigned
        let zero = UserId::new(0);
        assert_eq!(zero.checked_prev(), None);
        assert_eq!(id1.checked_prev(), Some(UserId::new(0)));
        
        let mut count = Count::new(0);
        count += Count::new(1);
        assert_eq!(count, Count::new(1));
    }
    
    #[test]
    fn test_float_specific_methods() {
        let distance = Kilometers::new(5.7);
        
        assert!(distance.is_finite());
        assert!(!distance.is_nan());
        assert_eq!(distance.round(), Kilometers::new(6.0));
        assert_eq!(distance.floor(), Kilometers::new(5.0));
        assert_eq!(distance.ceil(), Kilometers::new(6.0));
        
        let negative = Kilometers::new(-3.2);
        assert_eq!(negative.abs(), Kilometers::new(3.2));
    }
    
    #[test]
    fn test_display_and_debug() {
        let meters = Meters::new(42.5);
        assert_eq!(format!("{}", meters), "42.5");
        assert_eq!(format!("{:?}", meters), "Meters(42.5)");
    }
    
    #[test]
    fn test_type_safety() {
        let meters = Meters::new(100.0);
        let kilometers = Kilometers::new(1.0);
        let user_id = UserId::new(42);
        let score = Score::new(-10);
        
        // Esto NO compilaría, lo cual es exactamente lo que queremos:
        // let result = meters + kilometers; // Error de compilación
        // let bad_mix = user_id + score; // Error de compilación
        
        // Pero puedes convertir explícitamente si es necesario:
        let meters_from_km = Meters::new(kilometers.value() * 1000.0);
        let total = meters + meters_from_km;
        assert_eq!(total, Meters::new(1100.0));
        
        // Tests específicos para tipos con/sin signo
        assert_eq!(score.abs(), Score::new(10));
        assert_eq!(-score, Score::new(10));
        
        // Los unsigned no tienen negación ni abs
        assert!(user_id.is_even());
        assert_eq!(user_id.next(), UserId::new(43));
    }
}

/// Ejemplo de uso principal
fn main() {
    println!("=== Ejemplo de Tipos Numéricos con Macros ===\n");
    
    // Crear diferentes tipos de medidas
    let distance_m = Meters::new(1500.0);
    let distance_km = Kilometers::new(2.5);
    let time = Seconds::new(300.0);
    
    println!("Distancia en metros: {}", distance_m);
    println!("Distancia en kilómetros: {}", distance_km);
    println!("Tiempo en segundos: {}", time);
    
    // Operaciones aritméticas
    let double_distance = distance_m * Meters::new(2.0);
    println!("Distancia duplicada: {}", double_distance);
    
    // Tipos enteros
    let user1 = UserId::new(1001);
    let user2 = UserId::new(1002);
    let total_users = Count::new(50);
    
    println!("\nUsuario 1: {}", user1);
    println!("Usuario 2: {}", user2);
    println!("Total de usuarios: {}", total_users);
    
    // Operaciones con enteros
    let mut score = Score::new(100);
    score += Score::new(50);
    println!("Puntuación final: {}", score);
    
    // Métodos específicos para flotantes
    let precise_distance = Kilometers::new(5.789);
    println!("\nDistancia precisa: {}", precise_distance);
    println!("Redondeada: {}", precise_distance.round());
    println!("Hacia abajo: {}", precise_distance.floor());
    println!("Hacia arriba: {}", precise_distance.ceil());
    
    // Comparaciones
    if distance_m > Meters::new(1000.0) {
        println!("La distancia es mayor a 1000 metros");
    }
    
    // Conversiones
    let raw_value: f32 = distance_km.into();
    println!("Valor crudo de kilómetros: {}", raw_value);
    
    println!("\n=== Fin del ejemplo ===");
}