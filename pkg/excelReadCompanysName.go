package pkg

import (
	"log"
	"strconv"
	"sync"

	"github.com/xuri/excelize/v2"
)

func StartReadAndSearchInn() error {
	//var listOfCompanies []string
	var wtgr sync.WaitGroup
	f, err := excelize.OpenFile("datas.xlsx")
	if err != nil {
		return err
	}
	columns, err := f.GetCols("Лист1")
	if err != nil {
		return err
	}

	for _, value := range columns {
		if value[0] == "pagetitle" {
			for indxCompany, val := range value {
				if val == "pagetitle" {
					continue
				}
				wtgr.Add(1)

				go func(indxCompany int, val string) {
					list := FindInn(val)
					if err := f.SetCellValue("Лист1", "C"+strconv.Itoa(indxCompany+1), list); err != nil {
						log.Println(err)
					}
					wtgr.Done()
				}(indxCompany, val)

			}
			break
		}
	}

	wtgr.Wait()
	defer f.SaveAs("changed.xlsx")
	defer f.Close()
	return nil
}
