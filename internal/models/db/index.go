package db

type Index struct {
	Name     string   `json:"name"`
	Method   string   `json:"method"`
	Table    string   `json:"table"`
	Fields   []string `json:"fields"`
	IsUnique bool     `json:"is_unique"`
}
