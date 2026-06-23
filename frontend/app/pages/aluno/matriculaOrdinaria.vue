<!-- pages/matricula-ordinaria.vue -->
<template>
  <div>
    <Header />
    <Menu :items="menuItems" />
    
    <main class="container">
      <h1 class="titulo-pagina">Matrícula Ordinária</h1>
      <div class="principal">
        <TabelaTurmasAbertas />
        <!-- Tem que ter o componente de grade de horarios, materias optativas e obrigatorias -->
      </div>


      <ToastResultado
        :visivel="toastVisivel"
        :mensagem="toastMensagem"
        :tipo="toastTipo"
      />
    </main>
    
    <Footer />
  </div>
</template>

<script setup>
import { ref} from 'vue'

useHead({ title: 'Matrícula Ordinária - UnB' })

/* const turmasObrigatorias = ref([])
const turmasOptativas = ref([]) */
const modalConfirmacaoAberto = ref(false) //da pra usar aqui o msm modal eu acho
const toastVisivel = ref(false)
const toastMensagem = ref('')
const toastTipo = ref('')

async function enviarMatricula() {
  modalConfirmacaoAberto.value = false
  
  try {
    /* for (const ? of ?.value) {
      await $fetch('...', {

      })
    } */
    
    mostrarToast('success', 'Matrícula realizada com sucesso!')
    disciplinasSelecionadas.value = []
  } catch (error) {
    console.error("Erro na matrícula:", error)
    const mensagemErro = error.data || 'Erro interno ao processar matrícula.'
    mostrarToast('error', mensagemErro)
  }
}

function mostrarToast(tipo, mensagem) {
  toastTipo.value = tipo
  toastMensagem.value = mensagem
  toastVisivel.value = true

  setTimeout(() => {
    toastVisivel.value = false
  }, 3000)
}

const menuItems = [
  {
    label: 'Matrícula',
    children: [
      { label: 'Matrícula Ordinária', to: '/matricula-ordinaria' },
      { label: 'Trancamento de Matrícula' },
      { label: 'Histórico de Matrículas' },
    ]
  },
  {
    label: 'Disciplinas',
    children: [
      { label: 'Buscar Disciplinas' },
      { label: 'Grade Curricular' },
      { label: 'Disciplinas Optativas' },
      { label: 'Equivalências' },
    ]
  },
  {
    label: 'Horários',
    children: [
      { label: 'Minha Grade de Horários' },
      { label: 'Choques de Horário' },
    ]
  },
  {
    label: 'Situação Acadêmica',
    children: [
      { label: 'Coeficiente de Rendimento' },
      { label: 'Pendências Acadêmicas' },
      { label: 'Prazo de Conclusão' },
    ]
  },
  { label: 'Ajuda' },
]
</script>

<style scoped>
main {
  width: 95%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 1rem;
}

.titulo-pagina {
  font-size: 1.5rem;
  font-weight: 600;
  color: #1a365d;
  margin-bottom: 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 2px solid #3182ce;
}

.principal {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
  margin-bottom: 2rem;
}

@media (min-width: 768px) {
  main {
    width: 85%;
  }
  
  .principal {
    grid-template-columns: 1.5fr 1fr;
  }
}

@media (min-width: 1024px) {
  main {
    width: 75%;
  }
}
</style>
