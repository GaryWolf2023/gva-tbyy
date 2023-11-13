<template>
    <div class="insurance-box">
        <div class="ctrl-insurance">
            <el-button type="primary" plain @click="createInsuranceButton">
                新增<el-icon><Plus/></el-icon>
            </el-button>
            <el-input v-model="searchData" placeholder="根据名称检索保险条目" :style="{width:'200px',margin:'0 0 0 20px'}" clearable/>
            <el-button type="primary" plain @click="getInsuranceListFunc">
                刷新<el-icon><Refresh /></el-icon>
            </el-button>
        </div>
        <el-table :data="tableData" :style="{width:'100%',height:'calc(100% - 82px)'}">
            <el-table-column prop="name" label="保险名称" align="center" />
            <el-table-column prop="baseNum" label="缴费基数" width="280" align="center" />
            <el-table-column prop="percentUnit" label="单位缴纳比例" width="280" align="center" />
            <el-table-column prop="percentPerson" label="个人缴纳比例" width="280" align="center" />
            <el-table-column label="操作" width="380"  align="center">
                <template #default="scope">
                    <el-icon :style="{padding:'0 10px'}"><EditPen @click="editInsuranceButton(scope.row)"/></el-icon>
                    <el-icon :style="{padding:'0 10px'}"><Delete @click="deleteInsuranceFunc(scope.row)"/></el-icon>
                </template>
            </el-table-column>
        </el-table>
        <div class="page-ctrl">
            <el-pagination
              v-model:current-page="page"
              v-model:page-size="pageSize"
              :page-sizes="[20, 40, 60, 100]"
              :small="true"
              :background="true"
              layout="total, sizes, prev, pager, next, jumper"
              :total="total"
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
            />
        </div>
        <el-dialog
          v-model="dialogVisible"
          :title="dialogType==='create'?'新增保险条目':'编辑保险条目'"
          width="800px"
          :before-close="handleClose"
        >
          <div>
            <el-form :inline="true" class="demo-form-inline">
              <el-form-item label="保险名称">
                <el-input v-model="formData.name" clearable />
              </el-form-item>
              <el-form-item label="缴费基数">
                <el-input v-model="formData.baseNum" text="number" clearable />
              </el-form-item>
              <el-form-item label="单位缴纳比例">
                <el-input v-model="formData.percentUnit" text="number" clearable />
              </el-form-item>
              <el-form-item label="个人缴纳比例">
                <el-input v-model="formData.percentPerson" text="number" clearable />
              </el-form-item>
            </el-form>
          </div>
          <template #footer>
            <span class="dialog-footer">
              <el-button @click="dialogVisible=false">{{dialogType==='create'?'取 消 创 建':'取 消 编 辑'}}</el-button>
              <el-button type="primary" @click="ctrlInsurance">{{dialogType==='create'?'确 认 创 建':'确 认 编 辑'}}</el-button>
            </span>
          </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { EditPen, Refresh, Plus, Delete } from '@element-plus/icons-vue'
import { ref, onMounted, watch, computed, reactive } from 'vue'
import { ElMessage } from 'element-plus';
import { getInsuranceList, getInsuranceInfo, createInsurance, updateInsurance, deleteInsurance } from '@/api/insuranceManage.js'

const page = ref(0)
const pageSize = ref(20)
const total = ref(0)
const tableData = ref([])
const searchData = ref('')
const dialogVisible = ref(false)
const dialogType = ref('create')
const itemData = ref({})
const formData = reactive({
    name: '',
    baseNum: 0,
    percentUnit: 0,
    percentPerson: 0
})
const handleSizeChange = () => {
    getInsuranceListFunc()
}
const handleCurrentChange = () => {
    getInsuranceListFunc()
}
onMounted(() => {
    getInsuranceListFunc()
})
const getInsuranceListFunc = () => {
    getInsuranceList({
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchData.value
    }).then(res => {
        console.log(res)
        if (res.code === 0) {
            tableData.value = res.data.data
            total.value = res.data.total
        } else {
            ElMessage.warning(res.msg)
        }
    })
}
const handleClose = (down) => {
    itemData.value = {}
    clearFormData()
    down()
}
const clearFormData = () => {
    formData.name = ''
    formData.baseNum = 0
    formData.percentUnit = 0
    formData.percentPerson = 0
}
const editInsuranceButton = (val) => {
    dialogType.value = 'edit'
    for (const key in formData) {
        formData[key] = val[key]
    }
    itemData.value = val
    dialogVisible.value = true
}
const createInsuranceButton = () => {
    dialogType.value = 'create'
    dialogVisible.value = true
}
const ctrlInsurance = () => {
    console.log(dialogType.value);
    if (dialogType.value === 'create') {
        createInsuranceFunc()
    }
    if (dialogType.value === 'edit') {
        editInsuranceFunc()
    }
}
const createInsuranceFunc = () => {
    createInsurance({ ...formData, createUid: 2 }).then(res => {
        console.log(res)
        if (res.code === 0) {
            ElMessage.success(res.msg)
            getInsuranceListFunc()
        } else {
            ElMessage.warning(res.msg)
        }
    }).finally(() => {
        dialogVisible.value = false
    })
}
const editInsuranceFunc = (val) => {
    console.log(val)
    updateInsurance({ ...formData, id: itemData.value.id, writeUid: 2 }).then(res => {
        console.log(res)
        if (res.code === 0) {
            ElMessage.success(res.msg)
            getInsuranceListFunc()
        } else {
            ElMessage.warning(res.msg)
        }
    }).finally(() => {
        dialogVisible.value = false
    })
}
const deleteInsuranceFunc = (val) => { 
    console.log(val)
    deleteInsurance({ id: val.id }).then(res => {
        if (res.code === 0) {
            getInsuranceListFunc()
        }
    })
 }
</script>

<style lang="scss" scoped>
.insurance-box {
    height: calc(100vh - 124px);
    .ctrl-insurance {
        height: 46px;
        margin-bottom: 4px;
        border-radius: 8px;
        background-color: rgba(255, 255, 255, 0.801);
        display: flex;
        align-items: center;
        padding-left: 18px;
    }
    .page-ctrl {
        height: 32px;
        display: flex;
        justify-content: flex-start;
        align-items: center;
        .el-pagination {
            margin: 5px auto;
        }
    }
}
</style>