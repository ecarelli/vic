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

package container

import (
	"github.com/docker/docker/runconfig"
	containertypes "github.com/docker/engine-api/types/container"
)

// VicContainer is VIC's abridged version of Docker's container object.
type VicContainer struct {
	*runconfig.StreamConfig

	Name        string
	ID          string
	ContainerID string
	Config      *containertypes.Config
}

func NewVicContainer() *VicContainer {
	return &VicContainer{
		StreamConfig: runconfig.NewStreamConfig(),
		Config:       &containertypes.Config{},
	}
}