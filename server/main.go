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
	comPort  = "/dev/tty.usbserial-AQ018TY4"
	baudRate = 115200
)

type Frame struct {
	Temp int
	Hum  int
}

var frames []Frame

func createGraph() {
    p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"Temp", generateTempGraoh(),
		"Hum", generateHumGraoh())
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func generateTempGraoh() plotter.XYs {
	pts := make(plotter.XYs, len(frames)-1)
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = float64(frames[i].Temp)
	}
	return pts
}

func generateHumGraoh() plotter.XYs {
	pts := make(plotter.XYs, len(frames)-1)
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = float64(frames[i].Hum)
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
		// println(".", strings.TrimSpace(lines[i]), ".")
		if strings.Contains(lines[i], "Hell") {
			break
		}
		i++
	}
	// println("kek")
	// println(input)
	// println("sadas")
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
		frames = append(frames, Frame{
			Temp: temp,
			Hum:  hum,
		})
	}
	// println("lol")
}

func main() {
	// Setup Serial Port
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
            createGraph()
		}
	}
}
