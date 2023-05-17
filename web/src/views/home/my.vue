<template>
  <div class="page">
    <!-- 导航 -->
    <van-nav-bar fixed placeholder @click-right="onClickRight">
      <template #right>
        <van-icon name="photograph" size="22" />
      </template>
    </van-nav-bar>

    <van-cell is-link center to="/userinfo">
      <template #icon>
        <van-image class="pr-1" round width="45" height="45" :src="user.avatar|formatAvatar" />
      </template>
      <template #title>
        <div class="flex flex-column">
          <span class="text-dark font-lg font-weight-bold">{{nickname}}</span>
          <span class="text-light-muted">仿微信号：{{user.username}}</span>
        </div>
      </template>
      <template #default>
        <span class="iconfont font-lg text-light-muted">&#xe614;</span>
      </template>
    </van-cell>
    <van-divider />
    <van-cell title="支付" is-link>
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe66c;</span>
      </template>
    </van-cell>
    <van-divider />
    <van-cell title="收藏" is-link to="/fava">
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xea2d;</span>
      </template>
    </van-cell>
    <van-cell title="动态" is-link to="/moments">
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe608;</span>
      </template>
    </van-cell>
    <van-cell title="表情" is-link to="/emoticon_cat">
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe605;</span>
      </template>
    </van-cell>
    <van-divider />
    <van-cell title="设置" is-link to="/setting">
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe612;</span>
      </template>
    </van-cell>

  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { mapState } from 'vuex'
export default {
  mixins: [auth],
  data() {
    return {

    }
  },
  computed: {
    ...mapState({
      user: state => state.user.user
    }),
    nickname() {
      return this.user.nickname ? this.user.nickname : this.user.username
    }
  },
  methods: {
    onClickRight() {
      this.$router.push({ path: '/moments' })
    }
  }
}
</script>

<style>
</style>
