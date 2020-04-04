<template>
  <div class="article-row-wrapper">
    <div class="article-row-left-wrapper">
      <img
        class="article-row-image"
        :src="article.articleImage"
        alt="article"
      />
      <span
        v-if="type === ARTICLE_TYPE_RANK"
        class="article-mark rank"
        :class="{ gold: rank === 1, silver: rank === 2, bronze: rank === 3 }"
        >{{ rank }}</span
      >
      <span v-else class="article-mark new">new</span>
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
  name: 'ArticleRow',
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
  },
  data() {
    return {
      ARTICLE_TYPE_RANK: Const.ARTICLE_TYPE_RANK,
      ARTICLE_TYPE_NEW: Const.ARTICLE_TYPE_NEW
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

    .article-mark {
      position: absolute;
      top: 3px;
      left: 3px;
      display: inline-block;
      width: 26px;
      height: 26px;
      line-height: 26px;
      text-align: center;
      font-size: 13px;
      color: white;
      border: 1px solid white;
      border-radius: 50%;
    }

    .rank {
      background-color: gray;
    }

    .new {
      font-size: 10px;
      font-weight: bold;
      line-height: 23px;
      background-color: orange;
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
