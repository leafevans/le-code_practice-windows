<template>
  <div>
    <el-card>
      <template #header>
        <div class="card-header">
          <span>药品管理</span>
          <el-button type="primary" @click="dialogVisible = true">
            <el-icon><Plus /></el-icon> 添加药品
          </el-button>
        </div>
      </template>

      <el-form :inline="true">
        <el-form-item label="药品名称">
          <el-input
            v-model="searchForm.name"
            placeholder="请输入药品名称"
            clearable
          />
        </el-form-item>
        <el-form-item label="药品分类">
          <el-select
            v-model="searchForm.category"
            placeholder="请选择分类"
            clearable
          >
            <el-option label="抗生素" value="抗生素" />
            <el-option label="解热镇痛" value="解热镇痛" />
            <el-option label="感冒药" value="感冒药" />
            <el-option label="维生素" value="维生素" />
            <el-option label="中成药" value="中成药" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="searchMedicines">查询</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-divider />

      <el-table :data="medicineList" border style="width: 100%">
        <el-table-column prop="id" label="药品ID" width="80" />
        <el-table-column prop="name" label="药品名称" width="150" />
        <el-table-column prop="specification" label="规格" width="120" />
        <el-table-column prop="unit" label="单位" width="80" />
        <el-table-column prop="price" label="单价(元)" width="100" />
        <el-table-column prop="stock" label="库存" width="100">
          <template #default="scope">
            <el-tag
              :type="
                scope.row.stock > 100
                  ? 'success'
                  : scope.row.stock > 50
                  ? 'warning'
                  : 'danger'
              "
            >
              {{ scope.row.stock }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="category" label="分类" width="100" />
        <el-table-column label="操作" width="150">
          <template #default="scope">
            <el-button
              size="small"
              type="primary"
              @click="updateStock(scope.row)"
            >
              调整库存
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" title="添加药品" width="500px">
      <el-form :model="medicineForm" label-width="100px">
        <el-form-item label="药品名称" required>
          <el-input v-model="medicineForm.name" />
        </el-form-item>
        <el-form-item label="规格">
          <el-input
            v-model="medicineForm.specification"
            placeholder="例如：0.25g*24粒"
          />
        </el-form-item>
        <el-form-item label="单位">
          <el-input v-model="medicineForm.unit" placeholder="例如：盒、瓶" />
        </el-form-item>
        <el-form-item label="单价(元)" required>
          <el-input-number
            v-model="medicineForm.price"
            :min="0"
            :precision="2"
          />
        </el-form-item>
        <el-form-item label="初始库存">
          <el-input-number v-model="medicineForm.stock" :min="0" />
        </el-form-item>
        <el-form-item label="分类">
          <el-select v-model="medicineForm.category" placeholder="请选择分类">
            <el-option label="抗生素" value="抗生素" />
            <el-option label="解热镇痛" value="解热镇痛" />
            <el-option label="感冒药" value="感冒药" />
            <el-option label="维生素" value="维生素" />
            <el-option label="中成药" value="中成药" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveMedicine">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import apiFetch from "../api";

const dialogVisible = ref(false);
const medicineList = ref([]);

const searchForm = reactive({
  name: "",
  category: "",
});

const medicineForm = reactive({
  name: "",
  specification: "",
  unit: "盒",
  price: 0,
  stock: 0,
  category: "",
});

const searchMedicines = async () => {
  try {
    const params = new URLSearchParams();
    if (searchForm.name) params.append("name", searchForm.name);
    if (searchForm.category) params.append("category", searchForm.category);

    const data = await apiFetch(`/api/medicines?${params}`);
    medicineList.value = Array.isArray(data) ? data : [];
  } catch (error) {
    console.error("查询药品失败:", error);
    ElMessage.error("查询药品信息失败");
  }
};

const resetSearch = () => {
  searchForm.name = "";
  searchForm.category = "";
  searchMedicines();
};

const saveMedicine = async () => {
  if (!medicineForm.name || !medicineForm.price) {
    ElMessage.warning("请填写必填项！");
    return;
  }

  try {
    await apiFetch("/api/medicines", {
      method: "POST",
      body: JSON.stringify(medicineForm),
    });
    ElMessage.success("药品添加成功！");
    dialogVisible.value = false;
    Object.assign(medicineForm, {
      name: "",
      specification: "",
      unit: "盒",
      price: 0,
      stock: 0,
      category: "",
    });
    searchMedicines();
  } catch (error) {
    console.error("保存药品失败:", error);
    ElMessage.error(error.message || "保存药品信息失败");
  }
};

const updateStock = (medicine) => {
  ElMessageBox.prompt(
    `当前库存：${medicine.stock}，请输入新的库存数量`,
    "调整库存",
    {
      confirmButtonText: "确认",
      cancelButtonText: "取消",
      inputPattern: /^\d+$/,
      inputErrorMessage: "请输入有效的库存数量",
      inputValue: medicine.stock.toString(),
    }
  )
    .then(async ({ value }) => {
      try {
        await apiFetch(`/api/medicines/${medicine.id}`, {
          method: "PUT",
          body: JSON.stringify({ stock: parseInt(value, 10) }),
        });
        ElMessage.success("库存更新成功！");
        searchMedicines();
      } catch (error) {
        console.error("更新库存失败:", error);
        ElMessage.error(error.message || "更新库存失败");
      }
    })
    .catch(() => {
      ElMessage.info("已取消操作");
    });
};

onMounted(() => {
  searchMedicines();
});
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
