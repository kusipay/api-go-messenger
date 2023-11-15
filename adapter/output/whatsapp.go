package output

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/kusipay/api-go-messenger/domain/port"
	"github.com/kusipay/api-go-messenger/domain/types"
)

type wtsapp struct {
	logger port.LoggerRepository
}

func NewWhatsappRepository(logger port.LoggerRepository) port.WhatsappRepository {
	return &wtsapp{logger: logger}
}

func (w *wtsapp) SendWhatsapp(params types.SendWhatsappParams) error {
	w.logger.Debug("Whatsapp", "Send Whatsapp")

	token := "EAAMfZCGrCvQcBO9LzvN91VIOYkMD3nNzcVzCwrJznqTUwZAnKDJoMsGuD9aaIMsPOFyiSdZCA7geqgT5I5OwtZCJOY89ha8StI54TPAsxCBXZCG6hoZBErUgBqIVks0IwBcnZAWDEZAyRBWm94U9kgks6mMCxTIStghgQRZC3ifo4jTJ5bTmTCDl7BJe1B9If"

	url := "https://graph.facebook.com/v18.0/128041223734612/messages"

	rawBody := map[string]any{
		"messaging_product": "whatsapp",
		"to":                params.Phone,
		"type":              "template",
		"template": map[string]any{
			"name": params.TemplateName,
			"language": map[string]any{
				"code": params.LanguageCode,
			},
			"components": []any{
				map[string]any{
					"type": "body",
					"parameters": []any{
						map[string]any{"type": "text", "text": params.Variables["name"]},
					},
				},
			},
		},
	}

	body, err := json.Marshal(rawBody)
	if err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+token)

	response, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	bts, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	w.logger.Warn("SendWhatsapp |", string(bts))

	return errors.New("status code is not 200" + response.Status)
}
