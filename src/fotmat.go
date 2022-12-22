package src

import (
	"fmt"
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
