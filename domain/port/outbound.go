package port

import "github.com/kusipay/api-go-messenger/domain/types"

type LoggerRepository interface {
	Verbose(tag string, message ...string)
	Debug(tag string, message ...string)
	Info(tag string, message ...string)
	Warn(tag string, message ...string)
	Error(tag string, message ...string)
}

type SchedulerRepository interface {
	CreateSchedule(input types.CreateScheduleInput) error
}
