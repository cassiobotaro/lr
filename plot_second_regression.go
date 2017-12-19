package main

import (
	"image/color"
	"log"
	"math"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// we open the csv file from the disk
	f, err := os.Open("./datasets/kc_house_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	df := dataframe.ReadCSV(f)

	// pts will hold the values for plotting.
	pts := make(plotter.XYs, df.Nrow())

	yVals := df.Col("price").Float()
	for i, floatVal := range df.Col("grade").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	// Create the plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "grade"
	p.Y.Label.Text = "house price"
	p.Add(plotter.NewGrid())
	// Add the scatter plot points for the observations.
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Radius = vg.Points(2)
	s.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	curve := plotter.NewFunction(predict)
	curve.LineStyle.Width = vg.Points(3)
	curve.LineStyle.Dashes = []vg.Length{vg.Points(3), vg.Points(3)}
	curve.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	// Save the plot to a PNG file.
	p.Add(s, curve)
	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "./graphs/second_regression.png"); err != nil {
		log.Fatal(err)
	}
}

func predict(grade float64) float64 {
	return 1639674.31 + grade*-473161.41 + math.Pow(grade, 2)*42070.46
}
