<template>
  <el-container class="layout-container">
    <el-aside width="200px">
      <div class="logo">医院管理系统</div>
      <el-menu
        active-text-color="#ffd04b"
        background-color="#545c64"
        text-color="#fff"
        router
      >
        <el-menu-item index="/">
          <el-icon><House /></el-icon>
          <span>首页</span>
        </el-menu-item>
        <el-menu-item index="/registration">
          <el-icon><DocumentAdd /></el-icon>
          <span>挂号管理</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header>
        <span>欢迎您，管理员</span>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<style scoped>
.layout-container {
  height: 100%;
}
.el-aside {
  background-color: #545c64;
}
.logo {
  height: 60px;
  line-height: 60px;
  text-align: center;
  color: #fff;
  font-size: 20px;
  font-weight: bold;
}
.el-header {
  display: flex;
  align-items: center;
  background-color: #fff;
  border-bottom: 1px solid #ebeef5;
}
.el-main {
  background-color: #f4f6f9;
}
</style>
