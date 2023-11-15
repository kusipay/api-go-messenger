package sendwhatsapp

import (
	"context"

	"github.com/kusipay/api-go-messenger/adapter/output"
	"github.com/kusipay/api-go-messenger/domain/types"
	"github.com/kusipay/api-go-messenger/domain/usecase"
	"github.com/kusipay/api-go-messenger/util"
)

func fail(err error) (util.AnyResponse, error) {
	return util.AnyResponse{}, err
}

func success() (util.AnyResponse, error) {
	return util.AnyResponse{
		"message": "success",
	}, nil
}

func Handler(ctx context.Context, event types.SendWhatsappEvent) (util.AnyResponse, error) {
	logger := output.NewLoggerRepository()
	whatsappRepository := output.NewWhatsappRepository(logger)

	err := usecase.SendWhatsAppUseCase(logger, whatsappRepository).SendWhatsapp(event)
	if err != nil {
		return fail(err)
	}

	return success()
}
