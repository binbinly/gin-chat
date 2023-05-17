<template>
  <div style="z-index:9999;overflow:hidden;" v-if="status">
    <!-- 蒙版 -->
    <div v-if="mask" class="position-fixed top-0 left-0 right-0 bottom-0 z-index" :style="getMaskColor" @click="hide"></div>
    <!-- 弹出框内容 -->
    <div ref="popup" class="position-fixed z-index" :class="getBodyClass" :style="getBodyStyle">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    // 是否开启蒙版颜色
    maskColor: {
      type: Boolean,
      default: false
    },
    // 是否开启蒙版
    mask: {
      type: Boolean,
      default: true
    },
    // 是否居中
    center: {
      type: Boolean,
      default: false
    },
    // 是否处于底部
    bottom: {
      type: Boolean,
      default: false
    },
    // 弹出层内容宽度
    bodyWidth: {
      type: Number,
      default: 0
    },
    // 弹出层内容高度
    bodyHeight: {
      type: Number,
      default: 0
    },
    bodyBgColor: {
      type: String,
      default: "bg-dark"
    },
    transformOrigin: {
      type: String,
      default: "left top"
    },
    // tabbar高度
    tabbarHeight: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      status: false,
      x: -1,
      y: 1,
      maxX: 240,
      maxY: 500
    }
  },
  mounted() {
  },
  computed: {
    getMaskColor() {
      let i = this.maskColor ? 0.5 : 0
      return `background-color: rgba(0,0,0,${i});`
    },
    getBodyClass() {
      if (this.center) {
        return 'left-0 right-0 bottom-0 top-0 flex align-center justify-center'
      }
      let bottom = this.bottom ? 'left-0 right-0 bottom-0' : 'rounded border'
      return `${this.bodyBgColor} ${bottom}`
    },
    getBodyStyle() {
      let left = this.x > -1 ? `left:${this.x}px;` : ''
      let top = this.y > -1 ? `top:${this.y}px;` : ''
      return left + top
    }
  },
  methods: {
    show(x = -1, y = -1) {
      if (this.status) {
        return;
      }
      this.x = (x > this.maxX) ? this.maxX : x
      this.y = (y > this.maxY) ? this.maxY : y
      this.status = true
    },
    hide() {
      this.$emit('hide')
      this.status = false
    }
  }
}
</script>

<style scoped>
.z-index {
  z-index: 9999;
}
</style>
