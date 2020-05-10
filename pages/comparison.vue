<template>
  <section>
    <ListPageTitle
      title="徹底比較"
      sub-title="あなたの代わりにSS編集部が商品を購入して比較・検証"
    />
    <ArticleTab :type="articleType" />
    <ComparisonArticleList />
    <no-ssr>
      <infinite-loading spinner="spiral" @infinite="infiniteHandler">
        <div slot="no-more">これ以上記事は存在しません。</div>
      </infinite-loading>
    </no-ssr>
  </section>
</template>

<script lang="ts">
import Vue from 'vue'
import { ConstArticle } from '~/const/article'
import { Article } from '~/types/article'
import ListPageTitle from '~/components/common/ListPageTitle.vue'
import ComparisonArticleList from '~/components/common/ComparisonArticleList.vue'
import ArticleTab from '~/components/common/ArticleTab.vue'

export default Vue.extend({
  name: 'PopularPage',
  components: {
    ListPageTitle,
    ComparisonArticleList,
    ArticleTab
  },
  async fetch({ app }) {
    if (this.articles) return
    await app.$accessor.getComparisonArticles()
  },
  data() {
    return {
      isLoading: false as boolean,
      articleType: ConstArticle.ARTICLE_TYPE_COMPARISON,
      offset: 0 as number
    }
  },
  computed: {
    articles(): Article[] {
      return this.$accessor.comparisonArticles()
    }
  },
  methods: {
    async infiniteHandler($state: any) {
      if (this.isLoading) return
      this.isLoading = true
      try {
        this.offset += 10
        await this.$accessor.getComparisonArticles(this.offset)
        $state.loaded()
      } catch (e) {
        $state.complete()
      } finally {
        this.isLoading = false
      }
    }
  }
})
</script>
