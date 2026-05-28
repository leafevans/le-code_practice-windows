**目标陈述**
- 在“新创建的项目”中，按模块逐步复刻当前后端的能力，从注册/登录开始，后续覆盖简历、面试流程、预测、消息队列、向量检索等
- 保持 Hertz + Thrift 的“IDL 驱动”开发范式与目录结构一致，便于教学与维护
- thirft教程总结：https://awq7m8b63wy.feishu.cn/docx/UkAMdxsXhoLNsCxhhPucgLBsngf


**总体原则**
- 先定义 IDL，再代码生成，再实现 handler 逻辑，最后通过中间件与仓库层串起来
- 所有路由、模型来源于 Thrift 注解；不在生成目录里手改路由（变更通过 IDL + `hz update`）
- 配置分离：关键密钥与连接串来自环境/配置文件，业务逻辑不直接硬编码


**项目初始化**
- 在新项目根目录： cd backend
  - `go mod init <你的模块名>`
  - `go install github.com/cloudwego/hertz@latest`
  - `go install github.com/apache/thrift@v0.13.0`
  - 如需固定版本：`go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0`，再 `go mod tidy`
- Windows 环境：确保 `hz` 在 PATH，安装 `go install github.com/cloudwego/hertz/cmd/hz@latest`

**IDL 规划**
- 按当前仓库的结构拆分 IDL（保持可迭代）：
  - `idl/user/user.thrift`：注册、登录、资料接口
    - 参考现有字段设计与注解映射：`backend/idl/user/user.thrift:131-162,235-247,249-261,270-281`
  - `idl/interviews/interviews.thrift`：面试记录与列表
  - `idl/mianshi/mianshi.thrift`：面试过程、会话、流式接口
  - `idl/prediction/prediction.thrift`：预测相关接口
  - 汇总入口 `idl/api.thrift`：
    - `include "./user/user.thrift"` 等（参考 `backend/idl/api.thrift:1-4`）
    - `service UserService extends user.UserService {}`（参考 `backend/idl/api.thrift:9`）
- 注解约定：
  - `api.post="/api/user/register"`、`api.post="/api/user/login"`
  - `api.form="field"`、`api.query="field"`、`api.path="field"`

**代码生成工作流**
- 首次生成：
  - `hz new -f -module interview -idl idl/api.thrift --handler_dir api/handler --model_dir api/model --router_dir api/router`
- 变更 IDL 后：
  - `hz update --idl ./idl/api.thrift --enable_extends`
- 生成结果与现有仓库映射：
  - 路由入口：`api/router/register.go`（参考现有 `backend/api/router/register.go:11-15`）
  - 用户模块路由：`api/router/user/api.go`（布局风格参考 `backend/api/router/interview/api.go:65-70`）
