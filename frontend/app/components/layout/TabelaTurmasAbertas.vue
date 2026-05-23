<script setup lang="ts">
  import {ref , onMounted} from 'vue'
  defineProps<{
    turmasAbertas: any[]
  }>()


const turmas = ref<any[]>([])
const error = ref<string | null>(null)
const loading = ref(false)

const fetchTurmas = async () => {
  loading.value = true
  try {

    const response = await $fetch<any[]>('https://southamerica-east1-matriculas242.cloudfunctions.net/ListarTurmas')
    turmas.value = response

  } catch (err: any) {
    console.error('Erro ao buscar turmas:', err)
    error.value = err.message || 'Erro desconhecido'
  }finally {
    loading.value = false
  }
}

//Para testar o GET de turmas, descomentar isso aqui:
onMounted(() =>{
    fetchTurmas()
}) 

</script>

<template>
  <div class="turmas_container">
    <!-- <h2 class="titulo">Turmas Encontradas</h2> -->
    <div class="results-meta">
      <span class="results-label">Resultados</span>
      <span class="results-count">{{ turmas.length }} turmas encontradas</span>
    </div>

    <div v-if="loading" class="state-msg">
      Carregando turmas...
    </div>

    <div v-else-if="error" class="state-msg state-error">
      Erro ao carregar: {{ error }}
    </div>

    <div v-else-if="turmas.length === 0" class="state-msg">
      Nenhuma turma encontrada.
    </div>
    
    <table v-else class="turmas_table">
      <thead>
        <tr>
          <th>Turma</th>
          <th>Docente</th>
          <th>Horário</th>
          <th>Local</th>
          <th>Capacidade</th>
          <th>Ocupadas</th>
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
    margin-top: 24px;
    background-color: #fff;
    width: 100%;
    overflow-x: auto;
  }

.state-msg {
  padding: 32px;
  text-align: center;
  color: #888;
  font-size: 14px;
  border: 1px solid #e5e5e5;
  border-radius: 10px;
}

.state-error {
  color: #a32d2d;
  background: #fcebeb;
}

  .titulo{
    
    color: #ffffff;
    width: 30% fit-content;
    white-space: nowrap;
    font-size: 20px;
    margin-bottom: 10px;
    border-radius: 5px;
  }

  .turmas_table{
    width: 100%;
    border-collapse: collapse;
    border: 1px solid #e2e8f0;
    /* border-radius: 10px; */
    overflow: hidden;
    margin: auto;
    border-radius: 10px 10px 0px 0px;
    white-space: nowrap;
  }

  .turmas_table thead tr{
    background: #1a3a7a;
  }

  .turmas_table thead th{
  color: rgba(255, 255, 255, 0.85);
  text-align: left;
  font-size: 15px;
  letter-spacing: 0.06em;
  padding: 11px 16px;
  font-weight: 500;
  text-transform: uppercase;
}

.results-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10px;
}

.results-label {
  font-size: 13px;
  font-weight: 500;
  color: #555;
}

.results-count {
  font-size: 12px;
  color: #888;
  background: #f4f4f4;
  padding: 3px 10px;
  border-radius: 20px;
  border: 1px solid #e5e5e5;
}
</style>