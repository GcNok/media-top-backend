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
          alt="mainVisual"
        />
      </div>
      <div class="article-row-right-wrapper">
        <p class="article-row-title">
          {{ article.title }}
        </p>
      </div>
      <div class="product-image-wrapper">
        <img
          v-for="(imageUrl, imageUrlIndex) in article.productImageUrls"
          :key="imageUrlIndex"
          class="product-image"
          :src="imageUrl"
          alt="product-image"
        />
      </div>
    </a>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { ComparisonArticle } from '~/types/article'

export default Vue.extend({
  name: 'ComparisonArticleList',
  computed: {
    articles(): ComparisonArticle[] {
      return this.$accessor.comparisonArticles.filter((article) => {
        return article.productImageUrls.filter(
          (productImageUrl) => productImageUrl !== ''
        )
      })
    }
  }
})
</script>

<style lang="scss" scoped>
.article-row-wrapper {
  display: grid;
  padding: responsive-height(6) responsive-width(16);
  border-bottom: 2px solid $color-gray;
  color: unset;
  text-decoration: none;

  &:last-of-type {
    margin-bottom: responsive-height(20);
  }

  .article-row-left-wrapper {
    position: relative;
    grid-row: 1;
    grid-column: 1;

    .article-row-image {
      width: responsive-width(80);
      height: responsive-width(80);
      object-fit: cover;
    }
  }

  .article-row-right-wrapper {
    grid-row: 1;
    grid-column: 2;
    display: flex;
    align-items: center;
    padding-left: responsive-width(16);

    .article-row-title {
      font-size: responsive-width(14);
      font-weight: bold;
    }
  }

  .product-image-wrapper {
    grid-row: 2;
    grid-column: 1/3;
    margin-top: responsive-height(10);
    white-space: nowrap;
    overflow-x: scroll;
    .product-image {
      margin-right: responsive-width(5);
      width: responsive-width(50);
      height: responsive-width(50);
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
