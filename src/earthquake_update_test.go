package src_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
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
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseEarthquakeUpdate(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquakeUpdate([]byte(row))
		require.Error(t, err)
	})
}

func TestEarthquakeUpdateAssembly(t *testing.T) {
	testPath := filepath.Join(TEST_DATA_PATH, "32-35_06_09_100915_VXSE61.xml")

	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	ctx := context.Background()
	db, err := src.NewConnectMySQL(ctx)
	require.NoError(t, err)

	err = ea.Assembly(ctx, db)
	require.Error(t, err)

	t.Run("正しくパースされる", func(t *testing.T) {
		require.Equal(t, ea.NewName, "岩手県内陸南部")
		require.Equal(t, *ea.NewLat, 39.0)
		require.Equal(t, *ea.NewLon, 140.9)
		require.Equal(t, *ea.NewDepth, -10000)
		require.Equal(t, ea.NewLatLonStr, "北緯39.0度 東経140.9度 深さ 10km")

		require.Equal(t, ea.NewMagnitude, "M7.2")
	})

	// t.Run("更新前の情報は無いため")
}
