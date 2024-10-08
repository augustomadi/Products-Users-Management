<template>
  <div class="form-container">
    <form @submit.prevent="submitForm">
      <TitleSubtitle :title="formTitle" :subtitle="formSubtitle" />

      <div v-for="input in inputs" :key="input.key">
        <FormInput
          :key="input.key"
          v-model="formData[input.key]"
          :type="input.type"
          :placeholder="input.placeholder"
          :validator="input.validator"
          :errorMessage="input.errorMessage"
          :datalist="input.datalist"
          :listId="input.listId"
          :required="input.required"
          @input="validateField(input.key, formData[input.key])"
          @blur="validateField(input.key, formData[input.key])"
        />
        <span v-if="errors[input.key]" class="error-message">{{ errors[input.key] }}</span>
      </div>
      <button type="submit" class="submit-button" :disabled="isSubmitDisabled">&gt;</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios'
import TitleSubtitle from './TitleSubtitle.vue'
import FormInput from './FormInput.vue'

export default {
  components: {
    TitleSubtitle,
    FormInput
  },
  props: {
    formTitle: {
      type: String,
      required: true
    },
    formSubtitle: {
      type: String,
      required: true
    },
    inputs: {
      type: Array,
      required: true,
      default: () => []
    },
    submitUrl: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      formData: {},
      validFields: {},
      errors: {},
      fieldHasErrors: {},
      loading: false
    }
  },
  computed: {
    isFormValid() {
      return Object.values(this.validFields).every((isValid) => isValid)
    },
    isSubmitDisabled() {
      return !this.isFormValid || this.loading
    }
  },
  methods: {
    initializeFormData() {
      if (!Array.isArray(this.inputs)) {
        console.error('Inputs prop should be an array.')
        return
      }
      this.inputs.forEach((input) => {
        this.formData[input.key] = input.defaultValue || ''
        this.validFields[input.key] = !input.required
        this.fieldHasErrors[input.key] = false
      })
    },
    validateField(key, value) {
      const input = this.inputs.find((input) => input.key === key)
      if (input && input.validator) {
        const isValid = input.validator(value)
        this.validFields[key] = isValid
        this.fieldHasErrors[key] = !isValid
        this.errors[key] = isValid ? '' : input.errorMessage
      }
    },
    submitForm() {
      if (!this.isFormValid) {
        alert('Please correct the errors before submitting.')
        return
      }
      this.loading = true

      console.log(`Submitting to URL: ${this.submitUrl}`)

      axios
        .post(this.submitUrl, this.formData, {
          headers: {
            'Content-Type': 'application/json;charset=UTF-8'
          }
        })
        .then((response) => {
          this.$emit('formSubmitted', response.data)
          this.$router.push('/product')
        })
        .catch((error) => {
          console.error('Submission error:', error.response ? error.response.data : 'Network error')
        })
        .finally(() => {
          this.loading = false
        })
    }
  },
  mounted() {
    this.initializeFormData()
  }
}
</script>

<style scoped>
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
  opacity: 60%;
  background-color: #3c2414;
}

.input-group input {
  margin-bottom: 20px;
}

.error-message {
  color: red;
  font-size: 12px;
}
</style>
