package models

type RequestPlace struct {
	RequestPlace       	string		`db:"request_place" json:"request_place"`
	RequestType       	string		`db:"request_type" json:"request_type"`
	Total 				int			`db:"total" json:"total"`
	TotalRequestCount	int			`db:"total_request_count" json:"total_request_count"`
}
