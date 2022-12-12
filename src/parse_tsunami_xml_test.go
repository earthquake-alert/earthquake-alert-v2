package src_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/stretchr/testify/require"
)

const TEST_DATA_PATH = "../test_data/jma_xml/"

// 津波情報のテストデータ
var TestData = []string{
	"32-39_11_02_120615_VTSE41.xml",
	"32-39_11_03_120615_VTSE51.xml",
	"32-39_11_06_120615_VTSE41.xml",
	"32-39_11_08_120615_VTSE51.xml",
	"32-39_11_09_120615_VTSE41.xml",
	"32-39_11_10_120615_VTSE51.xml",
	"32-39_11_11_120615_VTSE41.xml",
	"32-39_11_13_120615_VTSE41.xml",

	"32-39_12_02_191025_VTSE41.xml",
	"32-39_12_03_191025_VTSE51.xml",
	"32-39_12_05_191025_VTSE52.xml",
	"32-39_12_06_191025_VTSE51.xml",
	"32-39_12_07_191025_VTSE41.xml",
	"32-39_12_08_191025_VTSE51.xml",
	"32-39_12_10_191025_VTSE51.xml",
	"32-39_12_11_191025_VTSE51.xml",
	"32-39_12_12_191025_VTSE51.xml",
	"32-39_12_13_191025_VTSE41.xml",
	"32-39_12_14_191025_VTSE51.xml",
	"32-39_13_02_191025_VTSE41.xml",
	"32-39_13_03_191025_VTSE51.xml",
	"32-39_13_07_191025_VTSE41.xml",

	"38-39_02_01_191025_VTSE41.xml",
	"38-39_02_02_191025_VTSE51.xml",
	"38-39_02_03_191025_VTSE51.xml",
	"38-39_02_04_191025_VTSE41.xml",
	"38-39_02_05_191025_VTSE51.xml",
	"38-39_03_01_210805_VTSE41.xml",
	"38-39_03_02_210805_VTSE51.xml",
	"38-39_03_03_210805_VTSE51.xml",
}

func TestParseTsunami(t *testing.T) {
	for _, d := range TestData {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseTsunami(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseTsunami([]byte(row))
		require.Error(t, err)
	})
}

func TestStatus(t *testing.T) {
	statusData := map[string]src.Status{
		"32-39_11_02_120615_VTSE41.xml": src.Common,
		"32-39_12_05_191025_VTSE52.xml": src.Training,
	}

	for d, s := range statusData {
		t.Run(fmt.Sprintf("Test IsStatus %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			tsunami, err := src.ParseTsunami(row)
			require.NoError(t, err)

			require.Equal(t, tsunami.Status(), s)
		})
	}
}
