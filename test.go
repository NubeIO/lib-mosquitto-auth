package main

import (
	"fmt"
	"github.com/NubeIO/lib-mosquitto-auth/mosquitto"
)

func main() {
	err := mosquitto.WriteConfig(&mosquitto.Config{
		Path: "mosquitto/mosquitto.conf.test",
		Security: mosquitto.Security{
			SSL:                false,
			ClientVerification: false,
		},
		AccessControl: mosquitto.AccessControl{
			Anonymous: false,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
