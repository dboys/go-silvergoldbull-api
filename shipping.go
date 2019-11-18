package sgb

import "encoding/json"

type Shipping struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *sgb) GetShippingList() ([]*Shipping, error) {
	const reqEntity = "shipping/method"
	var shm []*Shipping

	req, err := s.httpGetBytes(reqEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &shm); err != nil {
		return nil, err
	}

	return shm, err
}
