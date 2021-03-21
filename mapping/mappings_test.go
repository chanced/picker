package mapping_test

import (
	"testing"

	"github.com/chanced/esutil/mapping"
	"github.com/stretchr/testify/require"
)

func TestMappings(t *testing.T) {
	assert := require.New(t)
	_ = assert
	m := mapping.NewMappings()
	a := mapping.NewAliasField()
	m.Properties.AddField("alias", a)

}