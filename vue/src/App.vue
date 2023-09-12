<template>
  <header>
    <h1 class="title"> SAT Solver DPLL </h1>
  </header>
  <main>
    <div class="formula" contenteditable="true" @input="updateInput"> </div>
    <!-- <textarea id="formula" class="formula" type="text" v-model='input' placeholder="Enter the formula here" autofocus -->
    <!--   rows=10 cols=50>  </textarea> -->
    <button @click="solve" type="submit"> Solve </button>
    <ResultForm v-bind:trues='trues' v-bind:falses='falses' v-bind:dcs='dcs' v-bind:solvable='solvable' />
  </main>
</template>

<script setup>
import ResultForm from './components/ResultForm.vue';
</script >

<script>
const server_port = 3000
export default {
  data: function () {
    return {
      trues: '',
      falses: '',
      dcs: '',
      solvable: '',
      input: '',
    }
  },
  methods: {
    solve() {
      this.input = this.input
        .split('<div>')
        .map(line => line.replace(/&nbsp;/g, '').trim().replace('</div>', ''))
        .filter((line) => line !== "<br>")
        .sort()
        .join("\n");

      console.log(this.input)
      const options = {
        "method": "POST",
        "mode": "cors",
        "headers": {
          "Accept": "application/json",
          "Content-Type": "text/plain",
        },
        "body": this.input
      };

      fetch(`http://127.0.0.1:${server_port}/hello`, options)
        .then(resp => resp.json())
        .then((json) => {
          this.trues = json.trues.join();
          this.falses = json.falses.join();
          this.dcs = json.dcs.join();
          this.solvable = json.solved.toString();
        });
    },
    updateInput(event) {
      this.input = event.target.innerHTML;
    },
    getLines() {
    }
  },
};
</script>
<style scoped>
header {
  display: contents;
  border-radius: 20px;
}

header .title {
  margin-left: auto;
  margin-right: auto;
  padding-top: 20px;
  text-align: center;
}

.formula {
  margin-left: auto;
  margin-right: auto;
  width: 80%;
  min-height: 50px;
  background-color: white;
  padding: 15px;
  border: 1px solid white;
  border-radius: 20px;
}

main {
  margin-left: auto;
  margin-right: auto;
  width: 80%;
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
