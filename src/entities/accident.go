package entities

type Accident struct {
	TotalAccident int `json:"total_accident"`
	TotalDeath    int `json:"total_death"`
	TotalInvolved int `json:"total_involved"`
	TotalInjured  int `json:"total_injured"`
}
