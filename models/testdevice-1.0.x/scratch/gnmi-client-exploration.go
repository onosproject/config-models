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
	client := NewGnmiClient(gnmiConn)

	val := &gnmi.TypedValue{
		Value: &gnmi.TypedValue_StringVal{StringVal: "ABC-123"},
	}
	ctx := context.TODO()
	setRes, err := client.GnmiSetLeafAtTopLevel(ctx, val)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI SET:")
	fmt.Println(setRes)

	getRes, err := client.GnmiGetLeafAtTopLevel(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("gNMI GET:")
	fmt.Println(getRes)

}

// start of the generated code

func NewGnmiClient(conn *grpc.ClientConn) *gnmiClient {
	gnmi_client := gnmi.NewGNMIClient(conn)
	return &gnmiClient{client: gnmi_client}
}

func (c *gnmiClient) GnmiGetLeafAtTopLevel(ctx context.Context) (*gnmi.GetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	leafAtTopLevelPath := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "leafAtTopLevel",
				},
			},
			Target: "target-foo",
		},
	}

	getRequest := &gnmi.GetRequest{
		Prefix:    nil,
		Path:      leafAtTopLevelPath,
		Type:      0,
		Encoding:  gnmi.Encoding_PROTO,
		UseModels: nil,
		Extension: nil,
	}

	return c.client.Get(gnmiCtx, getRequest)
}

func (c *gnmiClient) GnmiSetLeafAtTopLevel(ctx context.Context, val *gnmi.TypedValue) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	leafAtTopLevelPath := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "leafAtTopLevel",
				},
			},
			Target: "target-foo",
		},
	}

	setRequest := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path:       leafAtTopLevelPath[0],
				Value:      nil,
				Val:        val,
				Duplicates: 0,
			},
		},
		Extension: nil,
	}

	return c.client.Set(gnmiCtx, setRequest)
}
