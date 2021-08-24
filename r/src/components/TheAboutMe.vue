<template>
    <div class="wr-wrap">
        <h1>Age: {{ age }}.</h1>
    </div>
</template>

<script lang="ts" setup>
const age = (new Date()).getFullYear() - 1995
</script>

<style scoped>
.wr-wrap {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #1174d8;
    margin-top: 60px;
}
</style>