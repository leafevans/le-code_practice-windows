<template>
  <div>
    <el-card>
      <template #header>
        <div class="card-header">
          <span>病历管理</span>
          <el-button
            type="primary"
            @click="dialogVisible = true"
            style="float: right"
          >
            <el-icon><Plus /></el-icon> 新建患者档案
          </el-button>
        </div>
      </template>

      <el-form :inline="true">
        <el-form-item label="患者姓名">
          <el-input
            v-model="searchForm.name"
            placeholder="请输入患者姓名"
            clearable
          />
        </el-form-item>
        <el-form-item label="身份证号">
          <el-input
            v-model="searchForm.idCard"
            placeholder="请输入身份证号"
            clearable
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="searchPatients">查询</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-divider />

      <el-table :data="patientList" border style="width: 100%">
        <el-table-column prop="id" label="患者ID" width="80" />
        <el-table-column prop="name" label="姓名" width="120" />
        <el-table-column prop="gender" label="性别" width="80" />
        <el-table-column prop="id_card" label="身份证号" width="180" />
        <el-table-column prop="phone" label="联系电话" width="140" />
        <el-table-column prop="birth_date" label="出生日期" width="120" />
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button
              size="small"
              type="primary"
              @click="viewRecords(scope.row)"
            >
              查看挂号记录
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div style="margin-top: 20px; display: flex; justify-content: flex-end">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <el-dialog v-model="dialogVisible" title="新建患者档案" width="500px">
      <el-form :model="patientForm" label-width="100px">
        <el-form-item label="姓名" required>
          <el-input v-model="patientForm.name" />
        </el-form-item>
        <el-form-item label="性别" required>
          <el-radio-group v-model="patientForm.gender">
            <el-radio label="男">男</el-radio>
            <el-radio label="女">女</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="身份证号" required>
          <el-input v-model="patientForm.id_card" />
        </el-form-item>
        <el-form-item label="联系电话" required>
          <el-input v-model="patientForm.phone" />
        </el-form-item>
        <el-form-item label="出生日期">
          <el-date-picker
            v-model="patientForm.birth_date"
            type="date"
            placeholder="选择日期"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="savePatient">保存</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="recordDialogVisible" title="患者挂号记录" width="800px">
      <el-descriptions :column="2" border v-if="currentPatient">
        <el-descriptions-item label="患者姓名">{{
          currentPatient.name
        }}</el-descriptions-item>
        <el-descriptions-item label="性别">{{
          currentPatient.gender
        }}</el-descriptions-item>
        <el-descriptions-item label="身份证号">{{
          currentPatient.id_card
        }}</el-descriptions-item>
        <el-descriptions-item label="联系电话">{{
          currentPatient.phone
        }}</el-descriptions-item>
      </el-descriptions>

      <el-divider />

      <h4>挂号记录</h4>
      <el-table :data="registrationRecords" border style="width: 100%">
        <el-table-column
          prop="registration_date"
          label="就诊日期"
          width="120"
        />
        <el-table-column prop="department_name" label="科室" width="100" />
        <el-table-column prop="doctor_name" label="医生" width="120" />
        <el-table-column prop="fee" label="挂号费(元)" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag
              :type="scope.row.status === 'completed' ? 'success' : 'warning'"
            >
              {{ scope.row.status === "completed" ? "已完成" : "待就诊" }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="挂号时间" />
      </el-table>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from "vue";
import { ElMessage } from "element-plus";
import apiFetch from "../api";

const dialogVisible = ref(false);
const recordDialogVisible = ref(false);
const patientList = ref([]);
const registrationRecords = ref([]);
const currentPatient = ref(null);

const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);

const searchForm = reactive({
  name: "",
  idCard: "",
});

const patientForm = reactive({
  name: "",
  gender: "男",
  id_card: "",
  phone: "",
  birth_date: "",
});

const searchPatients = async () => {
  try {
    const params = new URLSearchParams();
    if (searchForm.name) params.append("name", searchForm.name);
    if (searchForm.idCard) params.append("id_card", searchForm.idCard);
    params.append("page", currentPage.value);
    params.append("limit", pageSize.value);

    const data = await apiFetch(`/api/patients?${params}`);
    if (data && Array.isArray(data.list)) {
      patientList.value = data.list;
      total.value = data.total || 0;
    } else {
      patientList.value = Array.isArray(data) ? data : [];
      total.value = patientList.value.length;
    }
  } catch (error) {
    console.error("查询患者失败:", error);
    ElMessage.error("查询患者信息失败");
  }
};

const handleSizeChange = (val) => {
  pageSize.value = val;
  searchPatients();
};

const handleCurrentChange = (val) => {
  currentPage.value = val;
  searchPatients();
};

const resetSearch = () => {
  searchForm.name = "";
  searchForm.idCard = "";
  currentPage.value = 1;
  searchPatients();
};

const savePatient = async () => {
  if (!patientForm.name || !patientForm.id_card || !patientForm.phone) {
    ElMessage.warning("请填写必填项！");
    return;
  }

  try {
    try {
      await apiFetch("/api/patients", {
        method: "POST",
        body: JSON.stringify(patientForm),
      });
      ElMessage.success("患者档案创建成功！");
      dialogVisible.value = false;
      Object.assign(patientForm, {
        name: "",
        gender: "男",
        id_card: "",
        phone: "",
        birth_date: "",
      });
      searchPatients();
    } catch (err) {
      console.error("保存患者失败:", err);
      ElMessage.error(err.message || "保存患者信息失败");
    }
  } catch (error) {
    console.error("保存患者失败:", error);
    ElMessage.error("保存患者信息失败");
  }
};

const viewRecords = async (patient) => {
  currentPatient.value = patient;
  try {
    const data = await apiFetch(`/api/registrations/patient/${patient.id}`);
    registrationRecords.value = Array.isArray(data) ? data : [];
    recordDialogVisible.value = true;
  } catch (error) {
    console.error("查询挂号记录失败:", error);
    ElMessage.error("查询挂号记录失败");
  }
};

onMounted(() => {
  searchPatients();
});
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
