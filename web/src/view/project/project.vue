<template>
  <div>
    <!-- 表格按钮组 -->
    <div class="gva-search-box">
      <div class="box-top-line gva-btn-list">
        <div>
          <span>表格工具栏：</span>
        </div>
        <div>
          <el-button
            size="mini"
            type="primary"
            icon="plus"
            @click="createProjectFunc"
          >新增</el-button>
          <el-popover
            v-model:visible="deleteVisible"
            placement="top"
            width="160"
          >
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
                type="danger"
                style="margin-center: 10px"
                :disabled="!multipleSelection.length"
              >删除</el-button>
            </template>
          </el-popover>
          <el-button
            size="mini"
            type="primary"
            icon="search"
            @click="handleShowSearchBox"
          >
            搜索
          </el-button>
          <el-button
            size="mini"
            type="primary"
            icon="download"
            @click="handleExcelExport"
          >
            导出
          </el-button>
          <el-button
            size="mini"
            type="primary"
            icon="setting"
            @click="handleShowColCheckGroupBox"
          >
            列筛选
          </el-button>
        </div>
      </div>
    </div>
    <div v-if="isShowSearchBox">
      <div class="gva-search-box">
        <div class="box-top-line">
          <div>
            <span>按条件查询表格数据：</span>
          </div>
          <div>
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
            <el-button
              type="danger"
              size="mini"
              icon="circle-close"
              @click="handleHideSearchBox"
            >隐藏</el-button>
          </div>
        </div>
        <el-form
          v-if="isShowSearchBox"
          :model="searchInfo"
          size="medium"
          label-width="100px"
        >
          <el-row>
            <el-col :span="8">
              <el-form-item label="项目名称" label-width="100px">
                <el-input
                  v-model="searchInfo.name"
                  placeholder="请输入待匹配文本"
                  :style="{ width: '100%' }"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="项目类别" label-width="100px">
                <el-input
                  v-model="searchInfo.categories"
                  placeholder="请输入待匹配文本"
                  :style="{ width: '100%' }"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="项目码" label-width="100px">
                <el-input
                  v-model="searchInfo.projectCode"
                  placeholder="请输入待匹配文本"
                  :style="{ width: '100%' }"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="8">
              <el-form-item label="备案申请日期" label-width="100px">
                <el-date-picker
                  v-model="searchInfo.createdDateRange"
                  format="YYYY-MM-DD"
                  :style="{ width: '100%' }"
                  placeholder="请输入待匹配文本"
                  type="daterange"
                  unlink-panels
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="合同开始时间">
                <el-date-picker
                  v-model="searchInfo.contractStartDateRange"
                  format="YYYY-MM-DD"
                  :style="{ width: '100%' }"
                  placeholder="请输入待匹配文本"
                  type="daterange"
                  unlink-panels
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="培训开始时间">
                <el-date-picker
                  v-model="searchInfo.trainStartDateRange"
                  format="YYYY-MM-DD"
                  :style="{ width: '100%' }"
                  placeholder="请输入待匹配文本"
                  type="daterange"
                  unlink-panels
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col
              v-for="index in maxFormRowNum"
              :key="'client' + index"
              :span="24 / maxFormRowNum"
            >
              <el-form-item
                :label="`委托方${String(index)}名称`"
                label-width="100px"
              >
                <el-input
                  v-model="searchInfo.client[index - 1].name"
                  placeholder="请输入待匹配文本"
                  :style="{ width: '100%' }"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col
              v-for="index in maxFormRowNum"
              :key="'landingAgency' + index"
              :span="24 / maxFormRowNum"
            >
              <el-form-item
                :label="`落地方${String(index)}名称`"
                label-width="100px"
              >
                <el-input
                  v-model="searchInfo.landingAgency[index - 1].name"
                  placeholder="请输入待匹配文本"
                  :style="{ width: '100%' }"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col
              v-for="index in maxFormRowNum"
              :key="'partner' + index"
              :span="24 / maxFormRowNum"
            >
              <el-form-item
                :label="`技术方${String(index)}名称`"
                label-width="100px"
              >
                <el-input
                  v-model="searchInfo.partner[index - 1].name"
                  placeholder="请输入待匹配文本"
                  :style="{ width: '100%' }"
                />
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </div>
    </div>

    <!-- 列选择显示 -->
    <div v-if="isShowColCheckGroupBox">
      <div class="gva-search-box">
        <div class="box-top-line">
          <div>
            <span style="align: center">表格列显示设置：</span>
          </div>
          <div>
            <el-button
              size="mini"
              type="primary"
              icon="check"
              @click="updateManageSystemSettingFunc"
            >保存</el-button>
            <el-button
              type="danger"
              size="mini"
              icon="circle-close"
              @click="handleHideColCheckGroupBox"
            >隐藏</el-button>
          </div>
        </div>
        <div style="display: flex; flex-wrap: wrap">
          <!-- 一堆checkbox -->
          <div
            v-for="(col, index) in tableCols"
            :key="'checkbox' + '-' + index"
          >
            <div v-if="col.colType === 'multi'">
              <el-checkbox
                v-for="(show, showIndex) in col.showList"
                :key="col.prop + '-' + showIndex"
                v-model="col.showList[showIndex]"
                :label="col.label + String(showIndex + 1)"
              />
            </div>
            <div v-else>
              <el-checkbox v-model="col.show" :label="col.label" />
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="gva-table-box">
      <el-table
        v-if="tableCols !== undefined && tableData !== undefined"
        id="data-table"
        ref="multipleTable"
        :default-sort="{ prop: 'createdDate', order: 'descending' }"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        show-summary
        :summary-method="getSummries"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <!-- 当前项目状态 -->
        <el-table-column
          align="center"
          label="状态"
          width="100"
          prop="name"
          fixed="left"
        >
          <template #default="scope">
            <el-popover
              :visible="statusPopoverVisible[scope.$index]"
              placement="top-start"
              :width="220"
            >
              <template #reference>
                <el-button
                  plain
                  size="mini"
                  :type="getTagType(scope.row.status)"
                  @click="
                    handleShowProjectStatusPopover(
                      scope.$index,
                      scope.row.status
                    )
                  "
                >{{ getStatus(scope.row.status) }}</el-button>
              </template>
              <div style="display: flex; gap: 16px; flex-direction: column">
                <p>将项目状态修改为:</p>
                <div
                  style="
                    display: flex;
                    justify-content: space-between;
                    align-items: center;
                  "
                >
                  <el-select
                    v-model="nowProjectStatus"
                    size="small"
                    style="width: 115px"
                  >
                    <el-option
                      v-for="item in statusOptions"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    />
                  </el-select>
                  <div>
                    <el-button
                      plain
                      size="mini"
                      circle
                      icon="close"
                      type="danger"
                      @click="statusPopoverVisible[scope.$index] = false"
                    />
                    <el-button
                      plain
                      size="mini"
                      icon="check"
                      type="success"
                      circle
                      @click="
                        handleChangeProjectStatus(scope.$index, scope.row)
                      "
                    />
                  </div>
                </div>
              </div>
            </el-popover>
          </template>
        </el-table-column>
        <!-- 项目名称 左固定 -->
        <el-table-column
          align="center"
          label="项目名称"
          prop="name"
          fixed="left"
          width="200"
          :show-overflow-tooltip="true"
        />
        <!-- 备案申请日期 -->
        <el-table-column
          align="center"
          label="备案申请日期"
          width="150"
          sortable
          prop="createdDate"
        >
          <template #default="scope">{{
            formatTableVal(scope.row.createdDate, 'date')
          }}</template>
        </el-table-column>
        // todo clp
        <template v-for="col in tableCols">
          <!-- multi -->
          <template v-if="col.colType === 'multi'">
            <!-- showlist such as [true, false, false] -->
            <template v-for="(show, showIndex) in col.showList">
              <template v-if="show">
                <el-table-column
                  :key="col.prop + '-' + showIndex"
                  :label="col.label + String(showIndex + 1)"
                  align="center"
                >
                  <el-table-column
                    v-for="s in col.sub"
                    :key="col.label + '-' + showIndex + '-' + s.label"
                    :prop="s.prop"
                    :label="s.label"
                    align="center"
                    show-overflow-tooltip
                    :width="s.width"
                  >
                    <template #default="scope">
                      <span v-if="scope.row[col.prop][showIndex] != undefined">
                        {{
                          formatTableVal(
                            scope.row[col.prop][showIndex][s.prop],
                            s.format
                          )
                        }}
                      </span>
                    </template>
                  </el-table-column>
                </el-table-column>
              </template>
            </template>
          </template>
          <template v-else-if="col.colType === 'group'">
            <el-table-column
              v-if="col.show"
              :key="col.label"
              align="center"
              :label="col.label"
            >
              <el-table-column
                v-for="s in col.sub"
                :key="col.label + '-' + s.label"
                :label="s.label"
                :prop="s.prop"
                align="center"
                show-overflow-tooltip
                :width="s.width"
              >
                <template #default="scope">
                  <span v-if="scope.row[s.prop] != undefined">
                    {{ formatTableVal(scope.row[s.prop], s.format) }}
                  </span>
                </template>
              </el-table-column>
            </el-table-column>
          </template>
          <!-- 其他的正常渲染 -->
          <template v-else>
            <el-table-column
              v-if="col.show"
              :key="col.label"
              :prop="col.prop"
              :label="col.label"
              align="center"
              :show-overflow-tooltip="true"
              :width="col.width"
              :fixed="col.fixed"
              :sortable="col.sortable"
            >
              <template #default="scope">
                <!-- todo format -->
                <span>
                  {{ formatTableVal(scope.row[col.prop], col.format) }}
                </span>
              </template>
            </el-table-column>
          </template>
        </template>
        <!-- 查看项目的预算和支出 -->
        <el-table-column
          align="center"
          label="预算和支出"
          fixed="right"
          width="70"
        >
          <template #default="parentScope">
            <el-popover
              placement="top-start"
              :width="500"
              trigger="click"
              border
            >
              <template #reference>
                <el-button
                  size="medium"
                  class="table-button"
                  type="text"
                  icon="view"
                />
              </template>
              <el-table
                :data="parentScope.row.incomeAndOutcome"
                border
                :summary-method="getIncomeAndOutcomBias"
                show-summary
              >
                <el-table-column
                  type="index"
                  :index="indexMethod"
                  width="50"
                  align="center"
                  fixed="left"
                />
                <!-- todo 再创建一个模板，注意prop -->
                <el-table-column
                  v-for="col in incomeAndOutcomeCols"
                  :key="col.label"
                  :prop="col.prop"
                  :label="col.label"
                  :width="col.width"
                  align="center"
                >
                  <template #default="scope">
                    <!-- todo format -->
                    <span>
                      {{ formatTableVal(scope.row[col.prop], col.format) }}
                    </span>
                  </template>
                </el-table-column>
              </el-table>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" width="60">
          <template #default="scope">
            <el-popover placement="left" :width="150" trigger="click">
              <template #reference>
                <el-button size="medium" type="text" icon="setting" />
              </template>
              <el-button
                style="padding-left: 10px"
                type="text"
                icon="edit"
                size="medium"
                @click="updateProjectFunc(scope.row)"
              >编辑项目信息</el-button>
              <el-button
                type="text"
                icon="upload"
                size="medium"
                class="table-button"
                @click="uploadProjectFilesFunc(scope.row)"
              >上传项目附件</el-button>
              <el-button
                type="text"
                icon="plus"
                size="medium"
                @click="createIncomeStreamFunc(scope.row)"
              >添加收入流水</el-button>
              <el-button
                type="text"
                icon="plus"
                size="medium"
                @click="createOutcomeStreamFunc(scope.row)"
              >添加支出流水</el-button>
            </el-popover>
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
  </div>
</template>

<script>
export default {
  name: 'Project',
}
</script>

<script setup>
import {
  deleteProjectByIds,
  getProjectList,
  changeProjectStatus,
  exportToExcel,
} from '@/api/project'
import {
  findManageSystemSetting,
  updateManageSystemSetting,
} from '@/api/manageSystemSetting.js'

// 全量引入格式化工具 请按需保留
import { saveAs } from '@/utils/download'
import { formatTableVal } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// todo ,multi项的最大长度设置为showList的长度，如：showList:[true, true, false]
// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const tableCols = ref([])
const isShowSearchBox = ref(false)
const isShowColCheckGroupBox = ref(false)
// 委托方、落地机构和合作方的最大个数
const maxFormRowNum = ref(3)
const incomeAndOutcomeCols = ref([
  { prop: 'pg', label: '专家课酬', format: 'amount', width: 120 },
  { prop: 'ph', label: '方案费', format: 'amount', width: 120 },
  { prop: 'pi', label: '网络研修教学费', format: 'amount', width: 120 },
  { prop: 'pj', label: '管理及工作人员劳务费', format: 'amount', width: 120 },
  { prop: 'pk', label: '考察、跟岗、线下指导费', format: 'amount', width: 120 },
  { prop: 'pl', label: '专家食宿', format: 'amount', width: 120 },
  { prop: 'pm', label: '专家及工作人员交通费', format: 'amount', width: 120 },
  { prop: 'pn', label: '学员伙食费', format: 'amount', width: 120 },
  { prop: 'po', label: '学员住宿费', format: 'amount', width: 120 },
  { prop: 'pp', label: '学员交通费', format: 'amount', width: 120 },
  { prop: 'pq', label: '场租', format: 'amount', width: 120 },
  {
    prop: 'pr',
    label: '课程资源建设费(人员支出)',
    format: 'amount',
    width: 120,
  },
  { prop: 'ps', label: '课程资源建设费(购买)', format: 'amount', width: 120 },
  {
    prop: 'pt',
    label:
      '培训资料费、办公用品费、保险费、医药费及其他(含购买标书、中标服务费等)',
    format: 'amount',
    width: 200,
  },
  {
    prop: 'pTotal',
    label: '项目余额',
    format: 'amount',
    width: 150,
  },
])
// todo xiubug.
const searchInfo = ref({
  page: 1,
  pageSize: 10,
  name: '',
  categories: '',
  projectCode: '',
  createdDateRange: undefined,
  trainStartDateRange: undefined,
  contractStartDateRange: undefined,
  client: [
    { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
    { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
    { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
  ],
  landingAgency: [
    { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
    { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
    { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
  ],
  partner: [
    { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
    { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
    { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
  ],
})
const manageSystemSettingID = ref(1)
// status
const nowProjectStatus = ref(0)
const statusPopoverVisible = ref([])
const statusOptions = [
  {
    value: 0,
    label: '立项',
    disabled: true,
  },
  {
    value: 1,
    label: '进行中',
  },
  {
    value: 2,
    label: '中止',
  },
  {
    value: 3,
    label: '作废',
  },
  {
    value: 4,
    label: '结项',
  },
]

// 预算和支出的索引
const indexMethod = (index) => {
  switch (index) {
    case 0:
      return '预算'
    case 1:
      return '支出'
    default:
      return index
  }
}

// 计算对应预算和支出的差
const getIncomeAndOutcomBias = (param) => {
  const { columns, data } = param
  const sums = []
  columns.forEach((column, index) => {
    if (index === 0) {
      sums[index] = '差值'
      return
    }
    // 每一列，转成number类型
    const values = data.map((item) => Number(item[column.property]))
    // 里面的每个数都是number，则
    if (!values.every((value) => isNaN(value))) {
      sums[index] = formatTableVal(values[0] - values[1], 'amount')
    } else {
      sums[index] = 'N/A'
    }
  })

  return sums
}

const sum = (values) => {
  return values.reduce((prev, curr) => {
    if (!isNaN(Number(curr))) {
      return parseFloat(Number(prev + curr).toFixed(2))
    } else {
      return prev
    }
  }, 0)
}

// 计算表格合计
const getSummries = (param) => {
  const { columns, data } = param
  const sums = []
  let index = 0
  tableCols.value.forEach((col) => {
    switch (col.colType) {
      case 'multi':
        col.showList.forEach((show, i) => {
          if (show) {
            col.sub.forEach((s) => {
              const values = data.map((item) => {
                // 当显示的multi项较多时，对multi比较少的数据，越界访问会导致undefined
                return item[col.prop][i] === undefined
                  ? NaN
                  : Number(item[col.prop][i][s.prop])
              })
              sums[index] =
                !values.every((value) => isNaN(value)) && s.format === 'amount'
                  ? formatTableVal(sum(values), 'amount')
                  : ''
              index++
            })
          }
        })
        break
      case 'group':
        col.sub.forEach((s) => {
          if (s.show) {
            const values = data.map((item) => Number(item[s.prop]))
            sums[index] =
              !values.every((value) => isNaN(value)) && s.format === 'amount'
                ? formatTableVal(sum(values), 'amount')
                : ''
            index++
          }
        })
        break
      case 'json':
        col.sub.forEach((s) => {
          if (s.show) {
            const values = data.map((item) => Number(item[col.prop][s.prop]))
            sums[index] =
              !values.every((value) => isNaN(value)) && s.format === 'amount'
                ? formatTableVal(sum(values), 'amount')
                : ''
            index++
          }
        })
        break
      default:
        if (col.show) {
          const values = data.map((item) => Number(item[col.prop]))
          sums[index] =
            !values.every((value) => isNaN(value)) && col.format === 'amount'
              ? formatTableVal(sum(values), 'amount')
              : ''
          index++
        }
        break
    }
  })
  // todo change 项目名称、备案申请日期和状态不在tableCols里面，所以要加上
  sums.unshift('合计', '', '', '')
  return sums
}

// 显示|隐藏搜索框
const handleShowSearchBox = () => {
  isShowSearchBox.value = true
}
const handleHideSearchBox = () => {
  isShowSearchBox.value = false
}
// 显示|隐藏列筛选
const handleShowColCheckGroupBox = () => {
  isShowColCheckGroupBox.value = true
}
const handleHideColCheckGroupBox = () => {
  isShowColCheckGroupBox.value = false
}
// 重置筛选条件
const onReset = () => {
  searchInfo.value = {
    page: 1,
    pageSize: 10,
    name: '',
    categories: '',
    projectCode: '',
    createdDateRange: undefined,
    trainStartDateRange: undefined,
    contractStartDateRange: undefined,
    client: [
      { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
      { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
      { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
    ],
    landingAgency: [
      { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
      { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
      { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
    ],
    partner: [
      { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
      { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
      { name: '', radio: 0, incomeAmount: 0, outcomeAmount: 0 },
    ],
  }
}

// 搜索
const onSubmit = () => {
  // date l and r
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
  // 把所有查询的参数结合到一个对象中。
  searchInfo.value.page = page.value
  searchInfo.value.pageSize = pageSize.value
  const table = await getProjectList({
    searchInfo: JSON.stringify(searchInfo.value),
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    // init po list
    statusPopoverVisible.value = []
    for (let i = 0; i < tableData.value.length; i++) {
      statusPopoverVisible.value.push(false)
    }
  }
}

// 查询表格模板
const getTableCols = async() => {
  //
  const manageSystemSetting = await findManageSystemSetting({
    ID: manageSystemSettingID.value,
  })
  if (manageSystemSetting.code === 0) {
    tableCols.value =
      manageSystemSetting.data.remanageSystemSetting.projectTableCols
  }
}
// 获取表格列模板
getTableCols()
// 获取表格数据
getTableData()

// 根据状态让标签显示不同的颜色
const getTagType = (status) => {
  switch (status) {
    case 0:
      return 'info'
    case 1:
      return 'primary'
    case 2:
      return 'warning'
    case 3:
      return 'danger'
    case 4:
      return 'success'
    default:
      return 'info'
  }
}
const getStatus = (status) => {
  switch (status) {
    case 0:
      return '立项'
    case 1:
      return '进行中'
    case 2:
      return '中止'
    case 3:
      return '作废'
    case 4:
      return '结项'
    default:
      return '立项'
  }
}

// 显示修改项目状态pop时触发
const handleShowProjectStatusPopover = (rowIndex, status) => {
  nowProjectStatus.value = status
  statusPopoverVisible.value[rowIndex] = true
  console.log(statusPopoverVisible.value, rowIndex)
}
const handleChangeProjectStatus = (rowIndex, row) => {
  ElMessageBox.confirm(
    `确定将项目：${row.name}的状态更新为[${getStatus(
      nowProjectStatus.value
    )}]吗？`,
    '提示',
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'info',
    }
  )
    .then(() => {
      const params = {
        projectId: row.ID,
        status: nowProjectStatus.value,
      }
      changeProjectStatus(params).then((res) => {
        if (res.code === 0) {
          getTableData()
          ElMessage.success('更新项目状态成功')
          statusPopoverVisible.value[rowIndex] = false
        }
      })
    })
    .catch(() => {
      ElMessage('项目状态未更新')
    })
}

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

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      colType: 'warning',
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
      colType: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

const createProjectFunc = async() => {
  router.push({
    name: 'projectForm',
  })
}

const uploadProjectFilesFunc = async(row) => {
  if (row.status === 2 || row.status === 3) {
    ElMessage.error(
      `项目处于 [${getStatus(row.status)}] 状态，无法上传项目附件`
    )
    return
  }
  router.push({
    name: 'projectFiles',
    query: {
      id: row.ID,
    },
  })
}

// 更新行
const updateProjectFunc = async(row) => {
  // todo 根据项目状态判断是否允许更新。如果已经不是立项状态了，就不允许。
  if (row.status === 2 || row.status === 3) {
    ElMessage.error(`项目处于 [${getStatus(row.status)}] 状态，无法编辑`)
    return
  }
  router.push({
    name: 'projectForm',
    query: {
      id: row.ID,
    },
  })
}

// 更新系统设置
const updateManageSystemSettingFunc = async() => {
  // todo ：这里只传projectCol，gorm只更新一个字段
  const res = await updateManageSystemSetting({
    ID: 1, // 必须是1
    projectTableCols: tableCols.value,
    incomeStreamTableCols: [],
    outcomeStreamTableCols: [],
  })
  if (res.code === 0) {
    ElMessage({
      colType: 'success',
      message: '表格列显示字段配置保存成功',
    })
    getTableCols()
    getTableData()
  }
}

// 导出
const handleExcelExport = async() => {
  const res = await exportToExcel({
    searchInfo: JSON.stringify(searchInfo.value),
  })
  console.log(res)
  if (res.status === 200) {
    saveAs(res.data, '导出数据.xlsx')
  }
}

// 创建收入流水
const createIncomeStreamFunc = (row) => {
  if (row.status === 2 || row.status === 3) {
    ElMessage.error(
      `项目处于 [${getStatus(row.status)}] 状态，无法录入收入流水`
    )
    return
  }
  router.push({
    name: 'incomeStreamForm',
    query: {
      project: JSON.stringify(row),
    },
  })
}
// 支出流水
const createOutcomeStreamFunc = (row) => {
  if (row.status === 2 || row.status === 3) {
    ElMessage.error(
      `项目处于 [${getStatus(row.status)}] 状态，无法录入支出流水`
    )
    return
  }
  router.push({
    name: 'outcomeStreamForm',
    query: {
      project: JSON.stringify(row),
    },
  })
}
</script>

<style scoped>
.el-form-item {
  margin-bottom: 0;
}
.box-top-line {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 10px;
}
.box-top-line span {
  font-size: 16px;
}
</style>
