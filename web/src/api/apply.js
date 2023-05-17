import api from './index'

export function applyFriend(data) {
  return api.post(api.Apply.Friend, data)
}

export function applyList(p) {
  return api.get(api.Apply.List, { p })
}

export function applyCount() {
  return api.get(api.Apply.Count)
}

export function applyHandle(data) {
  return api.post(api.Apply.Handle, data)
}
