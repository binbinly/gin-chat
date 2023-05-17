<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="标签用户" fixed placeholder left-arrow @click-left="onClickLeft">
    </van-nav-bar>

    <van-cell v-for="(item,index) in list" :title="item.name" is-link center @click="openUser(item.id)">
      <template #icon>
        <van-image class="pr-1" round width="35" height="35" :src="item.avatar|formatAvatar" />
      </template>
    </van-cell>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { tagUserList } from '@/api/tag.js'
export default {
  mixins: [auth],
  data() {
    return {
      list: [],
      id: 0
    }
  },
  activated() {
    this.id = parseInt(this.$route.query.id)
    this.id > 0 && this.getData()
  },
  methods: {
    getData() {
      tagUserList(this.id).then(res => {
        this.list = res
      })
    },
    openUser(id) {
      this.$router.push({ path: '/user_base', query: { id } })
    }
  }
}
</script>

<style>
</style>
