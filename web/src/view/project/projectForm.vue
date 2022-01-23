<template>
  <div style="margin-right: 20px">
    <el-form
      ref="formRef"
      :model="formData"
      :rules="rules"
      size="medium"
      label-width="15px"
    >
      <el-row>
        <el-col :span="16">
          <el-form-item label-width="120px" label="项目名称" prop="name">
            <el-input
              v-model="formData.name"
              placeholder="请输入项目名称项目名称"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label-width="100px" label="项目类型" prop="categories">
            <el-input
              v-model="formData.categories"
              placeholder="请输入项目类型"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="16">
          <el-form-item
            label-width="120px"
            label="收费标准"
            prop="chargeStandard"
          >
            <el-input
              v-model="formData.chargeStandard"
              placeholder="请输入收费标准"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label-width="100px" label="备案日期" prop="createdDate">
            <el-date-picker
              v-model="formData.createdDate"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              :style="{ width: '100%' }"
              placeholder="请输入备案申请日期备案日期"
              clearable
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8">
          <el-form-item label-width="120px" label="项目负责人" prop="manager">
            <el-input
              v-model="formData.manager"
              placeholder="请输入项目负责人"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label-width="100px" label="项目码" prop="projectCode">
            <el-input
              v-model="formData.projectCode"
              placeholder="项目码"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label-width="100px" label="项目所属地" prop="area">
            <el-input
              v-model="formData.area"
              placeholder="项目所属地"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8">
          <el-form-item label-width="120px" label="培训模式：" prop="trainMode">
            <el-input
              v-model="formData.trainMode"
              placeholder="请输入培训模式："
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item
            label-width="100px"
            label="培训人数"
            prop="trainNumOfPerson"
          >
            <el-input
              v-model="formData.trainNumOfPerson"
              placeholder="请输入培训人数"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label-width="100px" label="培训学时数" prop="trainTime">
            <el-input
              v-model="formData.trainTime"
              placeholder="请输入培训学时数"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
      </el-row>
      <!-- 培训开始时间 -->
      <el-row>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="培训开始时间"
            prop="trainStartDate"
          >
            <el-date-picker
              v-model="formData.trainStartDate"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              :style="{ width: '100%' }"
              placeholder="请输入培训开始时间"
              clearable
            />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="培训结束时间"
            prop="trainEndDate"
          >
            <el-date-picker
              v-model="formData.trainEndDate"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              :style="{ width: '100%' }"
              placeholder="请选择培训结束时间"
              clearable
            />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="合同开始时间"
            prop="contractStartDate"
          >
            <el-date-picker
              v-model="formData.contractStartDate"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              :style="{ width: '100%' }"
              placeholder="请选择合同开始时间"
              clearable
            />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="合同结束时间"
            prop="contractEndDate"
          >
            <el-date-picker
              v-model="formData.contractEndDate"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              :style="{ width: '100%' }"
              placeholder="请输入合同结束时间"
              clearable
            />
          </el-form-item>
        </el-col>
      </el-row>
      <!-- 学校管理费、发展基金等 -->
      <el-row>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="学校管理费比例"
            prop="sRadio"
          >
            <el-input
              v-model="formData.sRadio"
              placeholder="请输入学校管理费比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label-width="120px" label="发展基金比例" prop="dRadio">
            <el-input
              v-model="formData.dRadio"
              placeholder="请输入发展基金比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label-width="120px" label="福利比例" prop="wRadio">
            <el-input
              v-model="formData.wRadio"
              placeholder="请输入福利比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label-width="120px" label="课酬比例" prop="cRadio">
            <el-input
              v-model="formData.cRadio"
              placeholder="请输入课酬比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>%</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <!-- 学校管理费、等 -->
      <el-row>
        <el-col :span="6">
          <el-form-item label-width="120px" label="学校管理费" prop="sAmount">
            <el-input
              v-model="formData.sAmount"
              placeholder="学校管理费"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label-width="120px" label="发展基金" prop="dAmount">
            <el-input
              v-model="formData.dAmount"
              placeholder="发展基金"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label-width="120px" label="福利" prop="wAmount">
            <el-input
              v-model="formData.wAmount"
              placeholder="福利"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label-width="120px" label="课酬" prop="cAmount">
            <el-input
              v-model="formData.cAmount"
              placeholder="课酬"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <!-- 委托方 -->
      <el-divider content-position="center">编辑委托方</el-divider>
      <el-row v-for="(c, index) in formData.client" :key="'c' + '-' + index">
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="名称"
            :prop="'client.' + index + '.name'"
            :rules="rules.clp.name"
          >
            <el-input
              v-model="c.name"
              placeholder="请输入名称"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="分成比例"
            :prop="'client.' + index + '.radio'"
            :rules="rules.clp.radio"
          >
            <el-input
              v-model="c.radio"
              placeholder="比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="分成"
            :prop="'client.' + index + '.amount'"
          >
            <el-input
              v-model="c.amount"
              placeholder="分成"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button
            size="small"
            type="primary"
            icon="plus"
            class="table-button"
            @click="addFormListFunc('client')"
          >添加</el-button>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button
            size="small"
            type="danger"
            icon="delete"
            class="table-button"
            @click="deleteFormListFunc('client', index)"
          >删除</el-button>
        </el-col>
      </el-row>
      <!-- 落地方 -->
      <el-divider content-position="center">编辑落地机构</el-divider>
      <el-row
        v-for="(l, index) in formData.landingAgency"
        :key="'l' + '-' + index"
      >
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="名称"
            :prop="'landingAgency.' + index + '.name'"
            :rules="rules.clp.name"
          >
            <el-input
              v-model="l.name"
              placeholder="请输入名称"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="分成比例"
            :prop="'landingAgency.' + index + '.radio'"
            :rules="rules.clp.radio"
          >
            <el-input
              v-model="l.radio"
              placeholder="比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="分成"
            :prop="'landingAgency.' + index + '.amount'"
          >
            <el-input
              v-model="l.amount"
              placeholder="分成"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button
            size="small"
            icon="plus"
            type="primary"
            class="table-button"
            @click="addFormListFunc('landingAgency')"
          >添加</el-button>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button
            size="small"
            type="danger"
            icon="delete"
            class="table-button"
            @click="deleteFormListFunc('landingAgency', index)"
          >删除</el-button>
        </el-col>
      </el-row>
      <!-- 技术方 -->
      <el-divider content-position="center">编辑技术合作方</el-divider>
      <el-row v-for="(p, index) in formData.partner" :key="'p' + '-' + index">
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="名称"
            :prop="'partner.' + index + '.name'"
            :rules="rules.clp.name"
          >
            <el-input
              v-model="p.name"
              placeholder="请输入名称"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="分成比例"
            :prop="'partner.' + index + '.radio'"
            :rules="rules.clp.radio"
          >
            <el-input
              v-model="p.radio"
              type="text"
              placeholder="比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            label-width="120px"
            label="分成"
            :prop="'partner.' + index + '.amount'"
          >
            <el-input
              v-model="p.amount"
              placeholder="分成"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button
            size="small"
            icon="plus"
            type="primary"
            class="table-button"
            @click="addFormListFunc('partner')"
          >添加</el-button>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button
            size="small"
            icon="delete"
            type="danger"
            class="table-button"
            @click="deleteFormListFunc('partner', index)"
          >删除</el-button>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="24">
          <el-form-item label-width="120px" label="备注" prop="remark">
            <el-input
              v-model="formData.remark"
              type="textarea"
              placeholder="请输入备注"
              :maxlength="200"
              show-word-limit
              :autosize="{ minRows: 4, maxRows: 4 }"
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12" :offset="6">
          <el-form-item size="large" align="center">
            <el-button
              type="primary"
              style="margin-right: 50px"
              @click="submitForm"
            >提交</el-button>
            <el-button
              v-if="type === 'create'"
              style="margin-left: 50px"
              @click="resetForm"
            >重置</el-button>
            <el-button
              v-if="type === 'update'"
              style="margin-left: 50px"
              @click="back"
            >返回</el-button>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'Project',
}
</script>

<script setup>
import { createProject, updateProject, findProject } from '@/api/project'

import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'

const route = useRoute()
const router = useRouter()
const type = ref('')
const maxFormRowNum = ref(3)
const formRef = ref()
const formData = ref({
  name: '',
  categories: '',
  area: '',
  chargeStandard: '',
  manager: '',
  createdDate: undefined,
  contractNum: '',
  contractStartDate: undefined,
  contractEndDate: undefined,
  paidAmount: 0,
  projectAmount: 0,
  projectCode: '',
  unpaidAmount: 0,
  trainMode: '',
  trainStartDate: undefined,
  trainEndDate: undefined,
  trainNumOfPerson: 0,
  trainTime: 0,
  client: [{ name: '', radio: 0, amount: 0 }],
  landingAgency: [{ name: '', radio: 0, amount: 0 }],
  partners: [{ name: '', radio: 0, amount: 0 }],
  sAmount: 0,
  dAmount: 0,
  wAmount: 0,
  cAmount: 0,
  sRadio: 0,
  dRadio: 0,
  wRadio: 0,
  cRadio: 0,
  incomeAndOutcome: [
    {
      pg: 0.0,
      ph: 0.0,
      pi: 0.0,
      pj: 0.0,
      pk: 0.0,
      pl: 0.0,
      pm: 0.0,
      pn: 0.0,
      po: 0.0,
      pp: 0.0,
      pq: 0.0,
      pr: 0.0,
      ps: 0.0,
      pt: 0.0,
    },
    {
      pg: 0.0,
      ph: 0.0,
      pi: 0.0,
      pj: 0.0,
      pk: 0.0,
      pl: 0.0,
      pm: 0.0,
      pn: 0.0,
      po: 0.0,
      pp: 0.0,
      pq: 0.0,
      pr: 0.0,
      ps: 0.0,
      pt: 0.0,
    },
  ],
  remark: '',
})
const rules = ref({
  name: [
    {
      required: true,
      message: '请输入项目名称项目名称',
      trigger: 'blur',
    },
  ],
  categories: [
    {
      required: true,
      message: '请输入项目类型',
      trigger: 'blur',
    },
  ],
  chargeStandard: [
    {
      required: true,
      message: '请输入收费标准',
      trigger: 'blur',
    },
  ],
  createdDate: [
    {
      required: true,
      message: '请输入备案申请日期备案日期',
      trigger: 'change',
    },
  ],
  manager: [],
  projectCode: [],
  area: [
    {
      required: true,
      message: '项目所属地',
      trigger: 'blur',
    },
  ],
  trainMode: [
    {
      required: true,
      message: '请输入培训模式：',
      trigger: 'blur',
    },
  ],
  trainNumOfPerson: [
    {
      required: true,
      message: '请输入培训人数',
      trigger: 'blur',
    },
    {
      pattern: /^\d+$/g,
      message: '该项为非负整数',
      trigger: 'blur',
    },
  ],
  trainTime: [
    {
      required: true,
      message: '请输入培训学时数',
      trigger: 'blur',
    },
    {
      pattern: /^\d+$/g,
      message: '该项为非负整数',
      trigger: 'blur',
    },
  ],
  trainStartDate: [],
  trainEndDate: [],
  contractStartDate: [],
  contractEndDate: [],
  sRadio: [
    {
      required: true,
      message: '请输入学校管理费比例',
      trigger: 'blur',
    },
    {
      pattern: /^(100|(([1-9]){1}[0-9]?|0{1})((\.)([0-9]){1,2})?)$/,
      message: '该项为0~100之间的百分数',
      trigger: 'blur',
    },
  ],
  sAmount: [],
  dRadio: [
    {
      required: true,
      message: '请输入发展基金比例',
      trigger: 'blur',
    },
    {
      pattern: /^(100|(([1-9]){1}[0-9]?|0{1})((\.)([0-9]){1,2})?)$/,
      message: '该项为0~100之间的百分数',
      trigger: 'blur',
    },
  ],
  dAmount: [],
  wRadio: [
    {
      required: true,
      message: '请输入福利比例',
      trigger: 'blur',
    },
    {
      pattern: /^(100|(([1-9]){1}[0-9]?|0{1})((\.)([0-9]){1,2})?)$/,
      message: '该项为0~100之间的百分数',
      trigger: 'blur',
    },
  ],
  wAmount: [],
  cRadio: [
    {
      required: true,
      message: '请输入课酬比例',
      trigger: 'blur',
    },
  ],
  cAmount: [],
  // 委托方、落地机构和技术方的校验
  clp: {
    name: [
      {
        required: true,
        message: '请输入名称',
        trigger: 'blur',
      },
    ],
    radio: [
      {
        required: true,
        message: '请输入分成比例',
        trigger: 'blur',
      },
      {
        pattern: /^(100|(([1-9]){1}[0-9]?|0{1})((\.)([0-9]){1,2})?)$/,
        message: '该项为0~100之间的百分数',
        trigger: 'blur',
      },
    ],
  },
})

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  // 这里不是直接把行赋给form，而是又查询了一次？
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

// 表单动态添加行
const addFormListFunc = (prop) => {
  try {
    if (formData.value[prop].length < maxFormRowNum.value) {
      formData.value[prop].push({ name: '', radio: 0, amount: 0 })
    } else {
      ElMessage({
        type: 'error',
        message: '该项最多只能添加三个',
      })
    }
  } catch (e) {
    ElMessage({
      type: 'error',
      message: e,
    })
    return
  }
}
// 表单动态删除行
const deleteFormListFunc = (prop, index) => {
  try {
    if (formData.value[prop].length > 1) {
      formData.value[prop].splice(index, 1)
    } else {
      ElMessage({
        type: 'error',
        message: '该项至少要有一个',
      })
    }
  } catch (e) {
    ElMessage({
      type: 'error',
      message: e,
    })
    return
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
        switch (type.value) {
          case 'update': {
            ElMessage({
              type: 'success',
              message: `成功更新项目：${formData.value.name}`,
            })
            back()
            break
          }
          case 'create': {
            ElMessage({
              type: 'success',
              message: `成功创建项目：${formData.value.name}`,
            })
            resetForm()
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

// 重置按钮
const resetForm = () => {
  formData.value = {
    name: '',
    categories: '',
    chargeStandard: '',
    createdDate: undefined,
    manager: '',
    projectCode: '',
    area: '',
    trainMode: '',
    trainNumOfPerson: 0,
    trainTime: 0,
    trainStartDate: undefined,
    trainEndDate: undefined,
    contractStartDate: undefined,
    contractEndDate: undefined,
    sRadio: 0,
    sAmount: 0,
    dRadio: 0,
    dAmount: 0,
    wRadio: 0,
    wAmount: 0,
    cRadio: 0,
    cAmount: 0,
    client: [{ name: '', radio: 0, amount: 0 }],
    landingAgency: [{ name: '', radio: 0, amount: 0 }],
    partner: [{ name: '', radio: 0, amount: 0 }],
    remark: '',
  }
}
</script>

<style></style>
