<template>
  <div class="input-group" :class="{ 'has-error': localHasError && isTouched }">
    <input
      ref="inputElement"
      :type="type"
      :value="internalValue"
      :placeholder="placeholder"
      :required="required"
      :aria-invalid="localHasError && isTouched"
      :list="listId"
      autocomplete="off"
      @input="updateValue($event.target.value)"
      @blur="handleBlur"
      @focus="handleFocus"
    />
    <datalist v-if="datalist && datalist.length > 0" :id="listId">
      <option v-for="item in datalist" :key="item.value" :value="item.value">
        {{ item.label }}
      </option>
    </datalist>
    <span class="error-message" v-if="localHasError && isTouched">
      {{ fullErrorMessage }}
    </span>
  </div>
</template>

<script>
export default {
  name: 'FormInput',
  props: {
    modelValue: [String, Number],
    placeholder: String,
    required: Boolean,
    type: {
      type: String,
      default: 'text'
    },
    hasError: Boolean,
    errorMessage: {
      type: String,
      default: 'Invalid input'
    },
    validator: {
      type: Function,
      default: () => true
    },
    datalist: {
      type: Array,
      default: () => []
    },
    listId: {
      type: String,
      default: () => `list-${Math.random().toString(36).substring(2, 15)}`
    }
  },
  emits: ['update:modelValue', 'update:hasError'],
  data() {
    return {
      internalValue: this.modelValue,
      localHasError: this.hasError,
      isTouched: false
    }
  },
  methods: {
    updateValue(value) {
      this.internalValue = value
      this.$emit('update:modelValue', value)
    },
    handleBlur() {
      this.isTouched = true
      this.localHasError = !this.validator(this.internalValue) && this.internalValue.length > 0
      this.$emit('update:hasError', this.localHasError)
    },
    handleFocus() {
      this.isTouched = true
    }
  },
  watch: {
    modelValue(newValue) {
      this.internalValue = newValue
    },
    hasError(newValue) {
      this.localHasError = newValue
    }
  },
  computed: {
    fullErrorMessage() {
      return this.errorMessage || 'Error'
    }
  }
}
</script>
<style scoped>
.input-group {
  position: relative;
}

.input-group input {
  width: 100%;
  padding: 20px;
  font-size: 20px;
  font-family: Arial, Helvetica, sans-serif;
  border: 0px;
  border-bottom: 1px solid #ddd;
  box-sizing: border-box;
}

.input-group input::placeholder {
  color: #999;
}

.input-group input:focus {
  outline: transparent;
  border-bottom: 1.25px solid #c5c5c5;
}

.error-message {
  color: red;
  font-size: 0.75rem;
  opacity: 0;
  transition: opacity 0.3s;
  position: 0 0 0;
}

.has-error .error-message {
  opacity: 1;
}

.input-group.has-error input {
  border-color: red;
}
</style>
