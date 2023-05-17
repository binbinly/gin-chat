<template>
  <div>
    <!-- 导航栏 -->
    <form>
      <van-search v-model="keyword" show-action fixed placeholder="搜索用户" @search="onSearch" @cancel="onCancel" />
    </form>

    <van-cell v-for="(item,index) in list" center @click="openUserBase(item.id)" :value="item.nickname" is-link>
      <template #icon>
        <van-image class="pr-1" round width="35" height="35" :src="item.avatar|formatAvatar" />
      </template>
      <template #title>
        <span class="custom-title pr-1">{{item.username}}</span>
        <van-tag type="danger">{{item.phone}}</van-tag>
      </template>
    </van-cell>
  </div>
</template>

<script>
import { searchUser } from '@/api/common.js'
export default {
  data() {
    return {
      keyword: "",
      list: []
    }
  },
  methods: {
    onSearch(val) {
      searchUser({
        keyword: this.keyword
      }).then(res => {
        this.list = res
      })
    },
    onCancel() {
      this.$router.back()
    },
    // 打开用户资料
    openUserBase(user_id) {
      this.$router.push({ path: '/user_base', query: { id: user_id } })
    }
  }
}
</script>

<style>
</style>
