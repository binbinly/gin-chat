import api from './index'

export function groupCreate(data) {
  return api.post(api.Group.Create, data)
}

export function groupInvite(data) {
  return api.post(api.Group.Invite, data)
}

export function groupList(p) {
  return api.get(api.Group.List, { p })
}

export function groupInfo(id) {
  return api.get(api.Group.Info, { id })
}

export function groupEdit(data) {
  return api.post(api.Group.Edit, data)
}

export function groupNickname(data) {
  return api.post(api.Group.Nickname, data)
}

export function groupQuit(data) {
  return api.get(api.Group.Quit, data)
}

export function groupKick(data) {
  return api.post(api.Group.KickOff, data)
}

export function groupUser(id) {
  return api.get(api.Group.User, { id })
}
