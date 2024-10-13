package db

type Entity struct {
	Name   string        `json:"name"`
	Fields []EntityField `json:"fields"`
}
