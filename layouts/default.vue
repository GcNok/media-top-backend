<template>
  <div>
    <Header />
    <div class="main-wrapper">
      <aside class="sidebar" />
      <nuxt class="page-content" />
    </div>
    <Footer />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Header from '@/components/common/Header.vue'
import Footer from '@/components/common/Footer.vue'
export default Vue.extend({
  name: 'Default',
  components: {
    Header,
    Footer
  }
})
</script>

<style lang="scss">
*,
*:before,
*:after {
  box-sizing: border-box;
  margin: 0;
}

@include media-query($pc) {
  .main-wrapper {
    display: flex;

    .sidebar {
      min-width: 400px;
      height: 100vh;
    }

    .page-content {
      max-width: 900px;
    }
  }
}
</style>
