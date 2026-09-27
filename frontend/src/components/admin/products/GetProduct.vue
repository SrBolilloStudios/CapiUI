<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { TableColumn } from '@nuxt/ui'
interface Product {
  id: number
  sku: string
  name: string
  price: number
  description: string
  img: string
  category: Category
  stock: number
}

interface Category {
  id: number,
  name: string
}

const url = 'http://localhost:8080/api/products'
const page = ref(1)
const limit = ref(10)

const products = ref<Product[] | null>(null)
const error = ref<Error | unknown | null>(null)

const columns: TableColumn<Product>[] = [
  { accessorKey: 'id', header: 'Id' },
  { accessorKey: 'sku', header: 'Sku' },
  { accessorKey: 'name', header: 'Name' },
  { accessorKey: 'price', header: 'Price' },
  { accessorKey: 'description', header: 'Description' },
  { accessorKey: 'img', header: 'Img' },
  {
    accessorKey: 'category',
    header: 'Category',
    cell: ({ row }) => row.original.category.name,
  },
  { accessorKey: 'stock', header: 'Stock' },
]

async function fetchProducts() {
  try {
    const params = new URLSearchParams({
      page: page.value.toString(),
      limit: limit.value.toString(),
    })

    const res = await fetch(`${url}?${params.toString()}`)

    if (!res.ok) {
      throw new Error('Error en la petición')
    }

    const data = (await res.json()) as Product[]
    products.value = data
  } catch (err) {
    error.value = err
  }
}

onMounted(() => {
  fetchProducts()
})
</script>

<template>
  <UTable :data="products ?? []" :columns="columns" class="flex-1" />
</template>

<style scoped></style>
