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

  async renameList({ commit }, {listID, newTitle, callback}) {
    try {
      const editedList = {
        id: listID,
        title: newTitle
      }
      await API.PATCH(`${API.URL}/board/${state.id}/${listID}`, editedList)
      commit("RENAME_LIST", {listID, newTitle})
    } catch(err) {
      alert("Error when renaming list ", + err)
    }
    callback()
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

  async dropTask({ commit }, { taskID, listID, ghostPos }) {
    try {
      const newList = state.lists.find(list => listID == list.id)
      const oldList = state.lists.find(list =>
        list.tasks.find(task => taskID == task.id)
      )
      let movedTask = oldList.tasks.find(task => task.id == taskID)

      if (newList == oldList && movedTask.pos == ghostPos) { return }
      if (newList == oldList && movedTask.pos < ghostPos) {
        ghostPos--;
      }

      commit('REMOVE_FROM_LIST', {listID: oldList.id, taskID,})
      commit('DROP_ON_LIST', {listID, movedTask, ghostPos})

      API.PATCH(`${API.URL}/board/${state.id}/${oldList.id}`, oldList)
      await API.PATCH(`${API.URL}/board/${state.id}/${newList.id}`, newList)

    }
    catch(err) {
      alert("Error when moving task " + err)
      router.go()
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
    state.lists.sort( (a, b) => a.pos - b.pos )
    state.lists.forEach( list => 
      list.tasks.sort( (a, b) => a.pos - b.pos )
    )
  },

  NEW_LIST: (state, newList) => state.lists.push(newList),
  NEW_TASK: (state, {listID, newTask}) => {
    const list = state.lists.find( list => listID === list.id )
    list.tasks.push(newTask)
  },

  RENAME_LIST: (state, {listID, newTitle}) => {
    const list = state.lists.find(list => list.id == listID)
    list.title = newTitle
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

  DROP_ON_LIST: (state, {listID, movedTask, ghostPos}) => {
    const list = state.lists.find(list => list.id == listID)

    list.tasks = [ 
      ...list.tasks.filter(task => task.pos < ghostPos),
      movedTask,
      ...list.tasks.filter(task => task.pos >= ghostPos)
    ]
    list.tasks.map((task, index) => task.pos = index )
  },
  REMOVE_FROM_LIST: (state, {listID, taskID}) => {
    const list = state.lists.find(list => list.id == listID)
    list.tasks = list.tasks.filter(task => task.id != taskID)
    list.tasks.map((task, index) => task.pos = index ) 
  }
}

export default{
  state,
  getters,
  actions,
  mutations
}
