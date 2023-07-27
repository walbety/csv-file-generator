package canonical

type (
	FieldType string
)

const (
	UUID      FieldType = "UUID"
	Number    FieldType = "NUMBER"
	Date      FieldType = "DATE"
	Timestamp FieldType = "TIMESTAMP"
	Text      FieldType = "TEXT"
	Float     FieldType = "FLOAT"

	MAX_LINES = 170
)

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int    `json:"size"`
}
