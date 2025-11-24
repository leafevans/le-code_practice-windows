<template>
  <div>
    <el-card>
      <template #header>
        <div class="card-header">
          <span>收费管理</span>
        </div>
      </template>

      <el-form :inline="true" :model="searchForm">
        <el-form-item label="患者姓名">
          <el-input
            v-model="searchForm.name"
            placeholder="请输入姓名"
            clearable
          />
        </el-form-item>
        <el-form-item label="身份证号">
          <el-input
            v-model="searchForm.id_card"
            placeholder="请输入身份证号"
            clearable
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="searchCharges">查询</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-divider />

      <el-table
        :data="chargeList"
        border
        style="width: 100%"
        empty-text="暂无待缴费记录"
      >
        <el-table-column prop="id" label="挂号ID" width="80" />
        <el-table-column prop="patient_name" label="患者姓名" width="120" />
        <el-table-column prop="id_card" label="身份证号" width="180" />
        <el-table-column prop="department_name" label="科室" width="120" />
        <el-table-column prop="doctor_name" label="医生" width="120" />
        <el-table-column
          prop="registration_date"
          label="就诊日期"
          width="120"
        />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button
              type="primary"
              size="small"
              @click="openPayDialog(scope.row)"
            >
              收费结算
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="payDialogVisible" title="费用结算" width="600px">
      <div v-if="currentCharge">
        <el-descriptions title="基本信息" :column="2" border>
          <el-descriptions-item label="患者姓名">{{
            currentCharge.registration.patient_name
          }}</el-descriptions-item>
          <el-descriptions-item label="挂号ID">{{
            currentCharge.registration.id
          }}</el-descriptions-item>
        </el-descriptions>

        <el-divider />

        <h4>费用明细</h4>
        <el-table
          :data="feeItems"
          border
          style="width: 100%; margin-bottom: 20px"
        >
          <el-table-column prop="name" label="项目" />
          <el-table-column prop="price" label="单价(元)" width="100" />
          <el-table-column prop="quantity" label="数量" width="80" />
          <el-table-column prop="amount" label="金额(元)" width="100" />
        </el-table>

        <div
          style="
            text-align: right;
            font-size: 18px;
            font-weight: bold;
            color: #f56c6c;
          "
        >
          应收总额：¥ {{ currentCharge.fees.total }}
        </div>
      </div>
      <template #footer>
        <el-button @click="payDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmPay">确认收费</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import apiFetch from "../api";

const chargeList = ref([]);
const payDialogVisible = ref(false);
const currentCharge = ref(null);

const searchForm = reactive({
  name: "",
  id_card: "",
});

const searchCharges = async () => {
  try {
    const params = new URLSearchParams();
    if (searchForm.name) params.append("name", searchForm.name);
    if (searchForm.id_card) params.append("id_card", searchForm.id_card);

    const data = await apiFetch(`/api/charges?${params}`);
    chargeList.value = Array.isArray(data) ? data : [];
  } catch (error) {
    console.error("查询待缴费列表失败:", error);
    ElMessage.error("查询失败");
  }
};

const resetSearch = () => {
  searchForm.name = "";
  searchForm.id_card = "";
  searchCharges();
};

const openPayDialog = async (row) => {
  try {
    const data = await apiFetch(`/api/charges/${row.id}/detail`);
    currentCharge.value = data;
    payDialogVisible.value = true;
  } catch (error) {
    console.error("获取费用详情失败:", error);
    ElMessage.error("获取费用详情失败");
  }
};

const feeItems = computed(() => {
  if (!currentCharge.value) return [];
  const items = [];

  if (currentCharge.value.registration.fee > 0) {
    items.push({
      name: "挂号费",
      price: currentCharge.value.registration.fee,
      quantity: 1,
      amount: currentCharge.value.registration.fee,
    });
  }

  currentCharge.value.prescriptions.forEach((p) => {
    items.push({
      name: p.medicine_name,
      price: p.price,
      quantity: p.quantity,
      amount: (p.price * p.quantity).toFixed(2),
    });
  });

  return items;
});

const confirmPay = () => {
  ElMessageBox.confirm(
    `确认收取费用 ¥${currentCharge.value.fees.total} 吗？`,
    "收费确认",
    {
      confirmButtonText: "确认收款",
      cancelButtonText: "取消",
      type: "warning",
    }
  )
    .then(async () => {
      try {
        await apiFetch(
          `/api/charges/${currentCharge.value.registration.id}/pay`,
          {
            method: "POST",
          }
        );
        ElMessage.success("缴费成功！");
        payDialogVisible.value = false;
        searchCharges();
      } catch (error) {
        console.error("缴费失败:", error);
        ElMessage.error("缴费失败");
      }
    })
    .catch(() => {});
};

onMounted(() => {
  searchCharges();
});
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
