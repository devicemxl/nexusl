## 🧠 `trunKV` - MVP *Sea of Semantic Nodes*

Especificación funcional para **Inspeccionar Commit** en el MVP de `trunKV`, con enfoque en un *sea of semantic nodes*

#### 🎯 Propósito:

Permitir inspeccionar versiones pasadas del grafo semántico (triplets y sus embeddings), navegando por commits anteriores, recuperando tanto el árbol lógico como las relaciones dinámicas asociadas a un punto en el tiempo.

> Ejemplo: *"Almacenar conocimiento semántico en forma de tripletas versionadas y enlazadas dinámicamente con embeddings."*

---

### 1. **`inspectCommit`**

* **Función:** `inspectCommit(commit_id: Hash): CommitSnapshot`
* **Acciones:**

  * Recuperar metadatos del commit (timestamp, mensaje, autor, padres)
  * Leer el árbol de triplets asociado al commit (estructura `FactTree`)
  * Cargar embeddings relacionados (opcional según snapshot o reconstrucción)
  * Reconstruir enlaces semánticos activos al momento del commit
  * Devolver una vista estructurada para análisis o visualización

---

### 2. **`getFactTreeAtCommit`**

* **Función:** `getFactTreeAtCommit(commit_id: Hash): Tree`
* **Acciones:**

  * Leer el árbol de hechos (`triplets`) asociado al commit
  * Validar integridad de nodos enlazados
  * Incluir estados internos como "deprecated", "active", o "shadowed"

---

### 3. **`getEmbeddingsAtCommit`**

* **Función:** `getEmbeddingsAtCommit(commit_id: Hash): Map<Symbol, Vector>`
* **Acciones:**

  * Consultar embeddings guardados con snapshot explícito
  * O regenerarlos desde los hechos si no están serializados
  * Asociar cada vector con su `Symbol` o `ConceptNode`

---

### 4. **`getSemanticLinksAtCommit`**

* **Función:** `getSemanticLinksAtCommit(commit_id: Hash): LinkGraph`
* **Acciones:**

  * Reconstruir el grafo de relaciones semánticas activas en ese commit
  * Usar embeddings y hechos para regenerar vínculos si necesario
  * Marcar vínculos "hardcoded" (declarativos) y "soft" (por similitud)

---

### 🔁 Flujo Esperado

```text
inspectCommit(commit_id)
  └─→ getFactTreeAtCommit(commit_id)
  └─→ getEmbeddingsAtCommit(commit_id)
  └─→ getSemanticLinksAtCommit(commit_id)
  └─→ renderCommitSnapshot(...)
```

---

### 🧩 Componentes clave

| Componente        | Descripción breve                                 |
| ----------------- | ------------------------------------------------- |
| `triplet store`   | Almacén jerárquico de hechos en formato `(s,p,o)` |
| `embedding index` | Vectores semánticos por símbolo o nodo            |
| `commit log`      | Lista encadenada de commits con árbol y metadata  |
| `link builder`    | Reconstruye relaciones semánticas entre nodos     |
| `snapshot viewer` | Visualiza commits como estados legibles del grafo |

---

### ✅ Buenas prácticas

* Commit inmutable = snapshot consistente del estado
* Separar metadatos del árbol para inspección eficiente
* Embeddings reconstruibles pero opcionalmente persistidos
* Inferencia de relaciones semánticas solo si necesario (lazily evaluated)
* API de inspección debe ser *read-only* y segura frente a cambios
