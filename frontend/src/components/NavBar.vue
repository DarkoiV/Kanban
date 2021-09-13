<template>
  <div class="navigation">
    <div class="left">
      <ul class="links">
        <li @click="goBoardList"> Boards list </li>
        <li @click="goCurrentBoard" :class="{disable: !onBoard}"> Current board </li>
      </ul>
    </div>

    <div class="center">
      {{ navTitle }}
    </div>

    <div class="right">
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  name: "NavBar",

  computed: {
    ...mapGetters(['navTitle', 'boardID']),
    
    onBoard () {
      return typeof this.boardID === 'number'
    }
  },

  methods: {
    ...mapActions(['setNavTitle']),

    goBoardList() {
      this.$router.push({ name: 'BoardList'})
    },

    goCurrentBoard() {
      if(this.onBoard) {
        this.$router.push({ name: 'Board', params: {id: this.boardID} })
      }
    }
  },

  created() {
    this.setNavTitle("")
  }

}
</script>

<style scoped>
.navigation {
  display: flex;
  align-content: center;

  width: 100%;

  position: fixed;
  top: 0;

  background: antiquewhite;
  font-family: Monaco, monospace;
  font-weight: 900;
  font-size: 24px;

  user-select: none;
}
.right {
  justify-self: right;
  margin: auto;
  width: 500px
}
.links {
  list-style-type: none;
} 
li {
  font-size: 16px;
  font-family: Avenir, Helvetica, Arial, sans-serif;

  display: block;
  width: 120px;
  padding-top: 10px;
  padding-bottom: 5px;
  text-align: center;
  background: seashell;

  margin-bottom: 5px;
  margin-left: 5px;
  float: left;
  cursor: pointer;
}
li:hover:not(.disable) {
  color: salmon;
  box-shadow: 3px 3px 3px 0px rgba(0, 0, 15, 0.15);
  transform: translateY(5px);
}
li.disable {
  color: grey;
  cursor: default;
}
.center {
  justify-self: center;
  margin: auto;
  width: 200px;
  text-align: center;
}
.left {
  justify-self: left;
  margin: auto;
  width: 500px
}
</style>
