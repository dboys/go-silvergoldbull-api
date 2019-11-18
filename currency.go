package sgb

import "encoding/json"

type Currency struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (s *sgb) GetCurrencyList() ([]*Currency, error) {
	const reqEntity = "currencies"
	var currList []*Currency

	req, err := s.httpGetBytes(reqEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &currList); err != nil {
		return nil, err
	}

	return currList, err
}
