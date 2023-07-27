package canonical

var (
	MapFileTypeStringToConst = map[string]FieldType{
		"UUID":   UUID,
		"NUMBER": Number,
		"DATE":   Date,
		"TEXT":   Text,
		"FLOAT":  Float,
	}
)
