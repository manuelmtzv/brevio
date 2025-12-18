package api

import (
	"context"

	appi18n "github.com/manuelmtzv/brevio/internal/i18n"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Localizer interface {
	Message(ctx context.Context, messageID string, defaultMsg string, data map[string]any) string
}

type localizer struct {
	bundle *i18n.Bundle
}

func NewLocalizer(bundle *i18n.Bundle) Localizer {
	return &localizer{bundle: bundle}
}

func (l *localizer) Message(
	ctx context.Context,
	messageID string,
	defaultMsg string,
	data map[string]any,
) string {
	lang := appi18n.LanguageFromContext(ctx)

	loc := i18n.NewLocalizer(l.bundle, lang)

	msg, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
		DefaultMessage: &i18n.Message{
			ID:    messageID,
			Other: defaultMsg,
		},
		TemplateData: data,
	})

	if err != nil {
		return defaultMsg
	}

	return msg
}
