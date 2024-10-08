<template>
  <div class="form-container">
    <form @submit.prevent="loginUser">
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
      <router-link to="/user">
        <button type="button" class="return-button">Crie um Usuario</button>
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
      formTitle: 'Login de Usuário',
      fields: [
        {
          key: 'email',
          type: 'email',
          placeholder: 'Digite o e-mail',
          validator: (value) => value.length > 0,
          errorMessage: 'Cadastre um e-mail válido',
          datalist: [],
          listId: '',
          required: true
        },
        {
          key: 'pass',
          type: 'password',
          placeholder: 'Defina sua Senha',
          validator: (value) => value.length >= 0,
          errorMessage: 'Senha deve ter no mínimo 8 caracteres.',
          required: true
        }
      ],
      formData: {
        email: '',
        pass: ''
      },
      fieldHasErrors: {},
      responseMessage: ''
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
    handleInputChange(key, value) {
      this.formData[key] = value
    },
    async loginUser() {
      try {
        const response = await axios.post('http://localhost:8080/user/login', {
          email: this.formData.email,
          password: this.formData.pass
        })
        this.responseMessage = response.data.message
        console.log(this.responseMessage)

        // Verifique se o token está presente na resposta e armazene-o no localStorage
        if (response.data.token) {
          localStorage.setItem('user-token', response.data.token)
          this.$router.push('/form') // Redireciona o usuário para a rota /products
        } else {
          alert('Login bem-sucedido, mas nenhum token foi retornado.')
        }
      } catch (error) {
        if (error.response) {
          this.responseMessage = error.response.data.message || error.response.data
        } else if (error.request) {
          this.responseMessage = 'Login failed: No response received from server.'
        } else {
          this.responseMessage = error.message
        }
        alert(this.responseMessage)
        console.log(this.formData.email, this.formData.pass)
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
