package divipola

import (
	"regexp"
	"strings"
)

func normalizarTexto(t string) string {
	t = strings.ToLower(t)
	t = regexp.MustCompile(`\s+?`).ReplaceAllString(t, " ")
	t = strings.Replace(t, "á", "a", -1)
	t = strings.Replace(t, "é", "e", -1)
	t = strings.Replace(t, "í", "i", -1)
	t = strings.Replace(t, "ó", "o", -1)
	t = strings.Replace(t, "ú", "u", -1)
	t = strings.Replace(t, "ü", "u", -1)
	return t
}

func normalizarParámetrosListado(prmtrs ...interface{}) (Orden, string) {
	var filtro string
	var orden Orden

	for _, prmtr := range prmtrs {
		switch prmtr.(type) {
		case Orden:
			orden = prmtr.(Orden)
		case Filtro:
			filtro = normalizarTexto(string(prmtr.(Filtro)))
		}
	}

	return orden, filtro
}
