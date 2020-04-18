<template>
  <aside class="sidebar-wrapper" :class="{ open: isDisplay }">
    <a class="sidebar-content-link" href="https://smashop.jp">お買いもの</a>
    <div class="sidebar-content-title">特集から探す</div>
    <nuxt-link class="sidebar-content-link" to="/popular">人気記事</nuxt-link>
    <nuxt-link class="sidebar-content-link" to="/latest">新着記事</nuxt-link>
    <nuxt-link class="sidebar-content-link" to="/comparison"
      >徹底比較記事</nuxt-link
    >
    <div class="sidebar-content-title">カテゴリから探す</div>
    <nuxt-link class="sidebar-content-link" to="/latest"
      >チルド食品・惣菜</nuxt-link
    >
    <nuxt-link class="sidebar-content-link" to="/latest">乳製品・卵</nuxt-link>
    <nuxt-link class="sidebar-content-link" to="/latest"
      >ドリンク・飲料</nuxt-link
    >
  </aside>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'Sidebar',
  computed: {
    isDisplay() {
      return this.$accessor.isDisplaySidebar
    }
  }
})
</script>

<style lang="scss" scoped>
@include media-query($sp) {
  .sidebar-wrapper {
    position: fixed;
    top: 0;
    left: responsive-width(-400);
    width: responsive-width(200);
    height: 100%;
    background-color: white;
    border-right: 1px solid $color-gray;
    z-index: 200;
    transition: left 100ms ease-in;

    .sidebar-content {
      display: flex;
      align-items: center;
      color: black;
      text-decoration: none;
      border-bottom: 1px solid $color-gray;
    }
    .sidebar-content-title {
      @extend .sidebar-content;
      padding-left: 10px;
      height: 30px;
      background-color: $color-light-orange;
      font-weight: bold;
      color: white;
    }
    .sidebar-content-link {
      @extend .sidebar-content;
      padding-left: 20px;
      height: 40px;
    }
  }
}

@include media-query($pc) {
  .sidebar-wrapper {
    min-width: 400px;
    height: 100vh;
  }
}
</style>
