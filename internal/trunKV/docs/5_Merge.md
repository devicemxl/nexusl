## 🧠 `trunKV` - MVP *Sea of Semantic Nodes*

**Especificación funcional** del proceso `merge` adaptado al sistema **trunKV**, donde los árboles representan estructuras semánticas versionadas con *triplets*, *embeddings* y *enlaces dinámicos*.

#### 🎯 Propósito:

Permitir la fusión coherente de dos versiones del árbol semántico, integrando tripletas y sus relaciones, actualizando embeddings y resolviendo conflictos conceptuales entre ramas de conocimiento.

> Ejemplo: *"Fusionar dos ramas de conocimiento versionado, resolviendo conflictos semánticos y preservando coherencia lógica del grafo expandido."*

---

### 1. **`loadSemanticTree(commitId)`**

* **Función:** `loadSemanticTree(commitId: CommitHash): SemanticTree`
* **Acciones:**

  * Recuperar el estado completo del árbol semántico (triplets, embeddings, links) de un commit dado.
  * Permitir acceso estructurado a los nodos y relaciones.

---

### 2. **`diffTriplets(treeA, treeBase)`**

* **Función:** `diffTriplets(a: SemanticTree, base: SemanticTree): TripletChanges`
* **Acciones:**

  * Identificar adiciones, eliminaciones y modificaciones de triplets entre `a` y `base`.
  * Categorizar cambios como un conjunto de operaciones semánticas: `added`, `deleted`, `changed`.

---

### 3. **`mergeTriplets(diffA, diffB)`**

* **Función:** `mergeTriplets(diffA, diffB): MergeResult`
* **Acciones:**

  * Comparar los cambios desde `base → A` y `base → B`.
  * Para cada triplet:

    * Si cambia en una sola rama → aplicar el cambio.
    * Si cambia en ambas ramas de forma distinta → marcar como conflicto.
  * Retornar:

    * `mergedTriplets`
    * `conflictList`

---

### 4. **`resolveConflicts(conflictList)`**

* **Función:** `resolveConflicts(conflictList: List[Conflict]) → ResolvedTriplets`
* **Acciones:**

  * (Manual o automática) Resolver conflictos entre versiones del mismo concepto.
  * Definir reglas semánticas (prioridad, dominancia, disyunción).
  * Generar triplets consistentes o anotaciones para revisión futura.

---

### 5. **`rebuildSemanticLinks(mergedTriplets)`**

* **Función:** `rebuildSemanticLinks(triplets: List[Triplet]) → LinkMap`
* **Acciones:**

  * Recalcular vínculos semánticos entre entidades en función de relaciones actualizadas.
  * Actualizar índices de accesos cruzados o links conceptuales.

---

### 6. **`updateEmbeddings(mergedTriplets)`**

* **Función:** `updateEmbeddings(triplets: List[Triplet]) → EmbeddingIndex`
* **Acciones:**

  * Calcular (o ajustar) embeddings para nuevos o modificados conceptos.
  * Actualizar vectores semánticos en el índice general.
  * Sincronizar espacio vectorial con la nueva topología semántica.

---

### 7. **`createMergeCommit(mergedTree, parentA, parentB)`**

* **Función:** `createMergeCommit(tree: SemanticTree, parentA: CommitHash, parentB: CommitHash) → CommitHash`
* **Acciones:**

  * Serializar el nuevo árbol fusionado.
  * Registrar un nuevo commit con doble parent (`A` y `B`).
  * Persistir con metadata de conflicto (si existieron) y referencia a árbol resultante.

---

### 🔁 Flujo Esperado

```text
loadSemanticTree(A) + loadSemanticTree(B) + loadSemanticTree(base)
        ↓
     diffTriplets(A, base) + diffTriplets(B, base)
        ↓
         mergeTriplets(diffA, diffB)
        ↓
      resolveConflicts(conflictList)
        ↓
     rebuildSemanticLinks(mergedTriplets)
        ↓
     updateEmbeddings(mergedTriplets)
        ↓
     createMergeCommit(newTree, A, B)
```

---

### 🧩 Componentes clave

| Componente          | Descripción breve                                             |
| ------------------- | ------------------------------------------------------------- |
| `triplet store`     | Base de datos de hechos versionados en forma de tripletas     |
| `embedding index`   | Vectores que representan el significado contextual            |
| `commit log`        | Historial versionado del sistema semántico                    |
| `link builder`      | Sistema que infiere y mantiene los enlaces semánticos activos |
| `conflict resolver` | Lógica de detección y resolución de colisiones semánticas     |

---

### ✅ Buenas prácticas

* Resolver conflictos conceptuales con ayuda de embeddings + LLM
* Permitir previsualización visual de conflictos semánticos
* Documentar merge con metadatos narrativos ("por qué se fusionó así")
* Guardar conflictos sin resolver como nodos anotados (de tipo `conflict`)

---
