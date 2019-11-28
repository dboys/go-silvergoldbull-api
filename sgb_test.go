package sgb

import (
	"os"
	"reflect"
	"strconv"
	"testing"
)

type testGetListMethod struct {
	fn   interface{}
	name string
}

type testGetMethod struct {
	fn          interface{}
	name        string
	inputID     int
	isErrorTest bool
}

func testsGetListMethods(s *sgb) []*testGetListMethod {
	return []*testGetListMethod{
		{s.GetCurrencyList, "GetCurrencyList"},
		{s.GetPaymentList, "GetPaymentList"},
		{s.GetShippingList, "GetShippingList"},
		{s.GetProductList, "GetProductList"},
		{s.GetOrderList, "GetOrderList"},
	}
}

func runGetListTests(t *testing.T, tests []*testGetListMethod, isErrorTest bool) {
	for _, test := range tests {
		result := reflect.ValueOf(test.fn).Call([]reflect.Value{})
		list, err := result[0].Interface(), result[1].Interface()

		if isErrorTest {
			if err == nil {
				t.Errorf("%s() didn't fail with wrong key: %s", test.name, err)
			}

			if reflect.ValueOf(list).Len() > 0 {
				t.Errorf("%s() returns non-empty result with wrong key", test.name)
			}
		} else {
			if err != nil {
				t.Errorf("%s() failed: %s", test.name, err)
			}

			if reflect.ValueOf(list).Len() == 0 {
				t.Errorf("%s() returns empty result", test.fn)
			}
		}
	}
}

func TestGetListMethods(t *testing.T) {
	apiKey := os.Getenv("SILVERGOLDBULL_API_KEY")

	if len(apiKey) > 0 {
		client := New(apiKey)
		tests := testsGetListMethods(client)

		runGetListTests(t, tests, false)
	}
}

func TestGetListMethodsWrongKey(t *testing.T) {
	client := New("wrong key")
	tests := testsGetListMethods(client)

	runGetListTests(t, tests, true)
}

func testsGetMethods(s *sgb) []*testGetMethod {
	var tests = []*testGetMethod{
		{s.GetProduct, "GetProduct", 0, true},
		{s.GetOrder, "GetOrder", 0, true},
	}

	orders, err := s.GetOrderList()
	if err == nil && len(orders) > 0 {
		order := orders[0]
		id, _ := strconv.Atoi(order.ID)
		tests = append(tests, &testGetMethod{s.GetOrder, "GetOrder", id, false})
	}

	products, err := s.GetProductList()
	if err == nil && len(products) > 0 {
		product := products[0]
		id, _ := strconv.Atoi(product.ID)
		tests = append(tests, &testGetMethod{s.GetProduct, "GetProduct", id, false})
	}

	return tests
}

func runGetTests(t *testing.T, tests []*testGetMethod) {
	for _, test := range tests {
		result := reflect.ValueOf(test.fn).Call([]reflect.Value{reflect.ValueOf(test.inputID)})
		_, err := result[0].Interface(), result[1].Interface()

		if test.isErrorTest {
			if err == nil {
				t.Errorf("%s(%d) failed: %s", test.name, test.inputID, err)
			}
		} else {
			if err != nil {
				t.Errorf("%s(%d) failed: %s", test.name, test.inputID, err)
			}
		}
	}
}

func TestGetMethods(t *testing.T) {
	apiKey := os.Getenv("SILVERGOLDBULL_API_KEY")

	if len(apiKey) > 0 {
		client := New(apiKey)
		tests := testsGetMethods(client)

		runGetTests(t, tests)
	}
}

func TestCreateMethods(t *testing.T) {
	apiKey := os.Getenv("SILVERGOLDBULL_API_KEY")

	if len(apiKey) > 0 {
		client := New(apiKey)

		addr := &Address{
			Email: "sales@silvergoldbull.com",
			Postcode: "T2P 	5C5",
			Region:    "AB",
			City:      "Calgary",
			FirstName: "John",
			LastName:  "Smith",
			Country:   "CA",
			Street:    "888 - 3 ST SW, 10 FLOOR - WEST TOWER",
			Phone:     "+1 (403) 668 8648",
		}
		items := []*Item{
			{
				ID:       "2706",
				QTY:      1,
				BidPrice: 468.37,
			},
			{
				ID:       "2580",
				QTY:      1,
				BidPrice: 2,
			},
		}
		quote := &Quote{
			Currency:      "USD",
			PaymentMethod: "paypall",
			ShipMethod:    "1YR_STORAGE",
			Declaration:   "TEST",
			Items:         items,
			Shipping:      addr,
			Billing:       addr,
		}

		_, err := client.Quote(quote)
		if err == nil {
			t.Errorf("Quote(%v) return result with wrong data", quote)
		}

		order := &Order{
			Currency:      "USD",
			PaymentMethod: "paypall",
			ShipMethod:    "1YR_STORAGE",
			Declaration:   "TEST",
			Items:         items,
			Shipping:      addr,
			Billing:       addr,
		}

		_, err = client.CreateOrder(order)
		if err == nil {
			t.Errorf("CreateOrder(%v) return result with wrong data", order)
		}
	}
}
