{{ $path := .}}
path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {{  range $p := $path -}}
            {
                Name: "{{ $p }}",
            },
        {{ end -}}
        },
        Target: target,
    },
}