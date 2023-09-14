package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pizza-hub/internal/transport"
	"testing"
)

func TestGetMenu(t *testing.T) {
	req, err := http.NewRequest("GET", "/menus", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := NewHandler()

	handler.GetMenu(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	var resp map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
	}

	respData := resp["data"].([]interface{})

	if len(respData) != 2 {
		t.Errorf("Expected pizza count is %d but got %d", 2, len(respData))
	}
}

func TestAddChef(t *testing.T) {
	chefReq := &transport.ChefReq{
		Name: "Chef Arnold",
	}

	reqBody, err := json.Marshal(chefReq)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/chef", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := NewHandler()

	handler.AddChef(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	var response transport.ChefRes
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
	}

}

func TestAddOrder(t *testing.T) {
	orderReq := []transport.OrderReq{
		{
			Name: "Pizza-customer",
			Type: "Pizza Cheese",
		},
	}
	reqBody, err := json.Marshal(orderReq)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/orders", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := NewHandler()

	handler.AddOrder(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	var response transport.OrderRes
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
	}

}
