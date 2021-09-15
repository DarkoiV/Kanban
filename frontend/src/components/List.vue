<template>
  <div class="list shadowBox" 
    @drop="onDrop($event, this.id)" 
    @dragenter.prevent
    @dragover="dragOver"
    @dragleave="dragLeave"
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
      maxlength="29"
      @keydown.enter.shift.exact.prevent
      @keydown.enter.shift.exact="this.$refs.titleInput.blur()"
      @keydown.enter.exact.prevent
      @keydown.enter.exact="this.$refs.titleInput.blur()"
      @keydown.escape.exact="this.$refs.titleInput.blur()"
      @blur="saveEdit"
    >

    <Task 
      @dragstart="startDrag($event, task.id)"
      @dragend="endDrag"
      @dragenter.prevent
      @dragover.prevent
      v-for="task in sortedTasks" 
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
    this.drgOver = false;
  },

  data() {
    return {
      newTitle: String,
      editable: Boolean,
      drgOver: Boolean,
      ghostPos: Number,
      leaveTicks: Number
    }
  },

  computed: {
    sortedTasks: function() {
      let sortedTasks = this.tasks

      if(this.drgOver) {
        sortedTasks = [
          ...sortedTasks.filter( task => task.pos < this.ghostPos), 
          {id: -1},
          ...sortedTasks.filter( task => task.pos >= this.ghostPos)
        ]
      } else {
        sortedTasks = sortedTasks.filter( task => task.id != -1 )
      }
      
      return sortedTasks
    }
  },

  methods: {
    ...mapActions(['newTask', 'dropTask', 'renameList', 'deleteList']),

    editTitle() {
      this.newTitle = this.title
      this.editable = true
      this.$nextTick(() => {
        this.$refs.titleInput.focus()
      })
    },

    saveEdit() {
      if (this.newTitle === "") {
        if (confirm(`Do you really want to delete ${this.title}?`)) {
          this.deleteList(this.id)
        } else {
          this.editable = false
        }
      } else if(this.newTitle != this.title) {
        this.renameList({
          listID: this.id,
          newTitle: this.newTitle,
          callback: () => {this.editable = false}
        })
      } else {
        this.editable = false
      }
    },

    startDrag(event, taskID) {
      event.dataTransfer.dropEffect = 'move'
      event.dataTransfer.effectAllowed = 'move'
      event.dataTransfer.setData('taskID', taskID)
      event.currentTarget.classList.add('dragged')
    },

    endDrag(event) {
      event.currentTarget.classList.remove('dragged')
    },

    dragOver(event) {
      this.leaveTicks = 0;
      this.drgOver = true;
      event.preventDefault()
      let elements = event.currentTarget.children
      this.ghostPos = 0
      elements.forEach(element => {
        if(element.classList.contains('task')) {
          const elementY = element.getBoundingClientRect().y
          const mouseY = event.clientY
          if(mouseY > (elementY + 100)) {
            this.ghostPos++;
          }
        }
      })
    },

    dragLeave(event) {
      this.leaveTicks = 1;
      event.preventDefault()
      setTimeout(() => {
        if(this.leaveTicks == 1){
            this.drgOver = false;
        }
      }, 100)

    },

    onDrop(event, listID) {
      this.drgOver = false;
      const taskID = event.dataTransfer.getData('taskID')
      const ghostPos = this.ghostPos
      this.dropTask({taskID, listID, ghostPos}) 
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
