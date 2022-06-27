package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Internet ResponseComponent `json:"internet"`
	Intra    ResponseComponent `json:"intra"`
	Cross    ResponseComponent `json:"cross"`
}

type ResponseComponent struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func printHeaders(r *http.Request) {
	if reqHeadersBytes, err := json.Marshal(r.Header); err != nil {
		log.Println("Could not Marshal Req Headers")
	} else {
		log.Println(string(reqHeadersBytes))
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(ResponseComponent{
		Status: "success",
		Code:   200,
	})
}

func statusAll(w http.ResponseWriter, r *http.Request, vars Vars) {
	if vars.Debug == "true" {
		printHeaders(r)
	}
	
	c := http.Client{Timeout: time.Duration(5) * time.Second}

	// Internet Test
	internetResp := ResponseComponent{}
	log.Print(fmt.Sprintf("INTERNET: Make request to: %s", vars.InternetUrl))
	resp, reqErr := c.Get(vars.InternetUrl)
	if reqErr != nil {
		internetResp.Code = 0
		internetResp.Status = "error"
		log.Println(fmt.Sprintf("INTERNET ERROR: %s", reqErr))
	} else {
		if resp.StatusCode == 200 {
			internetResp.Code = resp.StatusCode
			internetResp.Status = "success"
		} else {
			internetResp.Code = resp.StatusCode
			internetResp.Status = "error"
			log.Println(fmt.Sprintf("INTERNET ERROR: %d", resp.StatusCode))

		}
	}

	// Intra Test
	intraResp := ResponseComponent{}
	log.Print(fmt.Sprintf("Intra: Make request to: %s", vars.IntraUrl))
	if vars.AuthToken == "" {
		resp, reqErr = c.Get(vars.IntraUrl)
	} else {
		req, err := http.NewRequest("GET", vars.IntraUrl, nil)
		if err != nil {
			log.Printf("ERROR: %s", err)
		}

		req.Header.Set("X-Auth-Token", vars.AuthToken)
		resp, reqErr = c.Do(req)
	}

	if reqErr != nil {
		intraResp.Code = 0
		intraResp.Status = "error"
		log.Println(fmt.Sprintf("INTRA ERROR: %s", reqErr))
	} else {
		if resp.StatusCode == 200 {
			intraResp.Code = resp.StatusCode
			intraResp.Status = "success"
		} else {
			intraResp.Code = resp.StatusCode
			intraResp.Status = "error"
			log.Println(fmt.Sprintf("INTRA ERROR: %d", resp.StatusCode))

		}
	}

	// Cross Test
	crossResp := ResponseComponent{}
	log.Print(fmt.Sprintf("CROSS: Make request to: %s", vars.CrossUrl))
	if vars.AuthToken == "" {
		resp, reqErr = c.Get(vars.CrossUrl)
	} else {
		log.Println("CROSS: Using authtoken!")
		req, err := http.NewRequest("GET", vars.CrossUrl, nil)
		if err != nil {
			log.Printf("ERROR: %s", err)
		}

		req.Header.Set("X-Auth-Token", vars.AuthToken)
		resp, reqErr = c.Do(req)
	}
	if reqErr != nil {
		crossResp.Code = 0
		crossResp.Status = "error"
		log.Println(fmt.Sprintf("CROSS ERROR: %s", reqErr))
	} else {
		if resp.StatusCode == 200 {
			crossResp.Code = resp.StatusCode
			crossResp.Status = "success"
		} else {
			crossResp.Code = resp.StatusCode
			crossResp.Status = "error"
			log.Println(fmt.Sprintf("CROSS ERROR: %d", resp.StatusCode))
		}
	}

	response := Response{
		Internet: internetResp,
		Intra:    intraResp,
		Cross:    crossResp,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}
