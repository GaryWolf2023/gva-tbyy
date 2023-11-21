<template>
    <div class="staff-info">
        <!-- <div class="staff-info-basic">
            <img src="" alt="暂无图片" class="staff-img">
            <div class="basic-staff-info">
                <el-form
                    ref="basicInfo"
                    label-width="120px"
                >
                    <el-row>
                        <el-col :span="6" :offset="1">
                            <el-form-item label="姓名：">
                                <span>{{ staffInfo.name }}</span>
                            </el-form-item>
                        </el-col>
                        <el-col :span="6" :offset="1">
                            <el-form-item label="职位：">
                                <span>{{ staffInfo.jobTitle }}</span>
                            </el-form-item>
                        </el-col>
                        <el-col :span="6" :offset="1">
                            <el-form-item label="移动手机：">
                                <span>{{ staffInfo.mobilePhone }}</span>
                            </el-form-item>
                        </el-col>
                        <el-col :span="6" :offset="1">
                            <el-form-item label="办公邮箱：">
                                <span>{{ staffInfo.workEmail }}</span>
                            </el-form-item>
                        </el-col>
                        <el-col :span="6" :offset="1">
                            <el-form-item label="办公电话：">
                                <span>{{ staffInfo.workPhone }}</span>
                            </el-form-item>
                        </el-col>
                        <el-col :span="6" :offset="1">
                            <el-form-item label="公司名称：">
                                <span>{{ staffInfo.companyId }}</span>
                            </el-form-item>
                        </el-col>
                        <el-col :span="3" :offset="1">
                            <el-form-item label="管理员：">
                                <span>{{ getInfoReturnObject(staffInfo.parentId, 'name') }}</span>
                            </el-form-item>
                        </el-col>
                        <el-col :span="3" :offset="1">
                            <el-form-item label="教练：">
                                <span>{{ getInfoReturnObject(staffInfo.coachId, 'name') }}</span>
                            </el-form-item>
                        </el-col>
                        <el-col :span="6" :offset="18">
                            <el-button :style="{width:'100px'}" type="primary">编辑</el-button>
                        </el-col>
                    </el-row>
                </el-form>
            </div>
        </div> -->
        <div class="staff-info-plan">
            <div class="info-item">
                <p class="info-item-title"> 基础信息:</p>
                <div class="ctrl-info">
                    <el-button type="warning" size="small" :icon="Refresh" circle />
                    <!-- <el-button type="primary" size="small">编辑</el-button> -->
                    <el-button type="success" size="small" @click="saveData('basicInfo')">保存</el-button>
                </div>
                <InfoForm ref="infoForm" :staffInfo="staffInfo"></InfoForm>
            </div>
            <div class="info-item">
                <p class="info-item-title">系统配置:</p>
                <div class="ctrl-info">
                    <el-button type="warning" size="small" :icon="Refresh" circle />
                    <el-button type="primary" v-if="!editorSystem" size="small" @click="editorSystem=true">编辑</el-button>
                    <el-button type="success" v-else size="small" @click="saveData('systemConfiguration')">保存</el-button>
                </div>
                <div class="info-item-form">
                    <el-form
                        ref="basicInfo"
                        label-width="120px"
                    >
                        <el-row>
                            <el-divider content-position="left">预约人数配置</el-divider>
                            <el-col :span="6">
                                <el-form-item label="上午预约人数：">
                                    <span v-if="!editorSystem">{{ staffInfo.bookingNumAm }}</span>
                                    <el-input v-else v-model="staffInfo.bookingNumAm" />
                                </el-form-item>
                            </el-col>
                            <el-col :span="6">
                                <el-form-item label="下午预约人数：">
                                    <span v-if="!editorSystem">{{ staffInfo.bookingNumPm }}</span>
                                    <el-input v-else v-model="staffInfo.bookingNumPm" />
                                </el-form-item>
                            </el-col>
                            <el-col :span="6">
                                <el-form-item label="夜间预约人数：">
                                    <span v-if="!editorSystem">{{ staffInfo.bookingNumNg }}</span>
                                    <el-input v-else v-model="staffInfo.bookingNumNg" />
                                </el-form-item>
                            </el-col>
                            <el-divider content-position="left">诊室展示屏配置</el-divider>
                            <el-col :span="18" :offset="1">
                                <el-form-item label="医生简介:">
                                    <span v-if="!editorSystem">{{ staffInfo.introduction }}</span>
                                    <el-input
                                        type="textarea"
                                        :autosize="{ minRows: 2, maxRows: 4 }"
                                        v-else
                                        v-model="staffInfo.introduction"
                                    />
                                </el-form-item>
                            </el-col>
                            <el-col :span="16">
                                <el-form-item label="擅长疾病配置:">
                                    <div :style="{width:'100%'}">
                                        <AdeptIllness :id="route.query.id"></AdeptIllness>
                                    </div>
                                </el-form-item>
                            </el-col>
                        </el-row>
                    </el-form>
                </div>
            </div>
            <div class="info-item">
                <p class="info-item-title">员工保险配置:</p>
                <div class="info-item-form">
                    <InsuranceAllocation :staffInfo="staffInfo"></InsuranceAllocation>
                </div>
            </div>
            <div class="info-item">
                <p class="info-item-title">职称资质:</p>
                <div class="ctrl-info">
                    <el-button type="warning" size="small" :icon="Refresh" circle />
                    <el-button type="primary" size="small">编辑</el-button>
                    <el-button type="success" size="small">保存</el-button>
                </div>
                <div class="info-item-form">
                   <TitleManage :staffInfo="staffInfo"></TitleManage>
                </div>
            </div>
            <div class="info-item">
                <p class="info-item-title">兼职信息:</p>
                <div class="ctrl-info">
                    <el-button type="warning" size="small" :icon="Refresh" circle />
                    <el-button type="primary" size="small">编辑</el-button>
                    <el-button type="success" size="small">保存</el-button>
                </div>
                <div class="info-item-form">
                    <PartTimePosition :staffInfo="staffInfo"></PartTimePosition>
                </div>
            </div>
            <div class="info-item">
                <p class="info-item-title">简历:</p>
                <div class="ctrl-info">
                    <el-button type="warning" size="small" :icon="Refresh" circle />
                    <el-button type="primary" size="small">编辑</el-button>
                    <el-button type="success" size="small">保存</el-button>
                </div>
                <div class="info-item-form">
                    <el-divider content-position="left">技能证书</el-divider>
                    <el-table :data="skillCertificateData" style="width: 100%">
                      <el-table-column prop="date" label="Date" />
                      <el-table-column prop="name" label="Name" />
                      <el-table-column prop="address" label="Address" />
                    </el-table>
                </div>
            </div>
            <div class="info-item">
                <p class="info-item-title">工作信息:</p>
                <div class="ctrl-info">
                    <el-button type="warning" size="small" :icon="Refresh" circle />
                    <el-button type="primary" size="small">编辑</el-button>
                    <el-button type="success" size="small">保存</el-button>
                </div>
                <div class="info-item-form">
                    <el-form
                        ref="basicInfo"
                        label-width="120px"
                    >
                        <el-row>
                            <el-divider content-position="left">位置</el-divider>

                            <el-divider content-position="left">审批人</el-divider>

                            <el-divider content-position="left">安排</el-divider>

                            <el-divider content-position="left">组织图表</el-divider>
                        </el-row>
                    </el-form>
                    
                </div>
            </div>
            <div class="info-item">
                <p class="info-item-title">隐私信息:</p>
                <div class="ctrl-info">
                    <el-button type="warning" size="small" :icon="Refresh" circle />
                    <el-button type="primary" size="small">编辑</el-button>
                    <el-button type="success" size="small">保存</el-button>
                </div>
                <div class="info-item-form">
                    <el-form
                        ref="basicInfo"
                        label-width="120px"
                    >
                        <el-row>
                            <el-divider content-position="left">私人联系方式</el-divider>
                    
                            <el-divider content-position="left">身份信息</el-divider>

                            <el-divider content-position="left">护照信息</el-divider>

                            <el-divider content-position="left">家庭情况</el-divider>
                        </el-row>
                    </el-form>
                </div>
            </div>
            <div class="info-item">
                <p class="info-item-title">HR设置:</p>
                <div class="ctrl-info">
                    <el-button type="warning" size="small" :icon="Refresh" circle />
                    <el-button type="primary" size="small">编辑</el-button>
                    <el-button type="success" size="small">保存</el-button>
                </div>
                <div class="info-item-form">
                    <el-form
                        ref="basicInfo"
                        label-width="120px"
                    >
                        <el-row>
                            <el-divider content-position="left">员工状态</el-divider>
                    
                            <el-divider content-position="left">考勤/销售</el-divider>

                            <el-divider content-position="left">工资册</el-divider>

                            <el-divider content-position="left">护士能级</el-divider>

                            <el-divider content-position="left">应用程序设置</el-divider>
                        </el-row>
                    </el-form>
                </div>
            </div>
            <div class="info-item">
                <p class="info-item-title">员工历史记录:</p>
                <div class="info-item-form">
                    <el-button type="primary" plain bg :icon="Document">作业历史记录</el-button>
                    <el-button type="primary" plain bg :icon="Document">合同历史</el-button>
                    <el-button type="primary" plain bg :icon="Document">薪资历史记录</el-button>
                    <el-button type="primary" plain bg :icon="Document">时间表成本</el-button>
                    <el-button type="primary" plain bg :icon="Document">职称历史</el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import InfoForm from './infoForm.vue'
import InsuranceAllocation from '../othorAllocation/insuranceAllocation.vue'
import AdeptIllness from '../othorAllocation/adeptIllness.vue'
import TitleManage from './titleManage.vue'
import PartTimePosition from './partTimePosition.vue'
import { useRoute } from 'vue-router'
import { ref, computed, watch, onMounted } from 'vue'
import { Document, Refresh } from '@element-plus/icons-vue'
import { getStaffList, getStaffInfo, deleteStaff, updateStaff, updateStaffSystemConfig } from '@/api/staff.js'

const route = useRoute()
const staffInfo = ref({})
const page = ref(1)
const pageSize = ref(20)
const staffList = ref([])
const insuranceData = ref([]) // 保险列表
const partTimeJobData = ref([]) // 简直列表
const skillCertificateData = ref([])
const formData = ref(new FormData())
const infoForm = ref(null) // 子组件

const editorSystem = ref(false)

watch(() => route.query.id, (a) => {
    console.log(a);
    if (a) {
        getStaffInfo(a).then(res => {
        console.log(res);
            if (res.code === 0) { 
            console.log("员工详细信息：", res.data)
            staffInfo.value = res.data
        } else {
            ElMessage.warning("获取个人信息失败")
        }
    })
    }
}, { immediate: true })
const getInfoReturnObject = (id, key) => {
    if (!id) {
        return ''
    }
    let data = getStaffInfoFunc(id)
    console.log(data);
    return data[key]
}
const getStaffInfoFunc = async (id) => {
    const res = await getStaffInfo(id)
    console.log(res);
    return res.data
}
const getStaffListFunc = (str) => {
    getStaffList({
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchData.value
    }).then(res => {
        staffList.value = res.data.data
        total.value = res.data.total
    })
}
const onAddInsuranceFunc = () => {
    console.log('新增保险')
}
const deleteRowOfInsurance = () => {}
const onAddPartTimeJobFunc = () => {}
const deleteRowOfPartTimeJob = () => { }
const saveData = (type) => {
    switch (type) {
        case 'basicInfo':
            updateBasicInfo()
            break;
        case 'systemConfiguration':
            updateSystemConfiguration()
            break;
        default:
            break;
    }
}
const updateBasicInfo = () => {
    infoForm.value.updateData(route.query.id)
}
const updateSystemConfiguration = () => {
    var obj = {
        id: Number(route.query.id),
        bookingNumAm: Number(staffInfo.value.bookingNumAm),
        bookingNumPm:  Number(staffInfo.value.bookingNumPm),
        bookingNumNg:  Number(staffInfo.value.bookingNumNg),
        introduction:  staffInfo.value.introduction
    }
    console.log(obj);
    updateStaffSystemConfig(obj).then(res => {
        console.log(res)
    })
}
</script>

<style lang="scss" scoped>
.staff-info {
    height: calc(100vh - 116px);
    .staff-info-basic {
        height: 240px;
        width: calc(100% - 40px);
        padding: 0 20px;
        display: flex;
        align-items: center;
        justify-content: flex-start;
        .staff-img {
            height: 220px;
            width: 220px;
            border: 1px solid rgb(121, 121, 121);
            border-radius: 10px;
        }
        .basic-staff-info {
            width: calc(100% - 240px);
            height: calc(100% - 20px);
            padding-top: 10px;
        }
    }
    .staff-info-plan {
        height: calc(100%);
        padding: 0 40px;
        overflow-x: hidden;
        .info-item {
            min-height: 180px;
            position: relative;
            .ctrl-info {
                position: absolute;
                right: 30px;
                top: 0;
                bottom: 0;
                .refresh {
                    margin: 0 10px;
                    margin-top: 6px;
                    font-size: 18px;
                    line-height: 38px;
                }
            }
            .info-item-title {
                font-size: 16px;
                font-weight: bold;
                color: rgb(78, 78, 78);
                line-height: 36px;
                border-bottom: 2px solid rgba(46, 46, 46, 0.568);
                padding-left: 24px;
            }
            .info-item-form {
                padding: 16px 20px;
                width: calc(100% - 60px);
                margin: 10px;
                border: 2px solid rgba(207, 207, 207, 0.932);
                border-radius: 10px;
            }
        }
    }
}
</style>