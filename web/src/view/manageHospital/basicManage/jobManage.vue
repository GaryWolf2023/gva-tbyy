<template>
    <div>
        <ManageTemplate
            :showCreate="true"
            :total="total"
            @create="createButtonFunc"
            @getList="getListFunc"
        >
            <el-table :data="tableData" style="width: 100%" height="100%">
                <el-table-column fixed prop="name" label="岗位名称" align="center" width="160px"/>
                <el-table-column label="是否可用" width="80">
                  <template #default="scope">
                    <el-switch
                      v-model="scope.row.active"
                      class="ml-2"
                      style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                    />
                  </template>
                </el-table-column>
                <el-table-column prop="baseNum" width="160px" label="所属部门" align="center" />
                <el-table-column prop="noOfRecruitment" label="招聘目标" align="center" />
                <el-table-column prop="noOfEmployee" label="现有员工数量" align="center" />
                <el-table-column prop="userId" label="招聘人员" align="center" />
                <el-table-column prop="jobCode" label="岗位代码" align="center" />
                <el-table-column prop="categoryPost" label="岗位类别" align="center" />
                <el-table-column prop="naturePost" label="岗位性质" align="center" />
                <el-table-column prop="show" label="岗位表述" align="center" />
                <el-table-column label="描述" align="center" width="240px">
                    <template #default="scope">
                    </template>
                </el-table-column>
                <el-table-column fixed="right" label="" width="80">
                  <template #default="scope">
                    <el-button link type="text" size="small" @click.prevent="editFunc(scope.row)">
                        <el-icon><Edit /></el-icon>
                    </el-button>
                    <el-button link type="text" size="small" @click.prevent="deleteRowOfInsurance(scope.row)">
                        <el-icon><CloseBold /></el-icon>
                    </el-button>
                  </template>
                </el-table-column>
            </el-table>
        </ManageTemplate>
        <el-dialog
          v-model="dialogVisible"
          title="新增职位"
          width="860px"
          top="10vh"
          :before-close="handleClose"
          @open="opendialog"
        >
        <div class="insurance-list">
            <el-form :model="form" label-width="120px">
                <el-row>
                    <el-col :span="24">
                        <el-form-item label="岗位名称">
                            <el-input v-model="formData.name" />
                        </el-form-item>
                    </el-col>
                    <br/>
                    <el-col :span="12">
                        <el-form-item label="部门">
                            <el-input v-model="formData.departmentId" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="应用">
                            <el-input v-model="formData.name" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="目标员工数量">
                              <el-input v-model="formData.NoOfRecruitment" type="number" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="现有员工数量">
                            <el-input v-model="formData.NoOfRecruitment" type="number" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="招聘人员">
                            <el-input v-model="formData.NoOfRecruitment" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="岗位代码">
                            <el-input v-model="formData.jobCode" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="岗位性质">
                            <el-select v-model="formData.naturePost">
                                <el-option>{{ '专职' }}</el-option>
                                <el-option>{{ '全职' }}</el-option>
                                <el-option>{{ '兼职' }}</el-option>
                                <el-option>{{ '院外聘请' }}</el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="岗位类别">
                            <el-select v-model="formData.categoryPost">
                                <el-option>{{ '管理' }}</el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="24">
                        <el-form-item label="岗位表述">
                            <el-input v-model="formData.show" type="textarea" :rows="3" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="8">
                        <el-form-item label="是否启用">
                            <el-switch
                                v-model="formData.active"
                                class="ml-2"
                                style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                            />
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
        </div>
          <template #footer>
            <span class="dialog-footer">
              <el-button type="success" @click="dialogVisible = false">确认</el-button>
              <el-button type="info" @click="dialogVisible = false">取消</el-button>
            </span>
          </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from 'vue'
import ManageTemplate from '@/components/manageTemplate/index.vue'
import { getJobList, createJob, updateJob, deleteJob } from '@/api/partTimePosition.js'
import { ElMessage } from 'element-plus';
import { Edit, CloseBold } from '@element-plus/icons-vue'

onMounted(() => {
    getOptionsList()
})

const total = ref(0)
const dialogVisible = ref(false)
const typeOfDialog = ref('')
const formData = ref({})
const tableData = ref([])

const openDialog = () => {
    if (typeOfDialog === 'create') {
        
    } else if (typeOfDialog === 'edit') {
        
    }
}
const getOptionsList = () => {

}
const clearForm = () => {
    formData.value = {}
}
const getListFunc = (obj) => {
    console.log(obj);
    getJobList(obj).then(res => {
        console.log(res);
        if (res.code === 0) {
            let name = {}
            total.value = res.data.total
            tableData.value = res.data.data.map(it => {
                name = JSON.parse(it.name)
                it.name = name.zh_CN
                return it
            })
            ElMessage.success(res.msg)
        } else {
            ElMessage.error(res.msg)
        }
    })
}
const createButtonFunc = () => {
    clearForm()
    dialogVisible.value = true
}
const editFunc = (val) => {
    formData.value = val
    dialogVisible.value = true
}
const createJobFunc = () => { }
const updateJobFunc = () => { }
</script>

<style lang="scss" scoped>

</style>