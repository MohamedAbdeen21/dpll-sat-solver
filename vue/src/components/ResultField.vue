<template>
  <form v-show="values">
    <label for="field"> {{ label }} </label>
    <input id="field" type="text" v-bind:value="values" readonly>
  </form>
</template>

<script setup>
defineProps({
  label: {
    type: String,
    required: true
  },
  values: {
    type: String,
    required: true,
    default: ''
  },
})
</script>

<style scoped>
label {
  font-size: 20px;
  color: black;
  text-align: center;
  margin-left: auto;
  margin-right: auto;
}

input {
  background-color: var(--input);
  display: block;
  text-align: center;
  margin-left: auto;
  margin-right: auto;
  width: 80%;
  margin-left: auto;
  margin-top: 10px;
  border: 1px solid white;
  border-radius: 20px;
}

form {
  border-radius: 15px;
  background-color: var(--secondary);
  text-align: center;
  padding: 10px;
  width: 90%;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 10px;
  margin-top: 10px;
}
</style>
