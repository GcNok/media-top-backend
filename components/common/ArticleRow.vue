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
  padding: responsive-height(6) responsive-width(16);
  border-bottom: 2px solid $color-gray;

  .article-row-left-wrapper {
    position: relative;

    .article-row-image {
      width: responsive-width(80);
      height: responsive-width(80);
      object-fit: cover;
    }

    .article-mark {
      position: absolute;
      top: responsive-width(3);
      left: responsive-width(3);
      display: inline-block;
      width: responsive-width(26);
      height: responsive-width(26);
      line-height: responsive-width(26);
      text-align: center;
      font-size: responsive-width(13);
      color: white;
      border: 1px solid white;
      border-radius: 50%;
    }

    .rank {
      background-color: gray;
    }

    .new {
      font-size: responsive-width(10);
      font-weight: bold;
      line-height: responsive-width(23);
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
    margin-left: responsive-width(16);

    .article-row-title {
      font-size: responsive-width(14);
      font-weight: bold;
    }

    .article-row-info {
      display: flex;
      justify-content: flex-end;
      align-items: center;
      font-size: responsive-width(7);
    }
  }
}
</style>
