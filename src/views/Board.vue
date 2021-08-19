<template>
  <BoardHeader :boardName="name"/> 
  <div class="main-body">
    <List v-for="list in lists"
      :pos="list.pos"
      :title="list.title"
      :key="list.id"
      :tasks="list.tasks" 
      @update-description="updateDescription"
    />
  </div>
</template>

<script>
import List from '../components/List.vue'
import BoardHeader from '../components/BoardHeader.vue'

export default {
  name: 'Board',

  props: {
    id: Number,
  },

  components: {
    List,
    BoardHeader
  },

  data() {
    return { 
      name: String,
      lists: []
    }
  },
  
  // TMP populate multiple task list with predefined tasks
  created() {
    this.name = "Name of Board " + this.id
    this.lists = [
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
    this.lists.forEach(list => {
      list.tasks.sort( (t1, t2) => t1.pos - t2.pos)
    })
  },

  methods: {
    // Update description of task
    updateDescription(listPos, taskID, newDescription) {
      this.lists[listPos].tasks.map(task =>{
        if(task.id == taskID){
          task.description = newDescription;
        }
        return task;
      })
    }

  },
}
</script>

<style scope>
.main-body {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  padding: 25px;
}
</style>
