import axios from 'axios'
import store from '@/store/index.js'
import router from '@/router/index.js'
import { Toast } from 'vant'
// 根据环境不同引入不同api地址
import { baseApi } from '@/config'
import { getStorage } from '@/utils/index.js'
import { paramsSignMd5 } from '@/utils/sign.js'

// create an axios instance
const request = axios.create({
  baseURL: baseApi, // url = base api url + request url
  timeout: 5000, // request timeout
  headers: { 'Content-Type': 'application/json; charset=UTF-8' },
})

// request拦截器 request interceptor
request.interceptors.request.use(
  config => {
    // 不传递默认开启loading
    if (!config.hideloading) {
      Toast.loading({
        forbidClick: true
      })
    }
    if (config.auth === true) {
      const token = getStorage('token')
      if (!token) {
        return router.replace({ path: '/login' })
      }

      config.headers['Token'] = token || ''
    }
    const now = parseInt(new Date().getTime() / 1000)
    config.headers['Auth-Date'] =  now
    // sign 
    if (config.method == 'post') {
      const data = {"method":"POST","path":"/v1"+config.url, ...config.data}
      const signStr = paramsSignMd5(data, now)
      // const signStr = paramsSignMd5(data, now, "sha256")
      config.headers['Auth'] = "sign " + signStr
    } else {
      const data = {"method":"GET","path":"/v1"+config.url, ...config.params}
      const signStr = paramsSignMd5(data, now)
      // const signStr = paramsSignMd5(data, now, "sha256")
      config.headers['Auth'] = "sign " + signStr
    }
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
    if (response.status === 200) {
      const result = response.data
      if (result.code === 0) {
        return Promise.resolve(result.data)
      } else if (result.code == 20111 || result.code == 20120) {
        Toast('令牌已过期，请重新登录')
        store.dispatch('logout')
        return Promise.reject(result.msg)
      }
      Toast(result.msg)
      return Promise.reject(result.msg)
    } else {
      Toast('网络开小差了')
      console.log('response err', response)
      return Promise.reject(response.statusText)
    }
  },
  error => {
    Toast.clear()
    console.log('err', error) // for debug
    return Promise.reject(error)
  }
)

export default request
