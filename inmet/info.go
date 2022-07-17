package inmet

import (
	"fmt"
	"time"
)

func PrintAreas(sat Satellite) (err error) {
	areas, err := GetAreas(sat)
	if err != nil {
		return
	}

	sattext, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	fmt.Printf("Possible areas for '%s':\n", sattext)
	for _, area := range areas {
		areatext, err := area.GetAreaStr()
		if err != nil {
			return err
		}
		fmt.Printf("\t%s\n", areatext)
	}

	return
}

func PrintParams(sat Satellite, are Area) (err error) {
	params, err := GetParams(sat, are)
	if err != nil {
		return
	}

	sattext, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	aretext, err := are.GetAreaStr()
	if err != nil {
		return
	}

	fmt.Printf("Possible Params for satellite '%s' in '%s':\n", sattext, aretext)
	for _, param := range params {
		paramtext, err := param.GetParamStr()
		if err != nil {
			return err
		}
		fmt.Printf("\t%s\n", paramtext)
	}

	return
}

func PrintHours(sat Satellite, are Area, par Param, dateTime time.Time) (err error) {
	hours, err := GetHours(sat, are, par, dateTime)
	if err != nil {
		return
	}

	sattext, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	aretext, err := are.GetAreaStr()
	if err != nil {
		return
	}

	paramtext, err := par.GetParamStr()
	if err != nil {
		return
	}

	date := dateTime.Format("06-01-02")

	fmt.Printf("Possible hours for satellite '%s' in '%s' with '%s' at '%s'\n", sattext, aretext, paramtext, date)
	for _, hour := range hours {
		fmt.Printf("\t%s\n", hour)
	}

	return
}
