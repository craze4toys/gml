/*
 *  GML - Go QML
 *  Copyright (c) 2019 Roland Singer [roland.singer@deserbit.com]
 *  Copyright (c) 2019 Sebastian Borchers [sebastian@deserbit.com]
 */

package build

import "text/template"

var cppSourceTmpl = template.Must(template.New("t").Funcs(tmplFuncMap).Parse(cppSourceTmplText))

const cppSourceTmplText = `// This file is auto-generated by gml.
#include "{{.PackageName}}.h"

{{/* Struct loop */ -}}
{{range $struct := .Structs -}}
//###
//### {{$struct.Name}}
//###

{{$struct.CBaseName}} {{$struct.CBaseName}}_new() {
    auto _vv = new {{$struct.CPPBaseName}}();
    return (void*)_vv;
}

void {{$struct.CBaseName}}_free({{$struct.CBaseName}} _v) {
    if (_v == NULL) return;
    auto _vv = ({{$struct.CPPBaseName}}*)_v;
    delete _vv;
    _v = NULL;
}

{{/* Signals */ -}}
{{- range $signal := $struct.Signals }}
void {{$struct.CBaseName}}_{{$signal.Name}}({{$struct.CBaseName}} _v{{cParams $signal.Params false}}) {
    auto _vv = ({{$struct.CPPBaseName}}*)_v;
    emit _vv->{{$signal.CPPName}}();
}
{{end}}

{{- /* End of struct loop */ -}}
{{- end}}
`

var cppHeaderTmpl = template.Must(template.New("t").Funcs(tmplFuncMap).Parse(cppHeaderTmplText))

const cppHeaderTmplText = `// This file is auto-generated by gml.
#ifndef GML_GEN_CPP_{{.PackageName}}_H
#define GML_GEN_CPP_{{.PackageName}}_H

#include "../gen_c/{{.PackageName}}.h"
#include <QObject>

{{/* Struct loop */ -}}
{{range $struct := .Structs -}}
//###
//### {{$struct.Name}}
//###

class {{$struct.CPPBaseName}} : public QObject
{
    Q_OBJECT

signals:
{{- /* Signals */ -}}
{{- range $signal := $struct.Signals }}
    void {{$signal.CPPName}}({{cParams $signal.Params true}});
{{- end}}
};

{{- /* End of struct loop */ -}}
{{- end}}

#endif
`
