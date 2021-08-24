<template>
  <TheHeading />
  <RouterView />
</template>


<script lang="ts" setup>
//imported components are also directly usable in template
import { ref } from 'vue'
import { RouterView } from 'vue-router'
import TheHeading from './components/TheHeading.vue'
</script>
