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

package pipelinerun

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/klog"
	"tektoncd-demo/config"
)

func Run(name string, namespace string, pipelineName string) (res bool, err error) {
	var (
		piperun *v1beta1.PipelineRun
	)

	// k8s 配置文件
	restConfig, err := config.GetRestConf()
	if err != nil {
		klog.Fatal(err)
	}
	// pipline clientset
	tektonClient := versioned.NewForConfigOrDie(restConfig)
	piplinerun := &v1beta1.PipelineRun{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1beta1.PipelineRunSpec{
			Resources: []v1beta1.PipelineResourceBinding{
				{
					Name: name,
					ResourceRef: &v1beta1.PipelineResourceRef{
						Name: name,
					},
				},
			},
			PipelineRef: &v1beta1.PipelineRef{
				Name: pipelineName,
			},
		},
	}
	// 判断是否存在
	piperun, err = tektonClient.TektonV1beta1().PipelineRuns(namespace).Get(context.Background(), name, v1.GetOptions{})
	klog.Info(err)
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = tektonClient.TektonV1beta1().PipelineRuns(namespace).Create(context.Background(), piplinerun, v1.CreateOptions{})
			if err != nil {
				return res, err
			}
			return true, err
		}
		return res, err
	}
	piplinerun.ResourceVersion = piperun.ResourceVersion
	_, err = tektonClient.TektonV1beta1().PipelineRuns(namespace).Update(context.Background(), piplinerun, v1.UpdateOptions{})
	if err != nil {
		return res, err
	}

	return true, err
}
