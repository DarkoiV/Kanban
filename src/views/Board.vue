<template>
  <BoardHeader :boardName="name"/> 
  <div class="main-body">
    <List v-for="list in lists"
      :pos="list.pos"
      :title="list.title"
      :key="list.id"
      :tasks="list.tasks" 
      @update-description="updateDescription"
      @new-task="newTask"
    />
    <div>
    <p id="newlist" @click="newList"> Create new list </p>
    </div>
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
            createdAt: new Date()
          },
          { 
            id: 2,
            pos: 1,
            description: "Second task!",
            createdAt: new Date()
          },
          {
            id: 3,
            pos: 2,
            description: "LOREM IPSUM, I DO NOT REMEMBER FURTHER!!!",
            createdAt: new Date()
          },
          {
            id: 4,
            pos: 3,
            description: "Form is nice \n I LIKE FORM!",
            createdAt: new Date()
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
            createdAt: new Date()
          },
          { 
            id: 2,
            pos: 2,
            description: "Second task!",
            createdAt: new Date()
          },
          {
            id: 3,
            pos: 1,
            description: "LOREM IPSUM, I DO NOT REMEMBER FURTHER!!!",
            createdAt: new Date()
          },
          {
            id: 4,
            pos: 3,
            description: "Form is nice \n I LIKE FORM!",
            createdAt: new Date()
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
    updateDescription(listPos, taskID, newDescription) {
      this.lists[listPos].tasks.map(task =>{
        if(task.id == taskID){
          task.description = newDescription;
        }
        return task;
      })
    },
    
    newTask(listPos) {
      //TODO get ID from server
      const newPos = this.lists[listPos].tasks.length;
      const newID = newPos + listPos * 1000;

      const newTask = {
        id: newID,
        pos: newPos,
        description: "New task, double click to edit",
        createdAt: new Date()
      }
      console.log("New task", newTask)
      this.lists[listPos].tasks.push(newTask)
    },

    newList() {
      const newPos = this.lists.length
      const newID = newPos

      const newList = {
        id: newID,
        pos: newPos,
        title: "New List",
        tasks: []
      }

      console.log("New list", newList)
      this.lists.push(newList)
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
  padding-top: 50px;
}
#newlist {
  font-family: Monaco, monospace;
  font-weight: 500;
  font-size: 20px;

  min-width: 300px;
  margin: 15px;
}
#newlist:hover {
  color: salmon;
  cursor: pointer;
}
</style>
