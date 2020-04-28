export class ConstArticle {
  static readonly ARTICLE_TYPE_POPULAR = 'popular'
  static readonly ARTICLE_TYPE_NEW = 'latest'
  static readonly ARTICLE_TYPE_COMPARISON = 'comparison'
  static readonly ARTICLE_TAB_LIST = [
    {
      title: '人気記事',
      type: ConstArticle.ARTICLE_TYPE_POPULAR
    },
    {
      title: '新着記事',
      type: ConstArticle.ARTICLE_TYPE_NEW
    },
    {
      title: '徹底比較記事',
      type: ConstArticle.ARTICLE_TYPE_COMPARISON
    }
  ]
}
