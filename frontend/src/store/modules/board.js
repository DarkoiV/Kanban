import * as API from '../../api'
import router from '../../router'

const state = {
  id: Number,
  name: String,
  lists: [],
}

const getters = {
  name: state => state.name,
  lists: state => state.lists
}

const actions = {
  async fetchBoard({ commit }, id) {
    try {
      const board = await API.GET(API.URL + "/board/" + id)
      commit('loadBoard', board)
    } catch(err) {
      alert("Error when loading board " + err)
      router.push('/notfound')
    }
  },

  async newList({ commit }) {
    try {
      const response = await API.POST(API.URL + "/board/" + 1 + "/new")
      commit('newList', response)
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
      commit('updateTask', updatedTask)
    } catch(err) {
      alert("Error when updating task" + err)
    }
    callback()
  }
}

const mutations = {
  loadBoard: (state, data) => {
    state.id = data.id
    state.name = data.name
    state.lists = data.lists
    state.loaded = true
  },

  newList: (state, data) => state.lists.push(data),

  updateTask: (state, updatedTask) => {
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
