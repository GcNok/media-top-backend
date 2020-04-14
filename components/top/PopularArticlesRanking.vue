<template>
  <section class="popular-articles-wrapper">
    <SectionTitle
      :title="'人気記事ランキング'"
      :comment="'繰り返し読みたくなる暮らしに役立つ記事をご紹介'"
    />
    <ArticleRow />
    <LinkButton :title="'人気記事ランキング一覧を見る'" :link="'/ranking'" />
  </section>
</template>

<script lang="ts">
import Vue from 'vue'
import { Article } from '~/types/article'
import ArticleRow from '~/components/common/ArticleRow.vue'
import LinkButton from '~/components/common/LinkButton.vue'
import SectionTitle from '~/components/top/SectionTitle.vue'

export default Vue.extend({
  name: 'PopularArticlesRanking',
  components: {
    SectionTitle,
    ArticleRow,
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
.popular-articles-wrapper {
  margin-bottom: responsive-height(20);
}

@include media-query($pc) {
  .popular-articles-wrapper {
    margin-bottom: 20px;
  }
}
</style>
