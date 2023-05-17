<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar title="设置备注和标签" left-arrow @click-left="onClickLeft">
      <template #right>
        <van-button type="primary" size="small" @click="onSave">保存</van-button>
      </template>
    </van-nav-bar>

    <van-cell-group title="备注名">
      <van-field v-model="nickname" name="nickname" placeholder="请输入备注名" />
    </van-cell-group>
    <van-cell-group title="标签">
      <van-button icon="plus" size="small" class="m-1" type="primary" @click="show = true">自定义标签</van-button>
      <template v-for="(item,index) in tagList">
        <van-tag class="m-1" style="padding:5px;" closeable size="large" type="primary" @close="removeTag(item)">
          {{item}}</van-tag>
      </template>
    </van-cell-group>
    <van-cell-group title="可选常用标签">
      <van-tag class="m-1" type="primary" size="large" v-for="(item,index) in allTagList" :type="item.type" @click="addTag(item.name)">
        {{item.name}}</van-tag>
    </van-cell-group>

    <!-- 遮罩 添加自定义标签-->
    <van-overlay :show="show" @click="show = false">
      <div class="wrapper" @click.stop>
        <van-form @submit="onSubmitTag">
          <van-field v-model="tag" center clearable label="自定义标签" placeholder="请输入标签名" :rules="[{ required: true}]">
            <template #button>
              <van-button size="small" type="primary" native-type="submit">添加</van-button>
            </template>
          </van-field>
        </van-form>
      </div>
    </van-overlay>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js';
import { Toast } from 'vant';
import { friendRemarkTag } from '@/api/friend.js'
export default {
  mixins: [auth],
  data() {
    return {
      show: false,
      id: 0,
      nickname: "",
      tag: '',
      tagList: [],
      allTagList: ['家人', '亲戚', '朋友', '好友', '运动', '吃喝', '看书', '睡觉', '玩手机']
    }
  },
  mounted() {
    const query = this.$route.query
    if (!query) {
      return this.backToast()
    }
    this.id = parseInt(query.user_id)
    this.nickname = query.nickname
    this.tagList = query.tags == '' ? [] : query.tags.split(',')
    const t = ['primary', 'success', 'danger', 'warning']
    this.allTagList = this.allTagList.map(item => {
      const index = Math.floor((Math.random() * 4));
      return { name: item, type: t[index] }
    })
  },
  methods: {
    addTag(item) {
      if (this.tagList.indexOf(item) !== -1) {
        return Toast('标签已存在')
      }
      this.tagList.push(item)
    },
    removeTag(item) {
      this.tagList = this.tagList.filter(name => name != item)
    },
    onSubmitTag() {
      this.addTag(this.tag)
      this.tag = ''
      this.show = false
    },
    // 完成
    onSave() {
      friendRemarkTag({
        user_id: this.id,
        nickname: this.nickname,
        tags: this.tagList
      }).then(res => {
        Toast.success('保存成功')
        this.$router.back()
      })
    }
  }
}
</script>

<style>
</style>
