<template>
    <div class="create-temp">
        <div class="ctrl-emr">
            <el-button plain type="primary" @click="outputFiel">导出模板文件</el-button>
        </div>
        <div class="emr-box">
            <x-emr v-if="editorType==='xemr'" :docname="docName" @getEditor="getEditor"></x-emr>
            <create-form v-if="editorType==='form'"></create-form>
            <md-editor v-if="editorType==='md'"></md-editor>
        </div>
    </div>
</template>

<script setup>
import XEmr from '@/components/xemr/index.vue'
import CreateForm from '@/components/createForm/index.vue'
import MdEditor from '@/components/createForm/index.vue'
import { useRoute } from 'vue-router'
import { ref, computed, watch } from 'vue'

const route = useRoute()
console.log(route.query.emrType)
const editorType = computed(() => {
    return route.query.emrType
})
const docName = ref('')
const editorObj = ref(null)
const getEditor = (editor) => {
    // 如果是xemr的话获取文档事件对象
    editorObj.value = editor
}
const outputFiel = () => {
    switch (editorType.value) {
        case 'xemr':
            outputXemr()
            break;
            case 'form':
            outputForm()
            break;
            case 'md':
            outPutMd()
            break;
        default:
            break;
    }
}
const outputXemr = () => {
    // 导出为html， 想办法获取到这个文件
    editorObj.value.execCommand('exportHtml')
}
const outputForm = () => { }
const outPutMd = () => {}
</script>

<style lang="scss" scoped>
.create-temp {
    height: calc(100vh - 116px);
    .ctrl-emr {
        height: 40px;
    }
    .emr-box {
        height: calc(100% - 40px);
    }
}

</style>