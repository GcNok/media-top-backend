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
        <div class="article-row-info">
          <span>{{ article.last30daysPv }} views</span>
        </div>
      </div>
      <div class="product-image-wrapper">
        <img
          v-for="(imageUrl, imageUrlIndex) in article.productImageUrls"
          :key="imageUrlIndex"
          class="product-image"
          :src="imageUrl"
          alt="商品"
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
    flex-direction: column;
    padding-left: responsive-width(16);

    .article-row-title {
      font-size: responsive-width(14);
      font-weight: bold;
    }

    .article-row-info {
      display: flex;
      justify-content: flex-end;
      align-items: center;
      font-size: responsive-width(12);
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
      width: responsive-width(40);
      height: responsive-width(40);
    }
  }
}

@include media-query($pc) {
  .article-list-wrapper {
    display: grid;
    .article-row-wrapper {
      grid-template-columns: 150px 1fr;
      padding: 6px 16px;
      height: 150px;
      border-bottom: 2px solid $color-gray;

      &:last-of-type {
        margin-bottom: 20px;
      }

      .article-row-left-wrapper {
        position: relative;
        grid-row: 1/3;
        grid-column: 1;
        display: flex;
        align-items: center;

        .article-row-image {
          width: 110px;
          height: 110px;
        }
      }

      .article-row-right-wrapper {
        grid-row: 1;
        grid-column: 2;
        justify-content: space-evenly;
        padding: 0;
        height: 100%;

        .article-row-title {
          font-size: 16px;
        }

        .article-row-info {
          font-size: 12px;
        }
      }

      .product-image-wrapper {
        grid-row: 2;
        grid-column: 2;
        margin-top: 10px;
        .product-image {
          margin-right: 20px;
          width: 40px;
          height: 40px;
        }
      }
    }
  }
}
</style>
