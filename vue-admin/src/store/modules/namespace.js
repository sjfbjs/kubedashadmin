import nameSpaceSettings from '@/settings'

const { getNameSpace, setNameSpace } = nameSpaceSettings

const state = {
  getNameSpace: getNameSpace,
  setNameSpace: setNameSpace
}

const mutations = {
  SET_NAMESPACE: (state, { key, value }) => {
    if (state.hasOwnProperty(key)) {
      state[key] = value
    }
  }
}

const actions = {
  changeSetting({ commit }, data) {
    commit('SET_NAMESPACE', data)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

