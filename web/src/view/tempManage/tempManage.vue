<template>
    <div class="temp-manage">
        <div class="temp-box">
            <div class="ctrl-temp">
                <el-button type="primary" plain @click="updateTempFunc">上 传 模 板</el-button>
                <el-button type="primary" plain @click="createTempFunc">创 建 模 板</el-button>
                <el-button type="primary" plain v-if="showEditorDialog" @click.="showEditorDialog=false">模 板 列 表</el-button>
                <el-input v-model="searchData" clearable :style="{width:'400px', marginLeft:'20px'}">
                    <template #append>
                        <el-button :icon="Search" @click="searchTempList" />
                    </template>
                </el-input>
                <el-button type="primary" plain @click="refreshTempListFunc" :style="{marginLeft:'20px'}">刷 新</el-button>
            </div>
            <div class="temp-table">
                <div class="show-editor" :style="{height:'100%'}" v-show="showEditorDialog">
                    <x-emr v-if="route.query.type ==='xemr'" :docname="docName" @getEditor="getEditor"></x-emr>
                    <create-form v-if="route.query.type ==='form'"></create-form>
                    <md-editor v-if="route.query.type ==='md'"></md-editor>
                </div>
                <el-table v-show="!showEditorDialog" :data="tableData" style="width: 100%;" height="100%">
                    <el-table-column prop="tempName" label="模板名称" align="center"/>
                    <el-table-column prop="createDate" label="创建时间" width="200" align="center"/>
                    <el-table-column prop="createPerson" label="创建人" width="160" align="center"/>
                    <el-table-column prop="updatePerson" label="更新" width="160" align="center"/>
                    <el-table-column prop="tempType" label="模板类型" width="160" align="center"/>
                    <el-table-column prop="active" label="是否可用" width="120" align="center">
                      <template #default="scope">
                          <el-switch
                            v-model="scope.row.active"
                            class="ml-2"
                            disabled
                            style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                          />
                      </template>
                    </el-table-column>
                    <el-table-column label="操作" fixed="right" width="180" align="center">
                        <template #default="scope">
                            <el-button link type="primary" size="small" :style="{padding:'0',lineHeight:'14px'}" @click="editPayloadTemp(scope.row.id)">编辑</el-button>
                            <el-popconfirm
                                width="220"
                                confirm-button-text="确认"
                                cancel-button-text="取消"
                                :icon="InfoFilled"
                                icon-color="#626AEF"
                                title="确定删除?"
                                @confirm="deletePayloadTempFunc(scope.row.id)"
                            >
                                <template #reference>
                                    <el-button link type="primary" size="small" :style="{padding:'0',lineHeight:'14px'}">删除</el-button>
                                </template>
                            </el-popconfirm>
                            <el-button link type="primary" size="small" :style="{padding:'0',lineHeight:'14px'}" @click="getFileFunc(scope.row.id)">查看</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
            <div class="page-ctrl">
                <el-pagination
                  v-model:current-page="currentPage"
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
        <el-dialog
            v-model="updateDialog"
            :title="dialogTitle==='create'?'创建模板':'编辑模板'"
            width="60%"
            :before-close="handleClose"
        >
            <div class="crate-dialog">
                <div>
                    <el-form 
                        :inline="true" 
                        :model="payloadTemp" 
                        class="demo-form-inline" 
                        label-position="left"
                        :rules="rules"
                        status-icon
                        ref="ruleFormRef"
                    >
                        <el-row :gutter="20">
                            <el-col :span="6" :offset="3">
                                <el-form-item label="模板名称" prop="tempName">
                                    <el-input v-model="payloadTemp.tempName" clearable required />
                                </el-form-item>
                            </el-col>
                            <el-col :span="6" :offset="3">
                                <el-form-item label="模板类型" prop="tempType">
                                    <el-select
                                        v-model="payloadTemp.tempType"
                                        clearable
                                        required
                                    >
                                        <el-option label="xemr" value="xemr" />
                                        <el-option label="form" value="form" />
                                        <el-option label="markdown" value="md" />
                                    </el-select>
                                </el-form-item>
                            </el-col>
                            <el-col :span="6" :offset="3">
                                <el-form-item label="病历类型" prop="tempContentType">
                                    <el-select
                                        v-model="payloadTemp.tempContentType"
                                        clearable
                                        required
                                    >
                                        <el-option label="文件" value="file" />
                                        <el-option label="文本" value="string" />
                                    </el-select>
                                </el-form-item>
                            </el-col>
                            <el-col :span="6" :offset="3">
                                <el-form-item label="病历是否可用">
                                    <el-switch
                                        class="ml-2"
                                        v-model="payloadTemp.active"
                                        style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                                    />
                                </el-form-item>
                            </el-col>
                        </el-row>
                        <el-row :gutter="20" v-if="dialogTitle==='create'">
                            <el-col :span="16" :offset="3">
                                <el-form-item label="病历文件" :style="{width:'100%',height:'200px'}" prop="file">
                                    <el-upload
                                      ref="uploadRef"
                                      :file-list="fileList"
                                      class="upload-demo update"
                                      action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
                                      :auto-upload="false"
                                      :accept="acceptFile"
                                      :limit="1"
                                      :on-exceed="handleExceed"
                                      :on-remove="removeFile"
                                      :on-preview="previewFile"
                                      :on-change="changeFile"
                                    >
                                      <template #trigger :style="{width:'100%'}">
                                        <el-button type="primary" plain :style="{width:'100%'}">选择文件</el-button>
                                      </template>
                                      <template #tip>
                                        <div class="el-upload__tip">
                                          上传单文件
                                        </div>
                                      </template>
                                    </el-upload>
                                </el-form-item>
                            </el-col>
                        </el-row>
                        <el-row :gutter="20" v-if="dialogTitle==='edit'">
                            <el-col :span="16" :offset="3">
                                <el-alert title="修改文件需单独上传" type="warning"  :style="{margin:'20px 0'}" :closable="false" />
                            </el-col>
                            <el-col :span="16" :offset="3">
                                <div>
                                    <span>存放位置：</span>
                                    <el-input v-model="tempContent.Bucket" disabled />
                                </div>
                                <div>
                                    <span>存放路径：</span>
                                    <el-input v-model="tempContent.Key" disabled />
                                </div>
                            </el-col>
                            <el-col :span="16" :offset="3">
                                <el-form-item label="病历文件" :style="{width:'100%',height:'200px'}">
                                    <el-upload
                                      ref="uploadRef"
                                      :file-list="fileList"
                                      class="upload-demo update"
                                      action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
                                      :auto-upload="false"
                                      :accept="acceptFile"
                                      :limit="1"
                                      :on-exceed="handleExceed"
                                      :on-remove="removeFile"
                                      :on-preview="previewFile"
                                      :on-change="changeFile"
                                    >
                                      <template #trigger :style="{width:'100%'}">
                                        <el-button type="primary" plain :style="{width:'100%', margin:'0px 60px 0px 0px'}">选择文件</el-button>
                                      </template>
                                      <el-button class="ml-3" plain type="success" @click="submitUpload">
                                        上传到服务器
                                      </el-button>
                                      <template #tip>
                                        <div class="el-upload__tip">
                                          单文件上传
                                        </div>
                                      </template>
                                    </el-upload>
                                </el-form-item>
                            </el-col>
                        </el-row>
                    </el-form>
                </div>
            </div>
            <template #footer>
              <span class="dialog-footer">
                <el-button @click="cancleDialog">取消</el-button>
                <el-button type="primary" @click="ctrlTemplate(dialogTitle)">
                    {{ dialogTitle==='create'?'创建':'更新' }}
                </el-button>
              </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import XEmr from '@/components/xemr/index.vue'
import CreateForm from '@/components/createForm/index.vue'
import MdEditor from '@/components/createForm/index.vue'

import { ElButton, ElMessageBox } from 'element-plus'
import { Search, InfoFilled } from '@element-plus/icons-vue'
import { ref, reactive, watch, onMounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { createPayloadTemp, updatePayloadTemp, deletePayloadTemp, getPayloadTempList, getPayloadTemp, updatePayloadOfFile, tempManageGetFile } from '@/api/payloadTemp'

const route = useRoute()
const router = useRouter()
const currentPage = ref(1)
const pageSize = ref(20)
const total = ref(0)
const searchData = ref('')
const tableData = ref([])
const dialogTitle = ref('创建模板')
const tempData = ref({})
const showEditorDialog = ref(false)
const docName = ref('')
const editor = ref({})
const ruleFormRef = ref(null)

const rules = ({
    tempName: [
        { required: true, message: '请填写名称', trigger: 'blur' }
    ],
    tempType: [
        { required: true, message: '请选择模板类型', trigger: 'blur' },
    ],
    tempContentType: [
        { required: true, message: '请选择病历类型', trigger: 'blur' },
    ]
})

const getEditor = (e) => {
    console.log(e);
    editor.value = e.editor
}

onMounted(() => {
    getPayloadTempListFunc()
})

const handleSizeChange = (a) => {
    pageSize.value = a
    getPayloadTempListFunc()
}
const handleCurrentChange = (b) => {
    page.value = b
    getPayloadTempListFunc()
}

const updateDialog = ref(false)
const updateTempFunc = () => {
    dialogTitle.value = "create"
    updateDialog.value = true
}
const clearDialog = () => {
    payloadTemp.active = true
    payloadTemp.tempContentType = ''
    payloadTemp.tempName = ''
    payloadTemp.tempType = ''
    checkFileList.value = []
    fileList.value = []
}
const handleClose = (down) => {
// 这里写dialog关闭的选择
    clearDialog()
    down()
}
const handleClose2 = (down) => {
    // 这里写dialog关闭的选择
    down()
}
const cancleDialog = () => {
    updateDialog.value = false
    dialogTitle.value = ""
}
const createTempFunc = () => {
    router.push({
        path: '/layout/tempmanage/createEmrTemp',
        query: {
            emrType: route.query.type
        }
    })
}
const payloadTemp = reactive({
    tempName: '',
    tempType: '',
    tempContentType: '',
    active: true,
})

const checkFileList = ref([])
const fileList = ref([])
const uploadRef = ref()
// 监听路由中携带参数变化
watch(() => route.query.type,
    (val) => {
        getPayloadTempListFunc()
    }
)
const acceptFile = ref('')
watch(
    () => payloadTemp.tempType,
    (val) => {
        acceptFile.value = val
     },
)
// 删除文件
const removeFile = (uploadFile, uploadFiles) => { }
// 点击已在文件列表中文件
const previewFile = (uploadFile) => {}
// 选择文件
const changeFile = (uploadFile, uploadFiles) => {
    console.log(uploadFile)
    console.log(fileList.value)
    checkFileList.value = uploadFiles
    console.log(checkFileList.value[0])
}
//文件超出限制后执行的钩子函数
const handleExceed = (files) => {
    console.log(files)
    checkFileList.value[0] = files[0]
}
// 控制按钮函数，是选择创建还是更新
const ctrlTemplate = (type) => {
    if (type === 'create') {
        createPayloadTempFunc()
    } else {
        updatePayloadTempFunc()
    }
}
// 创建模板并上传数据方法
const createPayloadTempFunc = () => {
    if (!ruleFormRef.value) return ElMessage.warning('error')
    if (checkFileList.value.length === 0) {
        return ElMessage.warning('请先选择文件')
    }
    ruleFormRef.value.validate(valid => {
        if (valid) {
            let tempData = new FormData()
            tempData.append("file", checkFileList.value[0].raw)

            for (const key in payloadTemp) {
                if (Object.hasOwnProperty.call(payloadTemp, key)) {
                    tempData.append(key, payloadTemp[key])
                }
            }
            console.log({ ...payloadTemp, createPerson: "唐纳德·特朗普", file: tempData })

            createPayloadTemp(tempData).then(res => {
                console.log(res);
                if (res.code === 0) {
                    ElMessage.success(res.msg)
                    getPayloadTempListFunc()
                }
            }).finally(() => {
                updateDialog.value = false
                clearDialog()
            })
        } else {
            ElMessage.warning('表单校验失败，请检查是否填写完整、正确')
        }
    })
    
}
const tempContent = ref(null)
// 获取模板文件，并展示
const getFileFunc = (id) => {
    let obj = {}
    tableData.value.forEach(it => {
        if (it.id === id) {
            obj = JSON.parse(it.tempContent)
        }
    })
    tempManageGetFile({
        bucketName: obj.Bucket,
        objectName: obj.Key
    }).then(res => {
        console.log(res)
        showEditorDialog.value = true
        editor.value.loadHtml(res.data)
    })
}
// 点击编辑模板
const editPayloadTemp = (id) => {
    getPayloadTemp(id).then(res => {
        if (res.code !== 0) {
            ElMessage.error('获取模板信息失败')
        } else {
            tempData.value = res.data
            tempContent.value = JSON.parse(tempData.value.tempContent)
            console.log(tempContent.value);
            console.log(tempData.value)
            dialogTitle.value = "edit"
            updateDialog.value = true
            for (const key in payloadTemp) {
                if (Object.hasOwnProperty.call(payloadTemp, key)) {
                    payloadTemp[key] = tempData.value[key]
                }
            }
        }
    })
}
const updatePayloadTempFunc = () => {
    console.log(ruleFormRef.value)
    if (!ruleFormRef.value) return
    ruleFormRef.value.validate(valid => {
        if (valid) {
            updatePayloadTemp({...payloadTemp, updatePerson: '张三', id: tempData.value.id}).then(res => {
                console.log(res)
                if (res.code !== 0) {
                    ElMessage.error("更新失败")
                }
                ElMessage.success('更新成功')
            })
        } else {
            ElMessage.warning('表单校验失败，请检查是否填写完整、正确')
        }
    })
    
}
// 编辑时单独上传文件
const submitUpload = () => {
    let formData = new FormData()
    console.log(tempContent.value)
    formData.append('id', tempData.value.id)
    formData.append('bucketName', tempContent.value.Bucket)
    formData.append('oldContent', JSON.stringify(tempContent.value))
    formData.append('payloadTemplate', JSON.stringify(tempData.value))
    formData.append('version', tempData.value.tempVersion)
    formData.append("file", checkFileList.value[0].raw)
    let fileName = checkFileList.value[0].name
    let arr = tempContent.value.Key.split('/')
    arr[arr.length - 1] = fileName
    formData.append('objectName', arr.join('/'))
    console.log(arr.join('/'));
    updatePayloadOfFile(formData).then(res => {
        ElMessage.warning(res.msg)
        if (res.code === 0) {
            getPayloadTempFunc(tempData.value.id)
        }
    })
}
// 点击删除模板
const deletePayloadTempFunc = (id) => {
    deletePayloadTemp({ id })
    .then(res => {
        console.log(res);
        ElMessage.success(res.msg)
        getPayloadTempListFunc()
    }).catch(e => {
        ElMessage.error('删除失败')
    })
}
// 检索方法
const searchTempList = () => {
    currentPage.value = 1
    getPayloadTempListFunc()
}
// 刷新列表
const refreshTempListFunc = () => {
    currentPage.value = 1
    pageSize.value = 20
    searchData.value = ''
    getPayloadTempListFunc()
}
// 获取模板列表方法
const getPayloadTempListFunc = () => {
    const type = route.query.type
    getPayloadTempList({
        type,
        searchData: searchData.value,
        page: currentPage.value,     
        pageSize: pageSize.value
    }).then(res => {
        // console.log(res)
        // ElMessage.success(res.msg)
        tableData.value = res.data
    })
}
// 获取单个模板
const getPayloadTempFunc = (id) => {
    console.log(id)
    getPayloadTemp(id).then(res => {
        if (res.code !== 0) {
            ElMessage.error('获取模板信息失败')
        } else {
            tempData.value = res.data
            tempContent.value = JSON.parse(tempData.value.tempContent)
            console.log(tempContent.value);
            console.log(tempData.value)
            dialogTitle.value = "edit"
            // updateDialog.value = true
            for (const key in payloadTemp) {
                if (Object.hasOwnProperty.call(payloadTemp, key)) {
                    payloadTemp[key] = tempData.value[key]
                }
            }
        }
    })
}
 
</script>

<style lang="scss" scoped>
.temp-manage {
    height: calc(100vh - 116px);
    /* margin-top: 110px; */
    .temp-box {
        height: 100%;
        width: calc(100% - 20px);
        margin: 0 10px;
        background-color: rgba(206, 206, 206, 0.486);
        padding: 0;
        border-radius: 10px;
        overflow: hidden;
        position: relative;
        .ctrl-temp {
            display: flex;
            justify-content: flex-start;
            align-items: center;
            padding: 0 18px;
            height: 60px;
            .create-temp {
                font-size: .5rem;
                line-height: 18px;
                padding: 5px 16px;
                margin-right: 20px;
                border: 1px solid rgba(24, 124, 255, 0.863);
                border-radius: 3px;
                color: rgba(24, 124, 255, 0.863);
                font-weight: bold;
                &:hover {
                    background-color: rgba(24, 124, 255, 0.863);
                    color: aliceblue;
                }
            }
        }
        .temp-table {
            height: calc(100% - 100px);
            position: relative;
            .show-editor {
                position: absolute;
                width: 100%;
                top: 0;
                left: 0;
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
    }
    .emr-box {
        height: 500px;
    }
    .crate-dialog {
        .file {
            .update {
                height: 200px;
                padding: 20px;
                width: 60%;
                margin: 0 20%;
                border: 1px dashed rgb(129, 129, 129);
            }
        }
    }
}
</style>