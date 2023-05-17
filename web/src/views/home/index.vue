<template>
  <div>
    <!-- 导航栏 -->
    <nav-bar title="仿微信" :unread="totalunread"></nav-bar>

    <!-- 断线状态 -->
    <div class="bg-danger left-0 right-0 flex align-center justify-between" v-if="!chat || !chat.isOnline">
      <span class="text-white p-1" @click="chat.reconnectConfirm()">当前处于离线状态，点击重新连接</span>
    </div>

    <!-- 置顶列表 -->
    <template v-for="(item,index) in list">
      <msg-item v-if="item.is_top" :item="item" :index="index" @long="long" />
    </template>

    <!-- 非置顶列表 -->
    <template v-for="(item,index) in list">
      <msg-item v-if="!item.is_top" :item="item" :index="index" @long="long" />
    </template>

    <!-- 弹出层 -->
    <free-popup ref="extend" :bodyWidth="200" :bodyHeight="getMenusHeight">
      <div class="flex flex-column text-white p-1" style="width: 90px;" :style="getMenusStyle">
        <div class="flex-1 flex align-center justify-center" style="padding:5px;" v-for="(item,index) in menus" :key="index"
             @click="clickEvent(item.event)">
          <span class="font-sm">{{item.name}}</span>
        </div>
      </div>
    </free-popup>
  </div>
</template>

<script>
import FreePopup from '@/components/free-popup.vue';
import NavBar from "@/components/nav-bar.vue"
import MsgItem from "@/components/msg-item.vue"
import auth from '@/mixin/auth.js';
import { mapState } from 'vuex'
export default {
  mixins: [auth],
  components: {
    FreePopup,
    NavBar,
    MsgItem
  },
  data() {
    return {
      propIndex: -1,
      menus: [{
        name: "设为置顶",
        event: "setTop"
      },
      {
        name: "删除该聊天",
        event: "delChat"
      }
      ],
    }
  },
  mounted() {
    console.log('list', this.list)
  },
  computed: {
    ...mapState({
      list: state => state.user.chatList,
      totalunread: state => state.user.totalunread,
      chat: state => state.user.chat
    }),
    // 动态获取菜单高度
    getMenusHeight() {
      let H = 100
      return this.menus.length * H
    },
    // 获取菜单的样式
    getMenusStyle() {
      return `height: ${this.getMenusHeight}rpx;`
    }
  },
  methods: {
    long({
      x,
      y,
      index
    }) {
      // 初始化 索引
      this.propIndex = index
      // 拿到当前对象
      let item = this.list[index]
      console.log(x, y, index, item)
      // 判断之前是否处于置顶状态
      this.menus[0].name = item.is_top ? '取消置顶' : '设为置顶'
      this.$refs.extend.show(x, y)
    },
    // 分发菜单事件
    clickEvent(event) {
      switch (event) {
        case "setTop": // 置顶/取消置顶会话
          this.setTop()
          break;
        case "delChat": // 删除当前会话
          this.delChat()
          break;
      }
      this.$refs.extend.hide()
    },
    // 删除当前会话
    delChat() {
      let item = this.list[this.propIndex]
      this.chat.removeChatItem(item.id, item.chat_type)
    },
    // 置顶/取消置顶会话
    setTop() {
      let item = this.list[this.propIndex]
      item.is_top = !item.is_top
    }
  }
}
</script>

<style>
</style>
