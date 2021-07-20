<template>
  <div class="task">
    <p 
      style="white-space: pre-wrap;"
      v-if="editable" 
      @dblclick="editTask()"
    > 
      {{description}} 
    </p>

    <form v-else class="task-form"> 
      <textarea 
        v-model="formEdit"
        @keydown.enter.exact.prevent
        @keydown.enter.exact="saveEdit()"
      /> 
    </form>

    <p class="due"> {{ due }} </p>
  </div>
</template>

<script>

export default {
  name: 'Task',

  props: {
    id: Number,
    description: String,
    due: Date,
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
      this.formEdit = this.description;
      this.editable = !this.editable;
    },
    // Send edited task to parent
    saveEdit() {
      this.editable = !this.editable;
      console.log("TASK:", this.id, this.formEdit);
      this.$emit('update-description', this.id, this.formEdit)
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
  font-size: 14px;
  resize: vertical; 

  max-width: calc(100% - 10px);
  margin: 0;
}
.due {
  font-size: 70%;
  text-align: right;
}
</style>
