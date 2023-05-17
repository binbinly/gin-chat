import api from './index'

export function userEdit(data) {
  return api.post(api.User.Edit, data)
}

export function userLogout() {
  return api.get(api.Logout)
}
