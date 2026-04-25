<template>
  <div>
    <h1>Teste de API</h1>

    <p v-if="pending">
      Carregando...
    </p>

    <p v-else-if="error">
      Erro ao conectar à API. Recarregue a página
    </p>

    <pre v-else>
        <p>Ebaaa 👅</p>
        {{ data }}
    </pre>
  </div>
</template>

<script setup>
  const { data, pending, error } = await useFetch('/api/pong')

</script>