import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'
import audio from './modules/audio'
import user from './modules/user'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    audio,
    user
  },
  getters
})

export default store
