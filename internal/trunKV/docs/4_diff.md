## 🧠 `trunKV` - MVP Diff Semántico entre Árboles

Especificación funcional del flujo **Mostrar Diferencias entre Árboles (Diff)** adaptado al contexto semántico de `trunKV`, manteniendo la idea de un *sea of semantic nodes* y embebiendo la noción de tripletas, embeddings y versiones

#### 🎯 Propósito

Identificar las diferencias semánticas entre dos árboles versionados de conocimiento representado mediante tripletas enriquecidas con embeddings. Esto permite auditar cambios, detectar evolución conceptual o generar merges inteligentes.

> *Comparar versiones semánticas de un grafo, observando qué entidades/triplets fueron agregadas, eliminadas o semánticamente modificadas.*

---

### 1. **`diff(commitA, commitB)`**

* **Función:** `diff(commitA, commitB)`
* **Acciones:**

  * Cargar los árboles semánticos completos desde los commits dados.
  * Indexar tripletas por su `subject` para comparación eficiente.
  * Detectar cambios estructurales:

    * **Agregados:** tripletas presentes en `B` pero no en `A`.
    * **Eliminados:** tripletas presentes en `A` pero no en `B`.
    * **Modificados:** misma `subject` y `predicate` pero `object` distinto (o embedding significativamente diferente).
  * Calcular `semanticDistance(embeddingA, embeddingB)` si ambos existen.
  * Generar reporte agrupado por tipo de cambio (`added`, `deleted`, `modified`), con metainformación.

---

### 2. **`semanticDistance(embeddingA, embeddingB)`**

* **Función:** `semanticDistance(vecA, vecB)`
* **Acciones:**

  * Calcular la distancia coseno o métrica elegida entre dos vectores.
  * Determinar si el cambio es significativo (`threshold`).
  * Incluir esta diferencia en la salida de `diff()` como “cambio semántico suave”.

---

### 3. **`loadTree(commitHash)`**

* **Función:** `loadTree(commit)`
* **Acciones:**

  * Leer y deserializar el snapshot completo del grafo semántico (triplets + embeddings) apuntado por `commit`.
  * Verificar validez del árbol (estructura consistente, sin enlaces rotos).

---

### 🔁 Flujo Esperado (mínimo viable)

```plaintext
diff(commitA, commitB)
  → loadTree(commitA)
  → loadTree(commitB)
  → compareTriplets(treeA, treeB)
    ↳ if embeddings: semanticDistance(embA, embB)
  → generateDiffReport()
```

---

### 🧩 Componentes clave

| Componente        | Descripción breve                                     |
| ----------------- | ----------------------------------------------------- |
| `triplet store`   | Repositorio versionado de tripletas semánticas        |
| `embedding index` | Índice de vectores semánticos para entidades y hechos |
| `commit log`      | Registro de snapshots inmutables                      |
| `diff engine`     | Comparador estructural y semántico de árboles         |

---

### ✅ Buenas prácticas

* Realizar `diff` sólo sobre commits validados (firmados/hash correctos).
* Definir claramente el **umbral semántico** de cambio (e.g. `cos(θ) < 0.9`).
* Registrar diferencias con contexto (e.g. cuándo, por qué se hizo el cambio).
* Diferenciar entre cambios estructurales (objeto distinto) y suaves (misma idea con distinta expresión).
* Soportar visualización de `diff` como grafo o tabla temporal para debugging.

---

¿Deseas que lo desarrollemos también como CLI/API semántica o que avancemos ahora con `merge` en este formato?
