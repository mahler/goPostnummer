package zipcodes

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type zipcode struct {
	zip      int
	cityname string
}

func Initialize() []zips, error {
	var zip []zipcode
	f, err := excelize.OpenFile("../data/postnummerfil.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows("postnr")

	// SKip header columns
	rows = rows[2:]

	for _, row := range rows {
		number, _ := strconv.Atoi(row[0])
		name := row[1]

		zc := zipcode{zip: number, cityname: name}

		zip = append(zip, zc)
	}

	return zip, error{}
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
