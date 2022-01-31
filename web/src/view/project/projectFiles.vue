<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          size="mini"
          type="primary"
          icon="upload"
          @click="createProjectFileFunc"
        >上传</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？删除后无法撤销</p>
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
      </div>
      <!-- 查看某个类型的附件 -->
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="projectFileRecordTableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column label="名称" prop="fileName" align="center" />
        <el-table-column label="上传日期" prop="CreatedAt" align="center">
          <template #default="scope">
            <span>{{ formatTableVal(scope.row.CreatedAt, 'dateTime') }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="按钮组"
          fixed="right"
          width="200"
        >
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="downloadFileFunc(scope.row)"
            >下载
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[5, 10, 20, 50]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <!-- 选择某个类型的附件进行查看 -->
    <el-dialog
      v-model="isShowDialog"
      :title="`选择该项目的某个类型的附件进行查看`"
      width="60%"
      center
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div class="gva-table-box">
        <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
        >
          <el-table-column
            label="文件类型"
            prop="fileType"
            align="center"
            width="150"
          />
          <el-table-column
            label="文件描述"
            prop="discription"
            align="center"
            width="600"
          />
          <el-table-column align="center" label="操作" width="100">
            <template #default="scope">
              <el-button
                type="text"
                icon="edit"
                size="small"
                class="table-button"
                @click="showProjectFilesFunc(scope.row)"
              >查看
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-dialog>
    <!-- 上传附件 -->
    <el-dialog
      v-model="isShowUploadFilesDialog"
      title="上传附件"
      width="60%"
      center
    >
      <div class="display:flex">
        <el-upload
          action=""
          multiple
          :auto-upload="false"
          :on-change="(file, fileList) => handleFileChange(file, fileList)"
          :on-remove="(file, fileList) => handleFileRemove(file, fileList)"
          :file-list="filesToBeUpload"
          :show-file-list="false"
        >
          <template #trigger>
            <el-button
              size="small"
              type="primary"
              style="margin-left: 20px"
            >选取文件</el-button>
          </template>
          <el-button
            style="margin-left: 20px"
            size="small"
            type="success"
            :disabled="filesToBeUpload.length === 0"
            @click="uploadProjectFilesFunc()"
          >上传文件</el-button>
          <template #tip>
            <div class="el-upload__tip" style="margin-left: 20px">
              文件大小不超过200MB
            </div>
          </template>
        </el-upload>
        <!-- show files -->
        <div class="gva-table-box">
          <el-table
            border
            stripe
            style="width: 100%"
            fit
            highlight-current-row
            :data="filesToBeUpload"
          >
            <!-- 添加索引 -->
            <el-table-column
              type="index"
              width="50px"
              label="序号"
              align="center"
            />
            <!-- 文件名 -->
            <el-table-column
              width="500px"
              label="文件名"
              prop="name"
              align="center"
            />
            <el-table-column
              label="文件大小"
              prop="size"
              width="100px"
              align="center"
            >
              <template #default="scope">
                <span>{{ formatTableVal(scope.row.size, 'fileSize') }}</span>
              </template>
            </el-table-column>
            <el-table-column label="上传进度" align="center" width="200">
              <template #default="scope">
                <el-progress
                  :text-inside="true"
                  style="margin-top: 12px"
                  :stroke-width="10"
                  :percentage="scope.row.percentage"
                  :status="scope.row.percentage >= 100 ? 'success' : undefined"
                />
              </template>
            </el-table-column>
          </el-table>
        </div></div></el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Project',
}
</script>

<script setup>
// 全量引入格式化工具 请按需保留
import { formatTableVal } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, watch } from 'vue'
import SparkMD5 from 'spark-md5'

import {
  uploadProjectChunk,
  mergeProjectFileChunk,
  getProjectFileRecordList,
  deleteProjectFileByIds,
  downloadFile,
} from '@/api/project'

import { useRoute } from 'vue-router'

const route = useRoute()
const isShowDialog = ref(true)
const isShowUploadFilesDialog = ref(false)

// 哪个项目、哪个类型的附件
const fileTypeID = ref()
const projectID = ref()
// 上传文件的最大大小
const maxSize = 200 * 1024 * 1024
const uploadedNum = ref(0)
const successUploadNum = ref(0)

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
// 用来显示所有附件类型的静态表格数据
const tableData = ref([
  {
    ID: 1,
    fileType: '备案材料',
    discription:
      '办班申请表、责任书、培训方案、文件通知、项目审批表、专家论证报告、预算表（测算版）、合同（WORD版及签章扫描件）、培训手册',
  },
  {
    ID: 2,
    fileType: '物资准备',
    discription: '物资需求、办公用品申请表、课室申请、人员需求、疫情项目报备',
  },
  { ID: 3, fileType: '问卷调查', discription: '训前（培训需求）、训中、训后' },
  {
    ID: 4,
    fileType: '经费及支出',
    discription:
      '预算（正式版）、决算（正式版）、发票扫描件、学生助理信息、专家课酬、培训费用发票信息、日常报销、学员保单、学员住宿发票、学员饭卡转账单',
  },
  { ID: 5, fileType: '培训宣传', discription: '集体照、新闻稿' },
  { ID: 6, fileType: '学员档案资料', discription: '学员名单、签到表、考勤表' },
  { ID: 7, fileType: '各类函', discription: '各类函' },
  {
    ID: 8,
    fileType: '结业证书',
    discription: '证书登记表、各类证书模版',
  },
  {
    ID: 9,
    fileType: '培训成果',
    discription: '学员美篇、班主任总结、委托单位意见表',
  },
])
// 待上传的文件列表
const filesToBeUpload = ref([])
const projectFileRecordTableData = ref([])
// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  // 这里不是直接把行赋给form，而是又查询了一次？
  if (route.query.id) {
    projectID.value = route.query.id
  } else {
    ElMessage.error('初始化所属项目ID失败！')
  }
}
init()

// ----------------------------显示附件表格部分------------------------------------
const getTableData = async() => {
  const params = {
    page: page.value,
    pageSize: pageSize.value,
    projectID: projectID.value,
    fileTypeID: fileTypeID.value,
  }
  const table = await getProjectFileRecordList(params)
  console.log(params, table)
  if (table.code === 0) {
    projectFileRecordTableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

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
  const res = await deleteProjectFileByIds({ ids })
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

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
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

// 选中一个类型以后显示该类型的文件
const showProjectFilesFunc = (row) => {
  isShowDialog.value = false
  fileTypeID.value = row.ID
  // 获取项目附件列表
  getTableData()
}
const createProjectFileFunc = () => {
  isShowUploadFilesDialog.value = true
}

// upload component
const handleFileChange = (file, fileList) => {
  // 只有文件改变的时候才继续下面的操作
  if (file.status !== 'ready') {
    return
  }
  // todo: check file size
  if (file.size > maxSize) {
    ElMessage.error(
      `文件：${file.name}大于${formatTableVal(maxSize, 'fileSize')},无法上传`
    )
    fileList.pop()
  }
  let existFile = fileList
    .slice(0, fileList.length - 1)
    .find((f) => f.name === file.name)
  if (existFile) {
    ElMessage.error(`文件：${file.name} 已在待上传列表中`)
    fileList.pop()
  }
  existFile = projectFileRecordTableData.value.find(
    (f) => f.fileName === file.name
  )
  if (existFile) {
    ElMessage.error(`文件：${file.name} 已存在`)
    fileList.pop()
  }
  filesToBeUpload.value = fileList
}
const handleFileRemove = (file, fileList) => {
  filesToBeUpload.value = fileList
}

// 上传一个项目文件到服务器（断点续传）
const uploadFile = async(file) => {
  // file= {
  //   raw: File,
  //   percentage: 0,
  //   size : ..,
  //   ...
  // }
  const formDataList = [] // 分片存储的一个池子 丢全局
  const fileR = new FileReader() // 创建一个reader用来读取文件流
  fileR.readAsArrayBuffer(file.raw) // 把文件读成ArrayBuffer  主要为了保持跟后端的流一致
  fileR.onload = (e) => {
    // 读成arrayBuffer的回调 e 为方法自带参数 相当于 dom的e 流存在e.target.result 中
    const blob = e.target.result
    const spark = new SparkMD5.ArrayBuffer() // 创建md5制造工具 （md5用于检测文件一致性 这里不懂就打电话问我）
    spark.append(blob) // 文件流丢进工具
    const fileMd5 = spark.end() // 工具结束 产生一个a 总文件的md5
    const FileSliceCap = 2 * 1024 * 1024 // 分片字节数
    let start = 0 // 定义分片开始切的地方
    let end = 0 // 每片结束切的地方a
    let i = 0 // 第几片
    while (end < file.size) {
      // 当结尾数字大于文件总size的时候 结束切片
      start = i * FileSliceCap // 计算每片开始位置
      end = (i + 1) * FileSliceCap // 计算每片结束位置
      var fileSlice = file.raw.slice(start, end) // 开始切  file.slice 为 h5方法 对文件切片 参数为 起止字节数
      const formData = new window.FormData() // 创建FormData用于存储传给后端的信息
      formData.append('fileMd5', fileMd5) // 存储总文件的Md5 让后端知道自己是谁的切片
      formData.append('file', fileSlice) // 当前的切片
      formData.append('chunkNumber', i) // 当前是第几片
      formData.append('fileName', file.name) // 当前文件的文件名 用于后端文件切片的命名  formData.appen 为 formData对象添加参数的方法
      formDataList.push({ key: i, formData }) // 把当前切片信息 自己是第几片 存入我们方才准备好的池子
      i++
    }
    // 上传成功切片个数
    let uploadedChunks = 0
    // 把前面得到的切片数组一个个上传上去
    formDataList &&
      formDataList.forEach((item) => {
        const fileR = new FileReader() // 功能同上
        const fileF = item.formData.get('file')
        item.formData.append('chunkTotal', uploadedChunks)
        fileR.readAsArrayBuffer(fileF)
        fileR.onload = async(e) => {
          const spark = new SparkMD5.ArrayBuffer()
          spark.append(e.target.result)
          item.formData.append('chunkMd5', spark.end()) // 获取当前切片md5 后端用于验证切片完整性

          // 上传到db
          const fileRe = await uploadProjectChunk(item.formData)
          if (fileRe.code !== 0) {
            return
          }
          uploadedChunks++
          // 更新百分数
          file.percentage = parseFloat(
            ((uploadedChunks * 100) / formDataList.length).toFixed(2)
          )
          if (uploadedChunks === formDataList.length) {
            // 切片传完以后 合成文件
            const params = {
              fileName: file.name,
              fileMd5: fileMd5,
              projectID: projectID.value,
              fileTypeID: fileTypeID.value,
            }
            const res = await mergeProjectFileChunk(params)
            // 上传成功的切片
            if (res.code === 0) {
              ++successUploadNum.value
            }
            ++uploadedNum.value
          }
        }
        // 上传一个切片以后更新百分数
      })
  }
}

// 异步上传多个项目文件
const uploadProjectFilesFunc = async() => {
  await Promise.all(
    filesToBeUpload.value.map(async(file) => {
      // 分片上传
      await uploadFile(file)
    })
  )
}
// Get
watch(uploadedNum, () => {
  if (
    uploadedNum.value === filesToBeUpload.value.length &&
    uploadedNum.value !== 0
  ) {
    ElMessageBox.confirm(
      `上传完成，已上传${successUploadNum.value}/${filesToBeUpload.value.length}个文件`,
      'success',
      {
        confirmButtonText: '查看附件',
        cancelButtonText: '继续上传',
        type: 'success',
      }
    )
      .then(() => {
        getTableData()
        isShowUploadFilesDialog.value = false
        filesToBeUpload.value = []
        successUploadNum.value = 0
        uploadedNum.value = 0
      })
      .catch(() => {
        getTableData()
        filesToBeUpload.value = []
        successUploadNum.value = 0
        uploadedNum.value = 0
      })
  }
})

const saveAs = (data, fileName) => {
  const blob = new Blob([data])
  if ('download' in document.createElement('a')) {
    // 不是IE浏览器
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.style.display = 'none'
    link.href = url
    link.setAttribute('download', fileName)
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link) // 下载完成移除元素
    window.URL.revokeObjectURL(url) // 释放掉blob对象
  } else {
    // IE 10+
    window.navigator.msSaveBlob(blob, fileName)
  }
}
// 下载文件
const downloadFileFunc = async(row) => {
  const params = {
    fileName: row.fileName,
    projectID: projectID.value,
    fileTypeID: fileTypeID.value,
  }
  const res = await downloadFile(params)
  if (res.status === 200) {
    saveAs(res.data, row.fileName)
  }
}
</script>
