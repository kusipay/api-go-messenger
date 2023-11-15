package types

type SendWhatsappInput struct {
	Phone        string                 `json:"phone"`
	TemplateName string                 `json:"templateName"`
	LanguageCode string                 `json:"languageCode"`
	Variables    map[string]interface{} `json:"variables"`
}

type SendWhatsappParams = SendWhatsappInput

type SendWhatsappEvent = SendWhatsappInput
