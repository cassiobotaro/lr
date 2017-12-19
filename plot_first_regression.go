package main

import (
	"image/color"
	"log"
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
	// ptsPred will hold the predicted values for plotting.
	ptsPred := make(plotter.XYs, df.Nrow())

	yVals := df.Col("price").Float()
	for i, floatVal := range df.Col("grade").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
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
	// Add the line plot points for the predictions.
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(0.5)
	l.LineStyle.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	l.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	// Save the plot to a PNG file.
	p.Add(s, l)
	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "./graphs/first_regression.png"); err != nil {
		log.Fatal(err)
	}
}

func predict(grade float64) float64 {
	return -1044346.23 + grade*207717.23
}
