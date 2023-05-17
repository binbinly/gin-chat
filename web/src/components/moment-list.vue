<template>
  <div class="px-1 pt-1 flex align-start border-bottom border-light-secondary">
    <free-avater :src="item.user.avatar" :uid="item.user.id" />
    <div class="pl-1 flex-1 flex flex-column">
      <!-- 昵称 -->
      <span class="text-hover-primary" style="margin-bottom:5px;">{{item.user.name}}</span>
      <!-- 内容 -->
      <span v-if="item.content" class="text-dark font-sm">{{item.content}}</span>
      <!-- 图片 -->
      <div v-if="item.image" class="pt-1 flex flex-wrap">
        <template v-for="(image,imageIndex) in imgs">
          <!-- 单图 -->
          <van-image v-if="imgs.length === 1" :src="image" fit="cover" style="max-width:180px;max-height:240px;" imageClass="rounded bg-secondary"
                     @click="prediv(image)">
          </van-image>
          <!-- 多图 -->
          <van-image v-else :src="image" fit="cover" width="90" height="90" class="bg-secondary rounded" style="margin:0 5px 5px 0;"
                     @click="prediv(image)"></van-image>
        </template>
      </div>
      <!-- 视频 -->
      <!-- <div v-if="item.video" class="pt-1">
        <video :src="item.video" controls style="max-width:180px;max-height:300px;"></video>
      </div> -->
      <!-- 视频 -->
      <div v-if="item.video" class="position-relative rounded" @click="openVideo">
        <video :src="item.video" style="max-width:200px;max-height:300px;"></video>
        <span class="iconfont text-white position-absolute" style="font-size: 35px;width:35px;height:35px;" :style="posterIconStyle">&#xe737;</span>
      </div>
      <!-- 时间|操作 -->
      <div class="flex align-center justify-between">
        <span class="text-light-muted">{{item.created_at|formatTime}}</span>
        <div class="p-1">
          <van-popover v-model="showPopover" theme="dark" trigger="click" placement="left" :actions="actions" @select="onSelect">
            <template #reference>
              <van-icon name="ellipsis" size="20" />
            </template>
          </van-popover>
        </div>
      </div>
      <!-- 点赞列表|评论列表 -->
      <div class="bg-light mt-1" v-if="item.likes || item.comments">
        <!-- 点赞 -->
        <div v-if="item.likes.length" class="border-bottom flex align-start" style="padding:5px">
          <van-icon name="like-o" size="16" color="#0056b3" />
          <div class="flex flex-1 flex-wrap">
            <span v-for="(s,si) in item.likes" :key="si" class="text-hover-primary ml-1">{{s.name}}</span>
          </div>
        </div>
        <!-- 评论 -->
        <div v-if="item.comments.length" class="flex align-start" style="padding:5px;">
          <van-icon name="comment-o" size="16" color="#0056b3" />
          <div class="flex flex-column flex-1 ml-1">
            <div class="flex" v-for="(c,ci) in item.comments" :key="ci">
              <span v-if="!c.reply" class="text-hover-primary">{{c.user.name}}：</span>
              <div v-else class="flex align-center">
                <span class="text-hover-primary">{{c.user.name}} </span>
                <span class="text-muted" style="margin:0 2px;">回复</span>
                <span class="text-hover-primary">{{c.reply.name}}：</span>
              </div>
              <span class="text-dark flex-1" @click.stop="$emit('reply',{
								item,
								index,
								reply:c.user
							})">{{c.content}}</span>
            </div>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
import freeAvater from '@/components/free-avater.vue';
import { ImagePreview } from 'vant'
export default {
  components: {
    freeAvater
  },
  props: {
    item: Object,
    index: Number
  },
  data() {
    return {
      // 默认封面的宽高
      poster: {
        w: 100,
        h: 100
      },
      showPopover: false,
      actions: [{ text: '赞', event: 'like' }, { text: '评论', event: 'comment' }],
    }
  },
  computed: {
    imgs() {
      return this.item.image ? this.item.image.split(',') : []
    },
    // 短视频封面图标位置
    posterIconStyle() {
      let w = this.poster.w / 2 - 35 / 2
      let h = this.poster.h / 2 - 45 / 2
      return `left:${w}px;top:${h}px;`
    }
  },
  mounted() {
    if (this.item.video) {
      let video = document.querySelector('video');
      //canplay 事件，视频达到可以播放时触发；
      video.addEventListener('canplay', () => {
        this.loadPoster(video.videoWidth, video.videoHeight)
      });
    }
  },
  methods: {
    onSelect({ event }) {
      this.$emit('action', { event, item: this.item, index: this.index })
    },
    // 查看大图
    prediv(src) {
      let index = this.imgs.findIndex(item => {
        return item.url === src
      })
      if (index <= 0) {
        index = 0
      }
      //预览图片
      ImagePreview({
        images: this.imgs,
        startPosition: index,
        closeable: true
      });
    },
    openVideo() {// 播放视频
      this.$emit('openVideo', this.item.video)
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
    }
  },
}
</script>

<style>
</style>
