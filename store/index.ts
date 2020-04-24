import { getAccessorType } from 'typed-vuex'
import * as meta from '~/store/meta'
import { Article } from '~/types/article'

// これらは型推論に必要のため、空でも定義しておく
export const state = () => ({
  articles: [
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      articleImage: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    },
    {
      title: `【徹底比較】カップラーメン鉄板おすすめランキング10選
    【商品を購入して食べ比べました】`,
      articleURL: 'https://smashop.jp/subcategory/404000/special',
      articleImage: '/img/cup-noodle.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    },
    {
      title: `【徹底比較】発泡酒鉄板おすすめランキング11選
    【糖質あり・なし人気商品を購入して飲み比べ検証】`,
      articleURL: 'https://smashop.jp/subcategory/603030/special',
      articleImage: '/img/beer.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    },
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      articleImage: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    },
    {
      title: `【徹底比較】カップラーメン鉄板おすすめランキング10選
    【商品を購入して食べ比べました】`,
      articleURL: 'https://smashop.jp/subcategory/404000/special',
      articleImage: '/img/cup-noodle.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    },
    {
      title: `【徹底比較】発泡酒鉄板おすすめランキング11選
    【糖質あり・なし人気商品を購入して飲み比べ検証】`,
      articleURL: 'https://smashop.jp/subcategory/603030/special',
      articleImage: '/img/beer.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    },
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      articleImage: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    },
    {
      title: `【徹底比較】カップラーメン鉄板おすすめランキング10選
    【商品を購入して食べ比べました】`,
      articleURL: 'https://smashop.jp/subcategory/404000/special',
      articleImage: '/img/cup-noodle.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    },
    {
      title: `【徹底比較】発泡酒鉄板おすすめランキング11選
    【糖質あり・なし人気商品を購入して飲み比べ検証】`,
      articleURL: 'https://smashop.jp/subcategory/603030/special',
      articleImage: '/img/beer.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    },
    {
      title: `ふりかえ鉄板おすすめ人気ランキング10選
    【ごはん検定特Aライターが厳選！話題の商品紹介も】`,
      articleURL: 'https://smashop.jp/subcategory/403700/special',
      articleImage: '/img/hurikake2.jpg',
      writerImage: '/img/writer.jpg',
      writerName: '実用書ライター：小田原',
      updateTime: '1時間前',
      viewNum: '000,000'
    }
  ] as Article[]
})
export const getters = {}
export const mutations = {}
export const actions = {}

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
