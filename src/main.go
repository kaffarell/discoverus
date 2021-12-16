package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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
