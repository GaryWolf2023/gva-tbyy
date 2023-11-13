<template>
    <div class="job-title-box">
        <div class="total-box">
            <el-button type="primary" v-show="showCreate" plain @click="createFunc"> 创 建 </el-button>
            <el-input v-model="searchData" clearable :style="{ width: '400px', marginLeft: '20px' }" placeholder="通过名称检索">
                <template #append>
                    <el-button :icon="Search" @click="getListFunc" />
                </template>
            </el-input>
            <el-button type="primary" plain @click="refreshFunc" :style="{ marginLeft: '20px' }">刷 新</el-button>
        </div>
        <div class="box-total">
            <slot></slot>
        </div>
        <div class="page-ctrl">
            <el-pagination v-model:current-page="page" v-model:page-size="pageSize" :page-sizes="[20, 40, 60, 100]"
                :small="true" :background="true" layout="total, sizes, prev, pager, next, jumper" :total="props.total"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps(['showCreate', 'total'])
const emit = defineEmits(['create', 'getList'])

const page = ref(0)
const pageSize = ref(20)
import { value } from 'dom7';
const searchData = ref('')

const handleSizeChange = () => {
    emit('getList', {
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchData.value
    })
}
const handleCurrentChange = () => {
    emit('getList', {
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchData.value
    })
}
const createFunc = () => {
    emit('create')
}
const refreshFunc = () => {
    emit('getList', {
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchData.value
    })
}
const getListFunc = () => {
    emit('getList', {
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchData.value
    })
}
</script>

<style lang="scss" scoped>
.job-title-box {
    height: calc(100vh - 126px);

    .total-box {
        height: 36px;
    }

    .box-total {
        height: calc(100% - 76px);
        overflow-x: hidden;
    }

    .page-ctrl {
        height: 40px;
        display: flex;
        justify-content: flex-start;
        align-items: center;

        .el-pagination {
            margin: 5px auto;
        }
    }
}
</style>