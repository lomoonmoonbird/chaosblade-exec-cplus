/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"log"
	"os"

	"github.com/lomoonmoonbird/chaosblade-spec-go/spec"
	"github.com/lomoonmoonbird/chaosblade-spec-go/util"

	"github.com/lomoonmoonbird/chaosblade-exec-cplus/module"
)

func main() {
	if len(os.Args) != 2 {
		log.Panicln("less yaml file path")
	}
	err := util.CreateYamlFile(getModels(), os.Args[1])
	if err != nil {
		log.Panicf("create yaml file error, %v", err)
	}
}

// getModels returns experiment models in the project
func getModels() *spec.Models {
	modelCommandSpecs := []spec.ExpModelCommandSpec{
		module.NewCPlusCommandModelSpec(),
	}
	prepareModel := spec.ExpPrepareModel{
		PrepareType: "cplus",
		PrepareFlags: []spec.ExpFlag{
			{
				Name: "port",
				Desc: "server port to be listening",
			},
			{
				Name: "ip",
				Desc: "The ip address bound to the service",
			},
		},
		PrepareRequired: true,
	}
	specModels := make([]*spec.Models, 0)
	for _, modeSpec := range modelCommandSpecs {
		specModel := util.ConvertSpecToModels(modeSpec, prepareModel, "host")
		specModels = append(specModels, specModel)
	}
	return util.MergeModels(specModels...)
}
