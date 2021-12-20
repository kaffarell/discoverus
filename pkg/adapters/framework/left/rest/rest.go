package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/kaffarell/discoverus/pkg/application/core/instance"
	"github.com/kaffarell/discoverus/pkg/application/core/service"
	"github.com/kaffarell/discoverus/pkg/ports"
)

type Adapter struct {
	api ports.APIPort
}

// FIXME: find better solution
type ServiceJson struct {
	Id             string `json:"id"`
	ServiceType    string `json:"serviceType"`
	Ip             string `json:"ip"`
	Port           int    `json:"port"`
	HealthCheckUrl string `json:"healthCheckUrl"`
}

// NewAdapter creates a new Adapter
func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (adapter Adapter) PostRegister(writer http.ResponseWriter, req *http.Request, parameter httprouter.Params) {
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
	serviceId := parameter.ByName("id")

	// Check if service is already existing
	_, error := adapter.api.GetService(serviceId)
	if error != nil {
		// Create Service
		result := service.NewService(t.Id, t.ServiceType, t.HealthCheckUrl)
		err := adapter.api.InsertService(result)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(writer, "Error creating new Service")
			return
		}
	}
	// Add instance
	// Get the current unix time and set it as the last heartbeat of the instance
	currentTime := time.Now().Unix()

	uuidWithHyphen := uuid.New()
	// Get uuid without hyphens
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	// Create new instance
	instance := instance.NewInstance(uuid, t.Ip, t.Port, currentTime)
	adapter.api.AddInstance(serviceId, instance)

	// Return the instance
	json, _ := json.Marshal(instance)

	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, string(json))
}

func (adapter Adapter) GetInstances(writer http.ResponseWriter, req *http.Request, parameter httprouter.Params) {
	// Get appId
	serviceId := parameter.ByName("id")
	instances_array, err := adapter.api.GetInstances(serviceId)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Requested service not found")
		return
	}

	json, _ := json.Marshal([]instance.Instance(instances_array))
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, string(json))
}
func (adapter Adapter) GetServices(writer http.ResponseWriter, req *http.Request, parameter httprouter.Params) {
	json, _ := json.Marshal(adapter.api.GetServices())
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, string(json))
}

func (a Adapter) PutRenew(writer http.ResponseWriter, req *http.Request, parameter httprouter.Params) {
	writer.WriteHeader(http.StatusOK)
}

func (a Adapter) GetHC(writer http.ResponseWriter, req *http.Request, parameter httprouter.Params) {
	writer.WriteHeader(http.StatusOK)
}

func (a Adapter) Run() {
	router := httprouter.New()
	router.GET("/hc", a.GetHC)
	router.GET("/apps", a.GetServices)
	router.GET("/apps/:id", a.GetInstances)
	router.POST("/apps/:id", a.PostRegister)
	router.PUT("/apps/:id/:instance", a.PutRenew)

	// Serve status website
	router.ServeFiles("/status/*filepath", http.Dir("website/"))

	err := http.ListenAndServe(":2000", router)
	if err != nil {
		fmt.Println(err)
	}
}
