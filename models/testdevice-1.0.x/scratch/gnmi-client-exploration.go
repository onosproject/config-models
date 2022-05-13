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
	setRes, err := client.UpdateLeafAtTopLevel(ctx, target, val)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI SET:")
	fmt.Println(setRes)

	getRes, err := client.GetLeafAtTopLevel(ctx, target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI GET:")
	fmt.Println(getRes)

	val = &gnmi.TypedValue{
		Value: &gnmi.TypedValue_UintVal{UintVal: 2},
	}
	setRes, err = client.UpdateCont1ACont2ALeaf2A(ctx, target, val)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI SET nested:")
	fmt.Println(setRes)

	getNestedRes, err := client.GetCont1ACont2ALeaf2A(ctx, target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI GET nested:")
	fmt.Println(getNestedRes)

	cont1a, err := client.GetCont1A(context.TODO(), target)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI GET Cont1A:")
	fmt.Println(fmt.Sprintf("%v", cont1a))

	cont1a_cont2a, err := client.GetCont1A_Cont2A(context.TODO(), target)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI GET GetCont1A_Cont2A:")
	fmt.Println(fmt.Sprintf("%v", cont1a_cont2a))
}
