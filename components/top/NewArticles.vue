<template>
  <section class="new-articles-wrapper">
    <SectionTitle
      :title="'新着記事'"
      :comment="'暮らしのアイデアをSS編集部が毎日更新'"
    />
    <ArticleRow
      v-for="(article, index) in articles"
      :key="index"
      class="article-list-item"
      :article="article"
      :type="ARTICLE_TYPE_NEW"
    />
    <LinkButton :title="'新着記事一覧を見る'" :link="'/new'" />
  </section>
</template>

<script lang="ts">
import Vue from 'vue'
import { Article } from '~/types/article'
import ArticleRow from '~/components/common/ArticleRow.vue'
import LinkButton from '~/components/common/LinkButton.vue'
import SectionTitle from '~/components/top/SectionTitle.vue'
import { Const } from '~/const/const'

export default Vue.extend({
  name: 'NewArticles',
  components: {
    SectionTitle,
    ArticleRow,
    LinkButton
  },
  data() {
    return {
      ARTICLE_TYPE_NEW: Const.ARTICLE_TYPE_NEW
    }
  },
  computed: {
    articles(): Article[] {
      return this.$accessor.articles
    }
  }
})
</script>

<style lang="scss" scoped>
.new-articles-wrapper {
  margin-bottom: 20px;
  .article-list-item {
    &:last-of-type {
      margin-bottom: 20px;
    }
  }
}
</style>
