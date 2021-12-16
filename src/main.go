package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ServiceJson struct {
	Id             string `json:"id"`
	ServiceType    string `json:"serviceType"`
	Ip             string `json:"ip"`
	Port           int    `json:"port"`
	HealthCheckUrl string `json:"healthCheckUrl"`
}

func register(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//service := ps.ByName("id")
	/*
		JSON structure:
		{
			"id": "user"
			"serviceType": "service"
			"ip": "192_144.3.5"
			"port": 87
			"healthCheckUrl": "/hc"
		}
	*/
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var t ServiceJson
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Id)
	fmt.Println(t)

	w.WriteHeader(http.StatusOK)
}

func getService(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func renew(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func hc(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	router := httprouter.New()
	router.GET("/hc", hc)
	router.GET("/apps/:id", getService)
	router.POST("/apps/:id", register)
	router.PUT("/apps/:id/:instance", renew)

	http.ListenAndServe(":80", router)
}
