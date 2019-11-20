package sgb

import (
	"bytes"
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

type item struct {
	ID        string  `json:"id"`
	QTY       int     `json:"qty"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	TaxAmount float32 `json:"tax_amount"`
}

type order struct {
	ID         string      `json:"id"`
	Currency   string      `json:"currency"`
	Link       *link       `json:"links"`
	Status     string      `json:"status"`
	CreatedAt  createdDate `json:"created_at"`
	Total      float32     `json:"grand_total"`
	TaxAmount  float32     `json:"tax_amount"`
	ShipAmount float32     `json:"shipping_amount"`
	ShipMethod string      `json:"shipping_method"`
	Items      []*item     `json:"items"`
	Shipping   *Address    `json:"shipping"`
	Billing    *Address    `json:"billing"`
}

type Item struct {
	ID       string  `json:"id"`
	QTY      int     `json:"qty"`
	BidPrice float32 `json:"bid_price"`
}

type Quote struct {
	Currency      string   `json:"currency"`
	PaymentMethod string   `json:"payment_method"`
	ShipMethod    string   `json:"shipping_method"`
	Declaration   string   `json:"declaration"`
	Items         []*Item  `json:"items"`
	Shipping      *Address `json:"shipping"`
	Billing       *Address `json:"billing"`
}

type Order struct {
	Currency      string   `json:"currency"`
	PaymentMethod string   `json:"payment_method"`
	ShipMethod    string   `json:"shipping_method"`
	Declaration   string   `json:"declaration"`
	Items         []*Item  `json:"items"`
	Shipping      *Address `json:"shipping"`
	Billing       *Address `json:"billing"`
}

type createdDate struct {
	time.Time
}

const _orderEntity = "orders"

func (cd *createdDate) UnmarshalJSON(input []byte) error {
	str := string(input)
	str = strings.Trim(str, `"`)
	newTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}

	cd.Time = newTime
	return nil
}

func (s *sgb) GetOrderList() ([]*order, error) {
	var o []*order

	req, err := s.httpGetBytes(_orderEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &o); err != nil {
		return nil, err
	}

	return o, err
}

func (s *sgb) GetOrder(id int) (*order, error) {
	var reqEntity = filepath.Join(_orderEntity, strconv.Itoa(id))
	var o *order

	req, err := s.httpGetBytes(reqEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &o); err != nil {
		return nil, err
	}

	return o, err
}

func (s *sgb) Quote(q *Quote) ([]byte, error) {
	var reqEntity = filepath.Join(_orderEntity, "quote")

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(q)

	req, err := s.httpPostBytes(reqEntity, buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s *sgb) CreateOrder(o *Order) ([]byte, error) {
	var reqEntity = filepath.Join(_orderEntity, "create")

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(o)

	req, err := s.httpPostBytes(reqEntity, buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}
