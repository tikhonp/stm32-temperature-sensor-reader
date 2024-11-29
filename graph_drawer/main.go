package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	serial "github.com/tarm/goserial"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	comPort       = "/dev/tty.usbserial-AQ018TY4"
	baudRate      = 115200
	graphFileName = "points.png"
)

type ClimatFrame struct {
	Temp int
	Hum  int
}

var data []ClimatFrame

func generatePointsGraph() {
	p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"Temp", plotDataFrom(len(data), func(idx int) float64 { return float64(data[idx].Temp) }),
		"Hum", plotDataFrom(len(data), func(idx int) float64 { return float64(data[idx].Hum) }),
	)
	if err != nil {
		log.Fatalln(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, graphFileName); err != nil {
		log.Fatalln(err)
	}
}

func plotDataFrom(length int, getData func(idx int) float64) plotter.XYs {
	pts := make(plotter.XYs, length-1)
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = getData(i)
	}
	return pts
}

func remove(s string) string {
	return strings.Map(
		func(r rune) rune {
			if unicode.IsDigit(r) {
				return r
			}
			return -1
		},
		s,
	)
}

func processString(input string) {
	lines := strings.Split(input, "\n")
	i := 0
	for i < len(lines) {
		if strings.Contains(lines[i], "Hell") {
			break
		}
		i++
	}
	for ; i < len(lines)-2; i += 3 {
		var temp, hum int
		temp, err := strconv.Atoi(remove(lines[i+1]))
		if err != nil {
			log.Fatalln(err)
		}
		hum, err = strconv.Atoi(remove(lines[i+2]))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(temp, hum)
		data = append(data, ClimatFrame{
			Temp: temp,
			Hum:  hum,
		})
	}
}

func main() {
	c := &serial.Config{Name: comPort, Baud: baudRate}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	var data string
	for {
		buf := make([]byte, 2048)
		r, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		if i < 15 {
			data += string(buf[:r])
			i++
			println("i: ", i)
		} else {
			processString(data)
			i = 0
			data = ""
			generatePointsGraph()
		}
	}
}
