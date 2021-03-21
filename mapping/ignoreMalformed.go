package mapping

// WithIgnoreMalformed is a mapping with the ignore_malformed parameter
//
// Sometimes you don’t have much control over the data that you receive. One
// user may send a login field that is a date, and another sends a login field
// that is an email address.
//
// Trying to index the wrong data type into a field throws an exception by
// default, and rejects the whole document. The ignore_malformed parameter, if
// set to true, allows the exception to be ignored. The malformed field is not
// indexed, but other fields in the document are processed normally.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ignore-malformed.html
type WithIgnoreMalformed interface {
	// IgnoreMalformed determines if malformed numbers are ignored. If true,
	// malformed numbers are ignored. If false (default), malformed numbers
	// throw an exception and reject the whole document.
	IgnoreMalformed() bool
	// SetIgnoreMalformed sets IgnoreMalformed to v
	SetIgnoreMalformed(v bool)
}

// FieldWithIgnoreMalformed is a Field with the ignore_malformed parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ignore-malformed.html
type FieldWithIgnoreMalformed interface {
	Field
	WithIgnoreMalformed
}

// IgnoreMalformedParam is a mixin that adds the ignore_malformed parameter to
// mappings
//
// Sometimes you don’t have much control over the data that you receive. One
// user may send a login field that is a date, and another sends a login field
// that is an email address.
//
// Trying to index the wrong data type into a field throws an exception by
// default, and rejects the whole document. The ignore_malformed parameter, if
// set to true, allows the exception to be ignored. The malformed field is not
// indexed, but other fields in the document are processed normally.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ignore-malformed.html
type IgnoreMalformedParam struct {
	IgnoreMalformedValue *bool `bson:"ignore_malformed,omitempty" json:"ignore_malformed,omitempty"`
}

// IgnoreMalformed determines if malformed numbers are ignored. If true,
// malformed numbers are ignored. If false (default), malformed numbers throw an
// exception and reject the whole document.
func (im IgnoreMalformedParam) IgnoreMalformed() bool {
	if im.IgnoreMalformedValue == nil {
		return false
	}
	return *im.IgnoreMalformedValue
}

// SetIgnoreMalformed sets IgnoreMalformed to v
func (im *IgnoreMalformedParam) SetIgnoreMalformed(v bool) {
	im.IgnoreMalformedValue = &v
}