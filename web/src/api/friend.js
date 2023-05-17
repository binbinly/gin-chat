import api from './index'

export function friendList() {
  return api.get(api.Friend.List, null, true)
}

export function friendRead(id) {
  return api.get(api.Friend.Info, { id })
}

export function friendDestroy(user_id) {
  return api.post(api.Friend.Destroy, { user_id })
}

export function friendStar(data) {
  return api.post(api.Friend.Star, data)
}

export function friendBlack(data) {
  return api.post(api.Friend.Black, data)
}

export function friendRemarkTag(data) {
  return api.post(api.Friend.RemarkTag, data)
}

export function friendMomentAuth(data) {
  return api.post(api.Friend.MomentAuth, data)
}

export function userReport(data) {
  return api.post(api.Friend.Report, data)
}
