<template>
  <footer class="footer">©SmartShopping, Inc.</footer>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'Footer'
})
</script>

<style lang="scss">
.footer {
  display: flex;
  justify-content: center;
  align-items: center;
  height: responsive-height(50);
  background-color: #f4f4f4;

  @include media-query($pc) {
    height: 50px;
  }
}
</style>
