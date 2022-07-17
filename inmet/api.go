package inmet

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func GetAnimation(sat Satellite, area Area, param Param, dateTime time.Time, delay int, repeat bool, threads int) (anim gif.GIF, err error) {
	imgs, err := GetImages(sat, area, param, dateTime)
	if err != nil {
		return
	}

	if repeat {
		anim.LoopCount = 0
	} else {
		anim.LoopCount = -1
	}

	log.Println("Processing images...")

	anim.Image = make([]*image.Paletted, len(imgs))
	for range imgs {
		anim.Delay = append(anim.Delay, delay)
	}

	var wg sync.WaitGroup
	wg.Add(threads)

	perThreadLoad := len(imgs) / threads
	perThreadRest := len(imgs) % threads
	for i := 0; i < threads; i++ {
		go func(wg *sync.WaitGroup, id int) {
			offset := id * perThreadLoad
			for j := 0; j < perThreadLoad && offset+j < len(imgs); j++ {
				log.Printf("Image %d/%d\n", offset+j+1, len(imgs))
				bounds := imgs[offset+j].Bounds()
				palettedImg := image.NewPaletted(bounds, palette.Plan9)
				draw.Draw(palettedImg, palettedImg.Rect, imgs[offset+j], bounds.Min, draw.Over)
				anim.Image[offset+j] = palettedImg
			}

			if id < perThreadRest {
				extra := threads*perThreadLoad + id

				log.Printf("Image %d/%d\n", extra+1, len(imgs))
				bounds := imgs[extra].Bounds()
				palettedImg := image.NewPaletted(bounds, palette.Plan9)
				draw.Draw(palettedImg, palettedImg.Rect, imgs[extra], bounds.Min, draw.Over)
				anim.Image[extra] = palettedImg

			}
			wg.Done()
		}(&wg, i)
	}

	wg.Wait()
	log.Print("Done.")

	return
}

func GetImages(sat Satellite, area Area, param Param, dateTime time.Time) (images []image.Image, err error) {
	validAreas, err := GetAreas(sat)
	if err != nil {
		return
	}
	if !contains(validAreas, area) {
		sats, _ := sat.GetSatelliteStr()
		ares, _ := area.GetAreaStr()
		return nil, fmt.Errorf("invalid area '%s' for satellite '%s'", ares, sats)
	}

	validParams, err := GetParams(sat, area)
	if err != nil {
		return
	}
	if !contains(validParams, param) {
		sats, _ := sat.GetSatelliteStr()
		ares, _ := area.GetAreaStr()
		pars, _ := param.GetParamStr()
		return nil, fmt.Errorf("invalid param '%s' for satellite '%s' and area '%s'", pars, sats, ares)
	}

	url, err := getMReqUrl(sat, area, param, dateTime)
	if err != nil {
		return
	}

	log.Printf("Getting images from '%s'...", url)
	var result ImagesResp
	err = imagesRequest(url, &result)
	if err != nil {
		return
	}

	log.Println("Decoding images...")
	for i := len(result) - 1; i >= 0; i-- {
		imageInfo := strings.Split(result[i].Base64, ";base64,")
		mime := imageInfo[0]
		imgSrc, err := base64.StdEncoding.DecodeString(imageInfo[1])
		if err != nil {
			return nil, err
		}

		switch mime {
		case "data:image/jpg":
			img, err := jpeg.Decode(bytes.NewReader(imgSrc))
			if err != nil {
				return nil, err
			}
			images = append(images, img)

		default:
			return nil, fmt.Errorf("unexpected image format '%s'", mime)
		}
	}

	return
}

func imagesRequest(url string, result *ImagesResp) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("query failed: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(result)

	return
}

func GetAreas(sat Satellite) (ars []Area, err error) {
	url, err := getAReqUrl(sat)
	if err != nil {
		return
	}

	var result InfoResp
	err = getInfo(url, &result)
	if err != nil {
		return
	}

	for _, info := range result {
		var are Area
		are, err = GetAreaCode(info.Sigla)
		if err != nil {
			return
		}
		ars = append(ars, are)
	}

	return
}

func GetParams(sat Satellite, are Area) (prs []Param, err error) {
	url, err := getPReqUrl(sat, are)
	if err != nil {
		return
	}

	var result InfoResp
	err = getInfo(url, &result)
	if err != nil {
		return
	}

	for _, info := range result {
		var par Param
		par, err = GetParamCode(info.Sigla)
		if err != nil {
			return
		}
		prs = append(prs, par)
	}

	return
}

func GetHours(sat Satellite, are Area, par Param, dateTime time.Time) (dates []string, err error) {
	url, err := getHReqUrl(sat, are, par, dateTime)
	if err != nil {
		return
	}

	var result InfoResp
	err = getInfo(url, &result)
	if err != nil {
		return
	}

	for _, info := range result {
		dates = append(dates, info.Sigla)
	}

	return
}

func getInfo(url string, result *InfoResp) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("query failed: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(result)

	return
}

func contains[T comparable](arr []T, v T) bool {
	for _, el := range arr {
		if v == el {
			return true
		}
	}

	return false
}
