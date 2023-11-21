<template>
    <div class="table-info">
        <el-table :data="tableData" style="width: 100%" max-height="250">
            <el-table-column fixed prop="label" label="疾病名称" align="center" />
            <el-table-column label="疾病编码" prop="value" width="280" align="center"/>
            <el-table-column fixed="right" width="120">
              <template #default="scope">
                <el-button
                  link
                  type="primary"
                  size="small"
                  @click.prevent="deleteRow(scope.row)"
                >
                <el-icon><CloseBold /></el-icon>
                </el-button>
              </template>
            </el-table-column>
        </el-table>
        <el-button color="#626aef" class="mt-4" :style="{width:'100%'}" plain  size="large" @click="dialogVisible=true">新 增 条 目</el-button>
        <el-dialog
          v-model="dialogVisible"
          title="新增擅长疾病条目"
          width="60%"
          :before-close="handleClose"
          @open="openFunc"
        >
          <div class="dia-info-list">
            <div class="dia-list-search">
                <el-row>
                    <el-col :span="20" :offset="4">
                        <el-input v-model="searchData" clearable type="text" placeholder="根据疾病种类名称检索">
                            <template #append>
                              <el-button :icon="Search" />
                            </template>
                        </el-input>
                    </el-col>
                    <el-col :span="20" :offset="4">
                        <p :style="{textAlign:'center',fontSize:'12px',color:'red'}">*双击选中</p>
                    </el-col>
                </el-row>
            </div>
            <div class="null-list" v-if="illnessList.length==0">暂无数据</div>
            <ul class="illness-list" v-else :style="{height:'360px',width:'calc(100% - 40px)',margin:'10px 20px'}">
                <li v-for="it in illnessList" @dblclick="checkIllness(it)">
                    <span :style="{width:'50%',textAlign:'center'}">{{ it.label }}</span>
                    <span :style="{width:'50%',textAlign:'center'}">{{ it.value }}</span>
                </li>
            </ul>
          </div>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Search } from '@element-plus/icons-vue'
import { getNumberCode } from '@/api/common.js'
import { getAdeptIllnessList, addAdeptIllness, deleteAdeptIllness } from '@/api/staff.js'

const props = defineProps({
    id: Number
})

const dialogVisible = ref(false)
const tableData = ref([])
const illnessList = ref([])
const searchData = ref('')
onMounted(() => {
    getData()
})
const getData = async () => {
    const data = await getNumberCode('CV02.10.005')
    illnessList.value = data.data
    const data2 = await getAdeptIllnessList({ employeeId: props.id })
    if (data2.code === 0) {
        if (data2.data.length>0) {
            data2.data.forEach(it => {
                illnessList.value.forEach(item => {
                    if (item.value === it.illnessId) {
                        tableData.value.push(item)
                    }
                })
            })

        }
    }
    
}
watch(
    () => searchData.value,
    () => {
        illnessList.value = illnessList.value.filter(it => {
            if (it.label.includes(searchData.value)) {
                return it
            }
        })
    }
)
const getAdeptIllnessListFunc = () => {
    getAdeptIllnessList({employeeId: props.id}).then(res => {
        tableData.value = res.data
    })
}
const openFunc = () => {
    getNumberCode('CV02.10.005').then(res => {
        illnessList.value = res.data
    })
}
const handleClose = (down) => {
    down()
}
const deleteRow = (row) => {
    deleteAdeptIllness({ employeeId: props.id, illnessId: row.value }).then(res => {
        ElMessage.warning(res.msg)
        getData()
    })
} 
const checkIllness = (val) => {
    addAdeptIllness({
        employeeId: Number(props.id),
        illnessId: val.value
    }).then(res => {
        ElMessage.info(res.msg)
        if(res.code === 0){
            dialogVisible = false
            getAdeptIllnessListFunc()
        }
    })
}
</script>

<style lang="scss" scoped>
.table-info {

}
.dia-info-list {
    .dia-list-search {
        
    }
    .null-list {
        width: 100%;
        height: 200px;
        text-align: center;
        line-height: 200px;
    }
    .illness-list {
        overflow-x: hidden;
        >li {
            line-height: 26px;
            display: flex;
            justify-content: space-around;
            &:hover {
                background-color: rgba(140, 121, 194, 0.39);
            }
        }
    }
}
</style>