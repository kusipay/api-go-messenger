package types

type CreateScheduleInput struct {
	Id            string                 `json:"id"`
	StartAt       int64                  `json:"startAt"`
	EndAt         int64                  `json:"endAt"`
	FrecuencyType string                 `json:"frecuencyType"`
	Payload       map[string]interface{} `json:"payload"`
}
