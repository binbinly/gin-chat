<template>
  <div>
    <van-tabbar fixed route v-model="active" @change="handleChange" active-color="#08C060" inactive-color="#333333">
      <van-tabbar-item v-for="(item, index) in data" :to="item.to" :key="index" :badge="item.badge">
        {{ item.title }}
        <template #icon="props">
          <img :src="props.active ? item.iconS : item.icon" />
        </template>
      </van-tabbar-item>
    </van-tabbar>
  </div>
</template>
<script>
import event from '@/utils/event.js'
export default {
  name: 'TabBar',
  props: {
    defaultActive: {
      type: Number,
      default: 0
    },
    data: {
      type: Array,
      default: () => {
        return []
      }
    }
  },
  data() {
    return {
      active: this.defaultActive
    }
  },
  mounted() {
    // 开启监听聊天消息变化
    event.$on('tabBarBadge', (res) => {
      console.log('tabBarBadge', res)
      this.data[res.index].badge = res.text
    })
  },
  destroyed() {
    event.$off('tabBarBadge', () => { })
  },
  methods: {
    handleChange(value) {
      this.$emit('change', value)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
