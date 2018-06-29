# divipola
Codificación oficial de la división político-administrativa de Colombia. Municipios y departamentos.

```
go get -u github.com/saulortega/divipola
```

```go
import (
	"fmt"
	"github.com/saulortega/divipola"
)

func main() {
	// Obtener todos los departamentos y distrito capital, ordenados alfabéticamente:
	var deps1 = divipola.Departamentos()
	fmt.Println(deps1)

	// Obtener todos los departamentos y distrito capital ordenados según su código:
	var deps2 = divipola.Departamentos(divipola.OrdenNumérico)
	fmt.Println(deps2)

	// Obtener los departamentos o distrito capital cuyo nombre contiene "San":
	var deps3 = divipola.Departamentos(divipola.Filtro("San"))
	fmt.Println(deps3)

	// Obtener los departamentos o distrito capital cuyo nombre contiene "San" ordenados según su código:
	var deps4 = divipola.Departamentos(divipola.OrdenNumérico, divipola.Filtro("San"))
	fmt.Println(deps4)

	// Obtener los municipios cuyo nombre contiene "ocaña":
	var muns1 = divipola.Municipios(divipola.Filtro("ocaña"))
	fmt.Println(muns1)

	// Obtener el código de un municipio:
	var mun = muns1[0]
	fmt.Println(mun.Código())

	// Obtener los tres últimos dígitos del código de un municipio, descartando los dos de su departamento:
	fmt.Println(mun.Código3())

	// Obtener el nombre de un municipio:
	fmt.Println(mun.Nombre())

	// Verificar si un municipio es válido:
	fmt.Println(mun.Válido())

	// Obtener el departamento o distrito capital de un municipio:
	fmt.Println(mun.Departamento())

	// Obtener el código de un departamento o distrito capital:
	fmt.Println(mun.Departamento().Código())

	// Obtener el nombre de un departamento o distrito capital:
	fmt.Println(mun.Departamento().Nombre())

	// Verificar si un departamento o distrito capital es válido:
	fmt.Println(mun.Departamento().Válido())
}
```
