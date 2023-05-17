<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="选择" left-arrow @click-left="onClickLeft" fixed placeholder>
      <template #right>
        <van-button type="primary" size="small" @click="submit">{{buttonText}}</van-button>
      </template>
    </van-nav-bar>

    <!-- 通讯录列表 -->
    <template v-if="type === 'see'">
      <van-cell v-for="(item,index) in typeList" :title="item.name" center @click="typeIndex = index">
        <template #right-icon>
          <van-checkbox :value="typeIndex === index ? true : false" checked-color="#08c060" />
        </template>
      </van-cell>
    </template>
    <template v-if="type !== 'see' || (type === 'see' && (typeIndex === 1 || typeIndex === 2)) ">
      <!-- 侧边导航条 -->
      <van-index-bar :sticky-offset-top="45">
        <template v-for="(item,index) in list">
          <van-index-anchor :index="item.title" />
          <van-cell v-for="(item2,index2) in item.list" :title="item2.name" center @click="selectItem(item2)">
            <template #icon>
              <van-image class="pr-1" round width="35" height="35" :src="item2.avatar|formatAvatar" />
            </template>
            <template #right-icon>
              <van-checkbox v-model="item2.checked" checked-color="#08c060" ref="checkboxes" />
            </template>
          </van-cell>
        </template>
      </van-index-bar>
    </template>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import auth from '@/mixin/auth.js';
import { Toast } from 'vant';
import event from '@/utils/event.js';
import { groupCreate, groupInvite } from '@/api/group.js'
import $const from '@/const/index.js'
export default {
  mixins: [auth],
  data() {
    return {
      typeIndex: 0,
      typeList: [{
        name: "公开",
        key: "all"
      }, {
        name: "谁可以看",
        key: "only"
      }, {
        name: "不给谁看",
        key: "except"
      }, {
        name: "私密",
        key: "none"
      }],
      selectList: [],
      type: "",
      limit: 9,
      id: 0
    }
  },
  //keep-alive的生命周期 初次进入时：created > mounted > activated
  //只运行一次放此处
  mounted() {
    this.$store.dispatch('contactList')
  },
  //再次进入：只会触发 activated
  activated() {
    this.type = this.$route.query.type
    this.limit = parseInt(this.$route.query.limit)
    this.id = parseInt(this.$route.query.id)
    if (this.id) {
      if (this.type === 'inviteGroup') {
        this.limit = 1
      }
    }
  },
  computed: {
    ...mapState({
      list: state => {
        return state.user.mailList.map(item => {
          item.list = item.list.map(res => {
            return { ...res, checked: false }
          })
          return item
        })
      }
    }),
    buttonText() {
      let text = '发送'
      if (this.type === 'createGroup') {
        text = '创建群组'
      }
      return text + ' (' + this.selectCount + ')'
    },
    // 选中数量
    selectCount() {
      return this.selectList.length
    },
  },
  methods: {
    // 选中/取消选中
    selectItem(item) {
      if (!item.checked && this.selectCount === this.limit) {
        // 选中|限制选中数量
        return Toast('最多选中 ' + this.limit + ' 个')
      }
      item.checked = !item.checked
      if (item.checked) { // 选中
        this.selectList.push(item)
      } else { // 取消选中
        let index = this.selectList.findIndex(v => v === item)
        if (index > -1) {
          this.selectList.splice(index, 1)
        }
      }
    },
    submit() {
      if (this.type !== 'see' && this.selectCount === 0) {
        return Toast('请先选择')
      }
      switch (this.type) {
        case 'createGroup': // 创建群组
          groupCreate({
            ids: this.selectList.map(item => item.id)
          }).then(() => {
            Toast.success('创建群聊成功')
            this.$router.back()
          })
          break;
        case 'sendCard':
          let item = this.selectList[0]
          event.$emit('sendItem', {
            sendType: "card",
            content: item.name,
            type: $const.TYPE_CARD,
            options: {
              avatar: item.avatar,
              id: item.id
            }
          })
          this.$router.back()
          break;
        case 'remind':
          event.$emit('sendResult', {
            type: "remind",
            data: this.selectList
          })
          this.$router.back()
          break;
        case 'see':
          let k = this.typeList[this.typeIndex].key
          if (k !== 'all' && k !== 'none' && !this.selectCount) {
            return Toast('请先选择')
          }
          event.$emit('sendResult', {
            type: "see",
            data: {
              k,
              v: this.selectList
            }
          })
          this.$router.back()
          break;
        case 'inviteGroup':
          groupInvite({
            id: this.id,
            user_id: this.selectList[0].id
          }).then(() => {
            Toast.success('邀请成功')
            event.$emit('refreshGroupInfo')
            this.$router.back()
          })
          break;
      }
    }
  }
}
</script>

<style>
</style>
