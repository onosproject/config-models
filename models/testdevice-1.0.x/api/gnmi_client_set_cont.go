/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package api

import (
	"context"
	"github.com/openconfig/gnmi/proto/gnmi"
	"reflect"
	"time"
)

func (c *GnmiClient) SetCont1A_Cont2A(ctx context.Context, target string, data OnfTest1_Cont1A_Cont2A,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},
				{
					Name: "cont2a",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			//{
			//	Path: &gnmi.Path{Elem: leaf2apath, Target: target},
			//	Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_UintVal{UintVal: uint64(*data.Leaf2A)}},
			//},
		},
	}

	// NOTE in order to create the updates we need to iterate over the incoming
	// data structure and create all the updates, unluckily we have to use reflection for that

	rs := reflect.ValueOf(data)
	rt := rs.Type()
	for i := 0; i < rt.NumField(); i++ {
		var rpe = []*gnmi.PathElem{}
		fieldType := rt.Field(i)
		fieldValue := rs.Field(i)
		f_path, ok := fieldType.Tag.Lookup("path")
		if ok && f_path != "" {
			rpe = append(path[0].Elem, &gnmi.PathElem{Name: f_path})
		}

		if hasNonZeroField(fieldType, fieldValue) {
			// it means a value was assigned to that field

			var rval *gnmi.TypedValue
			val := reflect.Indirect(fieldValue)
			switch val.Kind() {
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				rval = &gnmi.TypedValue{
					Value: &gnmi.TypedValue_UintVal{
						UintVal: val.Uint(),
					},
				}
			}

			up := &gnmi.Update{
				Path: &gnmi.Path{Elem: rpe},
				Val:  rval,
			}
			req.Update = append(req.Update, up)
		}
	}

	return c.client.Set(gnmiCtx, req)

}

func hasNonZeroField(sf reflect.StructField, fv reflect.Value) bool {
	switch sf.Type.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if !fv.IsNil() {
			return true
		}
	// case reflect.Struct:
	// TODO: call recursively for nested structs
	// case reflect.Array:
	// TODO: call recursively for array elements
	default:
		if reflect.Zero(sf.Type).Interface() != fv.Interface() {
			return true
		}
	}
	return false
}
