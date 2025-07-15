import std/tables
import std/math
import std/options
import std/hashes

# --- DECLARACIÓN ADELANTADA DE AnyValue ---
# Esto le dice a Nim que "AnyValue" existe, pero su definición completa vendrá después.
# Debe ser un tipo de referencia para romper el ciclo de dependencia.
type AnyValue* = ref object 
type Number*   = ref object 

#[
================================================================================
                               CORE TYPES
================================================================================

Definición de tipos básicos para el DSL de transpilación a Nim.

Estos tipos forman la base del sistema de tipos del DSL y facilitan la 
transpilación a código Nim idiomático. Cada tipo aquí tiene un mapeo directo
con tipos nativos de Nim.

Pipeline: DSL → Lexer/Parser → AST → [ESTOS TIPOS] → Transpilación → Nim

Autor: [Tu nombre]
Fecha: [Fecha]
Versión: 1.0
================================================================================
]#

#[
--------------------------------------------------------------------------------
                            TIPOS NUMÉRICOS BÁSICOS
--------------------------------------------------------------------------------

Aliases para tipos numéricos de Nim que proporcionan:
1. Claridad semántica en el código del DSL
2. Facilidad para cambiar precisión si es necesario
3. Consistencia en la nomenclatura (CamelCase)

Nota: Estos tipos se transpilan directamente a sus equivalentes Nim.
]#

type
  # Tipos de punto flotante
  Double* = float64  ## Precisión doble (64 bits) - equivale a `float64` en Nim
  Simple* = float32  ## Precisión simple (32 bits) - equivale a `float32` en Nim
  
  # Tipos enteros con tamaño específico
  I8*  = int8   ## Entero con signo de 8 bits  - rango: -128 a 127
  I16* = int16  ## Entero con signo de 16 bits - rango: -32,768 a 32,767  
  I32* = int32  ## Entero con signo de 32 bits - rango: -2^31 a 2^31-1
  I64* = int64  ## Entero con signo de 64 bits - rango: -2^63 a 2^63-1

#[
--------------------------------------------------------------------------------
                              TIPOS DE TEXTO
--------------------------------------------------------------------------------

Manejo de cadenas de texto en el DSL.
]#

type
  Text* = string  ## Alias para string de Nim - UTF-8 por defecto

#[
--------------------------------------------------------------------------------
                             TIPOS DE VACUIDAD
--------------------------------------------------------------------------------

Manejo de valores ausentes, indefinidos o desconocidos en el DSL.

Estos tipos utilizan Option[T] de Nim para representar valores que pueden
o no estar presentes, evitando el uso de nil y proporcionando type safety.

IMPORTANTE: Aunque EmptyValue y UnknownValue son técnicamente idénticos,
tienen semánticas diferentes:
#[
- Empty: Representa un valor que intencionalmente está vacío
- Maybe: Representa un valor que existe pero no se conoce aún
]#
]#

type
  ## Valor existente pero desconocido
  Maybe[T] = object
    case hasValue: bool
    of true:
      value: T
    of false:
      discard


proc some[T](value: T): Maybe[T] =
  Maybe[T](hasValue: true, value: value)

template none[T](t: typedesc[T]): Maybe[T] = 
  Maybe[T](hasValue: false)

# Operadores
proc `==`[T](a: Maybe[T], b: Maybe[T]): bool =
  a.hasValue == b.hasValue and 
  (not a.hasValue or a.value == b.value)

# Uso
var x = none(int)
var y = some(42)
var z = Null

if x == none(int):
  echo "x es nulo"
if y == none(int):
  echo "y es nulo"
else:
  echo "y NO es nulo"

Comparable(3)

type
  ## Valor intencionalmente vacío
  Empty*   = Option[Maybe]


#[
--------------------------------------------------------------------------------
                           NOTAS DE IMPLEMENTACIÓN
--------------------------------------------------------------------------------

DECLARACIÓN ADELANTADA DE AnyValue:
- AnyValue se declara como `ref object` al inicio de las definiciones
- Esta declaración adelantada permite romper dependencias circulares
- La definición completa de AnyValue (con `case kind:`) viene después
- NO es necesario redeclarar `type AnyValue = ref object` múltiples veces

RAZÓN PARA ref object:
- Permite recursión en la definición de tipos
- Facilita el manejo de memoria automático
- Evita problemas de copia profunda en estructuras complejas

ALTERNATIVA A nil:
- En lugar de usar `nil` directamente, usamos Option[T]
- Esto proporciona type safety y evita NullPointerExceptions
- El compilador de Nim puede optimizar Option[T] eficientemente
]#

#[
--------------------------------------------------------------------------------
                            EJEMPLOS DE USO
--------------------------------------------------------------------------------

```nim
# Declaración de variables con tipos específicos
let temperatura: Double = 36.5
let edad: I32 = 25
let nombre: Text = "Juan"

# Manejo de valores opcionales
let valorVacio: EmptyValue = none(AnyValue)
let valorDesconocido: UnknownValue = none(AnyValue)

# Verificación de presencia de valores
if valorVacio.isSome():
  echo "El valor está presente: ", valorVacio.get()
else:
  echo "El valor está vacío"
```
]#
#[
* Math Mad Structs
* ==================================== # ==

]#
#
# ^ MATRIX
# ^ ------------
type
    matrix*[T] = ref object # Agregué '*' para exportar el tipo 'matrix' si se importa el módulo
        rows: int # Debe ser un tipo concreto, como 'int'
        cols: int # Debe ser un tipo concreto, como 'int'
        data: seq[T]

proc newMatrix*[T](rows, cols: int): matrix[T] =
    ## Crea una nueva matriz inicializada con valores por defecto.
    assert rows > 0 and cols > 0, "Matrix dimensions must be positive."
    result = matrix[T](rows:rows, cols:cols, data: newSeq[T](rows*cols))

proc `[]`*[T](m: matrix[T], row, col: int): T = # M: matrix[T] para el tipo genérico
    ## Acceso a elementos de la matriz (read).
    assert row >= 0 and row < m.rows and col >= 0 and col < m.cols,
           "matrix: out of limits (row: " & $row & ", col: " & $col & ")"
    result = m.data[row * m.cols + col]

proc `[]=`*[T](m: var matrix[T], row, col: int, value: T) =
    ## Asignación de elementos en la matriz (write).
    assert row >= 0 and row < m.rows and col >= 0 and col < m.cols,
           "matrix: out of limits (row: " & $row & ", col: " & $col & ")"
    m.data[row * m.cols + col] = value
#
# ^ VECTOR
# ^ ------------
# Un vector es un tipo especializado de matriz de una sola columna.
type
    vector*[T] = matrix[T] # Se beneficia de la estructura de matrix

proc newVector*[T](rows: int): vector[T] =
    ## Crea un nuevo vector de 'rows' elementos.
    # Un vector siempre será una matriz de una columna
    result = newMatrix[T](rows, 1) # Aquí creamos la matriz base con 1 columna

proc `[]`*[T](v: vector[T], row: int): T = # v: vector[T] para el tipo genérico
    ## Acceso a elementos del vector (read).
    assert row >= 0 and row < v.rows, "vector: out of limits (row: " & $row & ")"
    # Como es un vector, la columna es siempre 0.
    # El acceso m.data[row * m.cols + col] se simplifica a m.data[row * 1 + 0] = m.data[row]
    # Puedes usar el operador de la matriz base o acceder directamente a data
    result = v.data[row] # Acceso directo a los datos, ya que sabemos que cols es 1

proc `[]=`*[T](v: var vector[T], row: int, value: T) = # v: var vector[T] para el tipo genérico
    ## Asignación de elementos en el vector (write).
    assert row >= 0 and row < v.rows, "vector: out of limits (row: " & $row & ")"
    v.data[row] = value
#[
* Spatial Structs
* ==================================== # ==

]#
# * Point
type
  Point* = ref object
    x*: int  # Exportar para acceso directo
    y*: int
    
# Y añadir constructor
proc newPoint*(x, y: int): Point =
  Point(x: x, y: y)


# 
# ^ Line
# ^ ------------
#
type
  # Una línea definida por un punto de origen y un vector de dirección.
  # El vector de dirección debe ser del mismo tipo numérico que los componentes del punto.
  # Por simplicidad, asumiremos que x, y, z de Point podrían ser float64 para esto.
  # O bien, tu vector[T] podría ser vector[int8] si tus puntos usan int8.
  DirectedLine* = ref object
    origin: Point
    direction: vector[float64] # O vector[int8] si tus puntos lo son

proc newDirectedLine*(originPoint: Point, directionVector: vector[float64]): DirectedLine =
  ## Crea una nueva línea con un punto de origen y un vector de dirección.
  # Aquí podrías validar que el vector de dirección tenga una longitud apropiada (ej. 2 para 2D, 3 para 3D)
  assert directionVector.rows >= 2, "Direction vector must have at least 2 components (x,y)"
  result = DirectedLine(origin: originPoint, direction: directionVector)
# 
# ^ Polyline
# ^ ------------
#
type
  # Una Polyline es una secuencia de puntos, donde cada par de puntos forma un segmento.
  # Podríamos decir que es un 'vector' de 'Point's.
  Polyline* = vector[Point]

proc newPolyline*(numPoints: int): Polyline =
  ## Crea una nueva Polyline vacía con capacidad para `numPoints`.
  ## Esto es un vector de puntos, listos para ser llenados.
  assert numPoints > 0, "Polyline must have at least one Point."
  result = newVector[Point](numPoints)

# --- Métodos Útiles para Polyline ---

proc numSegments*(p: Polyline): int =
  ## Retorna el número de segmentos en la polilínea.
  ## Se requiere al menos 2 puntos para 1 segmento.
  if p.rows < 2:
    return 0
  return p.rows - 1

proc getSegment*(p: Polyline, index: int): DirectedLine =
  ## Obtiene un segmento de línea específico de la polilínea.
  ## `index` se refiere al índice del segmento (0 a numSegments-1).
  assert index >= 0 and index < p.numSegments(), "Polyline segment index out of bounds."

  let
    startPoint = p[index]      # Usamos el operador [] del vector para obtener el punto
    endPoint = p[index + 1]

  # Calculamos el vector de dirección entre los dos puntos.
  # Asumiendo que tus puntos y vectores tienen tipos numéricos compatibles (ej. float64)
  # Necesitas convertir int8 de Point a float64 para el vector de dirección.
  var directionVector = newVector[float64](2) # Para 2D (x,y)
  directionVector[0] = float64(endPoint.x - startPoint.x)
  directionVector[1] = float64(endPoint.y - startPoint.y)

  result = DirectedLine(origin: startPoint, direction: directionVector)
# 
# ^ Perimeter
# ^ ------------
#

## Calcula la distancia euclidiana entre dos puntos.
proc euclideanDistance*(p1, p2: Point): float64 =
  ## Los puntos (x, y) son int8, pero la distancia será un flotante.
  let
    dx = float64(p2.x - p1.x)
    dy = float64(p2.y - p1.y)
  result = sqrt(pow(dx, 2) + pow(dy, 2))

## Calcula el perímetro de una Polyline asumiendo que es un polígono cerrado.
proc perimeter*(p: Polyline): float64 =
  ## Suma la longitud de todos los segmentos, incluyendo el cierre entre el último y el primer punto.
  if p.rows < 2:
    # Un polígono necesita al menos 3 puntos para un área,
    # pero 2 puntos pueden formar una línea con longitud.
    # Para perímetro de una "figura", asumimos 2 o más.
    return 0.0 # O podrías lanzar un error si requieres al menos 3 puntos para un "polígono"

  var totalPerimeter: float64 = 0.0

  # Suma las longitudes de los segmentos consecutivos
  for i in 0 ..< p.rows - 1:
    totalPerimeter += euclideanDistance(p[i], p[i+1])

  # Suma la longitud del segmento que cierra el polígono (del último al primer punto)
  if p.rows > 1: # Solo si hay al menos dos puntos para cerrar
    totalPerimeter += euclideanDistance(p[p.rows - 1], p[0])

  return totalPerimeter
# 
# ^ Polygon
# ^ ------------
#
# Un polígono es una polilínea que se asume cerrada.
# Por simplicidad y reuso, podemos definirlo como un tipo distinto sobre Polyline.
type
  Polygon* = distinct Polyline

proc newPolygon*(numPoints: int): Polygon =
  ## Crea un nuevo polígono con espacio para `numPoints`.
  ## Requiere al menos 3 puntos para ser un polígono no degenerado.
  assert numPoints >= 3, "A polygon must have at least 3 Points."
  # Casteamos explícitamente el resultado de newPolyline a Polygon.
  result = Polygon(newPolyline(numPoints))

# --- Métodos Útiles para Polygon ---

proc `[]`*(p: Polygon, idx: int): Point =
  ## Acceso a un vértice del polígono.
  # Casteamos 'p' de vuelta a Polyline para usar su operador de acceso.
  result = Polyline(p)[idx]

proc `[]=`*(p: var Polygon, idx: int, value: Point) =
  ## Asigna un vértice en el polígono.
  # Casteamos 'p' de vuelta a Polyline para usar su operador de asignación.
  Polyline(p)[idx] = value

proc numVertices*(p: Polygon): int =
  ## Retorna el número de vértices del polígono.
  result = Polyline(p).rows # O también `int(p.base.len)` si Polyline fuera `seq[Point]`

proc getPerimeter*(p: Polygon): float64 =
  ## Calcula el perímetro del polígono.
  ## Reutiliza la función `perimeter` de Polyline, que ya maneja el cierre.
  result = perimeter(Polyline(p))

# --- Opcional: Calcular el área del polígono (usando la fórmula del lazo / Shoelace formula) ---
proc getArea*(p: Polygon): float64 =
  ## Calcula el área de un polígono 2D usando la fórmula del lazo (Shoelace formula).
  ## Asegúrate de que los vértices estén en orden (horario o antihorario).
  ## Retorna el valor absoluto del área.
  if p.numVertices() < 3:
    return 0.0 # Un polígono necesita al menos 3 vértices para tener área.

  var areaSum: float64 = 0.0
  let vertices = Polyline(p) # Acceso a la secuencia de puntos

  for i in 0 ..< vertices.rows:
    let
      p1 = vertices[i]
      # El siguiente punto es (i+1) o el primer punto si es el último vértice
      p2 = if i == vertices.rows - 1: vertices[0] else: vertices[i+1]

    # Fórmula: (x1 * y2) - (x2 * y1)
    areaSum += float64(p1.x) * float64(p2.y)
    areaSum -= float64(p2.x) * float64(p1.y)

  result = abs(areaSum / 2.0)

#
# Lazy - Indeterminate
type
  NumberKind = enum
    VkDouble, 
    VkSimple, 
    VkI8, 
    VkI16, 
    VkI32, 
    VkI64
  # Un objeto variante para representar cualquier valor en tu lista heterogénea
  `obj Number`* =   object # Nombre ficticio para la parte object a la que apunta el ref
    case kind:      NumberKind
    of VkDouble:
      VkDoubleVal:  Double
    of VkSimple:
      VkSimpleVal:  Simple
    of VkI8:
      VkI8Val:      I8
    of VkI16:
      VkI16Val:     I16
    of VkI32:
      VkI32Val:     I32
    of VkI64:
      VkI64Val:     I64



# --- Tu tipo Universal AnyValue (nLiValue) - DEFINICIÓN COMPLETA AL FINAL ---
# Ahora definimos el contenido completo del objeto al que 'AnyValue' (el ref) apunta.
# Notar que la definición es `object` sin `ref`, porque `AnyValue` ya es el `ref`.

type
  ValueKind = enum
    VkNumber
    Vktext, 
    VkBool,
    VkPoint, 
    VkLine, 
    VkPolyline, 
    VkPolygon,
    VkList,       # ¡Reincorporado!
    VkSymbol,     # ¡Reincorporado!
    VkObject,     # ¡Reincorporado!
    VkProc,       # ¡Reincorporado!
    VkTripletRef, # Renombrado para consistencia con 'Ref'
    # Tipos posibles en la lista


  # Aquí definimos la ESTRUCTURA del objeto al que apunta el 'ref AnyValue'
  # Puedes usar un nombre interno como `TAnyValueObj` o simplemente definirlo anónimamente.
  # Para claridad, usaré un nombre anónimo y el tipo AnyValue ya está declarado como ref object.
  # El compilador de Nim "sabe" que AnyValue se refiere a esta definición de object.
  # NOTA: si AnyValue fuera un tipo por valor, la recursión directa no sería posible.
  # La clave es que `AnyValue` (que es el `ref`) *apunta* a esta estructura.
type
  `obj AnyValue`* =     ref object
    case kind*:         ValueKind
    of VkNumber:
      VkNumberVal*:     Number
    of VkText:
      VkTextVal*:       Text
    of VkBool:
      VkBoolVal*:       bool
    of VkPoint:
      VkPointVal*:      Point
    of VkLine:
      VkLineVal*:       DirectedLine
    of VkPolyline:
      VkPolylineVal*:   Polyline
    of VkPolygon:
      VkPolygonVal*:    Polygon
    of VkTripletRef: # Usamos VkTripletRef como en ValueKind
      VkTripletRefVal*: Hash # ¡Volver a Hash es lo ideal para referencias!
    of VkList:
      VkListVal*:       seq[AnyValue]
    of VkSymbol:
      VkSymbolVal*:     Text # O Hash si los símbolos son hasheados
    of VkObject:
      VkObjectVal*:     Table[Text, AnyValue] # Clave Text, Valor AnyValue
    of VkProc:
      VkProcVal*:       Text # O un ID de procedimiento
