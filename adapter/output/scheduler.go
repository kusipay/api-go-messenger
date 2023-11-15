package output

import (
	"context"
	"encoding/json"
	"errors"
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

func (s *scdlr) CreateSchedule(params t.CreateScheduleParams) error {
	functionArn := os.Getenv("ENV_FUNCTION_ARN")
	if functionArn == "" {
		return errors.New("ENV_FUNCTION_ARN not found")
	}
	roleArn := os.Getenv("ENV_SCHEDULER_ROLE_ARN")
	if roleArn == "" {
		return errors.New("ENV_SCHEDULER_ROLE_ARN not found")
	}
	groupName := os.Getenv("ENV_SCHEDULE_GROUP_NAME")
	if groupName == "" {
		return errors.New("ENV_SCHEDULE_GROUP_NAME not found")
	}

	bts, err := json.Marshal(params.Payload)
	if err != nil {
		return err
	}

	startAt := time.Unix(params.StartAt, 0)
	endAt := time.Unix(params.EndAt, 0)

	cronExpression, err := getCronExpression(startAt, params.FrecuencyType)
	if err != nil {
		s.logger.Error("CreateSchedule |", err.Error())

		return err
	}

	_, err = s.client.CreateSchedule(s.ctx, &scheduler.CreateScheduleInput{
		Name:               aws.String(params.Id),
		Description:        aws.String(fmt.Sprintf("%s schedule for %s", params.FrecuencyType, params.Id)),
		ScheduleExpression: aws.String(cronExpression),
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

func getCronExpression(startAt time.Time, frecuency t.FrecuencyType) (string, error) {
	minutes := startAt.UTC().Minute()
	hours := startAt.UTC().Hour()
	dayOfTheWeek := int(startAt.UTC().Weekday())
	dayOfTheMonth := startAt.UTC().Day()

	switch frecuency {
	case t.FrecuencyTypeMonthly:
		return fmt.Sprintf("cron(%d %d %d * ? *)", minutes, hours, dayOfTheMonth), nil
	case t.FrecuencyTypeWeekly:
		return fmt.Sprintf("cron(%d %d ? * %d *)", minutes, hours, dayOfTheWeek), nil
	default:
		return "", fmt.Errorf("frecuencyType not supported: %s", frecuency)
	}
}
