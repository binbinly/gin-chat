<template>
  <div :class="item.is_top ? 'bg-light' : 'bg-white'">
    <van-cell center @click="onClick" v-longpress="long">
      <template #icon>
        <van-badge :content="item.unread || ''">
          <van-image round width="40" height="40" :src="getAvatar" />
        </van-badge>
      </template>
      <template #title>
        <div class="flex flex-column pl-1">
          <span class="font">{{item.name}}</span>
          <span class="text-inline font-small text-light-muted" style="width:220px;">{{item.content}}</span>
        </div>
      </template>
      <template #default>
        <span class="text-light-muted font-small">{{item.t|formatTime}}</span>
      </template>
    </van-cell>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  props: {
    item: Object,
    index: Number
  },
  computed: {
    ...mapState({
      chat: state => state.user.chat
    }),
    getAvatar() {
      if (this.item.avatar) {
        return this.item.avatar
      } else if (this.item.chat_type == 2) { // 群默认头像
        return require('@/assets/images/group.jpg')
      } else {
        return require('@/assets/images/userpic.png')
      }
    }
  },
  methods: {
    onClick() {
      this.$router.push({ path: '/chat', query: { id: this.item.id, type: this.item.chat_type } })
      this.chat.readChatItem(this.item.id, this.item.chat_type)
    },
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

<style>
</style>
