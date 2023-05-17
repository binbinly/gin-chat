<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar title="群聊列表" fixed placeholder left-arrow @click-left="onClickLeft" />

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list v-model="loading" :finished="finished" finished-text="没有更多了" @load="onLoad">
        <van-cell v-for="(item,index) in list" :title="item.name" center is-link @click="handle(item)">
          <template #icon>
            <van-image class="pr-1" round width="35" height="35" :src="item.avatar|formatAvatar" />
          </template>
        </van-cell>
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { groupList } from '@/api/group.js'
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
  methods: {
    onLoad() {
      groupList(this.page).then(res => {
        if (this.refreshing) {
          this.list = [];
          this.refreshing = false;
        }
        this.list = [...this.list, ...res]
        this.loading = false;
        if (res.length < 10) {
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
    handle(item) {
      this.$router.push({
        path: '/chat', query: {
          id: item.id,
          type: 2
        }
      })
    }
  }
}
</script>

<style>
</style>
