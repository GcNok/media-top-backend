<template>
  <section>
    <ListPageTitle
      title="徹底比較"
      sub-title="あなたの代わりにSS編集部が商品を購入して比較・検証"
    />
    <ArticleTab :type="articleType" />
    <ArticleList />
  </section>
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
      articleType: ConstArticle.ARTICLE_TYPE_COMPARISON
    }
  }
})
</script>
