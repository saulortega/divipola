// https://github.com/saulortega/divipola
// Saúl Ortega. 2018.

// Este paquete contiene el listado de municipios y departamentos de Colombia según la codificación oficial DIVIPOLA.
package divipola

// Filtro representa una cadena de texto para filtrar la búsqueda.
type Filtro string

// Orden representa el orden en que se debe listar los municipios o departamentos.
type Orden int

var (
	OrdenNumérico   Orden = 1
	OrdenAlfabético Orden = 2
)
