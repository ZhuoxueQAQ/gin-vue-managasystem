<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="项目名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目类别:">
          <el-input v-model="formData.categories" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目所属地:">
          <el-input v-model="formData.area" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收费标准:">
          <el-input v-model="formData.chargeStandard" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目负责人:">
          <el-input v-model="formData.manager" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备案申请日期:">
          <el-date-picker v-model="formData.createdDate" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="合同编号:">
          <el-input v-model="formData.contractNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同开始时间:">
          <el-date-picker v-model="formData.contractStartDate" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="合同结束时间:">
          <el-date-picker v-model="formData.contractEndDate" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="已到账费用:">
          <el-input-number v-model="formData.paidAmount" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="项目应收费用:">
          <el-input-number v-model="formData.projectAmount" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="项目码:">
          <el-input v-model="formData.projectCode" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="未到账金额:">
          <el-input-number v-model="formData.unpaidAmount" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="培训模式:">
          <el-input v-model="formData.trainMode" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="培训开始时间:">
          <el-date-picker v-model="formData.trainStartDate" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="培训结束时间:">
          <el-date-picker v-model="formData.trainEndDate" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="培训人数:">
          <el-input v-model.number="formData.trainNumOfPerson" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="培训学时数:">
          <el-input v-model.number="formData.trainTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="委托方:">
          <el-input v-model="formData.client" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="落地机构:">
          <el-input v-model="formData.landingAgency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="技术方:">
          <el-input v-model="formData.partners" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学校管理费、发展基金、福利、课酬:">
          <el-input v-model="formData.sDWCAmount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="预算和支出（根据项目对应的收入和支出流水汇总）:">
          <el-input v-model="formData.incomeAndOutput" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Project'
}
</script>

<script setup>
import {
  createProject,
  updateProject,
  findProject
} from '@/api/project'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        name: '',
        categories: '',
        area: '',
        chargeStandard: '',
        manager: '',
        createdDate: new Date(),
        contractNum: '',
        contractStartDate: new Date(),
        contractEndDate: new Date(),
        paidAmount: 0,
        projectAmount: 0,
        projectCode: '',
        unpaidAmount: 0,
        trainMode: '',
        trainStartDate: new Date(),
        trainEndDate: new Date(),
        trainNumOfPerson: 0,
        trainTime: 0,
        client: '',
        landingAgency: '',
        partners: '',
        sDWCAmount: '',
        incomeAndOutput: '',
        remark: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProject({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reproject
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createProject(formData.value)
          break
        case 'update':
          res = await updateProject(formData.value)
          break
        default:
          res = await createProject(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
