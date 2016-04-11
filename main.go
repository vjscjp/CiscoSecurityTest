package main

import (
	
	"fmt"
	"net/http"
	"crypto/tls"
	"encoding/json"
	"os"
	"strconv"
	marathon "github.com/gambol99/go-marathon"
)


type  App struct{
	
	App marathon.Application `json:"app"`
}



func main() {
	
	host := "shipped-tx3-worker-005"
	port := 31866
	
	argCount := len(os.Args[1:])
    //fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)
	if (argCount >=2){
		host = os.Args[1]
		port,_ = strconv.Atoi(os.Args[2])
	}else{
		os.Exit(0)
	}
	
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
	
	req, err := http.NewRequest("GET", "https://shipped-tx3-control-01.tx3.shipped-cisco.com/marathon/v2/tasks", nil)
	req.SetBasicAuth("synthetic-mon","VpYdy5abudqkk3Ts")
	resp, err := client.Do(req)
	if err != nil {
	    fmt.Printf("Error : %s", err)
	}
	//fmt.Println(resp)
	decoder := json.NewDecoder(resp.Body)
	var test marathon.Tasks
	decoder.Decode(&test)
	for _,t := range test.Tasks {
	
		if t.Host == host{
			//fmt.Println(t.AppID)
			for _,p := range t.Ports{
				if port == p{
					req, err = http.NewRequest("GET", "https://shipped-tx3-control-01.tx3.shipped-cisco.com/marathon/v2/apps"+t.AppID, nil)
					req.SetBasicAuth("synthetic-mon","VpYdy5abudqkk3Ts")
					//fmt.Println(req)
					resp, err = client.Do(req)
					if err != nil {
					    fmt.Printf("Error : %s", err)
					}
					decoder1 := json.NewDecoder(resp.Body)
					var applicant App
					decoder1.Decode(&applicant)
					fmt.Println("App Id : "+applicant.App.ID)
					//fmt.Println(applicant.App.Labels)
					for k,v :=range *applicant.App.Labels{
						fmt.Println(k+" : ",v)
					}
				}
			}
		}
	}
}