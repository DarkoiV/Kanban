<template>
  <div class="main-body">
    <List title="LIST ONE" :tasks="taskList" />
  </div>
</template>

<script>
import List from './components/List.vue'

export default {
  name: 'App',

  components: {
    List
  },

  // Store task list(s)
  data() {
    return { 
      taskList: []
    }
  },
  
  // TMP populate task list with predefined tasks
  created() {
    this.taskList = [
      {
        id: 1,
        description: "First task!",
        due: "2020/01/01"
      },
      { 
        id: 2,
        description: "Second task!",
        due: "2020/01/01"
      },
      {
        id: 3,
        description: "LOREM IPSUM, I DO NOT REMEMBER FURTHER!!!",
        due: "2021/08/07"
      }
    ]
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
