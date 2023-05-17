const getters = {
  name: state => state.user.nickname || state.user.username
}
export default getters
