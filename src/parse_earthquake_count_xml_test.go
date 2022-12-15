package src_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/stretchr/testify/require"
)

var TestEarthquakeCount = []string{
	"32-35_03_01_100514_VXSE60.xml",
	"32-35_10_02_220510_VXSE60.xml",
}

func TestParseEarthquakeCount(t *testing.T) {
	for _, d := range TestEarthquakeCount {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseEarthquakeCount(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquakeCount([]byte(row))
		require.Error(t, err)
	})
}
