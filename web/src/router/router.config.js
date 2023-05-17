/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/',
    component: () => import('@/views/layouts/index'),
    redirect: '/home',
    meta: {
      title: '首页',
      keepAlive: false
    },
    children: [
      {
        path: '/home',
        name: 'Home',
        component: () => import('@/views/home/index'),
        meta: { title: '首页', keepAlive: false, tabbar: true, badge: '' }
      },
      {
        path: '/contact',
        name: 'Contact',
        component: () => import('@/views/home/contact'),
        meta: { title: '通讯录', keepAlive: false, tabbar: true, badge: '' }
      },
      {
        path: '/find',
        name: 'Find',
        component: () => import('@/views/home/find'),
        meta: { title: '发现', keepAlive: false, tabbar: true, badge: '' }
      },
      {
        path: '/my',
        name: 'My',
        component: () => import('@/views/home/my'),
        meta: { title: '我的', keepAlive: false, tabbar: true, badge: '' }
      },
      {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/common/login'),
        meta: { title: '登录', keepAlive: false }
      },
      {
        path: '/apply_list',
        name: 'ApplyList',
        component: () => import('@/views/contact/apply-list'),
        meta: { title: '申请好友列表', keepAlive: true }
      },
      {
        path: '/add_friend',
        name: 'AddFriend',
        component: () => import('@/views/contact/add-friend'),
        meta: { title: '添加好友', keepAlive: true }
      },
      {
        path: '/search',
        name: 'Search',
        component: () => import('@/views/common/search'),
        meta: { title: '搜索', keepAlive: false }
      },
      {
        path: '/user_base',
        name: 'UserBase',
        component: () => import('@/views/contact/user-base'),
        meta: { title: '用户信息', keepAlive: true }
      },
      {
        path: '/user_remark_tag',
        name: 'UserRemarkTag',
        component: () => import('@/views/contact/user-remark-tag'),
        meta: { title: '设置备注和标签', keepAlive: true }
      },
      {
        path: '/chat_list',
        name: 'ChatList',
        component: () => import('@/views/chat/chat-list'),
        meta: { title: '好友列表', keepAlive: true }
      },
      {
        path: '/user_moments_auth',
        name: 'UserMomentsAuth',
        component: () => import('@/views/contact/user-moments-auth'),
        meta: { title: '设置朋友圈权限', keepAlive: true }
      },
      {
        path: '/user_report',
        name: 'UserReport',
        component: () => import('@/views/contact/user-report'),
        meta: { title: '投诉', keepAlive: true }
      },
      {
        path: '/tag_list',
        name: 'TagList',
        component: () => import('@/views/contact/tag-list'),
        meta: { title: '标签', keepAlive: true }
      },
      {
        path: '/tag_read',
        name: 'TagRead',
        component: () => import('@/views/contact/tag-read'),
        meta: { title: '标签用户', keepAlive: true }
      },
      {
        path: '/contact_list',
        name: 'ContactList',
        component: () => import('@/views/contact/contact-list'),
        meta: { title: '好友列表', keepAlive: true }
      },
      {
        path: '/group_list',
        name: 'GroupList',
        component: () => import('@/views/contact/group-list'),
        meta: { title: '群组列表', keepAlive: true }
      },
      {
        path: '/userinfo',
        name: 'Userinfo',
        component: () => import('@/views/my/userinfo'),
        meta: { title: '用户资料', keepAlive: true }
      },
      {
        path: '/fava',
        name: 'Fava',
        component: () => import('@/views/my/fava'),
        meta: { title: '收藏', keepAlive: true }
      },
      {
        path: '/setting',
        name: 'Setting',
        component: () => import('@/views/my/setting'),
        meta: { title: '设置', keepAlive: false }
      },
      {
        path: '/chat',
        name: 'Chat',
        component: () => import('@/views/chat/chat'),
        meta: { title: '聊天窗口', keepAlive: true }
      },
      {
        path: '/chat_set',
        name: 'ChatSet',
        component: () => import('@/views/chat/chat-set'),
        meta: { title: '聊天信息', keepAlive: true }
      },
      {
        path: '/chat_history',
        name: 'ChatHistory',
        component: () => import('@/views/chat/chat-history'),
        meta: { title: '聊天记录', keepAlive: true }
      },
      {
        path: '/group_user',
        name: 'GroupUser',
        component: () => import('@/views/chat/group-user'),
        meta: { title: '群聊用户', keepAlive: true }
      },
      {
        path: '/moments',
        name: 'Moments',
        component: () => import('@/views/moments/index'),
        meta: { title: '朋友圈', keepAlive: true }
      },
      {
        path: '/add_moment',
        name: 'AddMoment',
        component: () => import('@/views/moments/add'),
        meta: { title: '心情发布', keepAlive: true }
      },
      {
        path: '/emoticon_cat',
        name: 'EmoticonCat',
        component: () => import('@/views/common/emo-cat'),
        meta: { title: '表情包', keepAlive: true }
      },
      {
        path: '/emoticon',
        name: 'Emoticon',
        component: () => import('@/views/common/emo'),
        meta: { title: '表情', keepAlive: true }
      }
    ]
  }
]
