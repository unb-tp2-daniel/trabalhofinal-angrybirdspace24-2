<template>
    <div>
        <h1>Index</h1>

        <a @click="telaPing">Teste a API</a>

    </div>
</template>

<script setup>
    function telaPing() {
        navigateTo('/ping')
    }
    navigateTo('/login')
</script>