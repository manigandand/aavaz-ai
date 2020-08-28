package api

import (
	"aavaz/errors"
	"aavaz/respond"
	stre "aavaz/store"
	"aavaz/store/adapter"
	"aavaz/types"

	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// store connection copy
var store adapter.Store

// ServiceInfo stores basic service information
type ServiceInfo struct {
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Uptime  time.Time `json:"uptime"`
	Epoch   int64     `json:"epoch"`
}

// ServiceName holds the service which connected to
var ServiceName = ""
var serviceInfo *ServiceInfo

// InitAPI sets the service name
func InitAPI(name, version string) {
	store = stre.Store

	ServiceName = name
	serviceInfo = &ServiceInfo{
		Name:    name,
		Version: version,
		Uptime:  time.Now(),
		Epoch:   time.Now().Unix(),
	}
}

// API Handler's ---------------------------------------------------------------

// Handler custom api handler help us to handle all the errors in one place
type Handler func(w http.ResponseWriter, r *http.Request) *errors.AppError

func (f Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)
	if err != nil {
		// APP Level Error
		// TODO: handle 5XX, notify developers. Configurable
		log.Errorf("ServiceName: %s, StatusCode: %d, Error: %s\n DEBUG: %s\n",
			ServiceName, err.Status, err.Error(), err.Debug)
		respond.Fail(w, err)
	}
}

// Basic Handler func ----------------------------------------------------------

// IndexHandeler common index handler for all the service
func IndexHandeler(w http.ResponseWriter, r *http.Request) {
	respond.OK(w, types.JSON{
		"name":    serviceInfo.Name,
		"version": serviceInfo.Version,
	})
}

// HealthHandeler return basic service info
func HealthHandeler(w http.ResponseWriter, r *http.Request) {
	respond.OK(w, serviceInfo)
}
