<template>
    <div>
        <el-form :model="form" label-width="120px">
            <el-row :style="{margin:'5px 0'}">
                <el-col :span="6">
                    <el-form-item label="姓名:">
                        <span>{{ userInfo.name }}</span>
                    </el-form-item>
                </el-col>
                <el-col :span="6" :offset="3">
                    <el-form-item label="出生日期:">
                        <span>{{ userInfo.birthday }}</span>
                    </el-form-item>
                </el-col>
                <el-col :span="6" :offset="3">
                    <el-form-item label="出生地:">
                        <span>{{ userInfo.placeOfBirth }}</span>
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="职位:">
                        <span>{{ userInfo.jobTitle }}</span>
                    </el-form-item>
                </el-col>
                <el-col :span="6" :offset="3">
                    <el-form-item label="婚姻状况">
                        <span>{{ userInfo.marital }}</span>
                    </el-form-item>
                </el-col>
                <el-col :span="6" :offset="3">
                    <el-form-item label="出生地">
                        <el-input v-model="userInfo.placeOfBirth" disabled/>
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>
    </div>
</template>

<script setup>
import { ref, defineProps, computed, watch } from 'vue'

const props = defineProps(['staffInfo'])
const userInfo = computed(() => {
    console.log(props.staffInfo)
    return props.staffInfo
})
watch(userInfo, (a) => {
    console.log(a)
})
</script>

<style lang="scss" scoped></style>