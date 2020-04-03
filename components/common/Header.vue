<template>
  <header class="header">
    <i class="material-icons menu-icon">menu</i>
    <div class="search-bar">
      <i class="material-icons search-icon">search</i>
    </div>
  </header>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'Header'
})
</script>

<style lang="scss">
.header {
  display: flex;
  background-color: $color-black;
  padding: 7px 12px;
  width: 100%;

  .menu-icon {
    font-size: 34px;
    color: white;
  }

  .search-bar {
    display: flex;
    align-items: center;
    margin-left: 10px;
    padding-left: 10px;
    width: 100%;
    background-color: white;
    border-radius: 20px;

    .search-icon {
      font-size: 30px;
    }
  }
}
</style>
