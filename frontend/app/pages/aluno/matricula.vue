<template>
    <div>
    <Header />
    <Menu :items="menuItems" />
    <main class="container">
      <TabelaBuscas  @resultados="turmas = $event" /> 
      <TabelaTurmasAbertas :turmasAbertas="turmas" @selecionar="abrirModal"/>

      <ModalConfirmacao 
      :visivel="modalAberto"
      :turma="turmaSelecionada"
      @fechar="modalAberto = false"
      @confirmar="confirmarMatricula"/>

      <ToastResultado
      :visivel="toastVisivel"
      :mensagem="toastMensagem"
      :tipo="toastTipo" />

    </main>
    <Footer />
    </div>
</template>

<script setup>

useHead({ title: 'Matrícula Extraordinária - UnB' })

import { ref } from 'vue'
const turmas = ref([])
const toastVisivel = ref(false)
const toastMensagem = ref('')
const toastTipo = ref('')
const modalAberto = ref(false)
const turmaSelecionada = ref(null)

function abrirModal(turma) {
  turmaSelecionada.value = turma
  modalAberto.value = true
}

function mostrarToast(tipo, mensagem){
  toastTipo.value = tipo
  toastMensagem.value = mensagem
  toastVisivel.value = true

  setTimeout(() => {
    toastVisivel.value = false
  }, 3000)
}


  async function confirmarMatricula(turma) {
    modalAberto.value = false

    /* ELE RECEBE O ID COMO "Unb_MATRICULA" */
    const alunoId = "Unb_" + "211010057" // por enquanto, depois resgatar pelo usuario autenticado (n sei como faz)

    try {
      const res = await $fetch('https://southamerica-east1-matriculas242.cloudfunctions.net/MatricularExtraordinaria', {
        method: 'POST',
        body: {
          AlunoId: alunoId,
          TurmaId: turmaSelecionada.value.codigoTurma,
          Status: false,
          DataSolicitacao: null,
          Prioridades: null,
          Semestre: "20261", // POR ENQUANTO TAMBÉM
          PrioridadeNota: 0
        }
      })

      mostrarToast('success', 'Matrícula extraordinária confirmada com sucesso!')

      turma.vagasOcupadas++ // apenas altera localmente pra ficar bontio
    }

    catch (error) {
      console.error("Erro na matrícula:", error)
      
      const mensagemErro = error.data || 'Erro interno ao processar matrícula.'
      console.log(mensagemErro)
      mostrarToast(
        'error',
        mensagemErro
      )
    }
  }

  const menuItems = [
  {
    label: 'Matrícula',
    children: [
      { label: 'Trancamento de Matrícula'},
      { label: 'Histórico de Matrículas' },
    ]
  },
  {
    label: 'Disciplinas',
    children: [
      { label: 'Buscar Disciplinas'},
      { label: 'Grade Curricular'},
      { label: 'Disciplinas Optativas' },
      { label: 'Equivalências' },
    ]
  },
  {
    label: 'Horários',
    children: [
      { label: 'Minha Grade de Horários'},
      { label: 'Choques de Horário'},
    ]
  },
  {
    label: 'Situação Acadêmica',
    children: [
      { label: 'Coeficiente de Rendimento'},
      { label: 'Pendências Acadêmicas'},
      { label: 'Prazo de Conclusão'},
    ]
  },
  { label: 'Ajuda'},
]
</script>

<style scoped>
 main{
  width: 70%;
  margin: 0 auto;
 }
</style>