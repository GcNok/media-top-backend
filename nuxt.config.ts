import { Configuration } from '@nuxt/types'

const envPath = `config/.env.${process.env.ENV || 'local'}`
require('dotenv').config({ path: envPath })
const nuxtConfig: Configuration = {
  mode: 'universal',
  /*
   ** Headers of the page
   */
  head: {
    title: '日用品・食品の特集｜おすすめ・比較・ランキング',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content:
          '日用品・食品のネット通販を使いこなすお役立ちコンテンツ集。おすすめ商品、ランキング、賢い買い方、徹底比較。スマートショッピングは日本唯一の「送料込で価格比較」「残量を予想」する通販コンシェルジュ。'
      }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      {
        rel: 'stylesheet',
        href: 'https://fonts.googleapis.com/icon?family=Material+Icons'
      }
    ]
  },
  dotenv: {
    filename: envPath
  },
  /*
   ** Customize the progress-bar color
   */
  loading: { color: '#fff' },
  /*
   ** Global CSS
   */
  css: [],
  /*
   ** Plugins to load before mounting the App
   */
  plugins: [
    { src: '~/plugins/vue-carousel', ssr: false },
    { src: '~/plugins/vue-infinite-loading', ssr: false }
  ],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    '@nuxt/typescript-build',
    // Doc: https://github.com/nuxt-community/stylelint-module
    '@nuxtjs/stylelint-module',
    'nuxt-typed-vuex'
  ],
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    // Doc: https://github.com/nuxt-community/dotenv-module
    '@nuxtjs/dotenv',
    '@nuxtjs/style-resources'
  ],
  styleResources: {
    scss: ['~/assets/scss/style.scss', '~/assets/scss/function.scss']
  },
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {
    proxy: true,
    prefix: '/api'
  },
  /*
   ** Build configuration
   */
  proxy: {
    '/api': {
      target: process.env.BASE_URL,
      pathRewrite: {
        '^/api': ''
      }
    }
  },
  router: {
    middleware: ['init-sidebar']
  },
  build: {
    publicPath: '/contents/_nuxt/', // ビルドファイルの配置先を/_nuxtから変更
    transpile: [/typed-vuex/]
  }
}
export default nuxtConfig
