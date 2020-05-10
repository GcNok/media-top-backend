<template>
  <div>
    <ListPageTitle
      title="新着記事"
      sub-title="暮らしのアイデアをSS編集部が毎日更新"
    />
    <ArticleTab :type="articleType" />
    <ArticleList :type="articleType" />
    <no-ssr>
      <infinite-loading spinner="spiral" @infinite="infiniteHandler">
        <div slot="no-more">これ以上記事は存在しません。</div>
      </infinite-loading>
    </no-ssr>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Article } from '~/types/article'
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
  async fetch({ app }) {
    if (!this.articles) return
    await app.$accessor.getNewArticles()
  },
  data() {
    return {
      isLoading: false as boolean,
      articleType: ConstArticle.ARTICLE_TYPE_NEW,
      offset: 0 as number
    }
  },
  computed: {
    articles(): Article[] {
      return this.$accessor.newArticles()
    }
  },
  methods: {
    async infiniteHandler($state: any) {
      if (this.isLoading) return
      this.isLoading = true
      try {
        this.offset += 10
        await this.$accessor.getNewArticles(this.offset)
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
