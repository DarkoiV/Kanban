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
      apiURL: String,
      name: String,
      lists: []
    }
  },
  
  created() {
    this.apiURL = location.protocol + "//" + location.host + "/api/board/" + this.id

    fetch(this.apiURL).then(response => {
      if (!response.ok) {
        throw("HTTP status when requesting board data: " + response.status);
      }

      return response.json()
    })
    .then(data => {
      this.name = data.name
      this.lists = data.lists

      // Sort tasks
      this.lists.forEach(list => {
        list.tasks.sort( (t1, t2) => t1.pos - t2.pos)
      })
    })
    .catch(err => {
      alert(err)
      this.$router.push('/notfound')
    });

  },

  methods: {
    updateDescription(listPos, taskID, newDescription) {
      const tasks = this.lists[listPos].tasks

      tasks.map(task => {
        if (task.id == taskID) {
          task.description = newDescription;

          const listID = this.lists[listPos].id
          const taskURL = this.apiURL + "/" + listID + "/" + taskID
          const request = {
            method: "PUT",
            body: JSON.stringify(task)
          }
          
          fetch(taskURL, request)
          console.log(task)
        }
        return task;
      })
    },
    
    newTask(listPos) {
      //TODO get ID from server
      const newPos = this.lists[listPos].tasks.length;

      const newTask = {
        pos: newPos,
        description: "Double click to edit",
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
