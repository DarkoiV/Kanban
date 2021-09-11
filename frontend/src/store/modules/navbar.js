const state = {
  navTitle: String
}

const getters = {
  navTitle: state => state.navTitle,
}

const actions = {
  setNavTitle({ commit }, newTitle) {
    commit('SET_NAV_TITLE', newTitle)
  }
}

const mutations = {
  SET_NAV_TITLE: (state, newTitle) => {
    state.navTitle = newTitle 
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
