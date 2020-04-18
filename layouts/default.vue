<template>
  <div class="wrapper" :class="{ open: isDisplaySidebar }">
    <Header />
    <div class="main-wrapper">
      <Sidebar />
      <main class="page-content" @click="closeSidebar()">
        <nuxt />
      </main>
    </div>
    <Footer />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Header from '@/components/common/Header.vue'
import Sidebar from '@/components/common/Sidebar.vue'
import Footer from '@/components/common/Footer.vue'
export default Vue.extend({
  name: 'Default',
  components: {
    Header,
    Sidebar,
    Footer
  },
  computed: {
    isDisplaySidebar(): boolean {
      return this.$accessor.isDisplaySidebar
    }
  },
  methods: {
    closeSidebar(): void {
      this.$accessor.closeSidebar()
    }
  }
})
</script>

<style lang="scss">
*,
*:before,
*:after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
@include media-query($sp) {
  .wrapper {
    transform: translateX(0);
    transition: transform 100ms ease-in;
  }
  .open {
    transform: translateX(responsive-width(200));
  }
}

.main-wrapper {
  display: flex;
}

@include media-query($pc) {
  .main-wrapper {
    justify-content: center;
    .sidebar-wrapper {
      min-width: 400px;
      height: 100vh;
    }

    .page-content {
      max-width: 900px;
    }
  }
}
</style>
