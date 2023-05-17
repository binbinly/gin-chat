<template>
  <div>
    <div class="fixed-top" :style="navBarStyle">
      <!-- 状态栏 -->
      <div :style="'height:'+statusBarHeight+'px'"></div>
      <!-- 导航 -->
      <div class="w-100 flex align-center justify-between" style="height: 45px;">
        <!-- 左边 -->
        <div class="flex align-center">
          <!-- 返回按钮 -->
          <div class="flex align-center justify-center" hover-class="bg-hover-light" @click="back" style="height: 40px;width: 40px;">
            <van-icon name="arrow-left" class="font-lg" :color="buttonColor" />
          </div>
          <!-- 标题 -->
          <span v-if="title" class="font">{{title}}</span>
        </div>
        <!-- 右边 -->
        <div class="flex align-center">
          <div class="flex align-center justify-center" hover-class="bg-hover-light" style="height: 40px;width: 40px;" @click="$emit('clickRight')">
            <van-icon name="photograph" class="font-lg" :color="buttonColor" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: [String, Boolean],
      default: false
    },
    scrollTop: {
      type: [Number, String],
      default: 0
    }
  },
  data() {
    return {
      statusBarHeight: 0,
      navBarHeight: 0
    }
  },
  mounted() {
    this.navBarHeight = this.statusBarHeight + 90
  },
  computed: {
    // 变化 0 - 1
    changeNumber() {
      let start = 200
      let end = 280
      let H = end - start
      let num = 0
      if (this.scrollTop > start) {
        num = (this.scrollTop - start) / H
      }
      return num > 1 ? 1 : num
    },
    navBarStyle() {
      return `background-color: rgba(255,255,255,${this.changeNumber});`
    },
    buttonColor() {
      if (this.changeNumber > 0) {
        return '#000000';
      }
      return '#FFFFFF';
    }
  },
  methods: {
    // 返回
    back() {
      this.$router.back()
    }
  },
}
</script>

<style>
</style>
