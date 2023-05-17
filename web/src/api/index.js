import request from '@/utils/request'

const api = {
  //用户登录
  Login: '/login',
  // 用户注册
  Reg: '/reg',
  //用户注销
  Logout: '/user/logout',
  //搜索用户
  SearchUser: '/user/search',
  //表情包分类
  EmoticonCat: '/emoticon/cat',
  //表情
  Emoticon: '/emoticon/list',
  //文件上传
  Upload: '/upload/image',
  Apply: {
    //申请添加好友
    Friend: '/apply/friend',
    //申请列表
    List: '/apply/list',
    //申请处理
    Handle: '/apply/handle',
    //申请数量
    Count: '/apply/count'
  },
  Friend: {
    //好友列表
    List: '/friend/list',
    //好友资料
    Info: '/friend/info',
    //移入/移除黑名单
    Black: '/friend/black',
    //移入/移除星标好友
    Star: '/friend/star',
    //设置朋友圈权限
    MomentAuth: '/friend/auth',
    //设置备注和标签
    RemarkTag: '/friend/remark',
    //举报投诉好友/群组
    Report: '/user/report',
    //删除好友
    Destroy: '/friend/destroy'
  },
  Chat: {
    //发送消息
    Send: '/chat/send',
    //撤回消息
    Recall: '/chat/recall',
    //聊天回话详情
    Detail: '/chat/detail'
  },
  Group: {
    //创建群组
    Create: '/group/create',
    //群聊列表
    List: '/group/list',
    //群信息
    Info: '/group/info',
    //群成员
    User: '/group/user',
    //修改群信息
    Edit: '/group/edit',
    //删除并退出群聊
    Quit: '/group/quit',
    //修改我在本群中的昵称
    Nickname: '/group/nickname',
    //将某个群成员踢出
    KickOff: '/group/kickoff',
    //邀请加入群聊
    Invite: '/group/invite',
    //加入群聊
    Join: '/group/join'
  },
  Tag: {
    //标签列表
    List: '/user/tag',
    //标签用户列表
    UserList: '/friend/tag_list'
  },
  User: {
    //修改用户个人资料
    Edit: '/user/edit'
  },
  Collect: {
    //创建收藏
    Create: '/collect/create',
    //收藏列表
    List: '/collect/list',
    //删除收藏
    Destroy: '/collect/destroy'
  },
  Moment: {
    //发布朋友圈
    Create: '/moment/create',
    //点赞朋友圈
    Like: '/moment/like',
    //评论朋友圈
    Comment: '/moment/comment',
    //动态列表
    List: '/moment/list',
    //我的朋友圈
    Timeline: '/moment/timeline'
  },
  post(url, data, auth = true, hideloading = false) {
    return request({
      url,
      method: 'post',
      data,
      auth,
      hideloading
    })
  },
  get(url, params = null, hideloading = false) {
    return request({
      url,
      params,
      hideloading,
      auth: true
    })
  }
}

export default api
