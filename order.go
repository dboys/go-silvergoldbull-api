package sgb

import (
	"encoding/json"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Address struct {
	Region    string `json:"region"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
	Postcode  string `json:"postcode"`
	Street    string `json:"street"`
	Company   string `json:"company"`
}

type Item struct {
	ID        string  `json:"id"`
	QTY       int     `json:"qty"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	TaxAmount float32 `json:"tax_amount"`
}

type Order struct {
	ID         string      `json:"id"`
	Currency   string      `json:"currency"`
	Link       *Link       `json:"links"`
	Status     string      `json:"status"`
	CreatedAt  CreatedDate `json:"created_at"`
	Total      float32     `json:"grand_total"`
	TaxAmount  float32     `json:"tax_amount"`
	ShipAmount float32     `json:"shipping_amount"`
	ShipMethod string      `json:"shipping_method"`
	Items      []*Item     `json:"items"`
	Shipping   *Address    `json:"shipping"`
	Billing    *Address    `json:"billing"`
}

type CreatedDate struct {
	time.Time
}

const _orderEntity = "orders"

func (cd *CreatedDate) UnmarshalJSON(input []byte) error {
	str := string(input)
	str = strings.Trim(str, `"`)
	newTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}

	cd.Time = newTime
	return nil
}

func (s *sgb) GetOrderList() ([]*Order, error) {
	var o []*Order

	req, err := s.request(getMethod, _orderEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &o); err != nil {
		return nil, err
	}

	return o, err
}

func (s *sgb) GetOrder(id int) (*Order, error) {
	var reqEntity = filepath.Join(_orderEntity, strconv.Itoa(id))
	var o *Order

	req, err := s.request(getMethod, reqEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &o); err != nil {
		return nil, err
	}

	return o, err
}
