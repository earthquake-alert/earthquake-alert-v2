package src_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/stretchr/testify/require"
)

var TestEarthquakeActivity = []string{
	"32-35_09_01_191111_VXSE56.xml",
	"32-35_09_02_220316_VXSE56.xml",
}

func TestParseEarthquakeActivity(t *testing.T) {
	for _, d := range TestEarthquakeActivity {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseEarthquakeActivity(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquakeActivity([]byte(row))
		require.Error(t, err)
	})
}
