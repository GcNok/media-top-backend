<template>
  <div class="article-list-wrapper">
    <a
      v-for="(article, index) in articles"
      :key="index"
      :href="article.articleURL"
      class="article-row-wrapper"
    >
      <div class="article-row-left-wrapper">
        <img
          class="article-row-image"
          :src="article.mainVisual"
          alt="article"
        />
        <div
          v-if="type === ARTICLE_TYPE_POPULAR"
          class="article-mark"
          :class="{
            gold: index === 0,
            silver: index === 1,
            bronze: index === 2
          }"
        />
        <div v-else class="article-mark new-mark" />
        <span v-if="type === ARTICLE_TYPE_POPULAR" class="article-mark-text"
          >{{ index + 1 }}位</span
        >
        <span v-else class="article-mark-text new-text">new</span>
      </div>
      <div class="article-row-right-wrapper">
        <p class="article-row-title">
          {{ article.title }}
        </p>
        <div class="article-row-info">
          <span>{{ article.last30daysPv }} views</span>
        </div>
      </div>
    </a>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Article } from '~/types/article'
import { ConstArticle } from '~/const/article'

export default Vue.extend({
  name: 'ArticleList',
  props: {
    type: {
      type: String,
      default: ConstArticle.ARTICLE_TYPE_POPULAR
    }
  },
  data() {
    return {
      ARTICLE_TYPE_POPULAR: ConstArticle.ARTICLE_TYPE_POPULAR,
      ARTICLE_TYPE_NEW: ConstArticle.ARTICLE_TYPE_NEW
    }
  },
  computed: {
    articles(): Article[] {
      switch (this.type) {
        case ConstArticle.ARTICLE_TYPE_POPULAR:
          return this.$accessor.popularArticles
        case ConstArticle.ARTICLE_TYPE_NEW:
          return this.$accessor.newArticles
        default:
          return this.$accessor.articles
      }
    },
    popularArticles(): Article[] {
      return this.$accessor.popularArticles
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
  color: unset;
  text-decoration: none;

  &:last-of-type {
    margin-bottom: responsive-height(20);
  }

  .article-row-left-wrapper {
    position: relative;

    .article-row-image {
      width: responsive-width(80);
      height: responsive-width(80);
      object-fit: cover;
    }

    .article-mark {
      position: absolute;
      top: 0;
      border-top: responsive-width(22) solid gray;
      border-left: responsive-width(22) solid gray;
      border-right: responsive-width(22) solid transparent;
      border-bottom: responsive-width(22) solid transparent;
    }
    .article-mark-text {
      position: absolute;
      top: responsive-width(5);
      left: responsive-width(3);
      display: block;
      font-size: responsive-width(13);
      font-weight: bold;
      color: white;
      transform: rotate(-45deg);
      z-index: 100;
    }

    .new-mark {
      border-top-color: $color-orange;
      border-left-color: $color-orange;
    }
    .new-text {
      left: 0;
    }

    .gold {
      border-top-color: $color-gold;
      border-left-color: $color-gold;
    }

    .silver {
      border-top-color: silver;
      border-left-color: silver;
    }

    .bronze {
      border-top-color: brown;
      border-left-color: brown;
    }
  }

  .article-row-right-wrapper {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin-left: responsive-width(16);

    .article-row-title {
      font-size: responsive-width(14);
      font-weight: bold;
    }

    .article-row-info {
      display: flex;
      justify-content: flex-end;
      align-items: center;
      font-size: responsive-width(10);
    }
  }
}

@include media-query($pc) {
  .article-list-wrapper {
    display: grid;
    .article-row-wrapper {
      grid-column: 1/4;
      padding: 6px 16px;
      height: 130px;
      border-bottom: 2px solid $color-gray;

      &:nth-of-type(1) {
        grid-column: 1;
        height: 270px;
      }
      &:nth-of-type(2) {
        grid-column: 2;
        height: 270px;
      }
      &:nth-of-type(3) {
        grid-column: 3;
        height: 270px;
      }
      /* 最初の3記事共通スタイル */
      &:nth-of-type(-n + 3) {
        flex-direction: column;

        .article-row-left-wrapper {
          .article-row-image {
            width: 100%;
            height: 140px;
          }
        }
        .article-row-right-wrapper {
          margin: 0;
        }
      }
      &:last-of-type {
        margin-bottom: 20px;
      }

      .article-row-left-wrapper {
        position: relative;

        .article-row-image {
          width: 110px;
          height: 110px;
        }

        .article-mark {
          border-top-width: 30px;
          border-left-width: 30px;
          border-right-width: 30px;
          border-bottom-width: 30px;
        }
        .article-mark-text {
          position: absolute;
          top: 7px;
          left: 3px;
          display: block;
          font-size: 18px;
          color: white;
          transform: rotate(-45deg);
          z-index: 100;
        }
      }

      .article-row-right-wrapper {
        justify-content: space-evenly;
        margin-left: 16px;
        height: 100%;

        .article-row-title {
          font-size: 16px;
        }

        .article-row-info {
          font-size: 12px;
        }
      }
    }
  }
}
</style>
