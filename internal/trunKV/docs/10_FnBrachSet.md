## 🧠 RAZONAMIENTO MULTI-BRANCA CON LÓGICA FNS

**Razonamiento Multi-Branca con Lógica Fermateana Neutrosófica (FNS)**, aplicándolo a escenarios donde `trunKV` mantiene trayectorias paralelas —por ejemplo, distintas versiones de conocimiento, hipótesis, decisiones o comportamientos en un sistema simbólico o agente.

### 🎯 Objetivo

Evaluar **conjuntos de ramas** que representan *trayectorias paralelas* (multi-ramas) bajo lógica FNS, para:

* Detectar **consistencia transversal**
* Identificar **trayectorias dominantes**
* Medir **divergencia semántica**
* Evaluar decisiones hipotéticas o evoluciones posibles

---

### 🧩 Supuestos del sistema

Cada `Branch` tiene su FNS:

```math
FNS(Bᵢ) = (Tᵢ, Iᵢ, Fᵢ),  ∀ i ∈ {1,...,n}
```

Y un conjunto de ramas forma una **trayectoria paralela**:

```math
𝒫 = {B₁, B₂, ..., Bₙ}
```

Podemos evaluar ese conjunto como un todo:

```math
FNS(𝒫) = f(FNS(B₁), ..., FNS(Bₙ))
```

---

### 📐 I. AGREGACIÓN DE RAMAS

Definimos una **función de agregación** para obtener una FNS compuesta:

```pseudocode
function aggregateFNS(branches: list[Branch], p: float) -> FermateanNeutrosophicSet:
    T_list = [FNS(B).T for B in branches]
    I_list = [FNS(B).I for B in branches]
    F_list = [FNS(B).F for B in branches]

    T̄ = average(T_list)
    Ī = average(I_list)
    F̄ = average(F_list)

    if T̄^p + Ī^p + F̄^p > 1:
        (T̄, Ī, F̄) = normalizeToFermatean(T̄, Ī, F̄, p)

    return FermateanNeutrosophicSet(T̄, Ī, F̄, p)
```

> También puedes usar agregación *ponderada* si cada rama tiene un peso relativo.

---

### 🧠 II. INTERPRETACIÓN DEL CONJUNTO MULTI-BRANCA

| Valor Agregado | Interpretación                                             |
| -------------- | ---------------------------------------------------------- |
| Alto `T̄`      | Coherencia transversal alta (las ramas están de acuerdo)   |
| Alto `Ī`       | Gran incertidumbre entre ramas, contradicción o ambigüedad |
| Alto `F̄`      | Muchas ramas contienen errores o falsedades significativas |

---

### 🧮 III. DIVERGENCIA SEMÁNTICA ENTRE RAMAS

Para cada par de ramas $(Bᵢ, Bⱼ)$, definimos la **divergencia neutrosófica**:

```math
δ(Bᵢ, Bⱼ) = |Tᵢ - Tⱼ| + |Iᵢ - Iⱼ| + |Fᵢ - Fⱼ|
```

La **divergencia total** del conjunto es:

```math
Δ(𝒫) = average(δ(Bᵢ, Bⱼ)) for all i < j
```

#### Interpretación:

| Δ(𝒫)       | Significado                                                  |
| ----------- | ------------------------------------------------------------ |
| ≈ 0         | Las ramas son casi idénticas                                 |
| Moderado    | Trayectorias con variaciones suaves                          |
| Alto (>0.7) | Trayectorias semánticamente divergentes (hipótesis opuestas) |

---

### 📋 IV. DIAGRAMA DE FLUJO DEL RAZONAMIENTO MULTI-BRANCA (TEXTO)

```
[START]
   ↓
[Seleccionar conjunto 𝒫 de ramas paralelas]
   ↓
[Calcular FNS(Bᵢ) para cada rama]
   ↓
[Calcular FNS(𝒫) = agregación neutrosófica]
   ↓
[Calcular Δ(𝒫) = divergencia entre ramas]
   ↓
[Evaluar caso]
   ├─ Si T̄ alto y Δ bajo → consenso, mantener trayectoria
   ├─ Si T̄ alto y Δ alto → trayectorias posibles, continuar exploración
   ├─ Si F̄ alto → revisión de errores
   └─ Si Ī alto → pedir intervención de razonador externo o LLM
   ↓
[END]
```

---

### 🔬 V. CASOS DE USO EJEMPLARES

| Caso                       | Descripción                                                        | Ejemplo `trunKV`                                |
| -------------------------- | ------------------------------------------------------------------ | ----------------------------------------------- |
| *Consistencia histórica*   | Varias ramas representan evolución temporal de un concepto         | ¿Ha cambiado el significado de "democracia"?    |
| *Exploración de hipótesis* | Ramas representan posibles teorías o estrategias                   | Modelar alternativas a una política agrícola    |
| *Simulación de decisiones* | Ramas modelan caminos de decisión distintos                        | ¿Qué pasa si se invierte en tecnología A vs B?  |
| *Evaluación de agentes*    | Ramas representan comportamiento de agentes simbólicos o autónomos | Comparar comportamiento en distintos escenarios |

---

### 🧩 VI. POSIBLE EXTENSIÓN: MAPAS COGNITIVOS FERMATEANOS

En vez de trabajar solo con ramas sueltas, puedes construir un **grafo cognitivo de ramas**, donde los nodos son ramas y las aristas representan:

* Transiciones
* Influencias
* Contradicciones

Y aplicar razonamiento FNS sobre el **grafo entero**, como una especie de *Fermatean Neutrosophic Cognitive Map*.
