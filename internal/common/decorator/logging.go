package decorator

import "context"

type commandLoggingDecorator[C any] struct {
	base CommandHandler[C]
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	return d.base.Handle(ctx, cmd)
}