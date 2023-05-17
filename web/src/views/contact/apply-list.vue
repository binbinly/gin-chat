<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="好友申请列表" fixed placeholder left-arrow @click-left="onClickLeft" />

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list v-model="loading" :finished="finished" finished-text="没有更多了" @load="onLoad" :immediate-check="false">
        <van-cell v-for="(item,key) in list" :title="item.user.name" center>
          <template #icon>
            <van-image class="pr-1" round width="35" height="35" :src="item.user.avatar|formatAvatar" />
          </template>
          <template #default>
            <van-button v-if="item.status === 1" type="primary" size="small" @click="handle(item.user)">同意</van-button>
            <span v-else class="text-muted font-sm">{{ item | formatTitle }}</span>
          </template>
        </van-cell>
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { applyList } from '@/api/apply.js'
import $const from '@/const/index.js'
export default {
  mixins: [auth],
  data() {
    return {
      loading: false,
      finished: false,
      refreshing: false,
      page: 1,
      list: []
    }
  },
  filters: {
    formatTitle(value) {
      let obj = {
        3: "已通过",
        2: "已拒绝",
        4: "已忽略"
      }
      return obj[value.status];
    }
  },
  activated() {
    this.list = []
    this.onRefresh()
  },
  methods: {
    handle(user) {
      this.$router.push({ path: '/add_friend', query: { id: user.id, act: 'handle', nickname: user.name } })
    },
    onLoad() {
      applyList(this.page).then(res => {
        if (this.refreshing) {
          this.list = [];
          this.refreshing = false;
        }
        this.list = [...this.list, ...res]
        this.loading = false;
        if (res.length < $const.PAGE_SIZE) {
          this.finished = true;
        }
        this.page++
      })
    },
    onRefresh() {
      this.page = 1
      // 清空列表数据
      this.finished = false;
      // 重新加载数据
      // 将 loading 设置为 true，表示处于加载状态
      this.loading = true;
      this.onLoad();
    },
  }
}
</script>

<style>
</style>
