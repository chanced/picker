package mapping

// TODO: Locale?

var (
	_ WithDocValues       = (*DateField)(nil)
	_ WithFormat          = (*DateField)(nil)
	_ WithIgnoreMalformed = (*DateField)(nil)
	_ WithIndex           = (*DateField)(nil)
	_ WithNullValue       = (*DateField)(nil)
	_ WithStore           = (*DateField)(nil)
	_ WithMeta            = (*DateField)(nil)
)

func NewDateField() *DateField {
	return &DateField{BaseField: BaseField{MappingType: TypeDate}}
}

type DateField struct {
	BaseField            `bson:",inline" json:",inline"`
	DocValuesParam       `bson:",inline" json:",inline"`
	FormatParam          `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	IndexParam           `bson:",inline" json:",inline"`
	NullValueParam       `bson:",inline" json:",inline"`
	StoreParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

// DateNanoSecField is an addition to the DateField data type.
//
// However there is an important distinction between the two. The existing date
// data type stores dates in millisecond resolution. The date_nanos data type
// stores dates in nanosecond resolution, which limits its range of dates from
// roughly 1970 to 2262, as dates are still stored as a long representing
// nanoseconds since the epoch.
//
// Queries on nanoseconds are internally converted to range queries on this long
// representation, and the result of aggregations and stored fields is converted
// back to a string depending on the date format that is associated with the
// field.
//
// Date formats can be customised, but if no format is specified then it uses
// the default:
//
//  "strict_date_optional_time||epoch_millis"
//
// This means that it will accept dates with optional timestamps, which conform
// to the formats supported by strict_date_optional_time including up to nine
// second fractionals or milliseconds-since-the-epoch (thus losing precision on
// the nano second part). Using strict_date_optional_time will format the result
// up to only three second fractionals. To print and parse up to nine digits of
// resolution, use strict_date_optional_time_nanos.
//
// Limitations
//
// Aggregations are still on millisecond resolution, even when using a
// date_nanos field. This limitation also affects transforms.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/date_nanos.html
type DateNanoSecField struct {
	BaseField            `bson:",inline" json:",inline"`
	DocValuesParam       `bson:",inline" json:",inline"`
	FormatParam          `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	IndexParam           `bson:",inline" json:",inline"`
	NullValueParam       `bson:",inline" json:",inline"`
	StoreParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

func NewDateNanoSecField() *DateNanoSecField {
	return &DateNanoSecField{BaseField: BaseField{MappingType: TypeDateNanos}}
}