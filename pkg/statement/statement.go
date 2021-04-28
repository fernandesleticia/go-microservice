package statement

type statementService struct{}

func NewService() Service { return &statementService }

func (w *statementService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	doc := internal.Document{
		AccountDigit: "1",
		AccountNumber: "123",
		AccountType: "",
		BanckCode: "",
		PersonId: ""
	}
	return []internal.Document{doc}, nil
}
