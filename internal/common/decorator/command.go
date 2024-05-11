package decorator

import "context"

func ApplyCommandDecorators[H any](handler CommandHandler[H]) CommandHandler[H] {
	return commandLoggingDecorator[H]{
		base: handler,
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}