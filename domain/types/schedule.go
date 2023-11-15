package types

type FrecuencyType string

const (
	FrecuencyTypeMonthly FrecuencyType = "monthly"
	FrecuencyTypeWeekly  FrecuencyType = "weekly"
)

type CreateScheduleInput struct {
	Id            string            `json:"id"`
	StartAt       int64             `json:"startAt"`
	EndAt         int64             `json:"endAt"`
	FrecuencyType FrecuencyType     `json:"frecuencyType"`
	Payload       SendWhatsappEvent `json:"payload"`
}

type CreateScheduleParams = CreateScheduleInput

type CreateScheduleEvent = CreateScheduleInput
