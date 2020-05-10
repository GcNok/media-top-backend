<template>
  <div>
    <ListPageTitle
      title="人気記事ランキング"
      sub-title="繰り返し読みたくなる暮らしに役立つ記事をご紹介"
    />
    <ArticleTab :type="articleType" />
    <ArticleList />
    <no-ssr>
      <infinite-loading
        spinner="spiral"
        @infinite="infiniteHandler"
      ></infinite-loading>
    </no-ssr>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { ConstArticle } from '~/const/article'
import ListPageTitle from '~/components/common/ListPageTitle.vue'
import ArticleList from '~/components/common/ArticleList.vue'
import ArticleTab from '~/components/common/ArticleTab.vue'

export default Vue.extend({
  name: 'PopularPage',
  components: {
    ListPageTitle,
    ArticleList,
    ArticleTab
  },
  data() {
    return {
      isLoading: false,
      articleType: ConstArticle.ARTICLE_TYPE_POPULAR
    }
  },
  methods: {
    async infiniteHandler($state: any) {
      if (this.isLoading) return
      this.isLoading = true
      try {
        await this.$accessor.getPopularArticles()
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
