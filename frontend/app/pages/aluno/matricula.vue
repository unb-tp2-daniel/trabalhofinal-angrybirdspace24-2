<template>
    <div>
    <Header />
    <Menu :items="menuItems" />
    <main class="container">
      <TabelaBuscas  @resultados="turmas = $event" /> 
      <TabelaTurmasAbertas :turmasAbertas="turmas" @selecionar="abrirModal" @detalhes="abrirDetalhesMateria"/>

      <ModalConfirmacao 
      :visivel="modalAberto"
      :turma="turmaSelecionada"
      @fechar="modalAberto = false"
      @confirmar="confirmarMatricula"/>

      <ModalDetalhesMateria
      :visivel="modalDetalhesAberto"
      :materia="materiaSelecionada"
      @fechar="modalDetalhesAberto = false"/>

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

import { ref, computed } from 'vue'
import { useAuth } from '~/composables/useAuth'

const { matriculaUsuario } = useAuth() 

const modalDetalhesAberto = ref(false)
const materiaSelecionada = ref(null)
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
    const alunoId = "Unb_" + matriculaUsuario.value // por enquanto, depois resgatar pelo usuario autenticado (n sei como faz)
    console.log(alunoId)
    try {
      const res = await $fetch('https://southamerica-east1-matriculas242.cloudfunctions.net/MatricularExtraordinaria', {
        method: 'POST',
        body: {
          AlunoId: alunoId,
          TurmaId: turmaSelecionada.value.codigoTurma,
          MateriaId: turmaSelecionada.value.codigoTurma.split("_", 2).join("_"), // pega so o id da materia
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

  async function abrirDetalhesMateria(turma) {
    try {
      const materias = await $fetch(
        'https://southamerica-east1-matriculas242.cloudfunctions.net/ListarMaterias'
      )

      const codigoProcurado = turma.materiaId || turma.codigoTurma.split("_", 2).join("_")

      const materiaEncontrada = materias.find(m => m.codigo === codigoProcurado)

      if (materiaEncontrada) {
        materiaSelecionada.value = materiaEncontrada
        modalDetalhesAberto.value = true
      } else {
        alert("Esta matéria não foi encontrada no catálogo do Decanato.")
      }
    }
    catch(err) {
      console.error("Erro ao buscar matérias:", err)
      materiaSelecionada.value = null
    }
  }
</script>

<style scoped>
 main{
  width: 70%;
  margin: 0 auto;
 }
</style>