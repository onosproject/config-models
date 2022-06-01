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

// NOTE that these tests require a specific environment to be set up,
// so we're skipping them unless the GNMI_E2E environment var is set to true
func skipIfNotEnabled(t *testing.T) {
	gnmiE2e := os.Getenv("GNMI_E2E")

	if gnmiE2e != "true" {
		t.Skip()
	}
}

func setup(t *testing.T) *testdevice.GnmiClient {

	skipIfNotEnabled(t)

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

	return testdevice.NewTestdeviceGnmiClient(gnmiConn)
}

func TestLeafAtTopLevel(t *testing.T) {
	client := setup(t)
	ctx := context.TODO()

	str := "ABC-123"
	val := &gnmi.TypedValue{
		Value: &gnmi.TypedValue_StringVal{StringVal: str},
	}
	setRes, err := client.Update_LeafAtTopLevel(ctx, target, val)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err := client.Get_LeafAtTopLevel(ctx, target)
	assert.NoError(t, err)
	assert.Equal(t, str, getRes)

	delRes, err := client.Delete_LeafAtTopLevel(ctx, target)
	assert.NoError(t, err)
	resId, err = gnmi_utils.ExtractResponseID(delRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err = client.Get_LeafAtTopLevel(ctx, target)
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
	setRes, err := client.Update_Cont1ACont2ALeaf2A(ctx, target, val)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err := client.Get_Cont1ACont2ALeaf2A(ctx, target)
	assert.NoError(t, err)
	assert.Equal(t, v, getRes)

	delRes, err := client.Delete_Cont1ACont2ALeaf2A(ctx, target)
	assert.NoError(t, err)
	resId, err = gnmi_utils.ExtractResponseID(delRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err = client.Get_Cont1ACont2ALeaf2A(ctx, target)
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
	setRes, err := client.Update_Cont1A_Cont2A(ctx, target, c1a_c2a)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err := client.Get_Cont1A_Cont2A(context.TODO(), target)

	assert.NoError(t, err)
	assert.Equal(t, c1a_c2a.Leaf2A, getRes.Leaf2A)

	delRes, err := client.Delete_Cont1A_Cont2A(ctx, target)
	assert.NoError(t, err)
	resId, err = gnmi_utils.ExtractResponseID(delRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err = client.Get_Cont1A_Cont2A(ctx, target)
	assert.Error(t, err)
	assert.Nil(t, getRes)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.NotFound)
}

func TestNestedContainer(t *testing.T) {
	// NOTE this container does have a nested child (Cont1A_Cont2A)
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
	setRes, err := client.Update_Cont1A(ctx, target, c1a)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err := client.Get_Cont1A(context.TODO(), target)

	assert.NoError(t, err)
	assert.Equal(t, c1a.Cont2A.Leaf2A, getRes.Cont2A.Leaf2A)
	assert.Equal(t, c1a.Leaf1A, getRes.Leaf1A)

	delRes, err := client.Delete_Cont1A(ctx, target)
	assert.NoError(t, err)
	resId, err = gnmi_utils.ExtractResponseID(delRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes, err = client.Get_Cont1A(ctx, target)
	assert.Error(t, err)
	assert.Nil(t, getRes)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.NotFound)
}

func TestListSingleKeyItem(t *testing.T) {
	// FIXME seems like we can't create a single item in the list
	// test operations on a single item in a list
	client := setup(t)
	ctx := context.TODO()

	item := "item"
	missingItem := "missing-item"
	rm := uint8(12)

	listItem := testdevice.OnfTest1_Cont1A_List2A{Name: &item, RangeMax: &rm}

	// for some reason the delete is breaking the test, so read the list and delete every entry in it
	//_, err := client.Delete_Cont1A_List2A(ctx, target)
	//assert.NoError(t, err)

	res, err := client.Get_Cont1A_List2A_List(ctx, target)
	if code, ok := status.FromError(err); !ok && code.Code() != codes.NotFound {
		// the only error we accept is not found
		assert.NoError(t, err)
	}

	for _, item := range res {
		_, err := client.Delete_Cont1A_List2A(ctx, target, *item.Name)
		assert.NoError(t, err)
	}

	setRes, err := client.Update_Cont1A_List2A(ctx, target, listItem)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes1, err := client.Get_Cont1A_List2A(context.TODO(), target, item)
	assert.NoError(t, err)
	assert.Equal(t, listItem.Name, getRes1.Name, "item name is wrong")
	assert.Equal(t, listItem.RangeMax, getRes1.RangeMax, "item range-max is wrong")

	getRes2, err := client.Get_Cont1A_List2A(context.TODO(), target, missingItem)
	assert.Error(t, err)
	assert.Nil(t, getRes2)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.NotFound)
}

func TestListSingleKey(t *testing.T) {
	client := setup(t)
	ctx := context.TODO()

	item1 := "item1"
	item2 := "item2"
	list := map[string]*testdevice.OnfTest1_Cont1A_List2A{
		item1: {Name: &item1},
		item2: {Name: &item2},
	}

	res, err := client.Get_Cont1A_List2A_List(ctx, target)
	assert.NoError(t, err)

	for _, item := range res {
		_, err := client.Delete_Cont1A_List2A(ctx, target, *item.Name)
		assert.NoError(t, err)
	}

	setRes, err := client.Update_Cont1A_List2A_List(ctx, target, list)
	assert.NoError(t, err)

	resId, err := gnmi_utils.ExtractResponseID(setRes)
	assert.NoError(t, err)
	assert.NotNil(t, resId)

	getRes1, err := client.Get_Cont1A_List2A_List(ctx, target)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(getRes1))
	assert.Equal(t, list[item1].Name, getRes1[item1].Name, "item1 is missing")
	assert.Equal(t, list[item2].Name, getRes1[item2].Name, "item2 is missing")
}
