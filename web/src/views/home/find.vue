<template>
  <div class="page">
    <!-- 导航栏 -->
    <nav-bar title="发现"></nav-bar>

    <van-cell title="朋友圈" is-link center to="/moments">
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe667;</span>
      </template>
      <template #default>
        <div class="position-relative" v-if="notice.num > 0 && notice.user_id">
          <van-image round width="30" height="30" :src="notice.avatar|formatAvatar" />
          <span class="rounded-circle bg-danger position-absolute" style="width: 10px;height:10px;top: 0;right: 0;"></span>
        </div>
      </template>
    </van-cell>
    <van-divider />
    <van-cell title="扫一扫" is-link>
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe86d;</span>
      </template>
    </van-cell>
    <van-cell title="摇一摇" is-link>
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe63d;</span>
      </template>
    </van-cell>
    <van-divider />
    <van-cell title="看一看" is-link>
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe610;</span>
      </template>
    </van-cell>
    <van-cell title="搜一搜" is-link>
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe611;</span>
      </template>
    </van-cell>
    <van-divider />
    <van-cell title="购物" is-link>
      <template #icon>
        <span slot="icon" class="iconfont font-lg pr-1">&#xe658;</span>
      </template>
    </van-cell>

  </div>
</template>

<script>
import NavBar from "@/components/nav-bar.vue"
import auth from '@/mixin/auth.js';
import { mapState } from 'vuex'
export default {
  mixins: [auth],
  components: {
    NavBar
  },
  data() {
    return {
      notice: {}
    }
  },
  computed: {
    ...mapState({
      chat: state => state.user.chat
    })
  },
  mounted() {
    this.notice = this.chat.getNotice()
  },
  methods: {
  }
}
</script>

<style>
</style>
