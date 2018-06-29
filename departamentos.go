package divipola

import (
	"regexp"
)

// Departamento representa un departamento o distrito capital.
type Departamento string

// Código obtiene el código de 2 dígitos del departamento o distrito capital.
func (D Departamento) Código() string {
	return string(D)
}

// Válido comprueba si es un código válido.
func (D Departamento) Válido() bool {
	_, existe := mapaDepartamentos[string(D)]
	return existe
}

// Nombre representa el nombre completo del departamento o distrito capital.
func (D Departamento) Nombre() string {
	nombre, existe := mapaDepartamentos[string(D)]
	if !existe {
		return "(Desconocido)"
	}

	return nombre
}

func (D Departamento) String() string {
	return D.Nombre()
}

// Departamentos lista los departamentos y distrito capital.
// Se pueden enviar parámetros adicionales de ordenamiento y filtro.
func Departamentos(prmtrs ...interface{}) []Departamento {
	var array = departamentosOrdenadosAlfabéticamente
	var dptmts = []Departamento{}

	var orden, filtro = normalizarParámetrosListado(prmtrs...)

	if orden == OrdenNumérico {
		array = departamentosOrdenadosNuméricamente
	}

	if len(filtro) == 0 {
		for _, v := range array {
			dptmts = append(dptmts, Departamento(v))
		}
	} else {
		for _, v := range array {
			if regexp.MustCompile(filtro).MatchString(normalizarTexto(Departamento(v).String())) || regexp.MustCompile(filtro).MatchString(v) {
				dptmts = append(dptmts, Departamento(v))
			}
		}
	}

	return dptmts
}

var departamentosOrdenadosNuméricamente = []string{"05", "08", "11", "13", "15", "17", "18", "19", "20", "23", "25", "27", "41", "44", "47", "50", "52", "54", "63", "66", "68", "70", "73", "76", "81", "85", "86", "88", "91", "94", "95", "97", "99"}

var departamentosOrdenadosAlfabéticamente = []string{"05", "91", "81", "88", "08", "11", "13", "15", "17", "18", "85", "19", "20", "27", "23", "25", "94", "95", "41", "44", "47", "50", "52", "54", "86", "63", "66", "68", "70", "73", "76", "97", "99"}

var mapaDepartamentos = map[string]string{
	"05": "Antioquia",
	"08": "Atlántico",
	"11": "Bogotá, D.C.",
	"13": "Bolívar",
	"15": "Boyacá",
	"17": "Caldas",
	"18": "Caquetá",
	"19": "Cauca",
	"20": "Cesar",
	"23": "Córdoba",
	"25": "Cundinamarca",
	"27": "Chocó",
	"41": "Huila",
	"44": "La Guajira",
	"47": "Magdalena",
	"50": "Meta",
	"52": "Nariño",
	"54": "Norte de Santander",
	"63": "Quindío",
	"66": "Risaralda",
	"68": "Santander",
	"70": "Sucre",
	"73": "Tolima",
	"76": "Valle del Cauca",
	"81": "Arauca",
	"85": "Casanare",
	"86": "Putumayo",
	"88": "Archipiélago de San Andrés, Providencia y Santa Catalina",
	"91": "Amazonas",
	"94": "Guainía",
	"95": "Guaviare",
	"97": "Vaupés",
	"99": "Vichada",
}
