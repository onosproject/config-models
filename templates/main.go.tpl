// Code generated by model-compiler. DO NOT EDIT.

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
	"fmt"
	"github.com/onosproject/config-models/models/{{ .Name }}-{{ .Version }}/api"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("plugin")

type modelPlugin struct {
}

type server struct {
}

var modelData = []*gnmi.ModelData{
    {{- range .ModelData }}
	{Name: {{ .Name | quote }}, Organization: {{ .Organization | quote }}, Version: {{ .Version | quote }}},
	{{- end }}
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
			Name:           {{ .Name | quote }},
			Version:        {{ .Version | quote }},
			ModelData:      modelData,
			GetStateMode:   {{ .GetStateMode }},
		},
	}, nil
}

func (s server) ValidateConfig(ctx context.Context, request *admin.ValidateConfigRequest) (*admin.ValidateConfigResponse, error) {
	gostruct, err := s.UnmarshallConfigValues(request.Json)
	if err != nil {
		return nil, err
	}
	err = s.Validate(gostruct)
	if err != nil {
		return nil, err
	}
	return &admin.ValidateConfigResponse{Valid: true}, nil
}

// UnmarshallConfigValues allows Device to implement the Unmarshaller interface
func (s server) UnmarshallConfigValues(jsonTree []byte) (*ygot.ValidatedGoStruct, error) {
	device := &api.Device{}
	vgs := ygot.ValidatedGoStruct(device)

	if err := api.Unmarshal([]byte(jsonTree), device); err != nil {
		return nil, err
	}

	return &vgs, nil
}

func (s server) Validate(ygotModel *ygot.ValidatedGoStruct, opts ...ygot.ValidationOption) error {
	deviceDeref := *ygotModel
	device, ok := deviceDeref.(*api.Device)
	if !ok {
		return fmt.Errorf("Unable to convert model {{ .Name }}-{{ .Version }}")
	}
	return device.Validate()
}
