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

package resource

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"tektoncd-demo/config"
)

// 资源定义
func Resource(name string, namespace string, resourceType v1alpha1.PipelineResourceType, params []v1alpha1.ResourceParam) (res bool, err error) {
	var (
		pipelineresource *v1alpha1.PipelineResource
	)

	// k8s 配置文件
	restConfig, err := config.GetRestConf()
	if err != nil {
		return res, err
	}
	// pipeline clientset
	tektonClient := versioned.NewForConfigOrDie(restConfig)
	pipresource := &v1alpha1.PipelineResource{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.PipelineResourceSpec{
			Type:   resourceType,
			Params: params,
		},
		Status: nil,
	}
	// 判断是否存在resource
	pipelineresource, err = tektonClient.TektonV1alpha1().PipelineResources(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = tektonClient.TektonV1alpha1().PipelineResources(namespace).Create(context.Background(), pipresource, v1.CreateOptions{})
			if err != nil {
				return res, err
			}
			return true, err
		}
		return res, err
	}
	pipresource.ResourceVersion = pipelineresource.ResourceVersion
	_, err = tektonClient.TektonV1alpha1().PipelineResources(namespace).Update(context.Background(), pipresource, v1.UpdateOptions{})
	if err != nil {
		return res, err
	}

	return true, err
}
