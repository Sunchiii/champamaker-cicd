package responses

type Status struct {
	Status  int                    `json:"status"`
	Massage string                 `json:"massage"`
	Data    map[string]interface{} `json:"data"`
}
