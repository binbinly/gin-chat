<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="选择" left-arrow @click-left="onClickLeft">
      <template #right>
        <van-button type="primary" size="small" @click="handleNavBarBtn">{{muliSelect ? '发送 ('+selectCount+')' : '多选'}}</van-button>
      </template>
    </van-nav-bar>
    <!-- 搜索框 -->
    <van-search v-model="keyword" placeholder="搜索" input-align="center" />

    <!-- 好友列表 -->
    <van-cell :title="item.name" v-for="(item,index) in allList" center @click="selectItem(item)">
      <template #icon>
        <van-image class="pr-1" round width="35" height="35" :src="item.avatar|formatAvatar" />
      </template>
      <template #right-icon v-if="muliSelect">
        <van-checkbox v-model="item.checked" checked-color="#08c060" ref="checkboxes" />
      </template>
    </van-cell>
    <div v-if="keyword !== '' && searchList.length === 0" class="flex align-center justify-center" style="height: 50px;">
      <span class="text-light-muted">暂无搜索结果</span>
    </div>

    <!-- 弹出层 - 名片 -->
    <van-dialog v-model="show" title="发送给：" show-cancel-button @confirm="onConfirm">
      <div v-if="selectCount > 0">
        <div class="ml-1 flex flex-column" v-for="(item,index) in selectList">
          <van-image :src="item.avatar|formatAvatar" round width="35" height="35" />
          <span class="text-muted">{{item.name}}</span>
        </div>
      </div>
      <div class="flex flex-column" v-else>
        <van-image class="m-1" :src="sendItem.avatar|formatAvatar" width="35" height="35" />
        <span class="text-muted" style="margin-top:16px;">{{sendItem.name}}</span>
      </div>
      <div class="my-1 bg-light rounded p-1">
        <span class="text-light-muted">{{message}}</span>
      </div>
      <van-field v-model="content" class="bg-light" name="content" placeholder="给好友留言" />
    </van-dialog>
  </div>
</template>

<script>
import { Toast } from 'vant'
import { mapState } from 'vuex'
import auth from '@/mixin/auth.js';
import $const from '@/const/index.js';
export default {
  mixins: [auth],
  data() {
    return {
      show: false,
      keyword: "",
      muliSelect: false,
      list: [],
      detail: {},
      sendItem: {},
      content: ""
    }
  },
  activated() {
    this.list = this.chatList.map(item => {
      return {
        ...item,
        checked: false
      }
    })
    if (this.$route.query) {
      this.detail = this.$route.query
      if (this.detail.type) {
        this.detail.type = $const.TYPE_TRANS_LIST[parseInt(this.detail.type)]
      }
    }
    console.log('detail', this.detail)
  },
  computed: {
    ...mapState({
      chatList: state => state.user.chatList,
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
        return item.name.indexOf(this.keyword) !== -1
      })
    },
    // 选中列表
    selectList() {
      return this.list.filter(item => item.checked)
    },
    // 选中数量
    selectCount() {
      return this.selectList.length
    },
    message() {
      let obj = {
        image: "[图片]",
        video: "[视频]",
        audio: "[语音]",
        card: "[名片]",
        emoticon: "[表情]"
      }
      return this.detail.type === 'text' ? this.detail.content : obj[this.detail.type]
    }
  },
  methods: {
    // 点击导航栏按钮事件（群发）
    handleNavBarBtn() {
      // 切换成多选状态
      if (!this.muliSelect) {
        return this.muliSelect = true
      }
      // 发送
      if (this.selectCount === 0) {
        return Toast('请选择')
      }
      this.show = true
    },
    onConfirm() {
      // 发送
      this.selectList.forEach(item => {
        this.send(item)
        if (this.content) {
          this.send(item, this.content, 'text')
        }
      })
      this.$router.back()
    },
    // 选中/取消选中 | 发送
    selectItem(item) {
      // 选中/取消选中
      if (this.muliSelect) {
        if (!item.checked && this.selectCount === 9) {
          // 选中|限制选中数量
          return Toast('最多选中 9 个')
        }
        return item.checked = !item.checked
      }
      this.show = true
      this.sendItem = item
      item.checked = true
    },
    send(item, data = false, type = false) {
      const message = this.chat.formatSendData({
        to: { id: item.id, name: item.name, avatar: item.avatar },
        chat_type: item.chat_type || $const.CHAT_TYPE_USER,
        content: data || this.detail.content,
        type: type || this.detail.type,
        options: this.detail.options
      })
      this.chat.send(message)
    }
  }
}
</script>

<style>
</style>
