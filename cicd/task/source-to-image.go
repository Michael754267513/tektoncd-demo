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

package task

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"

	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"tektoncd-demo/config"
)

func SourceToImage(name string, namespace string, inputs []v1beta1.TaskResource, outputs []v1beta1.TaskResource, params []v1beta1.ParamSpec) (res bool, err error) {
	var (
		steps     []v1beta1.Step
		task_meta *v1beta1.Task
	)
	// k8s 配置文件
	restConfig, err := config.GetRestConf()
	if err != nil {
		klog.Fatal(err)
	}
	// pipline clientset
	tektonClient := versioned.NewForConfigOrDie(restConfig)
	// step
	steps = append(steps, v1beta1.Step{
		Container: corev1.Container{
			Name:    "build-and-push",
			Image:   "registry.cn-hangzhou.aliyuncs.com/knative-sample/kaniko-project-executor:v0.10.0",
			Command: []string{"/kaniko/executor"},
			Args: []string{
				"--dockerfile=$(params.pathToDockerFile)",
				//"--destination=$(params.imageUrl):$(inputs.params.imageTag)",
				"--destination=$(inputs.params.imageTag):shell",
				"--context=/workspace/cicd/$(inputs.params.pathToContext)",
			},
			// TODO
			//Env: []corev1.EnvVar{
			//	{Name: "DOCKER_CONFIG", Value: "/builder/home/.docker"},
			//},
		},
	})

	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1beta1.TaskSpec{
			Resources: &v1beta1.TaskResources{
				Inputs:  inputs,
				Outputs: outputs,
			},
			Params:      params,
			Description: "源编译",
			Steps:       steps,
		},
	}

	task_meta, err = tektonClient.TektonV1beta1().Tasks(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = tektonClient.TektonV1beta1().Tasks(namespace).Create(context.Background(), task, v1.CreateOptions{})
			if err != nil {
				return res, err
			}
			return true, err
		}
		return res, err
	}
	task.ResourceVersion = task_meta.ResourceVersion
	_, err = tektonClient.TektonV1beta1().Tasks(namespace).Update(context.Background(), task, v1.UpdateOptions{})
	if err != nil {
		return false, err
	}
	return true, err

}
