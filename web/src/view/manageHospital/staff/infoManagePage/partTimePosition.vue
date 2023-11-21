<template>
    <div>
        <el-table :data="partTimeJobData" style="width: 100%" max-height="320">
            <el-table-column fixed prop="date" label="工作岗位" />
            <el-table-column prop="name" label="部门" />
            <el-table-column prop="state" label="应用" />
            <el-table-column prop="city" label="目标" />
            <el-table-column prop="city" label="岗位代码" />
            <el-table-column prop="city" label="岗位性质" />
            <el-table-column prop="city" label="岗位类别" />
            <el-table-column prop="city" label="是否启用" />
            <el-table-column fixed="right" label="" width="120">
              <template #default="scope">
                <el-button
                  link
                  type="primary"
                  size="small"
                  @click.prevent="deleteRowOfPartTimeJob(scope.$index)"
                >
                <el-icon><CloseBold /></el-icon>
                </el-button>
              </template>
            </el-table-column>
        </el-table>
        <el-button class="mt-4" style="width: 100%" @click="onAddPartTimeJobFunc">新增兼职</el-button>
    </div>
</template>

<script setup>

</script>

<style lang="scss" scoped>

</style>