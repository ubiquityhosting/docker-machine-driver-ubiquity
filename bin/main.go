package main

import (
    "github.com/ubiquityhosting/docker-machine-driver-ubiquity"
    "github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
    plugin.RegisterDriver(ubiquity.NewDriver("", ""))
}