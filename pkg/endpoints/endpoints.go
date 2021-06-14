package endpoints

type Set struct {
	GetEndpoint           endpoint.Endpoint
	AddTransactionEndpoint   endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	StatmentEndpoint     endpoint.Endpoint
}

var logger log.Logger

func (s *Set) Status(ctx context.Context, transactionID string) (internal.Status, error) {
	resp, err := s.StatusEndpoint(ctx, StatusRequest{TransactionID: transactionID})
	if err != nil {
		return internal.Failed, err
	}
	stsResp := resp.(StatusResponse)
	if stsResp.Err != "" {
		return internal.Failed, errors.New(stsResp.Err)
	}
	return stsResp.Status, nil
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}