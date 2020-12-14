/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	tknResource "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	"tektoncd-demo/pkg/tekton"

	"k8s.io/klog"
	"tektoncd-demo/config"
	"tektoncd-demo/pkg/tekton/springboot"
)

func main() {
	// k8s 配置文件
	restConfig, err := config.GetRestConf()
	if err != nil {
		klog.Fatal(err)
	}
	tektonClient := versioned.NewForConfigOrDie(restConfig)
	tektonClientRes := tknResource.NewForConfigOrDie(restConfig)

	spboot := springboot.SpringBoot{
		Name:                 "wolong",
		NameSpace:            "test",
		Revision:             "Master",
		Url:                  "https://www.a.com",
		TektonClient:         tektonClient,
		TektonClientResource: tektonClientRes,
	}
	var cicd tekton.CICD
	cicd = &spboot
	if err = cicd.Run(); err != nil {
		klog.Error(err)
	}
}
