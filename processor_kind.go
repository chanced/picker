package picker

type ProcessorKind string

const (
	ProcessorKindAppend           ProcessorKind = "append"
	ProcessorKindBytes            ProcessorKind = "bytes"
	ProcessorKindCircle           ProcessorKind = "cirlce"
	ProcessorKindCommunityID      ProcessorKind = "community_id"
	ProcessorKindConvert          ProcessorKind = "convert"
	ProcessorKindCSV              ProcessorKind = "csv"
	ProcessorKindDate             ProcessorKind = "date"
	ProcessorKindDateIndexName    ProcessorKind = "date_index_name"
	ProcessorKindDissect          ProcessorKind = "dissect"
	ProcessorKindDotExpander      ProcessorKind = "dot_expander"
	ProcessorKindDrop             ProcessorKind = "drop"
	ProcessorKindEnrich           ProcessorKind = "enrich"
	ProcessorKindFail             ProcessorKind = "fail"
	ProcessorKindFingerprint      ProcessorKind = "fingerprint"
	ProcessorKindForeach          ProcessorKind = "foreach"
	ProcessorKindGeoIP            ProcessorKind = "geoip"
	ProcessorKindGrok             ProcessorKind = "grok"
	ProcessorKindGsub             ProcessorKind = "gsub"
	ProcessorKindHTMLStrip        ProcessorKind = "html_strip"
	ProcessorKindInference        ProcessorKind = "inference"
	ProcessorKindJoin             ProcessorKind = "join"
	ProcessorKindJSON             ProcessorKind = "json"
	ProcessorKindKV               ProcessorKind = "kv"
	ProcessorKindLowercase        ProcessorKind = "lowercase"
	ProcessorKindNetworkDirection ProcessorKind = "network_direction"
	ProcessorKindPipeline         ProcessorKind = "pipeline"
	ProcessorKindRemove           ProcessorKind = "remove"
	ProcessorKindRename           ProcessorKind = "rename"
	ProcessorKindScript           ProcessorKind = "script"
	ProcessorKindSet              ProcessorKind = "set"
	ProcessorKindSetSecurityUser  ProcessorKind = "set_security_user"
	ProcessorKindSort             ProcessorKind = "sort"
	ProcessorKindSplit            ProcessorKind = "split"
	ProcessorKindTrim             ProcessorKind = "trim"
	ProcessorKindUppercase        ProcessorKind = "uppercase"
	ProcessorKindURLDecode        ProcessorKind = "urldecode"
	ProcessorKindURIParts         ProcessorKind = "uri_parts"
	ProcessorKindUserAgent        ProcessorKind = "user_agent"
)

var processorHandlers = map[ProcessorKind]func() Processor{}
