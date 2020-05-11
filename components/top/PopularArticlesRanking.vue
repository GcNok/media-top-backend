<template>
  <section class="popular-articles-wrapper">
    <SectionTitle
      :title="'人気記事ランキング'"
      :comment="'繰り返し読みたくなる暮らしに役立つ記事をご紹介'"
    />
    <ArticleList :limit="10" />
    <LinkButton
      :title="'人気記事ランキング一覧を見る'"
      :link="'/contents/popular'"
    />
  </section>
</template>

<script lang="ts">
import Vue from 'vue'
import ArticleList from '~/components/common/ArticleList.vue'
import LinkButton from '~/components/common/LinkButton.vue'
import SectionTitle from '~/components/top/SectionTitle.vue'

export default Vue.extend({
  name: 'PopularArticlesRanking',
  components: {
    SectionTitle,
    ArticleList,
    LinkButton
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
