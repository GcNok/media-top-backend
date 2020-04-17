<template>
  <div class="list-page-title-wrapper">
    <h1 class="list-page-title">{{ title }}</h1>
    <h2 class="list-page-sub-title">
      {{ subTitle }}
    </h2>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'ListPageTitle',
  props: {
    title: {
      type: String,
      default: ''
    },
    subTitle: {
      type: String,
      default: ''
    }
  }
})
</script>

<style lang="scss" scoped>
.list-page-title-wrapper {
  margin: 20px;
  .list-page-title {
    font-size: responsive-width(20);
  }

  .list-page-sub-title {
    font-weight: 600;
    color: gray;
    font-size: responsive-width(12);
  }
}
</style>
