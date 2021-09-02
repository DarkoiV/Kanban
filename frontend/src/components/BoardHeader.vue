<template>
  <header>
    <p id="boardname">
      {{boardName}}
    </p>
  </header>
</template>

<script>
export default {
  name: "BoardHeader",

  props: {
    boardName: String
  },
}
</script>

<style scope>
header {
  position: fixed;
  width: 100vw;

  background: antiquewhite;
  padding: 5px;
}

#boardname {
  font-family: Monaco, monospace;
  font-weight: 600;
  font-size: 22px;

  text-align: center;
}

</style>
