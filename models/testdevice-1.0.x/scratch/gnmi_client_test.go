/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	testdevice "github.com/onosproject/config-models/models/testdevice-1.0.x/api"
	"github.com/onosproject/config-models/pkg/gnmi-client-gen/gnmi_utils"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"os"
	"testing"
)

const address = "localhost:5150"
const target = "target-foo"

func setup(t *testing.T) *testdevice.GnmiClient {
	cert, err := tls.X509KeyPair([]byte(certs.DefaultClientCrt), []byte(certs.DefaultClientKey))
	if err != nil {
		fmt.Println("Error while Handling certs")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			}),
		),
	}
	gnmiConn, err := grpc.Dial(address, options...)
	if err != nil {
		fmt.Println("Error while Dialing gRPC")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return testdevice.NewOnfTest1GnmiClient(gnmiConn)
}

func TestLeafAtTopLevel(t *testing.T) {
	client := setup(t)
	ctx := context.TODO()

	str := "ABC-123"
	val := &gnmi.TypedValue{
		Value: &gnmi.TypedValue_StringVal{StringVal: str},
	}
	setRes, err := client.UpdateLeafAtTopLevel(ctx, target, val)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err := client.GetLeafAtTopLevel(ctx, target)
	assert.NoError(t, err)
	assert.Equal(t, str, getRes)

	delRes, err := client.DeleteLeafAtTopLevel(ctx, target)
	assert.NoError(t, err)
	resId, err = gnmi_utils.ExtractResponseID(delRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err = client.GetLeafAtTopLevel(ctx, target)
	assert.Error(t, err)
	assert.Equal(t, "", getRes)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.NotFound)
}

func TestNestedLeaf(t *testing.T) {
	client := setup(t)
	ctx := context.TODO()

	v := uint8(2)
	val := &gnmi.TypedValue{
		Value: &gnmi.TypedValue_UintVal{UintVal: uint64(v)},
	}
	setRes, err := client.UpdateCont1ACont2ALeaf2A(ctx, target, val)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err := client.GetCont1ACont2ALeaf2A(ctx, target)
	assert.NoError(t, err)
	assert.Equal(t, v, getRes)

	delRes, err := client.DeleteCont1ACont2ALeaf2A(ctx, target)
	assert.NoError(t, err)
	resId, err = gnmi_utils.ExtractResponseID(delRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err = client.GetCont1ACont2ALeaf2A(ctx, target)
	assert.Error(t, err)
	assert.Equal(t, uint8(0), getRes)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.NotFound)
}

func TestBasicContainer(t *testing.T) {
	// NOTE this container does not have nested children(s)
	client := setup(t)
	ctx := context.TODO()

	leaf2a := uint8(1)
	c1a_c2a := testdevice.OnfTest1_Cont1A_Cont2A{
		Leaf2A: &leaf2a,
	}
	setRes, err := client.UpdateCont1A_Cont2A(ctx, target, c1a_c2a)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err := client.GetCont1A_Cont2A(context.TODO(), target)

	assert.NoError(t, err)
	assert.Equal(t, c1a_c2a.Leaf2A, getRes.Leaf2A)

	delRes, err := client.DeleteCont1A_Cont2A(ctx, target)
	assert.NoError(t, err)
	resId, err = gnmi_utils.ExtractResponseID(delRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err = client.GetCont1A_Cont2A(ctx, target)
	assert.Error(t, err)
	assert.Nil(t, getRes)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.NotFound)
}

func TestNestedContainer(t *testing.T) {
	// NOTE this container does has a nested child (Cont1A_Cont2A)
	client := setup(t)
	ctx := context.TODO()

	leaf2a := uint8(1)
	str := "string"
	c1a_c2a := testdevice.OnfTest1_Cont1A_Cont2A{
		Leaf2A: &leaf2a,
	}
	c1a := testdevice.OnfTest1_Cont1A{
		Leaf1A: &str,
		Cont2A: &c1a_c2a,
	}
	setRes, err := client.UpdateCont1A(ctx, target, c1a)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err := client.GetCont1A(context.TODO(), target)

	assert.NoError(t, err)
	assert.Equal(t, c1a.Cont2A.Leaf2A, getRes.Cont2A.Leaf2A)
	assert.Equal(t, c1a.Leaf1A, getRes.Leaf1A)

	delRes, err := client.DeleteCont1A(ctx, target)
	assert.NoError(t, err)
	resId, err = gnmi_utils.ExtractResponseID(delRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err = client.GetCont1A(ctx, target)
	assert.Error(t, err)
	assert.Nil(t, getRes)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.NotFound)
}
