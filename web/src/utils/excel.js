/* eslint-disable no-empty */
import XLSX from 'yxg-xlsx-style'

// 如果单元格是日期类型
function datenum(v, date1904) {
  if (date1904) v += 1462
  const epoch = Date.parse(v)
  return (epoch - new Date(Date.UTC(1899, 11, 30))) / (24 * 60 * 60 * 1000)
}

// 将二维数组(如：data[0][0])变换成excel数据格式(如：A1)
function sheet_from_array_of_arrays(data) {
  const ws = {}
  const range = {
    s: {
      c: 10000000,
      r: 10000000
    },
    e: {
      c: 0,
      r: 0
    }
  }
  for (let R = 0; R !== data.length; ++R) {
    for (let C = 0; C !== data[R].length; ++C) {
      if (range.s.r > R) range.s.r = R
      if (range.s.c > C) range.s.c = C
      if (range.e.r < R) range.e.r = R
      if (range.e.c < C) range.e.c = C

      const cell = {
        v: data[R][C],
        s: {
          font: {
            name: '微软雅黑',
            sz: 9,
            bold: false,
            // color: {
            //   auto: 1
            // },
            color: {
              rgb: ''
            }
          },
          border: {
            color: { auto: 1 }
          },
          alignment: {
            // / 自动换行
            wrapText: 1,
            // 居中
            horizontal: 'center',
            vertical: 'center',
            indent: 0
          }
        }
      }
      // 这里生成cell的时候，使用上面定义的默认样式
      if (cell.v == null) continue
      const cell_ref = XLSX.utils.encode_cell({
        c: C,
        r: R
      })

      if (typeof cell.v === 'number') cell.t = 'n'
      else if (typeof cell.v === 'boolean') cell.t = 'b'
      else if (cell.v instanceof Date) {
        cell.t = 'n'
        cell.z = XLSX.SSF._table[14]
        cell.v = datenum(cell.v)
      } else cell.t = 's'

      ws[cell_ref] = cell
    }
  }
  if (range.s.c < 10000000) ws['!ref'] = XLSX.utils.encode_range(range)
  return ws
}

// 新建excel文件(一个excel文件里面有很多工作表)
function Workbook() {
  if (!(this instanceof Workbook)) return new Workbook()
  this.SheetNames = []
  this.Sheets = {}
}

// 字符串转ArrayBuffer
function s2ab(s) {
  const buf = new ArrayBuffer(s.length)
  const view = new Uint8Array(buf)
  for (let i = 0; i !== s.length; ++i) view[i] = s.charCodeAt(i) & 0xff
  return buf
}

/**
 * 通用的打开下载对话框方法，没有测试过具体兼容性
 * @param url 下载地址，也可以是一个blob对象，必选
 * @param saveName 保存文件名，可选
 */
export function openDownloadDialog(url, saveName) {
  if (typeof url === 'object' && url instanceof Blob) {
    url = URL.createObjectURL(url) // 创建blob地址
  }
  const aLink = document.createElement('a')
  aLink.href = url
  aLink.download = saveName || '' // HTML5新增的属性，指定保存文件名，可以不要后缀，注意，file:///模式下不会生效
  let event
  if (window.MouseEvent) event = new MouseEvent('click')
  else {
    event = document.createEvent('MouseEvents')
    event.initMouseEvent(
      'click',
      true,
      false,
      window,
      0,
      0,
      0,
      0,
      0,
      false,
      false,
      false,
      false,
      0,
      null
    )
  }
  aLink.dispatchEvent(event)
}

// 生成excel文件总入口函数
export function export_json_to_excel({
  title,
  multiHeader = [],
  header,
  data,
  filename,
  merges = [],
  autoWidth = true,
  bookType = 'xlsx'
} = {}) {
  const getFormattedTime = () => {
    const today = new Date()
    const y = today.getFullYear()
    // JavaScript months are 0-based.
    const m = today.getMonth() + 1
    const d = today.getDate()
    const h = today.getHours()
    const mi = today.getMinutes()
    const s = today.getSeconds()
    return y + '-' + m + '-' + d + '-' + h + '-' + mi + '-' + s
  }
  filename = filename || getFormattedTime()
  data = [...data]
  data.unshift(header)
  if (merges.length > 0) {
    data.unshift(title)
  }
  for (let i = multiHeader.length - 1; i > -1; i--) {
    data.unshift(multiHeader[i])
  }

  const ws_name = 'SheetJS'
  const wb = new Workbook()
  const ws = sheet_from_array_of_arrays(data)

  if (merges.length > 0) {
    if (!ws['!merges']) ws['!merges'] = []
    merges.forEach(item => {
      ws['!merges'].push(XLSX.utils.decode_range(item))
    })
  }

  if (autoWidth) {
    // 设置worksheet每列的最大宽度
    const colWidth = data.map(row =>
      row.map(val => {
        if (val == null) {
          return {
            wch: 10
          }
        } else {
          let maxWch = 0
          let wch = 0
          try {
            const arrForVal = val.toString().split('\n')
            arrForVal.forEach(val2 => {
              if (val2 == null) {
                wch = 10
              } else if (val2.toString().charCodeAt(0) > 255) {
                wch = val2.toString().length * 2
              } else {
                wch = val2.toString().length
              }
              if (wch > maxWch) {
                maxWch = wch
              }
            })
            return {
              wch: maxWch
            }
          } catch {
            console.log(val)
          }
        }
      })
    )
    // 以第一行基准，第二行及其以后的对比第一行的每列宽度，选择最宽的
    const result = colWidth[0]
    for (let i = 1; i < colWidth.length; i++) {
      for (let j = 0; j < colWidth[i].length; j++) {
        if (result[j]['wch'] < colWidth[i][j]['wch']) {
          result[j]['wch'] = colWidth[i][j]['wch']
        }
      }
    }
    ws['!cols'] = result
  } else {
    // 设置worksheet每列的固定宽度
    if (data[0]) {
      ws['!cols'] = data[0].map(() => {
        return {
          wch: 15
        }
      })
    }
  }

  // add worksheet to workbook
  wb.SheetNames.push(ws_name)
  wb.Sheets[ws_name] = ws
  const dataInfo = wb.Sheets[wb.SheetNames[0]]

  // 单元格外侧框线
  const borderAll = {
    top: {
      style: 'thin',
      color: {
        rgb: 'e8eaec'
      }
    },
    bottom: {
      style: 'thin',
      color: {
        rgb: 'e8eaec'
      }
    },
    left: {
      style: 'thin',
      color: {
        rgb: 'e8eaec'
      }
    },
    right: {
      style: 'thin',
      color: {
        rgb: 'e8eaec'
      }
    },
    color: { auto: 1 }
  }

  // 给所有单元格加上边框
  for (const i in dataInfo) {
    if (merges.length > 0) {
      if (i === '!ref' || i === '!merges' || i === '!cols' || i === 'A1') {
      } else {
        dataInfo[i + ''].s.border = borderAll
      }
    } else {
      if (i === '!ref' || i === '!merges' || i === '!cols') {
      } else {
        dataInfo[i + ''].s.border = borderAll
      }
    }
  }

  // 去掉标题边框
  if (merges.length > 0) {
    const arr = [
      'A1',
      'B1',
      'C1',
      'D1',
      'E1',
      'F1',
      'G1',
      'H1',
      'I1',
      'J1',
      'K1',
      'L1',
      'M1',
      'N1',
      'O1',
      'P1',
      'Q1',
      'R1',
      'S1',
      'T1',
      'U1',
      'V1',
      'W1',
      'X1',
      'Y1',
      'Z1'
    ]
    arr.some(function(v) {
      const a = merges[0].split(':')
      if (v === a[1]) {
        dataInfo[v].s = {}
        return true
      } else {
        dataInfo[v].s = {}
      }
    })
  }

  if (merges.length > 0) {
    // 设置主标题样式
    dataInfo['A1'].s = {
      font: {
        name: '微软雅黑',
        sz: 9,
        color: { rgb: 'ff0000' },
        bold: true,
        italic: false,
        underline: false
      },
      alignment: {
        horizontal: 'center',
        vertical: 'center'
      }
      // fill: {
      //   fgColor: {rgb: "008000"},
      // },
    }
  }

  const wbout = XLSX.write(wb, {
    bookType: bookType,
    bookSST: false,
    type: 'binary'
  })

  const blob = new Blob([s2ab(wbout)], {
    type: 'application/octet-stream'
  })

  // saveAs(
  //   new Blob([s2ab(wbout)], {
  //     type: "application/octet-stream"
  //   }),
  //   `${filename}.${bookType}`
  // );
  // todo
  openDownloadDialog(blob, `${filename}.${bookType}`)
  // FileSaver.saveAs(
  //   new Blob([s2ab(wbout)], {
  //     type: "application/vnd.ms-excel"
  //   }),
  //   filename
  // );
}


const projectExportTableCols = [
    {prop:'name', label:'项目名称'},
    {prop:'createdDate', label:'备案申请日期'},
    {prop:'categories', label:'项目分类'},
    {prop:'area', label:'项目所属地'},
    {prop:'projectNum', label:'项目码'},
    {prop:'chargeStandard', label:'收费标准'},
    {prop:'manager', label:'负责人'},
    {prop:'trainMode', label:'培训模式'},
    {prop:'trainTime', label:'学时数'},
    {prop:'trainNumOfPerson', label:'培训人数'},
    {prop:'trainStartDate', label:'培训开始时间'},
    {prop:'trainEndDate', label:'培训结束时间'},
    {prop:'contractStartDate', label:'合同开始时间'},
    {prop:'contractEndDate', label:'合同结束时间'},
    {prop:'projectAmount', label:'项目应收费用'},
    {prop:'paidAmount', label:'已到账费用'},
    {prop:'unpaidAmount', label:'未到账金额'},
    {prop:'', label:''},
    {prop:'', label:''},
    {prop:'', label:''},
    {prop:'', label:''},


]


export const exportToExcel = (tableCols,tableData) => {
    const title = ["表"];
    // 表头（一维数组）
    const tableHead = []
    // 表头字段
    const propList = []

    // props.tableCol is a object array
    tableCols.forEach(item => {
      if (item.type == "multi") {
        // maxLengthIndex will not be undefined when type == 'multi'
        //
        const maxLengthIndex = item.maxSubItemNum;
        const insertIndex = item.insertIndex;
        // 生成表头和proplist
        for (let i = 1; i <= maxLengthIndex; i++) {
          item?.subTitle?.forEach(subItem => {
            if (subItem.isShow) {
              tableHead.push(
                labelFormatter(
                  insertIndex,
                  subItem.label,
                  i.toString(),
                  item.type
                )
              );
              // 这里的multi e.g = landingAgencies.0.agencyName
              propList.push(
                `${item.prop}.${i - 1}.${subItem.prop}.${subItem.formatType}`
              );
            }
          });
        }
      } else {
        if (item.isShow) {
          tableHead.push(item.label)
          // e.g schoolManageRadio.%
          propList.push(`${item.prop}.${item.formatType}`);
        }
      }
    });

    console.log("tableHead", tableHead, "propList", propList);
    console.log("tableData", tableData);

    // props.tableData is a OBJECT ARRAY
    // todo change
    const data = [];
    tableData.forEach(tableDataItem => {
      const rowList = [];
      propList.forEach(list => {
        const listStr = list.split(".");
        try {
          if (listStr.length > 2) {
            // console.log(listStr);
            // e.g tableData[0].landingAgencies.0.shareRadio, %
            rowList.push(
              // listStr[3] can be "currency" or "%"
              numberFormater(
                tableDataItem[listStr[0]][parseInt(listStr[1])][
                  listStr[2]
                ],
                listStr[3]
              )
            );
          } else {
            rowList.push(
              numberFormater(
                tableDataItem[listStr[0]],
                listStr[1]
              )
            );
          }
        } catch (error) {
          // e.g landingAgencies = {}
          rowList.push(" ");
        }
      });
      data.push(rowList);
    });
    console.log("data", data);

    const merges = [];
    excel.export_json_to_excel({
      title: title,
      header: tableHead,
      data,
      merges,
      filename: getFormattedTime(),
      autoWidth: true,
      bookType: "xlsx"
    });
};
