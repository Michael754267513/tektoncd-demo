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

package tekton

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TaskRun(tektonClient *versioned.Clientset, name string, namespace string) (err error) {
	taskRun := &v1beta1.TaskRun{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1beta1.TaskRunSpec{
			TaskRef: &v1beta1.TaskRef{
				Name: name,
			},
		},
	}
	if _, err := tektonClient.TektonV1beta1().TaskRuns(namespace).Get(context.Background(), name, v1.GetOptions{}); err != nil {
		if errors.IsNotFound(err) {
			_, err := tektonClient.TektonV1beta1().TaskRuns(namespace).Create(context.Background(), taskRun, v1.CreateOptions{})
			if err != nil {
				return err
			}
		}
	}

	_, err = tektonClient.TektonV1beta1().TaskRuns(namespace).Update(context.Background(), taskRun, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	return
}
