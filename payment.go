package sgb

import "encoding/json"

type payment struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	OrderMax     int      `json:"order_maximum"`
	ClearingDays int      `json:"clearing_days"`
	Currency     []string `json:"currency"`
	Country      []string `json:"country"`
}

func (s *sgb) GetPaymentList() ([]*payment, error) {
	const reqEntity = "payments/method"
	var pm []*payment

	req, err := s.httpGetBytes(reqEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &pm); err != nil {
		return nil, err
	}

	return pm, err
}
