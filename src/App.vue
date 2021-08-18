<template>
  <div class="main-body">
    <Board/>
  </div>
</template>

<script>
import Board from './components/Board.vue'

export default {
  name: 'App',

  components: {
    Board
  },

}
</script>

<style>
* {
  margin: 0px;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  padding: 30px;
}
.main-body {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
}
</style>
