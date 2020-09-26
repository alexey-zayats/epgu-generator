package consumer

import "context"

// Consumer интерфейс запуска консьюмеров
type Consumer interface {
	Run(ctx context.Context, args []string)
}
