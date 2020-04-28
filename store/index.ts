import { getAccessorType, mutationTree, actionTree } from 'typed-vuex'
import { ConstAPI } from '~/const/api'

import * as meta from '~/store/meta'
import { Article } from '~/types/article'

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
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】カップラーメン鉄板おすすめランキング10選
    【商品を購入して食べ比べました】`,
      articleURL: 'https://smashop.jp/subcategory/404000/special',
      mainVisual: '/img/cup-noodle.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】発泡酒鉄板おすすめランキング11選
    【糖質あり・なし人気商品を購入して飲み比べ検証】`,
      articleURL: 'https://smashop.jp/subcategory/603030/special',
      mainVisual: '/img/beer.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      mainVisual: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】カップラーメン鉄板おすすめランキング10選
    【商品を購入して食べ比べました】`,
      articleURL: 'https://smashop.jp/subcategory/404000/special',
      mainVisual: '/img/cup-noodle.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】発泡酒鉄板おすすめランキング11選
    【糖質あり・なし人気商品を購入して飲み比べ検証】`,
      articleURL: 'https://smashop.jp/subcategory/603030/special',
      mainVisual: '/img/beer.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      mainVisual: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】カップラーメン鉄板おすすめランキング10選
    【商品を購入して食べ比べました】`,
      articleURL: 'https://smashop.jp/subcategory/404000/special',
      mainVisual: '/img/cup-noodle.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `【徹底比較】発泡酒鉄板おすすめランキング11選
    【糖質あり・なし人気商品を購入して飲み比べ検証】`,
      articleURL: 'https://smashop.jp/subcategory/603030/special',
      mainVisual: '/img/beer.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    },
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      mainVisual: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      rewriteModified: '1時間前',
      last30daysPv: '000,000'
    }
  ] as Article[],
  popularArticles: [] as Article[],
  newArticles: [] as Article[]
})
export const getters = {}
export const mutations = mutationTree(state, {
  toggleSidebar(state): void {
    state.isDisplaySidebar = !state.isDisplaySidebar
  },
  closeSidebar(state): void {
    state.isDisplaySidebar = false
  },
  setPopularArticles(state, articles) {
    state.popularArticles = articles
  },
  setNewrArticles(state, articles) {
    state.newArticles = articles
  }
})
export const actions = actionTree(
  { state, mutations },
  {
    async getPopularArticles({ commit }) {
      const { data } = await this.app.$axios.get(
        `${ConstAPI.GET_POPULAR_ARTICLES}`
      )
      commit('setPopularArticles', data)
    },
    async getNewArticles({ commit }) {
      const { data } = await this.app.$axios.get(`${ConstAPI.GET_NEW_ARTICLES}`)
      commit('setNewrArticles', data)
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
