<template>
  <div class="article-row-wrapper">
    <div class="article-row-left-wrapper">
      <img
        class="article-row-image"
        :src="article.articleImage"
        alt="article"
      />
      <span
        class="raniking-num"
        :class="{ gold: rank === 1, silver: rank === 2, bronze: rank === 3 }"
        >{{ rank }}</span
      >
    </div>
    <div class="article-row-right-wrapper">
      <p class="article-row-title">
        {{ article.title }}
      </p>
      <div class="article-row-info">
        <span>{{ article.viewNum }} views</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Article } from '~/types/article'
import { Const } from '~/const/const'

export default Vue.extend({
  name: 'ArticleList',
  props: {
    article: {
      type: Object as () => Article,
      default: {}
    },
    type: {
      type: String,
      default: Const.ARTICLE_TYPE_RANK
    },
    rank: {
      type: Number,
      default: 0
    }
  }
})
</script>

<style lang="scss" scoped>
.article-row-wrapper {
  display: flex;
  align-items: center;
  padding: 0.4rem 1rem;
  border-bottom: 2px solid $color-gray;

  .article-row-left-wrapper {
    position: relative;

    .article-row-image {
      width: 5rem;
      height: 5rem;
      object-fit: cover;
    }

    .raniking-num {
      position: absolute;
      top: 3px;
      left: 3px;
      display: inline-block;
      width: 26px;
      height: 26px;
      line-height: 26px;
      text-align: center;
      font-size: 14px;
      font-weight: bold;
      color: white;
      background-color: gray;
      border: 1px solid white;
      border-radius: 50%;
    }

    .gold {
      background-color: rgb(255, 208, 0);
    }

    .silver {
      background-color: silver;
    }

    .bronze {
      background-color: brown;
    }
  }

  .article-row-right-wrapper {
    display: flex;
    flex-direction: column;
    margin-left: 1rem;

    .article-row-title {
      font-size: 0.9rem;
      font-weight: bold;
    }

    .article-row-info {
      display: flex;
      justify-content: flex-end;
      align-items: center;
      font-size: 0.4rem;
    }
  }
}
</style>
