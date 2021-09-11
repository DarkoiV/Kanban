<template>
  <div class="board-list shadowBox">
    <p> ------------------------------------</p>
    <div class=flex-container>
      <div
        class="board-element"
        v-for="board in boards" 
        :key="board.id"
        @click="toBoard(board.id)"
      >

        <span class="board-title">  {{ board.name }} </span>

        <span class="delete-board"
          @click="deleteBoard($event, board.id)"
        > 
          <svg width="15px" height="15px" viewBox="0 0 8 8" xmlns="http://www.w3.org/2000/svg">
            <path d="M1.41 0l-1.41 1.41.72.72 1.78 1.81-1.78 1.78-.72.69 1.41 1.44.72-.72 1.81-1.81 1.78 1.81.69.72 1.44-1.44-.72-.69-1.81-1.78 1.81-1.81.72-.72-1.44-1.41-.69.72-1.78 1.78-1.81-1.78-.72-.72z" />
          </svg> 
        </span>

        <span class="edit-board"
          @click="editBoard($event, board.id)"
        >
          <svg version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="15px" height="15px"
          viewBox="0 0 489.627 489.627" xml:space="preserve">
            <path d="M53.569,166.927v270.2h270.2v-118.4c13.4-8.8,25.6-18,37.2-27.9c2.8-2.2,8.3-6.7,15.4-13.5v186c0,14.5-11.8,26.3-26.3,26.3
              h-322.8c-14.5,0-26.3-11.8-26.3-26.3v-322.7c0-14.5,11.8-26.3,26.3-26.3h87.8c-9.6,16.4-17.8,34-24.7,52.6H53.569z M482.269,4.027
              c-9-1.5-20-2.5-32.5-3.2l0,0l-3.2-0.1c-1.9-0.1-3.8-0.2-5.8-0.3l0,0h-0.3l-4.8-0.2h-0.1c-87-2.9-147.6,23-147.6,23l0,0
              c-82.5,33.4-139.3,93.1-164.9,184.2c-8.3,29.5-5.7,60.2-1.5,90.4c0.5,2.8,2,6.9,10.8,1.3l0,0c17.5-47.8,46.7-92.5,81.4-131
              c43.9-48.6,93.2-87.3,154.3-111.6c11.2-4.5,21.9-7.9,33.3-11.8c0.4,1-0.5,2-1.4,2.4c-0.3,0.1-0.7,0.2-1,0.3
              c-32.5,12.5-61.5,31.1-88.5,52.9c-45.5,36.8-83.6,80.4-116.9,128.3c-38.6,55.5-69.6,114.8-89.9,179.4c0,0-3.6,16.1,8.6,4.4
              l53.8-66.9c0.6-1.2,1.4-2.1,2.3-2.8c1.1-1,2.5-1.7,4.2-2.3c24.5-9.2,49.3-17.6,73.1-28.4c32-14.5,61.4-32.6,87.6-55.1
              c2.9-2.2,32-24.9,62.6-71.9c0.5-0.6,1-1.3,1.4-2.1c0.2-0.3,0.4-0.7,0.6-1c7.4-11.7,14.9-24.8,22.2-39.4l-61.7-8.4
              c0,0,69.8-9.9,80.1-32.5c0.6-1.2,1.3-3.3,1.8-4.6c5.4-12.2,10.9-25.7,17-36.9l0,0c1.9-5.8,15.8-30.5,30-45
              C490.669,11.427,487.369,4.827,482.269,4.027z"/>
          </svg>
        </span>

      </div>
    </div>
    <p> ------------------------------------</p>
    <p class="list-new" @click="newBoard"> New Board </p>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  name: 'BoardList',

  computed: {
    ...mapGetters(["boards"])
  },

  created() {
    this.setNavTitle("BOARD LIST")
    this.getBoards()
  },

  methods: {
    ...mapActions(["setNavTitle", "getBoards", "newBoard"]),

    toBoard(id) {
      this.setNavTitle("")
      this.$router.push({ name: 'Board', params: {id: id}})
    },

    deleteBoard(e, id) {
      e.stopPropagation()
      if ( confirm("Do you really want to delete board " + id + "?")) {
        console.log("Deleted lmao")
      }
    },

    editBoard(e, id) {
      e.stopPropagation()
      confirm(id)
    }
  }
}
</script>

<style scoped>
.board-list {
  margin: auto;
  margin-top: 60px;

  font-family: Monaco, monospace;
  font-weight: 900;
  font-size: 20px;

  text-align: center;

  width: 700px;
  background: antiquewhite;
  border-radius: 3px;
  padding: 10px;
}
.flex-container {
  margin: auto;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
}
.board-element {
  margin: auto;
  margin-top: 5px;
  padding-top: 2px;
  padding-bottom: 2px;

  width: 80%;

  cursor: pointer;

  background: seashell;
  box-shadow: inset 3px 3px 3px 0px rgba(0, 0, 15, 0.15);
  border-radius: 3px;
  border: solid 2px seashell;
}
.board-element:hover {
  color: salmon;
  border: solid 2px seashell;

  box-shadow: 3px 3px 3px 0px rgba(0, 0, 15, 0.15);
}
.board-title {
  font: 20px;
  float: left;
  margin-left: 25px;
}
.edit-board {
  float: right;
  margin-right: 10px;
}
.edit-board:hover {
  fill: salmon;
}
.delete-board {
  float:right;
  margin-right: 10px;
}
.delete-board:hover {
  fill: red;
}
.list-new {
  cursor: pointer;
  margin-top: 10px;
  margin-bottom: 10px;
}
.list-new:hover {
  color: salmon;
}
</style>
