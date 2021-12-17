package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kaffarell/discoverus/src/service"
)

type ServiceJson struct {
	Id             string `json:"id"`
	ServiceType    string `json:"serviceType"`
	Ip             string `json:"ip"`
	Port           int    `json:"port"`
	HealthCheckUrl string `json:"healthCheckUrl"`
}

func register(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
	// Parse JSON
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var t ServiceJson
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	// Get appId
	serviceId := ps.ByName("id")

	// Check if service is already existing
	_, error := service.GetService(serviceId)
	if error != nil {
		// Create Service
		result := service.NewService(t.Id, t.ServiceType, t.HealthCheckUrl)

		if result == false {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error creating new Service")
			return
		}
	}
	// Add instance
	instance := service.NewInstance(0, t.Ip, t.Port)
	service.AddInstance(serviceId, instance)

	w.WriteHeader(http.StatusOK)
}

func getInstances(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// Get appId
	serviceId := ps.ByName("id")
	instances_array, err := service.GetInstances(serviceId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Requested service not found")
	}

	json, _ := json.Marshal([]service.Instance(instances_array))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(json))
}

func getServices(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	json, _ := json.Marshal(service.GetServices())
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(json))
}

func renew(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func hc(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	service.Services = make(map[service.Service]service.InstanceArray)

	router := httprouter.New()
	router.GET("/hc", hc)
	router.GET("/apps", getServices)
	router.GET("/apps/:id", getInstances)
	router.POST("/apps/:id", register)
	router.PUT("/apps/:id/:instance", renew)

	http.ListenAndServe(":80", router)
}
