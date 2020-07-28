package controllers

import (
	"fmt"
	"strconv"
	"unicode"

	"../models"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// ZipController holds the controller for Zipcodes
type ZipController struct{}

// New function to setup ZipController
func (z ZipController) New() *ZipController {
	return &ZipController{}
}

// Init reads data from XLS file into slice
func (z ZipController) Init() ([]models.Zipcode, error) {
	var zip []models.Zipcode
	f, err := excelize.OpenFile("../data/postnummerfil.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	rows, err := f.GetRows("postnr")

	// SKip header columns
	rows = rows[2:]

	for _, row := range rows {
		number, _ := strconv.Atoi(row[0])
		name := row[1]

		zc := models.Zipcode{Zip: number, Cityname: name}

		zip = append(zip, zc)
	}

	return zip, nil
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
