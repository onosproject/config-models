/*
 * SPDX-FileCopyrightText: 2023-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"context"
	"github.com/onosproject/config-models/models/testdevice-1.0.x/api"
	"github.com/onosproject/config-models/pkg/path"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"os"
	"testing"
	"time"
)

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func newTestService() (northbound.Service, error) {

	return &modelPlugin{}, nil
}

func createServerConnection(t *testing.T) *grpc.ClientConn {
	lis = bufconn.Listen(1024 * 1024)
	s, err := newTestService()
	assert.NoError(t, err)
	assert.NotNil(t, s)
	server := grpc.NewServer()
	s.Register(server)

	go func() {
		if err := server.Serve(lis); err != nil {
			assert.NoError(t, err, "Server exited with error: %v", err)
		}
	}()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	return conn
}

func TestMain(m *testing.M) {
	entries, err := api.UnzipSchema()
	if err != nil {
		log.Fatalf("Unable to extract model schema: %+v", err)
	}
	roPaths, rwPaths, namespaceMappings = path.ExtractPaths(entries)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestServer_GetModelInfo(t *testing.T) {
	conn := createServerConnection(t)
	client := admin.NewModelPluginServiceClient(conn)

	mi, err := client.GetModelInfo(context.Background(), &admin.ModelInfoRequest{})
	assert.NoError(t, err)
	assert.NotNil(t, mi)
	assert.Equal(t, "testdevice", mi.GetModelInfo().Name)
	assert.Equal(t, "1.0.x", mi.GetModelInfo().Version)
	assert.Equal(t, 3, len(mi.GetModelInfo().ReadOnlyPath))
	assert.Equal(t, 64, len(mi.GetModelInfo().ReadWritePath))
}

func TestServer_ValidateConfig(t *testing.T) {
	conn := createServerConnection(t)
	client := admin.NewModelPluginServiceClient(conn)

	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config.json")
	if err != nil {
		assert.NoError(t, err)
	}

	response, err := client.ValidateConfig(context.Background(), &admin.ValidateConfigRequest{
		Json: sampleConfig,
	})
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, true, response.Valid)
}

func TestServer_ValidateConfigChunked(t *testing.T) {
	conn := createServerConnection(t)
	defer conn.Close()
	client := admin.NewModelPluginServiceClient(conn)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.SetLevel(logging.DebugLevel)
	defer log.SetLevel(logging.InfoLevel)

	type testValidate struct {
		name       string
		jsonConfig []byte
		chunkSize  int
		expResult  bool
		expMsg     string
		expError   error
	}

	// jsonConfig is 1712 bytes
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config.json")
	if err != nil {
		assert.NoError(t, err)
	}

	validationTests := []testValidate{
		{
			name:       "one chunk test case",
			jsonConfig: sampleConfig,
			chunkSize:  65535,
			expResult:  true,
			expError:   nil,
		},
		{
			name:       "four chunks test case",
			jsonConfig: sampleConfig,
			chunkSize:  500,
			expResult:  true,
			expError:   nil,
		},
		{
			name:       "no data test case",
			jsonConfig: []byte{},
			chunkSize:  500,
			expError:   status.Error(codes.InvalidArgument, "Unable to unmarshal JSON: unexpected end of JSON input"),
		},
	}

	for _, validationTest := range validationTests {
		sender, err := client.ValidateConfigChunked(ctxTimeout)
		assert.NoError(t, err)
		assert.NotNil(t, sender)

		jsonLen := len(validationTest.jsonConfig)
		position := 0
		for position < jsonLen {
			var chunk []byte
			if position+validationTest.chunkSize < jsonLen {
				chunk = sampleConfig[position : position+validationTest.chunkSize]
				position += validationTest.chunkSize
			} else {
				chunk = sampleConfig[position:]
				position = jsonLen
			}
			err := sender.Send(&admin.ValidateConfigRequestChunk{
				Json: chunk,
			})
			if err != nil {
				t.Fatalf("test '%s': sending validation chunk failed. %v", validationTest.name, err)
			}
		}
		resp, err := sender.CloseAndRecv()
		if validationTest.expError != nil {
			assert.EqualError(t, err, validationTest.expError.Error(), validationTest.name)
			continue
		}
		assert.NoError(t, err, validationTest.name)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotNil(t, resp, validationTest.name)
		assert.Equal(t, validationTest.expResult, resp.Valid, validationTest.name)
	}
}

func TestServer_GetValueSelectionChunked(t *testing.T) {
	conn := createServerConnection(t)
	defer conn.Close()
	client := admin.NewModelPluginServiceClient(conn)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.SetLevel(logging.DebugLevel)
	defer log.SetLevel(logging.InfoLevel)

	type testGetValueSelection struct {
		name       string
		selection  string
		jsonConfig []byte
		chunkSize  int
		expResult  []string
		expMsg     string
		expError   error
	}

	// jsonConfig is 1712 bytes
	switchConfig, err := os.ReadFile("../testdata/switch-config-example-1.json")
	if err != nil {
		assert.NoError(t, err)
	}

	getValueSelectionTests := []testGetValueSelection{
		{
			name:       "one chunk test case",
			selection:  "/switch[switch-id=san-jose-edge-nic]/port[cage-number=1][channel-number=0]/speed",
			jsonConfig: switchConfig,
			chunkSize:  65535,
			expResult:  []string{"[speed-1g speed-10g speed-100g]"}, // TODO fix so it returns an actual array
		},
		{
			name:       "four chunks test case",
			selection:  "/switch[switch-id=san-jose-edge-nic]/port[cage-number=1][channel-number=0]/speed",
			jsonConfig: switchConfig,
			chunkSize:  500,
			expResult:  []string{"[speed-1g speed-10g speed-100g]"}, // TODO fix so it returns an actual array
		},
		{
			name:       "no data test case",
			selection:  "/switch[switch-id=san-jose-edge-nic]/port[cage-number=1][channel-number=0]/speed",
			jsonConfig: []byte{},
			chunkSize:  500,
			expError:   status.Error(codes.InvalidArgument, "Unable to unmarshal JSON: unexpected end of JSON input"),
		},
		{
			name:       "invalid selection test case",
			selection:  "invalid selection",
			jsonConfig: switchConfig,
			chunkSize:  500,
			expError:   status.Error(codes.InvalidArgument, "Unable to navigate to selection. navigatedValue path is invalid invalid selection"),
		},
	}

	for _, getValueSelectionTest := range getValueSelectionTests {
		sender, err := client.GetValueSelectionChunked(ctxTimeout)
		assert.NoError(t, err, getValueSelectionTest.name)
		assert.NotNil(t, sender, getValueSelectionTest.name)

		jsonLen := len(getValueSelectionTest.jsonConfig)
		position := 0
		for position < jsonLen {
			var chunk []byte
			if position+getValueSelectionTest.chunkSize < jsonLen {
				chunk = switchConfig[position : position+getValueSelectionTest.chunkSize]
				position += getValueSelectionTest.chunkSize
			} else {
				chunk = switchConfig[position:]
				position = jsonLen
			}
			err := sender.Send(&admin.ValueSelectionRequestChunk{
				SelectionPath: getValueSelectionTest.selection, // selection is sent with every chunk
				ConfigJson:    chunk,
			})
			if err != nil {
				t.Fatalf("test '%s': sending getValueSelection chunk failed. %v", getValueSelectionTest.name, err)
			}
		}
		resp, err := sender.CloseAndRecv()
		if getValueSelectionTest.expError != nil {
			assert.EqualError(t, err, getValueSelectionTest.expError.Error(), getValueSelectionTest.name)
			continue
		}
		assert.NoError(t, err, getValueSelectionTest.name)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotNil(t, resp, getValueSelectionTest.name)
		assert.Equal(t, getValueSelectionTest.expResult, resp.Selection, getValueSelectionTest.name)
	}
}
