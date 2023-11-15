package usecase

import (
	"github.com/kusipay/api-go-messenger/domain/port"
	"github.com/kusipay/api-go-messenger/domain/types"
)

type createschedule struct {
	logger              port.LoggerRepository
	schedulerRepository port.SchedulerRepository
}

func CreateScheduleUseCase(logger port.LoggerRepository, schedulerRepository port.SchedulerRepository) port.CreateScheduleService {
	return &createschedule{logger: logger, schedulerRepository: schedulerRepository}
}

func (c *createschedule) CreateSchedule(input types.CreateScheduleInput) error {
	return c.schedulerRepository.CreateSchedule(input)
}
