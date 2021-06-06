package statement

type statementService struct{}

func NewService() Service { return &statementService }

func (w *statementService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Transaction, error) {
	transaction := internal.Transaction{
		AccountDigit: "1",
		AccountNumber: "123",
		AccountType: "",
		BanckCode: "",
		PersonId: "",
		Ammount: ""
	}
	return []internal.Transaction{transaction}, nil
}

func (w *statementService) Status(_ context.Context, accountID string) (internal.Status, error) {
	return internal.InProgress, nil
}

func (w *statementService) Statement(_ context.Context, accountID string) (int, error) {
	return http.StatusOK, nil
}

func (w *statementService) AddTransaction(_ context.Context, transaction *internal.Transaction) (string, error) {
	return amount, nil
}

func (w *statementService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking service health")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
