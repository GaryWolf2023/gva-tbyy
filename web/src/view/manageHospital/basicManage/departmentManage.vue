<template>
    <div>
        <ManageTemplate
            :showCreate="true"
            :total="total"
            @create="createButtonFunc"
            @getList="getListFunc"
        >

        </ManageTemplate>
    </div>
</template>

<script setup>
import ManageTemplate from '@/components/manageTemplate/index.vue'


const createButtonFunc = () => { }
const getListFunc = () => {}
</script>

<style lang="scss" scoped>

</style>