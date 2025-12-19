package i18n

import "context"

type Localizer interface {
	Message(
		ctx context.Context,
		messageID string,
		defaultMsg string,
		data map[string]any,
	) string
}
