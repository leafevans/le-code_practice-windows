<template>
  <div>
    <el-card>
      <template #header>
        <div class="card-header">
          <span>门诊挂号</span>
        </div>
      </template>

      <el-form :inline="true" :model="formState">
        <el-form-item label="选择科室">
          <el-select
            v-model="formState.department"
            placeholder="请选择科室"
            style="width: 150px"
          >
            <el-option label="内科" value="internal" />
            <el-option label="外科" value="surgical" />
            <el-option label="儿科" value="pediatrics" />
          </el-select>
        </el-form-item>
        <el-form-item label="就诊日期">
          <el-date-picker
            v-model="formState.date"
            type="date"
            placeholder="选择日期"
            style="width: 150px"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="queryDoctors"
            >查询医生排班</el-button
          >
        </el-form-item>
      </el-form>

      <el-divider />

      <p><b>医生排班信息</b> (请先选择科室和日期后查询)</p>
      <el-table
        :data="doctorSchedule"
        border
        style="width: 100%"
        empty-text="暂无排班信息"
      >
        <el-table-column prop="doctorName" label="医生姓名" width="180" />
        <el-table-column prop="title" label="职称" width="180" />
        <el-table-column prop="availableSlots" label="剩余号源" />
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button
              size="small"
              type="primary"
              @click="handleRegister(scope.row)"
            >
              挂号
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { reactive, ref } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";

const formState = reactive({
  department: "",
  date: "",
});

const doctorSchedule = ref([]);

const queryDoctors = async () => {
  if (!formState.department || !formState.date) {
    ElMessage.warning("请先选择科室和就诊日期！");
    return;
  }

  try {
    const response = await fetch(
      `http://localhost:3000/api/doctors?department=${formState.department}`
    );
    if (!response.ok) {
      throw new Error("网络请求失败");
    }
    const data = await response.json();
    doctorSchedule.value = data;

    if (doctorSchedule.value.length === 0) {
      ElMessage.info("该科室当天暂无排班");
    }
  } catch (error) {
    console.error("获取医生数据失败:", error);
    ElMessage.error("获取医生排班信息失败，请检查后端服务是否开启。");
  }
};

const handleRegister = (doctor) => {
  ElMessageBox.prompt("请输入患者ID（在病历管理中创建）", "挂号确认", {
    confirmButtonText: "确认挂号",
    cancelButtonText: "取消",
    inputPattern: /^\d+$/,
    inputErrorMessage: "请输入有效的患者ID",
  })
    .then(async ({ value }) => {
      try {
        const response = await fetch(
          "http://localhost:3000/api/registrations",
          {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              patient_id: parseInt(value),
              doctor_id: doctor.id,
              department_code: formState.department,
              registration_date: formState.date,
            }),
          }
        );

        if (response.ok) {
          // 更新本地显示的号源数量
          doctor.availableSlots--;
          ElMessage.success("挂号成功！");
        } else {
          const error = await response.json();
          ElMessage.error(error.error || "挂号失败");
        }
      } catch (error) {
        console.error("挂号失败:", error);
        ElMessage.error("挂号失败，请稍后重试");
      }
    })
    .catch(() => {
      ElMessage.info("已取消挂号");
    });
};
</script>

<style scoped>
.card-header {
  font-weight: bold;
}
.el-form {
  margin-bottom: 20px;
}
</style>
