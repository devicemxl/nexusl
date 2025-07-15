// Gothic/metamodel/metamodel.go (ACTUALIZADO para usar símbolos cargados de ds)
package metamodel

import (
	"github.com/devicemxl/nexusl/Gothic/ds" // Importa el paquete ds que ahora contiene Symbol y ThingType
)

// MetamodelDefinitions actúa como un facade para consultar los símbolos del sistema
// cargados globalmente por el paquete `ds`. No guarda sus propios mapas duplicados.
type MetamodelDefinitions struct {
	// No necesitamos mapas internos aquí si ds.SymbolsByPublicName ya es la fuente de verdad.
	// Podrías tener un caché si las consultas a ds.LookupSymbolByPublicName fueran muy costosas,
	// pero para mapas en memoria, probablemente no sea necesario.
}

// NewMetamodelFacade crea una nueva instancia de MetamodelDefinitions que interactúa
// con los símbolos globales cargados por el paquete ds.
func NewMetamodelFacade() *MetamodelDefinitions {
	return &MetamodelDefinitions{}
}

// LookupScope busca una definición de scope por su nombre.
func (mm *MetamodelDefinitions) LookupScope(name string) (*ds.Symbol, bool) {
	sym, ok := ds.LookupSymbolByPublicName(name)
	if !ok || sym.Thing != ds.TripletScopeType {
		return nil, false
	}
	return sym, true
}

// LookupPredicate busca una definición de predicado por su nombre.
func (mm *MetamodelDefinitions) LookupPredicate(name string) (*ds.Symbol, bool) {
	sym, ok := ds.LookupSymbolByPublicName(name)
	if !ok || sym.Thing != ds.PredicateType {
		return nil, false
	}
	return sym, true
}

// Aquí no necesitamos NewTestMetamodel() ni LoadDefinitionsFromBytes()
// ya que ds.LoadSystemDefinitionsFromDB es el que carga los datos.
// Eliminamos:
/*
func NewTestMetamodel() *MetamodelDefinitions { ... }
func LoadDefinitionsFromBytes(data []byte) (*MetamodelDefinitions, error) { ... }
*/
