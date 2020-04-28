import { Context } from '@nuxt/types'
export default ({ app: { $accessor } }: Context) => {
  $accessor.closeSidebar()
}
