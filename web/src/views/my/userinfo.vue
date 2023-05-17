<template>
  <div>
    <!-- 导航栏 -->
    <van-nav-bar title="个人资料" fixed placeholder left-arrow @click-left="onClickLeft" />

    <van-cell is-link center @click="openAvatar" title="头像">
      <!-- 使用 title 插槽来自定义标题 -->
      <template #default>
        <div class="flex align-center">
          <van-image class="avatar" style="margin-left:auto" round :src="user.avatar|formatAvatar" />
        </div>
      </template>
    </van-cell>
    <van-cell title="昵称" is-link :value="user.nickname" @click="show = true" />
    <van-cell title="账号" is-link :value="user.username" />
    <van-cell title="二维码名片" is-link @click="creatQrCode">
      <template #default>
        <span class="iconfont">&#xe647;</span>
      </template>
    </van-cell>

    <!-- 上传控件 -->
    <van-uploader style="display:none" v-model="avatarFile" prediv-full-image :after-read="afterRead" accept="image/*" :max-count="1"
                  ref="uploadImage" />

    <!-- 遮罩 修改昵称-->
    <van-overlay :show="show" @click="show = false">
      <div @click.stop>
        <van-form @submit="onSubmit">
          <van-field v-model="nickname" center clearable label="昵称" placeholder="请输入新昵称" :rules="[{ required: true}]">
            <template #button>
              <van-button size="small" type="primary" native-type="submit">修改</van-button>
            </template>
          </van-field>
        </van-form>
      </div>
    </van-overlay>

    <van-overlay :show="showCode" @click="showCode = false">
      <div class="wrapper">
        <div class="bg-white rounded p-4">
          <div class="flex align-center mb-2 justify-center">
            <van-image :src="user.avatar | formatAvatar" class="avatar" />
            <div class="pl-1 flex flex-column">
              <span>{{user.username ? user.username : user.nickname}}</span>
            </div>
          </div>
          <div class="flex flex-column align-center justify-center">
            <div class="qrcode mb-2" ref="qrCodeUrl"></div>
            <span class="text-light-muted">扫一扫上边二维码图案，加我的仿微信</span>
          </div>
        </div>
      </div>
    </van-overlay>
  </div>
</template>

<script>
import auth from '@/mixin/auth.js'
import QRCode from 'qrcodejs2'
import { userEdit } from '@/api/user.js'
import { mapState } from 'vuex'
import { Toast } from 'vant'
import { uploadFile } from '@/api/common.js'
import { dataURLtoBlob } from '@/utils/index.js'
export default {
  mixins: [auth],
  data() {
    return {
      qrcode: false,
      avatarFile: [],
      show: false,
      showCode: false,
      nickname: ''
    }
  },
  computed: {
    ...mapState({
      user: state => state.user.user
    }),
    confirmTitle() {
      return this.confirmType == 'username' ? '修改账号' : '修改昵称'
    },
    placeholder() {
      return this.confirmType == 'username' ? '输入账号' : '输入昵称'
    },
  },
  methods: {
    creatQrCode() {
      this.showCode = true
      if (this.qrcode) {
        return;
      }
      new QRCode(this.$refs.qrCodeUrl, {
        text: JSON.stringify({
          id: this.user.id,
          type: 1
        }),
        width: 240,
        height: 240,
        colorDark: '#000000',
        colorLight: '#ffffff',
        correctLevel: QRCode.CorrectLevel.H
      })
      this.qrcode = true
    },
    //	文件读取完成后的回调函数
    afterRead(file) {
      //上传文件
      file.status = 'uploading';
      file.message = '上传中...';
      uploadFile(file.file).then(url => {
        file.status = 'done'
        userEdit({
          avatar: url
        }).then(() => {
          this.user.avatar = url
          Toast.success('修改头像成功')
          this.$store.commit('updateUser', {
            k: 'avatar',
            v: url
          })
        })
      }).catch(() => {
        file.status = 'failed';
        file.message = '上传失败';
      })
    },
    // 修改头像
    openAvatar() {
      this.$refs.uploadImage.chooseFile()
    },
    onSubmit() {
      userEdit({
        nickname: this.nickname
      }).then(() => {
        this.$store.commit('updateUser', {
          k: 'nickname',
          v: this.nickname
        })
        this.show = false
        this.user.nickname = this.nickname
        Toast.success('修改成功')
      })
    }
  }
}
</script>

<style>
.wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}
</style>
