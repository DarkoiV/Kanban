<template>
  <div class="task" ref="taskContainer" draggable="true" @dblclick="editTask">
    <p
      class="cardText"
      v-show="editable" 
    > 
      {{ taskObject.description }}
    </p>

    <form v-show="!editable" class="task-form"> 
      <textarea 
        v-model="formEdit"
        ref="textInput"
        @keydown.enter.shift.exact.prevent
        @keydown.enter.shift.exact="saveEdit"
        @keydown.escape.exact="saveEdit"
        @blur="saveEdit"
        @input="resize"
      /> 
    </form>

    <p class="due"> {{ taskObject.due }} </p>
  </div>
</template>

<script>

export default {
  name: 'Task',

  props: {
    taskObject: Object,
  },

  data() {
    return {
      editable: Boolean,
      formEdit: String,
    }
  },

  methods: {
    // Start editing task
    editTask() {
      this.formEdit = this.taskObject.description;
      this.editable = !this.editable;
      this.$refs.taskContainer.draggable = false;

      this.$nextTick(() => {
        // Resize it
        this.$refs.textInput.height = "auto"
        this.$refs.textInput.height = `${this.$refs.textInput.scrollHeight}px`

        // Grab input
        this.$refs.textInput.focus()
      })
    },

    // Send edited task to parent
    saveEdit() {
      this.editable = !this.editable;
      this.$refs.taskContainer.draggable = true;

      // Check if edition changed anything
      if(this.formEdit != this.taskObject.description) {
        console.log("TASK:", this.taskObject.id, this.formEdit);
        this.$emit('update-description', this.taskObject.id, this.formEdit)
      } else {
        console.log("Task have not changed")
      }
    },

    // Resize text area
    resize(e) {
      e.target.style.height = "auto"
      e.target.style.height = `${e.target.scrollHeight}px`
      console.log(e.target.scrollHeight)
    }
  },
}
</script>

<style scope>
.task {
  display: flex;
  flex-direction: column;

  background-color: white;
  border-radius: 2px;
  font-size: 14px;
  min-height: 75px;

  padding: 5px;
  padding-top: 10px;
  margin-bottom: 5px;
}
textarea {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  font-size: 14px;
  font-weight: 100;
  resize: none; 

  width: 100%;
  min-height: 60px;
  margin: 0;
  padding: 0;

  border: none;
}
textarea:focus {
  outline: none;
  outline-width: 0;
}
.cardText {
  width: 100%;
  white-space: pre-wrap;
  word-break: break-word;
}
.due {
  font-size: 70%;
  text-align: right;
  margin-top: auto;
}
</style>
