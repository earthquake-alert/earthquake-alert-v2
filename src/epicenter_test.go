package src_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/stretchr/testify/require"
)

func TestEpicenter(t *testing.T) {
	t.Run("通常", func(t *testing.T) {
		target := "32-39_11_05_120615_VXSE53.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := jma.ParseEarthquake(row)
		require.NoError(t, err)

		h, err := src.ParseEpicenter((*jma.Hypocenter)(&ea.Body.Earthquake.Hypocenter))
		require.NoError(t, err)

		t.Run("正しくパースできる", func(t *testing.T) {
			require.Equal(t, *h.Lat, 38.0)
			require.Equal(t, *h.Lon, 142.9)
			require.Equal(t, *h.Depth, -10000)
			require.True(t, h.IsDatumJapan)

			require.Equal(t, h.Name, "三陸沖（牡鹿半島の東南東130km付近）")
			require.Equal(t, h.Source, "")
		})
	})

	t.Run("遠地地震", func(t *testing.T) {
		target := "32-35_01_03_100514_VXSE53.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := jma.ParseEarthquake(row)
		require.NoError(t, err)

		h, err := src.ParseEpicenter((*jma.Hypocenter)(&ea.Body.Earthquake.Hypocenter))
		require.NoError(t, err)

		t.Run("正しくパースできる", func(t *testing.T) {
			require.Equal(t, *h.Lat, -17.2)
			require.Equal(t, *h.Lon, 178.6)
			require.Equal(t, *h.Depth, -570000)
			require.False(t, h.IsDatumJapan)

			require.Equal(t, h.Name, "南太平洋 フィジー諸島")
			require.Equal(t, h.Source, "PTWC")
		})
	})

	t.Run("震源要素不明", func(t *testing.T) {
		target := "32-39_05_01_100831_VXSE53_2.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := jma.ParseEarthquake(row)
		require.NoError(t, err)

		h, err := src.ParseEpicenter((*jma.Hypocenter)(&ea.Body.Earthquake.Hypocenter))
		require.NoError(t, err)

		t.Run("正しくパースできる", func(t *testing.T) {
			require.Nil(t, h.Lat)
			require.Nil(t, h.Lon)
			require.Nil(t, h.Depth)
			require.False(t, h.IsDatumJapan)
		})
	})
}

func TestFormatDepth(t *testing.T) {
	t.Run("通常", func(t *testing.T) {
		d := -10000
		h := src.Epicenter{
			Depth: &d,
		}

		require.Equal(t, h.FormatDepth(), "10km")
	})

	t.Run("小数点以下mは切り捨て", func(t *testing.T) {
		d := -20050
		h := src.Epicenter{
			Depth: &d,
		}

		require.Equal(t, h.FormatDepth(), "20km")
	})

	t.Run("5km以内", func(t *testing.T) {
		d := 0
		h := src.Epicenter{
			Depth: &d,
		}

		require.Equal(t, h.FormatDepth(), "ごく浅い")
	})

	t.Run("不明", func(t *testing.T) {
		h := src.Epicenter{
			Depth: nil,
		}

		require.Equal(t, h.FormatDepth(), "不明")
	})

	t.Run("不明2", func(t *testing.T) {
		d := 1
		h := src.Epicenter{
			Depth: &d,
		}

		require.Equal(t, h.FormatDepth(), "不明")
	})

	t.Run("700km以上", func(t *testing.T) {
		d := -700000
		h := src.Epicenter{
			Depth: &d,
		}

		require.Equal(t, h.FormatDepth(), "700km以上")
	})

	// 仕様的に無いが、一応
	t.Run("m単位", func(t *testing.T) {
		d := -500
		h := src.Epicenter{
			Depth: &d,
		}

		require.Equal(t, h.FormatDepth(), "500m")
	})

	t.Run("1000m", func(t *testing.T) {
		d := -1000
		h := src.Epicenter{
			Depth: &d,
		}

		require.Equal(t, h.FormatDepth(), "1000m")
	})
}

func TestParseCoordinate(t *testing.T) {
	t.Run("通常", func(t *testing.T) {
		v := "-17.2+178.6-570000/"

		lat, lon, depth, err := src.ParseCoordinate(v)
		require.NoError(t, err)

		require.Equal(t, lat, -17.2)
		require.Equal(t, lon, 178.6)
		require.Equal(t, depth, -570000)
	})

	t.Run("震源の深さが 5km より浅い場合", func(t *testing.T) {
		v := "+37.5+138.6+0/"

		lat, lon, depth, err := src.ParseCoordinate(v)
		require.NoError(t, err)

		require.Equal(t, lat, 37.5)
		require.Equal(t, lon, 138.6)
		require.Equal(t, depth, 0)
	})

	t.Run("震源の深さが 700km 以上の場合", func(t *testing.T) {
		v := "+37.5+138.6-700000/"

		lat, lon, depth, err := src.ParseCoordinate(v)
		require.NoError(t, err)

		require.Equal(t, lat, 37.5)
		require.Equal(t, lon, 138.6)
		require.Equal(t, depth, -700000)
	})

	t.Run("震源の深さが不明の場合", func(t *testing.T) {
		v := "+37.5+138.6/"

		lat, lon, depth, err := src.ParseCoordinate(v)
		require.NoError(t, err)

		require.Equal(t, lat, 37.5)
		require.Equal(t, lon, 138.6)
		require.Equal(t, depth, 1)
	})
}
