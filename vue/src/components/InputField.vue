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
const input_placeholder = "One line per clause, literals separated by space"

export default {
  data: function () {
    return {
      input: input_placeholder,
      formula: '',
      color: "grey",
      font_style: "italic"
    }
  },
  methods: {
    clearInput(event) {
      if (event.target.innerHTML === input_placeholder) {
        event.target.innerHTML = ""
        this.formula = ""
        this.color = "black"
        this.font_style = "normal"
      }
    },
    resetInput(event) {
      if (event.target.innerHTML === "") {
        event.target.innerHTML = input_placeholder
        this.formula = ""
        this.color = "grey"
        this.font_style = "italic"

      }
    },
    updateInput(event) {
      this.input = event.target.innerHTML
        .split('<div>')
        .map(line => line.replace(/&nbsp;/g, ' ').replace('</div>', '').trim())
        .filter((line) => line !== "<br>")
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
  background-color: white;
  padding: 15px;
  border: 1px solid white;
  border-radius: 20px;
  color: v-bind(color);
  font-style: v-bind(font_style);
}

hr {
  width: 80%;
  height: 2px;
  margin: 0 auto;
  margin-top: 15px;
  background-color: #A8577E;
  border-color: #A8577E;
}

button {
  background-color: #F4A5AE;
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
  background-color: #A8577E;
  border: 1px solid black;
}
</style>
