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
      boardURL: String,
      name: String,
      lists: []
    }
  },
  
  created() {
    if(process.env.NODE_ENV == 'development') {
      this.boardURL = "http://localhost:9000/api/board/" + this.id
    } else {
      this.boardURL = API.URL + "/board/" + this.id
    }

    API.GET(this.boardURL)
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
    newTask(listPos) {
      const newPos = this.lists[listPos].tasks.length;
      const newTask = {
        pos: newPos,
        description: "Double click to edit",
      }
      this.lists[listPos].tasks.push(newTask)
    },

    updateDescription(listPos, taskID, newDescription) {
      const list = this.lists[listPos]
      const task = list.tasks.find(task => task.id === taskID)
      const oldDescription = task.description 
      task.description = newDescription

      const URL = this.boardURL + '/' + list.id + '/' + task.id
      API.PUT(URL, task)
        .catch(err => {
          alert(err)
          task.description = oldDescription
        })
    },
    
    newList() {
      const newList = {
       pos: this.lists.length
      }
      const URL = this.boardURL + '/new'
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
