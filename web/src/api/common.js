import api from './index'

export function login(data) {
  return api.post(api.Login, data, false)
}

export function reg(data) {
  return api.post(api.Reg, data, false)
}

export function searchUser(data) {
  return api.post(api.SearchUser, data)
}

export function emoticonCat() {
  return api.get(api.EmoticonCat)
}

export function emoticon(cat) {
  return api.get(api.Emoticon, { cat }, true)
}

export function httpGet(url) {
  return api.get(url)
}

export function uploadFile(file) {
  console.log('upload file', file)
  return api
      .upload(api.Upload, { file })
      .then(res => {return res})
      .catch(e => {
        console.error('get upload url', e)
        reject(e)
      })
}
