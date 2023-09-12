<template>
  <div class="output">
    <ResultField label='Has solution?' v-bind:values="solvable" />
    <ResultField label='True' v-bind:values="trues" />
    <ResultField label='False' v-bind:values='falses' />
    <ResultField label="Don't Care" v-bind:values='dcs' />
  </div>
</template>

<script setup>
import ResultField from './ResultField.vue'
defineProps({
  solvable: {
    type: String,
    default: '',
  },
  trues: {
    type: String,
    default: ''
  },
  falses: {
    type: String,
    default: ''
  },
  dcs: {
    type: String,
    default: ''
  }
})
</script>

<style scoped>
.output {
  padding: 10px;
  margin-left: auto;
  margin-right: auto;
}
</style>
