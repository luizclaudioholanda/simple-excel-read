package main

import (
	"fmt"

	"github.com/luizclaudioholanda/simple-excel-read/src/utils"
)

func main() {

	fmt.Println("Read Excel file.")

	utils.Read("test.xlsx")
}
