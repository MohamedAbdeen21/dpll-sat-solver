<template>
  <header>
    <h1 class="title"> SAT Solver DPLL </h1>
    <ThemeSwitch />
  </header>
  <main>
    <InputField v-bind:getInput="getInput" />
    <ResultForm v-bind:trues='trues' v-bind:falses='falses' v-bind:dcs='dcs' v-bind:solvable='solvable' />
  </main>
</template>

<script setup>
import ResultForm from './components/ResultForm.vue';
import InputField from './components/InputField.vue';
import ThemeSwitch from './components/ThemeSwitch.vue';
</script>

<script>
const server_port = import.meta.env.VITE_SERVER_PORT;
const url = import.meta.env.VITE_URL;

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

      fetch(`${url}:${server_port}/hello`, options)
        .then(resp => resp.json())
        .then((json) => {
          if (json.solved) {
            this.solvable = "Yes";
            this.trues = json.trues.sort().join(', ');
            this.falses = json.falses.sort().join(', ');
            this.dcs = json.dcs.sort().join(', ');
          } else {
            this.solvable = "No";
          }
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
  display: flex;
  border-radius: 20px;
  width: 100%;
}

header .title {
  position: relative;
  margin-left: auto;
  margin-right: auto;
  margin-top: 40px;
  left: 30px;
  text-align: center;
}
</style>

<!-- Global styles -->
<style>
:root {
  --bg: #F5D7E3;
  --primary: #F4A5AE;
  --secondary: #A8577E;
  --input: white;
}

hr {
  width: 80%;
  height: 2px;
  margin: 0 auto;
  margin-top: 15px;
  background-color: var(--secondary);
  border-color: var(--secondary);
}

body {
  background-color: var(--bg);
  font-family: 'Fira Sans', sans-serif;
  margin-right: auto;
  margin-left: auto;
}

.app {
  background-color: var(--primary);
  margin-left: auto;
  margin-right: auto;
  margin-top: 40px;
  width: 70%;
  max-width: 500px;
  min-width: 200px;
  border-radius: 20px;
}
</style>
