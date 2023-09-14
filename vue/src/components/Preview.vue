<template>
  <div class="preview-container">
    <label for="preview"> Preview </label>
    <textarea id="preview" class="preview" v-bind:value="formula" rows="4" cols="50" readonly></textarea>
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

textarea {
  background-color: var(--input);
  text-align: center;
  width: 80%;
  margin-left: auto;
  margin-right: auto;
  margin-top: 10px;
  border-radius: 10px;
  resize: vertical;
}
</style>
