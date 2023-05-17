<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar title="设置" left-arrow @click-left="onClickLeft" />
    <!-- 退出登录 -->
    <van-divider />
    <van-button type="danger" block @click="logout">退出登录</van-button>

  </div>
</template>

<script>
import auth from '@/mixin/auth.js'
import { userLogout } from '@/api/user.js'
import { Dialog, Toast } from 'vant'
export default {
  mixins: [auth],
  data() {
    return {

    }
  },
  methods: {
    logout() {
      Dialog.confirm({
        message: '确定退出吗？',
      })
        .then(() => {
          userLogout().then(res => {
            Toast.success('退出登录成功')
            this.$store.dispatch('logout')
          })
        })
    }
  }
}
</script>

<style>
</style>
