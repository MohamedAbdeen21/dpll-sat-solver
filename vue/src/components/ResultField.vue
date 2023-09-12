<template>
  <div class="field" v-show="values">
    <form>
      <label> {{ label }} </label>
      <br>
      <input type="text" v-bind:value="values" readonly>
    </form>
  </div>
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

form {
  border-radius: 15px;
  background-color: #A8577E;
  padding: 10px;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 10px;
  margin-top: 10px;
}
</style>
