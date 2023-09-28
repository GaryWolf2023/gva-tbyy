<template>
    <div class="h-full">
        <x-emr @DocLoaded="onDocLoaded" @AfterInit="onAfterInit"></x-emr>
    </div>
</template>

<script setup>
import XEmr from './editor.vue'
import { ref, defineProps, defineEmits } from 'vue'
const props = defineProps({
    docname: {type:String, required: false}
})
const emits = defineEmits(['getEditor'])
const editor = ref(null)

//文档加载完成
const onDocLoaded = function(e) {
    editor.value = e.editor
    emits('getEditor', e.editor)
}
const onAfterInit = (e) => {
    console.log(e)
    e.editor.loadUrl('/src/assets/docMode/' + props.docname + '.html')
}
</script>

<style lang="scss" scoped>

</style>