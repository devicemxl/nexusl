## 🧠 `trunKV` - MVP: Manejo de Ramas en un Sistema de Nodos Semánticos

**Especificación funcional para el manejo de ramas** (`Branch`) en el MVP de `trunKV`, con enfoque en un **grafo semántico versionado** tipo *sea of semantic nodes*.


#### 🎯 Propósito:

Permitir múltiples líneas de evolución semántica independientes dentro de un sistema versionado de conocimiento basado en tripletas, donde cada rama representa un contexto o narrativa evolutiva propia. Esto facilita la experimentación, la divergencia de significado y la fusión posterior.

---

### 1. **Crear Rama**

* **Función:** `create_branch(name: str) -> Result`
* **Acciones:**

  * Verifica si ya existe una rama con ese nombre en `branchIndex`.
  * Crea una entrada en el registro de ramas (`branches`) apuntando al commit actual (`HEAD.commit`).
  * La rama hereda el grafo semántico hasta ese punto (`tripletStore`, `embeddingIndex`, `semanticLinks`).

---

### 2. **Cambiar de Rama**

* **Función:** `switch_branch(name: str) -> Result`
* **Acciones:**

  * Verifica que la rama exista en `branches`.
  * Actualiza el puntero `HEAD` para que apunte a esa rama.
  * Carga los índices de triplets, embeddings y semantic links según el commit al que apunta la rama.
  * Restaura el estado del *sea of semantic nodes* para esa narrativa.

---

### 3. **Eliminar Rama**

* **Función:** `delete_branch(name: str) -> Result`
* **Acciones:**

  * Verifica que la rama no esté activa (`HEAD`).
  * Elimina la entrada del índice de ramas (`branches[name]`).
  * No borra los commits ni los nodos semánticos compartidos (uso compartido permitido entre ramas).

---

### 🔁 Flujo Esperado

```plaintext
create_branch("bio-narrativa") ⟶ switch_branch("bio-narrativa")
⟶ insertTriplet("Anna is biologist")
⟶ updateEmbeddings("Anna") ⟶ buildSemanticLinks("Anna")
⟶ commit("added profession to Anna")
```

---

### 🧩 Componentes clave

| Componente       | Descripción breve                                                |
| ---------------- | ---------------------------------------------------------------- |
| `branches`       | Diccionario: `name → commitID`, usado para navegar narrativas    |
| `HEAD`           | Puntero activo: `{branch: name, commit: id}`                     |
| `tripletStore`   | Base de tripletas activas en esa rama                            |
| `embeddingIndex` | Embeddings semánticos de los conceptos de esa rama               |
| `semanticLinks`  | Grafo dinámico generado por co-ocurrencias o similitud semántica |
| `commitLog`      | Historial versionado (append-only) para cada rama                |

---

### ✅ Buenas prácticas

* Cada rama es una *vista semántica coherente* de un universo.
* No borrar ramas activas.
* Permitir fusión eventual si los conceptos convergen (a definir en `merge_branch()`).
* Las ramas pueden divergir en significado, útil para contextos ambiguos, contradictorios o exploratorios.
* Mantener sincronización mínima entre embeddings y commits para evitar inconsistencias.

---
