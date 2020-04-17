<template>
  <section>
    <ListPageTitle
      title="人気記事ランキング"
      sub-title="繰り返し読みたくなる暮らしに役立つ記事をご紹介"
    />
    <ArticleTab :type="articleType" />
    <ArticleList />
  </section>
</template>

<script lang="ts">
import Vue from 'vue'
import { Const } from '~/const/const'
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
      articleType: Const.ARTICLE_TYPE_POPULAR
    }
  }
})
</script>
