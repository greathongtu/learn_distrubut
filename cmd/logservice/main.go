package main

import (
	"context"
	"fmt"
	stlog "log"
	"se/log"
	"se/registry"
	"se/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName: "Log Service",
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(context.Background(), host, port,
		r, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shuting down log service.")
}
