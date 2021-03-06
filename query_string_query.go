package picker

type QueryStringer interface {
	QueryString() (*QueryStringQuery, error)
}

type QueryStringQueryParams struct {
	// (Required, string) Query string you wish to parse and use for search. See
	// Query string syntax.
	Query string `json:"query"`
	// (Optional, string) Default field you wish to search if no field is
	// provided in the query string.
	//
	// Defaults to the index.query.default_field index setting, which has a
	// default value of *. The * value extracts all fields that are eligible for
	// term queries and filters the metadata fields. All extracted fields are
	// then combined to build a query if no prefix is specified.
	//
	// Searching across all eligible fields does not include nested documents.
	// Use a nested query to search those documents.
	//
	// For mappings with a large number of fields, searching across all eligible
	// fields could be expensive.
	//
	// There is a limit on the number of fields that can be queried at once. It
	// is defined by the indices.query.bool.max_clause_count search setting,
	// which defaults to 1024.
	DefaultField string `json:"default_field,omitempty"`
	// (Optional, Boolean) If true, the wildcard characters * and ? are allowed
	// as the first character of the query string. Defaults to true.
	AllowLeadingWildcard interface{} `json:"allow_leading_wildcard,omitempty"`
	// (Optional, Boolean) If true, the query attempts to analyze wildcard terms
	// in the query string. Defaults to false.
	AnalyzeWildcard interface{} `json:"analyze_wildcard,omitempty"`
	// (Optional, string) Analyzer used to convert text in the query string into
	// tokens. Defaults to the index-time analyzer mapped for the default_field.
	// If no analyzer is mapped, the index’s default analyzer is used.
	Analyzer string `json:"analyzer,omitempty"`
	// (Optional, Boolean) If true, match phrase queries are automatically
	// created for multi-term synonyms. Defaults to true.
	AutoGenerateSynonymsPhraseQuery interface{} `json:"auto_generate_synonyms_phrase_query,omitempty"`
	// (Optional, float) Floating point number used to decrease or increase the
	// relevance scores of the query. Defaults to 1.0.
	//
	// Boost values are relative to the default value of 1.0. A boost value
	// between 0 and 1.0 decreases the relevance score. A value greater than 1.0
	// increases the relevance score.
	Boost interface{} `json:"boost,omitempty"`
	// (Optional, string) Default boolean logic used to interpret text in the
	// query string if no operators are specified. Valid values are:
	//
	// - OR (Default)
	//
	// - AND
	DefaultOperator Operator `json:"default_operator,omitempty"`
	// (Optional, Boolean) If true, enable position increments in queries
	// constructed from a query_string search. Defaults to true.
	EnablePositionIncrements interface{} `json:"enable_position_increments,omitempty"`

	// (Optional, array of strings) Array of fields you wish to search.
	//
	// You can use this parameter query to search across multiple fields. See
	// Search multiple fields.
	Fields []string `json:"fields,omitempty"`
	// (Optional, string) Maximum edit distance allowed for matching. See
	// Fuzziness for valid values and more information.
	Fuzziness string `json:"fuzziness,omitempty"`
	// (Optional, integer) Maximum number of terms to which the query expands
	// for fuzzy matching. Defaults to 50.
	FuzzyMaxExpansions interface{} `json:"fuzzy_max_expansions,omitempty"`
	// (Optional, Boolean) If true, edits for fuzzy matching include
	// transpositions of two adjacent characters (ab → ba). Defaults to true.
	FuzzyTranspositions interface{} `json:"fuzzy_transpositions,omitempty"`
	// (Optional, Boolean) If true, format-based errors, such as providing a
	// text value for a numeric field, are ignored. Defaults to false.
	Lenient bool `json:"lenient,omitempty"`
	// (Optional, integer) Maximum number of automaton states required for the
	// query. Default is 10000.
	//
	// Elasticsearch uses Apache Lucene internally to parse regular expressions.
	// Lucene converts each regular expression to a finite automaton containing
	// a number of determinized states.
	//
	// You can use this parameter to prevent that conversion from
	// unintentionally consuming too many resources. You may need to increase
	// this limit to run complex regular expressions.
	MaxDeterminizedStates interface{} `json:"max_determinized_states,omitempty"`
	// (Optional, string) Minimum number of clauses that must match for a
	// document to be returned. See the minimum_should_match parameter for valid
	// values and more information. See How minimum_should_match works for an
	// example.
	MinimumShouldMatch string `json:"minimum_should_match,omitempty"`
	// (Optional, string) Analyzer used to convert quoted text in the query
	// string into tokens. Defaults to the search_quote_analyzer mapped for the
	// default_field.
	//
	// For quoted text, this parameter overrides the analyzer specified in the
	// analyzer parameter.
	QuoteAnalyzer string `json:"quote_analyzer,omitempty"`
	// (Optional, integer) Maximum number of positions allowed between matching
	// tokens for phrases. Defaults to 0. If 0, exact phrase matches are
	// required. Transposed terms have a slop of 2.
	PhraseSlop interface{} `json:"phrase_slop,omitempty"`
	// (Optional, string) Suffix appended to quoted text in the query string.
	//
	// You can use this suffix to use a different analysis method for exact
	// matches. See Mixing exact search with stemming.
	QuoteFieldSuffix string `json:"quote_field_suffix,omitempty"`
	// (Optional, string) Method used to rewrite the query. For valid values and
	// more information, see the rewrite parameß
	Rewrite Rewrite `json:"rewrite,omitempty"`
	// (Optional, string) Coordinated Universal Time (UTC) offset or IANA time
	// zone used to convert date values in the query string to UTC.
	//
	// Valid values are ISO 8601 UTC offsets, such as +01:00 or -08:00, and IANA
	// time zone IDs, such as America/Los_Angeles.
	TimeZone   string      `json:"time_zone,omitempty"`
	TieBreaker interface{} `json:"tie_breaker,omitempty"`
	Name       string      `json:"_name,omitempty"`
	completeClause
}

func (QueryStringQueryParams) Kind() QueryKind {
	return QueryKindQueryString
}
func (p QueryStringQueryParams) Clause() (QueryClause, error) {
	return p.QueryString()
}
func (p QueryStringQueryParams) QueryString() (*QueryStringQuery, error) {
	q := &QueryStringQuery{}
	var err error
	q.SetAnalyzer(p.Analyzer)
	q.SetLenient(p.Lenient)
	q.SetTimeZone(p.TimeZone)
	q.SetDefaultField(p.DefaultField)
	q.SetFuzziness(p.Fuzziness)
	q.SetFields(p.Fields)
	q.SetName(p.Name)
	q.SetMinimumShouldMatch(p.MinimumShouldMatch)
	q.SetQuoteAnalyzer(p.QuoteAnalyzer)
	q.SetQuoteFieldSuffix(p.QuoteFieldSuffix)
	err = q.SetAllowLeadingWildcard(p.AllowLeadingWildcard)
	if err != nil {
		return q, err
	}
	err = q.SetAnalyzeWildcard(p.AnalyzeWildcard)
	if err != nil {
		return q, err
	}
	err = q.SetAutoGenerateSynonymsPhraseQuery(p.AutoGenerateSynonymsPhraseQuery)
	if err != nil {
		return q, err
	}
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, err
	}
	err = q.SetDefaultOperator(p.DefaultOperator)
	if err != nil {
		return q, err
	}
	err = q.SetEnablePositionIncrements(p.EnablePositionIncrements)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyMaxExpansions(p.FuzzyMaxExpansions)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyTranspositions(p.FuzzyTranspositions)
	if err != nil {
		return q, err
	}
	err = q.SetMaxDeterminizedStates(p.MaxDeterminizedStates)
	if err != nil {
		return q, err
	}
	err = q.SetPhraseSlop(p.PhraseSlop)
	if err != nil {
		return q, err
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return q, err
	}
	err = q.SetRewrite(p.Rewrite)
	if err != nil {
		return q, err
	}
	err = q.SetTieBreaker(p.TieBreaker)
	if err != nil {
		return q, err
	}
	return q, nil
}

type QueryStringQuery struct {
	query        string
	defaultField string
	fields       []string
	quoteFieldSuffixParam
	quoteAnalyzerParam
	phraseSlopParam
	maxDeterminizedStatesParam
	fuzzyMaxExpansionsParam
	allowLeadingWildcardParam
	analyzeWildcardParam
	defaultOperatorParam
	tieBreakerParam
	timeZoneParam
	fuzzinessParam
	fuzzyTranspositionsParam
	lenientParam
	minimumShouldMatchParam
	rewriteParam
	analyzerParam
	autoGenerateSynonymsPhraseQueryParam
	boostParam
	enablePositionIncrementsParam
	completeClause
	nameParam
}

func (qs *QueryStringQuery) Clause() (QueryClause, error) {
	return qs, nil
}
func (qs *QueryStringQuery) QueryString() (*QueryStringQuery, error) {
	return qs, nil
}
func (qs *QueryStringQuery) IsEmpty() bool {
	return qs == nil || len(qs.query) == 0
}
func (qs *QueryStringQuery) Clear() {
	if qs == nil {
		return
	}
	*qs = QueryStringQuery{}
}

func (QueryStringQuery) Kind() QueryKind {
	return QueryKindQueryString
}
func (qs QueryStringQuery) Query() string {
	return qs.query
}
func (qs *QueryStringQuery) SetQuery(q string) error {
	if len(q) == 0 {
		return ErrQueryRequired
	}
	qs.query = q
	return nil
}
func (qs QueryStringQuery) DefaultField() string {
	return qs.defaultField
}
func (qs *QueryStringQuery) SetDefaultField(field string) {
	qs.defaultField = field
}
func (qs QueryStringQuery) Fields() []string {
	return qs.fields
}
func (qs *QueryStringQuery) SetFields(fields []string) {
	qs.fields = fields
}

func (qs *QueryStringQuery) UnmarshalJSON(data []byte) error {
	q := queryStringQuery{}
	err := q.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	v, err := q.QueryString()
	if err != nil {
		return err
	}
	*qs = v
	return nil
}

func (qs QueryStringQuery) MarshalJSON() ([]byte, error) {
	return queryStringQuery{
		Query:                           qs.query,
		DefaultField:                    qs.defaultField,
		AllowLeadingWildcard:            qs.allowLeadingWildcard.Value(),
		AnalyzeWildcard:                 qs.analyzeWildcard.Value(),
		Analyzer:                        qs.analyzer,
		AutoGenerateSynonymsPhraseQuery: qs.autoGenerateSynonymsPhraseQuery.Value(),
		Boost:                           qs.boost.Value(),
		DefaultOperator:                 qs.defaultOperator,
		EnablePositionIncrements:        qs.enablePositionIncrements.Value(),
		Fields:                          qs.fields,
		Fuzziness:                       qs.fuzziness,
		FuzzyMaxExpansions:              qs.fuzzyMaxExpansions.Value(),
		FuzzyTranspositions:             qs.fuzzyTranspositions.Value(),
		Lenient:                         qs.lenient,
		MaxDeterminizedStates:           qs.maxDeterminizedStates.Value(),
		MinimumShouldMatch:              qs.minimumShouldMatch,
		QuoteAnalyzer:                   qs.quoteAnalyzer,
		PhraseSlop:                      qs.phraseSlop.Value(),
		QuoteFieldSuffix:                qs.quoteFieldSuffix,
		Rewrite:                         qs.rewrite,
		TimeZone:                        qs.timeZone,
		TieBreaker:                      qs.tieBreaker.Value(),
		Name:                            qs.name,
	}.MarshalJSON()
}

//easyjson:json
type queryStringQuery struct {
	Query                           string      `json:"query"`
	DefaultField                    string      `json:"default_field,omitempty"`
	AllowLeadingWildcard            interface{} `json:"allow_leading_wildcard,omitempty"`
	AnalyzeWildcard                 interface{} `json:"analyze_wildcard,omitempty"`
	Analyzer                        string      `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery interface{} `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           interface{} `json:"boost,omitempty"`
	DefaultOperator                 Operator    `json:"default_operator,omitempty"`
	EnablePositionIncrements        interface{} `json:"enable_position_increments,omitempty"`
	Fields                          []string    `json:"fields,omitempty"`
	Fuzziness                       string      `json:"fuzziness,omitempty"`
	FuzzyMaxExpansions              interface{} `json:"fuzzy_max_expansions,omitempty"`
	FuzzyTranspositions             interface{} `json:"fuzzy_transpositions,omitempty"`
	Lenient                         bool        `json:"lenient,omitempty"`
	MaxDeterminizedStates           interface{} `json:"max_determinized_states,omitempty"`
	MinimumShouldMatch              string      `json:"minimum_should_match,omitempty"`
	QuoteAnalyzer                   string      `json:"quote_analyzer,omitempty"`
	PhraseSlop                      interface{} `json:"phrase_slop,omitempty"`
	QuoteFieldSuffix                string      `json:"quote_field_suffix,omitempty"`
	Rewrite                         Rewrite     `json:"rewrite,omitempty"`
	TimeZone                        string      `json:"time_zone,omitempty"`
	TieBreaker                      interface{} `json:"tie_breaker,omitempty"`
	Name                            string      `json:"_name,omitempty"`
}

func (p queryStringQuery) QueryString() (QueryStringQuery, error) {
	q := QueryStringQuery{}
	var err error
	q.SetAnalyzer(p.Analyzer)
	q.SetLenient(p.Lenient)
	q.SetTimeZone(p.TimeZone)
	q.SetDefaultField(p.DefaultField)
	q.SetFuzziness(p.Fuzziness)
	q.SetFields(p.Fields)
	q.SetName(p.Name)
	q.SetMinimumShouldMatch(p.MinimumShouldMatch)
	q.SetQuoteAnalyzer(p.QuoteAnalyzer)
	q.SetQuoteFieldSuffix(p.QuoteFieldSuffix)
	err = q.SetAllowLeadingWildcard(p.AllowLeadingWildcard)
	if err != nil {
		return q, err
	}
	err = q.SetAnalyzeWildcard(p.AnalyzeWildcard)
	if err != nil {
		return q, err
	}
	err = q.SetAutoGenerateSynonymsPhraseQuery(p.AutoGenerateSynonymsPhraseQuery)
	if err != nil {
		return q, err
	}
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, err
	}
	err = q.SetDefaultOperator(p.DefaultOperator)
	if err != nil {
		return q, err
	}
	err = q.SetEnablePositionIncrements(p.EnablePositionIncrements)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyMaxExpansions(p.FuzzyMaxExpansions)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyTranspositions(p.FuzzyTranspositions)
	if err != nil {
		return q, err
	}
	err = q.SetMaxDeterminizedStates(p.MaxDeterminizedStates)
	if err != nil {
		return q, err
	}
	err = q.SetPhraseSlop(p.PhraseSlop)
	if err != nil {
		return q, err
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return q, err
	}
	err = q.SetRewrite(p.Rewrite)
	if err != nil {
		return q, err
	}
	err = q.SetTieBreaker(p.TieBreaker)
	if err != nil {
		return q, err
	}
	return q, nil
}
