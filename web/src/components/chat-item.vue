<template>
  <div>
    <!-- 时间显示 -->
    <div v-if="showTime" class="flex align-center justify-center pb-1 pt-1">
      <span class="text-light-muted">{{showTime}}</span>
    </div>
    <!-- 撤回消息 -->
    <div v-if="item.isremove" ref="isremove" class="flex align-center justify-center pb-1 pt-1">
      <span class="text-light-muted">{{ isself ? '你' : '对方' }}撤回了一条信息</span>
    </div>
    <!-- 系统消息 -->
    <div v-if="isSystem" ref="isremove" class="flex align-center justify-center pb-1 pt-1">
      <span class="text-light-muted">{{item.content}}</span>
    </div>
    <!-- 气泡 -->
    <div v-if="!isSystem && !item.isremove" class="flex align-start position-relative mb-2" :class="!isself ? 'justify-start' : 'justify-end'"
         v-longpress="long">
      <!-- 好友 -->
      <template v-if="!isself">
        <van-image class="pl-1" round width="35" height="35" @click="openUser" :src="item.from.avatar|formatAvatar" />

        <span v-if="hasLabelClass" class="iconfont font-smaller text-white position-absolute chat-left-icon"
              :style="show_name ? 'top:22px;':'top:10px;'">&#xe609;</span>
      </template>

      <div v-if="isself && item.sendStatus && item.sendStatus !== 'success'" class="flex flex-column p-1">
        <span :class="item.sendStatus === 'fail' ? 'text-danger' : 'text-muted'">{{item.sendStatus === 'fail' ? 'X' : ''}}</span>
      </div>
      <div class="flex flex-column">
        <!-- 昵称 -->
        <div v-if="show_name" class="flex" :class="nicknameClass" style="max-width:180px;margin:0 15px;" :style="labelStyle">
          <span class="text-muted">{{item.from.name}}</span>
        </div>

        <div class="rounded" :class="labelClass" style="max-width:200px;" :style="labelStyle">
          <!-- 文字 -->
          <span v-if="isText" class="font-sm">{{item.content}}</span>
          <!-- 表情包 | 图片-->
          <div class="flex flex-wrap" v-else-if="isEmoticon || isImage">
            <van-image style="max-width: 200px;max-height:240px;" fit="cover" imageClass="rounded bg-secondary" @click="prediv(item.content)"
                       :src="item.content" />
          </div>

          <!-- 音频 -->
          <div v-else-if="isAudio" class="flex align-center" @click="openAudio">
            <van-image v-if="isself" class="mx-1" width="25" height="25" round
                       :src="!audioPlaying ? require('@/assets/audio/audio3.png') : require('@/assets/audio/play.gif')" />
            <span>{{item.options.time + '"'}}</span>
            <van-image v-if="!isself" class="mx-1" width="25" height="25" round
                       :src="!audioPlaying ? require('@/assets/audio/audio3.png') : require('@/assets/audio/play.gif')" />
          </div>

          <!-- 视频 -->
          <div v-else-if="isVideo" class="position-relative rounded" @click="openVideo">
            <video :src="item.content" style="max-width:200px;max-height:300px;"></video>
            <span class="iconfont text-white position-absolute" style="font-size: 35px;width:35px;height:35px;"
                  :style="posterIconStyle">&#xe737;</span>
          </div>

          <!-- 名片 -->
          <div v-else-if="isCard" class="bg-white" style="width: 180px;" hover-class="bg-light" @click="openUserBase">
            <div class="p-1 flex align-center border-bottom border-light-secondary">
              <van-image class="pr-1" round width="35" height="35" :src="item.options.avatar|formatAvatar" />
              <span class="font ml-1">{{item.content}}</span>
            </div>
            <div class="flex align-center p-1">
              <span class="font-small text-muted">个人名片</span>
            </div>
          </div>
        </div>

      </div>

      <!-- 本人 -->
      <template v-if="isself">
        <span v-if="hasLabelClass" class="iconfont text-chat-item font-smaller position-absolute chat-right-icon"
              :style="show_name ? 'top:20px;':'top:10px;'">&#xe640;</span>
        <van-image class="pr-1" round width="35" height="35" @click="openMy" :src="user.avatar|formatAvatar" />
      </template>
    </div>
  </div>
</template>

<script>
import $T from "@/utils/time.js"
import { mapState, mapActions } from 'vuex'
import $const from '@/const/index.js'
export default {
  props: {
    item: Object,
    index: Number,
    // 上一条消息的时间戳
    pretime: [Number, String],
    show_name: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      test_url: '',
      innerAudioContext: null,
      audioPlaying: false,
      // 默认封面的宽高
      poster: {
        w: 100,
        h: 100
      }
    }
  },
  computed: {
    ...mapState({
      user: state => state.user.user
    }),
    isSystem() {//是否系统消息
      return this.item.type === $const.TYPE_SYSTEM
    },
    isText() {//是否文本消息
      return this.item.type === $const.TYPE_TEXT
    },
    isImage() {//是否图片消息
      return this.item.type === $const.TYPE_IMAGE
    },
    isVideo() {//是否视频消息
      return this.item.type === $const.TYPE_VIDEO
    },
    isAudio() {//是否音频消息
      return this.item.type === $const.TYPE_AUDIO
    },
    isEmoticon() {//是否表情消息
      return this.item.type === $const.TYPE_EMOTICON
    },
    isCard() {//是否名片消息
      return this.item.type === $const.TYPE_CARD
    },
    isself() {// 是否是本人
      return this.isSystem ? false : (this.item.self ? this.item.self : false)
    },
    showTime() {// 显示的时间
      return $T.getChatTime(this.item.t, this.pretime)
    },
    hasLabelClass() {// 是否需要气泡样式
      return this.isText || this.isAudio
    },
    labelClass() {// 气泡的样式
      const labelSelf = this.hasLabelClass ? 'bg-chat-item mr-1 p-1' : 'mr-1'
      const label = this.hasLabelClass ? 'bg-white ml-1 p-1' : 'ml-1'
      return this.isself ? labelSelf : label
    },
    nicknameClass() {
      let c = this.isself ? 'justify-end' : ''
      return c
    },
    labelStyle() {
      if (this.item.type === 'audio') {
        let time = this.item.options.time || 0
        let width = parseInt(time) / (60 / 200)
        width = width < 75 ? 75 : width
        return `width:${width}px;`
      }
    },
    // 短视频封面图标位置
    posterIconStyle() {
      let w = this.poster.w / 2 - 35 / 2
      let h = this.poster.h / 2 - 45 / 2
      return `left:${w}px;top:${h}px;`
    }
  },
  mounted() {
    // 注册全局事件
    if (this.isAudio) {
      this.audioOn(this.onPlayAudio)
    }
    if (this.isVideo) {
      let video = document.querySelector('video');
      //canplay 事件，视频达到可以播放时触发；
      video.addEventListener('canplay', () => {
        this.loadPoster(video.videoWidth, video.videoHeight)
      });
    }
  },
  // 组件销毁
  destroyed() {
    if (this.isAudio) {
      this.audioOff(this.onPlayAudio)
    }
    // 销毁音频
    if (this.innerAudioContext) {
      this.innerAudioContext.destroy()
      this.innerAudioContext = null
    }
  },
  methods: {
    ...mapActions(['audioOn', 'audioEmit', 'audioOff']),
    //有些视频第一帧是黑的，获取到的是纯黑图片，改成取视频的第n秒做封面：
    // genVideoCover(url, callback, second = 1) {
    //   const video = document.createElement('video')
    //   video.src = url
    //   video.setAttribute('crossorigin', 'anonymous')
    //   video.addEventListener('loadeddata', function() {
    //     // 设置currentTime
    //     this.currentTime = second
    //   })
    //   video.addEventListener('seeked', () => {
    //     const imageUrl = this.canvasToDataUrl(video)
    //     callback && callback(imageUrl)
    //   })
    //   video.addEventListener('error', function() {
    //     callback && callback()
    //   })
    // },
    // canvasToDataUrl(video) {
    //   const canvas = document.createElement('canvas')
    //   const ctx = canvas.getContext('2d')
    //   const imgHeight = video.videoHeight
    //   const imgWidth = video.videoWidth
    //   ctx.drawImage(video, 0, 0, imgWidth, imgHeight)
    //   this.loadPoster(imgWidth, imgHeight)
    //   // 设置图片质量为0.75，如果是缩略图一般取0.35
    //   return canvas.toDataURL('image/jpeg', 0.35)
    // },
    openUser() {
      const id = this.item.chat_type === $const.CHAT_TYPE_USER ? this.item.from.id : this.item.to_id
      this.$router.push({ path: '/user_base', query: { id } })
    },
    openMy() {
      const id = this.user.id
      this.$router.push({ path: '/user_base', query: { id } })
    },
    // 打开名片
    openUserBase() {
      console.log('id', this.item.options.id)
      this.$router.push({ path: '/user_base', query: { id: this.item.options.id } })
    },
    // 加载封面
    loadPoster(w, h) {
      const scale = w / h
      if (scale > 1) {//宽 > 高
        if (w > 200) { //缩放
          this.poster.w = 200
          this.poster.h = parseInt(200 / w * h)
          return
        }
      } else {
        if (h > 300) { //缩放
          this.poster.h = 300
          this.poster.h = parseInt(300 / h * w)
          return
        }
      }
      this.poster.w = w
      this.poster.h = h
    },
    // 监听播放音频全局事件
    onPlayAudio(index) {
      if (this.innerAudioContext) {
        if (this.index !== index) {
          this.innerAudioContext.pause()
        }
      }
    },
    openVideo() {// 播放视频
      this.$emit('openVideo', this.item.content)
    },
    // 播放音频
    openAudio() {
      // 通知停止其他音频
      this.audioEmit(this.index)
      if (!this.innerAudioContext) {
        this.innerAudioContext = uni.createInnerAudioContext();
        this.innerAudioContext.src = this.item.content
        this.innerAudioContext.play()
        // 监听播放
        this.innerAudioContext.onPlay(() => {
          this.audioPlaying = true
        })
        // 监听暂停
        this.innerAudioContext.onPause(() => {
          this.audioPlaying = false
        })
        // 监听停止
        this.innerAudioContext.onStop(() => {
          this.audioPlaying = false
        })
        // 监听错误
        this.innerAudioContext.onError(() => {
          this.audioPlaying = false
        })
      } else {
        this.innerAudioContext.stop()
        this.innerAudioContext.play()
      }
    },
    // 预览图片
    prediv(url) {
      this.$emit('prediv', url)
    },
    // 长按事件
    long(e) {
      const x = e.changedTouches[0].pageX
      const y = e.changedTouches[0].pageY
      this.$emit('long', {
        x,
        y,
        index: this.index
      })
    }
  }
}
</script>

<style scoped>
.chat-left-icon {
  left: 50px;
}
.chat-right-icon {
  right: 50px;
}
.mr-1 {
  margin-right: 13px;
}
.ml-1 {
  margin-left: 13px;
}
</style>
