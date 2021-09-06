<template>
  <p id="fourohfour"> 404 </p>
</template>

<script>
import { mapActions } from 'vuex'
export default {
  name: 'NotFound',

  methods: mapActions(['setNavTitle']),

  created() {
    this.setNavTitle('Not found')
  }
}
</script>

<style scoped>
#fourohfour {
  margin: 100px;
  text-align: center;
  color: salmon;

  font: Monaco, monospace;
  font-size: 60px;
  font-weight: 900;

  user-select: none;
}
</style>
