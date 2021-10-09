package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/NazeemNato/go-bank/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (s *CustomerHandlers) GetAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers, _ := s.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)
	}
}
