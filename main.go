package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// benchplot can be used to plot the output of a benchmark run on varying configuration. The configuration parameter which is
// subject to change needs to be part of the benchmark title separated by an underline. For example when trying the same
// benchmark with different configurations, the name in the output should look like:
//
// 	- Benchmark/configParam_1-<cores>  -> the "1" being the parameter which will be the value of the x-axis
//
// This example output can be found in ./example/bench.txt

func main() {
	var benchFilename, xLabel, title string

	flag.StringVar(&benchFilename, "file", "./example/bench.txt", "The path to the file with the benchmark result")
	flag.StringVar(&title, "title", "Benchmark", "Title of the plot")
	flag.StringVar(&xLabel, "xLabel", "Tick", "Label for the x-axis")
	flag.Parse()

	f, err := os.Open(benchFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pts := extractRuntimes(f)

	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = "Runtime [s]"

	err = plotutil.AddLinePoints(p, "", pts)
	if err != nil {
		log.Fatal(err)
	}

	// Save the plot to a PNG file
	outputFilename := strings.TrimSuffix(benchFilename, filepath.Ext(benchFilename)) + ".png"
	if err := p.Save(8*vg.Inch, 8*vg.Inch, outputFilename); err != nil {
		log.Fatal(err)
	}
}

func extractRuntimes(r io.Reader) plotter.XYs {
	pts := plotter.XYs{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "Benchmark") || strings.Contains(line, "FAIL") {
			continue
		}
		f := strings.Fields(line)
		split := strings.Split(f[0], "-")
		strX := split[0][strings.Index(split[0], "_")+1:]
		strY := f[2]
		x, err := strconv.ParseFloat(strX, 64)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.ParseFloat(strY, 64)
		if err != nil {
			log.Fatal(err)
		}
		pts = append(pts, plotter.XY{X: x, Y: y / 1e9})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return pts
}
