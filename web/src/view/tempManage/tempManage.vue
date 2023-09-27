<template>
    <!-- <div class="update-box">
        <el-button type="primary" size="small" plain>上传</el-button>
    </div> -->
    <div class="temp-box">
        <div class="ctrl-temp">
            <span class="create-temp" @click="updateTempFunc">上 传 模 板</span>
            <span class="create-temp" @click="createTempFunc">创 建 模 板</span>
        </div>
        <el-table :data="tableData" style="width: 100%;" height="100%">
          <el-table-column prop="tempName" label="模板名称" width="400" align="center"/>
          <el-table-column prop="createDate" label="创建时间" width="300" align="center"/>
          <el-table-column prop="updateDate" label="更新时间" width="300" align="center"/>
          <el-table-column prop="tempType" label="模板类型" align="center"/>
          <el-table-column prop="active" label="是否可用" align="center">
            <template #default>
                <el-switch
                  v-model="value2"
                  class="ml-2"
                  style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                />
            </template>
          </el-table-column>
          <el-table-column label="操作" fixed="right" width="160" align="center">
            <template #default>
                <el-button link type="primary" size="small" @click="handleClick">Detail</el-button>
                <el-button link type="primary" size="small">Edit</el-button>
            </template>
          </el-table-column>
        </el-table>
    </div>
    <el-dialog
        v-model="updateDialog"
        title="创建模板"
        width="60%"
        :before-close="handleClose"
    >
        <div class="crate-dialog">
            <div :style="{height:'100px'}">
                <el-form :inline="true" :model="payloadTemp" class="demo-form-inline" label-position="left">
                    <el-form-item label="模板名称">
                        <el-input v-model="payloadTemp.tempName" clearable />
                    </el-form-item>
                    <el-form-item label="模板类型">
                        <el-select
                            v-model="payloadTemp.tempType"
                            clearable
                        >
                            <el-option label="xemr" value="html" />
                            <el-option label="form" value="txt" />
                            <el-option label="markdown" value="md" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="病历类型">
                        <el-select
                            v-model="payloadTemp.tempContentType"
                            clearable
                        >
                            <el-option label="文件" value="file" />
                            <el-option label="文本" value="string" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="病历是否可用">
                        <el-switch
                            v-model="payloadTemp.active"
                            class="ml-2"
                            style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                        />
                    </el-form-item>
                </el-form>
            </div>
            <div class="file" :style="{height:'360px'}">
                <p :style="{lineHeight:'36px'}">文件上传：</p>
                <el-upload
                  ref="uploadRef"
                  class="upload-demo update"
                  action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
                  :auto-upload="false"
                  :accept="acceptFile"
                  :on-remove="removeFile"
                  :on-preview="previewFile"
                  :on-change="changeFile"
                >
                  <template #trigger :style="{width:'100%'}">
                    <el-button type="primary" plain :style="{width:'100%'}">选择文件</el-button>
                  </template>
                  <!-- <el-button class="ml-3" type="success" @click="submitUpload">
                    upload to server
                  </el-button> -->
                  <template #tip>
                    <div class="el-upload__tip">
                      文件大小不能超过 500kb
                    </div>
                  </template>
                </el-upload>
            </div>
        </div>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="dialogVisible = false">
                上传模板
            </el-button>
          </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { ElButton } from 'element-plus'
import { ref, reactive, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()

const updateDialog = ref(false)
const updateTempFunc = () => {
    updateDialog.value = true
}
const handleClose = (down) => {
// 这里写dialog关闭的选择
    down()
}
const createTempFunc = () => {
    router.push({
        path: '/layout/createEmrTemp',
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
    file: null
})
const updateFile = ref()
watch(() => {
  if (updateFile.value) {
    updateFile.value
  } else {
    // 此时还未挂载，或此元素已经被卸载（例如通过 v-if 控制）
  }
},
(val) => {
    console.log(val);
    let formData = new FormData();
    formData.append('file', file);
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
const previewFile = (uploadFile) => { }
// 选择文件
const changeFile = (uploadFile, uploadFile) => {}
</script>

<style lang="scss" scoped>
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

.update-box {
    position: fixed;
    top: 100px;
    width: calc(100% - 40px);
    height: 40px;
    padding: 0 20px;
    border-bottom: 1px solid #000;
    display: flex;
    justify-content: flex-start;
    .file-box {
        width: 300px;
        height: 100%;
    }
}
.temp-box {
    .ctrl-temp {
        display: flex;
        padding: 0 18px;
        .create-temp {
            font-size: .5rem;
            line-height: 18px;
            padding: 5px 16px;
            margin: 12px 0;
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
    /* min-height: 200px; */
    height: calc(100vh - 140px);
    width: calc(100% - 60px);
    margin: 0 20px;
    margin-top: 120px;
    background-color: rgba(206, 206, 206, 0.486);
    padding: 0;
    border-radius: 10px;
    overflow: hidden;
    position: relative;
}
</style>