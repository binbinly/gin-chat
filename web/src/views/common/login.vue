<template>
  <div class="page">
    <div v-if="!show" class="position-fixed top-0 bottom-0 left-0 right-0 bg-light flex align-center justify-center">
      <span class="text-muted font">正在加载...</span>
    </div>

    <div v-else>
      <div class="flex align-center justify-center" style="height: 175px;">
        <span style="font-size: 26px;">YOU-LOGO</span>
      </div>
      <div style="margin: 16px;">
        <van-form @submit="onSubmit">
          <template v-if="type === 'login'">
            <van-field v-model="username" name="username" placeholder="请输入用户名" :rules="[{ required: true }]" />
            <van-divider />
            <van-field v-model="password" type="password" name="password" placeholder="请输入密码" :rules="[{ required: true, validator: pwdValid }]" />
          </template>
          <template v-else>
            <van-field v-model="username" name="username" placeholder="请输入用户名" :rules="[{ required: true }]" />
            <van-divider />
            <van-field v-model="phone" name="phone" placeholder="请输入手机号" :rules="[{ required: true }]" />
            <van-divider />
            <van-field v-model="password" type="password" name="password" placeholder="请输入密码" :rules="[{ required: true, validator: pwdValid}]" />
            <van-divider />
            <van-field v-model="repassword" type="password" name="confirm_password" placeholder="请确认密码"
                       :rules="[{ required: true, validator: confirmPwdValid}]" />
          </template>
          <van-divider />
          <van-button class="text-white main-bg-color" hover-class="main-bg-hover-color" block native-type="submit">
            {{type === 'login' ? '登 录' : '注 册'}}
          </van-button>
        </van-form>
      </div>
      <div class="flex align-center justify-center">
        <span class="text-light-muted font p-2" @click="changeType">{{type === 'login' ?  '注册账号' : '马上登录'}}</span>
        <span class="text-light-muted font">|</span>
        <span class="text-light-muted font p-2" @click="resetPwd">忘记密码</span>
      </div>
    </div>

  </div>
</template>

<script>
import { getStorage } from '@/utils/index.js';
import { login, reg } from '@/api/common.js';
import { Toast } from 'vant';
export default {
  data() {
    return {
      type: "login",
      show: false,
      username: "",
      phone: "",
      password: "",
      repassword: ""
    }
  },
  //组件生命周期，在实例创建完成后被立即调用
  created() {
    let token = getStorage('token')
    if (!token) {
      // 用户未登录
      return this.show = true
    }
    this.$router.replace({ path: '/home' })
  },
  methods: {
    resetPwd() {
      Toast('待开发')
    },
    // 初始化表单
    initForm() {
      this.username = ''
      this.password = ''
      this.repassword = ''
      this.phone = ''
    },
    pwdValid(val) {
      return val.length >= 6 && val.length <= 20
    },
    confirmPwdValid(val) {
      return val === this.password
    },
    changeType() {
      this.type = this.type === 'login' ? 'reg' : 'login'
      this.initForm()
    },
    onSubmit() {
      if (this.type === 'login') {
        login({
          username: this.username,
          password: this.password
        }).then(res => {
          this.$store.dispatch('login', res)
          Toast.success('登录成功')
          this.$router.replace({ path: '/home' })
        })
      } else {
        reg({
          username: this.username,
          phone: this.phone,
          password: this.password,
          confirm_password: this.repassword
        }).then(() => {
          this.changeType()
          Toast('注册成功，去登录')
        })
      }
    }
  }
}
</script>

<style>
</style>
