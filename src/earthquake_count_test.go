package src_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/models"
	"github.com/stretchr/testify/require"
)

var TestEarthquakeCount = []string{
	"32-35_03_01_100514_VXSE60.xml",
	"32-35_10_02_220510_VXSE60.xml",
}

func TestParseEarthquakeCount(t *testing.T) {
	for _, d := range TestEarthquakeCount {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			row := LoadFile(d)

			_, err := src.ParseEarthquakeCount(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquakeCount([]byte(row))
		require.Error(t, err)
	})
}

func TestEarthquakeCountGetText(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		target := "32-35_03_01_100514_VXSE60.xml"
		row := LoadFile(target)

		ea, err := src.ParseEarthquakeCount(row)
		require.NoError(t, err)

		text, err := ea.GetText()
		require.NoError(t, err)

		require.Equal(t, text, "地震回数に関する情報をお知らせします。\n2008年8月24日15時ごろから2008年8月25日12時ごろまでの間に、1704回（うち震度1以上1回）の地震が発生しています。\n次の「地震回数に関する情報」は、26日18時00分頃に発表します。\n 8月24日15時過ぎから伊豆半島東方沖で地震が発生しています。この付近で発生した地震については、震度3以上の場合は「震源・震度情報」で発表しますが、震度2以下の場合は、「地震回数に関する情報」(本情報)で地震回数をまとめて発表します。")
	})

	t.Run("2", func(t *testing.T) {
		target := "32-35_10_02_220510_VXSE60.xml"
		row := LoadFile(target)

		ea, err := src.ParseEarthquakeCount(row)
		require.NoError(t, err)

		text, err := ea.GetText()
		require.NoError(t, err)

		require.Equal(t, text, "地震回数に関する情報を取り消します。\n先ほどの、地震回数に関する情報を取り消します。")
	})
}

func TestEarthquakeCountAssembly(t *testing.T) {
	ctx := context.Background()
	t.Run("DBに格納される", func(t *testing.T) {
		t.Run("1", func(t *testing.T) {
			target := "32-35_03_01_100514_VXSE60.xml"
			row := LoadFile(target)

			ea, err := src.ParseEarthquakeCount(row)
			require.NoError(t, err)

			err = ea.Assembly(ctx, DB)
			require.NoError(t, err)

			eventIds, err := ea.GetEventId()
			require.NoError(t, err)

			exists, err := models.EarthquakeCounts(
				models.EarthquakeCountWhere.EventID.EQ(eventIds[0]),
			).Exists(ctx, DB)
			require.NoError(t, err)
			require.True(t, exists)

			a, err := models.EarthquakeCounts(
				models.EarthquakeCountWhere.EventID.EQ(eventIds[0]),
			).One(ctx, DB)
			require.NoError(t, err)

			require.Equal(t, a.EventID, eventIds[0])
			require.NotNil(t, a.Created)
			require.NotNil(t, a.ID)
			require.NotNil(t, a.Date)
			require.Equal(t, a.Row, string(row))
		})

		t.Run("2", func(t *testing.T) {
			target := "32-35_10_02_220510_VXSE60.xml"
			row := LoadFile(target)

			ea, err := src.ParseEarthquakeCount(row)
			require.NoError(t, err)

			err = ea.Assembly(ctx, DB)
			require.NoError(t, err)

			eventIds, err := ea.GetEventId()
			require.NoError(t, err)

			exists, err := models.EarthquakeCounts(
				models.EarthquakeCountWhere.EventID.EQ(eventIds[0]),
			).Exists(ctx, DB)
			require.NoError(t, err)
			require.True(t, exists)
		})
	})
}
