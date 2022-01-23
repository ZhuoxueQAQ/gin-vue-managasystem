<template>
  <div>
    <el-form
      ref="elForm"
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
              format="yyyy-MM-dd"
              value-format="yyyy-MM-dd"
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
      <el-row>
        <el-col :span="12">
          <el-form-item
            label-width="120px"
            label="培训开始时间"
            prop="trainStartDate"
          >
            <el-date-picker
              v-model="formData.trainStartDate"
              format="yyyy-MM-dd"
              value-format="yyyy-MM-dd"
              :style="{ width: '100%' }"
              placeholder="请输入培训开始时间"
              clearable
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item
            label-width="120px"
            label="培训结束时间"
            prop="trainEndDate"
          >
            <el-date-picker
              v-model="formData.trainEndDate"
              format="yyyy-MM-dd"
              value-format="yyyy-MM-dd"
              :style="{ width: '100%' }"
              placeholder="请选择培训结束时间"
              clearable
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item
            label-width="120px"
            label="合同开始时间"
            prop="contractStartDate"
          >
            <el-date-picker
              v-model="formData.contractStartDate"
              format="yyyy-MM-dd"
              value-format="yyyy-MM-dd"
              :style="{ width: '100%' }"
              placeholder="请选择合同开始时间"
              clearable
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item
            label-width="120px"
            label="合同结束时间"
            prop="contractEndDate"
          >
            <el-date-picker
              v-model="formData.contractEndDate"
              format="yyyy-MM-dd"
              value-format="yyyy-MM-dd"
              :style="{ width: '100%' }"
              placeholder="请输入合同结束时间"
              clearable
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
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
              <template slot="append">%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label-width="120px" label="学校管理费" prop="sAmount">
            <el-input
              v-model="formData.sAmount"
              placeholder="学校管理费"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="prepend">￥</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item label-width="120px" label="发展基金比例" prop="dRadio">
            <el-input
              v-model="formData.dRadio"
              placeholder="请输入发展基金比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="append">%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label-width="120px" label="发展基金" prop="dAmount">
            <el-input
              v-model="formData.dAmount"
              placeholder="发展基金"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="prepend">￥</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item label-width="120px" label="福利比例" prop="wRadio">
            <el-input
              v-model="formData.wRadio"
              placeholder="请输入福利比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="append">%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label-width="120px" label="福利" prop="wAmount">
            <el-input
              v-model="formData.wAmount"
              placeholder="福利"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="prepend">￥</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item label-width="120px" label="课酬比例" prop="cRadio">
            <el-input
              v-model="formData.cRadio"
              placeholder="请输入课酬比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="append">%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label-width="120px" label="课酬" prop="cAmount">
            <el-input
              v-model="formData.cAmount"
              placeholder="课酬"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="prepend">￥</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8">
          <el-form-item label-width="120px" label="名称" prop="name">
            <el-input
              v-model="formData.name"
              placeholder="请输入名称"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label-width="50px" label="比例" prop="radio">
            <el-input
              v-model="formData.radio"
              placeholder="比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="append">%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="7">
          <el-form-item label-width="50px" label="分成" prop="amount">
            <el-input
              v-model="formData.amount"
              placeholder="分成"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="prepend">￥</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8">
          <el-form-item label-width="120px" label="名称" prop="field200">
            <el-input
              v-model="formData.field200"
              placeholder="请输入名称"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label-width="50px" label="比例" prop="field201">
            <el-input
              v-model="formData.field201"
              placeholder="比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="append">%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="7">
          <el-form-item label-width="50px" label="分成" prop="field202">
            <el-input
              v-model="formData.field202"
              placeholder="分成"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="prepend">￥</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8">
          <el-form-item label-width="120px" label="名称" prop="field204">
            <el-input
              v-model="formData.field204"
              placeholder="请输入名称"
              clearable
              :style="{ width: '100%' }"
            />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label-width="50px" label="比例" prop="field205">
            <el-input
              v-model="formData.field205"
              placeholder="比例"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="append">%</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="7">
          <el-form-item label-width="50px" label="分成" prop="field206">
            <el-input
              v-model="formData.field206"
              placeholder="分成"
              readonly
              :disabled="true"
              clearable
              :style="{ width: '100%' }"
            >
              <template slot="prepend">￥</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
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
      </el-row>
      <el-form-item size="large">
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  components: {},
  props: [],
  data() {
    return {
      formData: {
        name: undefined,
        categories: '',
        chargeStandard: undefined,
        createdDate: '2022-01-04',
        manager: undefined,
        projectCode: undefined,
        area: '',
        trainMode: undefined,
        trainNumOfPerson: '',
        trainTime: undefined,
        trainStartDate: undefined,
        trainEndDate: undefined,
        contractStartDate: undefined,
        contractEndDate: undefined,
        sRadio: undefined,
        sAmount: undefined,
        dRadio: undefined,
        dAmount: undefined,
        wRadio: undefined,
        wAmount: undefined,
        cRadio: undefined,
        cAmount: undefined,
        name: '',
        radio: '',
        amount: undefined,
        field200: undefined,
        field201: '',
        field202: undefined,
        field204: undefined,
        field205: '',
        field206: undefined,
        remark: undefined,
      },
      rules: {
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
            message: '比例',
            trigger: 'blur',
          },
        ],
        amount: [],
        field200: [
          {
            required: true,
            message: '请输入名称',
            trigger: 'blur',
          },
        ],
        field201: [
          {
            required: true,
            message: '比例',
            trigger: 'blur',
          },
        ],
        field202: [],
        field204: [
          {
            required: true,
            message: '请输入名称',
            trigger: 'blur',
          },
        ],
        field205: [
          {
            required: true,
            message: '比例',
            trigger: 'blur',
          },
        ],
        field206: [],
        remark: [],
      },
    }
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {},
  methods: {
    submitForm() {
      this.$refs['elForm'].validate((valid) => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
  },
}
</script>
<style></style>
