### 🧠 `trunKV` - MVP Merge Semántico con A-MEM

Especificación funcional para una **variación semántica del flujo de merge**, incorporando una **memoria semántica dinámica (tipo A-MEM)** en el sistema `trunKV`, en el contexto de un diseño *sea of semantic nodes*.

#### 🎯 Propósito:

Integrar dos ramas de conocimiento semántico considerando no solo las diferencias estructurales, sino también la *coherencia contextual* de los triplets, aprovechando una memoria semántica dinámica (tipo A-MEM) para resolver conflictos basados en significado y uso reciente.

> *“Fusionar ramas semánticas considerando contexto activo, peso de uso reciente y cercanía semántica, no solo igualdad textual.”*

---

### 1. **`semanticMerge(base, left, right)`**

* **Función:** `semanticMerge(base, left, right) → mergedCommit`
* **Acciones:**

  * Cargar árboles `base`, `left` y `right`.
  * Extraer cambios relativos al `base` en ambas ramas.
  * Invocar `resolveSemanticConflicts` para tratar conflictos significativos.
  * Generar nuevo árbol `merged` con triplets actualizados, normalizados y pesados semánticamente.
  * Crear nuevo `commit` apuntando a `merged`.

---

### 2. **`resolveSemanticConflicts(conflicts, aMemContext)`**

* **Función:** `resolveSemanticConflicts(conflicts, aMemContext) → resolvedSet`
* **Acciones:**

  * Recorrer `conflicts` identificados (e.g. mismo sujeto con `has:X` vs `has:Y`).
  * Consultar `aMemContext` (estructura de activación semántica):

    * Historial de uso reciente.
    * Peso semántico (embedding similarity o frecuencia en activaciones).
  * Elegir la versión con mayor coherencia contextual.
  * Marcar elementos ambiguos como `ambiguous` si no hay contexto suficiente.

> 🧠 *Nota:* A-MEM puede estar implementado como una cache semántica LRU ponderada con puntuación por similitud y acceso.

---

### 3. **`updateAMemContext(mergedTriplets)`**

* **Función:** `updateAMemContext(newTriplets)`
* **Acciones:**

  * Añadir los nuevos triplets del merge al contexto activo.
  * Recalcular pesos semánticos y actualizar el índice de activación.
  * Mantener tamaño acotado según política de retención.

---

### 🔁 Flujo Esperado

```text
diff(base, left) + diff(base, right)
      ↓
resolverSemanticConflicts()
      ↓
mergeTriplets → updateAMemContext()
      ↓
commit(mergedTree)
```

---

### 🧩 Componentes clave

| Componente          | Descripción breve                                  |
| ------------------- | -------------------------------------------------- |
| `triplet store`     | Almacén de hechos versionados por commit           |
| `embedding index`   | Similaridad semántica basada en vectores           |
| `aMemContext`       | Memoria dinámica que guarda activaciones recientes |
| `semantic diff`     | Detección de cambios con interpretación contextual |
| `semantic resolver` | Lógica de resolución de conflictos basada en A-MEM |
| `commit log`        | Registro de versiones semánticas                   |

---

### ✅ Buenas prácticas

* Resolver diferencias antes del merge estructural.
* Normalizar los triplets antes de evaluar conflicto semántico.
* Usar embeddings y activación previa como guía en conflictos.
* Generar anotaciones sobre decisiones del merge (rastro de decisión semántica).
* Mantener `aMemContext` aislado del histórico de commits (es volátil y contextual).

