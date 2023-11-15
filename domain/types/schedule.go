package types

type FrecuencyType string

const (
	FrecuencyTypeMonthly FrecuencyType = "monthly"
	FrecuencyTypeWeekly  FrecuencyType = "weekly"
)

type CreateScheduleInput struct {
	Id            string                 `json:"id"`
	StartAt       int64                  `json:"startAt"`
	EndAt         int64                  `json:"endAt"`
	FrecuencyType FrecuencyType          `json:"frecuencyType"`
	Payload       map[string]interface{} `json:"payload"`
}

type CreateScheduleParams = CreateScheduleInput

type CreateScheduleEvent = CreateScheduleInput
