import { getStorage } from '@/utils/index.js'
import { Toast } from 'vant'
export default {
  created() {
    let token = getStorage('token')
    if (!token) {
      return this.$router.push({ path: '/login' })
    }
  },
  methods: {
    onClickLeft() {
      this.$router.back()
    },
    push(path, params) {
      this.$router.push({ path, params })
    },
    draw() {
      Toast('待开发')
    },
    // 返回并提示
    backToast(msg = '非法参数') {
      this.toast(msg)
      setTimeout(() => {
        this.$router.back()
      }, 500)
    },
    toast(msg = '非法参数') {
      Toast({ title: msg, position: 'bottom' })
    }
  }
}
