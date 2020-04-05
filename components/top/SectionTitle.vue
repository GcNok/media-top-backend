<template>
  <div class="section-title-wrapper">
    <h2 class="section-title">
      {{ title }}
    </h2>
    <h3 class="section-comment">
      {{ comment }}
    </h3>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'SectionTitle',
  props: {
    title: {
      type: String,
      default: ''
    },
    comment: {
      type: String,
      default: ''
    }
  }
})
</script>

<style lang="scss">
.section-title-wrapper {
  padding: 0 responsive-height(16) responsive-width(3);
  border-bottom: 4px solid orange;

  .section-title {
    font-size: responsive-width(20);
  }

  .section-comment {
    font-weight: 600;
    color: gray;
    font-size: responsive-width(12);
  }
}
@include media-query($pc) {
  .section-title-wrapper {
    padding: 0 16px 3px;
    border-bottom: 4px solid orange;

    .section-title {
      font-size: 28px;
    }

    .section-comment {
      font-size: 16px;
    }
  }
}
</style>
