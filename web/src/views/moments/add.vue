<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="记忆心情" left-arrow @click-left="onClickLeft">
      <template #right>
        <van-button type="primary" size="small" @click="submit">发表</van-button>
      </template>
    </van-nav-bar>
    <!-- 文字 -->
    <van-field v-model="content" rows="4" autosize type="textarea" maxlength="500" placeholder="这一刻的想法" show-word-limit
               :rules="[{ required: true }]" />
    <!-- 多图上传 -->
    <div class="bg-white p-1" v-if="type === 'image'">
      <van-uploader v-model="imageList" name="image" preview-full-image multiple :after-read="afterRead" accept="image/*" :max-count="9"
                    ref="uploadImage" />
    </div>
    <!-- 视频 -->
    <video v-if="type === 'video' && video && video" :src="video" controls style="max-width:360px;"></video>
    <van-uploader v-if="type === 'video'" name="video" capture="camera" :prediv-image="false" :after-read="afterRead" accept=".mp4">
      <div v-if="video" class="my-1 flex align-center justify-center bg-light" hover-class="bg-hover-light" style="height: 50px;width:360px;">
        <span class="font-sm text-muted">点击切换视频</span>
      </div>
      <div v-else class="flex align-center justify-center bg-light px-1" style="height: 175px;width:360px;">
        <span class="text-muted" style="font-size: 50px;">+</span>
      </div>
    </van-uploader>

    <van-cell title="所在位置" value="位置" is-link @click="draw" />
    <van-cell title="提醒谁看" is-link center to="/contact_list?type=remind">
      <template #default>
        <div class="ml-1">
          <van-image style="margin-left:5px;" v-for="(item,index) in remindList" round width="25" height="25" :src="item.avatar|formatAvatar" />
        </div>
      </template>
    </van-cell>
    <van-cell title="谁可以看" :value="seeText" is-link center to="/contact_list?type=see" />
  </div>
</template>

<script>
import auth from '@/mixin/auth.js'
import event from '@/utils/event.js'
import { momentCreate } from '@/api/moment.js'
import { uploadFile } from '@/api/common.js'
import { Toast } from 'vant'
export default {
  mixins: [auth],
  data() {
    return {
      content: "",
      imageList: [],
      type: "image",
      video: "",
      remindList: [],
      seeType: 'all', //可见类型
      seeList: []  //可见用户id列表
    }
  },
  activated() {
    const t = this.$route.query.type
    if (t != this.type) {
      this.init()
    }
    this.type = t
  },
  mounted() {
    event.$on('sendResult', this.sendResult)
  },
  destroyed() {
    event.$off('sendResult', this.sendResult)
  },
  computed: {
    seeText() {
      const type = { all: "公开", none: "私密", only: "谁可以看", except: "不给谁看" }
      if (this.seeType === 'all' || this.seeType === 'none') {
        return type[this.seeType]
      }
      let names = (this.seeList.map(item => item.name)).join(',')
      return type[this.seeType] + ':' + names
    }
  },
  methods: {
    init() {
      this.content = ""
      this.imageList = []
      this.video = ''
      this.remindList = []
      this.seeType = 'all' //可见类型
      this.seeList = []  //可见用户id列表
    },
    sendResult(e) {
      console.log('e', e)
      if (e.type === 'remind') {
        this.remindList = e.data
      } else if (e.type === 'see') {
        this.seeType = e.data.k
        this.seeList = e.data.v
      }
    },
    submit() {
      const typeList = { 'text': 1, 'image': 2, 'video': 3 }
      const seeTypeList = { 'all': 1, 'none': 2, 'only': 3, 'except': 4 }
      if (this.type === 'text' && this.content == '') {
        return Toast('内容不能为空')
      } else if (this.type === 'image' && this.imageList.length == 0) {
        return Toast('请选择图片')
      } else if (this.type === 'video' && this.video == '') {
        return Toast('请上传短视频')
      }
      let imgs = this.imageList.map(item => item.url)
      momentCreate({
        content: this.content,
        image: imgs.join(','),
        video: this.video,
        type: typeList[this.type],
        location: "",
        remind: this.remindList.map(item => item.id),
        see_type: seeTypeList[this.seeType],
        see: this.seeList.map(item => item.id)
      }).then(() => {
        this.init()
        Toast.success('发布成功')
        event.$emit('refreshMoment')
        this.$router.back()
      })
    },
    // 文件读取完成后的回调函数
    afterRead(files, { name }) {
      //上传文件
      if (files instanceof Array) {
        files.forEach(item => {
          this.upload(item, name)
        })
      } else {
        this.upload(files, name)
      }
    },
    upload(file, name) {
      file.status = 'uploading';
      file.message = '上传中...';
      uploadFile(file.file).then(url => {
        file.status = 'done'
        if (name == 'video') {
          this.video = url
        } else {
          file.url = url
        }
      }).catch(() => {
        file.status = 'failed';
        file.message = '上传失败';
      })
    }
  }
}
</script>

<style>
</style>
