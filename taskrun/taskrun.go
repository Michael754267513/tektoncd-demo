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
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
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
	// pipline clientset
	tektonClient := versioned.NewForConfigOrDie(restConfig)

	taskRun := &v1beta1.TaskRun{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      "taskrun",
			Namespace: "default",
		},
		Spec: v1beta1.TaskRunSpec{
			TaskRef: &v1beta1.TaskRef{
				Name: "testrun",
			},
		},
	}

	meta, err := tektonClient.TektonV1beta1().TaskRuns("default").Create(context.Background(), taskRun, v1.CreateOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	klog.Info(meta)

}
