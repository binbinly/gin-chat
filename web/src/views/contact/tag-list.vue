<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="标签列表" fixed placeholder left-arrow @click-left="onClickLeft">
    </van-nav-bar>

    <van-cell is-link v-for="(item,index) in list" :title="item.name" @click="handle(item)"></van-cell>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { tagList } from '@/api/tag.js'
import { mapState } from 'vuex'
export default {
  mixins: [auth],
  data() {
    return {
      list: []
    }
  },
  computed: {

  },
  created() {
    this.getData()
  },
  methods: {
    getData() {
      tagList().then(res => {
        this.list = res
      })
    },
    handle(item) {
      this.$router.push({ path: 'tag_read', query: { id: item.id } })
    }
  }
}
</script>

<style>
</style>
