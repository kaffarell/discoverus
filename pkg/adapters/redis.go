package adapters

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/kaffarell/discoverus/pkg/types"
)

// Implements Storage interface
type RedisStorage struct {
	ctx                 context.Context
	redisServiceClient  *redis.Client
	redisInstanceClient *redis.Client
	redisRegistryClient *redis.Client
}

func (r *RedisStorage) New() {
	r.ctx = context.Background()
	// Create new clients for the service db, instance db and registry db
	r.redisServiceClient = redis.NewClient(&redis.Options{
		Addr:     "redis-services:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	r.redisInstanceClient = redis.NewClient(&redis.Options{
		Addr:     "redis-instances:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	r.redisRegistryClient = redis.NewClient(&redis.Options{
		Addr:     "redis-registry:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (r RedisStorage) AddService(service types.Service) error {
	// Add service with serviceId to services db
	jsonString, _ := json.Marshal(service)
	err := r.redisServiceClient.Set(r.ctx, service.Id, string(jsonString), 0).Err()

	// Create empty registry db entry
	// Create the service key with no instances
	// FIXME: better error handling (although we will most likely never get an error here)
	err = r.redisRegistryClient.Set(r.ctx, service.Id, "[]", 0).Err()
	// Error writing to redis storage or key (service) already existing
	return err
}

func (r RedisStorage) AddInstance(serviceId string, instance types.Instance) error {
	// TODO: Check if serviceId exists (I don't really know if we really have to do this, because the
	// service struct holds actually pretty meaningless data)

	// Add new instance to instance db
	jsonString, _ := json.Marshal(instance)
	err := r.redisInstanceClient.Set(r.ctx, instance.Id, jsonString, 0).Err()

	// Add new instance to corresponding registry
	// Get current instances of this service
	val, err := r.redisRegistryClient.Get(r.ctx, serviceId).Result()

	// Convert string to array with uuids
	var instances []string
	err = json.Unmarshal([]byte(val), &instances)

	// Add instance id to array
	instances = append(instances, instance.Id)

	// Convert array to json again
	jsonArrayString, _ := json.Marshal(instances)
	// Set new instances array to serviceid again
	err = r.redisRegistryClient.Set(r.ctx, serviceId, jsonArrayString, 0).Err()
	return err

}

func (r RedisStorage) RemoveInstance(serviceId string, instanceId string) error {
	// Remove instance from redis-registry
	err := r.redisInstanceClient.Del(r.ctx, instanceId).Err()

	// Remove instance from redis-instances
	// Get current instances of this service
	val, err := r.redisRegistryClient.Get(r.ctx, serviceId).Result()

	// Convert string to array with uuids
	var instances []string
	err = json.Unmarshal([]byte(val), &instances)

	// Remove instanceId from array
	instances = removeFromArray(instances, instanceId)

	// Convert array to json again
	jsonArrayString, _ := json.Marshal(instances)
	// Set new instances array to serviceid again
	err = r.redisRegistryClient.Set(r.ctx, serviceId, jsonArrayString, 0).Err()

	return err
}

func (r RedisStorage) GetInstances(serviceId string) ([]types.Instance, error) {
	instancesStringJson, err := r.redisRegistryClient.Get(r.ctx, serviceId).Result()

	// Convert string to array with uuids
	var instancesStrings []string
	json.Unmarshal([]byte(instancesStringJson), &instancesStrings)

	// For each instanceId get the instance object in redis-instances
	var instances []types.Instance
	instances = make([]types.Instance, 0)

	for _, s := range instancesStrings {
		val, err := r.redisInstanceClient.Get(r.ctx, s).Result()
		if err == nil {
			var newInstance types.Instance
			json.Unmarshal([]byte(val), &newInstance)
			instances = append(instances, newInstance)
		}
	}

	if err != nil {
		return nil, err
	}
	return instances, nil

}

func (r RedisStorage) GetRegistry() ([]string, error) {
	values, err := r.redisRegistryClient.Keys(r.ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (r RedisStorage) GetService(serviceId string) (types.Service, error) {
	val, err := r.redisServiceClient.Get(r.ctx, serviceId).Result()
	var service types.Service
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
