# 状态码说明

* 0  运行依赖仅仅支持clone和Make，代码拉取和编译
* 10 运行clone、Make、buildimage，进行代码编译打包上传镜像仓库
* 20 进行代码deploy 到测试环境

## TODO 环境 dag