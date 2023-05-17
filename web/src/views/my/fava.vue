<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar title="我的收藏" fixed placeholder left-arrow @click-left="onClickLeft" />

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list v-model="loading" :finished="finished" finished-text="没有更多了" @load="onLoad">
        <div class="pt-1 px-1" v-for="(item,index) in list" @click="send(item)">
          <van-swipe-cell>
            <div class="bg-white rounded p-1">
              <span v-if="isText(item.type)">{{item.content}}</span>
              <van-image v-else-if="isEmoticon(item.type) || isImage(item.type)" style="max-width: 200px;max-height:300px" :src="item.content" />
              <video v-else-if="isVideo(item.type)" :src="item.content" controls style="max-width: 200px;max-height:300px"></video>
            </div>
            <template #right>
              <van-button square type="danger" text="删除" style="height: 100%;" @click="onDelete(item.id, index)" />
            </template>
          </van-swipe-cell>
        </div>
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js'
import event from '@/utils/event.js'
import { collectList, collectDestroy } from '@/api/collect.js'
import { Dialog, Toast } from 'vant'
import $const from '@/const/index.js'
export default {
  mixins: [auth],
  data() {
    return {
      loading: false,
      finished: false,
      refreshing: false,
      page: 1,
      list: [],
      type: ""
    }
  },
  computed: {
    isText() {//是否文本消息
      return function(type) {
        return type === $const.TYPE_TEXT
      }
    },
    isImage() {//是否图片消息
      return function(type) {
        return type === $const.TYPE_IMAGE
      }
    },
    isVideo() {//是否视频消息
      return function(type) {
        return type === $const.TYPE_VIDEO
      }
    },
    isAudio() {//是否音频消息
      return function(type) {
        return type === $const.TYPE_AUDIO
      }
    },
    isEmoticon() {//是否表情消息
      return function(type) {
        return type === $const.TYPE_EMOTICON
      }
    },
    isCard() {//是否名片消息
      return function(type) {
        return type === $const.TYPE_CARD
      }
    }
  },
  activated() {
    this.type = this.$route.query.type || ''
  },
  methods: {
    onLoad() {
      this.getData()
    },
    onRefresh() {
      this.page = 1
      // 清空列表数据
      this.finished = false;
      // 重新加载数据
      // 将 loading 设置为 true，表示处于加载状态
      this.loading = true;
      this.onLoad();
    },
    send(item) {
      if (this.type !== 'send') {
        return
      }
      event.$emit('sendItem', {
        sendType: "fava",
        ...item
      })
      this.$router.back()
    },
    onDelete(id, index) {
      Dialog.confirm({
        message: '确定删除吗？',
      }).then(() => {
        collectDestroy(id).then(res => {
          this.list.splice(index, 1)
          Toast.success('删除成功')
        })
      });
    },
    getData() {
      collectList(this.page).then(res => {
        if (this.refreshing) {
          this.list = [];
          this.refreshing = false;
        }
        const list = res.map(item => {
          item.options = JSON.parse(item.options)
          return item
        })
        this.list = [...this.list, ...list]
        this.loading = false;
        if (res.length < 10) {
          this.finished = true;
        } else {
          this.page++
        }
      })
    }
  }
}
</script>

<style>
</style>
