package picker_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestFunctionScoreQuery(t *testing.T) {
	assert := require.New(t)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			FunctionScore: &picker.FunctionScoreQuery{
				Functions: picker.Funcs{
					picker.ExpFunc{},
				},
			},
		},
	})
	_ = s
	assert.ErrorIs(err, picker.ErrFieldRequired)

	_, err = picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			FunctionScore: &picker.FunctionScoreQuery{
				Functions: picker.Funcs{
					picker.ExpFunc{
						Field: "field",
					},
				},
			},
		},
	})
	assert.ErrorIs(err, picker.ErrOriginRequired)

	_, err = picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			FunctionScore: &picker.FunctionScoreQuery{
				Functions: picker.Funcs{
					picker.ExpFunc{
						Field:  "field",
						Origin: "sdf",
					},
				},
			},
		},
	})
	assert.ErrorIs(err, picker.ErrScaleRequired)

	s, err = picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			FunctionScore: &picker.FunctionScoreQuery{
				Query: &picker.QueryParams{
					Term: &picker.TermQuery{
						Field:           "query_term_field",
						Value:           "query_term_value",
						Boost:           3,
						CaseInsensitive: true,
						Name:            "query_term",
					},
				},
				Functions: picker.Funcs{
					picker.ExpFunc{
						Field:  "fieldName",
						Origin: "sdf",
						Scale:  34,
						Weight: 21,
						Offset: 7,
						Decay:  "34",
						Filter: picker.TermQuery{
							Field:           "term_field",
							Value:           "term_value",
							Boost:           34,
							CaseInsensitive: true,
							Name:            "term_name",
						},
					},
				},
			},
		},
	})
	assert.NoError(err)
	data, err := json.MarshalIndent(s.Query().FunctionScore(), "", "  ")
	assert.NoError(err)
	fmt.Println(string(data))
	var fs picker.FunctionScoreClause
	err = json.Unmarshal(data, &fs)
	assert.NoError(err)
}