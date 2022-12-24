package src_test

import (
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/stretchr/testify/require"
)

func TestParseEventID(t *testing.T) {
	t.Run("通常", func(t *testing.T) {
		e, err := src.ParseEventID("20090811050711")
		require.NoError(t, err)
		require.Equal(t, e, []uint64{20090811050711})
	})

	t.Run("複数", func(t *testing.T) {
		e, err := src.ParseEventID("20090811050711 20100125161221")
		require.NoError(t, err)
		require.Equal(t, e, []uint64{20090811050711, 20100125161221})
	})
}
