<template>
    <!--  -->
    <div class="job-title-box">
        <div class="total">
            <el-button type="primary" plain @click="openDialog('create', null)"> 创 建 </el-button>
            <el-input v-model="searchData" clearable :style="{width:'400px', marginLeft:'20px'}" placeholder="通过名称检索" >
                <template #append>
                    <el-button :icon="Search" @click="getListFunc" />
                </template>
            </el-input>
            <el-button type="primary" plain @click="getListFunc" :style="{marginLeft:'20px'}">刷 新</el-button>
        </div>
        <div class="box">
            <el-table :data="tableData" :style="{ height: '100%',width: '100%' }">
                <el-table-column prop="name" label="职称" align="center" />
                <el-table-column prop="category" label="职称类别" align="center" />
                <el-table-column prop="grade" label="职称级别" align="center" />
                <el-table-column prop="ceilingvalue" label="---------" align="center" />
                <el-table-column width="200px" align="center">
                    <template #default="scope">
                        <el-button link type="primary" size="small" :style="{ padding: '0', lineHeight: '14px' }" @click="openDialog('edit', scope.row)">编辑</el-button>
                        <el-popconfirm
                            width="220"
                            confirm-button-text="确认"
                            cancel-button-text="取消"
                            :icon="InfoFilled"
                            icon-color="#626AEF"
                            title="确定删除?"
                            @confirm="deleteStaffInfoFunc(scope.row.id)">
                            <template #reference>
                                <el-button link type="primary" size="small" :style="{ padding: '0', lineHeight: '14px' }">删除</el-button>
                            </template>
                        </el-popconfirm>
                    </template>
                </el-table-column>
            </el-table>
        </div>
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
          width="1000px"
          :before-close="handleClose"
        >
          <div>
            <el-form class="demo-form-inline" label-width="120px">
                <el-form-item label="名称">
                  <el-input v-model="formData.name" clearable />
                </el-form-item>
                <el-form-item label="指标值">
                    <el-input v-model="formData.indicatValue" text="number" clearable />
                </el-form-item>
                <el-form-item label="职称等级">
                    <el-select v-model="formData.grade">
                        <el-option label="正高级" value="正高级"></el-option>
                        <el-option label="副高级" value="副高级"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="职称类别">
                    <el-select v-model="formData.category">
                        <el-option label="护理" value="护理"></el-option>
                        <el-option label="技师" value="技师"></el-option>
                        <el-option label="医师" value="医师"></el-option>
                        <el-option label="工人" value="工人"></el-option>
                        <el-option label="其他" value="其他"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="下限">
                    <el-input v-model="formData.floorvalue" text="text" />
                </el-form-item>
                <el-form-item label="上限">
                    <el-input v-model="formData.ceilingvalue" text="text" />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="formData.note" type="textarea" />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="dialogButton">{{ dialogType==='create'?'新增':'更新' }}</el-button>
                </el-form-item>
            </el-form>
          </div>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getList, createJobTitle, updateJobTitle, deleteJobTitle } from '@/api/jobTitle'
import { ElMessage } from 'element-plus';
import { Search } from '@element-plus/icons-vue'
import { useUserStore } from '@/pinia/modules/user'
import { storeToRefs  } from 'pinia'

const userStroe = useUserStore()
const userInfo = storeToRefs(userStroe).userInfo

const page = ref(0)
const pageSize = ref(20)
const total = ref(0)
const searchData = ref('')
const tableData = ref([])
const dialogVisible = ref(false)
const dialogType = ref('')
const itemId = ref(0)
const formData = reactive({
    name: '',
    grade: '',
    indicatValue: 0.00,
    category: '',
    floorvalue: '',
    ceilingvalue: '',
    note: ''
}) 
const initForm = () => {
    formData.name = ''
    formData.grade = ''
    formData.indicatValue = 0.00
    formData.category = ''
    formData.floorvalue = ''
    formData.ceilingvalue = ''
    formData.note = ''
}
onMounted(() => {
    getListFunc()
})
const handleClose = (down) => {
    initForm()
    down()
}
const getListFunc = () => {
    getList({
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchData.value
    }).then(res => {
        console.log(res);
        if (res.code === 0) {
            ElMessage.success(res.msg)
            tableData.value = res.data.data
            total.value = res.data.total
        } else {
            ElMessage.error(res.msg)
        }
    })
}
const openDialog = (type, obj) => {
    dialogType.value = type
    if (type === 'creat') {
        initForm()
    }
    if (type === 'edit') {
        console.log(obj);
        itemId.value = obj.id
        for (const key in formData) {
            formData[key] = obj[key]
        }
    }
    dialogVisible.value = true
}

const handleSizeChange = () => {
    getListFunc()
}
const handleCurrentChange = () => {
    getListFunc()
}
const dialogButton = () => {
    if (dialogType.value === 'create') {
        createJobTitleFunc()
    }
    if (dialogType.value === 'edit') {
        updateOfTable()
    }
}
const createJobTitleFunc = () => {
    createJobTitle({
        name: formData.name,
        grade: formData.grade,
        indicatValue: parseFloat(formData.indicatValue),
        category: formData.category,
        floorvalue: formData.floorvalue,
        ceilingvalue: formData.ceilingvalue,
        note: formData.note,
        employeeId: Number(userInfo.value.employee_id)
    }).then(res => {
        console.log(res);
        if (res.code === 0) {
            ElMessage.success(res.msg)
        } else {
            ElMessage.error(res.msg)
        }
    })
}
const updateOfTable = () => {
    console.log({
        id: itemId.value,
        ...formData,
        employeeId: Number(userInfo.value.employee_id)
    });
    updateJobTitle({
        id: itemId.value,
        name: formData.name,
        grade: formData.grade,
        indicatValue: parseFloat(formData.indicatValue),
        category: formData.category,
        floorvalue: formData.floorvalue,
        ceilingvalue: formData.ceilingvalue,
        note: formData.note,
        employeeId: Number(userInfo.value.employee_id)
    }).then(res => {
        console.log(res);
        if (res.code === 0) {
            ElMessage.success(res.msg)
        } else {
            ElMessage.error(res.msg)
        }
    })
}
</script>

<style lang="scss" scoped>
.job-title-box {
    height: calc(100vh - 126px);
    .total {
        height: 36px;
    }
    .box {
        height: calc(100% - 76px);
    }
    .page-ctrl {
        height: 40px;
        display: flex;
        justify-content: flex-start;
        align-items: center;
        .el-pagination {
            margin: 5px auto;
        }
    }
}
</style>