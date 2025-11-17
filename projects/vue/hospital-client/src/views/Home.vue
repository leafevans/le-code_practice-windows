<template>
  <div class="home-container">
    <el-card class="welcome-card">
      <h1>医院门诊挂号与病历管理系统</h1>
      <p>欢迎使用本系统！</p>
    </el-card>

    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="今日挂号" :value="todayRegistrations">
            <template #prefix>
              <el-icon color="#409EFF"><DocumentAdd /></el-icon>
            </template>
            <template #suffix>次</template>
          </el-statistic>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="患者总数" :value="totalPatients">
            <template #prefix>
              <el-icon color="#67C23A"><User /></el-icon>
            </template>
            <template #suffix>人</template>
          </el-statistic>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="药品种类" :value="totalMedicines">
            <template #prefix>
              <el-icon color="#E6A23C"><Goods /></el-icon>
            </template>
            <template #suffix>种</template>
          </el-statistic>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="医生人数" :value="totalDoctors">
            <template #prefix>
              <el-icon color="#F56C6C"><Avatar /></el-icon>
            </template>
            <template #suffix>位</template>
          </el-statistic>
        </el-card>
      </el-col>
    </el-row>

    <el-card style="margin-top: 20px">
      <template #header>
        <span>快捷入口</span>
      </template>
      <el-row :gutter="20">
        <el-col :span="6">
          <el-button
            type="primary"
            @click="goTo('/registration')"
            style="width: 100%; height: 80px"
          >
            <div>
              <el-icon size="24"><DocumentAdd /></el-icon>
              <p>挂号管理</p>
            </div>
          </el-button>
        </el-col>
        <el-col :span="6">
          <el-button
            type="success"
            @click="goTo('/medical-record')"
            style="width: 100%; height: 80px"
          >
            <div>
              <el-icon size="24"><Folder /></el-icon>
              <p>病历管理</p>
            </div>
          </el-button>
        </el-col>
        <el-col :span="6">
          <el-button
            type="warning"
            @click="goTo('/medicine')"
            style="width: 100%; height: 80px"
          >
            <div>
              <el-icon size="24"><Goods /></el-icon>
              <p>药品管理</p>
            </div>
          </el-button>
        </el-col>
        <el-col :span="6">
          <el-button
            type="info"
            @click="goTo('/statistics')"
            style="width: 100%; height: 80px"
          >
            <div>
              <el-icon size="24"><DataAnalysis /></el-icon>
              <p>统计分析</p>
            </div>
          </el-button>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

const todayRegistrations = ref(0);
const totalPatients = ref(0);
const totalMedicines = ref(0);
const totalDoctors = ref(0);

const goTo = (path) => {
  router.push(path);
};

const loadStats = async () => {
  try {
    const patientsRes = await fetch("http://localhost:3000/api/patients");
    const patients = await patientsRes.json();
    totalPatients.value = patients.length;

    const medicinesRes = await fetch(
      "http://localhost:3000/api/statistics/medicines"
    );
    const medicines = await medicinesRes.json();
    totalMedicines.value = medicines.reduce((sum, item) => sum + item.count, 0);

    const doctorsRes = await fetch(
      "http://localhost:3000/api/doctors?department=internal"
    );
    totalDoctors.value = 4;

    todayRegistrations.value = 0;
  } catch (error) {
    console.error("加载统计数据失败:", error);
  }
};

onMounted(() => {
  loadStats();
});
</script>

<style scoped>
.home-container {
  padding: 20px;
}

.welcome-card {
  text-align: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.welcome-card h1 {
  color: white;
  margin: 20px 0;
}

.welcome-card p {
  color: rgba(255, 255, 255, 0.9);
  font-size: 16px;
}

.el-button div {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.el-button div p {
  margin: 0;
  font-size: 14px;
}
</style>
