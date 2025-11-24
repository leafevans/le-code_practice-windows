<template>
  <el-container class="layout-container">
    <el-aside width="200px">
      <div class="logo">医院管理系统</div>
      <el-menu
        :default-active="currentRoute"
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
        <el-menu-item index="/registration-list">
          <el-icon><List /></el-icon>
          <span>挂号记录</span>
        </el-menu-item>
        <el-menu-item index="/doctor-workbench">
          <el-icon><Monitor /></el-icon>
          <span>医生工作台</span>
        </el-menu-item>
        <el-menu-item index="/charge">
          <el-icon><Money /></el-icon>
          <span>收费管理</span>
        </el-menu-item>
        <el-menu-item index="/medical-record">
          <el-icon><Folder /></el-icon>
          <span>病历管理</span>
        </el-menu-item>
        <el-menu-item index="/doctor">
          <el-icon><Avatar /></el-icon>
          <span>医生管理</span>
        </el-menu-item>
        <el-menu-item index="/department">
          <el-icon><OfficeBuilding /></el-icon>
          <span>科室管理</span>
        </el-menu-item>
        <el-menu-item index="/medicine">
          <el-icon><Goods /></el-icon>
          <span>药品管理</span>
        </el-menu-item>
        <el-menu-item index="/statistics">
          <el-icon><DataAnalysis /></el-icon>
          <span>统计分析</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header>
        <div class="header-content">
          <span>欢迎您，{{ userName }}</span>
          <el-button type="danger" link @click="handleLogout"
            >退出登录</el-button
          >
        </div>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed, ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";

const route = useRoute();
const router = useRouter();
const userName = ref("管理员");

const currentRoute = computed(() => route.path);

onMounted(() => {
  const userStr = localStorage.getItem("user");
  if (userStr) {
    try {
      const user = JSON.parse(userStr);
      userName.value = user.name || "管理员";
    } catch (e) {
      console.error("解析用户信息失败", e);
    }
  }
});

const handleLogout = () => {
  localStorage.removeItem("token");
  localStorage.removeItem("user");
  ElMessage.success("已退出登录");
  router.push("/login");
};
</script>

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
  background-color: #fff;
  border-bottom: 1px solid #ebeef5;
  padding: 0 20px;
}
.header-content {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
}
.el-main {
  background-color: #f4f6f9;
}
</style>
