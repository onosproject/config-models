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
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"time"
)

type gnmiClient struct {
	client gnmi.GNMIClient
}

const target = "target-foo"

func main() {

	// setup the gRPC connection
	const address = "localhost:5150"
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

	// create an instance of the gNMI model client
	// this the one we want to autogenerate
	client := testdevice.NewOnfTest1GnmiClient(gnmiConn)
	ctx := context.TODO()

	val := &gnmi.TypedValue{
		Value: &gnmi.TypedValue_StringVal{StringVal: "ABC-123"},
	}
	setRes, err := client.UpdateLeafattoplevel(ctx, target, val)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI SET:")
	fmt.Println(setRes)

	getRes, err := client.GetLeafattoplevel(ctx, target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI GET:")
	fmt.Println(getRes)

	val = &gnmi.TypedValue{
		Value: &gnmi.TypedValue_UintVal{UintVal: 2},
	}
	setRes, err = client.UpdateCont1aCont2aLeaf2a(ctx, target, val)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI SET nested:")
	fmt.Println(setRes)

	getNestedRes, err := client.GetCont1aCont2aLeaf2a(ctx, target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI GET nested:")
	fmt.Println(getNestedRes)

	cont1a, err := GetCont1aJson(gnmi.NewGNMIClient(gnmiConn), context.TODO(), target)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI GET struct via json:")
	fmt.Println(fmt.Sprintf("%v", cont1a))

	cont1a_2, err := GetCont1aProto(gnmi.NewGNMIClient(gnmiConn), context.TODO(), target)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI GET struct via proto:")
	fmt.Println(fmt.Sprintf("%v", cont1a_2))
}

func GetCont1aJson(client gnmi.GNMIClient, ctx context.Context, target string,
) (*testdevice.OnfTest1_Cont1A, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_JSON,
		Path:     path,
	}
	res, err := client.Get(gnmiCtx, req)

	if err != nil {
		return nil, err
	}

	val, err := testdevice.GetResponseUpdate(res)

	if err != nil {
		return nil, err
	}

	json := val.GetJsonVal()
	st := testdevice.Device{}
	testdevice.Unmarshal(json, &st)

	return st.Cont1A, nil
}

func GetCont1aProto(client gnmi.GNMIClient, ctx context.Context, target string,
) (*testdevice.OnfTest1_Cont1A, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := client.Get(gnmiCtx, req)

	if err != nil {
		return nil, err
	}

	val, err := testdevice.GetResponseUpdate(res)

	if err != nil {
		return nil, err
	}

	json := val.GetJsonVal()
	st := testdevice.Device{}
	testdevice.Unmarshal(json, &st)

	return st.Cont1A, nil
}
