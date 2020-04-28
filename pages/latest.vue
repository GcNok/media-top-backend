<template>
  <div>
    <ListPageTitle
      title="新着記事"
      sub-title="暮らしのアイデアをSS編集部が毎日更新"
    />
    <ArticleTab :type="articleType" />
    <ArticleList :type="articleType" />
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
      articleType: ConstArticle.ARTICLE_TYPE_NEW
    }
  }
})
</script>
