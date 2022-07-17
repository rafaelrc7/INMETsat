package main

import (
	"flag"
	"image/gif"
	"log"
	"os"
	"time"

	"github.com/rafaelrc7/inmet-doppler/inmet"
)

func main() {
	var err error
	var outName, satelliteIn, areaIn, paramIn string
	var delay, threads int
	var satellite inmet.Satellite
	var area inmet.Area
	var param inmet.Param

	flag.StringVar(&outName, "output", "out.gif", "Output file name")
	flag.StringVar(&outName, "o", "out.gif", "Output file name")
	flag.StringVar(&satelliteIn, "satellite", "GOES", "Satellite type")
	flag.StringVar(&satelliteIn, "s", "GOES", "Satellite type")
	flag.StringVar(&areaIn, "area", "BR", "Area code")
	flag.StringVar(&areaIn, "a", "BR", "Area code")
	flag.StringVar(&paramIn, "param", "DEFAULT", "Parameter")
	flag.StringVar(&paramIn, "p", "DEFAULT", "Parameter")
	flag.IntVar(&delay, "delay", 5, "GIF frame delay")
	flag.IntVar(&delay, "d", 5, "GIF frame delay")
	flag.IntVar(&threads, "threads", 8, "Number of threads used for image processing")
	flag.IntVar(&threads, "t", 8, "Number of threads used for image processing")
	flag.Parse()

	if satellite, err = inmet.GetSatelliteCode(satelliteIn); err != nil {
		log.Fatal(err)
	}

	if area, err = inmet.GetAreaCode(areaIn); err != nil {
		log.Fatal(err)
	}

	if paramIn == "DEFAULT" {
		paramIn, err = inmet.GetDefaultParam(satellite)
		if err != nil {
			log.Fatal(err)
		}
	}

	if param, err = inmet.GetParamCode(paramIn); err != nil {
		log.Fatal(err)
	}

	anim, err := inmet.GetAnimation(satellite, area, param, time.Now(), delay, true, threads)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(outName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gif.EncodeAll(f, &anim)
}
