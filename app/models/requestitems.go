package models

type RequestItems struct {
	RequestType       string		`db:"request_type" json:"request_type"`
	Total 				int			`db:"total" json:"total"`
}
