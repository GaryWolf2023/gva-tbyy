<template>
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
        <p class="h-10">*上方工具皆为测试使用，并无实际意义</p>
    </div>
    <!-- <div class="h-80 w-full">
        <div v-for="it in payloadList">
            <span>{{ it.payload_id }}</span>
            -------------------------------------------------------------------
            <el-button>修改</el-button>
        </div>
    </div> -->
    <div class="show-editor">
        <x-emr v-if="editorType==='xemr'" :docname="docName" ></x-emr>
        <create-form v-if="editorType==='form'"></create-form>
        <md-editor v-if="editorType==='md'"></md-editor>
    </div>
</template>

<script setup>
import XEmr from '@/components/xemr/index.vue'
import CreateForm from '@/components/createForm/index.vue'
import MdEditor from '@/components/createForm/index.vue'
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { createPayload, getPayloadList, updatePayload, getPayloadById, deletePayload } from '@/api/ctrlPayload'

const route = useRoute()
const editorType = route.query.editorType
console.log(editorType);
const docName = route.query.docName
const payloadList = ref([])

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
const updatePayloadFunc = () => {
    updatePayload({
        payload_content: 'u====================================================b',
        payload_id: '1695015274364&asdasdhabuhsbduhbasda',
        save_his: false
    }).then(res => {
        console.log(res)
    })
}
const createPayloadFunc = () => {
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
.editor-top {
    margin-top: 96px;
    height: 36px;
    padding: 0 20px;
    padding-top: 10px;
    .el-button {
        margin: 0px 10px;
    }
}
.show-editor {
    width: 100%;
    height: calc(100vh - 132px);
}
</style>