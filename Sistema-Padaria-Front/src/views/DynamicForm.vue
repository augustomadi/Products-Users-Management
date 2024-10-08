<template>
  <div class="form-container">
    <form @submit.prevent="SearchProduct">
      <div class="form-header">
        <router-link to="/">
          <button type="button" class="back-button">Voltar</button>
        </router-link>
      </div>
      <TitleSubtitle :title="formTitle" :subtitle="formSubtitle" />

      <FormInput
        v-for="input in fields"
        :key="input.key"
        v-model="formData[input.key]"
        :type="input.type"
        :placeholder="input.placeholder"
        :validator="input.validator"
        :errorMessage="input.errorMessage"
        :datalist="input.datalist"
        :listId="input.listId"
        :required="input.required"
        v-model:hasError="fieldHasErrors[input.key]"
        @input="validateField(input.key, formData[input.key])"
        @blur="validateField(input.key, formData[input.key])"
      />

      <button type="submit" class="submit-button">&gt;</button>
      <router-link to="/newproduct">
        <button type="button" class="delete-button" @click="deleteProduct">Cadastrar novo</button>
      </router-link>
    </form>
  </div>
</template>

<script>
import TitleSubtitle from '../components/TitleSubtitle.vue'
import FormInput from '../components/FormInput.vue'
import axios from 'axios'

export default {
  components: {
    TitleSubtitle,
    FormInput
  },
  data() {
    return {
      formTitle: 'Pesquise seu produto',
      formSubtitle: 'Somente permitidos',
      fields: [
        {
          key: 'name',
          type: 'text',
          placeholder: 'Digite um produto',
          validator: (value) => value.length > 0,
          errorMessage: 'Digite um produto',
          datalist: [],
          listId: '',
          required: true
        }
      ],
      formData: {},
      fieldHasErrors: {}
    }
  },
  computed: {
    isFormValid() {
      return this.fields.every((field) =>
        field.required ? this.formData[field.key] && !this.fieldHasErrors[field.key] : true
      )
    }
  },
  mounted() {
    this.tokenExist()
  },
  methods: {
    async SearchProduct() {
      console.log('SearchProduct called') // Log de depuração
      console.log('Form Data:', this.formData) // Log para ver os dados do formulário

      try {
        const response = await axios.get('http://localhost:8080/search', {
          params: {
            search: this.formData.name
          },
          headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
          }
        })
        console.log('Product data:', response.data) // Log da resposta da API
        // Redireciona para a página de produtos com o parâmetro de consulta
        this.$router.push({ path: '/product', query: { q: this.formData.name } })
      } catch (error) {
        if (error.response) {
          alert(`Erro: ${error.response.data}`)
          console.error('Error response data:', error.response.data)
        } else if (error.request) {
          alert('Falha ao buscar produtos: Nenhuma resposta recebida do servidor.')
          console.error('Error request:', error.request)
        } else {
          alert(`Erro: ${error.message}`)
          console.error('Error message:', error.message)
        }
      }
    },
    async tokenExist() {
      var tokenUser = localStorage.getItem('user-token')
      if (!tokenUser) {
        this.$router.push('/')
      }
    },
    validateField(key, value) {
      const field = this.fields.find((f) => f.key === key)
      if (field && field.validator) {
        this.fieldHasErrors[key] = !field.validator(value)
      } else {
        this.fieldHasErrors[key] = false
      }
    },
    submitForm() {
      let hasError = false
      for (const field of this.fields) {
        this.validateField(field.key, this.formData[field.key])
        if (this.fieldHasErrors[field.key]) {
          hasError = true
        }
      }
      if (!hasError) {
        this.handleFormSubmission(this.formData)
      }
    },
    handleFormSubmission(data) {
      console.log('Form data:', data)
    },
    deleteProduct() {
      console.log('Cadastrar novo produto')
    }
  }
}
</script>

<style scoped>
.form-container {
  width: 100%;
  display: flex;
  justify-content: center;
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

.form-container form {
  background-color: #fff;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 25% !important;
  max-width: 600px;
  position: relative; /* Adicionado para permitir posicionamento do header */
}

.form-header {
  position: absolute;
  top: 30px; /* Ajuste conforme necessário */
  left: 40px; /* Ajuste conforme necessário */
}

.back-button {
  width: 60px;
  height: 45px;
  border: none;
  border-radius: 8px;
  background-color: #3c2414;
  color: #fff;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s;
}

.back-button:hover {
  background-color: #704d37;
}

.submit-button {
  bottom: 20px;
  width: 50px;
  height: 50px;
  border: none;
  border-radius: 50%;
  background-color: #3c2414;
  color: #fff;
  font-size: 25px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s;
  margin-top: 40px;
  margin-bottom: 0;
  float: right;
}

.submit-button:hover {
  background-color: #704d37;
}

.submit-button:disabled {
  color: #fff;
  cursor: not-allowed;
  opacity: 75%;
  background-color: #3c2414;
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
  margin-top: 40px;
}

.return-button {
  background-color: #3c2414;
}

.return-button:hover {
  background-color: #704d37;
}

.delete-button {
  background-color: #3c2414;
}

.delete-button:hover {
  background-color: #704d37;
}
</style>
