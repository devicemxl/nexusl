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
		// Implementar lógica de coincidencia simple por ahora.
		if (pattern.Subject == nil || pattern.Subject.Value == t.Subject.Value) &&
			(pattern.Predicate == nil || pattern.Predicate.Value == t.Predicate.Value) &&
			(pattern.Object == nil || pattern.Object.Value == t.Object.Value) {
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
