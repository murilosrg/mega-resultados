package service

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/murilosrg/mega-resultados/database"
	"github.com/murilosrg/mega-resultados/models"
	"github.com/murilosrg/mega-resultados/utils"
)

var wg sync.WaitGroup
var db *database.Database = database.New(database.Config{})

//Process file and insert in database
func Process(files []string) error {
	file := isMatchFile(files)
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))

	if err != nil {
		return err
	}

	extract(doc)

	return nil
}

func isMatchFile(files []string) string {
	for i := range files {
		if strings.HasSuffix(files[i], ".htm") {
			return files[i]
		}
	}

	return ""
}

func extract(doc *goquery.Document) {
	input := make(chan []string)
	output := make(chan models.Draw)

	go transform(input, output)
	go load(output)
	wg.Add(2)

	chanIndex := 0

	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(index int, trhtml *goquery.Selection) {
			var rows []string
			trhtml.Find("td").Each(func(index int, tdhtml *goquery.Selection) {
				rows = append(rows, tdhtml.Text())
			})

			if rows != nil && len(rows) >= 21 {
				input <- rows
				chanIndex++
				fmt.Println("Processing index:", chanIndex)
			}
		})

		wg.Done()
		wg.Done()
	})
}

func transform(input chan []string, output chan models.Draw) {
	for {
		select {
		case r, ok := <-input:
			if !ok {
				close(input)
				close(output)
				wg.Done()
				break
			}

			wg.Add(1)

			go func(row []string) {
				draw := models.Draw{
					Date:                utils.ConvertToDate(row[1]),
					Concourse:           utils.ConvertToInt(row[0]),
					Numbers:             utils.ConvertToIntSlice(row[2], row[3], row[4], row[5], row[6], row[7]),
					Collection:          utils.ConvertToDecimal(row[8]),
					Accumulated:         utils.ConvertToBool(row[17]),
					AccumulatedValue:    utils.ConvertToDecimal(row[18]),
					EstimatedPrize:      utils.ConvertToDecimal(row[19]),
					AccumulatedLastDraw: utils.ConvertToDecimal(row[20]),
					Winners: &[]models.Winner{
						{
							Number: 6,
							Count:  utils.ConvertToInt(row[9]),
							Amount: utils.ConvertToDecimal(row[12]),
						},
						{
							Number: 5,
							Count:  utils.ConvertToInt(row[13]),
							Amount: utils.ConvertToDecimal(row[14]),
						},
						{
							Number: 4,
							Count:  utils.ConvertToInt(row[15]),
							Amount: utils.ConvertToDecimal(row[16]),
						},
					},
				}

				output <- draw
				wg.Done()
			}(r)

		default:
		}
	}
}

func load(ch chan models.Draw) {
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				wg.Done()
				break
			}

			wg.Add(1)

			go func(draw models.Draw) {
				_, err := db.Insert("draw", draw)

				if err != nil {
					panic(err)
				}

				wg.Done()
			}(i)

		default:
		}
	}
}
