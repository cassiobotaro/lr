package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {
	// we open the csv file from the disk
	f, err := os.Open("./datasets/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// we create a new csv reader specifying
	// the number of columns it has
	salesData := csv.NewReader(f)
	salesData.FieldsPerRecord = 21

	// we read all the records
	records, err := salesData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// In this case we are going to try and model our house price (y)
	// by the grade feature.
	var r regression.Regression
	r.SetObserved("Price")
	r.SetVar(0, "Grade")
	r.SetVar(1, "Grade2")

	// Loop of records in the CSV, adding the training data to the regressionvalue.
	for i, record := range records {
		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the house price, "y".
		price, err := strconv.ParseFloat(records[i][2], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the grade value.
		grade, err := strconv.ParseFloat(record[11], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(price, []float64{grade, math.Pow(grade, 2)}))
	}

	// Train/fit the regression model.
	r.Run()
	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
}
