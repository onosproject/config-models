/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_utils

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	configapi "github.com/onosproject/onos-api/go/onos/config/v2"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"reflect"
)

var log = logging.GetLogger("gnmi-utils")

// GetResponseUpdate -- extract the single Update from the GetResponse
func GetResponseUpdate(gr *gnmi.GetResponse) (*gnmi.TypedValue, error) {
	if len(gr.Notification) != 1 {
		return nil, fmt.Errorf("unexpected number of GetResponse notifications %d", len(gr.Notification))
	}
	n0 := gr.Notification[0]
	if len(n0.Update) != 1 {
		return nil, fmt.Errorf("unexpected number of GetResponse notification updates %d", len(n0.Update))
	}
	u0 := n0.Update[0]
	if u0.Val == nil {
		return nil, nil
	}
	return &gnmi.TypedValue{
		Value: u0.Val.Value,
	}, nil
}

func hasNonZeroField(sf reflect.StructField, fv reflect.Value) bool {
	switch sf.Type.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if !fv.IsNil() {
			return true
		}
	// case reflect.Struct:
	// NOTE: if a nested struct is empty we want to return false, and that's covered by the default case
	// case reflect.Array:
	// TODO: call recursively for array elements
	default:
		ti := reflect.Zero(sf.Type).Interface() // default Nil value for this type
		vi := fv.Interface()                    // current value
		if ti != vi {
			return true
		}
	}
	return false
}

// PathToKey is a mapping to relate a particular path to it's key
// this is needed to generate the appropriate gnmi.SetRequest for lists
// as that information is not carried in the model
// NOTE we might consider to add in the Field Tags via YGOT (making a PR in the future)
type Path string
type Key interface{}
type PathToKey map[Path]Key

// NOTE in order to create the updates we need to iterate over the incoming
// data structure and create all the updates, unluckily we have to use reflection for that
func CreateGnmiSetForContainer(ctx context.Context, data interface{}, basePath *gnmi.Path, target string, pathKeys PathToKey) (*gnmi.SetRequest, error) {
	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{},
	}

	rv := reflect.ValueOf(data)
	rt := rv.Type()

	// iterate over all the fields in the model
	for i := 0; i < rt.NumField(); i++ {
		// reflect the field Type and Value
		fieldType := rt.Field(i)
		fieldValue := rv.Field(i)

		// construct the gNMI PathElement for this field
		var rpe = []*gnmi.PathElem{}
		f_path, ok := fieldType.Tag.Lookup("path")
		if ok && f_path != "" {
			rpe = append(basePath.Elem, &gnmi.PathElem{Name: f_path})
		}

		// if a value is set create a gNMI TypedValue and add it to the SetRequest
		if hasNonZeroField(fieldType, fieldValue) {
			// it means a value was assigned to that field

			var typedValue *gnmi.TypedValue
			val := reflect.Indirect(fieldValue)
			switch val.Kind() {
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				typedValue = &gnmi.TypedValue{
					Value: &gnmi.TypedValue_UintVal{
						UintVal: val.Uint(),
					},
				}
				up := &gnmi.Update{
					Path: &gnmi.Path{Elem: rpe, Target: target},
					Val:  typedValue,
				}
				req.Update = append(req.Update, up)
			case reflect.String:
				typedValue = &gnmi.TypedValue{
					Value: &gnmi.TypedValue_StringVal{
						StringVal: val.String(),
					},
				}
				up := &gnmi.Update{
					Path: &gnmi.Path{Elem: rpe, Target: target},
					Val:  typedValue,
				}
				req.Update = append(req.Update, up)
			case reflect.Struct:
				r, err := CreateGnmiSetForContainer(ctx, val.Interface(), &gnmi.Path{Elem: rpe, Target: target}, target, pathKeys)
				if err != nil {
					return nil, err
				}
				req.Update = append(req.Update, r.Update...)
			case reflect.Map:
				// if it is a Map (a.k.a YANG list) find the key name
				key, ok := pathKeys[Path(f_path)]
				if !ok {
					return nil, fmt.Errorf("cannot-find-key-name-for-element: %s", f_path)
				}

				// and then iterate over the elements in the list to create all the necessary updates
				for _, kv := range val.MapKeys() {

					newParentPath := &gnmi.Path{
						Elem: append(basePath.Elem, &gnmi.PathElem{
							Name: f_path,
							Key: map[string]string{
								fmt.Sprintf("%s", key): kv.String(),
							},
						}),
						Target: target,
					}
					itemInList := val.MapIndex(kv)
					item := reflect.Indirect(itemInList)

					r, err := CreateGnmiSetForContainer(ctx, item.Interface(), newParentPath, target, pathKeys)

					if err != nil {
						return nil, err
					}
					req.Update = append(req.Update, r.Update...)
				}

			default:
				log.Warnw("type-not-supported", "type", val.Kind())
			}

		}
	}

	return req, nil
}

// ExtractResponseID - the name of the change will be returned as extension 100
func ExtractResponseID(gnmiResponse *gnmi.SetResponse) (*string, error) {
	for _, ext := range gnmiResponse.Extension {
		switch extTyped := ext.Ext.(type) {
		case *gnmi_ext.Extension_RegisteredExt:
			// NOTE this is used in ONOS config
			if extTyped.RegisteredExt.Id == 100 {
				changeName := string(extTyped.RegisteredExt.Msg)
				return &changeName, nil
			}
			if extTyped.RegisteredExt.Id == configapi.TransactionInfoExtensionID {
				bytes := extTyped.RegisteredExt.Msg
				transactionInfo := &configapi.TransactionInfo{}
				err := proto.Unmarshal(bytes, transactionInfo)
				if err != nil {
					return nil, err
				}
				changeName := string(transactionInfo.ID)
				return &changeName, nil
			}
		}
	}

	return nil, fmt.Errorf("cannot find transaction ID in response")
}
