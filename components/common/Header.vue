<template>
  <header class="header">
    <i class="material-icons header-icon">menu</i>
    <img class="header-logo" src="/img/ss-logo.png" alt="ss-logo" />
    <i class="material-icons header-icon">search</i>
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
  justify-content: space-between;
  align-items: center;
  background-color: $color-black;
  padding: 7px 12px;
  width: 100%;

  .header-icon {
    font-size: 34px;
    color: white;
  }

  .header-logo {
    width: 180px;
  }
}
</style>
