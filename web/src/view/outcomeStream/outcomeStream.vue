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
            @click="createOutcomeStreamFunc"
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
              <el-form-item label="所属项目" label-width="100px">
                <el-input
                  v-model="searchInfo.projectName"
                  placeholder="请输入待匹配文本"
                  :style="{ width: '100%' }"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="入账日期">
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
              <el-form-item label="分账日期">
                <el-date-picker
                  v-model="searchInfo.splitAmountDateRange"
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
            <div v-else-if="col.colType === 'json'">
              <el-checkbox
                v-for="s in col.sub"
                :key="s.label + '-' + s.prop"
                v-model="s.show"
                :label="s.label"
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
        <!-- 项目名称 左固定 -->
        <el-table-column
          align="center"
          label="培训项目"
          prop="projectName"
          fixed="left"
          width="200"
          :show-overflow-tooltip="true"
        />
        <!-- 入账日期 -->
        <el-table-column
          align="center"
          label="支出日期"
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
          <!-- 对json类型的值 -->
          <template v-else-if="col.colType === 'json'">
            <template v-for="s in col.sub">
              <el-table-column
                v-if="s.show"
                :key="s.label"
                :prop="s.prop"
                :label="s.label"
                align="center"
                :show-overflow-tooltip="true"
                :width="s.width"
              >
                <template #default="scope">
                  <!-- todo format -->
                  <span>
                    {{ formatTableVal(scope.row[col.prop][s.prop], s.format) }}
                  </span>
                </template>
              </el-table-column>
            </template>
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
        <el-table-column align="center" label="操作" fixed="right" width="100">
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateOutcomeStreamFunc(scope.row)"
            >变更</el-button>
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
  name: 'OutcomeStream',
}
</script>

<script setup>
import {
  deleteOutcomeStreamByIds,
  getOutcomeStreamList,
} from '@/api/outcomeStream'
import {
  findManageSystemSetting,
  updateManageSystemSetting,
} from '@/api/manageSystemSetting.js'

// 全量引入格式化工具 请按需保留
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

const searchInfo = ref({
  page: 1,
  pageSize: 10,
  projectName: '',
  createdDateRange: undefined,
  client: [
    { name: '', radio: 0, amount: 0 },
    { name: '', radio: 0, amount: 0 },
    { name: '', radio: 0, amount: 0 },
  ],
  landingAgency: [
    { name: '', radio: 0, amount: 0 },
    { name: '', radio: 0, amount: 0 },
    { name: '', radio: 0, amount: 0 },
  ],
  partner: [
    { name: '', radio: 0, amount: 0 },
    { name: '', radio: 0, amount: 0 },
    { name: '', radio: 0, amount: 0 },
  ],
})
// 必须是1
const manageSystemSettingID = ref(1)

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
  // todo change 项目名称和入账日期不在tableCols里面，所以要加上
  sums.unshift('合计', '', '')
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
    projectName: '',
    createdDateRange: undefined,
    client: [
      { name: '', radio: 0, amount: 0 },
      { name: '', radio: 0, amount: 0 },
      { name: '', radio: 0, amount: 0 },
    ],
    landingAgency: [
      { name: '', radio: 0, amount: 0 },
      { name: '', radio: 0, amount: 0 },
      { name: '', radio: 0, amount: 0 },
    ],
    partner: [
      { name: '', radio: 0, amount: 0 },
      { name: '', radio: 0, amount: 0 },
      { name: '', radio: 0, amount: 0 },
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
  const table = await getOutcomeStreamList({
    searchInfo: JSON.stringify(searchInfo.value),
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 查询表格模板
const getTableCols = async() => {
  //
  const manageSystemSetting = await findManageSystemSetting({
    ID: manageSystemSettingID.value,
  })
  if (manageSystemSetting.code === 0) {
    // 是支出流水的模板
    tableCols.value =
      manageSystemSetting.data.remanageSystemSetting.outcomeStreamTableCols
  }
}
// 获取表格列模板
getTableCols()
// 获取表格数据
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
  const res = await deleteOutcomeStreamByIds({ ids })
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
// 前往培训项目选择一个项目后再创建收入流水
const createOutcomeStreamFunc = async() => {
  ElMessageBox.alert(
    '请前往“培训项目列表”中选择该流水所属的培训项目',
    '需选择待录入流水所属的项目',
    {
      confirmButtonText: '好的，这就去',
      callback: () => router.push({ name: 'project' }),
    }
  )
}

// 更新行
const updateOutcomeStreamFunc = async(row) => {
  router.push({
    name: 'outcomeStreamForm',
    query: {
      id: row.ID,
    },
  })
}

// 更新系统设置
const updateManageSystemSettingFunc = async() => {
  // 这里只传outcomeSteamTableCol，gorm只更新一个字段
  const res = await updateManageSystemSetting({
    ID: 1, // 必须是1
    projectTableCols: [],
    incomeStreamTableCols: [],
    outcomeStreamTableCols: tableCols.value,
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
const handleExcelExport = () => {}
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
