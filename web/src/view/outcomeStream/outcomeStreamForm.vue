<template>
  <div>
    <div class="gva-form-box">
      <el-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        size="medium"
        label-width="120px"
      >
        <el-row type="flex" justify="start" align="center">
          <el-col :span="24">
            <el-form-item label="所属项目" prop="projectName">
              <el-input
                v-model="formData.projectName"
                :disabled="true"
                :style="{ width: '100%' }"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row type="flex" justify="start" align="center">
          <el-col :span="12">
            <el-form-item label="支出日期" prop="createdDate">
              <el-date-picker
                v-model="formData.createdDate"
                format="YYYY-MM-DD"
                :style="{ width: '100%' }"
                placeholder="请输入入账日期"
                clearable
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="支出项目码" prop="outcomeProjectCode">
              <el-input
                v-model="formData.outcomeProjectCode"
                placeholder="请输入支出项目码"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col
            v-for="(c, index) in formData.client"
            :key="'client' + '-' + index + '-' + c.name"
            :span="parseInt(24 / formData.client.length)"
          >
            <el-form-item
              :label="`委托方${String(index + 1)}分成`"
              :prop="'client.' + index + '.amount'"
              :rules="rules.clp.amount"
            >
              <el-input
                v-model="c.amount"
                placeholder="分成"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col
            v-for="(c, index) in formData.landingAgency"
            :key="'landingAgency' + '-' + index + '-' + c.name"
            :span="parseInt(24 / formData.landingAgency.length)"
          >
            <el-form-item
              :label="`落地机构${String(index + 1)}分成`"
              :prop="'landingAgency.' + index + '.amount'"
              :rules="rules.clp.amount"
            >
              <el-input
                v-model="c.amount"
                placeholder="分成"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col
            v-for="(c, index) in formData.partner"
            :key="'partner' + '-' + index + '-' + c.name"
            :span="parseInt(24 / formData.partner.length)"
          >
            <el-form-item
              :label="`技术方${String(index + 1)}分成`"
              :prop="'partner.' + index + '.amount'"
              :rules="rules.clp.amount"
            >
              <el-input
                v-model="c.amount"
                placeholder="分成"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row type="flex" justify="start" align="center">
          <el-col :span="8">
            <el-form-item label="专家课酬" prop=".pays.pg">
              <el-input
                v-model="formData.pays.pg"
                placeholder="请输入专家课酬"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="方案费" prop=".pays.ph">
              <el-input
                v-model="formData.pays.ph"
                placeholder="请输入方案费"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="网络研修教学费" prop="pays.pi">
              <el-input
                v-model="formData.pays.pi"
                placeholder="请输入网络研修教学费"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row type="flex" justify="start" align="center">
          <el-col :span="8">
            <el-form-item label="管理及工作人员劳务费" prop="pays.pj">
              <el-input
                v-model="formData.pays.pj"
                placeholder="请输入管理及工作人员劳务费"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="考察、跟岗、线下指导费" prop="pays.pk">
              <el-input
                v-model="formData.pays.pk"
                placeholder="请输入考察、跟岗、线下指导费"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="专家食宿" prop="pays.pl">
              <el-input
                v-model="formData.pays.pl"
                placeholder="请输入专家食宿"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row type="flex" justify="start" align="center">
          <el-col :span="8">
            <el-form-item label="专家及工作人员交通费" prop="pays.pm">
              <el-input
                v-model="formData.pays.pm"
                placeholder="请输入专家及工作人员交通费"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="学员伙食费" prop="pays.pn">
              <el-input
                v-model="formData.pays.pn"
                placeholder="请输入学员伙食费"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="学员住宿费" prop="pays.po">
              <el-input
                v-model="formData.pays.po"
                placeholder="请输入学员住宿费"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row type="flex" justify="start" align="center">
          <el-col :span="8">
            <el-form-item label="学员交通费" prop="pays.pp">
              <el-input
                v-model="formData.pays.pp"
                placeholder="请输入学员交通费"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="场租" prop="pays.pq">
              <el-input
                v-model="formData.pays.pq"
                placeholder="请输入场租"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="课程资源建设费(人员支出)" prop="pays.pr">
              <el-input
                v-model="formData.pays.pr"
                placeholder="请输入课程资源建设费(人员支出)"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row type="flex" justify="start" align="center">
          <el-col :span="8">
            <el-form-item label="课程资源建设费(购买)" prop="pays.ps">
              <el-input
                v-model="formData.pays.ps"
                placeholder="请输入课程资源建设费(购买)"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="16">
            <el-form-item
              label-width="300px"
              label="培训资料费、办公用品费、保险费、医药费及其他(含购买标书、中标服务费等)"
              prop="pays.pt"
            >
              <el-input
                v-model="formData.pays.pt"
                placeholder="请输入培训资料费、办公用品费、保险费、医药费及其他(含购买标书、中标服务费等)"
                clearable
                :style="{ width: '100%' }"
              >
                <template #append>元</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row type="flex" justify="start" align="center">
          <el-col :span="24">
            <el-form-item label="备注" prop="remark">
              <el-input
                v-model="formData.pays.ps"
                placeholder="请输入备注"
                clearable
                :style="{ width: '100%' }"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col>
            <el-form-item size="large" align="center">
              <el-button
                type="primary"
                style="margin-right: 20px"
                @click="submitForm"
              >提交</el-button>
              <el-button
                v-if="type === 'create'"
                style="margin-left: 20px"
                @click="resetForm"
              >重置</el-button>
              <el-button
                v-if="type === 'update'"
                style="margin-left: 20px"
                @click="back"
              >返回</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'OutcomeStream',
}
</script>

<script setup>
import {
  createOutcomeStream,
  updateOutcomeStream,
  findOutcomeStream,
} from '@/api/outcomeStream'

// 自动获取字典
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

const formRef = ref()
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
  projectId: 0,
  projectName: '',
  categories: '',
  createdDate: undefined,
  outcomeProjectCode: '',
  client: [],
  landingAgency: [],
  partner: [],
  pays: {
    pg: 0,
    ph: 0,
    pi: 0,
    pj: 0,
    pk: 0,
    pl: 0,
    pm: 0,
    pn: 0,
    po: 0,
    pp: 0,
    pq: 0,
    pr: 0,
    ps: 0,
    pt: 0,
  },
})
const rules = ref({
  createdDate: [
    {
      required: true,
      message: '请输入支出日期',
      trigger: 'blur',
    },
  ],
  clp: {
    amount: [
      {
        required: true,
        message: '请输入分成金额',
        trigger: 'blur',
      },
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
  },
  // 一堆预算的校验
  pays: {
    pg: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    ph: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pi: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pj: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pk: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pl: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pm: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pn: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    po: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pp: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pq: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pr: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    ps: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
    pt: [
      {
        pattern: /^\d+(\.\d+)?$/,
        message: '金额为非负数',
        trigger: 'blur',
      },
    ],
  },
})

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findOutcomeStream({ ID: route.query.id })
    console.log(res)
    if (res.code === 0) {
      formData.value = res.data.reoutcome
      type.value = 'update'
    }
  } else if (route.query.project) {
    // 如果是从project那里点进来添加流水，合法，此时将所属项目的ID和名称先赋给该流水
    // todo 在这里赋值，收入流水所属的项目id、名称、委托方、学校管理费比例。。等
    const project = JSON.parse(route.query.project)
    // init
    formData.value.projectId = project.ID
    formData.value.projectName = project.name
    formData.value.categories = project.categories
    formData.value.client = []
    formData.value.landingAgency = []
    formData.value.partner = []
    // 委托方、落地机构、合作方
    project.client.forEach((c) => {
      formData.value.client.push({ name: c.name, radio: c.radio, amount: 0 })
    })
    project.landingAgency.forEach((c) => {
      formData.value.landingAgency.push({
        name: c.name,
        radio: c.radio,
        amount: 0,
      })
    })
    project.partner.forEach((c) => {
      formData.value.partner.push({
        name: c.name,
        radio: c.radio,
        amount: 0,
      })
    })
    type.value = 'create'
    console.log(formData.value)
  } else {
    ElMessageBox.alert(
      '请前往“培训项目列表”中选择该流水所属的培训项目',
      '需选择待录入流水所属的项目',
      {
        confirmButtonText: '好的，这就去',
        callback: () => router.push({ name: 'project' }),
      }
    )
  }
}

init()
// 保存按钮(提交表单)
const submitForm = async() => {
  let res
  formRef.value.validate(async(valid) => {
    if (valid) {
      switch (type.value) {
        case 'create':
          res = await createOutcomeStream(formData.value)
          break
        case 'update':
          res = await updateOutcomeStream(formData.value)
          break
        default:
          res = await createOutcomeStream(formData.value)
          break
      }
      console.log(res)
      if (res.code === 0) {
        switch (type.value) {
          case 'update': {
            ElMessage({
              type: 'success',
              message: `成功更新支出流水`,
            })
            back()
            break
          }
          case 'create': {
            ElMessageBox.confirm(
              `成功为项目：${formData.value.projectName}添加支出流水`,
              'success',
              {
                confirmButtonText: '继续添加',
                cancelButtonText: '查看流水',
                type: 'success',
              }
            )
              .then(() => {
                init()
              })
              .catch(() => {
                router.push({ name: 'outcomeStream' })
              })
            break
          }
          default:
            break
        }
      }
    } else {
      ElMessage({
        type: 'error',
        message: '表单提交失败，请根据提示信息进行修改',
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

// 重置
const resetForm = () => {}
</script>

<style></style>
