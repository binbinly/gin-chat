<template>
  <van-image round :width="size" :height="size" @click="openUser" :src="src|formatAvatar" :class="classStyle" />
</template>

<script>
export default {
  props: {
    uid: Number,
    size: {
      type: [String, Number],
      default: 40
    },
    src: {
      type: String,
      default: ''
    },
    classStyle: {
      type: String,
      default: ''
    }
  },
  methods: {
    openUser() {
      this.$router.push({ path: '/user_base', query: { id: this.uid } })
    },
  },
}
</script>

<style>
</style>
