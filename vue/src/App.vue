<template>
  <header>
    <h1 class="title"> SAT Solver DPLL </h1>
  </header>
  <main>
    <InputField v-bind:getInput="getInput" />
    <ResultForm v-bind:trues='trues' v-bind:falses='falses' v-bind:dcs='dcs' v-bind:solvable='solvable' />
  </main>
</template>

<script setup>
import ResultForm from './components/ResultForm.vue';
import InputField from './components/InputField.vue'
</script >

<script>
const server_port = 3000
export default {
  data: function () {
    return {
      trues: '',
      falses: '',
      dcs: '',
      solvable: ''
    }
  },
  methods: {
    solve(input) {
      const options = {
        "method": "POST",
        "mode": "cors",
        "headers": {
          "Accept": "application/json",
          "Content-Type": "text/plain",
        },
        "body": input
      };

      fetch(`http://127.0.0.1:${server_port}/hello`, options)
        .then(resp => resp.json())
        .then((json) => {
          this.solvable = json.solved.toString();
          this.trues = json.trues.sort().join();
          this.falses = json.falses.sort().join();
          this.dcs = json.dcs.sort().join();
        });
    },
    getInput(input) {
      this.solve(input);
    }
  },
};
</script>
<style scoped>
main {
  margin-left: auto;
  margin-right: auto;
  width: 80%;
}

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
</style>
