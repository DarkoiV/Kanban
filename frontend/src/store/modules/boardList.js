import * as API from '../../api'
import router from '../../router'

const state = {
  boards: []
}

const getters = {
  boards: state => state.boards
}

const actions = {
  async getBoards({ commit }) {
    try {
      const listOfBoards = await API.GET(`${API.URL}/board/list`)
      commit('SET_BOARDS_LIST', listOfBoards)
    } catch(err) {
      alert(err)
    }
  },

  async newBoard() {
    try {
      const board = {
        name: "NEW BOARD"
      }
      const newBoard = await API.POST(`${API.URL}/board`, board)
      router.push({name: 'Board', params: {id: newBoard.id}})
    } catch(err) {
      alert(err)
    }
  }
}

const mutations = {
  SET_BOARDS_LIST: (state, listOfBoards) => {
    state.boards = listOfBoards
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
