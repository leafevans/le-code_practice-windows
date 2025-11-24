<template>
  <div>
    <el-card>
      <template #header>
        <div class="card-header">
          <span>医生工作台</span>
          <div style="display: flex; align-items: center">
            <span style="margin-right: 10px; font-size: 14px">当前医生:</span>
            <el-select
              v-model="currentDoctorId"
              placeholder="请选择医生"
              style="width: 200px"
              @change="loadPendingRegistrations"
            >
              <el-option
                v-for="doc in doctorList"
                :key="doc.id"
                :label="doc.name + ' (' + doc.department_name + ')'"
                :value="doc.id"
              />
            </el-select>
          </div>
        </div>
      </template>

      <div
        v-if="!currentDoctorId"
        style="text-align: center; color: #999; padding: 40px"
      >
        请先选择医生以开始工作
      </div>

      <div v-else>
        <el-tabs v-model="activeTab">
          <el-tab-pane label="待接诊患者" name="pending">
            <el-table
              :data="pendingList"
              border
              style="width: 100%"
              empty-text="暂无待接诊患者"
            >
              <el-table-column
                prop="registration_date"
                label="就诊日期"
                width="120"
              />
              <el-table-column prop="patient_name" label="姓名" width="120" />
              <el-table-column prop="patient_gender" label="性别" width="80" />
              <el-table-column
                prop="patient_birth_date"
                label="出生日期"
                width="120"
              />
              <el-table-column label="操作">
                <template #default="scope">
                  <el-button
                    type="primary"
                    size="small"
                    @click="startConsultation(scope.row)"
                  >
                    接诊
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="已接诊记录" name="completed">
            <el-table
              :data="completedList"
              border
              style="width: 100%"
              empty-text="暂无接诊记录"
            >
              <el-table-column
                prop="registration_date"
                label="就诊日期"
                width="120"
              />
              <el-table-column prop="patient_name" label="姓名" width="120" />
              <el-table-column prop="diagnosis" label="诊断结果" />
              <el-table-column label="操作" width="100">
                <template #default="scope">
                  <el-button size="small" @click="viewDetail(scope.row)"
                    >查看</el-button
                  >
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-card>

    <el-dialog
      v-model="consultationVisible"
      title="接诊"
      width="800px"
      :close-on-click-modal="false"
    >
      <div v-if="currentPatient">
        <el-descriptions title="患者信息" :column="3" border>
          <el-descriptions-item label="姓名">{{
            currentPatient.patient_name
          }}</el-descriptions-item>
          <el-descriptions-item label="性别">{{
            currentPatient.patient_gender
          }}</el-descriptions-item>
          <el-descriptions-item label="出生日期">{{
            currentPatient.patient_birth_date
          }}</el-descriptions-item>
        </el-descriptions>

        <el-divider />

        <el-form :model="consultationForm" label-width="80px">
          <el-form-item label="主诉" required>
            <el-input
              v-model="consultationForm.symptom"
              type="textarea"
              rows="2"
              placeholder="患者描述的症状"
            />
          </el-form-item>
          <el-form-item label="诊断" required>
            <el-input
              v-model="consultationForm.diagnosis"
              type="textarea"
              rows="2"
              placeholder="医生的诊断结果"
            />
          </el-form-item>

          <el-divider content-position="left">处方开药</el-divider>

          <div
            v-for="(item, index) in consultationForm.prescriptions"
            :key="index"
            style="display: flex; gap: 10px; margin-bottom: 10px"
          >
            <el-select
              v-model="item.medicine_id"
              placeholder="选择药品"
              filterable
              style="width: 200px"
            >
              <el-option
                v-for="med in medicineList"
                :key="med.id"
                :label="med.name + ' (库存:' + med.stock + ')'"
                :value="med.id"
                :disabled="med.stock <= 0"
              />
            </el-select>
            <el-input-number
              v-model="item.quantity"
              :min="1"
              placeholder="数量"
              style="width: 120px"
            />
            <el-input
              v-model="item.dosage"
              placeholder="用法用量 (如: 每日3次, 每次1粒)"
              style="flex: 1"
            />
            <el-button
              type="danger"
              icon="Delete"
              circle
              @click="removeMedicine(index)"
            />
          </div>

          <el-button
            type="primary"
            plain
            size="small"
            @click="addMedicine"
            style="margin-top: 10px"
          >
            <el-icon><Plus /></el-icon> 添加药品
          </el-button>
        </el-form>
      </div>
      <template #footer>
        <el-button @click="consultationVisible = false">取消</el-button>
        <el-button type="primary" @click="submitConsultation"
          >完成接诊</el-button
        >
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from "vue";
import { ElMessage } from "element-plus";
import { Delete, Plus } from "@element-plus/icons-vue";
import apiFetch from "../api";

const currentDoctorId = ref(null);
const doctorList = ref([]);
const activeTab = ref("pending");
const pendingList = ref([]);
const completedList = ref([]);
const medicineList = ref([]);

const consultationVisible = ref(false);
const currentPatient = ref(null);
const consultationForm = reactive({
  symptom: "",
  diagnosis: "",
  prescriptions: [],
});

const loadDoctors = async () => {
  try {
    const data = await apiFetch("/api/doctors?limit=100");
    doctorList.value = data.list || [];
  } catch (error) {
    console.error("加载医生列表失败:", error);
  }
};

const loadMedicines = async () => {
  try {
    const data = await apiFetch("/api/medicines");
    medicineList.value = Array.isArray(data) ? data : [];
  } catch (error) {
    console.error("加载药品列表失败:", error);
  }
};

const loadPendingRegistrations = async () => {
  if (!currentDoctorId.value) return;
  try {
    const data = await apiFetch(
      `/api/registrations/doctor/${currentDoctorId.value}?status=pending`
    );
    pendingList.value = Array.isArray(data) ? data : [];

    const completedData = await apiFetch(
      `/api/registrations/doctor/${currentDoctorId.value}?status=completed`
    );
    completedList.value = Array.isArray(completedData) ? completedData : [];
  } catch (error) {
    console.error("加载挂号列表失败:", error);
    ElMessage.error("加载患者列表失败");
  }
};

const startConsultation = (row) => {
  currentPatient.value = row;
  consultationForm.symptom = "";
  consultationForm.diagnosis = "";
  consultationForm.prescriptions = [];
  consultationVisible.value = true;
  loadMedicines();
};

const addMedicine = () => {
  consultationForm.prescriptions.push({
    medicine_id: null,
    quantity: 1,
    dosage: "",
  });
};

const removeMedicine = (index) => {
  consultationForm.prescriptions.splice(index, 1);
};

const submitConsultation = async () => {
  if (!consultationForm.symptom || !consultationForm.diagnosis) {
    ElMessage.warning("请填写主诉和诊断结果");
    return;
  }

  for (const p of consultationForm.prescriptions) {
    if (!p.medicine_id || !p.quantity) {
      ElMessage.warning("请完善药品信息");
      return;
    }
  }

  try {
    await apiFetch("/api/consultations", {
      method: "POST",
      body: JSON.stringify({
        registration_id: currentPatient.value.id,
        symptom: consultationForm.symptom,
        diagnosis: consultationForm.diagnosis,
        prescriptions: consultationForm.prescriptions,
      }),
    });

    ElMessage.success("接诊完成！");
    consultationVisible.value = false;
    loadPendingRegistrations();
  } catch (error) {
    console.error("提交失败:", error);
    ElMessage.error(error.message || "提交失败");
  }
};

const viewDetail = (row) => {
  ElMessage.info("详情查看功能待开发");
};

watch(activeTab, () => {
  if (currentDoctorId.value) {
    loadPendingRegistrations();
  }
});

onMounted(() => {
  loadDoctors();
});
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
