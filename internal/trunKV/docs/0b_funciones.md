# especificación de flujo funcional

## 🔧 1. Crear Commit (Guardar el estado actual del árbol)

**Nombre:** `commit(tree, message)`
**Propósito:** Guardar una versión completa del árbol actual.
**Entradas:**

* Estado actual del árbol (nodos y hojas)
* Mensaje opcional de commit

**Flujo:**

1. Serializar el árbol como objeto (`tree`).
2. Guardar cada nodo como `sub-tree` si corresponde.
3. Crear objeto `commit` que apunte al `tree` raíz.
4. Si hay rama activa, actualizar su puntero al nuevo `commit`.
5. Actualizar `HEAD`.

**Salida:**

* Hash del nuevo commit.

### 📌 1. `commit(tree, message)`

| Elemento                | Especificación        |
| ----------------------- | ------------------------------------------------------------------------------------ |
| **Nombre**              | `commit`                                                                             |
| **Descripción**         | Guarda una versión completa y referenciable del árbol actual.                        |
| **Entradas**            | - `tree`: estructura actual (nodos, hojas)  <br> - `message`: texto opcional         |
| **Salidas**             | - `commit_hash`: identificador del nuevo commit                                      |
| **Efectos secundarios** | - Guarda blobs y trees en el almacenamiento <br> - Actualiza puntero de rama activa  |
| **Invariantes**         | - El commit apunta a un único árbol válido <br> - El árbol referenciado debe existir |
| **Errores posibles**    | - Árbol inválido o corrupto <br> - Falta de espacio o acceso de escritura            |

---

## 🔍 2. Inspeccionar Commit (Consultar versión pasada)

**Nombre:** `inspect(commit_hash)`
**Propósito:** Navegar y visualizar el árbol y metadatos de un commit.

**Entradas:**

* Hash del commit

**Flujo:**

1. Buscar y leer objeto `commit`.
2. Obtener el árbol raíz referenciado.
3. Leer recursivamente todos los nodos/subárboles.
4. Mostrar claves, tipos y contenidos.

**Salida:**

* Representación estructurada del árbol (clave/valor)
* Metadatos del commit (mensaje, padres, fecha)

### 🔍 2. `inspect(commit_hash)`

| Elemento                | Especificación                      |
| ----------------------- | ---------------------------- |
| **Nombre**              | `inspect`                                                            |
| **Descripción**         | Despliega la estructura completa del árbol y metadatos de un commit. |
| **Entradas**            | - `commit_hash`: identificador del commit                            |
| **Salidas**             | - Estructura del árbol <br> - Metadatos del commit                   |
| **Efectos secundarios** | Ninguno                                                              |
| **Invariantes**         | - El commit y el árbol apuntado deben existir                        |
| **Errores posibles**    | - Commit no encontrado <br> - Árbol no accesible                     |

---

## 🧯 3. Reset Duro (Volver a un commit anterior)

**Nombre:** `reset_hard(commit_hash)`
**Propósito:** Restaurar el estado del árbol a un commit anterior, eliminando todo lo posterior.

**Entradas:**

* Hash del commit destino

**Flujo:**

1. Leer el objeto `commit` destino.
2. Cargar el árbol completo desde ese commit.
3. Reemplazar el estado actual del árbol con el árbol cargado.
4. Actualizar rama activa para que apunte al commit restaurado.
5. Eliminar cualquier estado temporal o caché que no pertenezca al commit restaurado.

**Salida:**

* Confirmación del reset y hash del nuevo estado activo.

### 🧯 3. `reset_hard(commit_hash)`

| Elemento                | Especificación                           |
| ----------------------- | ---------------------------------------------- |
| **Nombre**              | `reset_hard`                                                               |
| **Descripción**         | Restaura el árbol y rama activa a un estado anterior, descartando cambios. |
| **Entradas**            | - `commit_hash`: commit al que se desea volver                             |
| **Salidas**             | - Nuevo estado del árbol activo                                            |
| **Efectos secundarios** | - Elimina el estado actual del árbol <br> - Cambia puntero de rama activa  |
| **Invariantes**         | - Commit debe existir <br> - La estructura restaurada debe ser coherente   |
| **Errores posibles**    | - Commit inválido o perdido <br> - Árbol corrupto                          |

---

## 🧾 4. Mostrar Diferencias entre Árboles (Diff)

**Nombre:** `diff(commitA, commitB)`
**Propósito:** Identificar diferencias entre dos versiones del árbol.

**Entradas:**

* Dos hashes de commit

**Flujo:**

1. Cargar árboles `treeA` y `treeB` desde los commits.
2. Recorrer recursivamente ambos árboles.
3. Detectar:

   * Archivos agregados
   * Archivos eliminados
   * Archivos modificados
4. Construir un reporte de diferencias.

**Salida:**

* Lista de cambios (`added`, `deleted`, `modified`)

### 🧾 4. `diff(commitA, commitB)`

| Elemento                | Especificación                            |
| ----------------------- | ------------------------------------------ |
| **Nombre**              | `diff`                                                                  |
| **Descripción**         | Muestra las diferencias estructurales y de contenido entre dos árboles. |
| **Entradas**            | - `commitA`, `commitB`: identificadores de commits                      |
| **Salidas**             | - Listado de cambios: `added`, `deleted`, `modified`                    |
| **Efectos secundarios** | Ninguno                                                                 |
| **Invariantes**         | - Ambos commits deben apuntar a árboles válidos                         |
| **Errores posibles**    | - Uno o ambos commits no existen <br> - Error al leer árboles           |

---

## 🔀 5. Merge (Fusionar dos árboles)

**Nombre:** `merge(commitA, commitB, commitBase)`
**Propósito:** Unir los cambios de dos versiones en una nueva, posiblemente con resolución de conflictos.

**Entradas:**

* Commits A y B
* Commit base común (opcional)

**Flujo:**

1. Cargar árboles desde A, B y base.
2. Determinar diferencias A↔base y B↔base.
3. Comparar cada clave:

   * Si cambia solo en una rama: aceptar el cambio.
   * Si cambia en ambas ramas de forma diferente: marcar como conflicto.
4. Construir un árbol resultante.
5. Crear un nuevo commit de merge (si no hay conflictos, o una vez resueltos).

**Salida:**

* Árbol fusionado (y hash)
* Lista de conflictos (si existen)

### 🔀 5. `merge(commitA, commitB, commitBase)`

| Elemento                | Especificación                      |
| ----------------------- | ----------------------------------- |
| **Nombre**              | `merge`                                                                                          |
| **Descripción**         | Fusiona dos versiones del árbol, resolviendo conflictos si es necesario.                         |
| **Entradas**            | - `commitA`, `commitB`: ramas a fusionar <br> - `commitBase`: ancestro común                     |
| **Salidas**             | - Árbol fusionado <br> - Lista de conflictos (si existen) <br> - (opcional) nuevo `commit_merge` |
| **Efectos secundarios** | - Se puede crear un commit nuevo si no hay conflictos                                            |
| **Invariantes**         | - Todos los árboles deben ser accesibles                                                         |
| **Errores posibles**    | - Conflictos irresolubles <br> - Árbol inconsistente tras fusión                                 |

---

## 🌿 6. Ramas (Branch)

**Nombre:** `create_branch(name)`, `switch_branch(name)`, `delete_branch(name)`
**Propósito:** Administrar múltiples punteros a distintas líneas de trabajo.

### `create_branch(name)`

**Entradas:** Nombre de la nueva rama
**Flujo:**

1. Verificar que no exista una rama con ese nombre.
2. Crear una nueva referencia que apunte al commit actual.

**Salida:** Confirmación y commit apuntado.

---

### `switch_branch(name)`

**Entradas:** Nombre de la rama a activar
**Flujo:**

1. Verificar que la rama exista.
2. Actualizar `HEAD` para que apunte a esa rama.
3. Cargar el árbol del commit apuntado por esa rama.

**Salida:** Estado nuevo del árbol.

---

### `delete_branch(name)`

**Entradas:** Nombre de la rama
**Flujo:**

1. Verificar que no sea la rama activa.
2. Eliminar la entrada del registro de ramas.

**Salida:** Confirmación.

#### 🌿 6. `branch` (creación, cambio, eliminación)

| Elemento                | Especificación                     |
| ----------------------- | ---------------------------------- |
| **Nombre**              | `create_branch(name)`, `switch_branch(name)`, `delete_branch(name)`           |
| **Descripción**         | Gestiona múltiples líneas de desarrollo apuntando a commits diferentes.       |
| **Entradas**            | - `name`: nombre de la rama                                                   |
| **Salidas**             | - Confirmación de operación o nuevo estado                                    |
| **Efectos secundarios** | - Cambia `HEAD` en `switch` <br> - Crea/borra punteros en `branches`          |
| **Invariantes**         | - Cada rama apunta a un commit válido <br> - No puede eliminar rama activa    |
| **Errores posibles**    | - Rama ya existe <br> - Rama no encontrada <br> - Intentar borrar rama activa |
