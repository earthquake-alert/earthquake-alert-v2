package src_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/stretchr/testify/require"
)

// 震度速報、震源に関する情報、震源・震度に関する情報のテストデータ
var TestEarthquakeData = []string{
	// 震度速報
	"32-35_01_01_100806_VXSE51.xml",
	"32-35_04_01_100831_VXSE51.xml",
	"32-35_04_02_100831_VXSE51.xml",
	"32-35_04_05_100831_VXSE51.xml",
	"32-35_07_02_100915_VXSE51.xml",
	"32-35_07_03_100915_VXSE51.xml",
	"32-35_07_04_100915_VXSE51.xml",
	"32-35_07_05_100915_VXSE51.xml",
	"32-35_08_01_100915_VXSE51.xml",
	"32-35_08_02_100915_VXSE51.xml",
	"32-35_08_03_100915_VXSE51.xml",
	"32-35_08_04_100915_VXSE51.xml",
	"32-35_08_05_100915_VXSE51.xml",
	"32-39_11_01_120615_VXSE51.xml",
	"32-35_10_01_220510_VXSE51.xml",
	// 震源に関する情報
	"32-35_01_02_100514_VXSE52.xml",
	"32-35_04_03_100831_VXSE52.xml",
	"32-35_06_01_100915_VXSE52.xml",
	"32-35_06_02_100915_VXSE52.xml",
	"32-35_08_06_100915_VXSE52.xml",
	"33_12_01_120615_VXSE52.xml",
	// 震源・震度に関する情報
	"32-35_01_03_100514_VXSE53.xml",
	"32-35_01_03_100806_VXSE53.xml",
	"32-35_04_04_100831_VXSE53.xml",
	"32-35_04_06_100831_VXSE53.xml",
	"32-35_06_03_100915_VXSE53.xml",
	"32-35_06_04_100915_VXSE53.xml",
	"32-35_06_05_100915_VXSE53.xml",
	"32-35_06_06_100915_VXSE53.xml",
	"32-35_07_06_100915_VXSE53.xml",
	"32-35_07_07_100915_VXSE53.xml",
	"32-35_08_07_100915_VXSE53.xml",
	"32-35_08_08_100915_VXSE53.xml",
	"32-39_05_01_100831_VXSE53.xml",
	"32-39_05_02_100831_VXSE53.xml",
	"32-39_05_04_100831_VXSE53.xml",
	"32-39_05_06_100831_VXSE53.xml",
	"32-39_05_07_100831_VXSE53.xml",
	"32-39_05_08_100831_VXSE53.xml",
	"32-39_05_12_100831_VXSE53.xml",
	"32-39_05_12_100915_VXSE53.xml",
	"32-39_11_05_120615_VXSE53.xml",
}

func TestParseEarthquake(t *testing.T) {
	for _, d := range TestEarthquakeData {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseEarthquake(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquake([]byte(row))
		require.Error(t, err)
	})
}
