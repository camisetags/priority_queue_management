package domain

type Job struct {
	ID     int    `json:"id"`
	Data   []int  `json:"data"`
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}
