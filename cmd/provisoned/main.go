package main

import (
	"context"
	"flag"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

var (
	apidAddr = flag.String("apidaddr", "battery_apid", "grpc address")
)

func main() {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.41"))
	// cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        panic(err)
    }
	// cliV := os.Getenv("")
	
	serviceName := "battery_apid"
	serviceID := "example_service.1"
	
	filter := filters.NewArgs()
    filter.Add("name", serviceName)
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{Filters: filter})
    if err != nil {
        panic(err)
    }
	
	for _, service := range services {
        if service.Spec.Name == serviceName {
            serviceID = service.ID
        }
    }

    // scale 명령어 실행
    _, err = cli.ServiceUpdate(context.Background(), serviceID, swarm.Version{

	}, swarm.ServiceSpec{
		
	}, types.ServiceUpdateOptions{})

    if err != nil {
        panic(err)
    }
}