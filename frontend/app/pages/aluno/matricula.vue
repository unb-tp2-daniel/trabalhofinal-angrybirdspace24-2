<template>
  <div>
    <Header />
    <Menu :items="menuItems" />
    <main class="page-container">
      <TabelaBuscas @resultados="turmas = $event" /> 
      
      <TabelaTurmasAbertas 
        :turmasAbertas="turmas" 
        @selecionar="abrirModal" 
        @detalhes="abrirDetalhesMateria"
      />

      <ModalConfirmacao 
        :visivel="modalAberto"
        :turma="turmaSelecionada"
        @fechar="modalAberto = false"
        @confirmar="confirmarMatricula"
      />

      <ModalDetalhesMateria
        :visivel="modalDetalhesAberto"
        :materia="materiaSelecionada"
        @fechar="modalDetalhesAberto = false"
      />

      <ToastResultado
        :visivel="toastVisivel"
        :mensagem="toastMensagem"
        :tipo="toastTipo" />
    </main>
    <Footer />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

useHead({ title: 'Matrícula Extraordinária - UnB' })

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

function mostrarToast(tipo, message) {
  toastTipo.value = tipo
  toastMensagem.value = message
  toastVisivel.value = true

  setTimeout(() => {
    toastVisivel.value = false
  }, 3000)
}

async function confirmarMatricula() {
  modalAberto.value = false

  // Como o useAuth expõe um ref reativo, acessamos o valor via .value
  const alunoId = matriculaUsuario.value 
  const aluno = await $fetch(`https://southamerica-east1-matriculas242.cloudfunctions.net/GetAlunoPorId?id=${alunoId}`)

  try {
    const res = await $fetch('https://southamerica-east1-matriculas242.cloudfunctions.net/MatricularExtraordinaria', {
      method: 'POST',
      body: {
        AlunoId: alunoId,
        TurmaId: turmaSelecionada.value.codigoTurma,
        MateriaId: turmaSelecionada.value.codigoTurma.split("_", 2).join("_"),
        CursoId: "",
        Status: false,
        DataSolicitacao: null,
        Prioridades: null,
        Semestre: "20261",
        PrioridadeNota: 0
      }
    })

    mostrarToast('success', 'Matrícula extraordinária confirmada com sucesso!')

    // CORREÇÃO DE REATIVIDADE: Altera a propriedade diretamente na referência selecionada
    if (turmaSelecionada.value) {
      turmaSelecionada.value.vagasOcupadas++
    }
  } catch (error) {
    console.error("Erro na matrícula:", error)
    const mensagemErro = error.data || 'Erro interno ao processar matrícula.'
    mostrarToast('error', mensagemErro)
  }
}

const menuItems = [
  {
    label: 'Matrícula',
    children: [
      { label: 'Página Inicial', to: '/aluno'},
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
    const codigoProcurado = turma.materiaId || turma.codigoTurma.split("_", 2).join("_")
    const materias = await $fetch(
      `https://southamerica-east1-matriculas242.cloudfunctions.net/ProcurarMateria?id=${codigoProcurado}`
    )

    if (materias) {
      materiaSelecionada.value = materias
      modalDetalhesAberto.value = true
    } else {
      alert("Esta matéria não foi encontrada no catálogo do Decanato.")
    }
  } catch(err) {
    console.error("Erro ao buscar matérias:", err)
    materiaSelecionada.value = null
  }
}
</script>

<style scoped>
main {
  width: 95%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 1rem;
}

@media (min-width: 768px) {
  main {
    width: 70%;
  }
}
</style>