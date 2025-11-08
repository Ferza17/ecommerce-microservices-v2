package context

import "context"

const ContextKeyCausationID = "X-Causation-Id"

func SetCausationIdToContext(ctx context.Context, correlationId string) context.Context {
	return context.WithValue(ctx, ContextKeyCausationID, correlationId)
}

func GetCausationIdFromContext(ctx context.Context) string {
	correlationId, ok := ctx.Value(ContextKeyCausationID).(string)
	if !ok {
		return ""
	}
	return correlationId
}
