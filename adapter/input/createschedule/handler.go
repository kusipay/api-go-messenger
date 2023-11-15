package createschedule

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
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

func Handler(ctx context.Context, event types.CreateScheduleEvent) (util.AnyResponse, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fail(err)
	}

	logger := output.NewLoggerRepository()
	scheduler := output.NewSchedulerRepository(ctx, cfg, logger)

	err = usecase.CreateScheduleUseCase(logger, scheduler).CreateSchedule(event)
	if err != nil {
		return fail(err)
	}

	return success()
}
