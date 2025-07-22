## 🔁 FLUJO COMPLETO — `trunKV` + Fermatean Neutrosophic Branch Evaluation

Flujo completo de evaluación de ramas semánticas usando Lógica Neutrosófica Fermateana (FNS)** como parte del motor de razonamiento de `trunKV`.

### 🧩 I. ESTRUCTURA BÁSICA DEL SISTEMA

#### 📘 Definiciones claves

```pseudocode
Branch:
    id: string
    triplets: list of Triplet
    embedding: vector[float]
    meta: {
        timestamp, creator, tags, ...
    }

FermateanNeutrosophicSet:
    T: float  # Truth Degree
    I: float  # Indeterminacy Degree
    F: float  # Falsity Degree
    p: float  # power ≥ 1, e.g. 2.5

    assert (T^p + I^p + F^p) ≤ 1
```

---

### 🧮 II. EVALUACIÓN DE CADA RAMA

#### **Objetivo: generar FNS(Bᵢ) para cada rama Bᵢ**

#### 📐 Paso 1: Coherencia Semántica (Tᵢ)

```pseudocode
function semanticCoherence(B: Branch) -> float:
    score = average_similarity_among_triplets(B.triplets)
    score += compatibility_with_canonical_embeddings(B.embedding)
    return normalize(score)
```

#### 📐 Paso 2: Indeterminación (Iᵢ)

```pseudocode
function indeterminacyScore(B: Branch) -> float:
    return normalize(ambiguity_in_links(B) + variance_in_embeddings(B))
```

#### 📐 Paso 3: Falsedad (Fᵢ)

```pseudocode
function falsityScore(B: Branch) -> float:
    contradictions = detect_contradictions(B.triplets)
    return normalize(contradictions)
```

#### 📐 Paso 4: Generar FNS con potencia `p`

```pseudocode
function generateFNS(B: Branch, p: float = 2.5) -> FermateanNeutrosophicSet:
    T = semanticCoherence(B)
    I = indeterminacyScore(B)
    F = falsityScore(B)
    
    if (T^p + I^p + F^p) > 1:
        (T, I, F) = normalizeToFermatean(T, I, F, p)
    
    return FermateanNeutrosophicSet(T, I, F, p)
```

---

### 🧮 III. COMPARACIÓN ENTRE RAMAS

#### 📋 Dominancia (una rama es mejor que otra)

```pseudocode
function dominates(A: FNS, B: FNS) -> bool:
    return A.T > B.T and A.I < B.I and A.F < B.F
```

#### 📋 Ranking

```pseudocode
function rankBranches(branches: list[Branch]) -> list[Branch]:
    scored = []
    for B in branches:
        fns = generateFNS(B)
        scored.append((B, fns))

    return sort_by(scored, key=lambda x: (x[1].T, -x[1].I, -x[1].F))
```

---

## 🔀 IV. FLUJO OPERATIVO (Texto)

```
[START]
   ↓
[RECOLECTAR ramas candidatas]
   ↓
[Para cada rama Bᵢ]
    ├─ Calcular Tᵢ = coherencia semántica
    ├─ Calcular Iᵢ = ambigüedad / indeterminación
    ├─ Calcular Fᵢ = contradicción / falsedad
    └─ Verificar Tᵢ^p + Iᵢ^p + Fᵢ^p ≤ 1
         └─ Si NO, normalizar
   ↓
[Obtener conjunto FNS(Bᵢ) para cada rama]
   ↓
[COMPARAR ramas usando dominancia o ranking neutrosófico]
   ↓
[SELECCIONAR ramas dominantes o fusionarlas]
   ↓
[Opcional: devolver feedback al sistema para entrenamiento]
   ↓
[END]
```

---

### 📊 EJEMPLO DE TABLA DE COMPARACIÓN

| Rama | Tᵢ   | Iᵢ   | Fᵢ   | Dominante? | Comentario                        |
| ---- | ---- | ---- | ---- | ---------- | --------------------------------- |
| A    | 0.81 | 0.1  | 0.04 | ✔️         | Muy coherente, baja contradicción |
| B    | 0.77 | 0.25 | 0.10 | ❌          | Buena, pero más ambigua           |
| C    | 0.72 | 0.12 | 0.15 | ❌          | Menor T y mayor F                 |

---

### 🧠 V. POSIBLES AMPLIACIONES

* **Fusión semántica**: combinar triplets de ramas dominantes o compatibles.
* **Evaluación probabilística por LLM**: utilizar LLMs ligeros para inferir T/I/F directamente desde embeddings.
* **Razonamiento multi-branca**: aplicar lógica FNS a ramas que representan trayectorias paralelas en sistemas dinámicos o decisiones.
