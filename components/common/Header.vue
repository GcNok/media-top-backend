<template>
  <header class="header">
    <i class="material-icons header-icon" @click="toggleSidebar()">menu</i>
    <nuxt-link class="header-link" to="/"
      ><img class="header-logo" src="/contents/img/ss-logo.png" alt="ss-logo"
    /></nuxt-link>
  </header>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'Header',
  methods: {
    toggleSidebar(): void {
      this.$accessor.toggleSidebar()
    }
  }
})
</script>

<style lang="scss">
.header {
  position: sticky;
  top: 0;
  display: flex;
  align-items: center;
  padding: responsive-height(7) responsive-width(12);
  width: 100%;
  background-color: $color-black;
  z-index: 200;

  @include media-query($pc) {
    padding: 7px 12px;
  }

  .header-icon {
    font-size: responsive-width(34);
    color: white;
  }
  .header-link {
    display: flex;
    .header-logo {
      margin-left: responsive-width(20);
      width: responsive-width(180);
    }
  }
}

@include media-query($pc) {
  .header {
    padding: 7px 12px;
    height: 50px;

    .header-icon {
      display: none;
      font-size: 34px;
    }
    .header-link {
      .header-logo {
        margin: 0 auto 0 30px;
        width: 220px;
      }
    }
  }
}
</style>
