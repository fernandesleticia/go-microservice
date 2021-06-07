package endpoints

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

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}