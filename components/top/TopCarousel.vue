<template>
  <section class="carousel-wrapper">
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
import CarouselSlide from '~/components/top/CarouselSlide.vue'
import { Article } from '~/types/article'
export default Vue.extend({
  name: 'TopCarousel',
  components: {
    CarouselSlide
  },
  computed: {
    articles(): Article[] {
      return this.$accessor.articles
    }
  }
})
</script>

<style lang="scss">
.carousel-wrapper {
  display: flex;

  .carousel {
    max-width: 100vw;
    padding: 16px 16px 0;
  }
}

.VueCarousel-wrapper {
  box-shadow: 0.2rem 0.2rem 10px rgba(0, 0, 0, 0.6);
  border-radius: 10px;
}
</style>
