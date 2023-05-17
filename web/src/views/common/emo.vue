<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar :left-text="cat" left-arrow @click-left="onClickLeft" fixed placeholder>
      <template #right>
        <span v-if="isAdd">已添加</span>
        <van-button v-else type="primary" size="small" @click="submit">添加表情</van-button>
      </template>
    </van-nav-bar>

    <div class="content-container">
      <vue-waterfall-easy :imgsArr="imgsArr" :loadingTimeOut="1"></vue-waterfall-easy>
    </div>
  </div>
</template>

<script>
import vueWaterfallEasy from 'vue-waterfall-easy'
import auth from '@/mixin/auth.js';
import { mapState } from 'vuex'
import { emoticon } from '@/api/common.js'
import event from '@/utils/event.js'
export default {
  components: {
    vueWaterfallEasy
  },
  mixins: [auth],
  data() {
    return {
      cat: '',
      isAdd: false, //是否已添加
      imgsArr: []
    }
  },
  computed: {
    ...mapState({
      emoCat: state => state.user.emoCat
    }),
  },
  activated() {
    this.cat = this.$route.query.cat
    const index = this.emoCat.indexOf(this.cat)
    this.isAdd = index > -1
    this.initData()
  },
  methods: {
    initData() {
      emoticon(this.cat).then(res => {
        this.imgsArr = res.map(item => {
          return { src: item.url, info: item.name }
        })
      })
    },
    submit() {
      this.$store.commit('addEmo', this.cat)
      this.isAdd = true
      event.$emit('onEmoticon', this.cat)
    }
  }
}
</script>

<style>
</style>
