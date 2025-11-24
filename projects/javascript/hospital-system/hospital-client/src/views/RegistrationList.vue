<template>
  <div>
    <el-card>
      <template #header>
        <div class="card-header">
          <span>挂号记录管理</span>
        </div>
      </template>

      <el-form :inline="true" :model="searchForm">
        <el-form-item label="就诊日期">
          <el-date-picker
            v-model="searchForm.date"
            type="date"
            placeholder="选择日期"
            value-format="YYYY-MM-DD"
            style="width: 150px"
          />
        </el-form-item>
        <el-form-item label="科室">
          <el-select
            v-model="searchForm.department"
            placeholder="全部科室"
            clearable
            style="width: 150px"
          >
            <el-option
              v-for="dept in departmentList"
              :key="dept.code"
              :label="dept.name"
              :value="dept.code"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select
            v-model="searchForm.status"
            placeholder="全部状态"
            clearable
            style="width: 120px"
          >
            <el-option label="待就诊" value="pending" />
            <el-option label="已完成" value="completed" />
            <el-option label="已取消" value="cancelled" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="searchRegistrations"
            >查询</el-button
          >
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-divider />

      <el-table :data="registrationList" border style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="patient_name" label="患者姓名" width="120" />
        <el-table-column prop="patient_gender" label="性别" width="80" />
        <el-table-column prop="department_name" label="科室" width="120" />
        <el-table-column prop="doctor_name" label="医生" width="120" />
        <el-table-column
          prop="registration_date"
          label="就诊日期"
          width="120"
        />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button
              v-if="scope.row.status === 'pending'"
              size="small"
              type="success"
              @click="updateStatus(scope.row, 'completed')"
            >
              完成就诊
            </el-button>
            <el-button
              v-if="scope.row.status === 'pending'"
              size="small"
              type="danger"
              @click="updateStatus(scope.row, 'cancelled')"
            >
              取消挂号
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div style="margin-top: 20px; display: flex; justify-content: flex-end">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import apiFetch from "../api";

const registrationList = ref([]);
const departmentList = ref([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);

const searchForm = reactive({
  date: "",
  department: "",
  status: "",
});

const getStatusType = (status) => {
  const map = {
    pending: "warning",
    completed: "success",
    cancelled: "info",
  };
  return map[status] || "info";
};

const getStatusText = (status) => {
  const map = {
    pending: "待就诊",
    completed: "已完成",
    cancelled: "已取消",
  };
  return map[status] || status;
};

const loadDepartments = async () => {
  try {
    const data = await apiFetch("/api/departments");
    departmentList.value = Array.isArray(data) ? data : [];
  } catch (error) {
    console.error("加载科室失败:", error);
  }
};

const searchRegistrations = async () => {
  try {
    const params = new URLSearchParams();
    if (searchForm.date) params.append("date", searchForm.date);
    if (searchForm.department)
      params.append("department", searchForm.department);
    if (searchForm.status) params.append("status", searchForm.status);
    params.append("page", currentPage.value);
    params.append("limit", pageSize.value);

    const data = await apiFetch(`/api/registrations?${params}`);
    if (data && Array.isArray(data.list)) {
      registrationList.value = data.list;
      total.value = data.total || 0;
    } else {
      registrationList.value = [];
      total.value = 0;
    }
  } catch (error) {
    console.error("查询挂号记录失败:", error);
    ElMessage.error("查询挂号记录失败");
  }
};

const resetSearch = () => {
  searchForm.date = "";
  searchForm.department = "";
  searchForm.status = "";
  currentPage.value = 1;
  searchRegistrations();
};

const handleSizeChange = (val) => {
  pageSize.value = val;
  searchRegistrations();
};

const handleCurrentChange = (val) => {
  currentPage.value = val;
  searchRegistrations();
};

const updateStatus = (row, status) => {
  const actionText = status === "completed" ? "完成就诊" : "取消挂号";
  ElMessageBox.confirm(
    `确定要将患者 "${row.patient_name}" 的挂号状态更新为 "${actionText}" 吗？`,
    "提示",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: status === "cancelled" ? "warning" : "primary",
    }
  )
    .then(async () => {
      try {
        await apiFetch(`/api/registrations/${row.id}/status`, {
          method: "PUT",
          body: JSON.stringify({ status }),
        });
        ElMessage.success("状态更新成功");
        searchRegistrations();
      } catch (error) {
        console.error("更新状态失败:", error);
        ElMessage.error(error.message || "更新状态失败");
      }
    })
    .catch(() => {});
};

onMounted(() => {
  loadDepartments();
  searchRegistrations();
});
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
