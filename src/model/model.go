package model

type Student struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Subject1 int  `json:"subject1"`
	Subject2 int  `json:"subject2"`
	Total    int     `json:"total"`
	Avg      float32 `json:"avg"`
	Rank     int     `json:"rank"`
}
