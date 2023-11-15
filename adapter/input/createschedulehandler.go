package input

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/kusipay/api-go-messenger/adapter/output"
	"github.com/kusipay/api-go-messenger/domain/types"
	"github.com/kusipay/api-go-messenger/domain/usecase"
)

func CreateScheduleHandler(ctx context.Context, event types.CreateScheduleInput) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}

	logger := output.NewLoggerRepository()
	scheduler := output.NewSchedulerRepository(ctx, cfg, logger)

	err = usecase.CreateScheduleUseCase(logger, scheduler).CreateSchedule(event)
	if err != nil {
		return "", err
	}

	return "Hello, World!", nil
}
