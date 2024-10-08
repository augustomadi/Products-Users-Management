<template>
  <div class="form-container">
    <form @submit.prevent="submitForm">
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

      <button type="submit" class="submit-button" :disabled="!isFormValid">&gt;</button>
    </form>
  </div>
</template>

<script>
import TitleSubtitle from '../components/TitleSubtitle.vue'
import FormInput from '../components/FormInput.vue'

export default {
  components: {
    TitleSubtitle,
    FormInput
  },
  data() {
    return {
      formTitle: 'Pesquise um produto',
      submitUrl: 'http://localhost:8080/search',
      fields: [
        {
          key: 'name',
          type: 'text',
          placeholder: 'Digite aqui',
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
  methods: {
    handleFormSubmission(data) {
      var requestOptions = {
        method: 'POST',
        body: JSON.stringify(data),
        redirect: 'follow'
      }

      fetch(this.submitUrl, requestOptions)
        .then((response) => response.text())
        .then((result) => {
          console.log(result)
          this.$router.push('/product')
        })
        .catch((error) => console.log('error', error))
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

.delete-button {
  background-color: #3c2414;
}

.delete-button:hover {
  background-color: #704d37;
}
</style>
