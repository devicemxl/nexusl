// Package store proporciona un almacén de conocimiento para NexusL.
package store

import (
	"sync"

	"github.com/devicemxl/nexusl/pkg/object"
)

// KnowledgeStore almacena las tripletas de conocimiento.
type KnowledgeStore struct {
	// mu es un mutex para sincronizar el acceso al almacén.
	mu sync.RWMutex
	// tripletas es una lista de tripletas de conocimiento.
	tripletas []*object.TripletaObject
	// Futuras optimizaciones: índices por sujeto, predicado, etc.
}

// NewKnowledgeStore devuelve un nuevo almacén de conocimiento.
func NewKnowledgeStore() *KnowledgeStore {
	return &KnowledgeStore{
		tripletas: make([]*object.TripletaObject, 0),
	}
}

// AddTripleta añade una tripleta al almacén.
func (ks *KnowledgeStore) AddTripleta(t *object.TripletaObject) {
	ks.mu.Lock()
	defer ks.mu.Unlock()
	ks.tripletas = append(ks.tripletas, t)
}

// FindTripletas busca tripletas que coincidan con un patrón.
func (ks *KnowledgeStore) FindTripletas(pattern *object.TripletaObject) []*object.TripletaObject {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	var results []*object.TripletaObject
	for _, t := range ks.tripletas {
		// Lógica de coincidencia para Pattern Matching / Unificación (simplificado por ahora)
		subjectMatch := (pattern.Subject == nil || pattern.Subject.Type == object.TermVariable || pattern.Subject.Value == t.Subject.Value)
		predicateMatch := (pattern.Predicate == nil || pattern.Predicate.Type == object.TermVariable || pattern.Predicate.Value == t.Predicate.Value)
		objectMatch := (pattern.Object == nil || pattern.Object.Type == object.TermVariable || pattern.Object.Value == t.Object.Value)

		if subjectMatch && predicateMatch && objectMatch {
			results = append(results, t)
		}
	}
	return results
}

// GetAllTripletas devuelve todas las tripletas almacenadas.
func (ks *KnowledgeStore) GetAllTripletas() []*object.TripletaObject {
	ks.mu.RLock()
	defer ks.mu.RUnlock()
	return ks.tripletas
}

// RemoveTripleta elimina tripletas que coincidan con el sujeto y predicado dados.
// Se usa para implementar la semántica de unicidad de 'assign' y 'def'.
func (ks *KnowledgeStore) RemoveTripleta(subject *object.Term, predicate *object.Term) {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	var newTripletas []*object.TripletaObject
	for _, t := range ks.tripletas {
		// Si el sujeto y el predicado coinciden, NO lo añadimos a la nueva lista (lo eliminamos).
		if !(t.Subject.Value == subject.Value && t.Predicate.Value == predicate.Value) {
			newTripletas = append(newTripletas, t)
		}
	}
	ks.tripletas = newTripletas
}
