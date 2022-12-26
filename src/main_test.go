package src_test

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/models"
	"github.com/stretchr/testify/require"
)

const TEST_DATA_PATH = "../test_data/jma_xml/"

var DB *sql.DB

// これをしないとテストが失敗するため追加している
// ref. https://stackoverflow.com/questions/27342973/custom-command-line-flags-in-gos-unit-tests
var _ = flag.Bool("test.sqldebug", false, "Turns on debug mode for SQL statements")
var _ = flag.String("test.config", "", "Overrides the default config")

func TestMain(m *testing.M) {
	src.Init("test")

	ctx := context.Background()
	db, err := src.NewConnectMySQL(ctx)
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	err = ResetDBTable(ctx, db)
	if err != nil {
		panic(err)
	}

	flag.Parse()

	code := m.Run()
	os.Exit(code)
}

// テスト用のDBを初期化する
func ResetDBTable(ctx context.Context, db *sql.DB) error {
	_, err := models.Earthquakes().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.JmaXmlEntries().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.TwitterThreads().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.TsunamiConnects().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.TsunamiInfos().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.EarthquakeInfos().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.EarthquakeEpicenters().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.EarthquakeReports().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.EarthquakeActivities().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.EarthquakeUpdates().DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	_, err = models.EarthquakeCounts().DeleteAll(ctx, db)
	if err != nil {
		return err
	}

	return nil
}

func TestMode(t *testing.T) {
	require.Equal(t, src.C.Mode, "test")
}

func LoadFile(file string) []byte {
	path := filepath.Join(TEST_DATA_PATH, file)

	row, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return row
}
