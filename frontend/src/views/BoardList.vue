<template>
  <div class="board-list shadowBox">
    <p class="list-title"> LIST OF BOARDS </p>
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
  },

  methods: {
    redirect(id) {
      this.$router.push({ name: 'Board', params: {id: id}})
    }
  }
}
</script>

<style scoped>
.board-list {
  margin: auto;
  margin-top: 5vh;

  font-family: Monaco, monospace;
  font-weight: 900;
  font-size: 20px;

  text-align: center;

  width: 700px;
  background: antiquewhite;
  border-radius: 3px;
  padding: 10px;
}
.list-title {
  font-size: 24px;
  margin-bottom: 10px;
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
  border: solid 2px white;
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
