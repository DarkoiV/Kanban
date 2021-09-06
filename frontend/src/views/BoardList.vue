<template>
  <div class="board-list shadowBox">
    <p> ------------------------------------</p>
    <div class=flex-container>
      <div
        class="board-element"
        v-for="board in boards" 
        :key="board.id"
        @click="redirect(board.id)"
        >
        <p class="board-title">  {{ board.title }} </p>
      </div>
    </div>
    <p> ------------------------------------</p>
    <p class="list-new"> New Board </p>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
export default {
  name: 'BoardList',

  data() {
    return {
      boards: []
    }
  },

  created() {
    this.boards = [
      ({title: "First Board", id: 1}),
      ({title: "Second Board", id: 2}),
      ({title: "Third Board", id: 3})
    ]
    this.setNavTitle("BOARD LIST")
  },

  methods: {
    ...mapActions(["setNavTitle"]),

    redirect(id) {
      this.setNavTitle("")
      this.$router.push({ name: 'Board', params: {id: id}})
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
  position: relative;
  top: 10%;
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
