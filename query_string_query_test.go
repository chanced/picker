package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestQueryStringQuery(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "query_string": {
			"query": "(new york city) OR (big apple)",
			"default_field": "content"
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: &picker.QueryParams{
			QueryString: picker.QueryStringQueryParams{
				Query:        "(new york city) OR (big apple)",
				DefaultField: "content",
			},
		},
	})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr *picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)
	sd2, err := sr.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd2), cmpjson.Diff(data, sd2))

	data = []byte(`{
		"query": {
		  "query_string" : {
			"fields" : ["content", "name^5"],
			"query" : "this AND that OR thus",
			"tie_breaker" : 0
		  }
		}
	  }`)
	s, err = picker.NewSearch(picker.SearchParams{
		Query: &picker.QueryParams{
			QueryString: picker.QueryStringQueryParams{
				Query:      "this AND that OR thus",
				Fields:     []string{"content", "name^5"},
				TieBreaker: 0,
			},
		},
	})
	assert.NoError(err)
	sd, err = s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)

}
