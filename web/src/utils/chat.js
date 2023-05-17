import { setStorage, getStorage, removeStorage } from '@/utils/index.js'
import $store from '@/store/index.js'
import { Dialog, Toast } from 'vant'
import { chatSend, chatRecall } from '@/api/chat.js'
import ev from '@/utils/event.js'
import $const from '@/const/index.js'

class chat {
  constructor(arg) {
    //ws地址
    this.url = arg.url
    //是否在线
    this.isOnline = false
    this.socket = null
    //自动重连次数
    this.reconnectCount = 0
    //是否自动重连
    this.isOpenReconnect = true
    //心跳
    this.heart = null
    // 获取当前用户相关信息
    let user = getStorage('user')
    if (user) {
      user = JSON.parse(user)
      this.user = { ...user }
    } else {
      this.user = {}
    }
    // 初始化聊天对象
    this.TO = false

    this.token = getStorage('token')
    // websocket连接和监听
    if (this.token) {
      this.connectSocket()
    }
  }
  // 断线重连
  reconnect() {
    if (this.isOnline) {
      return
    }
    console.log('reconnection', this.reconnectCount)
    if (this.reconnectCount >= 3) {
      //自动重连次数限制，超过后弹窗
      return this.reconnectConfirm()
    }
    this.reconnectCount += 1
    this.connectSocket()
  }
  // 连接socket
  connectSocket() {
    this.socket = new WebSocket(this.url + '?token=' + this.token)
    if (!this.socket) {
      console.log('socket empty')
      return
    }
    // 监听连接成功
    this.socket.onopen = () => this.onOpen()
    // 监听接收信息
    this.socket.onmessage = res => this.onMessage(res)
    // 监听断开
    this.socket.onclose = e => this.onClose(e)
    // 监听错误
    this.socket.onerror = e => this.onError(e)
  }
  //45s 心跳
  heartStart() {
    this.heart = setInterval(() => {
      this.socket.send('ping')
    }, 45000)
  }
  // 监听打开
  onOpen() {
    // 用户上线
    console.log('socket连接成功', this.socket.readyState)
    this.isOnline = true
    this.isOpenReconnect = true
    //心跳
    this.heartStart()
  }
  clearUser() {
    // 用户下线
    this.isOnline = false
    this.socket = null
    //清除心跳定时器
    clearInterval(this.heart)
    this.heart = null
  }
  // 监听关闭
  onClose(e) {
    console.log('socket连接关闭', e)
    this.clearUser()
    if (this.isOpenReconnect) {
      this.reconnect()
    }
  }
  // 监听连接错误
  onError(e) {
    console.log('socket连接错误', e)
    this.clearUser()
  }
  // 监听接收消息
  onMessage(data) {
    if (data.data == 'pong') {
      //心跳响应，无需处理
      return
    }
    console.log('data', data.data)
    const res = JSON.parse(data.data)

    switch (res.event) {
      case 'fail': // 错误消息
        Toast.fail(res.data)
        break
      case 'close': // 客户端关闭自动重连
        Toast.fail(res.data)
        this.close()
        $store.dispatch('logout')
        break
      case 'recall': // 撤回消息
        this.handleOnRecall(res.data)
        break
      case 'notify': // 通知
        if (res.data.type && res.data.type == 'apply') {
          // 申请加好友通知
          $store.dispatch('getApply')
        }
        break
      case 'moment': // 朋友圈更新
        this.handleMoment(res.data)
        break
      case 'chat': // 聊天消息
        this.handleOnMessage(res.data)
        break
      default:
        console.log('default')
    }
  }
  // 获取本地存储中的朋友圈动态通知
  getNotice() {
    let notice = getStorage('moment_' + this.user.id)
    return notice ? JSON.parse(notice) : { avatar: '', user_id: 0, num: 0 }
  }
  // 处理朋友圈通知
  async handleMoment(message) {
    let notice = this.getNotice()
    switch (message.type) {
      case 'new':
        if (message.user_id !== this.user.id) {
          notice.avatar = message.avatar
          notice.user_id = message.user_id
          notice.num += 1
          ev.$emit('tabBarBadge', {
            index: 2,
            text: ''
          })
        }
        break
      default:
        if (message.user_id !== this.user.id) {
          notice.avatar = message.avatar
          notice.user_id = message.user_id
          notice.num += 1
        }
        if (notice.num > 0) {
          ev.$emit('tabBarBadge', {
            index: 2,
            text: notice.num > 99 ? '99+' : notice.num.toString()
          })
        } else {
          ev.$emit('tabBarBadge', {
            index: 2,
            text: ''
          })
        }
        break
    }
    ev.$emit('momentNotice', notice)
    setStorage('moment_' + this.user.id, JSON.stringify(notice))
  }
  // 读取朋友圈动态
  async readMoments() {
    let notice = { avatar: '', user_id: 0, num: 0 }
    setStorage('moment_' + this.user.id, JSON.stringify(notice))
    ev.$emit('tabBarBadge', {
      index: 2,
      text: ''
    })
    ev.$emit('momentNotice', notice)
  }
  // 监听撤回消息处理
  async handleOnRecall(message) {
    // 通知聊天页撤回消息
    ev.$emit('onRecall', message)
    // 修改聊天记录
    let id = message.chat_type === $const.CHAT_TYPE_USER ? message.from_id : message.to_id
    // 删除本地消息
    this.deleteChatDetailItem(message.id, message.chat_type, id)
    // 当前会话最后一条消息的显示
    this.updateChatItem(
      {
        id,
        chat_type: message.chat_type
      },
      item => {
        item.content = '对方撤回了一条消息'
        item.t = new Date().getTime()
        return item
      }
    )
  }
  // 处理消息
  async handleOnMessage(message) {
    // 添加消息记录到本地存储中
    if (message.chat_type === $const.CHAT_TYPE_USER) {
      //单聊，接受者就是当前用户
      message.to = { id: this.user.id }
    }
    let { data } = this.addChatDetail(message, false)
    // 更新会话列表
    this.updateChatList(data, false)
    // 全局通知
    ev.$emit('onMessage', data)
  }
  // 关闭连接
  close() {
    if (this.socket) {
      this.socket.close()
    }
    this.isOpenReconnect = false
  }
  // 创建聊天对象
  createChatObject(detail) {
    this.TO = detail
    console.log('创建聊天对象', this.TO)
  }
  // 销毁聊天对象
  destroyChatObject() {
    this.TO = false
    console.log('销毁聊天对象')
  }
  // 断线重连提示
  reconnectConfirm() {
    Dialog.confirm({
      message: '你已经断线，是否重新连接？',
      confirmButtonText: '重新连接'
    }).then(() => {
      this.reconnectCount = 0
      this.connectSocket()
    })
  }
  // 验证是否上线
  checkOnline() {
    if (!this.isOnline) {
      // 断线重连提示
      this.reconnectConfirm()
      return false
    }
    return true
  }
  // 组织发送信息格式
  formatSendData(params) {
    return {
      id: '', // 唯一id，后端生成，用于撤回指定消息
      from: this.user,
      to: this.TO || params.to, // 接收人/群 id
      chat_type: params.chat_type || this.TO.chat_type, // 接收类型
      type: $const.TYPE_LIST[params.type], // 消息类型
      content: params.content, // 消息内容
      options: params.options ? params.options : {}, // 其他参数
      t: new Date().getTime(), // 创建时间
      self: true, //自己发送消息
      status: params.status ? params.status : 'pending' // 发送状态，success发送成功,fail发送失败,pending发送中
    }
  }
  // 撤回消息
  recall(message) {
    return new Promise((result, reject) => {
      chatRecall({
        to_id: this.TO.id,
        chat_type: message.chat_type,
        id: message.id
      })
        .then(res => {
          this.deleteChatDetailItem(message.id, message.chat_type, this.TO.id)
          result(res)
          // 更新会话最后一条消息显示
          this.updateChatItem(
            {
              id: message.TO.id,
              chat_type: message.chat_type
            },
            item => {
              item.content = '你撤回了一条消息'
              item.t = new Date().getTime()
              return item
            }
          )
        })
        .catch(err => {
          reject(err)
        })
    })
  }
  // 发送消息
  send(message, onProgress = false) {
    return new Promise(async (result, reject) => {
      // 验证是否上线
      if (!this.checkOnline()) return reject('未上线')
      // 添加消息历史记录
      const { k } = this.addChatDetail(message)
      // 提交到后端
      chatSend({
        to_id: this.TO.id || message.to.id,
        chat_type: message.chat_type || this.TO.chat_type,
        type: message.type,
        content: message.content,
        options: message.options
      })
        .then(res => {
          // 发送成功
          message.id = res.id
          message.status = 'success'
          // 更新指定历史记录
          console.log('更新指定历史记录', message, k)
          this.updateChatDetail(message, k)
          result(res)
        })
        .catch(err => {
          // 发送失败
          message.status = 'fail'
          // 更新指定历史记录
          this.updateChatDetail(message, k)
          // 断线重连提示
          reject(err)
        })
      // 更新会话列表
      this.updateChatList(message)
    })
  }
  // 添加聊天记录
  addChatDetail(message, isSend = true) {
    console.log('添加聊天记录')
    // 获取接受者id
    let id = message.chat_type === $const.CHAT_TYPE_USER ? (isSend ? message.to.id : message.from.id) : message.to.id
    // 获取原来的聊天记录
    let list = this.getChatDetail(message.chat_type, id)
    console.log('获取原来的聊天记录', list)
    // 标识
    message.k = 'k' + list.length
    list.push(message)
    // 加入缓存
    this.setChatDetail(message.chat_type, id, list)
    // 返回
    return {
      data: message,
      k: message.k
    }
  }
  /**
   * 删除本地指定聊天记录
   * @param {*} message_id 消息id
   * @param {*} chat_type 会话类型
   * @param {*} id 会话id
   */
  async deleteChatDetailItem(message_id, chat_type, id) {
    // 获取原来的聊天记录
    let list = this.getChatDetail(chat_type, id)
    // 根据k查找对应聊天记录
    let index = list.findIndex(item => item.id === message_id)
    if (index === -1) return
    list.splice(index, 1)
    // 存储
    this.setChatDetail(chat_type, id, list)
  }
  // 更新指定历史记录
  async updateChatDetail(message, k, isSend = true) {
    // 获取对方id
    let id = message.chat_type === $const.CHAT_TYPE_USER ? (isSend ? message.to.id : message.from.id) : message.to.id
    // 获取原来的聊天记录
    let list = this.getChatDetail(message.chat_type, id)
    console.log('获取原来的聊天记录', list)
    // 根据k查找对应聊天记录
    let index = list.findIndex(item => item.k === k)
    console.log('根据k查找对应聊天记录', index)
    if (index === -1) return
    list[index] = message
    // 存储
    this.setChatDetail(message.chat_type, id, list)
  }
  // 格式化会话最后一条消息显示
  formatChatItemData(message, isSend) {
    let content = message.content
    switch (message.type) {
      case $const.TYPE_EMOTICON:
        content = '[表情]'
        break
      case $const.TYPE_IMAGE:
        content = '[图片]'
        break
      case $const.TYPE_AUDIO:
        content = '[语音]'
        break
      case $const.TYPE_VIDEO:
        content = '[视频]'
        break
      case $const.TYPE_CARD:
        content = '[名片]'
        break
    }
    return isSend ? content : `${message.from.name}: ${content}`
  }
  // 更新会话列表
  updateChatList(message, isSend = true) {
    // 获取本地存储会话列表
    let list = this.getChatList()
    // 是否处于当前聊天中
    let isCurrentChat = false
    // 接收人/群
    let to = {}

    // 判断私聊还是群聊
    if (message.chat_type === $const.CHAT_TYPE_USER) {
      // 私聊
      // 聊天对象是否存在
      isCurrentChat = this.TO ? (isSend ? this.TO.id === message.to.id : this.TO.id === message.from.id) : false
      to = isSend ? this.TO : message.from
    } else {
      // 群聊
      isCurrentChat = this.TO && this.TO.id === message.to.id
      to = message.to
    }

    // 会话是否存在
    let index = list.findIndex(item => {
      return item.chat_type === message.chat_type && item.id === to.id
    })
    console.log('接收人/群', to, list, index)
    // 最后一条消息展现形式
    const content = this.formatChatItemData(message, isSend)
    // 会话不存在，创建会话
    // 未读数是否 + 1
    let unread = isSend || isCurrentChat ? 0 : 1
    if (index === -1) {
      let chatItem = {
        id: to.id,
        name: to.name,
        avatar: to.avatar,
        chat_type: message.chat_type, // 接收类型 user单聊 group群聊
        t: message.t || new Date().getTime(), // 最后一条消息的时间戳
        content, // 最后一条消息内容
        type: message.type, // 最后一条消息类型
        unread, // 未读数
        is_top: false, // 是否置顶
        show_name: false, // 是否显示昵称
        no_remind: false, // 消息免打扰
        is_remind: false // 是否开启强提醒
      }
      // 群聊
      if (message.chat_type === $const.CHAT_TYPE_GROUP) {
        chatItem.show_name = true
      }
      list.unshift(chatItem)
    } else {
      // 存在，更新会话
      // 拿到当前会话
      let item = list[index]
      // 更新该会话最后一条消息时间，内容，类型
      item.t = message.t || new Date().getTime()
      item.name = to.name
      item.avatar = to.avatar
      item.content = content
      item.type = message.type
      // 未读数更新
      item.unread += unread
      // 置顶会话
      list = this.listToFirst(list, index)
    }
    // 存储
    this.setChatList(list)
    // 更新未读数
    this.updateBadge(list)
    // 通知更新vuex中的聊天会话列表
    ev.$emit('onUpdateChatList', list)
    return list
  }
  // 更新未读数
  async updateBadge(list = false) {
    // 获取所有会话列表
    list = list || this.getChatList()
    // 统计所有未读数
    let total = 0
    list.forEach(item => {
      total += item.unread
    })
    // 设置底部导航栏角标
    if (total > 0) {
      ev.$emit('tabBarBadge', {
        index: 0,
        text: total <= 99 ? total.toString() : '99+'
      })
    } else {
      ev.$emit('tabBarBadge', {
        index: 0,
        text: ''
      })
    }
    ev.$emit('totalunread', total)
  }
  // 更新指定会话指定键值
  async updateChatItemKey(id, chat_type, key, value) {
    // 获取所有会话列表
    let list = this.getChatList()
    // 找到当前会话
    let index = list.findIndex(item => item.id === id && item.chat_type === chat_type)
    if (index === -1) return
    // 更新数据
    list[index][key] = value
    this.setChatList(list)
  }
  // 更新指定会话
  async updateChatItem(where, data) {
    // 获取所有会话列表
    let list = this.getChatList()
    // 找到当前会话
    let index = list.findIndex(item => item.id === where.id && item.chat_type === where.chat_type)
    if (index === -1) return
    // 更新数据
    if (typeof data === 'function') {
      list[index] = data(list[index])
    } else {
      list[index] = data
    }

    this.setChatList(list)

    // 更新会话列表状态
    ev.$emit('onUpdateChatList', list)
  }
  // 读取会话
  async readChatItem(id, chat_type) {
    // 获取所有会话列表
    let list = this.getChatList()
    // 找到当前会话
    let index = list.findIndex(item => item.id === id && item.chat_type === chat_type)
    if (index !== -1) {
      list[index].unread = 0
      this.setChatList(list)
      // 重新获取总未读数
      this.updateBadge()
      // 更新会话列表状态
      ev.$emit('onUpdateChatList', list)
    }
  }
  // 删除指定会话
  async removeChatItem(id, chat_type) {
    // 获取所有会话列表
    let list = this.getChatList()
    // 找到当前会话
    let index = list.findIndex(item => item.id === id && item.chat_type === chat_type)
    if (index !== -1) {
      list.splice(index, 1)

      this.setChatList(list)
      // 重新获取总未读数
      this.updateBadge()
      // 更新会话列表状态
      ev.$emit('onUpdateChatList', list)
    }
  }
  // 清空聊天记录
  async clearChatDetail(id, chat_type) {
    this.removeChatDetail(chat_type, id)

    // 获取所有会话列表
    let list = this.getChatList()
    // 找到当前会话
    let index = list.findIndex(item => item.id === id && item.chat_type === chat_type)
    if (index !== -1) {
      list[index].content = ''

      this.setChatList(list)
      // 更新会话列表状态
      ev.$emit('onUpdateChatList', list)
    }
  }
  chatKey(chat_type, to_id) {
    // key值：chatDetail_当前用户id_会话类型_接收人/群id
    return `chatDetail_${this.user.id}_${chat_type}_${to_id}`
  }
  // 获取聊天记录
  getChatDetail(chat_type, to_id) {
    chat_type = chat_type || this.TO.chat_type
    to_id = to_id || this.TO.id
    return this.getStorage(this.chatKey(chat_type, to_id))
  }
  // 设置聊天记录
  setChatDetail(chat_type, to_id, list) {
    return this.setStorage(this.chatKey(chat_type, to_id), list)
  }
  // 删除聊天记录
  removeChatDetail(chat_type, to_id) {
    return removeStorage(this.chatKey(chat_type, to_id))
  }
  // 获取本地存储会话列表
  getChatList() {
    const key = `chatList_${this.user.id}`
    return this.getStorage(key)
  }
  setChatList(list) {
    const key = `chatList_${this.user.id}`
    return this.setStorage(key, list)
  }
  // 获取指定会话
  getChatListItem(id, chat_type) {
    // 获取所有会话列表
    let list = this.getChatList()
    // 找到当前会话
    let index = list.findIndex(item => item.id == id && item.chat_type == chat_type)
    if (index !== -1) {
      return list[index]
    }
    return false
  }
  // 获取存储
  getStorage(key) {
    let list = getStorage(key)
    return list ? JSON.parse(list) : []
  }
  // 设置存储
  setStorage(key, value) {
    return setStorage(key, JSON.stringify(value))
  }
  // 数组置顶
  listToFirst(arr, index) {
    if (index != 0) {
      arr.unshift(arr.splice(index, 1)[0])
    }
    return arr
  }
}
export default chat
