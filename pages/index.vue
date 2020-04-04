<template>
  <main>
    <RecommendArticles />
    <PopularArticlesRanking />
    <NewArticles />
  </main>
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
  }
})
</script>
