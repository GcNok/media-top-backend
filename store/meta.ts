import { getterTree, mutationTree, actionTree } from 'typed-vuex'

export const state = () => ({
  title: '' as string,
  description: '' as string
})
export type RootState = ReturnType<typeof state>
export const getters = getterTree(state, {
  title: (state) => state.title,
  description: (state) => state.description
})

export const mutations = mutationTree(state, {
  setTitle(state, title: string): void {
    state.title = title
  },
  setDescription(state, description: string): void {
    state.description = description
  }
})

export const actions = actionTree(
  { state, getters, mutations },
  {
    getMetaInfo({ commit }): void {
      // TODO APIでここの値は更新したい
      commit('setTitle', '楽天モバイル')
      commit('setDescription', '楽天モバイルの説明')
    }
  }
)
