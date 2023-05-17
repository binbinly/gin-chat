<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar title="表情包" fixed placeholder left-arrow @click-left="onClickLeft" />

    <van-cell v-for="(item,index) in list" :title="item.category" center is-link @click="handle(item)">
      <template #icon>
        <van-image class="pr-1" lazy-load width="60" height="60" fit="cover" :src="item.url" />
      </template>
      <template #right-icon>
        <span v-if="item.isAdd">已添加</span>
        <van-button v-else type="primary" size="small" v-on:click.stop="submit(item)">添加</van-button>
      </template>
    </van-cell>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { mapState } from 'vuex'
import { emoticonCat } from '@/api/common.js'
import event from '@/utils/event.js'
export default {
  mixins: [auth],
  data() {
    return {
      list: []
    }
  },
  computed: {
    ...mapState({
      emoCat: state => state.user.emoCat
    }),
  },
  mounted() {
    this.initData()
    event.$on('onEmoticon', this.onEmoticon)
  },
  destroy() {
    event.$off('onEmoticon', this.onEmoticon)
  },
  methods: {
    onEmoticon(cat) {
      this.list.forEach(item => {
        if (item.category == cat) {
          item.isAdd = true
        }
      })
    },
    initData() {
      emoticonCat().then(res => {
        this.list = res.map(item => {
          let index = -1
          if (this.emoCat.length > 0) {
            index = this.emoCat.indexOf(item.category)
          }
          return { category: item.category, isAdd: index > -1, url: item.url }
        })
      })
    },
    handle(item) {
      console.log('handle')
      this.$router.push({
        path: '/emoticon', query: {
          cat: item.category,
        }
      })
    },
    submit(item) {
      this.$store.commit('addEmo', item.category)
      item.isAdd = true
      event.$emit('onEmoticon', item.category)
    }
  }
}
</script>

<style>
</style>
