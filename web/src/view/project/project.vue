<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="项目名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="项目类别">
          <el-input v-model="searchInfo.categories" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="备案申请日期">
          <el-input v-model="searchInfo.createdDate" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="合同开始时间">
          <el-input
            v-model="searchInfo.contractStartDate"
            placeholder="搜索条件"
          />
        </el-form-item>
        <el-form-item label="项目码">
          <el-input v-model="searchInfo.projectCode" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="培训开始时间">
          <el-input
            v-model="searchInfo.trainStartDate"
            placeholder="搜索条件"
          />
        </el-form-item>
        <el-form-item label="委托方">
          <el-input v-model="searchInfo.client" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="落地机构">
          <el-input v-model="searchInfo.landingAgency" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="技术方">
          <el-input v-model="searchInfo.partners" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button
            size="mini"
            type="primary"
            icon="search"
            @click="onSubmit"
          >查询</el-button>
          <el-button
            size="mini"
            icon="refresh"
            @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          size="mini"
          type="primary"
          icon="plus"
          @click="openDialog"
        >新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button
              size="mini"
              type="text"
              @click="deleteVisible = false"
            >取消</el-button>
            <el-button
              size="mini"
              type="primary"
              @click="onDelete"
            >确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="mini"
              style="margin-center: 10px"
              :disabled="!multipleSelection.length"
            >删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        v-if="tableData.length > 0"
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="center" label="创建日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.createdDate)
          }}</template>
        </el-table-column>
        // todo clp
        <template v-for="(col, index) in tableCols" :key="index">
          <!-- multi -->
          <template v-if="col.type === 'multi'">
            <template
              v-for="(dataItem, dataIndex) in tableData[0][col.prop]"
              :key="dataIndex"
            >
              <template v-for="(s, subIndex) in col.sub" :key="subIndex">
                <el-table-column
                  :prop="s.prop"
                  :label="s.label"
                  align="center"
                  show-overflow-tooltip
                  width="120"
                >
                  <!-- <template #default="scope">{{
                    valFormat(scope.row[col.prop][dataIndex][s.prop], s.format)
                  }}</template> -->
                  <template #default="scope">{{
                    scope.row[col.prop][dataIndex][s.prop]
                  }}</template>
                </el-table-column>
              </template>
            </template>
          </template>
          <!-- 其他的正常渲染 -->
          <template v-else>
            <el-table-column
              :prop="col.prop"
              :label="col.label"
              align="center"
              show-overflow-tooltip
              width="150"
            >
              <template #default="scope">
                <!-- 正常显示，row为当前某一行的数据对象，col.prop为数据对象的某个属性， scope.row[col.prop]为某行某列的一个值-->
                <!-- todo format -->
                <!-- <span>
                  {{ valFormat(scope.row[col.prop], col.format) }}
                </span> -->
                <span>
                  {{ scope.row[col.prop] }}
                </span>
              </template>
            </el-table-column>
          </template>
        </template>
        <el-table-column align="center" label="按钮组">
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateProjectFunc(scope.row)"
            >变更</el-button>
            <el-button
              type="text"
              icon="delete"
              size="mini"
              @click="deleteRow(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="弹窗操作"
    >
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="项目名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目类别:">
          <el-input
            v-model="formData.categories"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="项目所属地:">
          <el-input v-model="formData.area" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收费标准:">
          <el-input
            v-model="formData.chargeStandard"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="项目负责人:">
          <el-input v-model="formData.manager" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备案申请日期:">
          <el-date-picker
            v-model="formData.createdDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            clearable
          />
        </el-form-item>
        <el-form-item label="合同编号:">
          <el-input
            v-model="formData.contractNum"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="合同开始时间:">
          <el-date-picker
            v-model="formData.contractStartDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            clearable
          />
        </el-form-item>
        <el-form-item label="合同结束时间:">
          <el-date-picker
            v-model="formData.contractEndDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            clearable
          />
        </el-form-item>
        <el-form-item label="已到账费用:">
          <el-input-number
            v-model="formData.paidAmount"
            style="width: 100%"
            :precision="2"
            clearable
          />
        </el-form-item>
        <el-form-item label="项目应收费用:">
          <el-input-number
            v-model="formData.projectAmount"
            style="width: 100%"
            :precision="2"
            clearable
          />
        </el-form-item>
        <el-form-item label="项目码:">
          <el-input
            v-model="formData.projectCode"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="未到账金额:">
          <el-input-number
            v-model="formData.unpaidAmount"
            style="width: 100%"
            :precision="2"
            clearable
          />
        </el-form-item>
        <el-form-item label="培训模式:">
          <el-input
            v-model="formData.trainMode"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="培训开始时间:">
          <el-date-picker
            v-model="formData.trainStartDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            clearable
          />
        </el-form-item>
        <el-form-item label="培训结束时间:">
          <el-date-picker
            v-model="formData.trainEndDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            clearable
          />
        </el-form-item>
        <el-form-item label="培训人数:">
          <el-input
            v-model.number="formData.trainNumOfPerson"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="培训学时数:">
          <el-input
            v-model.number="formData.trainTime"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="委托方:">
          <el-input v-model="formData.client" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="落地机构:">
          <el-input
            v-model="formData.landingAgency"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="技术方:">
          <el-input
            v-model="formData.partners"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="学校管理费、发展基金、福利、课酬:">
          <el-input
            v-model="formData.sDWCAmount"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="预算和支出（根据项目对应的收入和支出流水汇总）:">
          <el-input
            v-model="formData.incomeAndOutput"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button
            size="small"
            type="primary"
            @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Project',
}
</script>

<script setup>
import {
  createProject,
  deleteProject,
  deleteProjectByIds,
  updateProject,
  findProject,
  getProjectList,
} from '@/api/project'

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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
  client: {},
  landingAgency: {},
  partners: {},
  sAmount: 0,
  dAmount: 0,
  wAmount: 0,
  cAmount: 0,
  sRadio: 0,
  dRadio: 0,
  wRadio: 0,
  cRadio: 0,
  incomeAndOutput: '',
  remark: '',
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const tableCols = ref([
  { prop: 'name', label: '项目名称', isShow: true },
  { prop: 'categories', label: '项目类别', isShow: true },
  { prop: 'projectNum', label: '项目码', isShow: true },
  { prop: 'manager', label: '项目负责人', isShow: true },
  { prop: 'area', label: '所属地', isShow: true },
  { prop: 'trainMode', label: '培训模式', isShow: true },
  { prop: 'chargeStandard', label: '收费标准', isShow: true },
  { prop: 'trainTime', label: '学时数', isShow: true },
  { prop: 'trainNumOfPerson', label: '培训人数', isShow: true },
  {
    prop: 'trainStartDate',
    label: '培训开始时间',
    isShow: true,
    format: 'date',
  },
  { prop: 'trainEndDate', label: '培训结束时间', isShow: true, format: 'date' },
  {
    prop: 'contractStartDate',
    label: '合同开始时间',
    isShow: true,
    format: 'date',
  },
  {
    prop: 'contractEndDate',
    label: '合同结束时间',
    isShow: true,
    format: 'date',
  },
  {
    prop: 'projectFund',
    label: '项目应收费用',
    isShow: true,
    format: 'amount',
  },
  { prop: 'paidAmount', label: '已到账费用', isShow: true, format: 'amount' },
  {
    prop: 'unpaidAmount',
    label: '未到账金额',
    isShow: true,
    format: 'amount',
  },
  {
    prop: 'client',
    type: 'multi',
    maxLengthIndex: 0,
    sub: [
      { prop: 'name', label: '委托方', isShow: true },
      { prop: 'radio', label: '委托方分成比例', format: 'radio', isShow: true },
      { prop: 'amount', label: '委托方分成', format: 'amount', isShow: true },
    ],
  },
  {
    prop: 'landingAgency',
    type: 'multi',
    maxLengthIndex: 0,
    sub: [
      { prop: 'name', label: '落地方', isShow: true },
      { prop: 'radio', label: '落地方分成比例', format: 'radio', isShow: true },
      { prop: 'amount', label: '落地方分成', format: 'amount', isShow: true },
    ],
  },
  {
    prop: 'partners',
    type: 'multi',
    maxLengthIndex: 0,
    sub: [
      { prop: 'name', label: '技术方', isShow: true },
      { prop: 'radio', label: '技术方分成比例', format: 'radio', isShow: true },
      { prop: 'amount', label: '技术方分成', format: 'amount', isShow: true },
    ],
  },
  {
    prop: 'sRadio',
    label: '学校管理费比例',
    format: 'radio',
    isShow: true,
  },
  { prop: 'sAmount', label: '学校管理费', format: 'amount', isShow: true },
  { prop: 'dRadio', label: '发展基金比例', format: 'radio', isShow: true },
  { prop: 'dAmount', label: '发展基金', format: 'amount', isShow: true },
  { prop: 'wRadio', label: '福利比例', format: 'radio', isShow: true },
  { prop: 'wAmount', label: '福利', format: 'amount', isShow: true },
  { prop: 'cRadio', label: '课酬比例', format: 'radio', isShow: true },
  { prop: 'cAmount', label: '课酬', format: 'amount', isShow: true },
  { prop: 'remark', label: '备注', isShow: true },
])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getProjectList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  console.log('table', table)
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    deleteProjectFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据',
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID)
    })
  const res = await deleteProjectByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateProjectFunc = async(row) => {
  console.log(row)
  // const res = await findProject({ ID: row.ID })
  // type.value = 'update'
  // if (res.code === 0) {
  //   formData.value = res.data.reproject
  //   dialogFormVisible.value = true
  // }
}

// 删除行
const deleteProjectFunc = async(row) => {
  const res = await deleteProject({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
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
  }
}
// 弹窗确定
const enterDialog = async() => {
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
      message: '创建/更改成功',
    })
    closeDialog()
    getTableData()
  }
}
</script>

<style></style>
