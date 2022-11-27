package main

import (
	"context"
	"distributed_system/log"
	"distributed_system/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed_system.log")
	host, port := "localhost", "4000"
	ctx, err := service.Start(context.Background(), "Log Service", host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<- ctx.Done()
	fmt.Println("shutting down log service.")
}
