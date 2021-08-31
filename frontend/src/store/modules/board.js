import * as API from '../../api'
import router from '../../router'

const state = {
  id: Number,
  loading: Boolean,
  name: String,
  lists: [],
}

const getters = {
  name: state => state.name,
  loading: state => state.loading,
  lists: state => state.lists
}

const actions = {
  async fetchBoard({ commit }, id) {
    commit('SET_LOADING', true)
    try {
      const board = await API.GET(API.URL + "/board/" + id)
      commit('LOAD_BOARD', board)
    } catch(err) {
      alert("Error when loading board " + err)
      router.push('/notfound')
    }
    commit('SET_LOADING', false)
  },

  async newList({ commit }) {
    try {
      const response = await API.POST(API.URL + "/board/" + 1 + "/new")
      commit('NEW_LIST', response)
    } catch(err) {
      alert("Error when creating new list " + err)
    }
  },

  async updateTask({ commit }, {taskID, newDescription, callback}) {
    try {
      let listID
      let newTask = {
        id: taskID,
        description: newDescription
      } 
      state.lists.forEach(list => {
        if ( list.tasks.find(task => taskID === task.id) ) {
          listID = list.id
          return;
        }
      })

      const updatedTask = await API.PUT(API.URL + "/board/" + state.id + "/" + listID + '/' + taskID, newTask)
      commit('UPDATE_TASK', updatedTask)
    } catch(err) {
      alert("Error when updating task" + err)
    }
    callback()
  }
}

const mutations = {
  SET_LOADING: (state, bool) => state.loading = bool,

  LOAD_BOARD: (state, data) => {
    state.id = data.id
    state.name = data.name
    state.lists = data.lists
  },

  NEW_LIST: (state, data) => state.lists.push(data),

  UPDATE_TASK: (state, updatedTask) => {
    state.lists.forEach(list => {
      list.tasks.forEach(task => {
        if(task.id === updatedTask.id) {
          task.description = updatedTask.description
        }
      })
    })
  }
}

export default{
  state,
  getters,
  actions,
  mutations
}
