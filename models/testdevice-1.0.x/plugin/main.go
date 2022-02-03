// Code generated by model-compiler. DO NOT EDIT.

// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"github.com/onosproject/config-models/models/testdevice/api"
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"google.golang.org/grpc"
	"os"
	"strconv"
)

var log = logging.GetLogger("plugin")

type modelPlugin struct {
}

type server struct {
}

var modelData = []*gnmi.ModelData{
	{Name: "onf-test1", Organization: "Open Networking Foundation", Version: "2018-02-20"},
	{Name: "onf-test1-extra", Organization: "Open Networking Foundation", Version: "2021-04-01"},
}

var encodings = []gnmi.Encoding{gnmi.Encoding_JSON_IETF}

func (p *modelPlugin) Register(gs *grpc.Server) {
	log.Info("Registering model plugin service")
	server := &server{}
	admin.RegisterModelPluginServiceServer(gs, server)
}

func main() {
	ready := make(chan bool)

	if len(os.Args) < 2 {
		log.Fatal("gRPC port argument is required")
		os.Exit(1)
	}

	i, err := strconv.ParseInt(os.Args[1], 10, 16)
	if err != nil {
		log.Fatal("specified gRPC port is invalid", err)
		os.Exit(1)
	}
	port := int16(i)

	roPaths, rwPaths = extractPaths()

	// Start gRPC server
	log.Info("Starting model plugin")
	p := modelPlugin{}
	if err := p.startNorthboundServer(port); err != nil {
		log.Fatal("Unable to start model plugin service", err)
	}

	// Serve
	<-ready
}

func (p *modelPlugin) startNorthboundServer(port int16) error {
	cfg := northbound.NewServerConfig("", "", "", port, true)
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
	log.Infof("Received model info request: %+v", request)
	return &admin.ModelInfoResponse{
		ModelInfo: &admin.ModelInfo{
			Name:               "testdevice",
			Version:            "1.0.0",
			ModelData:          modelData,
			SupportedEncodings: encodings,
			GetStateMode:       0,
			ReadOnlyPath:       roPaths,
			ReadWritePath:      rwPaths,
		},
	}, nil
}

func (s server) ValidateConfig(ctx context.Context, request *admin.ValidateConfigRequest) (*admin.ValidateConfigResponse, error) {
	log.Infof("Received validate config request: %s", request.String())
	gostruct, err := s.unmarshallConfigValues(request.Json)
	if err != nil {
		return nil, errors.Status(err).Err()
	}

	if err := s.validate(gostruct); err != nil {
		return nil, errors.Status(err).Err()
	}

	if err := s.validateMust(*gostruct); err != nil {
		return nil, errors.Status(err).Err()
	}
	return &admin.ValidateConfigResponse{Valid: true}, nil
}

func (s server) GetPathValues(ctx context.Context, request *admin.PathValuesRequest) (*admin.PathValuesResponse, error) {
	log.Infof("Received path values request: %+v", request)
	pathValues, err := getPathValues(request.PathPrefix, request.Json)
	if err != nil {
		return nil, errors.Status(errors.NewInvalid("Unable to get path values: %+v", err)).Err()
	}
	return &admin.PathValuesResponse{PathValues: pathValues}, nil
}

func (s server) unmarshallConfigValues(jsonTree []byte) (*ygot.ValidatedGoStruct, error) {
	device := &api.Device{}
	vgs := ygot.ValidatedGoStruct(device)
	if err := api.Unmarshal([]byte(jsonTree), device); err != nil {
		return nil, errors.NewInvalid("Unable to unmarshal JSON: %+v", err)
	}
	return &vgs, nil
}

func (s server) validate(ygotModel *ygot.ValidatedGoStruct, opts ...ygot.ValidationOption) error {
	deviceDeref := *ygotModel
	device, ok := deviceDeref.(*api.Device)
	if !ok {
		return errors.NewInvalid("Unable to convert model testdevice-1.0.0")
	}
	return device.Validate()
}

func (s server) validateMust(device ygot.ValidatedGoStruct) error {
	log.Infof("Received validateMust request for device: %v", device)
	schema, err := api.Schema()
	if err != nil {
		return errors.NewInvalid("Unable to get schema: %+v", err)
	}

	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	ynn, ok := nn.(*navigator.YangNodeNavigator)
	if !ok {
		return errors.NewInvalid("Cannot cast NodeNavigator to YangNodeNavigator")
	}
	return ynn.WalkAndValidateMust()
}

