<template>
    <div>
    <Header />
    <main class="container">
      <TabelaBuscas/>
        <div v-for="turma in turmas">
            <p>Turma: {{ turma.nomeMateria }}</p>
        </div>
    </main>
    <Footer />
    </div>
</template>

<script setup>
import {ref , onMounted} from 'vue'

const turmas = ref([])

const fetchTurmas = async () => {
  try {

    const response = await $fetch('https://southamerica-east1-matriculas242.cloudfunctions.net/ListarTurmas')
    turmas.value = response

  } catch (err) {
    console.error('Erro ao buscar turmas:', err)
    error.value = err.message || 'Erro desconhecido'
  }
}

//Para testar o GET de turmas, descomentar isso aqui:
onMounted(() =>{
    fetchTurmas()
})
</script>

<style scoped>
 main{
  width: 80%;
  margin: 0 auto;
 }
</style>