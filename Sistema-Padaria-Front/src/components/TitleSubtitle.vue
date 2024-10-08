<template>
  <div class="title-subtitle-container flex items-center justify-center flex-col">
    <img
      src="https://i.pinimg.com/originals/1f/8d/d5/1f8dd567294d173583e9135b5f788913.jpg"
      alt="Logo"
      class="logo"
    />
    <h1 class="title">{{ title }}</h1>
    <p class="subtitle">{{ subtitle }}</p>
  </div>
</template>
<script>
export default {
  name: 'TitleSubtitle',
  props: {
    title: String,
    subtitle: String
  }
}
</script>
<style>
.title-subtitle-container {
  text-align: center;
}

.logo {
  max-width: 150px;
  height: auto;
  margin-bottom: 1rem;
  transition:
    transform 0.3s ease-in-out,
    box shadow 0.3s ease-in-out;
}

.title {
  color: #000;
  font-size: 32px;
  font-family: Arial, Helvetica, sans-serif;
  margin-bottom: 0.25rem;
  margin-top: 0;
  letter-spacing: 1.5px;
}

.subtitle {
  color: #000;
  font-size: 20px;
  margin-top: 0;
  margin-bottom: 1.5rem;
  letter-spacing: 0.5px;
}
</style>
