<template>
    <iframe src="editor.html" v-bind="objectOfAttrs" />
</template>

<script setup>
import { ref, reactive } from 'vue'

const objectOfAttrs = reactive({
                         width:'100%',
                         height:'98%',
                         frameborder: 0
                    })
</script>

<style lang="scss" scoped>

</style>