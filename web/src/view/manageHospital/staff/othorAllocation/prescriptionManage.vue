<template>
    <div>
        <ManageTemplate
            @getList="GetList"
            @create="createDialog"
            :total="total"
            :showCreate="true"
        >
            <el-table :data="tableData" :style="{ height: '100%',width: '100%' }">
                <el-table-column prop="name" label="处方资质" align="center" />
                <el-table-column label="是否启用" align="center">
                    <template #default="scope">
                        <el-switch
                          v-model="scope.row.active"
                          class="ml-2"
                          style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                          @change="updateFunc(scope.row)"
                        />
                    </template>
                </el-table-column>
                <el-table-column width="200px" align="center">
                    <template #default="scope">
                        <el-button link type="primary" size="small" :style="{ padding: '0', lineHeight: '14px' }" @click="editDialog">编辑</el-button>
                        <el-popconfirm
                            width="220"
                            confirm-button-text="确认"
                            cancel-button-text="取消"
                            :icon="InfoFilled"
                            icon-color="#626AEF"
                            title="确定删除?"
                            @confirm="deleteFunc(scope.row.id)">
                            <template #reference>
                                <el-button link type="primary" size="small" :style="{ padding: '0', lineHeight: '14px' }">删除</el-button>
                            </template>
                        </el-popconfirm>
                    </template>
                </el-table-column>
            </el-table>
        </ManageTemplate>
        <el-dialog
          v-model="dialogVisible"
          :title="dialogType==='create'?'新增保险条目':'编辑保险条目'"
          width="800px"
          :before-close="handleClose"
        >
          <div>
            <el-form class="demo-form-inline">
                <el-form-item label="名称">
                  <el-input v-model="formData.name" placeholder="处方资质名称" clearable />
                </el-form-item>
                <el-form-item label="是否启用">
                    <el-switch
                          v-model="formData.active"
                          class="ml-2"
                          style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                        />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="dialogOk">{{ dialogType==='create'?'新增':'编辑完成' }}</el-button>
                </el-form-item>
            </el-form>
          </div>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import {InfoFilled} from '@element-plus/icons-vue'
import ManageTemplate from '@/components/manageTemplate/index.vue'
import { getListPQ, createPQ, updatePQ, deletePQ } from '@/api/prescriptionQualification'
import { ElMessage } from 'element-plus';
import { useUserStore } from '@/pinia/modules/user'
import { storeToRefs  } from 'pinia'

const userStroe = useUserStore()
const userInfo = storeToRefs(userStroe).userInfo

const total = ref(0)
const dialogVisible = ref(false)
const dialogType = ref('craete')
onMounted(() => {
    GetList()
})
const handleClose = (down) => {
    down()
}
const GetList = (obj) => {
    getListPQ(obj).then(res => {
        if (res.code === 0) {
            console.log(res.data.data);
            tableData.value = res.data.data
            total.value = res.data.total
        } else {
            ElMessage.error(res.msg)
        }
    })
}
const tableData = ref([])
const pqItem = ref({})
const formData = reactive({
    name: '',
    active: true
})
const editId = ref(0)
const clearForm = () => {
    formData.name = ''
    formData.active = true
}
const createDialog = () => {
    clearForm()
    dialogType.value = 'create'
    dialogVisible.value = true
}
const editDialog = (obj) => {
    dialogType.value = 'edit'
    formData.name = obj.name
    formData.active = obj.active
    editId.value = obj.id
    dialogVisible.value = true
    pqItem.value = obj
}
const dialogOk = () => {
    if (dialogType.value = 'create') {
        createFunc()
    } else {
        updateOfEdit()
    }
}
const createFunc = () => {
    createPQ({
        ...formData,
        employeeId: userInfo.employee_id
    }).then(res => {
        ElMessage.info(res.msg)
        GetList()
    })
}
const updateFunc = (obj) => {
    console.log(obj);
    updatePQ({
        name: obj.name,
        active: obj.active,
        id: obj.id,
        employeeId: userInfo.employee_id
    }).then(res => {
        if (res.code === 0) {
            ElMessage.success(res.msg)
        } else {
            ElMessage.error(res.msg)
        }
        GetList()
    })
}
const updateOfEdit = () =>  {
    console.log(obj);
    updatePQ({
        ...formData,
        id: editId.value,
        employeeId: userInfo.employee_id
    }).then(res => {
        if (res.code === 0) {
            ElMessage.success(res.msg)
        } else {
            ElMessage.error(res.msg)
        }
        GetList()
    })
}
const deleteFunc = (id) => {
    deletePQ({id}).then(res => {
        if (res.code === 0) {
            ElMessage.success(res.msg)
        } else {
            ElMessage.error(res.msg)
        }
        GetList()
    })
}
</script>

<style lang="scss" scoped>

</style>