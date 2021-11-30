// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("plugin")

type modelPlugin struct {
}

type server struct {
}

func (p *modelPlugin) Register(gs *grpc.Server) {
	server := &server{}
	admin.RegisterModelPluginServiceServer(gs, server)
}

func main() {
	ready := make(chan bool)

	// Start gRPC server
	log.Info("Starting model plugin")
	p := modelPlugin{}
	if err := p.startNorthboundServer(); err != nil {
		log.Fatal("Unable to start model plugin service", err)
	}

	// Serve
	<-ready
}

func (p *modelPlugin) startNorthboundServer() error {
	cfg := northbound.NewServerConfig("", "", "", 5152, false)
	s := northbound.NewServer(cfg)

	s.AddService(p)

	doneCh := make(chan error)
	go func() {
		err := s.Serve(func(started string) {
			log.Info("Started NBI on ", started)
			close(doneCh)
		})
		if err != nil {
			doneCh <- err
		}
	}()
	return <-doneCh
}

func (s server) GetModelInfo(ctx context.Context, request *admin.ModelInfoRequest) (*admin.ModelInfoResponse, error) {
	return &admin.ModelInfoResponse{
		ModelInfo: &admin.ModelInfo{
			Name:      {{ .Name | quote }},
			Version:   {{ .Version | quote }},
			ModelData: nil,
		},
	}, nil
}

func (s server) ValidateConfig(ctx context.Context, request *admin.ValidateConfigRequest) (*admin.ValidateConfigResponse, error) {
	panic("implement me")
}
