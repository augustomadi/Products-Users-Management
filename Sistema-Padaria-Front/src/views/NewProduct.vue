<template>
  <div class="form-container">
    <form @submit.prevent="registerProduct">
      <router-link to="/form">
        <button type="button" class="back-button">Voltar</button>
      </router-link>
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
      <button type="submit" class="delete-button" :disabled="!isFormValid">Atualizar</button>
    </form>
  </div>
</template>

<script>
import TitleSubtitle from '../components/TitleSubtitle.vue'
import FormInput from '../components/FormInput.vue'
import axios from 'axios'

axios.defaults.headers.post['Content-Type'] = 'application/json'

export default {
  components: {
    TitleSubtitle,
    FormInput
  },
  data() {
    return {
      formTitle: 'Cadastre novo produto',
      formSubtitle: '',
      fields: [
        {
          key: 'name',
          type: 'text',
          placeholder: 'Nome do produto',
          datalist: [],
          listId: '',
          required: true,
          validator: (value) => typeof value === 'string' && !/\d/.test(value),
          errorMessage: 'Nome inválido'
        },
        {
          key: 'quantity',
          type: 'number',
          placeholder: 'Quantidade em estoque',
          datalist: [],
          listId: '',
          required: true,
          validator: (value) => parseInt(value) > 0,
          errorMessage: 'Quantidade inválida'
        },
        {
          key: 'price',
          type: 'number',
          placeholder: 'Preço',
          datalist: [],
          listId: '',
          required: true,
          validator: (value) => parseFloat(value) > 0,
          errorMessage: 'Preço inválido'
        }
      ],
      formData: {
        name: '',
        quantity: '',
        price: ''
      },
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
  methods: {
    validateField(key, value) {
      const field = this.fields.find((f) => f.key === key)
      if (field && field.validator) {
        this.fieldHasErrors[key] = !field.validator(value)
      } else {
        this.fieldHasErrors[key] = false
      }
    },
    async registerProduct() {
      try {
        const productData = {
          name: this.formData.name,
          quantity: parseInt(this.formData.quantity, 10),
          price: parseFloat(this.formData.price)
        }

        console.log('Dados a serem enviados:', productData)

        await axios.post('http://localhost:8080/newproduct/create', productData, {
          headers: {
            'Content-Type': 'application/json'
          }
        })
        alert('Produto registrado com sucesso')
        console.log('Produto registrado com sucesso')
        this.$router.push('/form') // Redireciona o usuário para a página de sucesso
      } catch (error) {
        if (error.response) {
          alert(`Erro: ${error.response.data}`)
          console.error('Error response data:', error.response.data)
        } else if (error.request) {
          alert('Falha no registro: Nenhuma resposta recebida do servidor.')
          console.error('Error request:', error.request)
        } else {
          alert(`Erro: ${error.message}`)
          console.error('Error message:', error.message)
        }
      }
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

.delete-button {
  background-color: #3c2414;
}

.delete-button:hover {
  background-color: #704d37;
}
</style>
