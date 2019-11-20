package sgb

import "encoding/json"

type shipping struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *sgb) GetShippingList() ([]*shipping, error) {
	const reqEntity = "shipping/method"
	var shm []*shipping

	req, err := s.httpGetBytes(reqEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &shm); err != nil {
		return nil, err
	}

	return shm, err
}
