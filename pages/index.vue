<template>
  <main>
    <SectionTitle :section="sections[0]" />
    <TopCarousel />
    <SectionTitle :section="sections[1]" />
    <PopularRanking />
    <SectionTitle :section="sections[2]" />
    <SectionTitle :section="sections[3]" />
    <SectionTitle :section="sections[4]" />
  </main>
</template>

<script lang="ts">
import Vue from 'vue'
import SectionTitle from '~/components/top/SectionTitle.vue'
import TopCarousel from '~/components/top/TopCarousel.vue'
import PopularRanking from '~/components/top/PopularRanking.vue'

export default Vue.extend({
  components: {
    SectionTitle,
    TopCarousel,
    PopularRanking
  },
  data() {
    return {
      sections: [
        {
          title: '話題の記事',
          comment: 'SS編集部が厳選したおすすめの今すぐ読むべき記事まとめ'
        },
        {
          title: '人気記事ランキング',
          comment: '繰り返し読みたくなる暮らしに役立つ記事をご紹介'
        },
        {
          title: '新着記事',
          comment: '暮らしのアイデアをSS編集部が毎日更新'
        },
        {
          title: '今月の徹底比較記事ランキング',
          comment: 'あなたの代わりにSS編集部が商品を購入して比較・検証'
        },
        {
          title: '専門家監修のおすすめ記事一覧',
          comment: 'その道のプロや専門家の愛用品やおすすめを一挙大公開'
        },
        {
          title: '人気記事ランキング',
          comment: '繰り返し読みたくなる暮らしに役立つ記事をご紹介'
        }
      ]
    }
  }
})
</script>

<style>
/* Sample `apply` at-rules with Tailwind CSS
.container {
  @apply min-h-screen flex justify-center items-center text-center mx-auto;
}
*/
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
