package transaction

import "time"

func validateSearchTransactionsRequest(
	request SearchTransactionsRequest,
) error {
	if request.FromDate == "" || request.ToDate == "" {
		return nil
	}

	fromDate, err := time.Parse(
		"2006-01-02",
		request.FromDate,
	)
	if err != nil {
		return err
	}

	toDate, err := time.Parse(
		"2006-01-02",
		request.ToDate,
	)
	if err != nil {
		return err
	}

	if fromDate.After(toDate) {
		return ErrInvalidDateRange
	}

	return nil
}
