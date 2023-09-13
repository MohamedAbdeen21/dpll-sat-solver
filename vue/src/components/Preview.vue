<template>
  <div class="preview-container">
    <label for="preview"> Preview </label>
    <input id="preview" class="preview" type="text" v-bind:value="formula" readonly>
  </div>
</template>

<script setup>
defineProps({
  formula: {
    type: String,
    required: true,
    default: ''
  },
})
</script>

<style scoped>
.preview-container {
  padding-top: 10px;
  border-radius: 15px;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 10px;
  margin-top: 5px;
  text-align: center;
}

label {
  font-size: 20px;
  color: black;
}

input {
  background-color: var(--input);
  text-align: center;
  width: 80%;
  margin-left: auto;
  margin-right: auto;
  margin-top: 10px;
  border-radius: 10px;
}
</style>
