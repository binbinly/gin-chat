<template>
  <div class="app-container">
    <div class="layout-content">
      <keep-alive v-if="$route.meta.keepAlive">
        <router-view></router-view>
      </keep-alive>
      <router-view v-else></router-view>
    </div>
    <div class="layout-footer" v-if="$route.meta.tabbar">
      <TabBar :data="tabbars" @change="handleChange" />
    </div>
  </div>
</template>

<script>
import TabBar from '@/components/TabBar'
export default {
  name: 'AppLayout',
  data() {
    return {
      tabbars: [
        {
          title: '首页',
          to: {
            name: 'Home'
          },
          icon: require('@/assets/tabbar/index.png'),
          iconS: require('@/assets/tabbar/index-select.png')
        },
        {
          title: '通讯录',
          to: {
            name: 'Contact'
          },
          icon: require('@/assets/tabbar/contact.png'),
          iconS: require('@/assets/tabbar/contact-select.png')
        },
        {
          title: '发现',
          to: {
            name: 'Find'
          },
          icon: require('@/assets/tabbar/find.png'),
          iconS: require('@/assets/tabbar/find-select.png')
        },
        {
          title: '我的',
          to: {
            name: 'My'
          },
          icon: require('@/assets/tabbar/my.png'),
          iconS: require('@/assets/tabbar/my-select.png')
        }
      ]
    }
  },
  components: {
    TabBar
  },
  created() {
    // 初始化登录状态
    this.$store.dispatch('initLogin')
  },
  methods: {
    handleChange(v) {
      console.log('tab value:', v)
    }
  }
}
</script>
