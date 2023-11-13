<template>
    <div>
        <el-form ref="basicInfo" label-width="120px">
            <el-row>
                <el-divider content-position="left">职称</el-divider>
                <el-col :span="6">
                    <el-form-item label="职称：">
                        <el-select v-model="staff.appellation" clearable>
                            <el-option v-for="it in appellationList"></el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="职称证书编号：">
                        <el-input>{{ staff.appellationCode }}</el-input>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="职称获得日期：">
                        <el-date-picker v-model="staff.appellationTime" type="date"></el-date-picker>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="教学职称：">
                        <el-select v-model="staff.eduAppellation" clearable>
                            <el-option label="助教" value="助教"></el-option>
                            <el-option label="讲师" value="讲师"></el-option>
                            <el-option label="副教授" value="副教授"></el-option>
                            <el-option label="三级教授" value="三级教授"></el-option>
                            <el-option label="二级教授" value="二级教授"></el-option>
                            <el-option label="一级教授" value="一级教授"></el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="教学单位：">
                        <el-input v-model="staff.eduSchool" clearable></el-input>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="导师资质：">
                        <el-select v-model="staff.eduCategory" clearable>
                            <el-option label="博士后导师" value="1"></el-option>
                            <el-option label="博士生导师" value="2"></el-option>
                            <el-option label="硕士生导师" value="3"></el-option>
                            <el-option label="进修生带教" value="4"></el-option>
                            <el-option label="实习生带教" value="5"></el-option>
                            <el-option label="其他" value="99"></el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-divider content-position="left">资质</el-divider>
                <el-col :span="6">
                    <el-form-item label="执业类别:">
                        <el-select v-model="staff.practiceCateg" clearable>
                            <el-option label="临床" value="临床"></el-option>
                            <el-option label="检验" value="检验"></el-option>
                            <el-option label="放射" value="放射"></el-option>
                            <el-option label="药剂" value="药剂"></el-option>
                            <el-option label="护理" value="护理"></el-option>
                            <el-option label="工程技术" value="工程技术"></el-option>
                            <el-option label="康复治疗" value="康复治疗"></el-option>
                            <el-option label="研究实习" value="研究实习"></el-option>
                            <el-option label="其他" value="其他"></el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="执业医师资质:">
                        <span>{{ staff.physicianCert }}</span>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="执业助理医师资质:">
                        <span>{{ staff.physicianAidCert }}</span>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="执业药师资质:">
                        <span>{{ staff.pharmacistCert }}</span>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="执业护士资质:">
                        <span>{{ staff.nurseCert }}</span>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="操作资质:">
                        <el-select v-model="staff.operateCert" clearable>
                            <el-option label="X 线电子计算机断层扫描装置 (CT)医师" value="11"></el-option>
                            <el-option label="X 线电子计算机断层扫描装置 (CT)技师" value="12"></el-option>
                            <el-option label="医用磁共振成像装置(MRI)医师" value="21"></el-option>
                            <el-option label="医用磁共振成像装 (MRI) 技师" value="22"></el-option>
                            <el-option label="乳腺 X 线摄影装置技师" value="24"></el-option>
                            <el-option label="眼科准分子激光治疗装置(PRK/LASIK)医师" value="31"></el-option>
                            <el-option label="眼科准分子激光治疗装 (PRK/LASIK)技师" value="32"></el-option>
                            <el-option label="医用直线加速器 (LA) 医师" value="41"></el-option>
                            <el-option label="立体定向放射治疗装置( 刀、刀)技师" value="42"></el-option>
                            <el-option label="立体定向放治疗装(X 刀刀)物理师" value="43"></el-option>
                            <el-option label="彩色多普勒超声诊断装 (CDFI) 医师" value="51"></el-option>
                            <el-option label="彩色多音勒超声诊断装号(CDFI)技师" value="52"></el-option>
                            <el-option label="立体定向放射治疗装置 (X 刀刀)医师 (仅从事头部治疗的医师)" value="61"></el-option>
                            <el-option label="800 毫安以上数字减影血管造影X 光机(DSA) 技师" value="72"></el-option>
                            <el-option label="核医学影像装 (NIMl，包括 SPECT,PET) 医师" value="81"></el-option>
                            <el-option label="核医学影像装(NIMI，包括 SPECT,PET) 技师" value="82"></el-option>
                            <el-option label="核医学影像装(NIMl，包括SPECT,PET)物理师" value="83"></el-option>
                            <el-option label="核医学影像装(NIMI.包括SPECTPET)化学师" value="84"></el-option>
                            <el-option label="其他" value="99"></el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="手术资质:">
                        <span>{{ staff.surgryCert }}</span>
                        <el-select v-model="staff.surgryCert" clearable>
                            <el-option label="特大手术资质" value="01"></el-option>
                            <el-option label="中小手术资质" value="02"></el-option>
                            <el-option label="介入手术资质" value="03"></el-option>
                            <el-option label="腔镜手术资质" value="04"></el-option>
                            <el-option label="放射手术资质" value="05"></el-option>
                            <el-option label="其他" value="99"></el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="手术级别:">
                        <el-select v-model="staff.surgryGrade" clearable>
                            <el-option v-for="it in surgryGradeList" :label="it.label" :value="it.value"></el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-divider content-position="left">处方资质</el-divider>
                <div class="prescription-qualification" :style="{}">
                    <div class="table-show">
                        <el-table :data="tableData" :style="{ height: '100%',width: '100%' }">
                            <el-table-column prop="name" label="处方资质" align="center" />
                            <el-table-column prop="active" label="是否启用" align="center" />
                            <el-table-column width="200px" align="center">
                                <template #default="scope">
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
                    <div class="">
                        <el-button class="mt-4" style="width: 100%" @click="onAddItem">新增条目</el-button>
                    </div>
                </div>
            </el-row>
        </el-form>
        <el-dialog
          v-model="dialogVisible"
          title="新增条目"
          width="800px"
          :before-close="handleClose"
          @open="openDialog"
        >
          <div>
            <el-table :data="tableData1" :style="{ height: '100%',width: '100%' }">
                <el-table-column prop="name" label="处方资质" align="center" />
                <el-table-column prop="active" label="是否启用" align="center" />
                <el-table-column width="200px" align="center">
                    <template #default="scope">
                        <el-button icon="plus" @click="AddPPQFunc(scope.row)">add</el-button>
                    </template>
                </el-table-column>
            </el-table>
          </div>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { getNumberCode } from '@/api/common'
import { getListPQ, getPersonPQList, deletePersonPQ, addPersonPQ } from '@/api/prescriptionQualification'
import { useUserStore } from '@/pinia/modules/user'
import { storeToRefs  } from 'pinia'

const userStroe = useUserStore()
const userInfo = storeToRefs(userStroe).userInfo

onMounted(() => {
    getNumberCode('CV05.10.024').then(res => {
        surgryGradeList.value = res.data
    })
    getPersonPQList({ employeeId: Number(userInfo.value.employee_id) }).then(res => {
        console.log(res);
    })
})

const tableData = ref([])
const tableData1 = ref([])
const dialogVisible = ref(false)
const surgryGradeList = ref([])

const openDialog = () => {
    getListPQ({
        page: 1,
        pageSize: 99999
    }).then(res => {
        console.log(res);
        tableData1.value = res.data.data
    })
}

const props = defineProps(['staffInfo'])
const staff = computed(() => {
    return props.staffInfo
})
const deleteStaffInfoFunc = (row) => { }
const onAddItem = () => {
    dialogVisible.value = true
}
const AddPPQFunc = (obj) => {
    addPersonPQ({
        employeeId: Number(userInfo.value.employee_id),
        prescriptionId: obj.id,
    }).then(res => {

    })
}
</script>

<style lang="scss" scoped>
.prescription-qualification {
    height: 500px;
    width: 100%;
    .table-show {
        height: 460px;
    }
}
</style>