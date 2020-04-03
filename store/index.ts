import { getAccessorType } from 'typed-vuex'
import * as meta from '~/store/meta'

// これらは型推論に必要のため、空でも定義しておく
export const state = () => ({
  hoge: 'test' as string
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
