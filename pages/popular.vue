<template>
  <div>
    <ListPageTitle
      title="人気記事ランキング"
      sub-title="繰り返し読みたくなる暮らしに役立つ記事をご紹介"
    />
    <ArticleTab :type="articleType" />
    <ArticleList />
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
      articleType: ConstArticle.ARTICLE_TYPE_POPULAR
    }
  }
})
</script>
