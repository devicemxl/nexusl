## 🧠 trunKV - MVP Declaración de Hechos Semánticos Versionados

Especificación funcional de un MVP para declarar un `fact` (hecho) en **trunKV**, basado en una arquitectura tipo *sea of semantic nodes*, con tripletas, embeddings y versionado como ejes centrales

#### 🎯 Propósito

*"Almacenar conocimiento semántico en forma de tripletas versionadas y enlazadas dinámicamente con embeddings, permitiendo su recuperación, actualización y vinculación en un grafo semántico evolutivo."*

---

### 1. **Declarar Hecho**

* **Función:** `storeTriplet(subject, predicate, object)`
* **Acciones:**

  * Validar estructura del triplet (`Atom`-like).
  * Insertar en el almacén semántico (`tripletStore`).
  * Asignar timestamp y contexto (versión o rama activa).
  * Indexar por ID o hash único.
  * Marcar como "dirty" para actualización de embeddings.

---

### 2. **Actualizar Embedding**

* **Función:** `updateEmbedding(tripletID)`
* **Acciones:**

  * Recuperar el triplet.
  * Generar o actualizar su embedding semántico.
  * Insertar o reemplazar en el `embeddingIndex`.
  * Calcular distancia/relación con embeddings existentes.

---

### 3. **Construir Enlaces Semánticos**

* **Función:** `buildSemanticLinks(tripletID)`
* **Acciones:**

  * Inferir enlaces por similitud (coseno, k-NN, etc.).
  * Crear vínculos explícitos o tags relacionados.
  * Agregar entradas en `semanticGraph`.
  * Actualizar metadata del triplet con enlaces salientes/entrantes.

---

### 4. **Commit Semántico**

* **Función:** `commitSemanticState(message)`
* **Acciones:**

  * Capturar el snapshot del árbol semántico (triplets y enlaces).
  * Serializar estado y calcular hash.
  * Crear objeto de `commit` apuntando al estado actual.
  * Actualizar `HEAD` o rama activa con el nuevo commit.
  * Guardar en el `commitLog`.

---

### 5. **Indexar Triplet por Rutas**

* **Función:** `indexTriplet(tripletID)`
* **Acciones:**

  * Crear índice por `subject`, `predicate`, `object`, hash.
  * Permitir navegación semántica rápida.
  * Posibilitar autocompletado o expansión semántica posterior.

---

### 6. **Nota/Anotación Extendida**

* **Función:** `attachNote(tripletID, content)`
* **Acciones:**

  * Asociar nota a un triplet base.
  * Permitir expansión conceptual (como subgrafo textual).
  * Aumentar contexto semántico y enriquecer el embedding.

---

### 🔁 Flujo Esperado (Mínimo)

```
storeTriplet → updateEmbedding → buildSemanticLinks → indexTriplet → commitSemanticState
```

Opcionalmente:

```
attachNote → updateEmbedding → buildSemanticLinks → commitSemanticState
```

---

### 🧩 Componentes clave

| Componente        | Descripción breve                                 |
| ----------------- | ------------------------------------------------- |
| `tripletStore`    | Almacén de tripletas en forma `(s, p, o)`         |
| `embeddingIndex`  | Mapa `tripletID → vector` para búsqueda semántica |
| `semanticGraph`   | Grafo con enlaces inferidos y explícitos          |
| `commitLog`       | Historial versionado del grafo semántico          |
| `annotationStore` | Notas y metadata extendida para cada nodo         |

---

### ✅ Buenas prácticas

* **Normalización**: transformar sujetos y predicados a formas canónicas antes de insertar.
* **Embeddings asincrónicos**: permitir actualización perezosa o por lotes.
* **Tagging dinámico**: etiquetas generadas desde enlaces, no solo manualmente.
* **Visualización semántica**: herramienta opcional para inspección de significado.
* **Segmentación por rama/contexto**: útil para experimentación de submodelos.

---
