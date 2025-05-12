# UserCenter OpenSource
开源用户中心 没事干，用自己觉得比较良好的架构写一个开源实现的用户中心服务

## 基础架构

- kratos, 对http/grpc服务支持友好，架构清晰明了
- mage, Go语言实现的make替代，而且可以用go语言实现各类脚本，更加高效
- viper, 配置管理，支持变量注入，便于容器部署
- wire, 依赖注入，减少人工管理依赖
- bun, 轻量级orm，支持sql合并
- ddd, 业务层使用精简DDD架构
- docker, 容器化部署

