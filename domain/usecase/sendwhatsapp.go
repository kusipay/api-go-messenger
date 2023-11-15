package usecase

import (
	"github.com/kusipay/api-go-messenger/domain/port"
	"github.com/kusipay/api-go-messenger/domain/types"
)

type sendwhatsapp struct {
	logger             port.LoggerRepository
	whatsappRepository port.WhatsappRepository
}

func SendWhatsAppUseCase(loggerRepository port.LoggerRepository, whatsappRepository port.WhatsappRepository) port.SendWhatsappService {
	return &sendwhatsapp{logger: loggerRepository, whatsappRepository: whatsappRepository}
}

func (c *sendwhatsapp) SendWhatsapp(input types.SendWhatsappInput) error {
	return c.whatsappRepository.SendWhatsapp(input)
}
