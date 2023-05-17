<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="聊天记录" fixed placeholder left-arrow @click-left="onClickLeft" />
    <!-- 搜索框 -->
    <van-search v-model="keyword" placeholder="搜索" />

    <!-- 联系人列表 -->
    <div class="px-1 py-1 bg-white">
      <span class="text-muted">{{keyword ? '搜索结果' :'历史记录'}}</span>
    </div>

    <div v-for="(item,index) in allList" :key="index" :id="'chatItem_'+index">
      <chat-item :item="item" :index="index" ref="chatItem" :pretime=" index > 0 ? list[index-1].t : 0" :show_name="true">
      </chat-item>
    </div>

    <div v-if="keyword !== '' && searchList.length === 0" class="flex align-center justify-center" style="height: 50px;">
      <span class="text-light-muted">暂无搜索结果</span>
    </div>
  </div>
</template>

<script>
import ChatItem from '@/components/chat-item.vue';
import auth from '@/mixin/auth.js';
import { mapState } from 'vuex';
export default {
  mixins: [auth],
  components: {
    ChatItem
  },
  data() {
    return {
      keyword: "",
      top: 0,
      list: []
    }
  },
  activated() {
    this.list = this.chat.getChatDetail()
  },
  computed: {
    ...mapState({
      user: state => state.user.user,
      chat: state => state.user.chat
    }),
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
        return item.content.indexOf(this.keyword) !== -1
      })
    },
  },
  methods: {
  }
}
</script>

<style>
</style>
