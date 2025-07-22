## 🧠 trunKV + Fermatean Neutrosophic Evaluation (FNS) — Formalización Conceptual

Formalización inicial de cómo integrar ramas de un sistema tipo *trunKV* con una evaluación basada en **lógica neutrosófica fermateana**, usando una notación clara y modular

### 🎯 Propósito:

Usar ramas semánticas de `trunKV` como representaciones alternativas de conocimiento, hipótesis o escenarios; y evaluarlas bajo una lógica neutrosófica fermateana, donde cada rama es caracterizada por su grado de verdad, falsedad y neutralidad (con potencias > 1).

---

### 🧩 Componentes involucrados:

| Componente            | Rol                                                       |
| --------------------- | --------------------------------------------------------- |
| `Branch`              | Conjunto versionado de tripletas y embeddings             |
| `FNS_Vector(branch)`  | Vector ⟨T, I, F⟩ ∈ \[0,1]³^p con valores Fermateanos      |
| `Evaluator`           | Módulo que genera valores T, I, F basados en propiedades  |
| `Semantic Comparator` | Calcula similitud, contradicción y ambigüedad entre ramas |
| `Decision Engine`     | Selecciona ramas "dominantes" bajo reglas neutrosóficas   |

---

### 🧮 Definición de Evaluación Fermateana para Branches

Cada rama `Bᵢ` se representa por un vector:

$$
FNS(Bᵢ) = \langle Tᵢ, Iᵢ, Fᵢ \rangle \quad \text{tal que} \quad Tᵢ^p + Iᵢ^p + Fᵢ^p \leq 1, \quad p > 1
$$

Donde:

* `Tᵢ`: Veracidad semántica (e.g. consistencia interna + compatibilidad con otras ramas "canon").
* `Iᵢ`: Indeterminación (e.g. tripletas contradictorias, ambigüedad léxica, enlaces débiles).
* `Fᵢ`: Falsedad inferida (e.g. contradicciones duras, outliers lógicos o factuales).

---

### 🔢 Paso a paso del Flujo Evaluativo

#### 1. Recolección de Ramas:

$$
\mathcal{B} = \{B₁, B₂, ..., Bₙ\} \quad \text{candidatas para evaluación}
$$

#### 2. Evaluación Semántica:

Cada `Bᵢ` es evaluada y recibe:

* `Tᵢ = semantic_coherence(Bᵢ)`
* `Iᵢ = ambiguity_score(Bᵢ)`
* `Fᵢ = contradiction_penalty(Bᵢ)`

Todo se normaliza y se eleva a potencia `p > 1` (e.g. `p = 2.5` en conjuntos Fermateanos).

---

#### 3. Dominancia Neutrosófica

Una rama `Bᵢ` domina a otra `Bⱼ` si:

$$
Tᵢ > Tⱼ \quad \text{y} \quad Iᵢ < Iⱼ \quad \text{y} \quad Fᵢ < Fⱼ
$$

Con posibilidad de incluir ponderaciones si se desea priorizar T sobre I y F.

---

### 🧪 Ejemplo (simplificado)

| Rama | Tᵢ  | Iᵢ   | Fᵢ   | Dominante |
| ---- | --- | ---- | ---- | --------- |
| A    | 0.8 | 0.1  | 0.05 | ✔️        |
| B    | 0.7 | 0.3  | 0.2  | ❌         |
| C    | 0.6 | 0.15 | 0.1  | ❌         |

Resultado: `A` es la rama más semánticamente confiable.

---

### 📘 Ventajas de esta Integración

* 📈 **Comparación semántica** entre hipótesis (ramas) en un marco matemáticamente fundado.
* 🔄 **Evaluación dinámica:** FNS puede recalcularse al actualizar ramas.
* 🤖 **Integración con LLMs:** los valores T, I y F pueden inferirse a partir de embeddings, contradicciones detectadas o distancia semántica.
* ⚖️ **Tolerancia a incertidumbre**, ambigüedad y contradicción, propio del razonamiento flexible.

---
