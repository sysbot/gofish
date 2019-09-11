package main

import (
	"fmt"

	gofish "github.com/stmcginnis/gofish"
	redfish "github.com/stmcginnis/gofish/redfish"
)

func main() {
	// Create a new instance of gofish client, ignoring self-signed certs
	config := gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	}

	c, err := gofish.Connect(config)
	if err != nil {
		panic(err)
	}
	defer c.Logout()

	// Attached the client to service root
	service := c.Service

	// Query the computer systems
	ss, err := service.Systems()
	if err != nil {
		panic(err)
	}

	// Creates a boot override to pxe once
	bootOverride := redfish.Boot{
		BootSourceOverrideTarget:  redfish.PxeBootSourceOverrideTarget,
		BootSourceOverrideEnabled: redfish.OnceBootSourceOverrideEnabled,
	}

	for _, system := range ss {
		fmt.Printf("System: %#v\n\n", system)
		_, err := system.SetBoot(bootOverride)
		if err != nil {
			panic(err)
		}
		_, err = system.Reset(redfish.ForceRestartResetType)
		if err != nil {
			panic(err)
		}
	}
}
