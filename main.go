package main

import (
	"flag"
	"fmt"
	"image/gif"
	"log"
	"os"
	"time"

	"github.com/rafaelrc7/inmetsat/inmet"
)

const (
	version = "1.0.0"
	name    = "INMETsat"
	mail    = "contact@rafaelrc.com"
)

const usage = `Usage: %s [OPTION...] [MODE]
	Simple client for the INMET satellite API.

	Modes:
	animation					Default. Generates animation based on settings.
	areas						Print available areas.
	params						Print available params.
	hours						Print available areas.

	Options:
	-v, --version               Shows current app version
	-o, --output FILE           Set output gif name, by default "out.gif".
	-s, --satellite MODE        Select satellite mode: GOES, GOESIM or SATELLITE.
	-a, --area AREA             Region code, depends on satellite mode.
	-p, --param PARAM           Query parameter, depends both on satellite and area.
	-d, --delay DELAY           Set gif frame delay, in 100ths of a second.
	-t, --threads THREADS       Set number of threads used for image processing.

	Report bugs to <%s>
`

const (
	animation = iota
	areas
	params
	hours
)

func main() {
	var err error
	var outName, satelliteIn, areaIn, paramIn string
	var showVersion bool
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
	flag.BoolVar(&showVersion, "version", false, "Show app version")
	flag.BoolVar(&showVersion, "v", false, "Show app version")
	flag.Usage = func() {
		if len(os.Args) > 0 {
			fmt.Printf(usage, os.Args[0], mail)
		} else {
			fmt.Printf(usage, "inmetsat", mail)
		}
	}
	flag.Parse()

	if showVersion {
		fmt.Printf("%s %s\n", name, version)
		os.Exit(0)
	}

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

	args := flag.NArg()
	mode := animation
	if args > 0 {
		off := len(os.Args) - args
		switch os.Args[off] {
		case "animation":
			mode = animation
		case "areas":
			mode = areas
		case "params":
			mode = params
		case "hours":
			mode = hours
		default:
			log.Fatalf("Unexpected mode '%s'\n", os.Args[off])
		}
	}

	switch mode {
	case animation:
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

	case areas:
		inmet.PrintAreas(satellite)

	case params:
		inmet.PrintParams(satellite, area)

	case hours:
		inmet.PrintHours(satellite, hours, param, time.Now())
	}
}
