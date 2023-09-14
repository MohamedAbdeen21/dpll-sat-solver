<template>
  <div class="input" v-once contenteditable="true" @focusin="clearInput" @input="updateInput" @focusout="resetInput"
    v-html="input">
  </div>
  <hr>
  <Preview v-bind:formula="formula" />
  <button @click="getInput(input)" type="submit"> Solve </button>
</template>

<script setup>
import Preview from './Preview.vue'
defineProps({
  getInput: Function
})
</script>

<script>
const INPUT_PLACEHOLDER = "One line per clause, separate literals by spaces"

export default {
  data: function () {
    return {
      input: INPUT_PLACEHOLDER,
      formula: '',
      color: "grey",
      fontStyle: "italic"
    }
  },
  methods: {
    clearInput(event) {
      if (event.target.innerText.trim() !== INPUT_PLACEHOLDER) {
        return
      }
      event.target.innerHTML = ""
      this.formula = ""
      this.color = "black"
      this.fontStyle = "normal"
    },
    resetInput(event) {
      if (event.target.innerText.trim() !== "") {
        return
      }
      event.target.innerHTML = INPUT_PLACEHOLDER
      this.formula = ""
      this.color = "grey"
      this.fontStyle = "italic"
    },
    updateInput(event) {
      this.input = event.target.innerHTML
        .split('<div>')
        .map(line => line.replace(/&nbsp;/g, ' ').replace(/\s+/g, ' ').replace('</div>', '').replace('<br>', '').trim())
        .join("\n");

      this.formula = this.input
        .split('\n')
        .filter((line) => line !== "")
        .map((line) => `(${line})`.replace(/\s/g, ' v '))
        .join(' ^ ');
    },
  },
};
</script>
<style scoped>
.input {
  margin-left: auto;
  margin-right: auto;
  width: 80%;
  min-height: 50px;
  background-color: var(--input);
  padding: 15px;
  border: 1px solid white;
  border-radius: 20px;
  color: v-bind(color);
  font-style: v-bind(fontStyle);
}

button {
  background-color: var(--primary);
  border-radius: 20px;
  border: 1px solid white;
  padding: 15px 32px;
  text-align: center;
  display: flex;
  margin-left: auto;
  margin-right: auto;
  margin-top: 10px;
  margin-bottom: 10px;
  font-size: 16px;
}

button:hover {
  transition: 500ms ease;
  background-color: var(--secondary);
  border: 1px solid black;
  color: white;
}

button:active {
  transform: scale(0.9);
}
</style>
