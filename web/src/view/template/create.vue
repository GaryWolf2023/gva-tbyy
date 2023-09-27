<template>
    <div class="emr-box">
        <x-emr v-if="editorType==='xemr'" :docname="docName" ></x-emr>
        <create-form v-if="editorType==='form'"></create-form>
        <md-editor v-if="editorType==='md'"></md-editor>
    </div>
</template>

<script setup>
import XEmr from '@/components/xemr/index.vue'
import CreateForm from '@/components/createForm/index.vue'
import MdEditor from '@/components/createForm/index.vue'
import { useRoute } from 'vue-router'
import { ref, onMounted, watch } from 'vue'

const route = useRoute()
const editorType = ref()
const docName = ref('')
watch(() => route.query, (val) => {
    console.log(val)
    editorType.value = val.emrType
}, { immediate: true })

</script>

<style lang="scss" scoped>
.emr-box {
    height: calc(100vh - 120px);
    margin-top: 110px;
}
</style>