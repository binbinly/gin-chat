<template>
  <div id="content">
    <div class="body" @click="show=false">
      <transparent-bar :scrollTop="scrollTop" @clickRight="clickRight" :title="title"></transparent-bar>
      <div class="position-relative" style="height: 620rpx;">
        <van-image :src="require('@/assets/images/bg.jpg')" fit="cover" height="280" class="bg-secondary w-100"></van-image>
        <van-image :src="userinfo.avatar|formatAvatar" width="60" height="60" style="right: 15px;bottom:-20px;" fit="cover"
                   class="bg-secondary rounded position-absolute">
        </van-image>
        <span class="text-white font-sm position-absolute" style="bottom: 10px;right: 80px;">{{userinfo.name}}</span>
      </div>

      <div class="w-100 text-center mt-1" v-if="showNotice">
        <van-tag round type="primary" size="medium" @click="onRefresh">有新动态哦</van-tag>
      </div>
      <!-- 朋友圈列表 -->
      <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
        <van-list v-model="loading" :finished="finished" finished-text="没有更多了" @load="onLoad">
          <moment-list v-for="(item,index) in list" :key="index" :item="item" :index="index" @action="doAction" @reply="replyEvent"
                       @openVideo="openVideo"></moment-list>
        </van-list>
      </van-pull-refresh>
    </div>

    <!-- 评论弹出层 -->
    <van-popup v-model="show" position="bottom" :overlay="false" :lock-scroll="false">
      <div style="height: 50px;" class="bg-light border-top flex align-center px-1">
        <van-field v-model="content" rows="1" type="textarea" :placeholder="placeholder || '文明发言'" style="height:43px;width:75%;" />
        <icon-button :icon="'\ue605'" @click="changeFaceModal" />
        <van-button type="primary" size="small" :disabled="content.length === 0" @click="send">发送</van-button>
      </div>
      <div v-if="faceModal" class="flex flex-wrap" style="height: 200px;overflow: auto;">
        <div style="width: 45px;height: 45px;" class="flex align-center justify-center" hover-class="bg-white" v-for="(item,index) in faceList"
             @click="addFace(item)">
          <span style="font-size:24px;">{{item}}</span>
        </div>
      </div>
    </van-popup>

    <van-overlay :show="showVideo" @click="showVideo = false">
      <div class="wrapper">
        <video :src="videoUrl" controls class="w-100"></video>
      </div>
    </van-overlay>

    <van-action-sheet v-model="showAction" :actions="actions" cancel-text="取消" close-on-click-action @cancel="onCancel" @select="onSelect" />
  </div>
</template>

<script>
import IconButton from '@/components/icon-button.vue';
import TransparentBar from '@/components/transparent-bar.vue';
import MomentList from '@/components/moment-list.vue';
import { mapState } from 'vuex'
import { Toast } from 'vant';
import event from '@/utils/event.js'
import $const from '@/const/index.js'
import { momentLike, momentComment, momentList } from '@/api/moment.js'
export default {
  components: {
    IconButton,
    TransparentBar,
    MomentList
  },
  data() {
    return {
      showVideo: false,
      videoUrl: '',
      showNotice: false,
      showAction: false,
      actions: [{
        name: "图文",
        type: "image"
      }, {
        name: "短视频",
        type: "video"
      }, {
        name: "文字",
        type: "text"
      }],
      show: false,
      placeholder: '',
      content: "",
      scrollTop: 0,
      list: [],

      faceModal: false,
      faceList: ["😀", "😁", "😂", "😃", "😄", "😅", "😆", "😉", "😊", "😋", "😎", "😍", "😘", "😗", "😙", "😚", "😇", "😐", "😑", "😶", "😏", "😣", "😥", "😮", "😯", "😪", "😫", "😴", "😌", "😛", "😜", "😝", "😒", "😓", "😔", "😕", "😲", "😷", "😖", "😞", "😟", "😤", "😢", "😭", "😦", "😧", "😨", "😬", "😰", "😱", "😳", "😵", "😡", "😠"],
      // 评论的对象
      commentIndex: -1,

      loading: false,
      finished: false,
      refreshing: false,
      page: 1,
      reply_user: false,

      user_id: 0,
      userinfo: {
        id: 0,
        name: '',
        avatar: ''
      },
      title: ''
    }
  },
  computed: {
    ...mapState({
      user: state => state.user.user,
      chat: state => state.user.chat
    }),
    nickname() {
      if (!this.params) {
        return this.user.nickname || this.user.username
      }
      return this.params.name
    },
    avatar() {
      const avatar = !this.params ? this.user.avatar : this.params.avatar
      return avatar ? avatar : require('@/assets/images/userpic.png')
    }
  },
  activated() {
    const id = parseInt(this.$route.query.id) || 0
    console.log('store', this.user_id)
    if (id) {
      if (this.user_id != id) {
        this.user_id = id
        this.refresh()
      }
    }
    console.log('query', id)
  },
  mounted() {
    this.chat.readMoments()
    event.$on('momentNotice', this.momentNotice)
    event.$on('refreshMoment', this.onRefresh)
    //绑定滚动事件
    window.addEventListener('scroll', this.scroll, true);
  },
  destroyed() {
    event.$off('momentNotice', this.momentNotice)
    event.$off('refreshMoment', this.onRefresh)
    window.removeEventListener('scroll', this.scroll); // 离开页面清除（移除）滚轮滚动事件
  },
  methods: {
    momentNotice(notice) {
      if (notice.user_id && notice.num) {
        this.showNotice = true
      }
    },
    scroll(e) {
      //可滚动总高度
      // const clientHeight = document.documentElement.clientHeight || document.body.clientHeight;
      //距离顶部高度
      this.scrollTop = document.documentElement.scrollTop
      if (this.scrollTop > 240) {
        this.title = '朋友圈'
      } else {
        this.title = ''
      }
    },
    onLoad() {
      this.showNotice = false
      this.getData()
    },
    onRefresh() {
      this.chat.readMoments()
      this.refresh()
    },
    refresh() {
      this.page = 1
      // 清空列表数据
      this.finished = false;
      // 重新加载数据
      // 将 loading 设置为 true，表示处于加载状态
      this.loading = true;
      this.onLoad();
    },
    getData() {
      momentList(this.user_id, this.page).then(res => {
        if (this.refreshing) {
          this.list = [];
          this.refreshing = false;
        }
        if (this.page == 1) {
          this.userinfo = res['user']
        }
        this.list = this.page === 1 ? res['list'] : [...this.list, ...res['list']]
        this.loading = false;
        if (res['list'].length < $const.PAGE_SIZE) {
          this.finished = true;
        } else {
          this.page++
        }
      })
    },
    // 点击操作菜单
    doAction({ event, item, index }) {
      if (event === 'like') {//点赞
        return this.doSupport(item)
      } else if (event === 'comment') {//评论
        this.show = true
        this.commentIndex = index
        this.reply_user = false
      }
    },
    openVideo(url) {
      this.showVideo = true
      this.videoUrl = url
    },
    initComment() {
      this.content = ''
      this.faceModal = false
      this.reply_user = false
    },
    // 点赞
    doSupport(item) {
      momentLike({
        id: item.id
      }).then(() => {
        let i = item.likes.findIndex(val => val.id === this.user.id)
        if (i !== -1) { // 取消点赞
          item.likes.splice(i, 1)
        } else { // 点赞
          item.likes.push({
            id: this.user.id,
            name: this.user.nickname || this.user.username
          })
        }
        Toast.success(i !== -1 ? '取消点赞成功' : '点赞成功')
      })
    },
    // 添加表情
    addFace(item) {
      this.content += item
    },
    // 开启/关闭表情包面板
    changeFaceModal() {
      setTimeout(() => {
        this.faceModal = !this.faceModal
      }, 100)
    },
    // 发送
    send() {
      let item = this.list[this.commentIndex]
      momentComment({
        id: item.id,
        content: this.content,
        reply_id: this.reply_user ? this.reply_user.id : 0
      }).then(() => {
        item.comments.push({
          content: this.content,
          user: {
            id: this.user.id,
            name: this.user.nickname || this.user.username
          },
          reply: this.reply ? this.reply : null
        })
        Toast.success('评论成功')
        this.initComment()
      })
      this.show = false
    },
    // 选择发表朋友圈类型
    clickRight() {
      this.showAction = true
    },
    replyEvent(e) {
      this.content = ''
      this.faceModal = false
      this.commentIndex = e.index
      this.reply_user = e.reply
      this.show = true
      this.placeholder = '回复' + e.reply.name + ':'
    },
    onCancel() {

    },
    onSelect({ type }) {
      this.$router.push({ path: '/add_moment', query: { type } })
    }
  }
}
</script>

<style>
</style>
