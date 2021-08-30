package cmd

import (
	"fmt"
	"tugasgolang/routes"
	"tugasgolang/services"
)

func Run() {
	fmt.Println("masuk cmd")
	routes.Routes()
	services.ServiceGetCar()
}
