<template>
  <div class="page">
    <!-- 导航栏 -->
    <van-nav-bar :left-text="titleText" fixed placeholder left-arrow @click-left="onClickLeft" />

    <van-form @submit="onSubmit">
      <van-cell-group title="备注名">
        <van-field v-model="nickname" placeholder="请填写备注名" :rules="[{ required: true }]" />
      </van-cell-group>
      <van-cell-group title="朋友圈权限">
        <van-cell center title="不让他看我">
          <template #right-icon>
            <van-switch v-model="look_me" size="24" active-color="#08c060" />
          </template>
        </van-cell>
        <van-cell center title="不看他">
          <template #right-icon>
            <van-switch v-model="look_him" size="24" active-color="#08c060" />
          </template>
        </van-cell>
      </van-cell-group>
      <div style="margin: 16px;">
        <van-button block type="primary" native-type="submit">提交</van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { applyFriend, applyHandle } from '@/api/apply.js'
import { Toast } from 'vant';
export default {
  mixins: [auth],
  data() {
    return {
      look_me: true,
      look_him: true,
      nickname: "",
      friend_id: 0,
      act: 'apply'
    }
  },
  computed: {
    titleText() {
      return this.act == 'apply' ? '添加好友' : '申请处理'
    }
  },
  activated() {
    this.friend_id = parseInt(this.$route.query.id)
    if (!this.friend_id) {
      Toast('参数非法')
      return this.$router.back()
    }
    this.nickname = this.$route.query.nickname
    this.act = this.$route.query.act || 'apply'
  },
  methods: {
    onSubmit() {
      // 添加好友
      if (this.act === 'apply') {
        applyFriend({
          look_me: this.look_me ? 1 : 0,
          look_him: this.look_him ? 1 : 0,
          nickname: this.nickname,
          friend_id: this.friend_id,
        }).then(() => {
          Toast.success('申请成功')
          this.$router.back()
        })
      } else {
        // 处理好友申请
        applyHandle({
          look_me: this.look_me ? 1 : 0,
          look_him: this.look_him ? 1 : 0,
          nickname: this.nickname,
          friend_id: this.friend_id
        }, this.id).then(() => {
          Toast.success('处理成功')
          this.$store.dispatch('contactList')
          this.$store.dispatch('getApply')
          this.$router.back()
        })
      }
    }
  }
}
</script>

<style>
</style>
