package port

import "github.com/kusipay/api-go-messenger/domain/types"

type CreateScheduleService interface {
	CreateSchedule(input types.CreateScheduleInput) error
}

type SendWhatsappService interface {
	SendWhatsapp(input types.SendWhatsappInput) error
}
