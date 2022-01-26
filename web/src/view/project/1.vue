<template>
  <!-- table item filter -->
  <div class="table-item-filter">
    <div class="table-item-filter-line1">
      <div class="line1-left">
        <el-button type="text">
          <div class="display:flex;align-items:flex-end;">
            <!-- <img src="../../public/icon-tabletitle.png" alt="" /> -->
            <span>{{ tableTitle + '列表' }}</span>
          </div>
        </el-button>
        <span>{{ '显示并列出目前所有的' + tableTitle + '列表' }}</span>
      </div>
    </div>
    <!-- search by keyword, and add\upload\export data -->
    <div class="table-item-filter-line2">
      <div class="line2-left">
        <p class="line2-left-item">选择列名进行筛选：</p>
        <!-- select -->
        <el-select
          v-model="seletedRowName"
          placeholder="选择列名"
          size="mini"
          style="width: 150px"
          class="line2-left-item"
        >
          <el-option
            v-for="(item, index) in selectRowNameList"
            :key="index"
            :value="item.prop"
            :label="item.label"
          ></el-option>
        </el-select>
        <!-- search keyword -->
        <el-date-picker
          v-if="
            seletedRowName.includes('Date') || seletedRowName.includes('Time')
          "
          :disabled="seletedRowName == undefined || seletedRowName.length == 0"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          v-model="createdDateKeyword"
          size="mini"
          class="line2-left-item"
        >
        </el-date-picker>
        <el-input
          v-else-if="
            !seletedRowName.includes('Date') && !seletedRowName.includes('Time')
          "
          :disabled="seletedRowName == undefined || seletedRowName.length == 0"
          placeholder="输入关键字搜索"
          v-model="keyword"
          style="width: 250px"
          size="mini"
          class="line2-left-item"
          clearable
        >
        </el-input>
      </div>
      <div class="line2-right" v-if="!isReadonly">
        <el-upload
          id="upload-excel"
          action=""
          name="excelFile"
          :show-file-list="false"
          accept="application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
          :before-upload="(file) => beforeUpload(file, excelNameRegExp)"
          :http-request="uploadExcel"
          :limit="1"
        >
          <el-button size="mini" icon="el-icon-upload" :disabled="true"
            >导入暂不开放</el-button
          >
        </el-upload>
        <el-button
          id="export-to-excel"
          size="mini"
          @click="exportToExcel(tableCols, pagedAndFiltedTableData)"
          icon="el-icon-download"
          >导出</el-button
        >
        <el-button
          icon="el-icon-circle-plus-outline"
          @click="handleAdd"
          size="mini"
          >添加</el-button
        >
        <el-button icon="el-icon-setting" @click="handleShowCols" size="mini"
          >显示字段设置</el-button
        >
      </div>
    </div>
  </div>
  <!-- checkbox to select columns to show -->
  <div class="table-col-selector" v-if="isShowColCheckGroup">
    <div class="top">
      <div class="left">表格显示字段设置:</div>
      <div class="right">
        <el-button @click="handleShowAllCols" size="mini" icon="el-icon-thumb"
          >全选</el-button
        >
        <el-button
          @click="handleCancelShowAllCols"
          size="mini"
          icon="el-icon-refresh"
          >取消全选</el-button
        >
        <el-button size="mini" icon="el-icon-turn-off" @click="handleShowCols"
          >隐藏</el-button
        >
        <el-button
          size="mini"
          icon="el-icon-lock"
          @click="handleSaveColSettings"
          >保存设置</el-button
        >
      </div>
    </div>
    <div class="checkboxgroup">
      <div class="checkbox-item" v-for="(col, index) in tableCols" :key="index">
        <div
          v-if="
            tableData[col.maxLengthIndex] != undefined &&
            tableData[col.maxLengthIndex][col.prop] != undefined &&
            (col.type === 'multi' || col.type === 'object')
          "
        >
          <div
            v-for="(dataItem, dataIndex) in tableData[col.maxLengthIndex][
              col.prop
            ]"
            :key="dataIndex"
          >
            <el-checkbox
              v-for="(subTitleItem, subIndex) in col.subTitle"
              :key="subIndex"
              :label="
                labelFormatter(
                  col.insertIndex,
                  subTitleItem.label,
                  (dataIndex + 1).toString(),
                  col.type
                )
              "
              v-model="subTitleItem.isShow"
              @change="handleCheckboxChange"
            ></el-checkbox>
          </div>
        </div>
        <div v-else>
          <el-checkbox
            :label="col.label"
            v-model="col.isShow"
            @change="handleCheckboxChange"
          ></el-checkbox>
        </div>
      </div>
    </div>
  </div>

  <div class="table-main">
    <el-table
      border
      :show-summary="!isReadonly"
      :summary-method="getSummaries"
      ref="tableRef"
      :data="pagedAndFiltedTableData"
      style="width: 100%"
      max-height="1000px"
      stripe
      :id="tableID"
      @row-dblclick="handleEdit"
      @selection-change="handleSelectionChange"
      row-key="key"
      @sort-change="handleSortChange"
    >
      <!-- 添加索引 -->
      <el-table-column
        type="selection"
        width="80"
        align="center"
        v-if="!isReadonly"
      >
      </el-table-column>
      <!-- 表头 -->
      <template v-for="(col, index) in tableCols" :key="index">
        <!-- 日期排序列 -->
        <template v-if="col.prop == 'createdDate' && col.isShow">
          <el-table-column
            :prop="col.prop"
            :label="col.label"
            align="center"
            width="150px"
            sortable="custorm"
            :sort-method="tableSort"
          >
            <template v-slot="scope">
              <!-- 正常显示，row为当前某一行的数据对象，col.prop为数据对象的某个属性， scope.row[col.prop]为某行某列的一个值-->
              <span v-if="scope.row[col.prop] != undefined">{{
                scope.row[col.prop]
              }}</span>
            </template>
          </el-table-column>
        </template>
        <!-- 查看项目资源 -->
        <el-table-column
          label="附件"
          align="center"
          v-else-if="
            col.prop == 'uploadProjectResources' && !isReadonly && col.isShow
          "
          width="50"
        >
          <template v-slot="scope">
            <el-button
              class="el-button-projectresource"
              type="text"
              size="mini"
              @click="uploadProjectResources(scope.row)"
              >查看</el-button
            >
          </template>
        </el-table-column>
        <!-- el-upload -->
        <template v-else-if="col.prop == 'name' && col.isShow">
          <el-table-column label="文件" align="center" width="auto">
            <template v-slot="scope">
              <el-button
                type="text"
                @click="downloadResource(scope.row)"
                class="el-button-download"
                >{{ scope.row.name }}</el-button
              >
            </template>
          </el-table-column>
        </template>
        <!-- if multi -->
        <!-- 在vue渲染机制中，刚开始tableData是没有数据的，此时不能渲染，否则会报错TypeError: Cannot read property 'xxx' of undefined： -->
        <!-- object相当于multi-item只有一个 -->
        <template
          v-else-if="
            (col.type === 'multi' || col.type === 'object') &&
            tableData.length > 0 &&
            tableData[col.maxLengthIndex] != undefined &&
            tableData[col.maxLengthIndex][col.prop] != undefined
          "
        >
          <!-- here the dataItem is useless, only need the dataIndex -->
          <!-- 比如，落地机构对应的列数为所有项目中落地机构最多的那个流水的落地机构个数 -->
          <template
            v-for="(dataItem, dataIndex) in tableData[col.maxLengthIndex][
              col.prop
            ]"
            :key="dataIndex"
          >
            <!-- e.g label=落地机构1分成比例 -->
            <!-- 当它们处于同一节点，v-for 的优先级比 v-if 更高，这意味着 v-if 将分别重复运行于每个 v-for 循环中。当你想为仅有的一些项渲染节点时，这种优先级的机制会十分有用 -->
            <el-table-column
              v-for="(subTitleItem, subIndex) in columnFilter(col.subTitle)"
              :key="subIndex"
              :prop="subTitleItem.prop"
              :label="
                labelFormatter(
                  col.insertIndex,
                  subTitleItem.label,
                  (dataIndex + 1).toString(),
                  col.type
                )
              "
              align="center"
              :width="
                flexColumnWidth(
                  col.prop,
                  dataIndex,
                  subTitleItem.prop,
                  'multi',
                  subTitleItem.formatType
                )
              "
            >
              <!-- e.g tableData.landingAgencies[0].agencyName -->
              <template v-slot="scope">
                <!-- 格式化显示，如加%和千分位 -->
                <!-- 如果当前流水的multi项的长度小于最大长度，不显示 -->
                <span
                  v-if="
                    scope.row[col.prop] != undefined &&
                    scope.row[col.prop][dataIndex] != undefined &&
                    scope.row[col.prop].length > 0
                  "
                >
                  {{
                    numberFormater(
                      scope.row[col.prop][dataIndex][subTitleItem.prop],
                      subTitleItem.formatType
                    )
                  }}</span
                >
              </template>
            </el-table-column>
          </template>
        </template>
        <!-- not -->
        <template v-else>
          <el-table-column
            :prop="col.prop"
            :label="col.label"
            align="center"
            :width="flexColumnWidth(col.prop, 0, '', '', col.formatType)"
            v-if="col.isShow"
            :fixed="col.isFixed == true"
          >
            <template v-slot="scope">
              <!-- 正常显示，row为当前某一行的数据对象，col.prop为数据对象的某个属性， scope.row[col.prop]为某行某列的一个值-->
              <!-- <div
              v-if="col.format != undefined"
              v-format="`${col.format}`"
            ></div> -->
              <span v-if="scope.row[col.prop] != undefined">
                {{ numberFormater(scope.row[col.prop], col.formatType) }}
              </span>
            </template>
          </el-table-column>
        </template>
      </template>
    </el-table>
    <div class="table-bottom">
      <!-- 分页 -->
      <div class="table-bottom-left" v-if="!isReadonly">
        <p>{{ `已选择 ${selectedRowCount} 项` }}</p>
        <el-button size="mini" @click="handleDelete">批量删除</el-button>
      </div>
      <div class="table-bottom-right">
        <el-pagination
          class="my-el-pagination"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          :page-size="pageSize"
          :page-sizes="[5, 10, 15, 20]"
          layout="slot, sizes, prev, pager, next, jumper"
          background
          :total="total"
        >
          <span class="total-span">{{ totalText }} </span>
        </el-pagination>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable vue/no-unused-vars */
import { computed, defineComponent, PropType, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  numberFormater,
  labelFormatter,
  saveAs,
  uploadExcel,
  exportToExcel,
  beforeUpload,
  numberToThousands,
} from '@/functions/general'
import store from '@/store'
import { Resource, SystemSetting } from '@/interface/entity'
import { TablePattern } from '@/interface/general'
import { useRoute } from 'vue-router'
import { constObjs, fileNamePatternRegExp } from '@/const/const'
import { updateSystemSetting, querySystemSetting } from '@/api/SystemSettingApi'
export default defineComponent({
  name: 'MyTable',
  props: {
    tableTitle: {
      type: String,
      required: true,
    },
    tableData: {
      type: Array,
      required: true,
    },
    tableCols: {
      type: Array as PropType<TablePattern[]>,
      required: true,
    },
    tableID: {
      type: String,
      required: true,
    },
    tableRowAdd: {
      type: Function as PropType<() => {}>,
    },
    tableRowDelete: {
      // 指定函数类型
      type: Function as PropType<(arg1: object) => {}>,
    },
    tableRowEdit: {
      type: Function as PropType<(arg1: object, arg2: number) => {}>,
    },
    uploadProjectResource: {
      type: Function as PropType<(arg1: object) => {}>,
    },
    isReadonly: {
      type: Boolean,
      required: true,
    },
    tableSort: {
      type: Function as PropType<(arg1: object, arg2: object) => {}>,
    },
  },
  setup(props) {
    const currentPage = ref<number>(1)
    const pageSize = ref<number>(5)
    // will be used when deleting rows
    const selectedRows = ref<object[]>()
    // keyword to search table item
    const keyword = ref<string>('')
    // when search by createdDate
    const createdDateKeyword = ref<Date[]>([])
    // the row name to search
    const seletedRowName = ref<string>('')
    // route
    const route = useRoute()
    //
    const isShowColCheckGroup = ref<boolean>(false)
    //
    const nowTableCols = ref<TablePattern[]>([])
    //
    const isShowAllCols = ref<boolean>(false)

    // 检测到路由变化时清除关键字
    watch(
      () => route.matched,
      (oldVal, newVal) => {
        if (oldVal != newVal) {
          // 更新tablePattern
          nowTableCols.value = JSON.parse(
            JSON.stringify(props.tableCols)
          ) as TablePattern[]
          seletedRowName.value = ''
          keyword.value = ''
          // console.log(props.tableCols);
        }
      }
    )
    // init tableCol
    nowTableCols.value = JSON.parse(
      JSON.stringify(props.tableCols)
    ) as TablePattern[]

    const columnFilter = (titles: TablePattern[]) => {
      return titles.filter((title) => title.isShow)
    }

    //
    const flexColumnWidth = (
      prop: string,
      subIndex: number,
      subProp: string,
      type: string,
      formatType: string
    ) => {
      let flexWidth = 0
      let maxLengthColumn = ''
      // traverse the tableData row, and get data by prop, then find the max length
      if (type != 'multi') {
        ;(props.tableData as object[]).forEach((data) => {
          const tmp = numberFormater(
            String(data[prop as keyof object]),
            formatType
          )

          if (tmp.length > maxLengthColumn.length) {
            maxLengthColumn = tmp
          }
        })
      } else {
        ;(props.tableData as object[]).forEach((data) => {
          try {
            const tmp = numberFormater(
              String(data[prop as keyof object][subIndex][subProp]),
              formatType
            )
            if (tmp == 'undefined') {
              return 100
            } else if (tmp.length > maxLengthColumn.length) {
              maxLengthColumn = tmp
            }
          } catch (error) {
            // console.log(data[prop as keyof object][subIndex]);
            return 100
          }
        })
      }
      // 计算宽度
      if (maxLengthColumn.includes('￥')) {
        flexWidth += 32
      }
      for (const ch of maxLengthColumn) {
        // chinese
        if ((ch >= '\u4e00' && ch <= '\u9fa5') || ch == '￥') {
          flexWidth += 17
        } else flexWidth += 10
      }
      if (flexWidth < 100) return 100
      if (flexWidth > 500) return 500
      return flexWidth
    }

    // 排序的时候不针对当前页 而是所有数据排序
    const handleSortChange = (column: object) => {
      // 排序后跳转当前页
      console.log(column)
      currentPage.value = 1
      //
      const prop = column['prop' as keyof object],
        order = column['order' as keyof object]
      // compare func（by date）
      const compareFunc = (a: object, b: object): number => {
        if (
          a[prop as keyof object] == undefined ||
          b[prop as keyof object] == undefined
        ) {
          return -1
        }
        return a[prop as keyof object] > b[prop as keyof object] ? 1 : -1
      }
      if (order == 'descending') {
        ;(props.tableData as object[]).sort((a, b) => compareFunc(a, b))
      } else if (order == 'ascending') {
        ;(props.tableData as object[])
          .sort((a, b) => compareFunc(a, b))
          .reverse()
      }
    }

    // pagination
    const handleCurrentChange = (val: number) => {
      currentPage.value = val
    }
    const handleSizeChange = (val: number) => {
      currentPage.value = 1
      pageSize.value = val
    }

    const uploadProjectResources = (row: object) => {
      if (props.uploadProjectResource != undefined) {
        props.uploadProjectResource(row)
      }
    }

    const handleAdd = () => {
      if (props.tableRowAdd != undefined) {
        props.tableRowAdd()
      }
    }

    // table
    const handleEdit = (row: object, index: number) => {
      // 调用父组件传来的编辑行的函数
      if (props.tableRowEdit != undefined) {
        props.tableRowEdit(row, index)
      }
    }
    const handleDelete = () => {
      if (props.tableRowDelete != undefined) {
        // 显示消息提示，并调用父组件的删除行的函数
        if (
          selectedRows.value == undefined ||
          selectedRows.value.length == undefined
        ) {
          ElMessage.error('请至少选择一个要删除的记录')
        } else {
          ElMessageBox.confirm(
            `您确定要删除${selectedRows.value.length}条记录吗？删除后不可撤销`,
            '警告',
            {
              confirmButtonText: '删除',
              cancelButtonText: '取消',
              type: 'warning',
            }
          ).then(() => {
            // delete
            ;(props.tableRowDelete as Function)(selectedRows.value)
            // reset
            selectedRows.value = []
            //console.log(row, index);
          })
        }
      }
    }
    const tableRowFilter = (data: object[]) => {
      const filetedData = (data as object[]).filter((item) => {
        if (keyword.value != '') {
          const tmp = item[seletedRowName.value as keyof object]
          const tmpStr = tmp as string
          try {
            return tmpStr.toLowerCase().includes(keyword.value.toLowerCase())
          } catch (error) {
            console.log(tmpStr, typeof tmp)
            return true
          }
        } else if (
          seletedRowName.value.includes('Date') ||
          seletedRowName.value.includes('Time')
        ) {
          if (
            createdDateKeyword.value != null &&
            createdDateKeyword.value.length > 0
          ) {
            // convert to Date so it can be compared with date object
            const nowDate = new Date(
              item[seletedRowName.value as keyof object] as string
            )
            return (
              nowDate >= createdDateKeyword.value[0] &&
              nowDate <= createdDateKeyword.value[1]
            )
          }
          // if now seleted a date range
          return true
        }
        // if without keyword
        return true
      })
      return filetedData == undefined ? [] : filetedData
    }

    const handleShowCols = () => {
      isShowColCheckGroup.value = !isShowColCheckGroup.value
    }

    const handleShowAllCols = () => {
      isShowAllCols.value = true
      props.tableCols.forEach((col) => {
        if (col.type != 'multi') {
          col.isShow = true
        } else {
          col.subTitle?.forEach((title) => (title.isShow = true))
        }
      })
    }

    const handleCancelShowAllCols = () => {
      isShowAllCols.value = false
      if (nowTableCols.value.length == 0) {
        return
      }
      props.tableCols.forEach((col, index) => {
        try {
          if (col.type != 'multi') {
            col.isShow = nowTableCols.value[index].isShow
          } else {
            col.subTitle?.forEach(
              (title, subIndex) =>
                // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
                (title.isShow =
                  nowTableCols.value[index].subTitle![subIndex].isShow)
            )
          }
        } catch (error) {
          console.log(col, nowTableCols.value[index])
        }
      })
    }

    const handleCheckboxChange = () => {
      if (!isShowAllCols.value) {
        nowTableCols.value = JSON.parse(
          JSON.stringify(props.tableCols)
        ) as TablePattern[]
      }
    }

    const handleSaveColSettings = async () => {
      switch (route.name) {
        // todo :pattern store in db but not code
        case 'projectsystem': {
          // change store
          constObjs.projectsTablePattern = JSON.parse(
            JSON.stringify(props.tableCols)
          ) as TablePattern[]
          constObjs.projectsTablePattern.forEach((item) => {
            if (item.type == 'multi') {
              item.subTitle?.forEach((title) => {
                title.maxLengthIndex = 0
                title.maxSubItemNum = 1
              })
            }
          })
          break
        }
        case 'incomestreamsystem': {
          constObjs.incomeStreamTablePattern = JSON.parse(
            JSON.stringify(props.tableCols)
          ) as TablePattern[]
          constObjs.incomeStreamTablePattern.forEach((item) => {
            if (item.type == 'multi') {
              item.subTitle?.forEach((title) => {
                title.maxLengthIndex = 0
                title.maxSubItemNum = 1
              })
            }
          })
          break
        }
        case 'outlaystreamsystem': {
          constObjs.outlayStreamTablePattern = JSON.parse(
            JSON.stringify(props.tableCols)
          ) as TablePattern[]

          break
        }
        case 'adminsystem': {
          constObjs.userPermissionTablePattern = JSON.parse(
            JSON.stringify(props.tableCols)
          ) as TablePattern[]
          break
        }
        default: {
          ElMessage.error('路由不匹配')
          return
        }
      }
      // console.log(tmpPattern, props.tableCols);
      const systemSetting = {
        guid: '123',
        projectTablePattern: JSON.stringify(constObjs.projectsTablePattern),
        incomeStreamTablePattern: JSON.stringify(
          constObjs.incomeStreamTablePattern
        ),
        outlayStreamTablePattern: JSON.stringify(
          constObjs.outlayStreamTablePattern
        ),
        adminTablePattern: JSON.stringify(constObjs.userPermissionTablePattern),
      } as SystemSetting
      await updateSystemSetting(systemSetting)
      await querySystemSetting('123')
    }

    // when using checkbox
    const handleSelectionChange = (selection: object[]) => {
      // save now selected cols
      selectedRows.value = selection
    }

    // only for download resource
    const downloadResource = (row: Resource) => {
      saveAs(row.url, row.name)
    }

    // to get the sum of currency
    // this method will be used twice when table changed. in this case, when first used, the data will change but pattern not, then a bug will occur
    const getSummaries = (params: object) => {
      // 只读时，不显示合计（筛选培训项目时）
      if (props.isReadonly == true) {
        return
      }
      // data
      // const data = params["data" as keyof object] as object[];
      const data = tableRowFilter(props.tableData as object[])
      // the array will be return
      const sums = [] as string[]
      // the number of columns which will be shown in the table
      let columnCount = 0
      props.tableCols.forEach((column) => {
        //
        // 将data中每一行的column.property对应的值转换成number作为数组返回
        const nowKey = column.prop
        let values = [] as number[]
        // 正常列
        if (column.type != 'multi') {
          // 选择当前显示的列
          if (column.isShow) {
            // 对所有行的某一列求和，如果是currency就求和显示，不是显示“”
            values = data.map((item) => Number(item[nowKey as keyof object]))
            sums[columnCount] = ''
            if (
              !values.every((value) => isNaN(value)) &&
              column.formatType == 'currency'
            ) {
              // float add and convert to currency format
              sums[columnCount] = numberToThousands(
                values.reduce((prev, curr) => {
                  if (!isNaN(Number(curr))) {
                    return parseFloat((prev + curr).toFixed(2))
                  } else {
                    return prev
                  }
                }, 0)
              )
            }
            columnCount++
          }
          // if multi column ,such landangAgencies and partners
          // i is the max num of multiitem such as laxxx
        } else {
          for (let i = 0; i < (column.maxSubItemNum as number); i++) {
            // when first called, the data and colpattern is not matched and  will return in this case
            try {
              // traverse the subtitle to get the key, and find the data by it
              column.subTitle?.forEach((title) => {
                if (title.isShow) {
                  values = data.map((item) => {
                    const nowMultiItem = (
                      item[nowKey as keyof object] as object[]
                    )[i]
                    return nowMultiItem == undefined
                      ? NaN
                      : nowMultiItem[title.prop as keyof object]
                  })
                  if (
                    !values.every((value) => isNaN(value)) &&
                    title.formatType == 'currency'
                  ) {
                    // float add and convert to currency format
                    sums[columnCount] = numberToThousands(
                      values.reduce((prev, curr) => {
                        if (!isNaN(Number(curr))) {
                          return parseFloat((prev + curr).toFixed(2))
                        } else {
                          return prev
                        }
                      }, 0)
                    )
                  } else {
                    sums[columnCount] = ''
                  }
                  columnCount++
                }
              })
            } catch (error) {
              return
            }
          }
        }
      })
      sums.unshift('合计')
      return sums.filter((item) => item != null && item != undefined)
    }

    return {
      tableRowFilter,
      //el-table data绑定的数据必须是确定的，因此需要在动态拿到确定数据后，再绑定到data上,否则无法多选
      pagedAndFiltedTableData: computed(() =>
        tableRowFilter(props.tableData as object[]).slice(
          (currentPage.value - 1) * pageSize.value,
          currentPage.value * pageSize.value
        )
      ),
      columnFilter,
      route,
      keyword,
      createdDateKeyword,
      seletedRowName,
      // 表格总行数
      selectedRows,
      isShowColCheckGroup,
      isShowAllCols,
      // 获取表列（不包含落地机构、合作方等）
      selectRowNameList: computed(() =>
        props.tableCols.filter((col) => {
          let flag = true
          if (
            col.type == 'multi' ||
            col.prop == 'uploadProjectResources' ||
            col.prop.includes('Radio')
          ) {
            flag = false
          }
          return flag
        })
      ),
      selectedRowCount: computed(() =>
        selectedRows.value?.length == undefined
          ? 0
          : selectedRows.value?.length.toString()
      ),
      total: computed(() => tableRowFilter(props.tableData as object[]).length),
      totalText: computed(
        () =>
          `共 ${props.tableData.length} 条，已筛选 ${
            tableRowFilter(props.tableData as object[]).length
          } 条`
      ),
      nowSystem: computed(() => store.state.nowSystem),
      flexColumnWidth,
      currentPage,
      pageSize,
      handleSortChange,
      handleSelectionChange,
      handleCurrentChange,
      handleSizeChange,
      numberFormater: computed(() => numberFormater),
      labelFormatter: computed(() => labelFormatter),
      excelNameRegExp: computed(() => fileNamePatternRegExp.excel),
      handleAdd,
      handleEdit,
      handleDelete,
      handleShowCols,
      handleShowAllCols,
      handleSaveColSettings,
      handleCancelShowAllCols,
      handleCheckboxChange,
      uploadProjectResources,
      downloadResource,
      uploadExcel,
      exportToExcel,
      beforeUpload,
      getSummaries,
    }
  },
})
</script>

<style>
.el-button-projectresource,
.el-button-edit {
  font-size: 13px;
  color: rgb(96, 98, 112);
  padding: 0;
}
.el-button-projectresource:hover,
.el-button-edit:hover {
  color: rgb(96, 98, 112);
  border-bottom: 1px solid rgb(96, 98, 112);
}
/* -------------------------table------------------------ */
/* -------------------------table------------------------ */
.table-main {
  margin-top: 20px;
  padding: 20px;
  background-color: #ffffff;
  box-shadow: -1px 1px 1px 0px #ccc, 1px -1px 1px 0px #ccc,
    -1px -1px 1px 0px #ccc, 1px 1px 1px 0px #ccc;
}
/* -------------table item filter-------------- */
/* -------------table item filter-------------- */
.table-item-filter,
.table-item-filter-line1,
.table-item-filter-line2,
.line1-left,
.line1-right,
.line2-left,
.line2-right {
  display: flex;
  align-items: center;
}
.table-item-filter {
  flex-direction: column;
  box-shadow: -1px 1px 1px 0px #ccc, 1px -1px 1px 0px #ccc,
    -1px -1px 1px 0px #ccc, 1px 1px 1px 0px #ccc;
}
/* line1 */
.table-item-filter-line1 {
  width: 100%;
  background-color: #f5f5f6;
}
.line1-left {
  margin-left: 20px;
}
.table-item-filter-line1 .el-button {
  font-size: 18px;
  font-weight: bold;
  font-family: 'Microsoft YaHei';
  color: #0d92d7;
}
.table-item-filter-line1 .el-button:hover {
  border-bottom: 3px solid #0d92d7;
}
.line1-left > span {
  font-size: 14px;
  color: #000000;
  margin-left: 20px;
}
.line1-right,
.line2-right {
  margin-left: auto;
  margin-right: 22px;
}
/* line2 */

.table-item-filter-line2 {
  width: 100%;
  background-color: #ffffff;
  padding: 15px 0;
}
.line2-left {
  margin-left: 22px;
}
.line2-left .line2-left-item {
  margin-right: 10px;
  font-size: 14px;
}
.line2-left p {
  font-size: 13px;
}
/* >>> 是vue的深度选择器，vue引用了第三方组件，需要在组件中局部修改第三方组件的样式，而又不想去除scoped属性造成组件之间的样式污染。此时只能通过 >>> ，穿透scoped。 */
.line2-left >>> .el-input__inner {
  border-radius: 0;
  border: 1px solid #cccccc;
  box-shadow: 0 0 1px 1px #cccccc;
}
.line2-right .el-button {
  border: 1px solid #0d92d7;
  font-size: 12px;
  color: #0eb7ef;
  margin-left: 12px;
  /* 修改自定义padding */
  padding: 2px 10px;
}

/* -----------------table col selector ------------------*/
.table-col-selector {
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
  margin-top: 20px;
  background: #ffffff;
  box-shadow: -1px 1px 1px 0px #ccc, 1px -1px 1px 0px #ccc,
    -1px -1px 1px 0px #ccc, 1px 1px 1px 0px #ccc;
}
.table-col-selector .top {
  display: flex;
  padding-left: 20px;
  align-items: center;
  background-color: #f5f5f6;
  height: 50px;
}
.top .left {
  font-size: 15px;
}
.top .right {
  margin: 0 20px 0 auto;
}
.top .right .el-button {
  font-size: 12px;
  color: #0eb7ef;
  border: 1px solid #0d92d7;
  /* 修改自定义padding */
  padding: 2px 10px;
}
.checkboxgroup {
  display: flex;
  flex-wrap: wrap;
  margin: 10px 20px;
}
.checkboxgroup .el-checkbox {
  margin: 5px 30px 5px 0;
  width: 150px;
}
/* 文字太长时换行 */
.el-checkbox__label {
  display: inline-grid;
  white-space: pre-line;
  word-wrap: break-word;
  overflow: hidden;
  line-height: 20px;
}

/* ------------------------------table-------------------- */
.el-button-download {
  color: #606266;
  padding: 0;
}
.el-button-download:hover {
  color: black;
  font-weight: bold;
  border-bottom: 1px solid black;
}
.table-bottom {
  display: flex;
  align-items: center;
  width: 100%;
  margin-top: 20px;
}
.table-bottom-left {
  display: flex;
  align-items: center;
  margin-left: 20px;
  height: 30px;
  color: #cccccc;
}
.table-bottom-left .el-button {
  border: 1px solid #0d92d7;
  font-size: 12px;
  color: #0eb7ef;
  margin-left: 12px;
  /* 修改自定义padding */
  padding: 2px 10px;
}

.table-bottom-right {
  margin-left: auto;
}
.el-table th.gutter {
  display: table-cell !important;
}
.el-pagination .total-span {
  font-weight: normal;
  color: #606266;
  margin-right: 10px;
}
/* 字体纯黑色 */
.el-table span,
.el-table__footer-wrapper tbody td,
.el-table__header-wrapper tbody td {
  color: #000000;
}
/* 表头样式 */
.el-table .el-table__header-wrapper tr th,
.el-table .el-table__fixed-header-wrapper tr th {
  background-color: #f2f8ff;
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB',
    'Microsoft YaHei', '微软雅黑', Arial, sans-serif;
  color: #20a6d3;
}
/* 表头不换行 */
/* .el-table .el-table__header .cell {
  text-overflow: unset !important;
  white-space: nowrap !important;
} */
.el-table .el-table__body tr:hover > td {
  background-color: rgb(234, 250, 255) !important;
}
/* .el-table td {
  border-bottom: 1px solid #606266;
} */
.el-table td,
.el-table th {
  padding: 11px 0;
}
.el-table--striped .el-table__body tr.el-table__row--striped td {
  background: rgb(245, 245, 246);
}

/*  滚动条的宽度 */
.el-table__body-wrapper::-webkit-scrollbar {
  /* 横向滚动条 */
  width: 20px;
  /* 纵向滚动条 必写 */
  height: 12px;
}

/*  滚动条的滑块 */
.el-table__body-wrapper::-webkit-scrollbar-thumb {
  background-color: #ddd;
  border-radius: 6px;
  min-height: 50px;
  background-clip: padding-box;
  width: 900px;
}
.el-table__body-wrapper::-webkit-scrollbar-thumb:hover {
  background-color: rgb(200, 200, 200);
}
</style>
