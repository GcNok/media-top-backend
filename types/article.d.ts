export type Article = {
  title: string
  articleURL: string
  mainVisual: string
  writerImage: string
  writerName: string
  last30daysPv: string
  updated: string
}

export type ComparisonArticle = {
  title: string
  articleURL: string
  mainVisual: string
  writerImage: string
  writerName: string
  last30daysPv: string
  productImageUrls: string[]
  updated: string
}
