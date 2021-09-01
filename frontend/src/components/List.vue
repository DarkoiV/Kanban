<template>
  <div class="list shadowBox" 
    @drop="onDrop($event, this.id)" 
    @dragenter.prevent 
    @dragover.prevent
    >
    <p id="title"> {{title}} </p>
    <Task 
      @dragstart="startDrag($event, task.id)"
      @dragend="endDrag"
      v-for="task in tasks" 
      :key="task.id"
      :taskObject = "task"
      />
    <span id="newtask" @click="newTask(this.id)"> Create new task </span>
  </div>
</template>

<script>
import Task from './Task.vue'
import { mapActions } from 'vuex'

export default {
  name: 'List',

  components: {
    Task
  },

  methods: {
    ...mapActions(['newTask', 'dropTask']),

    startDrag(event, taskID) {
      event.currentTarget.classList.add('dragged')
      event.dataTransfer.dropEffect = 'move'
      event.dataTransfer.effectAllowed = 'move'
      event.dataTransfer.setData('taskID', taskID)
    },

    endDrag(event) {
      event.currentTarget.classList.remove('dragged')
    },

    onDrop(event, listID) {
      const taskID = event.dataTransfer.getData('taskID')
      this.dropTask({taskID, listID}) 
    },
  },

  props: {
    id: Number,
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

  user-select: none;
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
