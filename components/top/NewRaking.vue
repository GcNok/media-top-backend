<template>
  <section>
    <ArticleList
      v-for="(article, index) in articles"
      :key="index"
      class="article-list-item"
      :article="article"
      :rank="index + 1"
    />
    <LinkButton :title="'人気記事ランキング一覧を見る'" :link="'/ranking'" />
  </section>
</template>

<script lang="ts">
import Vue from 'vue'
import { Article } from '~/types/article'
import ArticleList from '~/components/common/ArticleList.vue'
import LinkButton from '~/components/common/LinkButton.vue'

export default Vue.extend({
  name: 'NewRanking',
  components: {
    ArticleList,
    LinkButton
  },
  computed: {
    articles(): Article[] {
      return this.$accessor.articles
    }
  }
})
</script>

<style lang="scss" scoped>
.article-list-item {
  &:last-of-type {
    margin-bottom: 20px;
  }
}
</style>
