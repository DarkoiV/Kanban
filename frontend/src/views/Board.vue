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
import * as API from '../api'

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

    API.GET(this.apiURL)
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
      const list = this.lists[listPos]
      const task = list.tasks.find(task => task.id === taskID)
      const oldDescription = task.description 
      task.description = newDescription

      const URL = this.apiURL + '/' + list.id + '/' + task.id
      API.PUT(URL, task)
        .catch(err => {
          alert(err)
          task.description = oldDescription
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
      const newList = {
       pos: this.lists.length
      }
      const URL = this.apiURL + '/new'
      API.POST(URL, newList)
        .then(data => {
          this.lists.push(data)
        })
        .catch(err => {
          alert(err)
        });
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
