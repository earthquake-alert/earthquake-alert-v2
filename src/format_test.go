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
