package main

import (
	"github.com/onosproject/onos-config-model/{{ .Model.Name }}_{{ .Model.Version | replace "." "_" }}/model"
)

var ConfigModelPlugin configmodel.ConfigModelPlugin
