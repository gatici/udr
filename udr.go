// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"fmt"
	"os"

	"github.com/omec-project/udr/logger"
	udr_service "github.com/omec-project/udr/service"
	"github.com/urfave/cli"
)

var UDR = &udr_service.UDR{}

func main() {
	app := cli.NewApp()
	app.Name = "udr"
	logger.AppLog.Infoln(app.Name)
	app.Usage = "-free5gccfg common configuration file -udrcfg udr configuration file"
	app.Action = action
	app.Flags = UDR.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		logger.AppLog.Errorf("UDR Run error: %v", err)
	}
}

func action(c *cli.Context) error {
	if err := UDR.Initialize(c); err != nil {
		logger.CfgLog.Errorf("%+v", err)
		return fmt.Errorf("failed to initialize")
	}

	UDR.Start()

	return nil
}
