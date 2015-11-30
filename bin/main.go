package main

import (
    "github.com/ubiquityhosting/dm_driver"
    "github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
    plugin.RegisterDriver(ubiquity.NewDriver("", ""))
}