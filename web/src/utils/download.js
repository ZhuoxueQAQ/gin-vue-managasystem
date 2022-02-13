export const saveAs = (data, fileName) => {
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
