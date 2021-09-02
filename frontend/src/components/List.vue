<template>
  <div class="list shadowBox" 
    @drop="onDrop($event, this.id)" 
    @dragenter.prevent 
    @dragover.prevent
    >
    <p 
      v-if="!editable" 
      id="title"
      @dblclick="editTitle"
    > 
      {{title}} 
    </p>
    <input
      ref="titleInput"
      v-else 
      id="title-in" 
      type="text" 
      v-model="newTitle"
      @keydown.enter.shift.exact.prevent
      @keydown.enter.shift.exact="this.$refs.titleInput.blur()"
      @keydown.escape.exact="this.$refs.titleInput.blur()"
      @blur="saveEdit"
    >

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

  created() {
    this.newTitle = ""
    this.editable = false;
  },

  data() {
    return {
      newTitle: String,
      editable: Boolean
    }
  },

  methods: {
    ...mapActions(['newTask', 'dropTask', 'renameList']),

    editTitle() {
      this.newTitle = this.title
      this.editable = true
      this.$nextTick(() => {
        this.$refs.titleInput.focus()
      })
    },

    saveEdit() {
      if(this.newTitle != this.title) {
        this.renameList({
          listID: this.id,
          newTitle: this.newTitle,
          callback: () => {
            this.editable = false
          }
        })
      } else {
        this.editable = false
      }
    },

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
#title-in {
  font-family: Monaco, monospace;
  font-weight: 500;
  font-size: 20px;

  margin: 0;
  margin-top: 10px;
  margin-bottom: 5px;
  padding: 5px;
  width: calc(100% - 10px);

  border: none;
  border-radius: 3px;
}
#title-in:focus {
  outline: none;
  outline-width: 0;
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
