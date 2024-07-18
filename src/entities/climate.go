package entities

type Climate struct {
	TotalAccident int `json:"total_accident"`
	TotalDeath    int `json:"total_death"`
	TotalInvolved int `json:"total_involved"`
}
