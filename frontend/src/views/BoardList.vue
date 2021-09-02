<template>
  <div class="board-list shadowBox">
    <p class="list-title"> LIST OF BOARDS </p>
    <div
      class="board-element"
      v-for="board in boards" 
      :key="board.id"
      @click="redirect(board.id)"
      >
      <p>  {{ board.title }} </p>
    </div>
    <p class="list-new"> Create new board </p>
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
      ({title: "Second Board", id: 2})
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
.board-element {
  cursor: pointer;
}
.board-element:hover {
  color: salmon;
}
.list-new {
  cursor: pointer;
  margin-top: 10px;
}
.list-new:hover {
  color: salmon;
}
</style>
