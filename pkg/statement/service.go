package statement

import (
	"context"
	"github.com/fernandesleticia/go-microservice/internal"
)

type Service interface {
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Transaction, error)
	Status(ctx context.Context, TransactionID string) (internal.Status, error)
	Statement(ctx context.Context, TransactionID string) (int, error)
	AddTransaction(ctx context.Context, doc *internal.Transaction) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}