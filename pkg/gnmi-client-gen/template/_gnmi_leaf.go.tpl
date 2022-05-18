{{ $ep := . }}
func (c *GnmiClient) {{ $ep.MethodName }}(ctx context.Context, target string,{{ if eq $ep.Method "update"}} val *gnmi.TypedValue,{{end}}
) ({{ if eq $ep.Method "get"}}{{ $ep.GoType }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()

{{ template "_gnmi_path.go.tpl" $ep.Path }}

{{ if eq $ep.Method "get" -}}
    req := &gnmi.GetRequest{
    Encoding:  gnmi.Encoding_PROTO,
    Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return {{ $ep.GoEmptyReturnType }}, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return {{ $ep.GoEmptyReturnType }}, err
    }

    if {{ $ep.GoReturnType }} ==  {{ $ep.GoEmptyReturnType }} {
    return {{ $ep.GoEmptyReturnType }}, status.Error(codes.NotFound, "{{ $ep.ModelName }}-not-found")
    }

    return {{ $ep.GoReturnType }}, nil
{{ end -}}

{{ if eq $ep.Method "update" -}}
    req := &gnmi.SetRequest{
    Update: []*gnmi.Update{
    {
    Path: path[0],
    Val:  val,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
{{ end -}}

{{ if eq $ep.Method "delete" -}}
    req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
{{ end -}}
}