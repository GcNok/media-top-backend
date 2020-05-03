import { getAccessorType, mutationTree, actionTree } from 'typed-vuex'
import { ConstAPI } from '~/const/api'

import * as meta from '~/store/meta'
import { Article, ComparisonArticle } from '~/types/article'

// これらは型推論に必要のため、空でも定義しておく
export const state = () => ({
  isDisplaySidebar: false as boolean,
  isSpScreen: true as boolean,
  articles: [
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      mainVisual: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】カップラーメン鉄板おすすめランキング10選
    【商品を購入して食べ比べました】`,
      articleURL: 'https://smashop.jp/subcategory/404000/special',
      mainVisual: '/img/cup-noodle.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】発泡酒鉄板おすすめランキング11選
    【糖質あり・なし人気商品を購入して飲み比べ検証】`,
      articleURL: 'https://smashop.jp/subcategory/603030/special',
      mainVisual: '/img/beer.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      mainVisual: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】カップラーメン鉄板おすすめランキング10選
    【商品を購入して食べ比べました】`,
      articleURL: 'https://smashop.jp/subcategory/404000/special',
      mainVisual: '/img/cup-noodle.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】発泡酒鉄板おすすめランキング11選
    【糖質あり・なし人気商品を購入して飲み比べ検証】`,
      articleURL: 'https://smashop.jp/subcategory/603030/special',
      mainVisual: '/img/beer.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      mainVisual: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】カップラーメン鉄板おすすめランキング10選
    【商品を購入して食べ比べました】`,
      articleURL: 'https://smashop.jp/subcategory/404000/special',
      mainVisual: '/img/cup-noodle.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】発泡酒鉄板おすすめランキング11選
    【糖質あり・なし人気商品を購入して飲み比べ検証】`,
      articleURL: 'https://smashop.jp/subcategory/603030/special',
      mainVisual: '/img/beer.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      mainVisual: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updated: '1時間前',
      last30daysPv: '000,000'
    }
  ] as Article[],
  recommendArticles: [] as Article[],
  popularArticles: [] as Article[],
  newArticles: [] as Article[],
  comparisonArticles: [] as ComparisonArticle[]
})
export const getters = {}
export const mutations = mutationTree(state, {
  toggleSidebar(state): void {
    state.isDisplaySidebar = !state.isDisplaySidebar
  },
  closeSidebar(state): void {
    state.isDisplaySidebar = false
  },
  setRecommendArticles(state, articles) {
    state.recommendArticles = articles
  },
  setPopularArticles(state, articles) {
    state.popularArticles = articles
  },
  setNewrArticles(state, articles) {
    state.newArticles = articles
  },
  setComparisonArticles(state, articles) {
    state.comparisonArticles = articles
  }
})
export const actions = actionTree(
  { state, mutations },
  {
    async getRecommendArticles({ commit }) {
      const { data } = await this.app.$axios.get(
        `${ConstAPI.GET_RECOMMEND_ARTICLES}`
      )
      commit('setRecommendArticles', data)
    },
    async getPopularArticles({ commit }) {
      const { data } = await this.app.$axios.get(
        `${ConstAPI.GET_POPULAR_ARTICLES}`
      )
      commit('setPopularArticles', data)
    },
    async getNewArticles({ commit }) {
      const { data } = await this.app.$axios.get(`${ConstAPI.GET_NEW_ARTICLES}`)
      commit('setNewrArticles', data)
    },
    async getComparisonArticles({ commit }) {
      const { data } = await this.app.$axios.get(
        `${ConstAPI.GET_COMARISON_ARTICLES}`
      )
      commit('setComparisonArticles', data)
    }
  }
)
export const accessorType = getAccessorType({
  state,
  getters,
  mutations,
  actions,
  modules: {
    // インポートしたサブモジュールを定義
    meta
  }
})
