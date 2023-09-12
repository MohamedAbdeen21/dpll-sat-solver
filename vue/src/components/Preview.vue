<template>
  <div class="preview-container">
    <form>
      <label for="preview"> Preview </label>
      <input id="preview" class="preview" type="text" v-bind:value="formula" readonly>
    </form>
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
  text-align: center;
  padding-top: 10px;
  border-radius: 15px;
  display: contents;
  padding: 10px;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 10px;
  margin-top: 10px;
  text-align: center;
}

label {
  font-size: 20px;
  color: black;
}

input {
  display: block;
  text-align: center;
  width: 80%;
  margin-left: auto;
  margin-right: auto;
  margin-top: 10px;
  border-radius: 5px;
}
</style>
