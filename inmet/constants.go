package inmet

import (
	"fmt"
	"strings"
	"time"
)

type Area int
type Satellite int
type Param int

const (
	AS Area = iota
	BR Area = iota
	CO Area = iota
	DF Area = iota
	N  Area = iota
	NE Area = iota
	S  Area = iota
	SE Area = iota
)

const (
	GOES     Satellite = iota
	GOESIM   Satellite = iota
	SATELITE Satellite = iota
)

const (
	// GOES params
	IV Param = iota
	TN Param = iota
	VA Param = iota
	VI Param = iota
	VP Param = iota

	// GOESIM params
	CH Param = iota

	// SATELITE parmas
	P    Param = iota
	V10  Param = iota
	V200 Param = iota
	V500 Param = iota
	V700 Param = iota
	V850 Param = iota
)

const apisatBaseUrl string = `https://apisat.inmet.gov.br`

type Info struct {
	Sigla, Nome string
}

type InfoResp []Info

type ImageResp struct {
	Nome, Satelite, Parametro, Data, Hora, Base64 string
}

type ImagesResp []ImageResp

func getSReqUrl(sat Satellite, are Area, par Param, dateTime time.Time) (url string, err error) {
	satstr, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	arestr, err := are.GetAreaStr()
	if err != nil {
		return
	}

	parstr, err := par.GetParamStr()
	if err != nil {
		return
	}

	utcdate := getUTCDate(dateTime)
	time := getTime(dateTime)

	url = fmt.Sprintf("%s/%s/%s/%s/%s/%s", apisatBaseUrl, satstr, arestr, parstr, utcdate, time)

	return
}

func getMReqUrl(sat Satellite, are Area, par Param, dateTime time.Time) (url string, err error) {
	satstr, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	arestr, err := are.GetAreaStr()
	if err != nil {
		return
	}

	parstr, err := par.GetParamStr()
	if err != nil {
		return
	}

	date := getDate(dateTime)

	url = fmt.Sprintf("%s/%s/%s/%s/%s", apisatBaseUrl, satstr, arestr, parstr, date)

	return
}

func getAReqUrl(sat Satellite) (url string, err error) {
	satstr, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	url = fmt.Sprintf("%s/areas/%s", apisatBaseUrl, satstr)

	return
}

func getPReqUrl(sat Satellite, are Area) (url string, err error) {
	satstr, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	arestr, err := are.GetAreaStr()
	if err != nil {
		return
	}

	url = fmt.Sprintf("%s/parametros/%s/%s", apisatBaseUrl, satstr, arestr)

	return
}

func getHReqUrl(sat Satellite, are Area, par Param, dateTime time.Time) (url string, err error) {
	satstr, err := sat.GetSatelliteStr()
	if err != nil {
		return
	}

	arestr, err := are.GetAreaStr()
	if err != nil {
		return
	}

	parstr, err := par.GetParamStr()
	if err != nil {
		return
	}

	utcdate := getUTCDate(dateTime)

	url = fmt.Sprintf("%s/horas/%s/%s/%s/%s", apisatBaseUrl, satstr, arestr, parstr, utcdate)

	return
}

func getDate(dateTime time.Time) string {
	return dateTime.Format("2006-01-02")
}

func getTime(dateTime time.Time) string {
	return dateTime.Format("15:04")
}

func getUTCDate(dateTime time.Time) string {
	return getDate(dateTime) + "T03:00:00.000Z"
}

func (are Area) GetAreaStr() (str string, err error) {
	switch are {
	case AS:
		str = "AS"

	case BR:
		str = "BR"

	case CO:
		str = "CO"

	case DF:
		str = "DF"

	case N:
		str = "N"

	case NE:
		str = "NE"

	case S:
		str = "S"

	case SE:
		str = "SE"

	default:
		str = "XX"
		err = fmt.Errorf("invalid area code '%d'", are)
	}

	return
}

func (sat Satellite) GetSatelliteStr() (str string, err error) {

	switch sat {
	case GOES:
		str = "GOES"

	case GOESIM:
		str = "GOESIM"

	case SATELITE:
		str = "SATELITE"

	default:
		str = "XX"
		err = fmt.Errorf("invalid satellite code '%d'", sat)
	}

	return
}

func (par Param) GetParamStr() (str string, err error) {
	switch par {
	case IV:
		str = "IV"

	case TN:
		str = "TN"

	case VA:
		str = "VA"

	case VI:
		str = "VI"

	case VP:
		str = "VP"

	case CH:
		str = "CH"

	case P:
		str = "P"

	case V10:
		str = "v10"

	case V200:
		str = "v200"

	case V500:
		str = "v500"

	case V700:
		str = "v700"

	case V850:
		str = "v850"

	default:
		str = "XX"
		err = fmt.Errorf("invalid param code '%d", par)
	}

	return
}

func GetAreaCode(str string) (are Area, err error) {
	switch strings.ToUpper(str) {
	case "AS":
		are = AS

	case "BR":
		are = BR

	case "CO":
		are = CO

	case "DF":
		are = DF

	case "N":
		are = N

	case "NE":
		are = NE

	case "S":
		are = S

	case "SE":
		are = SE

	default:
		err = fmt.Errorf("invalid area string '%s'", str)
	}

	return
}

func GetSatelliteCode(str string) (sat Satellite, err error) {

	switch strings.ToUpper(str) {
	case "GOES":
		sat = GOES

	case "GOESIM":
		sat = GOESIM

	case "SATELITE":
		sat = SATELITE

	default:
		err = fmt.Errorf("invalid satellite string '%s'", str)
	}

	return
}

func GetParamCode(str string) (par Param, err error) {
	switch strings.ToUpper(str) {
	case "IV":
		par = IV

	case "TN":
		par = TN

	case "VA":
		par = VA

	case "VI":
		par = VI

	case "VP":
		par = VP

	case "CH":
		par = CH

	case "P":
		par = P

	case "V10":
		par = V10

	case "V200":
		par = V200

	case "V500":
		par = V500

	case "V700":
		par = V700

	case "V850":
		par = V850

	default:
		err = fmt.Errorf("invalid param string '%s'", str)
	}

	return
}

func GetDefaultParam(sat Satellite) (par string, err error) {
	switch sat {
	case GOES:
		par = "IV"
	case GOESIM:
		par = "CH"
	case SATELITE:
		par = "P"
	default:
		err = fmt.Errorf("invalid satellite code '%d'", sat)
	}

	return
}
