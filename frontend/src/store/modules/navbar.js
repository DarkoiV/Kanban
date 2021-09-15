const state = {
  navTitle: String,
  onBoard: Boolean
}

const getters = {
  navTitle: state => state.navTitle,
  onBoard: state => state.onBoard
}

const actions = {
  setNavTitle({ commit }, newTitle) {
    commit('SET_NAV_TITLE', newTitle)
  },

}

const mutations = {
  SET_NAV_TITLE: (state, newTitle) => {
    state.navTitle = newTitle 
  },

  SET_ON_BOARD: (state, onBoard) => {
    state.onBoard = onBoard
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
