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
              查看病历
            </el-button>
            <el-button
              size="small"
              type="success"
              @click="addRecord(scope.row)"
            >
              添加病历
            </el-button>
          </template>
        </el-table-column>
      </el-table>
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
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from "vue";
import { ElMessage } from "element-plus";

const dialogVisible = ref(false);
const patientList = ref([]);

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

    const response = await fetch(
      `http://localhost:3000/api/patients?${params}`
    );
    const data = await response.json();
    patientList.value = data;
  } catch (error) {
    console.error("查询患者失败:", error);
    ElMessage.error("查询患者信息失败");
  }
};

const resetSearch = () => {
  searchForm.name = "";
  searchForm.idCard = "";
  searchPatients();
};

const savePatient = async () => {
  if (!patientForm.name || !patientForm.id_card || !patientForm.phone) {
    ElMessage.warning("请填写必填项！");
    return;
  }

  try {
    const response = await fetch("http://localhost:3000/api/patients", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(patientForm),
    });

    if (response.ok) {
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
    }
  } catch (error) {
    console.error("保存患者失败:", error);
    ElMessage.error("保存患者信息失败");
  }
};

const viewRecords = (patient) => {
  ElMessage.info(`查看患者 ${patient.name} 的病历（功能待实现）`);
};

const addRecord = (patient) => {
  ElMessage.info(`为患者 ${patient.name} 添加病历（功能待实现）`);
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
