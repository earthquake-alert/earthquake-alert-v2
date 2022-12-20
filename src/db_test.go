package src_test

import (
	"context"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/models"
	"github.com/stretchr/testify/require"
)

func TestDB(t *testing.T) {
	t.Run("接続可能", func(t *testing.T) {
		ctx := context.Background()

		db, err := src.NewConnectMySQL(ctx)
		require.NoError(t, err)

		_, err = models.Earthquakes().All(ctx, db)
		require.NoError(t, err)
	})
}
