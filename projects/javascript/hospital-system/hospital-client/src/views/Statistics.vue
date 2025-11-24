<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="12">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>门诊量统计（按日期）</span>
            </div>
          </template>
          <v-chart :option="registrationOption" style="height: 400px" />
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>科室挂号分布</span>
            </div>
          </template>
          <v-chart :option="departmentOption" style="height: 400px" />
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="12">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>药品库存统计（按分类）</span>
            </div>
          </template>
          <v-chart :option="medicineOption" style="height: 400px" />
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>数据概览</span>
            </div>
          </template>
          <el-statistic-group>
            <el-statistic title="总挂号数" :value="totalRegistrations">
              <template #suffix>次</template>
            </el-statistic>
            <el-statistic title="药品种类" :value="totalMedicines">
              <template #suffix>种</template>
            </el-statistic>
          </el-statistic-group>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { ElMessage } from "element-plus";
import apiFetch from "../api";

const totalRegistrations = ref(0);
const totalMedicines = ref(0);

const registrationOption = ref({
  title: { text: "挂号趋势" },
  tooltip: { trigger: "axis" },
  xAxis: {
    type: "category",
    data: [],
  },
  yAxis: { type: "value" },
  series: [
    {
      name: "挂号数量",
      type: "line",
      data: [],
      smooth: true,
      itemStyle: { color: "#409EFF" },
    },
  ],
});

const departmentOption = ref({
  title: { text: "科室分布", left: "center" },
  tooltip: { trigger: "item" },
  legend: { orient: "vertical", left: "left" },
  series: [
    {
      name: "挂号数量",
      type: "pie",
      radius: "50%",
      data: [],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: "rgba(0, 0, 0, 0.5)",
        },
      },
    },
  ],
});

const medicineOption = ref({
  title: { text: "药品库存" },
  tooltip: { trigger: "axis" },
  xAxis: {
    type: "category",
    data: [],
  },
  yAxis: { type: "value" },
  series: [
    {
      name: "库存数量",
      type: "bar",
      data: [],
      itemStyle: { color: "#67C23A" },
    },
  ],
});

const loadRegistrationStats = async () => {
  try {
    const data = await apiFetch("/api/statistics/registrations");
    registrationOption.value.xAxis.data = Array.isArray(data)
      ? data.map((item) => item.date)
      : [];
    registrationOption.value.series[0].data = Array.isArray(data)
      ? data.map((item) => item.count)
      : [];
    totalRegistrations.value = Array.isArray(data)
      ? data.reduce((sum, item) => sum + (item.count || 0), 0)
      : 0;
  } catch (error) {
    console.error("加载门诊量统计失败:", error);
    ElMessage.error("加载门诊量统计失败");
  }
};

const loadDepartmentStats = async () => {
  try {
    const data = await apiFetch("/api/statistics/departments");
    departmentOption.value.series[0].data = Array.isArray(data)
      ? data.map((item) => ({ name: item.departmentName, value: item.count }))
      : [];
  } catch (error) {
    console.error("加载科室统计失败:", error);
    ElMessage.error("加载科室统计失败");
  }
};

const loadMedicineStats = async () => {
  try {
    const data = await apiFetch("/api/statistics/medicines");
    medicineOption.value.xAxis.data = Array.isArray(data)
      ? data.map((item) => item.category)
      : [];
    medicineOption.value.series[0].data = Array.isArray(data)
      ? data.map((item) => item.totalStock)
      : [];
    totalMedicines.value = Array.isArray(data)
      ? data.reduce((sum, item) => sum + (item.count || 0), 0)
      : 0;
  } catch (error) {
    console.error("加载药品统计失败:", error);
    ElMessage.error("加载药品统计失败");
  }
};

onMounted(() => {
  loadRegistrationStats();
  loadDepartmentStats();
  loadMedicineStats();
});
</script>

<style scoped>
.card-header {
  font-weight: bold;
}
</style>
