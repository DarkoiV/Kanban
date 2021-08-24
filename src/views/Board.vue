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
  
  created() {
    this.name = "Name of Board " + this.id
    const boardDataUrl = "http://" + location.host + "/api/" + this.id
    console.log(boardDataUrl)
    fetch(boardDataUrl)
      .then(response => response.json())
      .then(data => {
        this.name = data.name
        this.lists = data.lists

        // Sort tasks
        this.lists.forEach(list => {
          list.tasks.sort( (t1, t2) => t1.pos - t2.pos)
        })
      })
      .catch(err => console.log(err))

  },

  methods: {
    updateDescription(listPos, taskID, newDescription) {
      const tasks = this.lists[listPos].tasks

      tasks.map(task => {
        if (task.id == taskID) {
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
