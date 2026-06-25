<script setup lang="ts">
  import { ref, onMounted, watch, computed } from 'vue'
  const props = defineProps<{
    turmasAbertas: any[]
  }>()

  const turmas = ref<any[]>([])
  const error = ref<string | null>(null)
  const loading = ref(false)
  
  const turmasDisponiveis = computed(() => 
    turmas.value.filter(t => t.vagasOcupadas < t.vagasTotais)
  )

  const itensPorPagina = 10
  const paginaAtual = ref(1)

  const totalPaginas = computed(() =>
    Math.ceil(turmasDisponiveis.value.length / itensPorPagina)
  )

  const turmasPaginadas = computed(() => {
    const inicio = (paginaAtual.value - 1) * itensPorPagina
    const fim = inicio + itensPorPagina

    return turmasDisponiveis.value.slice(inicio, fim)
  })

  watch(() => props.turmasAbertas, (novas) => {
    turmas.value = novas
    paginaAtual.value = 1
  })

  const fetchTurmas = async () => {
    loading.value = true
    try {
      const response = await $fetch<any[]>('https://southamerica-east1-matriculas242.cloudfunctions.net/ListarTurmas')
      turmas.value = response
    } catch (err: any) {
      console.error('Erro ao buscar turmas:', err)
      error.value = err.message || 'Erro desconhecido'
    } finally {
      loading.value = false
    }
  }

  onMounted(() => {
    fetchTurmas()
  }) 

  const emit = defineEmits([
    'selecionar',
    'detalhes'
  ])

  watch(() => props.turmasAbertas, (novas) => {
    turmas.value = novas
  })
</script>

<template>
  <div class="turmas_container">
    <!-- Os textos ficam fora da área de scroll e permanecem fixos -->
    <div class="results-meta">
      <span class="results-label">Resultados</span>
      <span class="results-count">{{ turmasDisponiveis.length }} turmas encontradas</span>
    </div>

    <div v-if="loading" class="state-msg">
      Carregando turmas...
    </div>

    <div v-else-if="error" class="state-msg state-error">
      Erro ao carregar: {{ error }}
    </div>

    <div v-else-if="turmasDisponiveis.length === 0" class="state-msg">
      Nenhuma turma encontrada.
    </div>
    
    <!-- Nova div envelopando apenas a tabela para isolar o scroll horizontal -->
    <div v-else class="table_responsive_wrapper">
      <table class="turmas_table">
        <thead>
          <tr>
            <th>Turma</th>
            <th>Docente</th>
            <th>Horário</th>
            <th>Local</th>
            <th>Capacidade</th>
            <th>Ocupadas</th>
            <th>Ações</th>
          </tr>
        </thead>

        <tbody>
          <LinhaTurma v-for="turma in turmasPaginadas"
            :key="turma.codigoTurma" 
            :turma="turma"
            @selecionar="emit('selecionar', $event)"
            @detalhes="emit('detalhes', $event)"
          />
        </tbody>
      </table>
    </div>
      <div v-if="totalPaginas >= 1" class="pagination">
      <button @click="paginaAtual--" :disabled="paginaAtual === 1"> ← </button>
      <span>
        Página {{ paginaAtual }} de {{ totalPaginas }}
      </span>
      <button @click="paginaAtual++" :disabled="paginaAtual === totalPaginas"> → </button>
    </div>
  </div>
</template>

<style scoped>
  .turmas_container {
    margin-top: 24px;
    background-color: #fff;
    width: 100%;
    /* Removemos o overflow-x daqui para os textos não scrollarem */
  }

  /* Esta nova classe controla o scroll somente da tabela */
  .table_responsive_wrapper {
    width: 100%;
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
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

  .turmas_table {
    width: 100%;
    border-collapse: collapse;
    border: 1px solid #e2e8f0;
    overflow: hidden;
    margin: auto;
    border-radius: 10px 10px 0px 0px;
    white-space: nowrap;
  }

  .turmas_table thead tr {
    background: #1a3a7a;
  }

  .turmas_table thead th {
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

  .pagination{
  color: #555;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
  padding: 16px;
}

.pagination button {
  border: none;
  background: #1a3a7a;
  color: white;
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
}

.pagination button:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}
</style>