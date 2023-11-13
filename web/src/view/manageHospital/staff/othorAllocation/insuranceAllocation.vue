<template>
    <div>
        <el-table :data="tableData" style="width: 100%" max-height="250">
            <el-table-column fixed prop="name" label="保险类型名称" />
            <el-table-column prop="baseNum" label="缴费基数" />
            <el-table-column prop="percentUnit" label="单位缴纳比例" />
            <el-table-column prop="percentPerson" label="个人缴纳比例" />
            <el-table-column fixed="right" label="" width="120">
              <template #default="scope">
                <el-button
                  link
                  type="primary"
                  size="small"
                  @click.prevent="deleteRowOfInsurance(scope.row)"
                >
                <el-icon><CloseBold /></el-icon>
                </el-button>
              </template>
            </el-table-column>
        </el-table>
        <el-button color="#626aef" class="mt-4" style="width: 100%" plain @click="dialogVisible=true" size="large">新 增 保 险 条 目</el-button>
        <el-dialog
          v-model="dialogVisible"
          title="新增保险条目"
          width="860px"
          top="1vh"
          :before-close="handleClose"
          @open="Opendialog"
        >
        <div class="insurance-list">
            <div class="ctrl-insurance">
               <el-input v-model="searchData" placeholder="根据名称检索保险条目" :style="{width:'200px',margin:'0 0 0 20px'}" clearable/>
               <el-button type="primary" plain @click="getInsuranceListFunc">
                   刷新<el-icon><Refresh /></el-icon>
               </el-button>
            </div>
            <el-table :data="tableData1" style="width: 100%" max-height="540" >
                <el-table-column fixed prop="name" label="保险类型名称" />
                <el-table-column prop="baseNum" label="缴费基数" />
                <el-table-column prop="percentUnit" label="单位缴纳比例" />
                <el-table-column prop="percentPerson" label="个人缴纳比例" />
                <el-table-column fixed="right" label="" width="120">
                  <template #default="scope">
                    <el-button
                      link
                      type="primary"
                      size="small"
                      @click.prevent="addInsuranceFunc(scope.row)"
                    >
                    <el-icon><Plus /></el-icon>
                    </el-button>
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
        </div>
          <template #footer>
            <span class="dialog-footer">
              <el-button type="primary" @click="dialogVisible = false">管理保险条目</el-button>
              <el-button type="success" @click="dialogVisible = false">确认</el-button>
              <el-button type="info" @click="dialogVisible = false">取消</el-button>
            </span>
          </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { getInsuranceList, getPersonInsuranceList, addPersonInsurance, deletePersonInsurance } from '@/api/insuranceManage.js'
import { ElMessage } from 'element-plus';
import { useUserStore } from '@/pinia/modules/user'
import { storeToRefs  } from 'pinia'

const userStroe = useUserStore()
const userInfo = storeToRefs(userStroe).userInfo

const dialogVisible = ref(false)
const tableData = ref([])
const tableData1 = ref([])
const searchData = ref('')
const page = ref(0)
const pageSize = ref(20)
const total = ref(0)

onMounted(() => {
    getPersonInsuranceListFunc()
})

const Opendialog = () => {
    getInsuranceListFunc()
}

const getPersonInsuranceListFunc = () => {
    getPersonInsuranceList({id: Number(userInfo.value.employee_id)}).then(res => {
        if (res.code === 0) {
            tableData.value = res.data
        } else {
            ElMessage.warning(res.msg)
        }
    })
}

const getInsuranceListFunc = () => {
    getInsuranceList({
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchData.value
    }).then(res => {
        if (res.code === 0) {
            tableData1.value = res.data.data
            total.value = res.data.total
        } else {
            ElMessage.warning(res.msg)
        }
    })
}

const addInsuranceFunc = (row) => {
    addPersonInsurance({ employeeId: Number(userInfo.value.employee_id), insuranceId: row.id }).then(res => {
        if (res.code === 0) {
            ElMessage.success(res.msg)
        } else {
            ElMessage.warning(res.msg)
        }
    })
}

const handleClose = (down) => {
    down()
}

const deleteRowOfInsurance = (raw) => {
    deletePersonInsurance({ employeeId: Number(userInfo.value.employee_id), insuranceId: Number(raw.id) }).then(res => {
        if (res.code === 0) {  
            ElMessage.success(res.msg)
            getPersonInsuranceListFunc()
        } else {
            ElMessage.error(res.msg)
        }
    })
}
</script>

<style lang="scss" scoped>
    .insurance-list {
        height: 600px;
        width: 100%;
        .el-table {
            height: calc(100% - 64px);
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