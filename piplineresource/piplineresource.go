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
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	"tektoncd-demo/config"
)

func main() {

	// k8s 配置文件
	restConfig, err := config.GetRestConf()
	if err != nil {
		klog.Fatal(err)
	}
	// pipeline clientset
	tektonClient := versioned.NewForConfigOrDie(restConfig)
	pipresource := &v1alpha1.PipelineResource{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      "pipresource",
			Namespace: "default",
		},
		Spec: v1alpha1.PipelineResourceSpec{
			Description: "this is demo!",
			Type:        v1alpha1.PipelineResourceTypeGit,
			Params: []v1alpha1.ResourceParam{
				{
					Name:  "url",
					Value: "https://github.com/knative-sample/tekton-knative.git",
				},
				{
					Name:  "revision",
					Value: "master",
				},
			},
		},
		Status: nil,
	}
	meta, err := tektonClient.TektonV1alpha1().PipelineResources("default").Create(context.Background(), pipresource, v1.CreateOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	klog.Info(meta)
}
