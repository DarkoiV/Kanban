import * as API from '../../api'

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
      alert("Error getting boards " + err)
    }
  },

  async renameBoard({commit}, {boardID, newName}) {
    try {
      const reqBody = {
        name: newName
      }
      await API.PATCH(`${API.URL}/board/${boardID}`, reqBody)
      commit("RENAME_BOARD", {boardID, newName})
    } catch(err) {
      alert("Error renaming board " + err)
    }
  },

  async newBoard({dispatch}) {
    try {
      const board = {
        name: "NEW BOARD"
      }
      await API.POST(`${API.URL}/board`, board)
      dispatch('getBoards')
    } catch(err) {
      alert(err)
    }
  },

  async removeBoard({commit}, id) {
    try {
      await API.DELETE(`${API.URL}/board/${id}`)
      const listOfBoards = await API.GET(`${API.URL}/board/list`)
      commit('SET_BOARDS_LIST', listOfBoards)
    } catch(err) {
      alert("Error deleting board " + err)
    }
  }
}

const mutations = {
  SET_BOARDS_LIST: (state, listOfBoards) => {
    state.boards = listOfBoards
  },

  RENAME_BOARD: (state, {boardID, newName}) => {
    const board = state.boards.find(board => board.id == boardID)
    board.name = newName
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
