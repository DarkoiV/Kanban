import { createStore } from 'vuex'
import board from './modules/board'
import boardList from './modules/boardList'
import navbar from './modules/navbar'

const store = createStore({
  modules: {
    board,
    boardList,
    navbar
  }
})

export default store
