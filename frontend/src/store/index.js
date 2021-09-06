import { createStore } from 'vuex'
import board from './modules/board'
import navbar from './modules/navbar'

const store = createStore({
  modules: {
    board,
    navbar
  }
})

export default store
