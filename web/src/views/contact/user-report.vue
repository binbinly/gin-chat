<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="用户投诉" left-arrow @click-left="onClickLeft">
      <template #right>
        <van-button type="primary" size="small" @click="onReport">提交</van-button>
      </template>
    </van-nav-bar>
    <van-field readonly clickable :value="form.category" label="分类" @click="picker = true" placeholder="请选择分类" :rules="[{ required: true}]" />
    <van-field v-model="form.content" rows="3" autosize label="意见箱" type="textarea" maxlength="500" placeholder="请填写投诉内容" show-word-limit
               :rules="[{ required: true}]" />
    <van-popup v-model="picker" round position="bottom">
      <van-picker show-toolbar :columns="categoryList" @cancel="picker = false" @confirm="onConfirm" />
    </van-popup>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { userReport } from '@/api/friend.js'
import { Toast } from 'vant';
export default {
  mixins: [auth],
  data() {
    return {
      picker: false,
      categoryList: ["分类一", "分类二", "分类三", "分类四", "分类五"],
      form: {
        user_id: 0,
        type: 1,
        category: "",
        content: ""
      }
    }
  },
  activated() {
    if (!this.$route.query) {
      return this.backToast()
    }
    this.form.user_id = parseInt(this.$route.query.user_id)
    this.form.type = parseInt(this.$route.query.type)
  },
  methods: {
    onConfirm(value) {
      this.form.category = value
      this.picker = false
    },
    onReport() {
      if (!this.form.category) {
        return Toast('请选择分类')
      }
      if (!this.form.content) {
        return Toast('请填写投诉内容')
      }
      // 请求服务器
      userReport(this.form).then(res => {
        Toast.success('投诉成功')
        this.$router.back()
      })
    }
  }
}
</script>

<style>
</style>
