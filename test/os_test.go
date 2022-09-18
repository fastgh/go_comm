package test

import (
	"testing"

	"github.com/fastgh/go-comm"
	"github.com/stretchr/testify/require"
)

func Test_EnvironMap_happy(t *testing.T) {
	a := require.New(t)

	actual := comm.EnvironMap(nil)
	a.True(len(actual) > 0)
	a.True(len(actual["PATH"]) > 0)

	overrides := map[string]string{
		"k1":   "v1",
		"k2":   "v2",
		"PATH": "overrided_path",
	}

	actual = comm.EnvironMap(overrides)
	a.Equal("v1", actual["k1"])
	a.Equal("v2", actual["k2"])
	a.Equal("overrided_path", actual["PATH"])
}