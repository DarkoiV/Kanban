import { createStore } from 'vuex'
import board from './modules/board'

const store = createStore({
  modules: {
    board
  }
})

export default store
