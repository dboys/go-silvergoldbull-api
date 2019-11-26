package sgb

import (
	"encoding/json"
	"path/filepath"
	"strconv"
)

type link struct {
	Self *struct {
		URL string `json:"href"`
	}
}

type tier struct {
	QTY   int     `json:"qty"`
	Price float32 `json:"price"`
}

type price struct {
	To           float32 `json:"to"`
	From         float32 `json:"from"`
	CashDiscount bool    `json:"cash_discount"`
	Currency     string  `json:"currency"`
	Tiers        []*tier `json:"tiers"`
}

type tax struct {
	Name     string  `json:"name"`
	Percent  float32 `json:"percent"`
	Region   string  `json:"region"`
	Priority int     `json:"priority"`
	TaxClass int     `json:"customer_tax_class_id"`
}

type product struct {
	Name         string  `json:"name"`
	ID           string  `json:"id"`
	InStock      bool    `json:"in_stock"`
	Link         *link   `json:"links"`
	Description  string  `json:"description"`
	DaysSinceNew int     `json:"days_since_new"`
	URL          string  `json:"url"`
	Country      string  `json:"country"`
	Material     string  `json:"material"`
	Weigth       int     `json:"metal_weight_troy_oz"`
	Image        string  `json:"image"`
	Mintage      float32 `json:"mintage"`
	QTY          int     `json:"qty"`
	Purity       string  `json:"purity"`
	Shape        string  `json:"shape"`
	Reverse      string  `json:"reverse"`
	Obverse      string  `json:"obverse"`
	TotalWeight  float32 `json:"total_weight_troy_oz"`
	Manufacturer string  `json:"manufacturer"`
	Tender       string  `json:"legal_tender"`
	Price        *price  `json:"prices"`
	Tax          []*tax  `json:"tax"`
}

const _productEntity = "products"

func (s *sgb) GetProductList() ([]*product, error) {
	var p []*product

	req, err := s.httpGetBytes(_productEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &p); err != nil {
		return nil, err
	}

	return p, err
}

func (s *sgb) GetProduct(id int) (*product, error) {
	var reqEntity = filepath.Join(_productEntity, strconv.Itoa(id))
	var p *product

	req, err := s.httpGetBytes(reqEntity)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &p); err != nil {
		return nil, err
	}

	return p, err
}
