import {
  getAccessorType,
  getterTree,
  mutationTree,
  actionTree
} from 'typed-vuex'
import { ConstAPI } from '~/const/api'

import * as meta from '~/store/meta'
import { Article, ComparisonArticle } from '~/types/article'

// これらは型推論に必要のため、空でも定義しておく
export const state = () => ({
  isDisplaySidebar: false as boolean,
  isSpScreen: true as boolean,
  recommendArticles: [] as Article[],
  popularArticles: [] as Article[],
  newArticles: [] as Article[],
  comparisonArticles: [] as ComparisonArticle[]
})
export const getters = getterTree(state, {
  popularArticles: (state) => (limit: number = 0) => {
    return limit ? state.popularArticles.slice(0, limit) : state.popularArticles
  },
  newArticles: (state) => (limit: number = 0) => {
    return limit ? state.newArticles.slice(0, limit) : state.newArticles
  },
  comparisonArticles: (state) => (limit: number = 0) => {
    let articles: ComparisonArticle[]
    if (limit) {
      articles = state.comparisonArticles.slice(0, limit)
    } else {
      articles = state.comparisonArticles
    }
    return articles.filter((article) => {
      if (!article.productImageUrls) {
        return article
      }
      return article.productImageUrls.filter(
        (productImageUrl) => productImageUrl !== ''
      )
    })
  }
})
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
  addPopularArticles(state, articles) {
    state.popularArticles.push(...articles)
  },
  setNewrArticles(state, articles) {
    state.newArticles = articles
  },
  addNewArticles(state, articles) {
    state.newArticles.push(...articles)
  },
  setComparisonArticles(state, articles) {
    state.comparisonArticles = articles
  },
  addComparisonArticles(state, articles) {
    state.comparisonArticles.push(...articles)
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
    async getPopularArticles({ commit, state }, offset: number = 0) {
      const { data } = await this.app.$axios.get(
        `${ConstAPI.GET_POPULAR_ARTICLES}`,
        {
          params: {
            offset
          }
        }
      )
      if (state.popularArticles.length === 0) {
        commit('setPopularArticles', data)
      } else {
        commit('addPopularArticles', data)
      }
    },
    async getNewArticles({ commit, state }, offset: number = 0) {
      const { data } = await this.app.$axios.get(
        `${ConstAPI.GET_NEW_ARTICLES}`,
        {
          params: {
            offset
          }
        }
      )
      if (state.newArticles.length === 0) {
        commit('setNewrArticles', data)
      } else {
        commit('addNewArticles', data)
      }
    },
    async getComparisonArticles({ commit, state }, offset: number = 0) {
      const { data } = await this.app.$axios.get(
        `${ConstAPI.GET_COMARISON_ARTICLES}`,
        {
          params: {
            offset
          }
        }
      )
      if (state.comparisonArticles.length === 0) {
        commit('setComparisonArticles', data)
      } else {
        commit('addComparisonArticles', data)
      }
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
