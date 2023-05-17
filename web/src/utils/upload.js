import axios from 'axios'
import { Toast } from 'vant'
// 根据环境不同引入不同api地址
import { baseApi } from '@/config'

// create an axios instance
const request = axios.create({
  baseURL: uploadUrl, // url = base api url + request url
  url: '',
  timeout: 5000, // request timeout
  headers: { 'Content-Type': 'application/json;charset=UTF-8' },
  data: {
    output: 'json'
  }
})

// request拦截器 request interceptor
request.interceptors.request.use(
  config => {
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)
// respone拦截器
request.interceptors.response.use(
  response => {
    Toast.clear()
    if (response.status !== 200) {
      Toast('上传失败')
      console.log('response err', response)

      return Promise.reject(response.statusText)
    }
    return Promise.resolve(response.data)
  },
  error => {
    Toast.clear()
    console.log('err', error) // for debug
    return Promise.reject(error)
  }
)

export default request
