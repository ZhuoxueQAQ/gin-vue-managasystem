import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}
export const formatDate = (time, pattern) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, pattern)
  } else {
    return ''
  }
}

export const myDateFormat = (arg1) => {
  // 这里必须用参数arg去new一个Date出来，不然会提示.xxx is not a function！！
  if (arg1 != null && arg1 !== '' && arg1 !== undefined) {
    return new Date(arg1).toLocaleDateString().replace(/(\/)/g, '-')
  } else {
    return ''
  }
}

export const numberToThousands = (num) => {
  try {
    const parts = String(num).split('.')
    parts[0] = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ',')
    let res = parts.join('.')
    // Two decimal places
    if (parts[1] == null || parts[1] === undefined) {
      res += '.00'
    } else if (parts[1] !== undefined && parts[1].length === 1) {
      res += '0'
    }
    return '￥' + res
  } catch (error) {
    console.log(error, typeof num)
    return 'error'
  }
}

export const numberToPercentage = (radio) => {
  return radio.toString().concat('%')
}
export const percentageToNumber = (str) => {
  return parseFloat(str.split('%').join(''))
}

export const formatFileSize = (size) => {
  if (size < 0) {
    return 'invalid size!'
  }
  // 1024^x = y
  const pow = Math.floor(Math.log(size) / Math.log(1024))
  switch (pow) {
    case 0:
      return size.toFixed(2) + 'B'
    case 1:
      return (size / 1024).toFixed(2) + 'KB'
    case 2:
      return (size / Math.pow(1024, 2)).toFixed(2) + 'MB'
    default:
      return (size / Math.pow(1024, 3)).toFixed(2) + 'GB'
  }
}

export const formatTableVal = (value, format) => {
  try {
    switch (format) {
      case 'date': {
        const res = formatDate(value, 'yyyy-MM-dd')
        if (res === '1-01-01') {
          return ''
        }
        return res
      }
      case 'dateTime':
        return formatDate(value, 'yyyy-MM-dd hh:mm:ss')
      case 'amount':
        return numberToThousands(value)
      case 'radio':
        return numberToPercentage(value)
      case 'fileSize':
        return formatFileSize(value)
      default:
        return value
    }
  } catch (e) {
    console.log(value, format)
    return ''
  }
}

export const filterDict = (value, options) => {
  const rowLabel = options && options.filter(item => item.value === value)
  return rowLabel && rowLabel[0] && rowLabel[0].label
}

export const getDictFunc = async(type) => {
  const dicts = await getDict(type)
  return dicts
}
