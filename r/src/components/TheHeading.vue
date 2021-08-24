<template>
    <div class="wr-wrap">
        <RouterLink to="/">{{ home }}</RouterLink>
        <RouterLink to="/tagged">{{ tagged }}</RouterLink>
        <RouterLink to="/about">{{ about }}</RouterLink>
    </div>
</template>


<script lang="ts" setup>
import { RouterLink } from 'vue-router'

const home = "首页",
    tagged = "标签",
    about = "关于"
</script>


<style scoped>
.wr-wrap {
    display: -webkit-flex; /* Safari */
    display: flex;
}
</style>