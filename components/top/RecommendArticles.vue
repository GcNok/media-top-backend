<template>
  <section class="reccoment-article-wrapper">
    <SectionTitle
      :title="'話題の記事'"
      :comment="'SS編集部が厳選したおすすめの今すぐ読むべき記事まとめ'"
    />
    <carousel
      class="carousel"
      autoplay
      :autoplay-timeout="4000"
      autoplay-hover-pause
      :per-page="1"
      :per-page-custom="[
        [768, 1],
        [1024, 2]
      ]"
      :loop="true"
      center-mode
      :pagination-padding="5"
    >
      <slide v-for="(article, index) in articles" :key="index">
        <CarouselSlide :article="article" />
      </slide>
    </carousel>
  </section>
</template>

<script lang="ts">
import Vue from 'vue'
import Carousel from 'vue-carousel/src/Carousel.vue'
import Slide from 'vue-carousel/src/Slide.vue'
import SectionTitle from '~/components/top/SectionTitle.vue'
import CarouselSlide from '~/components/top/CarouselSlide.vue'
import { Article } from '~/types/article'
export default Vue.extend({
  name: 'RecommendArticles',
  components: {
    SectionTitle,
    CarouselSlide,
    Carousel,
    Slide
  },
  computed: {
    articles(): Article[] {
      return this.$accessor.recommendArticles
    }
  }
})
</script>

<style lang="scss">
.reccoment-article-wrapper {
  margin: responsive-height(20) 0;

  .carousel {
    max-width: 100vw;
    padding: responsive-height(16) responsive-width(16) 0;
  }
}

@include media-query($pc) {
  .reccoment-article-wrapper {
    margin: 20px 0;

    .carousel {
      padding: 16px 30px;
      background-color: whitesmoke;
    }
  }
}

@include media-query($sp) {
  .VueCarousel-wrapper {
    border-radius: 10px;
  }
}
</style>
