export type Content = {
  id: number
  provider: string
  provider_id: string
  title: string
  type: "video" | "article"
  tags: string[]
  published_at: string
  score: number
}
