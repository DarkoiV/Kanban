<template>
  <div class="task">
    <p 
      style="white-space: pre-wrap;"
      v-if="editable" 
      @dblclick="editTask()"
    > 
      {{ taskObject.description }} 
    </p>

    <form v-else class="task-form"> 
      <textarea 
        v-model="formEdit"
        @keydown.enter.exact.prevent
        @keydown.enter.exact="saveEdit()"
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
    },
    // Send edited task to parent
    saveEdit() {
      this.editable = !this.editable;
      console.log("TASK:", this.taskObject.id, this.formEdit);
      this.$emit('update-description', this.taskObject.id, this.formEdit)
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
.due {
  font-size: 70%;
  text-align: right;
}
</style>
