// Copyright 2016 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"github.com/vmware/vic/cmd/vic-machine/create"
	uninstall "github.com/vmware/vic/cmd/vic-machine/delete"
	"github.com/vmware/vic/cmd/vic-machine/inspect"
	"github.com/vmware/vic/pkg/errors"
)

var (
	Version  string
	BuildID  string
	CommitID string
)

func main() {
	app := cli.NewApp()

	app.Name = filepath.Base(os.Args[0])
	app.Usage = "Create and manage Virtual Container Hosts"
	app.EnableBashCompletion = true

	create := create.NewCreate()
	uninstall := uninstall.NewUninstall()
	inspect := inspect.NewInspect()
	app.Commands = []cli.Command{
		{
			Name:   "create",
			Usage:  "Deploy VCH",
			Action: create.Run,
			Flags:  create.Flags(),
		},
		{
			Name:   "delete",
			Usage:  "Delete VCH and associated resources",
			Action: uninstall.Run,
			Flags:  uninstall.Flags(),
		},
		{
			Name:   "inspect",
			Usage:  "Inspect VCH",
			Action: inspect.Run,
			Flags:  inspect.Flags(),
		},
		{
			Name:   "version",
			Usage:  "Show VIC version information",
			Action: showVersion,
		},
	}
	if Version != "" {
		app.Version = fmt.Sprintf("%s-%s-%s", Version, BuildID, CommitID)
	} else {
		app.Version = fmt.Sprintf("%s-%s", BuildID, CommitID)
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println("--------------------")
		fmt.Printf("%s failed: %s\n", app.Name, errors.ErrorStack(err))
	}
}

func showVersion(cli *cli.Context) error {
	fmt.Fprintf(cli.App.Writer, "%v version %v\n", cli.App.Name, cli.App.Version)
	return nil
}