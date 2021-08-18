<template>  
    <List v-for="taskList in boardLists"
      :pos="taskList.pos"
      :title="taskList.title"
      :key="taskList.id"
      :tasks="taskList.tasks" 
      @update-description="updateDescription"
    />
</template>

<script>
import List from './List.vue'

export default {
  name: 'Board',

  components: {
    List
  },

  data() {
    return { 
      boardLists: []
    }
  },
  
  // TMP populate multiple task list with predefined tasks
  created() {
    this.boardLists = [
      {
        pos: 0,
        title: "TO DO",
        tasks: [
          {
            id: 1,
            pos: 0,
            description: "First task!",
            due: "2020/01/01"
          },
          { 
            id: 2,
            pos: 1,
            description: "Second task!",
            due: "2020/01/01"
          },
          {
            id: 3,
            pos: 2,
            description: "LOREM IPSUM, I DO NOT REMEMBER FURTHER!!!",
            due: "2021/08/07"
          },
          {
            id: 4,
            pos: 3,
            description: "Form is nice \n I LIKE FORM!",
            due: "2021/08/07"
          }
        ]
      },
      {
        pos: 1,
        title: "In Progress",
        tasks: [
          {
            id: 1,
            pos: 0,
            description: "First task!",
            due: "2020/01/01"
          },
          { 
            id: 2,
            pos: 2,
            description: "Second task!",
            due: "2020/01/01"
          },
          {
            id: 3,
            pos: 1,
            description: "LOREM IPSUM, I DO NOT REMEMBER FURTHER!!!",
            due: "2021/08/07"
          },
          {
            id: 4,
            pos: 3,
            description: "Form is nice \n I LIKE FORM!",
            due: "2021/08/07"
          }
        ]
      }
    ]

    // Sort tasks
    this.boardLists.forEach(list => {
      list.tasks.sort( (t1, t2) => t1.pos - t2.pos)
    })
  },

  methods: {
    // Update description of task
    updateDescription(listPos, taskID, newDescription) {
      this.boardLists[listPos].tasks.map(task =>{
        if(task.id == taskID){
          task.description = newDescription;
        }
        return task;
      })
    }

  },
}
</script>
