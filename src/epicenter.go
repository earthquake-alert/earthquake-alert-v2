package src

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

var COORDINATE_REGEXP = regexp.MustCompile(`([+-][0-9.]+)([+-][0-9.]+)([+-][0-9]+)?/`)

type Epicenter struct {
	// 緯度経度
	// 震源要素が不明な場合はnilになる
	Lat *float64
	Lon *float64
	// 日本測地系かどうか
	IsDatumJapan bool

	// 深さ
	// 震源要素が不明だったり、深さ不明だった場合はnilになる
	Depth *int

	Name string
	// 遠地地震の場合、ここに震源決定機関が入る
	Source string

	H *jma.Hypocenter
}

func ParseEpicenter(h *jma.Hypocenter) (*Epicenter, error) {
	name := h.Area.Name
	if h.Area.NameFromMark != "" {
		// 日本の地震の場合で、NameFromMarkがある場合は追加する
		// `三陸沖（牡鹿半島の東南東130km付近）` のようになる
		name = fmt.Sprintf("%s（%s）", name, Convert(h.Area.NameFromMark, true))
	} else if h.Area.DetailedName != "" {
		// 遠地地震の場合、DetailedNameを使用する
		// `南太平洋 フィジー諸島`のようになる`
		name = fmt.Sprintf("%s %s", name, Convert(h.Area.DetailedName, true))
	}

	isJapanDatum := false
	if h.Area.Coordinate.Datum == "日本測地系" {
		isJapanDatum = true
	}

	var latG *float64
	var lonG *float64
	var depthG *int
	if h.Area.Coordinate.Value != "" {
		lat, lon, depth, err := ParseCoordinate(h.Area.Coordinate.Value)
		if err != nil {
			return nil, err
		}

		latG = &lat
		lonG = &lon
		depthG = &depth
	}

	return &Epicenter{
		Lat:          latG,
		Lon:          lonG,
		IsDatumJapan: isJapanDatum,

		Depth: depthG,

		Name:   name,
		Source: Convert(h.Source, true),

		H: h,
	}, nil
}

// 深さを返す
func (e *Epicenter) FormatDepth() string {
	if e.Depth == nil || *e.Depth == 1 {
		return "不明"
	}

	if *e.Depth == 0 {
		return "ごく浅い"
	}
	if *e.Depth <= -700000 {
		return "700km以上"
	}

	if *e.Depth < -1000 {
		return fmt.Sprintf("%dkm", -(*e.Depth)/1000)
	}
	return fmt.Sprintf("%dm", -(*e.Depth))
}

// 震源要素を解析する。
// 例:
// "-17.2+178.6-570000/"
// --> (-17.2, 178.6, -570000, nil)
//
// returns: (lat, lon, depth, error)
// depthは不明の場合は1、ごく浅い場合は0
func ParseCoordinate(v string) (float64, float64, int, error) {
	r := COORDINATE_REGEXP.FindStringSubmatch(v)

	lat, err := strconv.ParseFloat(r[1], 64)
	if err != nil {
		return 0, 0, 1, err
	}
	lon, err := strconv.ParseFloat(r[2], 64)
	if err != nil {
		return 0, 0, 1, err
	}

	depth := 1
	if len(r) == 4 && r[3] != "" {
		depth, err = strconv.Atoi(r[3])
		if err != nil {
			return 0, 0, 1, err
		}
	}

	return lat, lon, depth, nil
}
