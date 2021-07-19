<template>
  <div class="task">
    <p> {{description}} </p>
    <p class="due"> {{ due }} </p>
  </div>
</template>

<script>

export default {
  name: 'Task',

  props: {
    description: String,
    due: Date,
  },
}
</script>

<style>
.task {
  background-color: white;
  border-radius: 2px;

  padding: 5px;
  padding-top: 10px;
  margin-bottom: 5px;
}
.due {
  font-size: 70%;
  text-align: right;
}
</style>
