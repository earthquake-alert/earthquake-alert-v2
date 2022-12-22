package src_test

import (
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/stretchr/testify/require"
)

func TestFormatMagnitude(t *testing.T) {
	t.Run("通常", func(t *testing.T) {
		m := &jma.Magnitude{
			Value:       "5.6",
			Type:        jma.JMAMagnitude,
			Description: "M6.6",
			Condition:   "",
		}
		require.Equal(t, src.FormatMagnitude(m), "M5.6")
	})

	t.Run("M8以上", func(t *testing.T) {
		m := &jma.Magnitude{
			Value:       "NaN",
			Type:        jma.JMAMagnitude,
			Description: "Ｍ８を超える巨大地震",
			Condition:   "不明",
		}
		require.Equal(t, src.FormatMagnitude(m), "M8以上")
	})

	t.Run("不明", func(t *testing.T) {
		m := &jma.Magnitude{
			Value:       "NaN",
			Type:        jma.JMAMagnitude,
			Description: "Ｍ不明",
			Condition:   "不明",
		}
		require.Equal(t, src.FormatMagnitude(m), "不明")
	})

	t.Run("モーメントマグニチュード", func(t *testing.T) {
		m := &jma.Magnitude{
			Value:       "5.6",
			Type:        jma.MomentMagnitude,
			Description: "M6.6",
			Condition:   "",
		}
		require.Equal(t, src.FormatMagnitude(m), "Mw5.6")
	})
}

func TestFormatDepth(t *testing.T) {
	t.Run("通常", func(t *testing.T) {
		d := -10000
		require.Equal(t, src.FormatDepth(&d), "10km")
	})

	t.Run("小数点以下mは切り捨て", func(t *testing.T) {
		d := -20050
		require.Equal(t, src.FormatDepth(&d), "20km")
	})

	t.Run("5km以内", func(t *testing.T) {
		d := 0
		require.Equal(t, src.FormatDepth(&d), "ごく浅い")
	})

	t.Run("不明", func(t *testing.T) {
		var d *int = nil
		require.Equal(t, src.FormatDepth(d), "不明")
	})

	t.Run("不明2", func(t *testing.T) {
		d := 1
		require.Equal(t, src.FormatDepth(&d), "不明")
	})

	t.Run("700km以上", func(t *testing.T) {
		d := -700000
		require.Equal(t, src.FormatDepth(&d), "700km以上")
	})

	// 仕様的に無いが、一応
	t.Run("m単位", func(t *testing.T) {
		d := -500
		require.Equal(t, src.FormatDepth(&d), "500m")
	})

	t.Run("1000m", func(t *testing.T) {
		d := -1000
		require.Equal(t, src.FormatDepth(&d), "1000m")
	})
}

func TestConvert(t *testing.T) {
	t.Run("全角英数を半角にする", func(t *testing.T) {
		text := "ＡＢＣ１４０３"

		require.Equal(t, src.Convert(text, false), "ABC1403")
	})

	t.Run("全角スペースを半角スペースにする", func(t *testing.T) {
		text := "123 あsだ　aaa"

		require.Equal(t, src.Convert(text, false), "123 あsだ aaa")
	})

	t.Run("改行を消す", func(t *testing.T) {
		text := "123\nabc\ncdf"

		require.Equal(t, src.Convert(text, true), "123abccdf")
	})

	t.Run("空文字", func(t *testing.T) {
		require.Equal(t, src.Convert("", false), "")
	})
}
