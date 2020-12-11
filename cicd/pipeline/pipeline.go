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

package pipeline

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	"tektoncd-demo/config"
)

func BuildPipeline(name string, namespace string, taskName string) (res bool, err error) {

	var (
		pip *v1beta1.Pipeline
	)

	// k8s 配置文件
	restConfig, err := config.GetRestConf()
	if err != nil {
		klog.Fatal(err)
	}
	// pipline clientset
	tektonClient := versioned.NewForConfigOrDie(restConfig)
	pipline := &v1beta1.Pipeline{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1beta1.PipelineSpec{
			Resources: []v1beta1.PipelineDeclaredResource{
				{
					Name: name,
					Type: v1beta1.PipelineResourceTypeGit,
				},
			},
			Tasks: []v1beta1.PipelineTask{
				{
					Name: name,
					TaskRef: &v1beta1.TaskRef{
						Name: taskName,
					},
					Resources: &v1beta1.PipelineTaskResources{
						Inputs: []v1beta1.PipelineTaskInputResource{
							{
								Name:     name,
								Resource: name,
							},
						},
					},
				},
			},
		},
	}
	pip, err = tektonClient.TektonV1beta1().Pipelines(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = tektonClient.TektonV1beta1().Pipelines(namespace).Create(context.Background(), pipline, v1.CreateOptions{})
			if err != nil {
				return res, err
			}
			return true, err
		}
		return res, err
	}
	pipline.ResourceVersion = pip.ResourceVersion
	_, err = tektonClient.TektonV1beta1().Pipelines(namespace).Update(context.Background(), pipline, v1.UpdateOptions{})
	if err != nil {
		return false, err
	}
	return true, err
}
