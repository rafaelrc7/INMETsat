package inmet

import (
	"fmt"
	"time"
)

func PrintAreas(sat Satellite) (err error) {
	sattext, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	url, err := getAReqUrl(sat)
	if err != nil {
		return
	}

	var areas InfoResp
	err = getInfo(url, &areas)
	if err != nil {
		return
	}

	fmt.Printf("Possible areas for '%s':\n", sattext)
	for _, area := range areas {
		fmt.Printf("\t%s\t%s\n", area.Sigla, area.Nome)
	}

	return
}

func PrintParams(sat Satellite, are Area) (err error) {
	sattext, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	aretext, err := are.GetAreaStr()
	if err != nil {
		return
	}

	url, err := getPReqUrl(sat, are)
	if err != nil {
		return
	}

	var params InfoResp
	err = getInfo(url, &params)
	if err != nil {
		return
	}

	fmt.Printf("Possible Params for satellite '%s' in '%s':\n", sattext, aretext)
	for _, param := range params {
		fmt.Printf("\t%s\t%s\n", param.Sigla, param.Nome)
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
