import api from './index'

export function tagList() {
  return api.get(api.Tag.List)
}

export function tagUserList(id) {
  return api.get(api.Tag.UserList, { id })
}
