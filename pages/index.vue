<template>
  <div>
    <RecommendArticles />
    <PopularArticlesRanking />
    <NewArticles />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import RecommendArticles from '~/components/top/RecommendArticles.vue'
import PopularArticlesRanking from '~/components/top/PopularArticlesRanking.vue'
import NewArticles from '~/components/top/NewArticles.vue'

export default Vue.extend({
  components: {
    RecommendArticles,
    PopularArticlesRanking,
    NewArticles
  },
  async fetch({ app }) {
    await app.$accessor.getPopularArticles()
    await app.$accessor.getNewArticles()
  }
})
</script>
