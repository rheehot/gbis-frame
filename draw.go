package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"path/filepath"
	"reflect"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	panelW = 480
	panelH = 800
)

var (
	lastBuses []busArrival
	firstDraw = true

	icons map[string]*image.Image
	fonts map[float64]*font.Face
)

func init() {
	icons = make(map[string]*image.Image)
	fonts = make(map[float64]*font.Face)
}

func drawBusArrivalInfo(buses []busArrival) {
	if flagImageOut == "" && !flagUpdatePanel {
		return
	}

	if !firstDraw && len(buses) == len(lastBuses) {
		same := true
		for i, b := range buses {
			if !reflect.DeepEqual(b, lastBuses[i]) {
				same = false
				break
			}
		}
		if same {
			log.Println("same contents. skip drawing")
			return
		}
		// log.Println("update drawing")
	}
	firstDraw = false

	dc := gg.NewContext(panelW, panelH)
	dc.SetColor(color.White)
	dc.Clear()

	// dc.SetColor(color.Black)
	// dc.DrawRectangle(0, 0, 480, 80)
	// dc.Fill()

	drawStringAnchored(dc, stationName, 40, panelW/2, 30, 0.5, 0.5, color.Black) // 역이름

	var yOffset float64
	for _, b := range buses {
		if yOffset >= panelH {
			break
		}
		// yOffset := float64(160 * i)
		yOffset += 20
		drawImage(dc, filepath.Join("_resource", "directions_bus-60px.png"), 12, 75+yOffset) // 아이콘
		drawStringAnchored(dc, findBusNo(b.RouteID), 48,
			70, 80+24-5+yOffset, 0, 0.4, color.Black,
		) // 버스번호
		yOffset += 65
		if b.PredictTime1 != "" && b.LocationNo1 != "" {
			drawString(dc, "다음버스", 30,
				75, 80+24-5+yOffset,
			)
			drawString(dc, fmt.Sprintf("%s분 후 (%s 전)", b.PredictTime1, b.LocationNo1), 40,
				75, 80+24+30+10+yOffset,
			)
			yOffset += 90
		}
		if b.PredictTime2 != "" && b.LocationNo2 != "" {
			drawString(dc, "다다음버스", 30,
				75, 80+24-5+yOffset,
			)
			drawString(dc, fmt.Sprintf("%s분 후 (%s 전)", b.PredictTime2, b.LocationNo2), 40,
				75, 80+24+30+10+yOffset,
			)
			yOffset += 90
		}
		// yOffset += 5
	}

	lastBuses = buses
	if flagImageOut != "" {
		log.Println("drawBusArrivalInfo to", flagImageOut)
		dc.SavePNG(flagImageOut)
	}
	if flagUpdatePanel {
		log.Println("update Panel start")
		updatePanel(dc.Image())
		log.Println("update Panel done")
	}
}

func drawImage(dc *gg.Context, imgName string, x, y float64) {
	var img *image.Image
	var err error
	if i, ok := icons[imgName]; ok {
		img = i
	} else {
		img, err = loadImage(imgName)
		if err != nil {
			panic(err)
		}
		icons[imgName] = img
	}

	dc.DrawImage(*img, int(x), int(y))
	drawDebugCrossHair(dc, x, y)
}

func drawString(dc *gg.Context, text string, fontSize, x, y float64) {
	dc.SetRGB(0, 0, 0)
	// ff, err := loadFontFace(filepath.Join("_resource", "BMDOHYEON_ttf.ttf"), fontSize)
	ff, err := loadFontFace(fontSize)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(ff)
	dc.DrawString(text, x, y)
	drawDebugCrossHair(dc, x, y)
}

func drawStringAnchored(dc *gg.Context, text string, fontSize, x, y, ax, ay float64, c color.Color) {
	dc.SetColor(c)
	ff, err := loadFontFace(fontSize)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(ff)
	dc.DrawStringAnchored(text, x, y, ax, ay)
	drawDebugCrossHair(dc, x, y)
}

func drawDebugCrossHair(dc *gg.Context, x, y float64) {
	if !flagDebugGG {
		return // do nothing
	}
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(1)
	dc.DrawLine(x-10, y, x+10, y)
	dc.Stroke()
	dc.DrawLine(x, y-10, x, y+10)
	dc.Stroke()
}

func loadImage(name string) (*image.Image, error) {
	r := bytes.NewReader(MustAsset(name))
	im, _, err := image.Decode(r)
	return &im, err
}

func loadFontFace(points float64) (font.Face, error) {
	if ff, ok := fonts[points]; ok {
		return *ff, nil
	}
	path := filepath.Join("_resource", "BMDOHYEON_ttf.ttf")
	f, err := truetype.Parse(MustAsset(path))
	if err != nil {
		return nil, err
	}
	nface := truetype.NewFace(f, &truetype.Options{
		Size:    points,
		Hinting: font.HintingFull,
		// Hinting: font.HintingNone,
	})
	fonts[points] = &nface
	return nface, nil
}
