package db

type FieldKind int

const (
	FIELD_VALUE FieldKind = iota
	FIELD_REF
)

type FieldType struct {
	Kind  FieldKind `json:"kind"`
	Type  string    `json:"type,omitempty"`
	Table string    `json:"table,omitempty"`
	Field string    `json:"field,omitempty"`
}
