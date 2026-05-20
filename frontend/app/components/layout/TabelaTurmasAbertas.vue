<script setup lang="ts">
  import {ref , onMounted} from 'vue'
  defineProps<{
    turmasAbertas: any[]
  }>()


const turmas = ref<any[]>([])
const error = ref<string | null>(null)

const fetchTurmas = async () => {
  try {

    const response = await $fetch<any[]>('https://southamerica-east1-matriculas242.cloudfunctions.net/ListarTurmas')
    turmas.value = response

  } catch (err: any) {
    console.error('Erro ao buscar turmas:', err)
    error.value = err.message || 'Erro desconhecido'
  }
}

//Para testar o GET de turmas, descomentar isso aqui:
onMounted(() =>{
    fetchTurmas()
}) 

</script>

<template>
  <div class="turmas_container">
    <h2 class="titulo"></h2>
    
    <table>
      <thead>
        <tr>
          <th>Turma</th>
          <th>Docente</th>
          <th>Horário</th>
          <th>Local</th>
          <th>Capacidade</th>
        </tr>
      </thead>

      <tbody>
        <LinhaTurma v-for="turma in turmas"
        :key="turma.codigoTurma" 
        :turma="turma"
        />
      </tbody>
    </table>
  </div>
</template>

<style scoped>
  .turmas_container{
    margin-top: 20px;
    background-color: #fff;
  }

  .titulo{
    
    color: #2d2d2d;
    font-size: 20px;
    margin-bottom: 10px;
  }

  table{
    width: 100%;
    border-collapse: collapse;
  }

  thead{
    background-color: #dfe7f3;
  }

  th{
  color: #222;
  text-align: left;
  font-size: 18px;
  padding: 10px;
}
</style>