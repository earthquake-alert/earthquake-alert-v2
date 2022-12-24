package src_test

import (
	"testing"
	"time"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/stretchr/testify/require"
)

func TestParseEventID(t *testing.T) {
	t.Run("通常", func(t *testing.T) {
		e, err := src.ParseEventID("20090811050711")
		require.NoError(t, err)
		require.Equal(t, e, []uint64{20090811050711})
	})

	t.Run("複数", func(t *testing.T) {
		e, err := src.ParseEventID("20090811050711 20100125161221")
		require.NoError(t, err)
		require.Equal(t, e, []uint64{20090811050711, 20100125161221})
	})
}

func TestParseDate(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		parsed, err := src.ParseDate("2022-12-15T06:12:47+09:00")
		require.NoError(t, err)

		require.Equal(t, parsed.Year(), 2022)
		require.Equal(t, parsed.Month(), time.Month(12))
		require.Equal(t, parsed.Day(), 15)
		require.Equal(t, parsed.Hour(), 6)
		require.Equal(t, parsed.Minute(), 12)
	})
}

func TestParsePosition(t *testing.T) {
	t.Run("正しく変換できている", func(t *testing.T) {
		require.Equal(t, src.ParsePosition(34.8), 35.33)
		require.Equal(t, src.ParsePosition(138.5), 138.83)
	})

	t.Run("負の値でも正しく変換できる", func(t *testing.T) {
		require.Equal(t, src.ParsePosition(-36.1), -35.83)
		require.Equal(t, src.ParsePosition(-072.6), -71.0)
	})

	t.Run("分が0の場合", func(t *testing.T) {
		require.Equal(t, src.ParsePosition(40.0), 40.0)
		require.Equal(t, src.ParsePosition(-40.0), -40.0)
	})
}
