package main

import (
	managementapi "github.com/breeve/cloudovn/pkg/api/management/v1"
	"github.com/breeve/cloudovn/pkg/utils"
	"github.com/spf13/cobra"
)

func main() {
	_ = cobra.Command{}
	utils.LogInit()
	managementapi.NewVPCServiceClient(nil)
}
