<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar left-text="聊天信息" fixed placeholder left-arrow @click-left="onClickLeft" />

    <div class="flex flex-wrap py-1 bg-white">
      <div v-if="isUser" class="flex flex-column align-center justify-center mb-1 ml-1">
        <van-image :src="detail.avatar|formatAvatar" round width="50" height="50"></van-image>
        <span class="text-muted" style="margin-top:5px;">{{detail.name}}</span>
      </div>

      <div v-else class="flex flex-column align-center justify-center mb-1 ml-1" v-for="(item,index) in list">
        <van-image :src="item.avatar|formatAvatar" round width="50" height="50"></van-image>
        <span class="text-muted" style="margin-top:5px;">{{item.name}}</span>
      </div>

      <div class="mb-1 ml-1">
        <div class="border" @click="addUser">
          <van-icon class="text-light-muted p-1" name="plus" size="30" />
        </div>
      </div>
      <div class="mb-1 ml-1" v-if="isGroup">
        <div class="border" @click="deleteUser">
          <van-icon class="text-light-muted p-1" name="minus" size="30" />
        </div>
      </div>

    </div>
    <div v-if="isGroup">
      <van-divider />
      <van-cell title="群聊名称" is-link @click="show = true;editType='name'">
        <template #default>
          <span class="text-inline">{{detail.name}}</span>
        </template>
      </van-cell>
      <van-cell title="群二维码" is-link @click="creatQrCode">
        <template #default>
          <span class="iconfont">&#xe647;</span>
        </template>
      </van-cell>
      <van-cell title="群公告" :value="detail.remark" is-link @click="show=true;editType='remark'" />
    </div>
    <van-divider />
    <van-cell title="查找聊天记录" is-link @click="openHistory" />
    <van-divider />
    <van-cell center title="消息免打扰">
      <template #right-icon>
        <van-switch v-model="setting.no_remind" size="24" active-color="#08c060" @change="updateChatItem(setting.no_remind, 'no_remind')" />
      </template>
    </van-cell>
    <van-cell center title="置顶聊天">
      <template #right-icon>
        <van-switch v-model="setting.is_top" size="24" active-color="#08c060" @change="updateChatItem(setting.is_top, 'is_top')" />
      </template>
    </van-cell>
    <van-cell center title="强提醒">
      <template #right-icon>
        <van-switch v-model="setting.is_remind" size="24" active-color="#08c060" @change="updateChatItem(setting.is_remind, 'is_remind')" />
      </template>
    </van-cell>

    <div v-if="isGroup">
      <van-divider />
      <van-cell title="我在本群的昵称" is-link :value="nickname" @click="show=true;editType='nickname'" />
      <van-cell center title="显示群成员昵称">
        <template #right-icon>
          <van-switch v-model="setting.show_name" size="24" active-color="#08c060" @change="updateChatItem(setting.show_name, 'show_name')" />
        </template>
      </van-cell>
    </div>
    <van-divider />
    <van-cell title="清空聊天记录" is-link @click="clear" />
    <van-divider />
    <van-cell title="投诉" is-link @click="openReport" />
    <van-divider />
    <van-button v-if="isGroup" block icon="delete" type="danger" @click="quit">删除并退出</van-button>
    <div style="height: 50px;"></div>

    <!-- 遮罩 修改昵称-->
    <van-overlay :show="show" @click="show = false">
      <div class="wrapper-bottom">
        <div class="bg-white rounded w-100" @click.stop>
          <van-form @submit="onSubmit">
            <van-field v-if="editType == 'name'" v-model="detail.name" label="群聊名称" clearable placeholder="请输入群聊名称" :rules="[{ required: true }]" />
            <van-field v-else-if="editType == 'nickname'" v-model="nickname" label="群昵称" clearable placeholder="请输入群昵称"
                       :rules="[{ required: true }]" />
            <van-field v-else-if="editType == 'remark'" v-model="detail.remark" type="textarea" clearable rows="2" maxlength="240" label="群公告"
                       placeholder="请输入群公告" :rules="[{ required: true }]" />
            <div style="margin: 16px;">
              <van-button round block type="primary" native-type="submit">提交</van-button>
            </div>
          </van-form>
        </div>
      </div>
    </van-overlay>
    <!-- 遮罩 二维码-->
    <van-overlay v-if="isGroup" :show="showCode" @click="showCode = false">
      <div class="wrapper">
        <div class="bg-white rounded p-4">
          <div class="flex align-center mb-2 justify-center">
            <van-image :src="detail.avatar | formatAvatar" class="avatar" />
            <div class="pl-1 flex flex-column">
              <span>{{detail.name}}</span>
            </div>
          </div>
          <div class="flex flex-column align-center justify-center">
            <div class="qrcode mb-2" ref="groupUrl"></div>
            <span class="text-light-muted">扫一扫上边二维码图案，加我的仿微信</span>
          </div>
        </div>
      </div>
    </van-overlay>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import auth from '@/mixin/auth.js';
import QRCode from 'qrcodejs2'
import { groupInfo, groupEdit, groupNickname, groupQuit } from '@/api/group.js'
import { Toast, Dialog } from 'vant';
import event from '@/utils/event.js'
import $const from '@/const/index.js'
export default {
  mixins: [auth],
  data() {
    return {
      show: false,
      showCode: false,
      editType: '',
      qrcode: false,
      list: [],
      nickname: "", // 我在本群的昵称
      id: 0,// 接收人/群 id
      chat_type: 1, // 接收类型 1 单聊 2 群聊
      detail: {
        id: 0, // 接收人/群 id
        avatar: '', // 接收人/群 头像
        name: '', // 接收人/群 昵称
        user_id: 0, // 群管理员id
        remark: "", // 群公告
        invite_confirm: 0, // 邀请确认
      },
      setting: {
        is_top: false, // 是否置顶
        show_name: false, // 是否显示昵称
        no_remind: false, // 消息免打扰
        is_remind: false, // 是否开启强提醒
      }
    }
  },
  computed: {
    ...mapState({
      chat: state => state.user.chat,
      user: state => state.user.user
    }),
    isUser() {
      return this.chat_type === $const.CHAT_TYPE_USER
    },
    isGroup() {
      return this.chat_type === $const.CHAT_TYPE_GROUP
    }
  },
  activated() {
    const { id, chat_type } = this.$route.query
    if (!id || !chat_type) {
      return this.backToast()
    }
    if (this.id != id || this.chat_type != chat_type) { // 参数变化，重新初始化
      this.chat_type = parseInt(chat_type)
      this.id = parseInt(id)
      this.initData()
    }
  },
  mounted() {
    event.$on('refreshGroupInfo', this.initData)
  },
  destroyed() {
    event.$off('refreshGroupInfo', this.initData)
  },
  methods: {
    creatQrCode() {
      this.showCode = true
      if (this.qrcode) {
        return;
      }
      new QRCode(this.$refs.groupUrl, {
        text: JSON.stringify({
          id: this.id,
          type: 2
        }),
        width: 240,
        height: 240,
        colorDark: '#000000',
        colorLight: '#ffffff'
      })
      this.qrcode = true
    },
    initData() {
      // 获取当前会话详细资料
      const detail = this.chat.getChatListItem(this.id, this.chat_type)
      console.log('detail', detail)
      if (!detail) {
        return this.backToast()
      }
      this.setting.is_top = detail.is_top
      this.setting.show_name = detail.show_name
      this.setting.no_remind = detail.no_remind
      this.setting.is_remind = detail.is_remind
      if (this.isGroup) {
        groupInfo(this.id).then(res => {
          this.list = res.users.slice(0, 4)
          this.detail = res.info
          this.nickname = res.nickname
        })
      }
    },
    updateChatItem(value, key) {
      this.chat.updateChatItemKey(this.id, this.chat_type, key, value)
    },
    addUser() {
      this.$router.push({
        path: '/contact_list', query: {
          id: this.id,
          type: this.isUser ? 'createGroup' : 'inviteGroup'
        }
      })
    },
    openReport() {
      this.$router.push({
        path: '/user_report', query: {
          user_id: this.id,
          type: this.chat_type
        }
      })
    },
    quit() {
      Dialog.confirm({
        message: '是否要删除或退出该群聊？',
      })
        .then(() => {
          groupQuit({
            id: this.id
          }).then(() => {
            Toast.success('退出成功')
            this.$router.replace({ path: '/home' })
          })
        })
    },
    clear() {
      Dialog.confirm({
        message: '是否要清空聊天记录？',
      })
        .then(() => {
          this.chat.clearChatDetail(this.id, this.chat_type)
          Toast.success('清空成功')
          event.$emit('updateHistory')
        })
    },
    openHistory() {
      this.$router.push({
        path: '/chat_history', query: {
          chat_type: this.chat_type,
          id: this.id
        }
      })
    },
    deleteUser() {
      this.$router.push({ path: 'group_user', query: { id: this.id } })
    },
    onSubmit() {
      if (this.editType == 'nickname') {//修改本群昵称
        groupNickname({
          id: this.id,
          nickname: this.nickname
        }).then(() => {
          Toast.success('修改成功')
        })
      } else if (this.editType == 'remark') {
        groupEdit({
          id: this.id,
          remark: this.detail.remark
        }).then(() => {
          Toast.success('修改成功')
        })
      } else {//修改群信息
        groupEdit({
          id: this.id,
          name: this.detail.name,
        }).then(() => {
          Toast.success('修改成功')
        })
      }
      this.editType = ''
      this.show = false
    }
  }
}
</script>

<style>
</style>
