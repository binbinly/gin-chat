<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar :title="nickname" fixed placeholder left-arrow @click-left="onClickLeft" @click-right="onClickRight">
      <template #right>
        <van-icon name="ellipsis" size="24" color="#0E151D" />
      </template>
    </van-nav-bar>
    <van-cell center>
      <template #icon>
        <van-image class="pr-1" round width="45" height="45" :src="info.avatar|formatAvatar" />
      </template>
      <template #title>
        <div class="flex flex-column">
          <span class="text-dark font-md font-weight-bold">{{nickname}}</span>
          <span class="text-light-muted">账号：{{info.username}}</span>
        </div>
      </template>
    </van-cell>
    <van-cell is-link v-if="is_friend" @click="openTag">
      <!-- 使用 title 插槽来自定义标题 -->
      <template #title>
        <div class="flex align-center">
          <span class="text-dark mr-1">标签</span>
          <div style="width:260px;" class="text-inline" v-if="friend.tags.length > 0">
            <span class="text-light-muted mr-1" v-for="(item,index) in friend.tags">{{item}}</span>
          </div>
          <span class="text-light-muted" v-else>未设置</span>
        </div>
      </template>
    </van-cell>
    <van-divider />
    <van-cell title="朋友圈" center is-link @click="openMoments">
      <template #default v-if="friend.moments">
        <span v-if="friend.moments[0].content && !friend.moments[0].image.length" class="text-secondary">{{friend.moments[0].content}}</span>
        <van-image v-for="(item,index) in friend.moments[0].image" :src="item" width="40" height="40" style="margin:5px;" />
      </template>
    </van-cell>
    <van-cell title="更多信息" is-link />
    <van-divider />
    <van-button v-if="is_friend" block :icon="!friend.is_black ? 'chat-o' : ''" type="default" @click="doEvent">{{friend.is_black ? '移出黑名单' : '发信息'}}
    </van-button>
    <van-button v-if="user.id != user_id && !is_friend" block icon="plus" type="default" @click="addFriend">添加好友</van-button>

    <!-- 操作菜单 -->
    <van-popup v-model="show" position="bottom" round style="height:40%">
      <van-cell is-link v-for="(item,index) in actions" :title="item.title" @click="popupEvent(item)">
        <!-- 使用 title 插槽来自定义标题 -->
        <template #icon>
          <span slot="icon" class="iconfont font-lg pr-1">{{item.icon}}</span>
        </template>
      </van-cell>
    </van-popup>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import event from '@/utils/event.js'
import { mapState } from 'vuex'
import { friendRead, friendDestroy, friendStar, friendBlack } from '@/api/friend.js'
import { Dialog, Toast } from 'vant';
import $const from '@/const/index.js';
export default {
  mixins: [auth],
  data() {
    return {
      show: false,
      user_id: 0,
      is_friend: false,
      info: {
        id: 0,
        username: "",
        nickname: "",
        avatar: "",
        sex: "",
        sign: "",
        area: "",
      },
      friend: {
        nickname: "",
        look_me: 1,
        look_him: 1,
        is_star: 0,
        is_black: 0,
        tags: []
      }
    }
  },
  activated() {
    this.user_id = parseInt(this.$route.query.id)
    if (!this.user_id) {
      return this.backToast()
    }
    // 获取当前用户资料
    this.getData()
  },
  mounted() {
    event.$on('saveRemarkTag', (e) => {
      this.nickname = e.nickname
      this.tagList = e.tagList
    })
  },
  computed: {
    ...mapState({
      chat: state => state.user.chat,
      user: state => state.user.user
    }),
    nickname() {
      return this.friend.nickname || this.info.nickname || this.info.username
    },
    actions() {
      return [{
        icon: "\ue6b3",
        title: "设置备注和标签",
        type: "push",
        path: '/user_remark_tag',
        query: {
          user_id: this.user_id,
          nickname: this.friend.nickname,
          tags: this.friend.tags ? this.friend.tags.join(',') : ''
        }
      }, {
        icon: "\ue613",
        title: "把他推荐给朋友",
        type: "push",
        path: "/chat_list",
        query: {
          type: $const.TYPE_CARD,
          content: this.info.nickname || this.info.username,
          options: { id: this.user_id, avatar: this.info.avatar }
        }
      }, {
        icon: "\ue6b0",
        title: this.friend.is_star ? '取消星标好友' : "设为星标朋友",
        type: "event",
        event: "setStar"
      }, {
        icon: "\ue667",
        title: "设置朋友圈和动态权限",
        type: "push",
        path: "/user_moments_auth",
        query: {
          id: this.user_id,
          look_me: this.friend.look_me,
          look_him: this.friend.look_him,
        }
      }, {
        icon: "\ue638",
        title: this.friend.is_black ? '移出黑名单' : "加入黑名单",
        type: "event",
        event: "setBlack"
      }, {
        icon: "\ue61c",
        title: "投诉",
        type: "push",
        path: "/user_report",
        query: {
          user_id: this.user_id,
          type: 1
        }
      }, {
        icon: "\ue638",
        title: "删除",
        type: "event",
        event: "deleteUser"
      }]
    }
  },
  destroyed() {
    event.$off('saveRemarkTag')
  },
  methods: {
    onClickRight() {
      this.show = true
    },
    getData() {
      friendRead(this.user_id).then(res => {
        if (res.moments && res.moments[0]['image']) {
          res.moments[0].image = res.moments[0].image.split(',')
          if (res.moments[0].image.length >= 3) { //截取4个元素
            res.moments[0].image = res.moments[0].image.splice(0, 3)
          }
        }
        this.is_friend = res.is_friend
        this.info = res.user
        if (res.is_friend) {
          this.friend = res.friend
        }
        console.log('detail', res)
      })
    },
    addFriend() {
      return this.$router.push({ path: '/add_friend', query: { id: this.user_id, act: 'apply', nickname: this.nickname } })
    },
    // 操作菜单事件
    popupEvent(e) {
      this.show = false
      if (!e.type) {
        return
      }
      switch (e.type) {
        case 'push':
          this.$router.push({ path: e.path, query: e.query })
          break;
        case 'event':
          this[e.event](e)
          break;
      }
    },
    // 删除好友
    deleteUser() {
      Dialog.confirm({
        message: '是否要删除该好友？',
      })
        .then(() => {
          friendDestroy(this.user_id).then(() => {
            Toast.success('删除好友成功')
            this.$router.replace({ path: '/home' })
          })
        }).catch(() => { })
    },
    // 设为星标
    setStar(e) {
      let star = this.friend.is_star == 0 ? 1 : 0
      friendStar({ user_id: this.user_id, star }).then(res => {
        this.friend.is_star = star
        e.title = this.friend.is_star ? '取消星标好友' : "设为星标朋友"
        Toast.success('操作成功')
      })
    },
    // 加入黑名单
    setBlack() {
      let msg = this.friend.is_black ? '移出黑名单' : '加入黑名单'
      Dialog.confirm({
        message: '是否要' + msg + '？',
      })
        .then(() => {
          const black = this.friend.is_black == 0 ? 1 : 0
          friendBlack({ user_id: this.user_id, black }).then(res => {
            this.friend.is_black = black
            Toast.success(msg + '成功')
          })
        }).catch(() => { })
    },
    doEvent() {
      if (this.friend.is_black) {
        return this.setBlack()
      }
      this.$router.push({ path: '/chat', query: { id: this.user_id, type: $const.CHAT_TYPE_USER } })
      this.chat.readChatItem(this.user_id, 'user')
    },
    openMoments() {
      this.$router.push({
        path: '/moments', query: {
          id: this.user_id,
        }
      })
    },
    openTag() {
      this.$router.push({
        path: '/user_remark_tag',
        query: {
          user_id: this.user_id,
          nickname: this.friend.nickname,
          tags: this.friend.tags ? this.friend.tags.join(',') : ''
        }
      })
    }
  }
}
</script>

<style>
</style>
