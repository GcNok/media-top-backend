<template>
  <section class="comparison-articles-wrapper">
    <SectionTitle
      :title="'今月の徹底比較記事ランキング'"
      :comment="'あなたの代わりにSS編集部が商品を購入して比較・検証'"
    />
    <ComparisonArticleList :limit="10" />
    <LinkButton :title="'徹底比較記事一覧を見る'" :link="'/comparison'" />
  </section>
</template>

<script lang="ts">
import Vue from 'vue'
import { Article } from '~/types/article'
import ComparisonArticleList from '~/components/common/ComparisonArticleList.vue'
import LinkButton from '~/components/common/LinkButton.vue'
import SectionTitle from '~/components/top/SectionTitle.vue'

export default Vue.extend({
  name: 'ComparisonArticles',
  components: {
    SectionTitle,
    ComparisonArticleList,
    LinkButton
  }
})
</script>

<style lang="scss" scoped>
.comparison-articles-wrapper {
  margin-bottom: responsive-height(20);
}

@include media-query($pc) {
  .comparison-articles-wrapper {
    margin-bottom: 20px;
  }
}
</style>
