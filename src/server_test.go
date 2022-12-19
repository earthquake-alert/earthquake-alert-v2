package src_test

import (
	"flag"
	"os"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/stretchr/testify/require"
)

// これをしないとテストが失敗するため追加している
// ref. https://stackoverflow.com/questions/27342973/custom-command-line-flags-in-gos-unit-tests
var _ = flag.Bool("test.sqldebug", false, "Turns on debug mode for SQL statements")
var _ = flag.String("test.config", "", "Overrides the default config")

func TestMain(m *testing.M) {
	src.Init("test")

	flag.Parse()

	code := m.Run()
	os.Exit(code)

}

func TestMode(t *testing.T) {
	require.Equal(t, src.Mode, "test")
}
