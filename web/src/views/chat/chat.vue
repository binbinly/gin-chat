<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar :title="detail.name" left-arrow @click-left="onClickLeft" @click-right="onChatSet">
      <template #right>
        <van-icon name="ellipsis" size="24" color="#0E151D" />
      </template>
    </van-nav-bar>

    <!-- 聊天内容区域 -->
    <div ref="chat" @click="hidePopup" class="w-100" style="overflow: auto;" :style="'height:'+chatHeight+'px;'">
      <!-- 聊天信息列表组件 -->
      <div v-for="(item,index) in list" :key="index" :id="'chatItem_'+index">
        <chat-item :item="item" :index="index" ref="chatItem" :pretime=" index > 0 ? list[index-1].t : 0" @long="long" @prediv="predivImage"
                   @openVideo="openVideo" :show_name="currentChatItem.show_name"></chat-item>
      </div>
    </div>

    <!-- 底部输入框 -->
    <div class="position-fixed left-0 right-0 border-top flex align-center" style="background-color: #F7F7F6;height: 45px;"
         :style="'bottom:'+KeyboardHeight+'px;'">
      <icon-button v-if="mode === 'audio'" :icon="'\ue607'" @click="changeVoiceOrText"></icon-button>
      <icon-button v-else :icon="'\ue606'" @click="changeVoiceOrText"></icon-button>
      <div class="flex-1">
        <div v-if="mode === 'audio'" class="rounded flex align-center justify-center" style="height: 45px;"
             :class="isRecording?'bg-hover-light':'bg-white'" @touchstart="voiceTouchStart" @touchend="voiceTouchEnd" @touchcancel="voiceTouchCancel"
             @touchmove="voiceTouchMove">
          <span>{{isRecording ? '松开 结束':'按住 说话'}}</span>
        </div>
        <van-field v-else v-model="text" rows="1" type="textarea" placeholder="请输入" @focus="textFocus" style="height:43px;width:240px;" />
      </div>
      <!-- 表情 -->
      <icon-button :icon="'\ue605'" @click="openActionOrEmoticon('emoticon')"></icon-button>
      <template v-if="text.length === 0">
        <!-- 扩展菜单 -->
        <icon-button :icon="'\ue603'" @click="openActionOrEmoticon('action')"></icon-button>
      </template>
      <div v-else class="flex-shrink" style="width:60px;">
        <!-- 发送按钮 -->
        <van-button type="primary" size="small" @click="send('text')">发送</van-button>
      </div>
    </div>

    <!-- 扩展菜单 -->
    <van-popup v-model="show" position="bottom" :overlay="false" transition-appea>
      <div style="height: 203px;" class="border-top border-light-secondary bg-light flex flex-wrap">
        <template v-if="mode=='action'">
          <van-swipe :loop=" false">
            <van-swipe-item v-for="(item,index) in emoticonOrActionList" :key="index">
              <van-grid :column-num="3">
                <van-grid-item v-for="(item2,index2) in item" @click="actionEvent(item2)" :text="item2.name">
                  <template #icon>
                    <van-image :src="item2.icon" fit="center" width="50" height="50" />
                  </template>
                </van-grid-item>
              </van-grid>
            </van-swipe-item>
          </van-swipe>
        </template>
        <template v-else>
          <van-tabs class="w-100" @change="beforeChange" animated swipeable sticky>
            <van-tab title="经典">
              <template v-for="(item,index) in emoticonList">
                <span style="font-size:24px;margin:5px;" @click="addFace(item)">{{item}}</span>
              </template>
            </van-tab>
            <van-tab v-for="(item,index) in emoCat" :title="item.name|substr" style="height:160px;overflow: auto;">
              <template v-for="(item2,index2) in item['list']">
                <van-image class="ml-1 mt-1" :src="item2" fit="center" width="80" height="80" @click="sendEmoticon(item2)" />
              </template>
            </van-tab>
            <van-tab title="更多" style="text-align:center;margin-top:60px;">
              <van-button type="primary" to="/emoticon_cat">添加表情</van-button>
            </van-tab>
          </van-tabs>
        </template>
      </div>
    </van-popup>

    <!-- 上传图片  -->
    <van-uploader style="display:none" name="image" :prediv-image="false" multiple :max-count="9" :after-read="afterRead" accept="image/*"
                  ref="uploadImage" />
    <!-- 拍摄  -->
    <van-uploader style="display:none" name="video" capture="camera" :prediv-image="false" :after-read="afterRead" accept=".mp4" ref="uploadVideo" />
    <!-- 扩展菜单  -->
    <free-popup ref="extend" :bodyWidth="240" :bodyHeight="450" :tabbarHeight="105">
      <div class="flex flex-column text-white p-1" style="width:90px;" :style="getMenusStyle">
        <div class="flex-1 flex align-center justify-center" style="padding:5px;" v-for="(item,index) in menusList" @click="clickEvent(item.event)">
          <span class="font-sm">{{item.name}}</span>
        </div>
      </div>
    </free-popup>

    <van-overlay :show="showVideo" @click="showVideo = false">
      <div class="wrapper">
        <video :src="videoUrl" controls class="w-100"></video>
      </div>
    </van-overlay>

    <!-- 录音提示 -->
    <div v-if="isRecording" class="position-fixed top-0 left-0 right-0 flex align-center justify-center" style="bottom: 45px;">
      <div style="width: 140px;height: 140px;background-color: rgba(0,0,0,0.5);" class="rounded flex flex-column align-center justify-center">
        <van-image :src="require('@/assets/audio/recording.gif')" width="75" height="75"></van-image>
        <span class="text-white mt-1">{{unRecord ? '松开手指，取消发送':'手指上滑，取消发送'}}</span>
      </div>
    </div>
  </div>
</template>

<script>
import ChatItem from '@/components/chat-item.vue';
import IconButton from '@/components/icon-button.vue';
import FreePopup from '@/components/free-popup.vue';
import { chatDetail } from '@/api/chat.js'
import { collectCreate } from '@/api/collect.js'
import event from '@/utils/event.js'
import { uploadFile } from '@/api/common.js'
import { mapState, mapMutations } from 'vuex'
import auth from '@/mixin/auth.js';
import { Dialog, Toast, ImagePreview } from 'vant';
import { emoticon } from '@/api/common.js'
import $const from '@/const/index.js'
export default {
  mixins: [auth],
  components: {
    ChatItem,
    IconButton,
    FreePopup
  },
  data() {
    return {
      show: false,
      showVideo: false,
      videoUrl: '',
      chatHeight: 575,
      // 模式 text输入文字，emoticon表情，action操作，audio音频
      mode: "text",
      // 扩展菜单列表
      actionList: [
        [{
          name: "相册",
          icon: require("@/assets/images/extends/pic.png"),
          event: "uploadImage"
        }, {
          name: "拍摄",
          icon: require("@/assets/images/extends/video.png"),
          event: "uploadVideo"
        }, {
          name: "收藏",
          icon: require("@/assets/images/extends/shoucan.png"),
          event: "openFava"
        }, {
          name: "名片",
          icon: require("@/assets/images/extends/man.png"),
          event: "sendCard"
        }, {
          name: "语音通话",
          icon: require("@/assets/images/extends/phone.png"),
          event: ""
        }, {
          name: "位置",
          icon: require("@/assets/images/extends/path.png"),
          event: ""
        }]
      ],
      emoCat: [],
      emoticonList: ["😀", "😁", "😂", "😃", "😄", "😅", "😆", "😉", "😊", "😋", "😎", "😍", "😘", "😗", "😙", "😚", "😇", "😐", "😑", "😶", "😏", "😣", "😥", "😮", "😯", "😪", "😫", "😴", "😌", "😛", "😜", "😝", "😒", "😓", "😔", "😕", "😲", "😷", "😖", "😞", "😟", "😤", "😢", "😭", "😦", "😧", "😨", "😬", "😰", "😱", "😳", "😵", "😡", "😠"],
      // 键盘高度
      KeyboardHeight: 0,
      menusList: [],
      navBarHeight: 46,
      list: [],
      // 当前操作的气泡索引
      propIndex: -1,
      // 输入文字
      text: "",

      // 音频录制状态
      isRecording: false,
      RecordingStartY: 0,
      // 取消录音
      unRecord: false,

      detail: {
        id: 0,
        name: "",
        avatar: "",
        chat_type: 1
      },
    }
  },
  //挂载到实例上去之后调用
  created() {
    // 注册发送音频事件
    this.regSendVoiceEvent((url) => {
      if (!this.unRecord) {
        this.send('audio', url, {
          time: this.RecordTime
        })
      }
    })
  },
  computed: {
    ...mapState({
      chatList: state => state.user.chatList,
      RECORD: state => state.audio.RECORD,
      RecordTime: state => state.audio.RecordTime,
      chat: state => state.user.chat,
      totalunread: state => state.user.totalunread,
      user: state => state.user.user,
    }),
    // 当前会话配置信息
    currentChatItem() {
      let index = this.chatList.findIndex(item => item.id === this.detail.id && item.chat_type === this.detail.chat_type)
      if (index !== -1) {
        return this.chatList[index]
      }
      return {}
    },
    // 动态获取菜单高度
    getMenusHeight() {
      const H = 30
      return this.menusList.length * H
    },
    // 获取菜单的样式
    getMenusStyle() {
      return `height: ${this.getMenusHeight}rpx;`
    },
    // 判断是否操作本人信息
    isdoSelf() {
      // 获取本人id（假设拿到了）
      let id = 1
      let user_id = this.propIndex > -1 ? this.list[this.propIndex].user_id : 0
      return user_id === id
    },
    // 获取操作或者表情列表
    emoticonOrActionList() {
      return (this.mode === 'emoticon' || this.mode === 'action') ? this[this.mode + 'List'] : []
    },
    // 所有信息的图片地址
    imageList() {
      let arr = []
      this.list.forEach((item) => {
        if (item.type === $const.TYPE_EMOTICON || item.type === $const.TYPE_IMAGE) {
          arr.push({ url: item.content, type: item.type })
        }
      })
      return arr
    },
    images() {
      return this.imageList.map(item => item.url)
    }
  },
  activated() {
    const { id, type } = this.$route.query
    if (!id) {
      return this.backToast()
    }
    this.detail.id = parseInt(id)
    if (type) {
      this.detail.chat_type = parseInt(type)
    }
    // 初始化
    this.initData()
    // 创建聊天对象
    this.chat.createChatObject(this.detail)
    // 获取历史记录
    this.list = this.chat.getChatDetail(this.detail.chat_type, this.detail.id)
    // 触底
    setTimeout(() => {
      this.pageToBottom()
    }, 500);
  },
  mounted() {
    // 监听接收聊天信息
    event.$on('onMessage', this.onMessage)
    // 监听撤销聊天消息
    event.$on('onRecall', this.onRecall)

    event.$on('updateHistory', this.updateHistory)
    // 监听发送收藏和名片
    event.$on('sendItem', this.onSendItem)
    event.$on('onEmoticon', this.onEmoticon)
  },
  destroyed() {
    // 销毁聊天对象
    this.chat.destroyChatObject()
    // 销毁监听接收聊天消息
    event.$off('onMessage', this.onMessage)
    // 销毁监听撤销聊天消息
    event.$off('onRecall', this.onRecall)

    event.$off('updateHistory', this.updateHistory)

    event.$off('sendItem', this.onSendItem)
    event.$off('onEmoticon', this.onEmoticon)
  },
  methods: {
    ...mapMutations(['regSendVoiceEvent']),
    beforeChange(index) {
      const len = this.emoCat.length
      if (index == 0 || len == 0 || len + 1 == index) {
        return true
      }
      index--;
      if (this.emoCat[index].list.length == 0) {
        emoticon(this.emoCat[index].name).then(res => {
          this.emoCat[index].list = res.map(item => {
            return item.url
          })
        })
      }
      return true
    },
    textFocus() {
      this.mode = 'text'
      this.hidePopup()
    },
    hidePopup() {
      this.show = false
      this.KeyboardHeight = 0
      this.chatHeight = 575
    },
    onEmoticon(cat) {
      this.emoCat.push({ name: cat, list: [] })
    },
    onSendItem(e) {
      console.log('e', e)
      if (e.sendType === 'fava' || e.sendType === 'card') {
        this.send($const.TYPE_TRANS_LIST[e.type], e.content, e.options)
      }
    },
    // 添加表情
    addFace(item) {
      this.text += item
    },
    updateHistory(isclear = true) {
      if (isclear) {
        this.list = []
      } else {
        this.list = this.chat.getChatDetail()
      }
    },
    onRecall(message) {
      console.log('[聊天页] 监听撤销聊天信息', message);
      // 撤回消息
      const index = this.list.findIndex(item => item.id === message.id)
      if (index !== -1) {
        this.list.splice(index, 1)
      }
    },
    onMessage(message) {
      console.log('[聊天页] 监听接收聊天信息', message);
      if ((message.from.id === this.detail.id && message.chat_type === $const.CHAT_TYPE_USER) || (message.chat_type === $const.CHAT_TYPE_GROUP && message.to.id === this.detail.id)) {
        this.list.push(message)
        // 置于底部
        return this.pageToBottom()
      }
    },
    initData() {
      if (this.emoCat.length == 0 && this.$store.state.user.emoCat) {
        this.emoCat = this.$store.state.user.emoCat.map(item => {
          return { name: item, list: [] }
        })
        console.log('emoCat', this.emoCat)
      }
      chatDetail({
        id: this.detail.id,
        type: this.detail.chat_type
      }).then(res => {
        this.detail.name = res.name;
        this.detail.avatar = res.avatar;
      })
    },
    // 打开扩展菜单或者表情包
    openActionOrEmoticon(mode = 'action') {
      this.mode = mode
      this.show = true
      this.KeyboardHeight = 203
      this.chatHeight = 370
      this.pageToBottom()
    },
    // 发送
    send(type, content = '', options = {}) {
      // 组织数据格式
      switch (type) {
        case 'text':
          content = content || this.text
          break;
      }
      let message = this.chat.formatSendData({
        chat_type: this.detail.chat_type,
        type,
        content,
        options
      })
      // 渲染到页面
      let index = this.list.length
      message.progress = 0
      this.list.push(message)
      // 监听上传进度
      let onProgress = false
      if (type !== 'text' && type !== 'emoticon' && type !== 'card' && !message.content.startsWith('http')) {
        onProgress = (progress) => {
          this.list[index].progress = progress;
        }
      }
      // 发送到服务端
      this.chat.send(message, onProgress).then(res => {
        // 发送成功
        this.list[index].id = res.id
        this.list[index].from.name = res.from.name
        this.list[index].status = 'success'
      }).catch(err => {
        // 发送失败
        this.list[index].status = 'fail'
        console.log(err);
      })
      // 发送文字成功，清空输入框
      if (type === 'text') {
        this.text = ''
      }
      // 置于底部
      this.pageToBottom()
    },
    // 回到底部
    pageToBottom() {
      this.$nextTick(() => {
        const msg = this.$refs.chat // 获取对象
        if (msg) {
          msg.scrollTop = msg.scrollHeight // 滚动高度
        }
      })
    },
    // 长按消息气泡
    long({ x, y, index }) {
      // 初始化 索引
      this.propIndex = index
      // 组装菜单
      let menus = [{
        name: "发送给朋友",
        event: 'sendToChatItem'
      }, {
        name: '复制',
        event: 'copy'
      }, {
        name: "收藏",
        event: 'fava'
      }, {
        name: "删除",
        event: 'delete'
      }]
      let item = this.list[this.propIndex]
      let isSelf = this.user.id === item.from.id
      if (isSelf) {
        menus.push({
          name: "撤回",
          event: 'removeChatItem'
        })
      }
      this.menusList = menus
      // 显示扩展菜单
      this.$refs.extend.show(x, y)
    },
    // 操作菜单方法分发
    clickEvent(event) {
      console.log('event', event)
      let item = this.list[this.propIndex]
      let isSelf = this.user.id === item.from.id
      switch (event) {
        case 'removeChatItem': // 撤回消息
          // 拿到当前被操作的信息
          this.chat.recall(item).then(() => {
            this.list.splice(this.propIndex, 1)
          })
          break;
        case 'sendToChatItem':
          this.$router.push({ path: '/chat_list', query: { ...item } })
          break;
        case 'copy': // 复制
          this.$copyText(item.data).then(() => {
            Toast.success('复制成功')
          })
          break;
        case 'delete':
          Dialog.confirm({
            message: '是否要删除该记录？',
          })
            .then(() => {
              this.chat.deleteChatDetailItem(item.id, this.detail.chat_type, this.detail.id)
              this.list.splice(this.propIndex, 1)
              // 删除最后一条消息
              if (this.list.length === this.propIndex) {
                this.chat.updateChatItem({
                  id: this.detail.id,
                  chat_type: this.detail.chat_type
                }, (v) => {
                  let o = this.list[this.propIndex - 1]
                  let data = ''
                  if (o) {
                    data = this.chat.formatChatItemData(o, isSelf)
                  }
                  v.content = data
                  return v
                })
              }
            })
          break;
        case 'fava': // 加入收藏
          Dialog.confirm({
            message: '是否要加入收藏？',
          })
            .then(() => {
              let options = ''
              if (item.type !== $const.TYPE_VIDEO) {
                options = item.options
              }
              collectCreate({
                type: item.type,
                content: item.content,
                options,
              }).then(res => {
                Toast.success('加入收藏成功')
              })
            })
          break;
      }
      // 关闭菜单
      this.$refs.extend.hide()
    },
    sendEmoticon(url) {
      this.send('emoticon', url)
    },
    // 扩展菜单
    actionEvent(e) {
      switch (e.event) {
        case 'uploadImage': // 选择相册
          this.$refs.uploadImage.chooseFile()
          break;
        case 'uploadVideo': // 发送短视频
          this.$refs.uploadVideo.chooseFile()
          break;
        case 'openFava': // 打开收藏
          this.$router.push({ path: '/fava', query: { type: 'send' } })
          break;
        case 'sendCard': // 发送名片
          this.$router.push({
            path: '/contact_list',
            query: { type: 'sendCard', limit: 1 }
          })
          break;
        default:
          Toast('待开发')
          break;
      }
    },
    // 点击页面
    clickPage() {
      this.mode = ''
    },
    openVideo(url) {
      this.showVideo = true
      this.videoUrl = url
    },
    // 预览图片
    predivImage(url) {
      console.log('image', this.imageList)
      let index = this.imageList.findIndex(item => {
        return item.url === url
      })
      if (index <= 0) {
        index = 0
      }
      ImagePreview({
        images: this.images,
        startPosition: index,
        closeable: true
      });
    },
    // 切换音频录制和文本输入
    changeVoiceOrText() {
      this.mode = this.mode !== 'audio' ? 'audio' : 'text'
    },
    // 录音相关
    // 录音开始
    voiceTouchStart(e) {
      // 初始化
      this.isRecording = true
      this.RecordingStartY = e.changedTouches[0].screenY
      this.unRecord = false
      // 开始录音
      this.RECORD.start({
        format: "mp3"
      })
    },
    // 录音结束
    voiceTouchEnd() {
      this.isRecording = false
      // 停止录音
      this.RECORD.stop()
    },
    // 录音被打断
    voiceTouchCancel() {
      this.isRecording = false
      this.unRecord = true
      // 停止录音
      this.RECORD.stop()
    },
    voiceTouchMove(e) {
      let Y = Math.abs(e.changedTouches[0].screenY - this.RecordingStartY)
      this.unRecord = (Y >= 50)
    },
    // 打开聊天信息设置
    onChatSet() {
      this.$router.push({
        path: '/chat_set', query: {
          id: this.detail.id,
          chat_type: this.detail.chat_type
        }
      })
    },
    //	文件读取完成后的回调函数
    afterRead(files, { name }) {
      //上传文件
      if (files instanceof Array) {
        files.forEach(item => {
          this.upload(item, name)
        })
      } else {
        this.upload(files, name)
      }
    },
    upload(file, name) {
      file.status = 'uploading';
      file.message = '上传中...';
      uploadFile(file.file).then(url => {
        file.status = 'done'
        this.send(name, url)
      }).catch(rej => {
        Toast.fail('上传失败')
      })
    },
    onSelect(action, index) {
      console.log('select', action, index)
    }
  }
}
</script>

<style>
.van-image-prediv__cover {
  top: 50%;
  margin-top: -25%;
}
</style>
