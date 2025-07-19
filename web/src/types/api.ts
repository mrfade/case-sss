export type ApiResponse<T> = {
  success: boolean
  message: string
  data: {
    meta: {
      searchable: string[]
      filterable: string[]
      sortable: string[]
    }
    pagination: {
      pageNumber: number
      pageSize: number
      totalPages: number
      totalRecords: number
      hasNext: boolean
      hasPrev: boolean
    }
    items: T
  }
}
