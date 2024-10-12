package db

type EntityField struct {
	Name string    `json:"name"`
	Tags []string  `json:"tags"`
	Type FieldType `json:"type"`
}
