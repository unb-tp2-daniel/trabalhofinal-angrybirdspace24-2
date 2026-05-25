<template>
    <div>
    <Header />
    <!-- <Menu/> -->
    <main class="container">
      <TabelaBuscas  @resultados="turmas = $event" />
         <!-- <div v-for="turma in turmas">
            <p>Turma: {{ turma.nomeMateria }}</p>
        </div> --> 
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


function confirmarMatricula(turma){
  const vagasRestantes = turma.vagasTotais - turma.vagasOcupadas
  if(vagasRestantes > 0){
    turma.vagasOcupadas++

    mostrarToast(
      'success',
      'Matrícula realizada com sucesso!'
    )

  } else {
    mostrarToast(
      'error',
      'Não há mais vagas disponíveis.'
    )
  }
  modalAberto.value = false
}
</script>

<style scoped>
 main{
  width: 70%;
  margin: 0 auto;
 }
</style>