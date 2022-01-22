import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}
export const formatDate = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd')
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
  return (radio * 100).toString().concat('%')
}
export const percentageToNumber = (str) => {
  return parseFloat(str.split('%').join(''))
}

export const formatTableVal = (value, format) => {
  switch (format) {
    case 'date':
      return myDateFormat(value)
    case 'amount':
      return numberToThousands(value)
    case 'radio':
      return numberToPercentage(value)
    default:
      return value
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
