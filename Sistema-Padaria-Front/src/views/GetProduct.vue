<template>
  <div class="form-container">
    <TitleSubtitle :title="formTitle" :subtitle="formSubtitle" />
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
      <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
        <thead
          class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
        >
          <tr>
            <th scope="col" class="px-6 py-3">Nome do Produto</th>
            <th scope="col" class="px-6 py-3">Preço</th>
            <th scope="col" class="px-6 py-3">Ações</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="product in products"
            :key="product.id"
            class="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700"
          >
            <th
              scope="row"
              class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
            >
              {{ product.name }}
            </th>
            <td class="px-6 py-4">{{ product.price }}</td>
            <td class="px-6 py-4">
              <a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
                >Edit</a
              >
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="button-container">
      <router-link to="/form">
        <button type="button" class="return-button">Voltar</button>
      </router-link>
      <button type="button" class="delete-button" @click="deleteProduct">Deletar</button>
    </div>
  </div>
</template>

<script>
import TitleSubtitle from '../components/TitleSubtitle.vue'
import axios from 'axios'

export default {
  components: {
    TitleSubtitle
  },
  data() {
    return {
      formTitle: 'Produtos Encontrados',
      formSubtitle: 'Lista de produtos encontrados na pesquisa',
      products: [] // Array para armazenar os produtos
    }
  },
  created() {
    this.fetchProducts()
  },
  methods: {
    async fetchProducts() {
      const query = this.$route.query.q
      try {
        const token = localStorage.getItem('user-token')
        const response = await axios.get(`http://localhost:8080/products?search=${query}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        this.products = response.data
      } catch (error) {
        console.log('Error fetching products:', error)
      }
    },
    deleteProduct() {
      console.log('Deletar produto')
    }
  }
}
</script>

<style scoped>
.form-container {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 100vh;
  box-sizing: border-box;
  overflow: hidden;
  font-family: Arial, Helvetica, sans-serif;
}

.form-container::before {
  content: '';
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(300deg, #d2b48c, #f5f5dc, #8b4513);
  background-size: 250% 400%;
  animation: gradientAnimation 20s ease infinite;
  z-index: -1;
}

@keyframes gradientAnimation {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

.form-container {
  background-color: #fff;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 50%;
  max-width: 600px;
  text-align: left;
}

.product-item {
  margin-bottom: 20px;
}

.button-container {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

.return-button,
.delete-button {
  width: 100px;
  height: 50px;
  border: none;
  border-radius: 8px;
  color: #fff;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s;
}

.return-button {
  background-color: #3c3c3c;
}

.return-button:hover {
  background-color: #2a2a2a;
}

.delete-button {
  background-color: #ff4d4d;
}

.delete-button:hover {
  background-color: #cc0000;
}
</style>
