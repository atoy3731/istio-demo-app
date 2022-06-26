package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Status: "success",
		Code:   200,
	})
}

func internetStatus(w http.ResponseWriter, r *http.Request, vars Vars) {
	w.Header().Set("Content-Type", "application/json")

	var respStatus string
	var respCode int

	c := http.Client{Timeout: time.Duration(5) * time.Second}
	log.Print(fmt.Sprintf("INTERNET: Make request to: %s", vars.InternetUrl))
	resp, reqErr := c.Get(vars.InternetUrl)
	if reqErr != nil {
		respCode = 0
		respStatus = "error"
	} else {
		if resp.StatusCode == 200 {
			respCode = resp.StatusCode
			respStatus = "success"
		} else {
			respCode = resp.StatusCode
			respStatus = "error"
		}
	}

	err := json.NewEncoder(w).Encode(Response{
		Status: respStatus,
		Code:   respCode,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}

func intraStatus(w http.ResponseWriter, r *http.Request, vars Vars) {
	w.Header().Set("Content-Type", "application/json")

	var respStatus string
	var respCode int

	c := http.Client{Timeout: time.Duration(5) * time.Second}
	log.Print(fmt.Sprintf("INTRA-NAMESPACE: Make request to: %s", vars.IntraUrl))
	resp, reqErr := c.Get(vars.IntraUrl)
	if reqErr != nil {
		respCode = 0
		respStatus = "error"
	} else {
		if resp.StatusCode == 200 {
			respCode = resp.StatusCode
			respStatus = "success"
		} else {
			respCode = resp.StatusCode
			respStatus = "error"
		}
	}

	err := json.NewEncoder(w).Encode(Response{
		Status: respStatus,
		Code:   respCode,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}

func crossStatus(w http.ResponseWriter, r *http.Request, vars Vars) {
	w.Header().Set("Content-Type", "application/json")

	var respStatus string
	var respCode int

	c := http.Client{Timeout: time.Duration(5) * time.Second}
	log.Print(fmt.Sprintf("CROSS-NAMESPACE: Make request to: %s", vars.CrossUrl))
	resp, reqErr := c.Get(vars.CrossUrl)
	if reqErr != nil {
		respCode = 0
		respStatus = "error"
	} else {
		if resp.StatusCode == 200 {
			respCode = resp.StatusCode
			respStatus = "success"
		} else {
			respCode = resp.StatusCode
			respStatus = "error"
		}
	}

	err := json.NewEncoder(w).Encode(Response{
		Status: respStatus,
		Code:   respCode,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}
