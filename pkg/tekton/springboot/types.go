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
	"k8s.io/apimachinery/pkg/selection"
)

type SpringBoot struct {
	Name                 string                 `json:"name"`                   // 包名
	Url                  string                 `json:"url"`                    // 代码地址
	Revision             string                 `json:"revision"`               // 代码分支
	Status               string                 `json:"status"`                 // 状态码用于dag
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

// 发布
func (sb *SpringBoot) Deploy() (err error) {
	panic("implement me")
}

// 资源清理
func (sb *SpringBoot) Clean() (err error) {
	panic("implement me")
}

// 代码克隆
func (sb *SpringBoot) Clone() (err error, name string) {
	var (
		steps []v1beta1.Step
	)
	name = sb.Name + "-" + "clone"
	steps = append(steps, v1beta1.Step{
		Container: corev1.Container{
			Name:    "clone",
			Image:   "busybox",
			Command: []string{"echo"},
			Args: []string{
				" 代码克隆clone",
			},
		},
	})

	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: sb.NameSpace,
		},
		Spec: v1beta1.TaskSpec{
			Description: "代码clone",
			Steps:       steps,
		},
	}

	task_meta, err := sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Create(context.Background(), task, v1.CreateOptions{})
			if err != nil {
				return
			}
			return
		}
		return
	}

	task.ResourceVersion = task_meta.ResourceVersion
	_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Update(context.Background(), task, v1.UpdateOptions{})
	if err != nil {
		return
	}
	return
}

// 编译代码
func (sb *SpringBoot) Make() (err error, name string) {
	var (
		steps []v1beta1.Step
	)
	name = sb.Name + "-" + "make"
	steps = append(steps, v1beta1.Step{
		Container: corev1.Container{
			Name:    "clone",
			Image:   "busybox",
			Command: []string{"echo"},
			Args: []string{
				"Make",
			},
		},
	})

	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: sb.NameSpace,
		},
		Spec: v1beta1.TaskSpec{
			Description: "Make",
			Steps:       steps,
		},
	}

	task_meta, err := sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Create(context.Background(), task, v1.CreateOptions{})
			if err != nil {
				return err, name
			}
			return err, name
		}
		return err, name
	}

	task.ResourceVersion = task_meta.ResourceVersion
	_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Update(context.Background(), task, v1.UpdateOptions{})
	if err != nil {
		return err, name
	}
	return
}

// 打包镜像
func (sb *SpringBoot) BuildImage() (err error, name string) {
	var (
		steps []v1beta1.Step
	)
	name = sb.Name + "-" + "buildimage"
	steps = append(steps, v1beta1.Step{
		Container: corev1.Container{
			Name:    "clone",
			Image:   "busybox",
			Command: []string{"echo"},
			Args: []string{
				"Make",
			},
		},
	})

	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: sb.NameSpace,
		},
		Spec: v1beta1.TaskSpec{
			Description: "Make",
			Steps:       steps,
		},
	}

	task_meta, err := sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Create(context.Background(), task, v1.CreateOptions{})
			if err != nil {
				return
			}
			return
		}
		return
	}

	task.ResourceVersion = task_meta.ResourceVersion
	_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Update(context.Background(), task, v1.UpdateOptions{})
	if err != nil {
		return
	}
	return
}

// 消息通知
func (sb *SpringBoot) Notice() (err error, name string) {
	var (
		steps []v1beta1.Step
	)
	name = sb.Name + "-" + "notice"
	steps = append(steps, v1beta1.Step{
		Container: corev1.Container{
			Name:    "clone",
			Image:   "busybox",
			Command: []string{"echo"},
			Args: []string{
				"notice",
			},
		},
	})

	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: sb.NameSpace,
		},
		Spec: v1beta1.TaskSpec{
			Description: "notice",
			Steps:       steps,
		},
	}

	task_meta, err := sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Create(context.Background(), task, v1.CreateOptions{})
			if err != nil {
				return
			}
			return
		}
		return
	}

	task.ResourceVersion = task_meta.ResourceVersion
	_, err = sb.TektonClient.TektonV1beta1().Tasks(sb.NameSpace).Update(context.Background(), task, v1.UpdateOptions{})
	if err != nil {
		return
	}
	return

	// 项目阶段消息发送
}

// 运行整个流程
func (sb *SpringBoot) Run() (err error) {
	var (
		piperun *v1beta1.PipelineRun
	)

	err, clone_name := sb.Clone()
	if err != nil {
	}

	err, make_name := sb.Make()
	if err != nil {
	}

	err, buildimage_name := sb.BuildImage()
	if err != nil {
	}

	err, notice := sb.Notice()
	if err != nil {
	}

	piplinerun := &v1beta1.PipelineRun{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      sb.Name,
			Namespace: sb.NameSpace,
		},
		Spec: v1beta1.PipelineRunSpec{
			PipelineSpec: &v1beta1.PipelineSpec{
				Description: "pipline " + sb.Name,
				Tasks: []v1beta1.PipelineTask{
					{
						Name: clone_name,
						TaskRef: &v1beta1.TaskRef{
							Name: clone_name,
						},
						WhenExpressions: []v1beta1.WhenExpression{
							{
								Input:    sb.Status,
								Operator: selection.In,
								Values:   []string{"0", "10", "20"},
							},
						},
					},
					{
						Name: make_name,
						TaskRef: &v1beta1.TaskRef{
							Name: make_name,
						},
						WhenExpressions: []v1beta1.WhenExpression{
							{
								Input:    sb.Status,
								Operator: selection.In,
								Values:   []string{"0", "10", "20"},
							},
						},
						RunAfter: []string{clone_name},
					},
					{
						Name: buildimage_name,
						TaskRef: &v1beta1.TaskRef{
							Name: buildimage_name,
						},
						WhenExpressions: []v1beta1.WhenExpression{
							{
								Input:    sb.Status,
								Operator: selection.In,
								Values:   []string{"10", "20"},
							},
						},
						RunAfter: []string{make_name},
					},
				},
				Finally: []v1beta1.PipelineTask{

					{
						Name: notice,
						TaskRef: &v1beta1.TaskRef{
							Name: notice,
						},
					},
				},
			},
		},
	}

	//// 删除run
	sb.TektonClient.TektonV1beta1().PipelineRuns(sb.NameSpace).Delete(context.Background(), sb.Name, v1.DeleteOptions{})
	for {
		_, err = sb.TektonClient.TektonV1beta1().PipelineRuns(sb.NameSpace).Get(context.Background(), sb.Name, v1.GetOptions{})
		if errors.IsNotFound(err) {
			break
		}
	}
	if piperun, err = sb.TektonClient.TektonV1beta1().PipelineRuns(sb.NameSpace).Get(context.Background(), sb.Name, v1.GetOptions{}); err != nil {
		if errors.IsNotFound(err) {
			_, err = sb.TektonClient.TektonV1beta1().PipelineRuns(sb.NameSpace).Create(context.Background(), piplinerun, v1.CreateOptions{})
			if err != nil {
				return
			}
		}
		return
	} else {
		piplinerun.ResourceVersion = piperun.ResourceVersion
		_, err = sb.TektonClient.TektonV1beta1().PipelineRuns(sb.NameSpace).Update(context.Background(), piplinerun, v1.UpdateOptions{})
		if err != nil {
			return
		}
	}
	return
}
