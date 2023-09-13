<template>
  <button type="submit" class="theme-switch" v-if="isDark" @click="{ toggle(); enableLight() }">
    <i class="fa-solid fa-sun"></i>
    <i class="fa-solid fa-sun fa-bounce"></i>
    <i class="fa-regular fa-moon"></i>
  </button>
  <button type="submit" class="theme-switch" v-if="!isDark" @click="{ toggle(); enableDark() }">
    <i class="fa-solid fa-moon"></i>
    <i class="fa-solid fa-moon fa-bounce"></i>
    <i class="fa-regular fa-sun"></i>
  </button>
</template>

<script setup>
import { useDark, useToggle } from '@vueuse/core';
</script >

<script>
const isDark = useDark();
const toggle = useToggle(isDark);

export default {
  methods: {
    enableDark() {
      document.documentElement.style.setProperty("--bg", "#352F44");
      document.documentElement.style.setProperty("--primary", "#B9B4C7");
      document.documentElement.style.setProperty("--secondary", "#5C5470");
      document.documentElement.style.setProperty("--input", "#CCCCCC");
    },
    enableLight() {
      document.documentElement.style.setProperty("--bg", "#F5D7E3");
      document.documentElement.style.setProperty("--primary", "#F4A5AE");
      document.documentElement.style.setProperty("--secondary", "#A8577E");
      document.documentElement.style.setProperty("--input", "#FFFFFF");
    },
  },
  mounted() {
    isDark.value ? this.enableDark() : this.enableLight();
  },
}
</script >

<style scoped>
.theme-switch {
  height: 20px;
  width: 15px;
  margin-left: 15px;
  margin-right: 15px;
  margin-top: 15px;
  justify-content: center;
  display: flex;
  background: none;
  border: none;
}

.theme-switch i:nth-child(2) {
  display: none;
}

.theme-switch i:nth-child(3) {
  display: none;
}

.theme-switch:hover i:nth-child(1) {
  display: none;
}

.theme-switch:hover i:nth-child(2) {
  display: block;
}

.theme-switch:active i:nth-child(1) {
  display: none;
  transition-duration: 1s linear;
}

.theme-switch:active i:nth-child(2) {
  display: none;
  transition-duration: 1s linear;
}

.theme-switch:active i:nth-child(3) {
  display: block;
  transition-duration: 1s linear;
}
</style>
