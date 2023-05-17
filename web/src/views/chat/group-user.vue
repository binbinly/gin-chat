<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="群聊用户" left-arrow @click-left="onClickLeft" />

    <!-- 搜索框 -->
    <van-search v-model="keyword" placeholder="搜索" />
    <!-- 联系人列表 -->
    <div class="px-1 py-1 bg-white">
      <span class="text-muted">{{keyword ? '搜索结果' :'用户'}}</span>
    </div>

    <van-cell v-for="(item,index) in allList" :title="item.name" center @click="selectItem(item)">
      <template #icon>
        <van-image class="pr-1" round width="35" height="35" :src="item.avatar|formatAvatar" />
      </template>
    </van-cell>

    <div v-if="keyword !== '' && searchList.length === 0" class="flex align-center justify-center" style="height: 100rpx;">
      <span class="text-light-muted">暂无搜索结果</span>
    </div>

  </div>
</template>

<script>
import { groupUser, groupKick } from '@/api/group.js'
import auth from '@/mixin/auth.js';
import { Dialog, Toast } from 'vant';
import event from '@/utils/event.js';
export default {
  mixins: [auth],
  data() {
    return {
      keyword: "",
      top: 0,
      list: [],
      group_id: 0
    }
  },
  activated() {
    if (this.$route.query.id) {
      this.group_id = parseInt(this.$route.query.id)
      groupUser(this.group_id).then(res => {
        this.list = res
      })
    }
  },
  computed: {
    // 最终列表
    allList() {
      return this.keyword === '' ? this.list : this.searchList
    },
    // 搜索结果列表
    searchList() {
      if (this.keyword === '') {
        return []
      }
      return this.list.filter(item => {
        return item.name.indexOf(this.keyword) !== -1
      })
    }
  },
  methods: {
    // 选中/取消选中 | 发送
    selectItem(item) {
      Dialog.confirm({
        message: '是否要踢出该成员？',
      })
        .then(() => {
          groupKick({
            id: this.group_id,
            user_id: item.id
          }).then(() => {
            Toast.success('踢出成功')
            event.$emit('refreshGroupInfo')
            this.$router.back()
          })
        })
    },
  }
}
</script>

<style>
</style>
