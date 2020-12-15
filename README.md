# tektoncd-demo
测试tektoncd pipeline client
## 环境介绍
kubernetes v1.18.2
pipline v0.18.1


## demo 
* cicd  官方demo封装
* config k8s配置获取
* pipline client go demo
* piplineresource client go demo
* piplinerun client go demo
* task client go demo
* taskrun client go demo
* cmd和pkg 自行封装CICD流程demo 

## when 表达式
```cassandraql
const (
	DoesNotExist Operator = "!"
	Equals       Operator = "="
	DoubleEquals Operator = "=="
	In           Operator = "in"
	NotEquals    Operator = "!="
	NotIn        Operator = "notin"
	Exists       Operator = "exists"
	GreaterThan  Operator = "gt"
	LessThan     Operator = "lt"
)
```
## TODO
*  DAG 状态--> 根据when --> 判断流程走到哪里(完成)
*  pipeline 状态回调

