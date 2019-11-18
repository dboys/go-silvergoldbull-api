package sgb

import (
	"encoding/json"
	"path/filepath"
	"strconv"
)

type Link struct {
	Self *struct {
		URL string `json:"href"`
	}
}

type Tier struct {
	QTY   int     `json:"qty"`
	Price float32 `json:"price"`
}

type Price struct {
	To           float32 `json:"to"`
	From         float32 `json:"from"`
	CashDiscount bool    `json:"cash_discount"`
	Currency     string  `json:"currency`
	Tiers        []*Tier `json:"tiers"`
}

type Tax struct {
	Name     string  `json:"name"`
	Percent  float32 `json:"percent"`
	Region   string  `json:"region"`
	priority int     `json:"priority"`
	TaxClass int     `json:"customer_tax_class_id"`
}

type Product struct {
	Name         string  `json:"name"`
	ID           string  `json:"id"`
	InStock      bool    `json:"in_stock"`
	Link         *Link   `json:"links"`
	Description  string  `json:"description"`
	DaysSinceNew int     `json:"days_since_new"`
	URL          string  `json:"url"`
	Country      string  `json:"country"`
	Material     string  `json:"material"`
	Weigth       int     `json:"metal_weight_troy_oz"`
	Image        string  `json:"image"`
	Mintage      string  `json:"mintage"`
	QTY          int     `json:"qty"`
	Purity       string  `json:"purity"`
	Shape        string  `json:"shape"`
	Reverse      string  `json:"reverse"`
	Obverse      string  `json:"obverse"`
	TotalWeight  float32 `json:"total_weight_troy_oz"`
	Manufacturer string  `json:"manufacturer"`
	Tender       string  `json:"legal_tender"`
	Price        *Price  `json:"prices"`
	Tax          []*Tax  `json:"tax"`
}

const _productEntity = "products"

func (s *sgb) GetProductList() ([]*Product, error) {
	var p []*Product

	req, err := s.httpGetBytes(_productEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &p); err != nil {
		return nil, err
	}

	return p, err
}

func (s *sgb) GetProduct(id int) (*Product, error) {
	var reqEntity = filepath.Join(_productEntity, strconv.Itoa(id))
	var p *Product

	req, err := s.httpGetBytes(reqEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &p); err != nil {
		return nil, err
	}

	return p, err
}
