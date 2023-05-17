<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="设置朋友圈动态权限" left-arrow @click-left="onClickLeft" />

    <van-cell center title="不让他看我">
      <template #right-icon>
        <van-switch v-model="form.lookme" size="24" active-color="#08c060" @change="onChangeMe" />
      </template>
    </van-cell>
    <van-cell center title="不看他">
      <template #right-icon>
        <van-switch v-model="form.lookhim" size="24" active-color="#08c060" @change="onChangeHim" />
      </template>
    </van-cell>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { friendMomentAuth } from '@/api/friend.js'
import { Toast } from 'vant';
export default {
  mixins: [auth],
  data() {
    return {
      id: 0,
      form: {
        lookme: true,
        lookhim: true
      }
    }
  },
  mounted() {
    this.id = parseInt(this.$route.query.id)
    this.form.lookhim = this.$route.query.lookhim == 1 ? true : false
    this.form.lookme = this.$route.query.lookme == 1 ? true : false
  },
  methods: {
    onChangeMe(value) {
      this.form.lookme = value
      this.submit()
    },
    onChangeHim(value) {
      this.form.lookhim = value
      this.submit()
    },
    submit() {
      friendMomentAuth({
        user_id: this.id,
        lookme: this.form.lookme ? 1 : 0,
        lookhim: this.form.lookhim ? 1 : 0,
      }).then(res => {
        Toast.success('修改成功')
      })
    }
  }
}
</script>

<style>
</style>
