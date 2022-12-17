package src_test

import (
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/stretchr/testify/require"
)

type TestTemplateObj struct {
	Test string
}

func TestTemplate(t *testing.T) {
	fileName := "test.tmpl"

	obj := TestTemplateObj{Test: "hoge"}

	value, err := src.Template(fileName, obj)
	require.NoError(t, err)

	require.Equal(t, value, "123 abc hoge\n")
}
