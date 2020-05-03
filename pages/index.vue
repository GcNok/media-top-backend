<template>
  <div>
    <RecommendArticles />
    <PopularArticlesRanking />
    <NewArticles />
    <ComparisonArticles />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import RecommendArticles from '~/components/top/RecommendArticles.vue'
import PopularArticlesRanking from '~/components/top/PopularArticlesRanking.vue'
import NewArticles from '~/components/top/NewArticles.vue'
import ComparisonArticles from '~/components/top/ComparisonArticles.vue'

export default Vue.extend({
  components: {
    RecommendArticles,
    PopularArticlesRanking,
    NewArticles,
    ComparisonArticles
  },
  async fetch({ app }) {
    await app.$accessor.getPopularArticles()
    await app.$accessor.getNewArticles()
    await app.$accessor.getRecommendArticles()
    await app.$accessor.getComparisonArticles()
  }
})
</script>
