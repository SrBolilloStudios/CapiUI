<script setup lang="ts">
import { ref, onMounted } from 'vue';

interface Product {
    id: number
    name: string
    price: number
}

const data = ref<Product[] | null>(null)
const status = ref<string | null>("")
const isLoading = ref<boolean | null>(true)
async function fetchProducts() {
    try {
        const url = "http://localhost:8080/api/test/products"
        const request = await fetch(url)
        if (!request.ok) {
            status.value = "Error de peticion"
            return
        }
        data.value = await request.json()
        isLoading.value = false
    }
    catch (err) {
        status.value = String(err)
        isLoading.value = false
    }
}
onMounted(() => fetchProducts())
</script>
<template>
    <div class="m-8" flex>
        <UTable :data="data" class="flex-1"/>
        <UButton>Button</UButton>
    </div>
</template>