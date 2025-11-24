<template>
  <div>
    <el-card>
      <template #header>
        <div class="card-header">
          <span>医生管理</span>
          <el-button type="primary" @click="openDialog()" style="float: right">
            <el-icon><Plus /></el-icon> 新增医生
          </el-button>
        </div>
      </template>

      <el-form :inline="true">
        <el-form-item label="医生姓名">
          <el-input
            v-model="searchForm.name"
            placeholder="请输入医生姓名"
            clearable
          />
        </el-form-item>
        <el-form-item label="所属科室">
          <el-select
            v-model="searchForm.department"
            placeholder="请选择科室"
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
        <el-form-item>
          <el-button type="primary" @click="searchDoctors">查询</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-divider />

      <el-table :data="doctorList" border style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="姓名" width="120" />
        <el-table-column prop="title" label="职称" width="150" />
        <el-table-column prop="department_name" label="科室" width="120" />
        <el-table-column prop="available_slots" label="剩余号源" width="120" />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button
              size="small"
              type="primary"
              @click="openDialog(scope.row)"
            >
              编辑
            </el-button>
            <el-button
              size="small"
              type="danger"
              @click="handleDelete(scope.row)"
            >
              删除
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

    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? '编辑医生' : '新增医生'"
      width="500px"
    >
      <el-form :model="doctorForm" label-width="100px">
        <el-form-item label="姓名" required>
          <el-input v-model="doctorForm.name" />
        </el-form-item>
        <el-form-item label="职称" required>
          <el-select v-model="doctorForm.title" placeholder="请选择职称">
            <el-option label="主任医师" value="主任医师" />
            <el-option label="副主任医师" value="副主任医师" />
            <el-option label="主治医师" value="主治医师" />
            <el-option label="医师" value="医师" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属科室" required>
          <el-select
            v-model="doctorForm.department_code"
            placeholder="请选择科室"
          >
            <el-option
              v-for="dept in departmentList"
              :key="dept.code"
              :label="dept.name"
              :value="dept.code"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="初始号源" required>
          <el-input-number v-model="doctorForm.available_slots" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveDoctor">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import apiFetch from "../api";

const dialogVisible = ref(false);
const isEdit = ref(false);
const doctorList = ref([]);
const departmentList = ref([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);

const searchForm = reactive({
  name: "",
  department: "",
});

const doctorForm = reactive({
  id: null,
  name: "",
  title: "",
  department_code: "",
  available_slots: 20,
});

const loadDepartments = async () => {
  try {
    const data = await apiFetch("/api/departments");
    departmentList.value = Array.isArray(data) ? data : [];
  } catch (error) {
    console.error("加载科室失败:", error);
  }
};

const searchDoctors = async () => {
  try {
    const params = new URLSearchParams();
    if (searchForm.name) params.append("name", searchForm.name);
    if (searchForm.department)
      params.append("department", searchForm.department);
    params.append("page", currentPage.value);
    params.append("limit", pageSize.value);

    const data = await apiFetch(`/api/doctors?${params}`);
    if (data && Array.isArray(data.list)) {
      doctorList.value = data.list;
      total.value = data.total || 0;
    } else {
      doctorList.value = [];
      total.value = 0;
    }
  } catch (error) {
    console.error("查询医生失败:", error);
    ElMessage.error("查询医生列表失败");
  }
};

const resetSearch = () => {
  searchForm.name = "";
  searchForm.department = "";
  currentPage.value = 1;
  searchDoctors();
};

const handleSizeChange = (val) => {
  pageSize.value = val;
  searchDoctors();
};

const handleCurrentChange = (val) => {
  currentPage.value = val;
  searchDoctors();
};

const openDialog = (doctor = null) => {
  if (doctor) {
    isEdit.value = true;
    Object.assign(doctorForm, doctor);
  } else {
    isEdit.value = false;
    Object.assign(doctorForm, {
      id: null,
      name: "",
      title: "",
      department_code: "",
      available_slots: 20,
    });
  }
  dialogVisible.value = true;
};

const saveDoctor = async () => {
  if (!doctorForm.name || !doctorForm.title || !doctorForm.department_code) {
    ElMessage.warning("请填写必填项！");
    return;
  }

  try {
    const url = isEdit.value ? `/api/doctors/${doctorForm.id}` : "/api/doctors";
    const method = isEdit.value ? "PUT" : "POST";

    await apiFetch(url, {
      method,
      body: JSON.stringify(doctorForm),
    });

    ElMessage.success(isEdit.value ? "更新成功" : "添加成功");
    dialogVisible.value = false;
    searchDoctors();
  } catch (error) {
    console.error("保存医生失败:", error);
    ElMessage.error(error.message || "保存失败");
  }
};

const handleDelete = (doctor) => {
  ElMessageBox.confirm(`确定要删除医生 "${doctor.name}" 吗？`, "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      try {
        await apiFetch(`/api/doctors/${doctor.id}`, {
          method: "DELETE",
        });
        ElMessage.success("删除成功");
        searchDoctors();
      } catch (error) {
        console.error("删除失败:", error);
        ElMessage.error("删除失败");
      }
    })
    .catch(() => {});
};

onMounted(() => {
  loadDepartments();
  searchDoctors();
});
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
