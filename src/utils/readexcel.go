package utils

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Read Read an excel file
func Read(fileName string) {

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	firstSheet := f.WorkBook.Sheets.Sheet[0].Name
	fmt.Printf("'%s' is first sheet of %d sheets.\n", firstSheet, f.SheetCount)
	rows := f.GetRows(firstSheet)

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
