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

type CICD interface {
	Clone() (err error, name string)      // 代码克隆
	Make() (err error, name string)       // 代码编译
	BuildImage() (err error, name string) // 编译镜像
	Cache() (err error)                   // 缓存cache

	Scan() (err error)        // 代码扫描
	UnitTest() (err error)    // 单元测试
	CodeQuqlity() (err error) // 代码qos

	Deploy() (err error) // 部署

	Notice() (err error) // 消息通知
	Clean() (err error)  // 资源清理
	Run() (err error)    // 运行pipline
}
