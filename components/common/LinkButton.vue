<template>
  <nuxt-link class="link-button" :to="link">{{ title }}</nuxt-link>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'LinkButton',
  props: {
    title: {
      type: String,
      default: ''
    },
    link: {
      type: String,
      default: ''
    }
  }
})
</script>

<style lang="scss" scoped>
.link-button {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0 auto;
  width: 250px;
  height: 50px;
  background-color: orange;
  color: white;
  font-size: 15px;
  border-radius: 10px;
  box-shadow: 0px 3px 6px 1px grey;
  text-decoration: none;
}
</style>
