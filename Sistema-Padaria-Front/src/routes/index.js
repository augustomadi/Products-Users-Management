// src/routes/index.js

import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/HomeView.vue' // Usando alias @ para o diretório src
import Product from '@/views/GetProduct.vue'
import User from '@/views/NewUser.vue'
import Newproduct from '@/views/NewProduct.vue'
import DynamicForm from '@/views/DynamicForm.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/user', component: User },
  { path: '/newproduct', component: Newproduct },
  { path: '/product', component: Product },
  { path: '/form', component: DynamicForm }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
