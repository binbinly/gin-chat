<template>
  <div>
    <!-- 导航 -->
    <van-nav-bar fixed placeholder>
      <template #left v-if="title">
        <span class="font">{{getTitle}}</span>
      </template>
      <template #right>
        <van-icon class="mr-1" name="search" @click="search" size="24" />
        <!-- 扩展菜单 -->
        <van-popover v-model="showPopover" :offset="[-50, 20]" placement="bottom" theme="dark" trigger="click" :actions="menus" @select="onSelect">
          <template #reference>
            <van-icon name="add-o" size="22" />
          </template>
        </van-popover>
      </template>
    </van-nav-bar>
  </div>
</template>
<script>
import { Toast } from 'vant'
export default {
  props: {
    title: {
      type: [String, Boolean],
      default: false
    },
    unread: {
      type: [Number, String],
      default: 0
    },
  },
  data() {
    return {
      showPopover: false,
      menus: [
        {
          text: "发起群聊",
          className: "group",
          icon: "chat-o",
          path: '/contact_list',
          query: {
            type: 'createGroup'
          }
        },
        {
          text: "添加好友",
          pclassName: "add_user",
          icon: "friends-o",
          path: '/search',
          query: {}
        },
        {
          text: "扫一扫",
          className: "scan",
          icon: "scan"
        },
        {
          text: "收付款",
          className: "pay",
          icon: "gold-coin-o"
        },
        {
          text: "帮助与反馈",
          className: "about",
          icon: "question-o"
        }
      ],
    }
  },
  computed: {
    getTitle() {
      let unread = this.unread > 0 ? '（' + this.unread + '）' : ''
      return this.title + unread
    }
  },
  methods: {
    search() {
      this.$router.push({ path: '/search' })
    },
    onSelect(action) {
      if (action.path) {
        this.$router.push({ path: action.path, query: action.query })
      } else {
        Toast('待开发')
      }
    }
  }
}
</script>

<style>
.van-popover[data-popper-placement='bottom'] .van-popover__arrow {
  left: 85%;
}
.van-popover__action {
  width: 4rem;
}
.van-nav-bar {
  line-height: unset;
}
</style>
