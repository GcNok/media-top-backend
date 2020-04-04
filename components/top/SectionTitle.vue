<template>
  <section class="section-wrapper">
    <h2 class="section-title">
      {{ section.title }}
    </h2>
    <h3 class="section-comment">
      {{ section.comment }}
    </h3>
  </section>
</template>

<script lang="ts">
import Vue from 'vue'

type Section = {
  title: string
  comment: string
}

export default Vue.extend({
  name: 'SectionTitle',
  props: {
    section: {
      type: Object as () => Section,
      default: {}
    }
  }
})
</script>

<style lang="scss">
.section-wrapper {
  padding: 16px 16px 3px;
  border-bottom: 4px solid orange;

  .section-title {
    font-size: 20px;
  }

  .section-comment {
    font-weight: 600;
    color: gray;
    font-size: 12px;
  }
}
</style>
