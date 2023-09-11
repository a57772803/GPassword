<script lang="ts" setup>
import Sidebar from "./components/Sidebar.vue";
</script>



<template>
  <div class="flex-1 flex">
    <Sidebar></Sidebar>
    <router-view></router-view>
  </div>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
