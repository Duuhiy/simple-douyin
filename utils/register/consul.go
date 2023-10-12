package register

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type Registry struct {
	Host string
	Port int
}

func (r Registry) Register(address string, port int, name string, tags []string, id string) error {
	//TODO implement me
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		zap.L().Info("Register:api.NewClient(cfg): " + err.Error())
		return err
	}
	// 健康检查对象
	check := &api.AgentServiceCheck{
		Interval:                       "5s",
		Timeout:                        "5s",
		HTTP:                           fmt.Sprintf("http://%s:%d", address, port),
		DeregisterCriticalServiceAfter: "10s",
	}
	// 注册对象
	registration := api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: address,
		Check:   check,
	}
	// 注册到consul
	err = client.Agent().ServiceRegister(&registration)
	if err != nil {
		zap.L().Info("Register:client.Agent().ServiceRegister(&registration): " + err.Error())
		return err
	}
	return nil
}

func (r Registry) DeRegister(serviceId string) error {
	//TODO implement me
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		zap.L().Info("DeRegister:DeRegister: " + err.Error())
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		zap.L().Info("DeRegister:client.Agent().ServiceDeregister(serviceId): " + err.Error())
	}
	return err
}

type RegisterClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

func NewRegistry(host string, port int) RegisterClient {
	return &Registry{host, port}
}
