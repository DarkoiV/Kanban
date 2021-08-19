<template>
  <div class="list shadowBox">
    <p id="title"> {{title}} </p>
    <Task 
      v-for="task in tasks" 
      :key="task.id"
      :taskObject = "task"
      @update-description="updateDescription"
      />
    <p id="newtask" @click="newTask"> Create new task </p>
  </div>
</template>

<script>
import Task from './Task.vue'

export default {
  name: 'List',

  components: {
    Task
  },

  methods: {
    updateDescription(taskID, newDescription) {
      this.$emit('update-description', this.pos, taskID, newDescription);
    },

    newTask() {
      this.$emit("new-task", this.pos);
    }
  },

  props: {
    pos: Number,
    title: String,
    tasks: Array,
  }
}
</script>

<style>
#title {
  font-family: Monaco, monospace;
  font-weight: 500;
  font-size: 20px;

  margin: 5px;
}
.list {
  align-self: flex-start;

  background-color: antiquewhite;
  border-radius: 5px;
  margin: 10px;

  padding: 7px;
  padding-top: 0px;

  min-width: 300px;
  width: 300px;
}
#newtask {
  margin-top: 10px;
  margin-left: 3px;
  user-select: none;

  font-weight: bold;
}
#newtask:hover {
  color: salmon;
  cursor: pointer;
}
</style>
