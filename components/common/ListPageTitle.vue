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
  margin: responsive-height(20) responsive-width(20);
  .list-page-title {
    font-size: responsive-width(20);
  }

  .list-page-sub-title {
    font-weight: 600;
    color: gray;
    font-size: responsive-width(12);
  }
}

@include media-query($pc) {
  .list-page-title-wrapper {
    padding: 0 16px 3px;

    .list-page-title {
      font-size: 28px;
    }

    .list-page-sub-title {
      font-size: 16px;
    }
  }
}
</style>
