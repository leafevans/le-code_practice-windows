<template>
  <div>
    <el-card>
      <template #header>
        <div class="card-header">
          <span>科室管理</span>
          <el-button type="primary" @click="openDialog()" style="float: right">
            <el-icon><Plus /></el-icon> 新增科室
          </el-button>
        </div>
      </template>

      <el-table :data="departmentList" border style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="科室名称" />
        <el-table-column prop="code" label="科室代码" />
        <el-table-column label="操作" width="200">
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
    </el-card>

    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? '编辑科室' : '新增科室'"
      width="500px"
    >
      <el-form :model="deptForm" label-width="100px">
        <el-form-item label="科室名称" required>
          <el-input v-model="deptForm.name" placeholder="例如：内科" />
        </el-form-item>
        <el-form-item label="科室代码" required>
          <el-input
            v-model="deptForm.code"
            placeholder="例如：internal"
            :disabled="isEdit"
          />
          <span
            v-if="isEdit"
            style="font-size: 12px; color: #999; margin-left: 10px"
            >代码不可修改</span
          >
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveDepartment">保存</el-button>
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
const departmentList = ref([]);

const deptForm = reactive({
  id: null,
  name: "",
  code: "",
});

const loadDepartments = async () => {
  try {
    const data = await apiFetch("/api/departments");
    departmentList.value = Array.isArray(data) ? data : [];
  } catch (error) {
    console.error("加载科室失败:", error);
    ElMessage.error("加载科室列表失败");
  }
};

const openDialog = (dept = null) => {
  if (dept) {
    isEdit.value = true;
    Object.assign(deptForm, dept);
  } else {
    isEdit.value = false;
    Object.assign(deptForm, {
      id: null,
      name: "",
      code: "",
    });
  }
  dialogVisible.value = true;
};

const saveDepartment = async () => {
  if (!deptForm.name || !deptForm.code) {
    ElMessage.warning("请填写必填项！");
    return;
  }

  try {
    const url = isEdit.value
      ? `/api/departments/${deptForm.id}`
      : "/api/departments";
    const method = isEdit.value ? "PUT" : "POST";

    await apiFetch(url, {
      method,
      body: JSON.stringify(deptForm),
    });

    ElMessage.success(isEdit.value ? "更新成功" : "添加成功");
    dialogVisible.value = false;
    loadDepartments();
  } catch (error) {
    console.error("保存科室失败:", error);
    ElMessage.error(error.message || "保存失败");
  }
};

const handleDelete = (dept) => {
  ElMessageBox.confirm(`确定要删除科室 "${dept.name}" 吗？`, "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      try {
        await apiFetch(`/api/departments/${dept.id}`, {
          method: "DELETE",
        });
        ElMessage.success("删除成功");
        loadDepartments();
      } catch (error) {
        console.error("删除失败:", error);
        ElMessage.error(error.message || "删除失败");
      }
    })
    .catch(() => {});
};

onMounted(() => {
  loadDepartments();
});
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
