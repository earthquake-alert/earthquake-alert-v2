package src

import (
	"fmt"
	"math"
	"strings"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/ktnyt/go-moji"
)

// マグニチュード表記をいい感じにする
func FormatMagnitude(m *jma.Magnitude) string {
	value := m.Value

	// 気象庁マグニチュード（Mj）の場合は、`M`と表記する
	prefix := "M"
	if m.Type == jma.MomentMagnitude {
		prefix = "Mw"
	}

	if value == "NaN" && m.Condition == "不明" {
		if Convert(m.Description, true) == "M不明" {
			return "不明"
		}
		return prefix + "8以上"
	}

	return prefix + value
}

// 震源の深さをいい感じにする
func FormatDepth(d *int) string {
	if d == nil || *d == 1 {
		return "不明"
	}

	if *d == 0 {
		return "ごく浅い"
	}
	if *d <= -700000 {
		return "700km以上"
	}

	if *d < -1000 {
		return fmt.Sprintf("%dkm", -(*d)/1000)
	}
	return fmt.Sprintf("%dm", -(*d))
}

// 日本測地系の緯度経度をいい感じにする
// lat, lonは度分秒で表す
func FormatLatLonDepth(lat float64, lon float64, depth *int) string {
	latPrefix := "北緯"
	if lat < 0 {
		latPrefix = "南緯"
	}
	lonPrefix := "東経"
	if lon < 0 {
		lonPrefix = "西経"
	}

	return fmt.Sprintf(
		"%s%.1f度 %s%.1f度 深さ %s",
		latPrefix,
		math.Abs(lat),
		lonPrefix,
		math.Abs(lon),
		FormatDepth(depth),
	)
}

// 文字列をイイカンジに整形する
//
// - 全ての改行を消す (isDeleteReturnsがtrueの場合)
// - 全角英数字を半角英数字に変換する
// - 全角スペースを半角スペースに変換する
func Convert(text string, isDeleteReturns bool) string {
	if text == "" {
		return ""
	}

	if isDeleteReturns {
		text = strings.ReplaceAll(text, "\n", "")
	}

	return moji.Convert(
		moji.Convert(
			text,
			moji.ZS,
			moji.HS,
		),
		moji.ZE,
		moji.HE,
	)
}
