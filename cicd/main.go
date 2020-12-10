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
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/klog"
	"tektoncd-demo/cicd/resource"
	"tektoncd-demo/cicd/task"
)

func main() {
	var (
		name           string                        = "cicd"
		namespace      string                        = "default"
		resourceType   v1alpha1.PipelineResourceType = v1alpha1.PipelineResourceTypeGit
		resourceParams []v1alpha1.ResourceParam      = []v1alpha1.ResourceParam{
			{
				Name:  "revision",
				Value: "master",
			},
			{
				Name:  "url",
				Value: "https://github.com/knative-sample/tekton-knative.git",
			},
		}
		inputs []v1beta1.TaskResource = []v1beta1.TaskResource{
			{},
		}
	)
	// 创建资源
	if res, err := resource.Resource(name, namespace, resourceType, resourceParams); !res {
		klog.Fatal(err)
	}
	task.SourceToImage(name, namespace)

}
