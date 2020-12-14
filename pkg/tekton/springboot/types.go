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

package springboot

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	tknResource "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SpringBoot struct {
	Name                 string                 `json:"name"`                   // 包名
	Url                  string                 `json:"url"`                    // 代码地址
	Revision             string                 `json:"revision"`               // 代码分支
	TektonClient         *versioned.Clientset   `json:"tekton_client"`          // tekton client
	TektonClientResource *tknResource.Clientset `json:"tekton_client_resource"` // tekton resource client
	NameSpace            string                 `json:"namespace"`              // 运行环境

}

func (sb *SpringBoot) Cache() (err error) {
	panic("implement me")
}

func (sb *SpringBoot) Scan() (err error) {
	panic("implement me")
}

func (sb *SpringBoot) UnitTest() (err error) {
	panic("implement me")
}

func (sb *SpringBoot) CodeQuqlity() (err error) {
	panic("implement me")
}

func (sb *SpringBoot) Deploy() (err error) {
	panic("implement me")
}

func (sb *SpringBoot) Clean() (err error) {
	panic("implement me")
}

func (sb *SpringBoot) Clone() (err error) {
	var (
		steps []v1beta1.Step
	)
	sb.Name = sb.Name + "-" + "clone"
	steps = append(steps, v1beta1.Step{
		Container: corev1.Container{
			Name:    "clone",
			Image:   "busybox",
			Command: []string{"/usr/bin/echo"},
			Args: []string{
				" 代码克隆clone",
			},
		},
	})

	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      sb.Name,
			Namespace: sb.NameSpace,
		},
		Spec: v1beta1.TaskSpec{
			Description: "代码clone",
			Steps:       steps,
		},
	}

	task_meta, err := sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Get(context.Background(), sb.Name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Create(context.Background(), task, v1.CreateOptions{})
			if err != nil {
				return err
			}
			return err
		}
		return err
	}

	task.ResourceVersion = task_meta.ResourceVersion
	_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Update(context.Background(), task, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	return
}

func (sb *SpringBoot) Make() (err error) {
	var (
		steps []v1beta1.Step
	)
	sb.Name = sb.Name + "-" + "make"
	steps = append(steps, v1beta1.Step{
		Container: corev1.Container{
			Name:    "clone",
			Image:   "busybox",
			Command: []string{"/usr/bin/echo"},
			Args: []string{
				"Make",
			},
		},
	})

	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      sb.Name,
			Namespace: sb.NameSpace,
		},
		Spec: v1beta1.TaskSpec{
			Description: "Make",
			Steps:       steps,
		},
	}

	task_meta, err := sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Get(context.Background(), sb.Name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Create(context.Background(), task, v1.CreateOptions{})
			if err != nil {
				return err
			}
			return err
		}
		return err
	}

	task.ResourceVersion = task_meta.ResourceVersion
	_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Update(context.Background(), task, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	return
}

func (sb *SpringBoot) BuildImage() (err error) {
	var (
		steps []v1beta1.Step
	)
	sb.Name = sb.Name + "-" + "buildimage"
	steps = append(steps, v1beta1.Step{
		Container: corev1.Container{
			Name:    "clone",
			Image:   "busybox",
			Command: []string{"/usr/bin/echo"},
			Args: []string{
				"Make",
			},
		},
	})

	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      sb.Name,
			Namespace: sb.NameSpace,
		},
		Spec: v1beta1.TaskSpec{
			Description: "Make",
			Steps:       steps,
		},
	}

	task_meta, err := sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Get(context.Background(), sb.Name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Create(context.Background(), task, v1.CreateOptions{})
			if err != nil {
				return err
			}
			return err
		}
		return err
	}

	task.ResourceVersion = task_meta.ResourceVersion
	_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Update(context.Background(), task, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	return
}

func (sb *SpringBoot) Notice() (err error) {
	var (
		steps []v1beta1.Step
	)
	sb.Name = sb.Name + "-" + "notice"
	steps = append(steps, v1beta1.Step{
		Container: corev1.Container{
			Name:    "clone",
			Image:   "busybox",
			Command: []string{"/usr/bin/echo"},
			Args: []string{
				"notice",
			},
		},
	})

	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      sb.Name,
			Namespace: sb.NameSpace,
		},
		Spec: v1beta1.TaskSpec{
			Description: "notice",
			Steps:       steps,
		},
	}

	task_meta, err := sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Get(context.Background(), sb.Name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Create(context.Background(), task, v1.CreateOptions{})
			if err != nil {
				return err
			}
			return err
		}
		return err
	}

	task.ResourceVersion = task_meta.ResourceVersion
	_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Update(context.Background(), task, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	return

	// 项目阶段消息发送
}

func (sb *SpringBoot) Run() (err error) {
	if sb.Clone() == nil {
	}
	if sb.Make() == nil {
	}
	if sb.BuildImage() == nil {
	}
	if sb.Notice() == nil {
	}

	return
}
