<template>
    <div class="staff-manage">
        <div class="ctrl-list">
            <el-button type="primary" plain @click="createStaffFunc">创 建 员 工</el-button>
            <!-- <el-button type="primary" plain v-if="showEditorDialog" @click.="showEditorDialog=false">模 板 列 表</el-button> -->
            <el-input v-model="searchData" clearable :style="{width:'400px', marginLeft:'20px'}" placeholder="通过名称检索员工" >
                <template #append>
                    <el-button :icon="Search" @click="getStaffListFunc" />
                </template>
            </el-input>
            <el-button type="primary" plain @click="getStaffListFunc" :style="{marginLeft:'20px'}">刷 新</el-button>
        </div>
        <div class="table-show">
            <el-table :data="tableData" style="width: 100%" :style="{height:'100%'}">
                <el-table-column prop="id" label="ID" width="80" align="center" />
                <el-table-column prop="name" label="姓名" width="180" align="center" />
                <el-table-column prop="name" label="姓名" width="180" align="center">
                    <template #default="scope">
                        <el-switch
                          v-model="scope.row.active"
                          class="ml-2"
                          @change="updateStaffInfo(scope.row)"
                          style="--el-switch-off-color: #ff4949; --el-switch-on-color: #13ce66;"
                        />
                    </template>
                </el-table-column>
                <el-table-column prop="jobTitle" label="工作岗位" align="center" />
                <el-table-column prop="performanceclass" label="职位" align="center" />
                <el-table-column prop="workPhone" label="工作电话" align="center" />
                <el-table-column prop="tocompile" label="在职类型" align="center" />
                <el-table-column label="更多" width="120px" align="center">
                    <template #default="scope">
                            <el-button link type="primary" size="small" :style="{padding:'0',lineHeight:'14px'}" @click="editStaffInfoFunc(scope.row)">编辑</el-button>
                            <el-popconfirm
                                width="220"
                                confirm-button-text="确认"
                                cancel-button-text="取消"
                                :icon="InfoFilled"
                                icon-color="#626AEF"
                                title="确定删除?"
                                @confirm="deleteStaffInfoFunc(scope.row.id)"
                            >
                                <template #reference>
                                    <el-button link type="primary" size="small" :style="{padding:'0',lineHeight:'14px'}">删除</el-button>
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
              small
              background
              layout="total, sizes, prev, pager, next, jumper"
              :total="total"
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
            />
        </div>
    </div>
</template>

<script setup>
import { getStaffList, getStaffInfo, deleteStaff, updateStaff } from '@/api/staff.js'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElButton, ElMessage, ElMessageBox } from 'element-plus'
import { Search, InfoFilled } from '@element-plus/icons-vue'

const router = useRouter()
const page = ref(1)
const pageSize = ref(20)
const total = ref(0)
const searchData = ref('')
const tableData = ref([])
const staffInfo = ref({})

onMounted(() => {
    getStaffListFunc()

})
const handleClose = (down) => {
    ElMessageBox.confirm(
    '关闭后所填写信息将会被清空，请谨慎操作?',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
      return down()
    })
    .catch(() => {
    })
}
const handleSizeChange = (a) => {
    pageSize.value = a
    getStaffListFunc()
}
const handleCurrentChange = (b) => {
    page.value = b
    getStaffListFunc()
}
const getStaffListFunc = () => {
    getStaffList({
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchData.value
    }).then(res => {
        tableData.value = res.data.data
        total.value = res.data.total
    })
}
const editStaffInfoFunc = (raw) => {
    router.push({
                path: '/layout/hosManage/staffManage/editorStaff',
                query: {id: raw.id}
            })
    
}
const createStaffFunc = () => {
    console.log('创建员工')
}
const updateStaffInfo = (info) => {
    console.log(info)
    updateStaff(info.value).then(res => {
        console.log(res)
        if (res.code === 0) {
            ElMessage.success(res.msg)
        } else {
            
        }
    })
}
const deleteStaffInfoFunc = (id) => {
    deleteStaff(id).then(res => {
        console.log(res)
    })
}
const getStaffInfoFunc = (id) => {
        getStaffInfo(id).then(res => {
            console.log(res);
            if (res.code === 0) { 
                staffInfo.value = res.data
            } else {
                ElMessage.warning("获取个人信息失败")
            }
        })
}
</script>

<style lang="scss" scoped>
.staff-manage {
    height: calc(100vh - 120px);
    .ctrl-list {
        height: 36px;
        margin-bottom: 10px;
    }
    .table-show {
        height: calc(100% - 84px);
        .el-table {
            padding: 0;
        }
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
    :deep(.el-dialog > .el-dialog__body) {
        padding: 0 10px;
    }
}
</style>