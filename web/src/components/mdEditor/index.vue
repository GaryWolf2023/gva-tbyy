<template>
    <div>
        markdown编辑器
    </div>
</template>

<script setup>
import { ref, defineProps } from 'vue'
const props = defineProps({
    docname: {type:String, required: false}
})
console.log(props.editortype);
</script>

<style lang="scss" scoped>

</style>