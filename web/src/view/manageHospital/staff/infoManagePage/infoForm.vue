
<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="right" label-width="101px" size="medium">
    <el-row>
      <el-col :span="24" class="grid-cell">
        <div class="static-content-item">
          <el-divider direction="horizontal" content-position="left"></el-divider>
        </div>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="姓名" prop="name" class="required">
          <el-input v-model="formData.name" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="职位" prop="jobTitle" class="required">
          <el-input v-model="formData.jobTitle" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="移动电话" prop="mobilePhone" class="required">
          <el-input v-model="formData.mobilePhone" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="办公邮箱" prop="workEmail" class="required">
          <el-input v-model="formData.workEmail" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="办公电话" prop="workPhone" class="required">
          <el-input v-model="formData.workPhone" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="民族" prop="nation" class="required">
          <el-select v-model="formData.nation" class="full-width-input" clearable filterable automatic-dropdown remote>
            <el-option v-for="(item, index) in nationOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="国籍" prop="countryId" class="required">
          <el-select v-model="formData.countryId" class="full-width-input" clearable>
            <el-option v-for="(item, index) in countryIdOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="性别" prop="gender" class="required">
          <el-select v-model="formData.gender" class="full-width-input">
            <el-option v-for="(item, index) in genderOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="出生日期" prop="birthday" class="required">
          <el-date-picker v-model="formData.birthday" type="date" class="full-width-input"  format="YYYY-MM-DD" value-format="YYYY-MM-DD" :default-time="defaultTime" clearable></el-date-picker>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="出生地" prop="placeOfBirth" class="required">
          <el-input v-model="formData.placeOfBirth" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="出生国" prop="countryOfBirth" class="required">
          <el-select v-model="formData.countryOfBirth" class="full-width-input" clearable>
            <el-option v-for="(item, index) in countryOfBirthOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="籍贯（省）" prop="nativePlaceP" class="required">
          <el-select v-model="formData.nativePlaceP" class="full-width-input" clearable>
            <el-option v-for="(item, index) in nativePlacePOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="籍贯（市）" prop="nativePlaceC" class="required">
          <el-select v-model="formData.nativePlaceC" class="full-width-input" clearable>
            <el-option v-for="(item, index) in nativePlaceCOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="政治面貌" prop="politicsStatus" class="required">
          <el-select v-model="formData.politicsStatus" class="full-width-input" clearable>
            <el-option v-for="(item, index) in politicsStatusOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="24" class="grid-cell">
        <div class="static-content-item">
          <el-divider direction="horizontal" content-position="left"></el-divider>
        </div>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="居住地址" prop="addressHomeId">
          <el-input v-model="formData.addressHomeId" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <!-- <el-col :span="6" class="grid-cell">
        <el-form-item label="私人邮箱" prop="privateEmail" class="required">
          <el-input v-model="formData.privateEmail" type="text" clearable></el-input>
        </el-form-item>
      </el-col> -->
      <!-- <el-col :span="6" class="grid-cell">
        <el-form-item label="私人电话" prop="phone" class="required">
          <el-input v-model="formData.phone" type="text" clearable></el-input>
        </el-form-item>
      </el-col> -->
      <!-- <el-col :span="6" class="grid-cell">
        <el-form-item label="银行卡账户" prop="bankAccountId" class="required">
          <el-input v-model="formData.bankAccountId" type="text" clearable></el-input>
        </el-form-item>
      </el-col> -->
      <el-col :span="6" class="grid-cell">
        <el-form-item label="家到公司距离" prop="kmHomeWork">
          <el-input v-model="formData.kmHomeWork" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="身份证号" prop="identificationId" class="required">
          <el-input v-model="formData.identificationId" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="到期时间" prop="idExpiryDate" class="required">
          <el-date-picker v-model="formData.idExpiryDate" type="date" class="full-width-input" format="YYYY-MM-DD" :default-time="defaultTime"
            value-format="YYYY-MM-DD" clearable></el-date-picker>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="护照号" prop="passportId">
          <el-input v-model="formData.passportId" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" :offset="6" class="grid-cell">
        <el-form-item label="身份证附件" prop="idAttachmentId">
          <el-upload :file-list="idAttachmentIdFileList" :headers="idAttachmentIdUploadHeaders"
            :data="idAttachmentIdUploadData" list-type="picture-card" show-file-list :limit="2">
            <template #default>
              <i class="el-icon-plus"></i>
            </template>
          </el-upload>
        </el-form-item>
      </el-col>
      <el-col :span="24" class="grid-cell">
        <div class="static-content-item">
          <el-divider direction="horizontal" content-position="left"></el-divider>
        </div>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="婚姻状况" prop="marital">
          <el-select v-model="formData.marital" class="full-width-input" clearable>
            <el-option v-for="(item, index) in maritalOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="配偶全名" prop="spouseCompleteName">
          <el-input v-model="formData.spouseCompleteName" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="配偶生日" prop="spouseBirthdate">
          <el-date-picker v-model="formData.spouseBirthdate" type="date" class="full-width-input" format="YYYY-MM-DD" :default-time="defaultTime"
            value-format="YYYY-MM-DD" clearable></el-date-picker>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="受抚养子女数" prop="chlidren">
          <el-input v-model="formData.children" type="number" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="24" class="grid-cell">
        <div class="static-content-item">
          <el-divider direction="horizontal" content-position="left"></el-divider>
        </div>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="联系人姓名" prop="emergencyContact" class="required">
          <el-input v-model="formData.emergencyContact" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="紧急联系方式" prop="emergencyPhone" class="required">
          <el-input v-model="formData.emergencyPhone" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="24" class="grid-cell">
        <div class="static-content-item">
          <el-divider direction="horizontal" content-position="left"></el-divider>
        </div>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="学位" prop="certificate" class="required">
          <el-select v-model="formData.certificate" class="full-width-input" clearable>
            <el-option v-for="(item, index) in certificateOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="学历" prop="schoolgrade" class="required">
          <el-select v-model="formData.schoolGrade" class="full-width-input" clearable>
            <el-option v-for="(item, index) in schoolgradeOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="毕业时间" prop="graduationDate" class="required">
          <el-date-picker v-model="formData.graduationDate" type="date" class="full-width-input" format="YYYY-MM-DD" :default-time="defaultTime"
            value-format="YYYY-MM-DD" clearable></el-date-picker>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="研究领域" prop="studyField" class="required">
          <el-input v-model="formData.studyField" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="毕业院校" prop="studySchool" class="required">
          <el-input v-model="formData.studySchool" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="24" class="grid-cell">
        <div class="static-content-item">
          <el-divider direction="horizontal" content-position="left"></el-divider>
        </div>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="签证号" prop="visaNo">
          <el-input v-model="formData.visaNo" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="签证编号" prop="permitNo">
          <el-input v-model="formData.permitNo" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="签证到期日期" prop="visaExpire">
          <el-date-picker v-model="formData.visaExpire" type="date" class="full-width-input" format="YYYY-MM-DD" :default-time="defaultTime"
            value-format="YYYY-MM-DD" clearable></el-date-picker>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="工作许可证日期" prop="workPermitExpiration">
          <el-date-picker v-model="formData.workPermitExpiration" type="date" class="full-width-input" format="YYYY-MM-DD" :default-time="defaultTime"
            value-format="YYYY-MM-DD" clearable></el-date-picker>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell">
        <el-form-item label="许可证文件图片" prop="hasWorkPermit">
          <el-upload :file-list="hasWorkPermitFileList" :headers="hasWorkPermitUploadHeaders"
            :data="hasWorkPermitUploadData" list-type="picture-card" show-file-list :limit="5"> <template #default><i
                class="el-icon-plus"></i></template></el-upload>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>
  
<script setup>
import { reactive, ref, onMounted, watch } from 'vue'
import { getNumberCode, getArea, getCity, getCountry, getProvince } from '@/api/common.js'
import { updateStaff } from '@/api/staff'
import { ElMessage } from 'element-plus';

const vForm = ref(null)
const props = defineProps(["staffInfo"])

const formData = reactive({
  name: "",
  jobTitle: "",
  mobilePhone: "",
  workEmail: "",
  workPhone: "",
  nation: "",
  countryId: null,
  gender: null,
  birthday: null,
  placeOfBirth: "",
  countryOfBirth: null,
  nativePlaceP: null,
  nativePlaceC: null,
  politicsStatus: "",
  addressHomeId: "",
  // privateEmail: "",
  // phone: "",
  // bankAccountId: "",
  kmHomeWork: null,
  identificationId: "",
  idExpiryDate: null,
  passportId: "",
  marital: null,
  spouseCompleteName: "",
  spouseBirthdate: null,
  children: null,
  emergencyContact: "",
  emergencyPhone: "",
  certificate: null,
  schoolGrade: null,
  graduationDate: null,
  studyField: "",
  studySchool: "",
  visaNo: "",
  permitNo: "",
  visaExpire: null,
})
const rules = reactive({
  name: [{
    required: true,
    message: '字段值不可为空',
  }],
  jobTitle: [{
    required: true,
    message: '字段值不可为空',
  }],
  mobilePhone: [{
    required: true,
    message: '字段值不可为空',
  }],
  workEmail: [{
    required: true,
    message: '字段值不可为空',
  }],
  workPhone: [{
    required: true,
    message: '字段值不可为空',
  }],
  nation: [{
    required: true,
    message: '字段值不可为空',
  }],
  countryId: [{
    required: true,
    message: '字段值不可为空',
  }],
  gender: [{
    required: true,
    message: '字段值不可为空',
  }],
  birthday: [{
    required: true,
    message: '字段值不可为空',
  }],
  placeOfBirth: [{
    required: true,
    message: '字段值不可为空',
  }],
  countryOfBirth: [{
    required: true,
    message: '字段值不可为空',
  }],
  nativePlaceP: [{
    required: true,
    message: '字段值不可为空',
  }],
  nativePlaceC: [{
    required: true,
    message: '字段值不可为空',
  }],
  politicsStatus: [{
    required: true,
    message: '字段值不可为空',
  }],
  privateEmail: [{
    required: true,
    message: '字段值不可为空',
  }],
  phone: [{
    required: true,
    message: '字段值不可为空',
  }],
  bankAccountId: [{
    required: true,
    message: '字段值不可为空',
  }],
  indentificationId: [{
    required: true,
    message: '字段值不可为空',
  }],
  idExpiryDate: [{
    required: true,
    message: '字段值不可为空',
  }],
  emergencyContact: [{
    required: true,
    message: '字段值不可为空',
  }],
  emergencyPhone: [{
    required: true,
    message: '字段值不可为空',
  }],
  certificate: [{
    required: true,
    message: '字段值不可为空',
  }],
  schoolgrade: [{
    required: true,
    message: '字段值不可为空',
  }],
  graduationDate: [{
    required: true,
    message: '字段值不可为空',
  }],
  studyField: [{
    required: true,
    message: '字段值不可为空',
  }],
  studySchool: [{
    required: true,
    message: '字段值不可为空',
  }],
})
const nationOptions = ref([])
const countryIdOptions = ref([])
const genderOptions = ref([])
const countryOfBirthOptions = ref([])
const nativePlacePOptions = ref([])
const nativePlaceCOptions = ref([])
const politicsStatusOptions = ref([
  {label: "中共党员",value:"1"},
  {label: "中共预备党员",value:"2"},
  {label: "共青团员",value:"3"},
  {label: "民革党员",value:"4"},
  {label: "民盟盟员",value:"5"},
  {label: "民建会员",value:"6"},
  {label: "民进会员",value:"7"},
  {label: "农工党党员",value:"8"},
  {label: "致公党党员",value:"9"},
  {label: "九三学社社员",value:"10"},
  {label: "台盟盟员",value:"11"},
  {label: "无党派人士",value:"12"},
  {label: "群众",value:"13"}
])
const maritalOptions = ref([])
const certificateOptions = ref([])
const schoolgradeOptions = ref([])
const idAttachmentIdFileList = ref([])
const idAttachmentIdUploadHeaders = ref([])
const idAttachmentIdUploadData = ref([])
const hasWorkPermitFileList = ref([])
const hasWorkPermitUploadHeaders = ref([])
const hasWorkPermitUploadData = ref([])

const defaultTime = ref<[Date, Date]>([
  new Date(2000, 1, 1, 0, 0, 0),
  new Date(2000, 2, 1, 23, 59, 59),
])

const updateData = (id) => {
  updateStaff({ id: Number(id), ...formData }).then(res => {
    if (res.code === 0) {
      ElMessage.success(res.msg)
    } else {
      ElMessage.warning(res.msg)
    }
  })
}
onMounted(async () => {
  getCountry().then(res => {
    res.data.forEach((it, index) => {
      let obj = JSON.parse(it.label)
      countryIdOptions.value[index] = {
        label: obj["zh_CN"],
        value: it.value
      }
      countryOfBirthOptions.value[index] = {
        label: obj["zh_CN"],
        value: it.value
      }
    })
  })
  // 获取省市去
  const res1 = await getNumberCode("GB/T 2261.1-2003")
  genderOptions.value = res1.data
  const res2 = await getNumberCode("GB/T 4658-2006")
  schoolgradeOptions.value = res2.data
  const res3 = await getNumberCode("GB/T 2261.2-2003")
  res3.data.forEach((it, index) => {
    maritalOptions.value[index] = {
        label: it.label,
        value: String(it.value)
      }
  })
  const res4 = await getNumberCode("GB/T 3304-1991")
  nationOptions.value = res4.data
  const res5 = await getNumberCode("GB/T 6864-2003")
  certificateOptions.value = res5.data
})
watch(() => props.staffInfo, (val) => {
  if (val) {
    for (const key in formData) {
      formData[key] = val[key]
    }
  }
},{immediate: true})
watch(
  () => formData.countryOfBirth,
  (val, oldval) => {
    if (!!val) {
      getProvince(val).then(res => {
        nativePlacePOptions.value = res.data
      })
    }
    if (!!oldval) {
      formData.nativePlaceP = null
    }
  }, {immediate: true})
watch(
  () => formData.nativePlaceP,
  (val,oldval) => {
    formData.nativePlaceC = null
    getCity(val).then(res => {
      nativePlaceCOptions.value = res.data
    })
    if (!!oldval) {
      formData.nativePlaceC = null
    }
}, {immediate: true})
defineExpose({
  updateData
})
</script>
  
<style lang="scss">
.el-input-number.full-width-input,
.el-cascader.full-width-input {
  width: 100% !important;
}

.el-form-item--medium {
  .el-radio {
    line-height: 36px !important;
  }

  .el-rate {
    margin-top: 6px;
  }
}

.el-form-item--small {
  .el-radio {
    line-height: 32px !important;
  }

  .el-rate {
    margin-top: 6px;
  }
}

.el-form-item--mini {
  .el-radio {
    line-height: 26px !important;
  }

  .el-rate {
    margin-top: 4px;
  }
}

.clear-fix:before,
.clear-fix:after {
  display: table;
  content: "";
}

.clear-fix:after {
  clear: both;
}

.float-right {
  float: right;
}
</style>
  
<style lang="scss" scoped>
div.table-container {
  table.table-layout {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;

    td.table-cell {
      display: table-cell;
      height: 36px;
      border: 1px solid #e1e2e3;
    }
  }
}

div.tab-container {}

.label-left-align :deep(.el-form-item__label) {
  text-align: left;
}

.label-center-align :deep(.el-form-item__label) {
  text-align: center;
}

.label-right-align :deep(.el-form-item__label) {
  text-align: right;
}

.custom-label {}

.static-content-item {
  min-height: 20px;
  display: flex;
  align-items: center;

  :deep(.el-divider--horizontal) {
    margin: 0;
  }
}
</style>