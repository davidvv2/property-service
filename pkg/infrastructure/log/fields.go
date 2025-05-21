package log

import "go.uber.org/zap"

type Fields map[string]interface{}

// toZapFields converts a Fields map to a slice of zap.Field.
func toZapFields(fields Fields) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}
	return zapFields
}
