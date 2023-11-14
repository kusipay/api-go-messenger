package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/kusipay/api-go-messenger/pkg/util"
	"github.com/mefellows/vesper"
)

// EventLog logs event and response
func EventLog() vesper.Middleware {
	return func(next vesper.LambdaFunc) vesper.LambdaFunc {
		return func(ctx context.Context, event any) (any, error) {
			logEnvironments()
			logAny("Event", event)

			response, err := next(ctx, event)
			if err != nil {
				logError(err)
			} else {
				logAny("Response", response)
			}

			return response, err
		}
	}
}

func logEnvironments() {
	environments := os.Environ()

	envs := strings.Join(environments, "\r")

	util.Log("Environment |", envs)
}

func logAny(tag string, event any) {
	bytes, err := json.MarshalIndent(event, "", "  ")

	var text string
	if err != nil {
		text = fmt.Sprintf("%+v", event)
	} else {
		text = string(bytes)
	}

	util.Log(tag+" |", text)
}

func logError(err error) {
	if err != nil {
		util.Log("Error |", err.Error())
	}
}
