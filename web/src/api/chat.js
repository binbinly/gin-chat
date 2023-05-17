import api from './index'

export function chatSend(data) {
  return api.post(api.Chat.Send, data, true, true)
}

export function chatMessage() {
  return api.post(api.Chat.Message)
}

export function chatRecall(data) {
  return api.post(api.Chat.Recall, data)
}

export function chatDetail(data) {
  return api.post(api.Chat.Detail, data)
}
