<template>
  <a
    class="carousel-image-wrapper"
    :href="article.articleURL"
    :style="{ backgroundImage: 'url(' + article.mainVisual + ')' }"
  >
    <div class="article-image">
      <p class="article-title">
        {{ article.title }}
      </p>
      <div class="article-row">
        <div class="writer-wrapper">
          <img class="writer-image" :src="article.writerImage" />
          <span class="writer-name">{{ article.writerName }}</span>
        </div>
        <div class="article-info-wrapper">
          <span>最終更新：{{ article.rewriteModified }}</span>
          <span>{{ article.last30daysPv }} views</span>
        </div>
      </div>
    </div>
  </a>
</template>

<script lang="ts">
import Vue from 'vue'
import { Article } from '~/types/article'
export default Vue.extend({
  name: 'CarouselSlide',
  props: {
    article: {
      type: Object as () => Article,
      default: {}
    }
  }
})
</script>

<style lang="scss" scoped>
.carousel-image-wrapper {
  display: flex;
  align-items: flex-end;
  height: responsive-width(190);
  min-width: responsive-width(272);
  box-shadow: inset 0 -96px 32px rgba(0, 0, 0, 0.6);
  background-size: cover;
  text-decoration: none;
  border-radius: 10px;

  .article-image {
    display: grid;
    padding: responsive-height(10) responsive-width(10);
    color: white;

    .article-title {
      grid-row: 1;
      margin-bottom: responsive-height(7);
      font-size: responsive-width(13);
      font-weight: bold;
    }

    .article-row {
      display: flex;
      justify-content: space-between;

      .writer-wrapper {
        display: flex;
        align-items: center;

        .writer-image {
          margin-right: responsive-width(3);
          width: responsive-width(32);
          border-radius: 50%;
        }

        .writer-name {
          font-size: responsive-width(10);
          font-weight: bold;
        }
      }

      .article-info-wrapper {
        display: flex;
        flex-direction: column;
        align-items: flex-end;
        font-size: responsive-width(10);
      }
    }
  }
}

@include media-query($pc) {
  .carousel-image-wrapper {
    height: 250px;
    width: 400px;
    min-width: unset;
    box-shadow: inset 0 -96px 32px rgba(0, 0, 0, 0.6);
    background-size: cover;
    text-decoration: none;
    border-radius: 10px;

    .article-image {
      display: grid;
      padding: 10px 10px;
      color: white;

      .article-title {
        grid-row: 1;
        margin-bottom: 7px;
        font-size: 13px;
        font-weight: bold;
      }

      .article-row {
        display: flex;
        justify-content: space-between;

        .writer-wrapper {
          display: flex;
          align-items: center;

          .writer-image {
            margin-right: 3px;
            width: 32px;
            border-radius: 50%;
          }

          .writer-name {
            font-size: 10px;
            font-weight: bold;
          }
        }

        .article-info-wrapper {
          display: flex;
          flex-direction: column;
          align-items: flex-end;
          font-size: 10px;
        }
      }
    }
  }
}
</style>
