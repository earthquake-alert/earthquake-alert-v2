package src_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/models"
	"github.com/stretchr/testify/require"
)

var TestEarthquakeUpdate = []string{
	"32-35_06_09_100915_VXSE61.xml",
	"32-35_06_10_100915_VXSE61.xml",
	"32-35_07_09_100915_VXSE61.xml",
}

func TestParseEarthquakeUpdate(t *testing.T) {
	for _, d := range TestEarthquakeUpdate {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			row := LoadFile(d)

			_, err := src.ParseEarthquakeUpdate(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquakeUpdate([]byte(row))
		require.Error(t, err)
	})
}

func TestEarthquakeGetTitle(t *testing.T) {
	t.Run("通常", func(t *testing.T) {
		row := LoadFile("32-35_06_09_100915_VXSE61.xml")

		ea, err := src.ParseEarthquakeUpdate(row)
		require.NoError(t, err)

		require.Equal(t, ea.GetTitle(), "震源要素更新のお知らせ")
	})

	t.Run("取消報", func(t *testing.T) {
		row := LoadFile("32-35_06_10_100915_VXSE61.xml")

		ea, err := src.ParseEarthquakeUpdate(row)
		require.NoError(t, err)

		require.Equal(t, ea.GetTitle(), "【取消】震源要素更新のお知らせ")
	})
}

func TestEarthquakeUpdateAssembly(t *testing.T) {
	row := LoadFile("32-35_06_09_100915_VXSE61.xml")

	ea, err := src.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	eventId := RandomEventID()
	ea.Parsed.Head.EventID = fmt.Sprint(eventId)

	ctx := context.Background()

	// Earthquakeテーブルに同じEventIDのカラムを追加しておく
	err = InsertEarthquake(ctx, eventId)
	require.NoError(t, err)

	err = ea.Assembly(ctx, DB)
	require.NoError(t, err)

	t.Run("正しくパースされる", func(t *testing.T) {
		require.Equal(t, ea.NewName, "岩手県内陸南部")
		require.Equal(t, *ea.NewLat, 39.0)
		require.Equal(t, *ea.NewLon, 140.9)
		require.Equal(t, *ea.NewDepth, -10000)
		require.Equal(t, ea.NewLatLonStr, "北緯39.0度 東経140.9度 深さ 10km")

		require.Equal(t, ea.NewMagnitude, "M7.2")
	})

	t.Run("EarthquakeUpdatesテーブルに格納されている", func(t *testing.T) {
		update, err := models.EarthquakeUpdates(
			models.EarthquakeUpdateWhere.EventID.EQ(eventId),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, update.Lat.Float64, *ea.NewLat)
		require.Equal(t, update.Lon.Float64, *ea.NewLon)
		require.Equal(t, update.Depth.Int, *ea.NewDepth)
		require.Equal(t, update.Magnitude.String, "M7.2")
	})

	t.Run("Earthquakeテーブルに格納されている", func(t *testing.T) {
		e, err := models.Earthquakes(
			models.EarthquakeWhere.EventID.EQ(eventId),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, e.Lat.Float64, *ea.NewLat)
		require.Equal(t, e.Lon.Float64, *ea.NewLon)
		require.Equal(t, e.Depth.Int, *ea.NewDepth)
		require.Equal(t, e.Magnitude.String, "M7.2")
	})
}

func TestEarthquakeUpdateAssemblyNoEarthquakeTableColumns(t *testing.T) {
	row := LoadFile("32-35_06_09_100915_VXSE61.xml")

	ea, err := src.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	eventId := RandomEventID()
	ea.Parsed.Head.EventID = fmt.Sprint(eventId)

	ctx := context.Background()

	err = ea.Assembly(ctx, DB)
	require.NoError(t, err)

	t.Run("EarthquakeUpdatesテーブルには格納されている", func(t *testing.T) {
		update, err := models.EarthquakeUpdates(
			models.EarthquakeUpdateWhere.EventID.EQ(eventId),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, update.Lat.Float64, *ea.NewLat)
		require.Equal(t, update.Lon.Float64, *ea.NewLon)
		require.Equal(t, update.Depth.Int, *ea.NewDepth)
		require.Equal(t, update.Magnitude.String, "M7.2")
	})

	t.Run("Earthquakeテーブルは存在しないまま", func(t *testing.T) {
		e, err := models.Earthquakes(
			models.EarthquakeWhere.EventID.EQ(eventId),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, e)
	})
}

func TestEarthquakeUpdateAssemblyCancel(t *testing.T) {
	row := LoadFile("32-35_06_10_100915_VXSE61.xml")

	ea, err := src.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	eventId := RandomEventID()
	ea.Parsed.Head.EventID = fmt.Sprint(eventId)

	ctx := context.Background()

	// Earthquakeテーブルに同じEventIDのカラムを追加しておく
	err = InsertEarthquake(ctx, eventId)
	require.NoError(t, err)

	err = ea.Assembly(ctx, DB)
	require.NoError(t, err)

	t.Run("正しくパースされる", func(t *testing.T) {
		require.Equal(t, ea.NewName, "")
		require.Nil(t, ea.NewLat)
		require.Nil(t, ea.NewLon)
		require.Nil(t, ea.NewDepth)
		require.Equal(t, ea.NewLatLonStr, "")

		require.Equal(t, ea.NewMagnitude, "")
	})

	t.Run("取消報でもEarthquakeUpdatesテーブルには格納されている", func(t *testing.T) {
		update, err := models.EarthquakeUpdates(
			models.EarthquakeUpdateWhere.EventID.EQ(eventId),
		).One(ctx, DB)
		require.NoError(t, err)

		require.False(t, update.Lat.Valid)
		require.False(t, update.Lon.Valid)
		require.False(t, update.Depth.Valid)
		require.Equal(t, update.Magnitude.String, "")
	})
}
