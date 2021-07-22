<template>
  <div class="task">
    <p
      class="cardText"
      v-show="editable" 
      @dblclick="editTask"
    > 
      {{ taskObject.description }} 
    </p>

    <form v-show="!editable" class="task-form"> 
      <textarea 
        v-model="formEdit"
        ref="textInput"
        @keydown.enter.exact.prevent
        @keydown.enter.exact="saveEdit"
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
      console.log("TASK:", this.taskObject.id, this.formEdit);
      this.$emit('update-description', this.taskObject.id, this.formEdit)
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
  background-color: white;
  border-radius: 2px;
  font-size: 14px;

  padding: 5px;
  padding-top: 10px;
  margin-bottom: 5px;
}
textarea {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  font-size: 14px;
  resize: none; 

  width: calc(100% - 12px);
  margin: 0;

  border: none;
}
.cardText {
  width: calc(100% - 12px);
  white-space: pre-wrap;
  word-break: break-all;
}
.due {
  font-size: 70%;
  text-align: right;
}
</style>
