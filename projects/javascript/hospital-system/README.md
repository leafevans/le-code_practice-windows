# 医院门诊挂号与病历管理系统

一个基于 Vue 3 + Node.js + SQLite 的医院管理系统，支持挂号管理、病历管理、药品管理和统计分析等功能。

## 📋 项目简介

本系统是一个功能完整的医院门诊管理系统，采用前后端分离架构，提供直观的 Web 界面，支持一键启动和打包成独立可执行文件。

### 主要功能

- ✅ **系统登录** - 多角色登录（管理员/医生）、路由权限控制
- ✅ **挂号管理** - 科室选择、医生排班查询、患者挂号
- ✅ **医生工作台** - 医生接诊、诊断录入、处方开具
- ✅ **收费管理** - 费用计算、处方缴费、支付状态管理
- ✅ **病历管理** - 患者档案创建、历史病历查询
- ✅ **药品管理** - 药品信息维护、库存管理
- ✅ **统计分析** - 门诊量/科室/药品数据可视化

## 🛠️ 技术栈

### 前端

- Vue 3 - 渐进式 JavaScript 框架
- Element Plus - UI 组件库
- Vue Router - 路由管理
- ECharts - 数据可视化图表库
- Vite - 构建工具

### 后端

- Node.js - JavaScript 运行环境
- Express - Web 应用框架
- SQLite3 - 轻量级数据库
- CORS - 跨域资源共享

## 📁 项目结构

```
vue/
├── hospital-client/          # 前端项目
│   ├── src/
│   │   ├── components/      # 公共组件
│   │   ├── views/           # 页面组件
│   │   │   ├── Login.vue          # 登录页
│   │   │   ├── Home.vue           # 首页
│   │   │   ├── Registration.vue   # 挂号管理
│   │   │   ├── DoctorWorkbench.vue # 医生工作台
│   │   │   ├── Charge.vue         # 收费管理
│   │   │   ├── MedicalRecord.vue  # 病历管理
│   │   │   ├── Medicine.vue       # 药品管理
│   │   │   └── Statistics.vue     # 统计分析
│   │   ├── layout/          # 布局组件
│   │   ├── router/          # 路由配置
│   │   ├── App.vue          # 根组件
│   │   └── main.js          # 入口文件
│   ├── dist/                # 构建输出目录
│   ├── package.json
│   └── vite.config.js
│
├── hospital-server/         # 后端项目
│   ├── server.js           # 服务器主文件
│   ├── database.js         # 数据库配置
│   ├── hospital.db         # SQLite 数据库文件
│   └── package.json
│
├── 启动医院管理系统.bat    # 一键启动脚本
└── README.md
```

## 🚀 快速开始

### 环境要求

- Node.js >= 18.0.0
- npm >= 9.0.0

### 开发模式运行

#### 1. 克隆项目

```bash
git clone <repository-url>
cd vue
```

#### 2. 安装依赖

```bash
# 安装前端依赖
cd hospital-client
npm install

# 安装后端依赖
cd ../hospital-server
npm install
```

#### 3. 启动后端服务

```bash
cd hospital-server
node server.js
```

服务器将在 `http://localhost:3000` 启动

#### 4. 启动前端开发服务器

```bash
cd hospital-client
npm run dev
```

前端开发服务器将在 `http://localhost:5173` 启动

### 生产模式部署

#### 方法一：使用批处理文件（推荐）

1. **构建前端**

   ```bash
   cd hospital-client
   npm run build
   ```

2. **双击 `启动医院管理系统.bat`** 即可运行

#### 方法二：打包成独立 EXE

1. **安装打包工具**

   ```bash
   npm install -g pkg
   ```

2. **构建前端**

   ```bash
   cd hospital-client
   npm run build
   ```

3. **打包后端**

   ```bash
   cd hospital-server
   pkg . --targets node18-win-x64 --output hospital-system.exe
   ```

4. **整理发布文件**

   ```
   医院管理系统/
   ├── hospital-system.exe
   └── hospital-client/
       └── dist/
   ```

5. **运行** - 双击 `hospital-system.exe`

## 📊 数据库结构

系统使用 SQLite 数据库，包含以下数据表：

- `departments` - 科室表
- `doctors` - 医生表
- `patients` - 患者表
- `registrations` - 挂号记录表
- `medicines` - 药品表
- `prescriptions` - 处方表

### 初始测试数据

系统首次运行时会自动创建测试数据：

**科室**

- 内科 (internal)
- 外科 (surgical)
- 儿科 (pediatrics)

**医生**

- 张三丰 - 内科主任医师
- 李时珍 - 内科副主任医师
- 华佗 - 外科主任医师
- 扁鹊 - 儿科主治医师

**药品**

- 阿莫西林胶囊
- 布洛芬片
- 感冒灵颗粒
- 维生素 C 片
- 板蓝根颗粒

## 🎯 使用指南

### 1. 系统登录

- **管理员账号**: `admin` / `123456`
- **医生账号**: `doctor` / `123456`

### 2. 挂号流程（管理员）

1. 进入「挂号管理」页面
2. 选择就诊科室和日期
3. 点击「查询医生」查看医生排班
4. 点击「挂号」按钮，输入患者 ID（需先在病历管理中创建患者档案）
5. 确认挂号，生成待就诊记录

### 3. 医生接诊（医生）

1. 登录医生账号，进入「医生工作台」
2. 在左侧列表中选择「待就诊」患者
3. 填写诊断信息（主诉、诊断结果）
4. 添加处方药品（自动扣减库存）
5. 点击「完成接诊」，系统生成待缴费记录

### 4. 收费结算（管理员）

1. 进入「收费管理」页面
2. 查看待缴费患者列表
3. 点击「收费结算」查看费用明细（挂号费 + 药品费）
4. 确认收款后，完成整个就诊流程

### 5. 病历管理

1. 进入「病历管理」页面
2. 点击「新建患者档案」
3. 填写患者基本信息（姓名、性别、身份证号、联系电话等）
4. 保存后记录患者 ID，用于挂号
5. 点击「查看挂号记录」可查看患者的所有就诊记录

### 6. 药品管理

1. 进入「药品管理」页面
2. 查看当前药品库存
3. 点击「添加药品」新增药品信息
4. 点击「调整库存」修改药品库存数量

### 7. 统计分析

1. 进入「统计分析」页面
2. 查看门诊量趋势图表
3. 查看科室挂号分布饼图
4. 查看药品库存柱状图
5. 查看系统数据概览

## 🔧 配置说明

### 修改端口

在 [`hospital-server/server.js`](hospital-server/server.js) 中修改：

```javascript
const PORT = 3000; // 修改为其他端口
```

### 修改数据库路径

在 [`hospital-server/database.js`](hospital-server/database.js) 中修改：

```javascript
const dbPath = path.join(__dirname, "hospital.db");
```

### 修改挂号费用

在数据库表 `registrations` 中，`fee` 字段默认值为 10.0 元

## ⚠️ 注意事项

1. **端口占用** - 确保 3000 端口未被其他程序占用
2. **数据备份** - 定期备份 `hospital.db` 数据库文件
3. **浏览器兼容** - 推荐使用 Chrome、Edge 等现代浏览器
4. **数据安全** - 本系统为演示项目，实际使用需加强安全措施

## 🐛 常见问题

### Q: 启动后浏览器无法访问？

A: 请检查：

1. 后端服务是否正常启动
2. 防火墙是否阻止了 3000 端口
3. 等待 3-5 秒让服务完全启动后再访问

### Q: 数据库连接失败？

A: 请检查：

1. `hospital.db` 文件是否有读写权限
2. 数据库文件路径是否正确
3. 是否有多个实例同时访问数据库

### Q: 挂号时提示号源不足？

A: 医生初始号源为 10 个，挂满后需要：

1. 重启系统重置号源
2. 或手动修改数据库中医生的 `available_slots` 字段

### Q: 无法登录系统？

A: 请使用默认测试账号：

- 管理员：admin / 123456
- 医生：doctor / 123456

### Q: 打包后的 exe 无法运行？

A: 请检查：

1. 是否将 `hospital-client/dist` 文件夹放在 exe 同级目录
2. 是否安装了 Visual C++ 运行库
3. 查看命令行窗口的错误信息
