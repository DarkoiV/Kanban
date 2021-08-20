<template>
  <p id="fourohfour"> 404 </p>
  <p id="notfound">Not Found </p>
</template>

<script>
export default {
  name: 'NotFound',
}
</script>

<style scoped>
#fourohfour {
  margin: 30px;
  text-align: center;
  color: salmon;

  font: Monaco, monospace;
  font-size: 40px;
  font-weight: 600;

  user-select: none;
}
#notfound {
  margin: 30px;
  text-align: center;
  color: salmon;

  font: Monaco, monospace;
  font-size: 24px;
  font-weight: 600;

  user-select: none;
}
</style>
