package sgb

import "encoding/json"

type currency struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (s *sgb) GetCurrencyList() ([]*currency, error) {
	const reqEntity = "currencies"
	var currList []*currency

	req, err := s.httpGetBytes(reqEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &currList); err != nil {
		return nil, err
	}

	return currList, err
}
