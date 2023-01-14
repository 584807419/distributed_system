package main

import (
	"context"
	"distributed_system/log"
	"distributed_system/registry"
	"distributed_system/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed_system.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
	}
	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("shutting down log service.")
}
