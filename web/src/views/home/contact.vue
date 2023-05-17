<template>
  <div>
    <!-- 导航栏 -->
    <nav-bar title="通讯录"></nav-bar>

    <!-- 通讯录列表 -->
    <div>
      <van-cell v-for="(item,index) in topList" is-link center :title="item.title" :to="item.path">
        <template #icon>
          <van-image fit="cover" :src="item.icon" class="pr-1" style="width: 35px;height: 35px;" />
        </template>
        <template #default v-if="item.id === 'friend' && applyCount > 0">
          <van-badge :content="applyCount" />
        </template>
      </van-cell>
    </div>
    <!-- 侧边导航条 -->
    <van-index-bar>
      <template v-for="(item,index) in list">
        <van-index-anchor :index="item.title" />
        <van-cell v-for="(item2,index2) in item.list" :title="item2.name" center @click="openUser(item2.id)">
          <template #icon>
            <van-image class="pr-1" round width="35" height="35" :src="item2.avatar|formatAvatar" />
          </template>
        </van-cell>
      </template>
    </van-index-bar>
  </div>
</template>

<script>
import NavBar from "@/components/nav-bar.vue"
import auth from '@/mixin/auth.js';
import { mapState } from 'vuex'
export default {
  mixins: [auth],
  components: {
    NavBar
  },
  data() {
    return {
      topList: [
        {
          id: "friend",
          title: "新的朋友",
          icon: require("@/assets/images/mail/friend.png"),
          path: "/apply_list"
        },
        {
          id: "group",
          title: "群聊",
          icon: require("@/assets/images/mail/group.png"),
          path: "/group_list"
        },
        {
          id: "tag",
          title: "标签",
          icon: require("@/assets/images/mail/tag.png"),
          path: "/tag_list"
        }
      ],

      top: 0,
      scrollHeight: 0,
      scrollInto: '',
      current: ''
    }
  },
  mounted() {
    this.$store.dispatch('contactList')
  },
  computed: {
    ...mapState({
      applyCount: state => state.user.apply.count,
      list: state => state.user.mailList
    }),
  },
  methods: {
    openUser(id) {
      this.$router.push({ path: '/user_base', query: { id } })
    }
  }
}
</script>

<style>
</style>
