import { setStorage, getStorage, removeStorage } from '@/utils/index.js'
import { applyCount } from '@/api/apply.js'
import { friendList } from '@/api/friend.js'
import $config from '@/config/index.js'
import $router from '@/router/index.js'
import Chat from '@/utils/chat.js'
import ev from '@/utils/event.js'
import SortWord from 'sort-word'
export default {
  state: {
    user: false,
    token: false,

    apply: {
      list: [],
      count: 0
    },

    //我的好友
    mailList: [],

    //我的表情
    emoCat:[],

    chat: null,

    // 会话列表
    chatList: [],

    // 总未读数
    totalunread: 0,

    notice: {
      avatar: '',
      user_id: 0,
      num: 0
    }
  },
  mutations: {
    updateUser(state, { k, v }) {
      if (state.user) {
        console.log('k:', k, ';v:', v)
        console.log('start:', state.user)
        state.user[k] = v
        console.log('end:', state.user)
        setStorage('user', JSON.stringify(state.user))
      }
    },
    addEmo(state, name) {//添加表情包
      state.emoCat.push(name)
      setStorage('emoticon_cat', JSON.stringify(state.emoCat))
    },
    delEmo(state, name) {//删除表情包
      state.emoCat.splice(state.emoCat.findIndex(item => item === name), 1)
      setStorage('emoticon_cat', JSON.stringify(state.emoCat))
    }
  },
  actions: {
    // 登录后处理
    login({ state, dispatch }, data) {
      // 存到状态中
      data.user.name = data.user.nickname || data.user.username
      state.user = data.user
      state.token = data.token
      // 存储到本地存储中
      setStorage('token', data.token)
      setStorage('user', JSON.stringify(data.user))
      setStorage('user_id', data.user.id)
      // 获取好友申请列表
      dispatch('getApply')
      // 连接socket
      state.chat = new Chat({
        url: $config.socketUrl
      })
      // 获取会话列表
      dispatch('getChatList')
      // 初始化总未读数角标
      dispatch('updateBadge')
      // 获取朋友圈动态通知
      dispatch('getNotice')
    },
    // 退出登录处理
    logout({ state }) {
      // 清除登录状态
      state.user = false
      // 清除本地存储数据
      removeStorage('token')
      removeStorage('user')
      removeStorage('user_id')
      // 关闭socket连接
      if (state.chat) {
        state.chat.close()
        state.chat = null
      }
      // 跳转到登录页
      $router.replace({ path: '/login' })
      // 注销监听事件
      ev.$off('onUpdateChatList')
      ev.$off('momentNotice')
      ev.$off('totalunread')
    },
    // 初始化登录状态
    initLogin({ state, dispatch }) {
      // 拿到存储
      const user = getStorage('user')
      if (user) {
        // 初始化登录状态
        state.user = JSON.parse(user)
        // 连接socket
        state.chat = new Chat({
          url: $config.socketUrl
        })
        state.emoCat = getStorage('emoticon_cat') ? JSON.parse(getStorage('emoticon_cat')) : []
        // 获取会话列表
        dispatch('getChatList')
        // 获取离线信息
        // 获取好友申请列表
        dispatch('getApply')
        // 初始化总未读数角标
        dispatch('updateBadge')
        // 获取朋友圈动态通知
        dispatch('getNotice')
      }
    },
    // 获取好友申请列表
    getApply({ state, dispatch }, page = 1) {
      applyCount().then(res => {
        console.log('applyCount', res)
        state.apply.count = res
        // 更新通讯录角标提示
        dispatch('updateMailBadge')
      })
    },
    // 更新通讯录角标提示
    updateMailBadge({ state }) {
      let count = ''
      if (state.apply.count > 99) {
        count = '99+'
      } else if (state.apply.count > 0) {
        count = state.apply.count + ''
      }
      ev.$emit('tabBarBadge', {
        index: 1,
        text: count
      })
    },
    // 获取通讯录列表
    contactList({ state }) {
      friendList().then(res => {
        if (res.length > 0) {
          const list = new SortWord(res, 'name')
          state.mailList = list.newList
        }
        console.log('mailList', state.mailList)
      })
    },
    // 获取会话列表
    getChatList({ state }) {
      state.chatList = state.chat.getChatList()
      // 监听会话列表变化
      ev.$on('onUpdateChatList', list => {
        state.chatList = list
      })
    },
    // 获取朋友圈动态通知
    getNotice({ state }) {
      state.notice = state.chat.getNotice()
      if (state.notice.num > 0) {
        ev.$emit('tabBarBadge', {
          index: 2,
          text: state.notice.num > 99 ? '99+' : state.notice.num.toString()
        })
      } else {
        ev.$emit('tabBarBadge', {
          index: 2,
          text: ''
        })
      }
      ev.$on('momentNotice', notice => {
        state.notice = notice
      })
    },
    // 初始化总未读数角标
    updateBadge({ state }) {
      // 开启监听总未读数变化
      ev.$on('totalunread', num => {
        state.totalunread = num
      })
      state.chat.updateBadge()
    },
    // 断线自动重连
    reconnect({ state }) {
      if (state.user && state.chat) {
        state.chat.reconnect()
      }
    }
  }
}
