import api from './index'

export function momentCreate(data) {
  return api.post(api.Moment.Create, data)
}

export function momentLike(data) {
  return api.post(api.Moment.Like, data)
}

export function momentComment(data) {
  return api.post(api.Moment.Comment, data)
}

export function momentList(user_id, p) {
  if (user_id > 0) {
    return api.get(api.Moment.List, { user_id, p })
  }
  return api.get(api.Moment.Timeline, { p })
}
