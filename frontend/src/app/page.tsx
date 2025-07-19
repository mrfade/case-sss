"use client"

import { useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { Search, Filter, SortAsc, SortDesc, Calendar, Tag, Star } from "lucide-react"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Skeleton } from "@/components/ui/skeleton"
import { Pagination } from "@/components/pagination"
import { ThemeToggle } from "@/components/theme-toggle"
import type { ApiResponse } from "@/types/api"
import type { Content } from "@/types/content"

const fetchContent = async (params: {
  search?: string
  type?: string
  sortBy?: string
  sortOrder?: "asc" | "desc"
  page?: number
  pageSize?: number
}): Promise<ApiResponse<Content[]>> => {
  const query = new URLSearchParams({
    ...(params.search && { 'search[title]': params.search }),
    ...(params.type && { 'filter[type]': params.type }),
    ...(params.sortBy && params.sortOrder && { sort: `${params.sortOrder === "asc" ? "" : "-"}${params.sortBy}` }),
    ...(params.page && { 'page[number]': params.page }),
    ...(params.pageSize && { 'page[size]': params.pageSize }),
  } as Record<string, string>)
  const response = await fetch(`http://localhost:8080/api/v1/contents?${query.toString()}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    }
  })

  if (!response.ok) {
    throw new Error("Network response was not ok")
  }

  return response.json()
}

export default function SearchDashboard() {
  const [search, setSearch] = useState("")
  const [typeFilter, setTypeFilter] = useState("all")
  const [sortOrder, setSortOrder] = useState<"asc" | "desc">("desc")
  const [currentPage, setCurrentPage] = useState(1)

  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["content", search, typeFilter, sortOrder, currentPage],
    queryFn: () =>
      fetchContent({
        search: search || undefined,
        type: typeFilter === "all" ? undefined : typeFilter,
        sortBy: "score",
        sortOrder,
        page: currentPage,
        pageSize: 5,
      }),
    staleTime: 5 * 60 * 1000, // 5 minutes
  })

  const handleSearch = (value: string) => {
    setSearch(value)
    setCurrentPage(1)
  }

  const handleTypeFilter = (value: string) => {
    setTypeFilter(value)
    setCurrentPage(1)
  }

  const toggleSortOrder = () => {
    setSortOrder((prev) => (prev === "desc" ? "asc" : "desc"))
  }

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("tr-TR", {
      year: "numeric",
      month: "short",
      day: "numeric",
    })
  }

  const formatScore = (score: number) => {
    return score.toFixed(2)
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
        <div className="container mx-auto p-6">
          <Card className="border-red-200 dark:border-red-800 bg-white dark:bg-gray-800">
            <CardContent className="p-8 text-center">
              <p className="text-red-600 dark:text-red-400 text-lg font-medium mb-4">
                Bir hata oluştu. Lütfen tekrar deneyin.
              </p>
              <Button onClick={() => refetch()} className="bg-red-600 hover:bg-red-700 text-white">
                Tekrar Dene
              </Button>
            </CardContent>
          </Card>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      <div className="container mx-auto p-6 space-y-6">
        <div className="text-center space-y-4 relative">
          <div className="absolute top-0 right-0">
            <ThemeToggle />
          </div>
          <h1 className="text-3xl font-bold text-gray-900 dark:text-gray-100">İçerik Arama Motoru</h1>
          <p className="text-gray-600 dark:text-gray-400">Farklı sağlayıcılardan gelen içerikleri arayın ve keşfedin</p>
        </div>

        <Card className="border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800">
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-gray-900 dark:text-gray-100">
              <Search className="w-5 h-5" />
              Arama ve Filtreler
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="flex flex-col md:flex-row gap-4">
              <div className="flex-1">
                <Input
                  placeholder="İçerik başlığında ara..."
                  value={search}
                  onChange={(e) => handleSearch(e.target.value)}
                  className="w-full border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100"
                />
              </div>

              <div className="flex items-center gap-2">
                <Filter className="w-4 h-4 text-gray-600 dark:text-gray-400" />
                <Select value={typeFilter} onValueChange={handleTypeFilter}>
                  <SelectTrigger className="w-40 border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100">
                    <SelectValue placeholder="İçerik Türü" />
                  </SelectTrigger>
                  <SelectContent className="bg-white dark:bg-gray-800 border-gray-200 dark:border-gray-700">
                    <SelectItem value="all" className="text-gray-900 dark:text-gray-100">
                      Tümü
                    </SelectItem>
                    <SelectItem value="video" className="text-gray-900 dark:text-gray-100">
                      Video
                    </SelectItem>
                    <SelectItem value="article" className="text-gray-900 dark:text-gray-100">
                      Makale
                    </SelectItem>
                  </SelectContent>
                </Select>
              </div>

              <Button
                variant="outline"
                onClick={toggleSortOrder}
                className="flex items-center gap-2 border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100 hover:bg-gray-50 dark:hover:bg-gray-700"
              >
                {sortOrder === "desc" ? <SortDesc className="w-4 h-4" /> : <SortAsc className="w-4 h-4" />}
                Skor {sortOrder === "desc" ? "Azalan" : "Artan"}
              </Button>
            </div>
          </CardContent>
        </Card>

        <Card className="border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800">
          <CardHeader>
            <CardTitle className="flex items-center justify-between text-gray-900 dark:text-gray-100">
              <span>Sonuçlar</span>
              {data && (
                <Badge className="bg-gray-100 dark:bg-gray-700 text-gray-900 dark:text-gray-100">
                  {data.data.pagination.totalRecords} sonuç
                </Badge>
              )}
            </CardTitle>
          </CardHeader>
          <CardContent className="p-6">
            {isLoading ? (
              <div className="space-y-4">
                {Array.from({ length: 5 }).map((_, i) => (
                  <div
                    key={i}
                    className="border border-gray-200 dark:border-gray-700 rounded-lg p-4 space-y-3 bg-white dark:bg-gray-800"
                  >
                    <Skeleton className="h-6 w-3/4 bg-gray-200 dark:bg-gray-700" />
                    <div className="flex gap-2">
                      <Skeleton className="h-5 w-16 bg-gray-200 dark:bg-gray-700" />
                      <Skeleton className="h-5 w-20 bg-gray-200 dark:bg-gray-700" />
                    </div>
                    <div className="flex justify-between items-center">
                      <Skeleton className="h-4 w-32 bg-gray-200 dark:bg-gray-700" />
                      <Skeleton className="h-6 w-12 bg-gray-200 dark:bg-gray-700" />
                    </div>
                  </div>
                ))}
              </div>
            ) : data?.data.items.length === 0 ? (
              <div className="text-center py-12">
                <p className="text-gray-600 dark:text-gray-400 text-lg">Arama kriterlerinize uygun sonuç bulunamadı.</p>
              </div>
            ) : (
              <div className="space-y-4">
                {data?.data.items.map((item) => (
                  <div
                    key={item.id}
                    className="border border-gray-200 dark:border-gray-700 rounded-lg p-6 hover:shadow-md transition-shadow bg-white dark:bg-gray-800"
                  >
                    <div className="flex justify-between items-start mb-4">
                      <h3 className="text-xl font-semibold text-gray-900 dark:text-gray-100 hover:text-blue-600 dark:hover:text-blue-400 transition-colors">
                        {item.title}
                      </h3>
                      <div className="flex items-center gap-1 bg-gray-100 dark:bg-gray-700 px-3 py-1 rounded-full">
                        <Star className="w-4 h-4 text-blue-600 dark:text-blue-400" />
                        <span className="text-sm font-medium text-gray-900 dark:text-gray-100">
                          {formatScore(item.score)}
                        </span>
                      </div>
                    </div>

                    <div className="flex flex-wrap gap-2 mb-4">
                      <Badge
                        className={`${
                          item.type === "video"
                            ? "bg-blue-600 text-white hover:bg-blue-700"
                            : "bg-gray-200 dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                        }`}
                      >
                        {item.type === "video" ? "Video" : "Metin"}
                      </Badge>
                      <Badge className="border-gray-200 dark:border-gray-700 text-gray-600 dark:text-gray-400 bg-transparent">
                        {item.provider}
                      </Badge>
                    </div>

                    <div className="flex flex-wrap gap-1 mb-4">
                      {item.tags.map((tag) => (
                        <Badge
                          key={tag}
                          className="text-xs border-gray-200 dark:border-gray-700 text-gray-600 dark:text-gray-400 bg-transparent"
                        >
                          <Tag className="w-3 h-3 mr-1" />
                          {tag}
                        </Badge>
                      ))}
                    </div>

                    <div className="flex justify-between items-center text-sm text-gray-600 dark:text-gray-400">
                      <div className="flex items-center gap-1">
                        <Calendar className="w-4 h-4" />
                        <span>{formatDate(item.published_at)}</span>
                      </div>
                      <div className="text-xs bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded text-gray-600 dark:text-gray-400">
                        ID: {item.provider_id}
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </CardContent>
        </Card>

        {data && data.data.pagination.totalPages > 1 && (
          <Card className="border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800">
            <CardContent className="p-6">
              <div className="flex flex-col items-center space-y-4">
                <Pagination
                  currentPage={data.data.pagination.pageNumber}
                  totalPages={data.data.pagination.totalPages}
                  onPageChange={setCurrentPage}
                  hasNext={data.data.pagination.hasNext}
                  hasPrev={data.data.pagination.hasPrev}
                />
                <div className="text-sm text-gray-600 dark:text-gray-400">
                  Toplam {data.data.pagination.totalRecords} sonuç
                </div>
              </div>
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  )
}
