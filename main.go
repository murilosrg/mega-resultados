package main

import (
	"github.com/murilosrg/mega-resultados/file"
	"github.com/murilosrg/mega-resultados/service"
)

func main() {
	fileURL := "http://www1.caixa.gov.br/loterias/_arquivos/loterias/D_megase.zip"
	fileResult := "results.zip"

	files, err := file.Download(fileURL, fileResult)

	checkErr(err)

	err = service.Process(files)

	checkErr(err)

	err = file.Remove(fileResult)

	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
