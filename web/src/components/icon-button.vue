<template>
  <div class="flex align-center justify-center" hover-class="bg-hover-light" @click="$emit('click')" style="height: 45px;width: 45px;">
    <span class="iconfont font-md">{{icon}}</span>
  </div>
</template>

<script>
export default {
  props: {
    icon: {
      type: String,
      default: ''
    },
  },
}
</script>

<style>
</style>
