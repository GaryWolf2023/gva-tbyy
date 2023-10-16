<template>
    <div class="emr-editor">
        <div class="editor-top">
            <el-button type="primary" plain size="small" @click="createPayloadFunc">上传</el-button>
            <el-button type="primary" plain size="small" @click="getPayloadListFunc">获取列表</el-button>
            <el-button type="primary" plain size="small" @click="getOnePayloadFunc">获取单条记录</el-button>
            <el-button type="primary" plain size="small" @click="updatePayloadFunc">更改</el-button>
            <el-button type="primary" plain size="small" @click="deletePayloadFunc">删除</el-button>
            <el-button type="primary" plain size="small">检索</el-button>
            <el-dropdown>
                <span class="el-dropdown-link">
                    <el-button type="primary" plain size="small">
                        更多操作
                        <el-icon class="el-icon--right">
                            <arrow-down />
                        </el-icon>
                    </el-button>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item>查看更多病历</el-dropdown-item>
                    <el-dropdown-item>查看当前病历历史记录</el-dropdown-item>
                    <el-dropdown-item>当前病历存档</el-dropdown-item>
                    <el-dropdown-item>分享当前病历</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
            </el-dropdown>
            <el-dropdown>
                <span class="el-dropdown-link">
                    <el-button type="primary" plain size="small">
                        辅助工具
                        <el-icon class="el-icon--right">
                            <arrow-down />
                        </el-icon>
                    </el-button>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item>月经工具</el-dropdown-item>
                    <el-dropdown-item>特殊符号添加</el-dropdown-item>
                    <el-dropdown-item>体温三测表</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
            </el-dropdown>
            <!-- <p class="h-10">*上方工具皆为测试使用，并无实际意义</p> -->
        </div>
        <!-- <div class="h-80 w-full">
            <div v-for="it in payloadList">
                <span>{{ it.payload_id }}</span>
                -------------------------------------------------------------------
                <el-button>修改</el-button>
            </div>
        </div> -->
        <div class="show-editor">
            <x-emr v-if="editorType==='xemr'" :docname="docName" @getEditor="getEditor" ></x-emr>
            <create-form v-if="editorType==='form'"></create-form>
            <md-editor v-if="editorType==='md'"></md-editor>
        </div>
    </div>
</template>

<script setup>
import XEmr from '@/components/xemr/index.vue'
import CreateForm from '@/components/createForm/index.vue'
import MdEditor from '@/components/createForm/index.vue'
import { ref, onMounted, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { createPayload, getPayloadList, updatePayload, getPayloadById, deletePayload } from '@/api/ctrlPayload'
import { getFileOfTemp } from '@/api/payloadTemp'
import { ElMessage } from 'element-plus'

const route = useRoute()
const editorType = route.query.editorType
const docName = route.query.docName

const payloadList = ref([])
const editor = ref(null)
const fileData = ref(null)
const getEditor = (e) => {
    console.log(e);
    editor.value = e.editor
    if (fileData.value) {
        editor.value.loadHtml(fileData.value)
    } else {
        ElMessage.warning('未查到相应模板信息，请检查模板是否存在')
        editor.value.loadHtml()
    }
}
onMounted(() => {
    getFileOfTemp({
        tempName: route.query.docName,
        tempType: route.query.editorType
    }).then(res => {
        console.log(res)
        fileData.value = res.data
    })
})

const getPayloadListFunc = () => {
    getPayloadList({
        job_id: 'asdasdhabuhsbduhbasda',
        // unique: '222222222222'
    }).then(res => {
        console.log(res)
        // payloadList.value = res.data
    })
}
const getOnePayloadFunc = () => {
    getPayloadById({
        payload_id: '1695005627732&asdasdhabuhsbduhbasda'
    }).then(res => {
        console.log(res)
        // payloadList.value = res.data
    })
}
// 更新病历
const updatePayloadFunc = () => {
    updatePayload({
        payload_content: 'u====================================================b',
        payload_id: '1695015274364&asdasdhabuhsbduhbasda',
        save_his: false
    }).then(res => {
        console.log(res)
    })
}
// 创建病历
const createPayloadFunc = () => {
    let data = editor.value.getHtml()
    let data2 = editor.value.getBindObject()
    console.log('html', data)
    console.log('json', data2)
    return
    createPayload({
        job_id: 'asdasdhabuhsbduhbasda',
        unique: '222222222222',
        business: 'bnmenzhen',
        order: 'asdasdasd',
        payload_content: 'asbduasduasbdouhabhusdb',
        doc_name: "患者基本信息",
        doc_type: "xemr"
    }).then(res => {
        console.log(res);
    })
}
const deletePayloadFunc = () => {
    deletePayload({payload_id: '1695015274364&asdasdhabuhsbduhbasda'}).then(res => {
        console.log(res)
    })
}
</script>

<style lang="scss" scoped>
.emr-editor {
    height: calc(100vh - 116px);
    .editor-top {
        height: 32px;
        /* padding-top: 10px; */
        .el-button {
            margin: 0px 10px;
        }
    }
    .show-editor {
        height: calc(100% - 42px);
        width: 100%;
    }
}

</style>