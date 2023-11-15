package output

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/kusipay/api-go-messenger/domain/port"
	t "github.com/kusipay/api-go-messenger/domain/types"
)

type scdlr struct {
	ctx    context.Context
	logger port.LoggerRepository
	client *scheduler.Client
}

func NewSchedulerRepository(ctx context.Context, cfg aws.Config, logger port.LoggerRepository) port.SchedulerRepository {
	client := scheduler.NewFromConfig(cfg)

	return &scdlr{ctx: ctx, logger: logger, client: client}
}

func (s *scdlr) CreateSchedule(input t.CreateScheduleInput) error {
	functionArn := os.Getenv("ENV_FUNCTION_ARN")
	roleArn := os.Getenv("ENV_SCHEDULER_ROLE_ARN")
	groupName := os.Getenv("ENV_SCHEDULE_GROUP_NAME")

	bts, err := json.Marshal(input.Payload)
	if err != nil {
		return err
	}

	startAt := time.Unix(input.StartAt, 0)
	endAt := time.Unix(input.EndAt, 0)

	_, err = s.client.CreateSchedule(s.ctx, &scheduler.CreateScheduleInput{
		Name:               aws.String(input.Id),
		Description:        aws.String(fmt.Sprintf("%s schedule for %s", input.FrecuencyType, input.Id)),
		ScheduleExpression: aws.String(getCronExpression(startAt, input.FrecuencyType)),
		FlexibleTimeWindow: &types.FlexibleTimeWindow{
			Mode:                   types.FlexibleTimeWindowModeFlexible,
			MaximumWindowInMinutes: aws.Int32(60),
		},
		Target: &types.Target{
			Arn:     aws.String(functionArn),
			RoleArn: aws.String(roleArn),
			Input:   aws.String(string(bts)),
			RetryPolicy: &types.RetryPolicy{
				MaximumRetryAttempts: aws.Int32(2),
			},
		},
		ActionAfterCompletion: types.ActionAfterCompletionDelete,
		State:                 types.ScheduleStateEnabled,
		StartDate:             aws.Time(startAt),
		EndDate:               aws.Time(endAt),
		GroupName:             aws.String(groupName),
	})
	if err != nil {
		return err
	}

	return nil
}

func getCronExpression(startAt time.Time, frecuency string) string {
	minutes := startAt.UTC().Minute()
	hours := startAt.UTC().Hour()
	dayOfTheWeek := startAt.UTC().Weekday()
	dayOfTheMonth := startAt.UTC().Day()

	if frecuency == "monthly" {
		return fmt.Sprintf("cron(%d %d %d ? * *)", minutes, hours, dayOfTheMonth)
	}

	return fmt.Sprintf("cron(%d %d ? * %d *)", minutes, hours, dayOfTheWeek)
}
