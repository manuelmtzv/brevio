package i18n

import "context"

type contextKey string

const languageKey contextKey = "language"

func WithLanguage(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, languageKey, lang)
}

func LanguageFromContext(ctx context.Context) string {
	if lang, ok := ctx.Value(languageKey).(string); ok && lang != "" {
		return lang
	}
	return "en"
}
