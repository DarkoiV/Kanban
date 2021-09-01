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
      const newListBody = {
        title: "New List",
        pos: state.lists.length
      }
      const newList = await API.POST(API.URL + "/board/" + state.id + "/new", newListBody)
      newList.tasks = []
      commit('NEW_LIST', newList)
    } catch(err) {
      alert("Error when creating new list " + err)
    }
  },

  async newTask({ commit }, listID) {
    try {
      const list = state.lists.find(list => list.id === listID)
      const newTaskBody = {
        description: "Double click to edit",
        pos: list.tasks.length
      }
      const newTask = await API.POST(`${API.URL}/board/${state.id}/${listID}/new`, newTaskBody)
      commit('NEW_TASK', {listID: listID, newTask: newTask})
    } catch(err) {
      alert("Error when creating new task" + err)
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
      alert("Error when updating task " + err)
    }
    callback()
  },

  async dropTask({ commit }, { taskID, listID }) {
    try {
      const newList = state.lists.find(list => listID == list.id)
      const oldList = state.lists.find(list =>
        list.tasks.find(task => taskID == task.id)
      )
      let movedTask = oldList.tasks.find(task => task.id == taskID)

      movedTask.pos = newList.tasks.lenght
      movedTask.listID = listID

      await API.PUT(`${API.URL}/board/${state.id}/${oldList.id}/${taskID}`, movedTask)

      console.log(oldList)
      commit('DROP_ON_LIST', {listID, movedTask})
      commit('REMOVE_FROM_LIST', {listID: oldList.id, taskID,})
    }
    catch(err) {
      alert("Error when moving task " + err)
    }

  }
}

const mutations = {
  SET_LOADING: (state, bool) => state.loading = bool,

  LOAD_BOARD: (state, data) => {
    state.id = data.id
    state.name = data.name
    state.lists = data.lists
    
    // Sort
    state.lists.forEach( list => 
      list.tasks.sort( (a, b) => a.pos - b.pos)
    )
  },

  NEW_LIST: (state, newList) => state.lists.push(newList),
  NEW_TASK: (state, {listID, newTask}) => {
    const list = state.lists.find( list => listID === list.id )
    list.tasks.push(newTask)
  },

  UPDATE_TASK: (state, updatedTask) => {
    state.lists.forEach(list => {
      list.tasks.forEach(task => {
        if(task.id === updatedTask.id) {
          task.description = updatedTask.description
        }
      })
    })
  },

  DROP_ON_LIST: (state, {listID, movedTask}) => {
    const list = state.lists.find(list => list.id == listID)
    list.tasks.push(movedTask)
  },
  REMOVE_FROM_LIST: (state, {listID, taskID}) => {
    const list = state.lists.find(list => list.id == listID)
    list.tasks = list.tasks.filter(task => task.id != taskID)
  }
}

export default{
  state,
  getters,
  actions,
  mutations
}
