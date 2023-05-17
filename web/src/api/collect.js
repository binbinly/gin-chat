import api from './index'

export function collectCreate(data) {
  return api.post(api.Collect.Create, data)
}

export function collectList(p) {
  return api.get(api.Collect.List, { p })
}

export function collectDestroy(id) {
  return api.post(api.Collect.Destroy, { id })
}
