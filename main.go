package main

import (
	"github.com/murilosrg/mega-resultados/file"
)

func main() {
	fileURL := "http://www1.caixa.gov.br/loterias/_arquivos/loterias/D_megase.zip"
	fileResult := "results.zip"

	file.Download(fileURL, fileResult)
}
