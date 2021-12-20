package db

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"
	// FIXME: should not be here
	// Maybe use interface{} as type and use Services in application/api
	"github.com/kaffarell/discoverus/pkg/application/core/instance"
	"github.com/kaffarell/discoverus/pkg/application/core/service"
)

// Adapter implements the DbPort interface
type Adapter struct {
	ctx           context.Context
	redisService  *redis.Client
	redisInstance *redis.Client
	redisRegistry *redis.Client
}

// NewAdapter creates a new Adapter
func NewAdapter() (*Adapter, error) {
	ctx := context.Background()
	// Create new clients for the service db, instance db and registry db
	redisService := redis.NewClient(&redis.Options{
		Addr:     "redis-services:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisInstance := redis.NewClient(&redis.Options{
		Addr:     "redis-instances:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisRegistry := redis.NewClient(&redis.Options{
		Addr:     "redis-registry:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test if all databases are up
	dbStatus := false
	dbStatus = pingDatabase(ctx, *redisService)
	dbStatus = pingDatabase(ctx, *redisInstance)
	dbStatus = pingDatabase(ctx, *redisRegistry)

	if dbStatus == false {
		return nil, errors.New("One or more databases are not up!")
	} else {
		return &Adapter{
			ctx:           ctx,
			redisService:  redisService,
			redisInstance: redisInstance,
			redisRegistry: redisRegistry}, nil

	}

}

func pingDatabase(ctx context.Context, rdb redis.Client) bool {
	_, err := rdb.Ping(ctx).Result()
	if err == nil {
		return true
	} else {
		return false
	}
}

func (a Adapter) AddService(service service.Service) error {
	// Add service with serviceId to services db
	jsonString, _ := json.Marshal(service)
	err := a.redisService.Set(a.ctx, service.Id, string(jsonString), 0).Err()

	// Create empty registry db entry
	// Create the service key with no instances
	// FIXME: better error handling (although we will most likely never get an error here)
	err = a.redisRegistry.Set(a.ctx, service.Id, "[]", 0).Err()
	// Error writing to redis storage or key (service) already existing
	return err
}

func (a Adapter) AddInstance(serviceId string, instance instance.Instance) error {
	// TODO: Check if serviceId exists (I don't really know if we really have to do this, because the
	// service struct holds actually pretty meaningless data)

	// Add new instance to instance db
	jsonString, _ := json.Marshal(instance)
	err := a.redisInstance.Set(a.ctx, instance.Id, jsonString, 0).Err()

	// Add new instance to corresponding registry
	// Get current instances of this service
	val, err := a.redisRegistry.Get(a.ctx, serviceId).Result()

	// Convert string to array with uuids
	var instances []string
	err = json.Unmarshal([]byte(val), &instances)

	// Add instance id to array
	instances = append(instances, instance.Id)

	// Convert array to json again
	jsonArrayString, _ := json.Marshal(instances)
	// Set new instances array to serviceid again
	err = a.redisRegistry.Set(a.ctx, serviceId, jsonArrayString, 0).Err()
	return err

}

func (a Adapter) RemoveInstance(serviceId string, instanceId string) error {
	// Remove instance from redis-registry
	err := a.redisInstance.Del(a.ctx, instanceId).Err()

	// Remove instance from redis-instances
	// Get current instances of this service
	val, err := a.redisRegistry.Get(a.ctx, serviceId).Result()

	// Convert string to array with uuids
	var instances []string
	err = json.Unmarshal([]byte(val), &instances)

	// Remove instanceId from array
	instances = removeFromArray(instances, instanceId)

	// Convert array to json again
	jsonArrayString, _ := json.Marshal(instances)
	// Set new instances array to serviceid again
	err = a.redisRegistry.Set(a.ctx, serviceId, jsonArrayString, 0).Err()

	return err
}

func (a Adapter) GetInstances(serviceId string) ([]instance.Instance, error) {
	instancesStringJson, err := a.redisRegistry.Get(a.ctx, serviceId).Result()

	// Convert string to array with uuids
	var instancesStrings []string
	json.Unmarshal([]byte(instancesStringJson), &instancesStrings)

	// For each instanceId get the instance object in redis-instances
	var instances []instance.Instance
	instances = make([]instance.Instance, 0)

	for _, s := range instancesStrings {
		val, err := a.redisInstance.Get(a.ctx, s).Result()
		if err == nil {
			var newInstance instance.Instance
			json.Unmarshal([]byte(val), &newInstance)
			instances = append(instances, newInstance)
		}
	}

	if err != nil {
		return nil, err
	}
	return instances, nil

}

func (a Adapter) GetRegistry() ([]string, error) {
	values, err := a.redisRegistry.Keys(a.ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (a Adapter) GetService(serviceId string) (service.Service, error) {
	val, err := a.redisService.Get(a.ctx, serviceId).Result()
	var service service.Service
	json.Unmarshal([]byte(val), &service)
	return service, err
}

func removeFromArray(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
