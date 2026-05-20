<template>
    <div>
    <Header />
    <main class="container">
      <TabelaBuscas/>
        <!-- <div v-for="turma in turmas">
            <p>Turma: {{ turma.nomeMateria }}</p>
        </div> -->
        <TabelaTurmasAbertas :turmasAbertas="turma"/>
    </main>
    <Footer />
    </div>
</template>

<script setup>
/* import {ref , onMounted} from 'vue'

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
}) */

const turma = ref([
  {
    codigoTurma: 'CIC0229',
    docente: 'Marcelo Mandelli',
    horario: '24T34',
    local: 'PJC 085',
    capacidade: 40
  },

  {
    codigoTurma: 'CIC0002',
    docente: 'Vinicius Ruela',
    horario: '35M12',
    local: 'PJC 012',
    capacidade: 35
  },

  {
    codigoTurma: 'MAT0026',
    docente: 'Henrique Reis',
    horario: '24M34',
    local: 'BSAS 101',
    capacidade: 60
  },

  {
    codigoTurma: 'EST0023',
    docente: 'Esqueci o nome',
    horario: '46T23',
    local: 'ICC Sul Anf 7',
    capacidade: 45
  },

  {
    codigoTurma: 'CIC0198',
    docente: 'Daniel Porto',
    horario: '24M12',
    local: 'PAT BT 20',
    capacidade: 30
  }
])
</script>

<style scoped>
 main{
  width: 70%;
  margin: 0 auto;
 }
</style>