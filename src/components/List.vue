<template>
  <div class="list">
    <p id="title"> {{title}} </p>
    <Task 
      v-for="task in tasks" 
      :key="task.id"
      :taskObject = "task"
      @update-description="updateDescription"
      />
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
    updateDescription(id, formEdit){
      this.$emit('update-description', this.pos, id, formEdit);
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

  margin: 0px;
}
.list {
  align-self: flex-start;

  background-color: antiquewhite;
  border-radius: 3px;
  margin: 5px;

  padding: 5px;
  padding-top: 0px;
  padding-bottom: 10px;

  width: 300px;
  overflow: auto;
}
</style>
