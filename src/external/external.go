package external

import "context"

type requestIDKey struct{}

func GetRequestID(ctx context.Context) (int, bool) {
	// requestIDKey型をKeyとすることで、パッケージ内のみ有効なKeyとできる
	// requestIDKey型はプライベートなので他パッケージで参照できない為、contextのValueの
	// KeyがrequestIDKey型の操作ができるのがこのパッケージのみになる
	r, ok := ctx.Value(requestIDKey{}).(int)
	if ok {
		return r, true
	}
	return 0, false
}

func WithRequestID(ctx context.Context, reqID int) context.Context {
	return context.WithValue(ctx, requestIDKey{}, reqID)
}
