package statement

import (
	"context"
	"github.com/fernandesleticia/go-microservice/internal"
)

type Service interface {
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
	Status(ctx context.Context, id string) (internal.Status, error)
	Statement(ctx context.Context, id string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}