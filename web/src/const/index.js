const system = 1
const text = 2
const image = 3
const video = 4
const audio = 5
const emoticon = 6
const card = 7

export default {
  PAGE_SIZE: 20, //分页大小

  CHAT_TYPE_USER: 1, //好友聊天
  CHAT_TYPE_GROUP: 2, //群聊

  TYPE_SYSTEM: system, //系统消息
  TYPE_TEXT: text, //文本消息
  TYPE_IMAGE: image, //图片消息
  TYPE_VIDEO: video, //视频消息
  TYPE_AUDIO: audio, //音频消息
  TYPE_EMOTICON: emoticon, //表情
  TYPE_CARD: card, //名片

  TYPE_LIST: {
    system: system,
    text: text,
    image: image,
    video: video,
    audio: audio,
    emoticon: emoticon,
    card: card
  },
  TYPE_TRANS_LIST: {
    1: 'system',
    2: 'text',
    3: 'image',
    4: 'video',
    5: 'audio',
    6: 'emoticon',
    7: 'card'
  }
}
