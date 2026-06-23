<template>
    <div>
        <Header />
        <Menu :items="menuItems" />
        <div class="container">
          <ProfileSidebarReal :aluno="aluno" :email="email" :curso="curso" :ch="ch"/>
          <main class="content">
            <TabelaMatriculas :materias="materias"/>
          </main>
        </div>
        <Footer/>
    </div>
</template>

<script setup>
  import { ref, watch, onMounted } from 'vue'
  import { useAuth } from '~/composables/useAuth'

  const { user, matriculaUsuario } = useAuth()

  useHead({ title: 'Aluno - UnB' })

  const aluno = ref(null)
  const curso = ref(null)
  const ch = ref(null)
  const email = ref(null)
  const materias = ref([])

  const buscarDadosDoAluno = async (uid) => {
    try {
      const res_aluno = await $fetch(`https://southamerica-east1-matriculas242.cloudfunctions.net/GetAlunoPorId?id=${uid}`)
      aluno.value = res_aluno
      console.log("Aluno encontrado:", aluno.value)

      const res_curso = await $fetch(`https://southamerica-east1-matriculas242.cloudfunctions.net/GetCursoPorId?id=${aluno.value.cursoId}`)
      curso.value = res_curso
      console.log("Curso encontrado:", curso.value)

      const res_ch = await $fetch(`https://southamerica-east1-matriculas242.cloudfunctions.net/CalcularCH?idAluno=${uid}`)
      ch.value = res_ch
      console.log("Carga horária encontrada:", ch.value)

    } catch (error) {
      console.error("Erro ao encontrar aluno:", error)
    }
  }

  const carregarGrade = async (uid) => {
    try {
      const res = await $fetch(`https://southamerica-east1-matriculas242.cloudfunctions.net/ListarTurmasMatriculadas?id=Unb_${uid}`)
      
      console.log(res)

      materias.value = res || []
      console.log("Grade horária real carregada:", materias.value)
    }
    
    catch (error) {
      console.error("Erro ao carregar a grade do aluno:", error)
    }
  }

  // o watch monitora o estado do usuário, assim que o firebase injetar o usuário no cliente, ele dispara a busca automaticamente
  watch(user, (novoUsuario) => {
    if (novoUsuario) {
      email.value = novoUsuario.email
      buscarDadosDoAluno(novoUsuario.uid)
      carregarGrade(novoUsuario.uid)
    }
  }, { immediate: true })

  

  onMounted(() => {
    // se o usuário já veio do localstorage no momento em que a página montou, força a busca direto
    if (user.value) {
      email.value = user.value.email
      buscarDadosDoAluno(user.value.uid)
      carregarGrade(user.value.uid)
    }
  })

const menuItems = [
  {
    label: 'Ensino',
    children: [
      { label: 'Realizar Matrícula Extraordinária', to: '/aluno/matricula' },
      { label: 'Realizar Matrícula', to: '/aluno/matriculaOrdinaria' },
      { label: 'Trancamento de Matrícula' },
      { label: 'Trancamento Geral de Matrícula' },
      { label: 'Grade Curricular' },
      { label: 'Coeficiente de Rendimento' },
    ]
  },
  {
    label: 'Financeiro',
    children: [
      { label: 'Extrado do Restaurante' },
      { label: 'Histórico de Pagamentos' },
      { label: 'Isenções e Descontos' },
    ]
  },
  {
    label: 'Documentos',
    children: [
      { label: 'Declaração de Matrícula' },
      { label: 'Atestado de Frequência' },
      { label: 'Histórico Acadêmico' },
      { label: 'Diploma Digital' },
    ]
  },
  {
    label: 'Estágio',
    children: [
      { label: 'Cadastrar Estágio' },
      { label: 'Acompanhar Estágio' },
      { label: 'Encerrar Estágio' },
    ]
  },
  {
    label: 'Bolsas',
    children: [
      { label: 'Bolsas Disponíveis' },
      { label: 'Minhas Bolsas' },
      { label: 'Renovação de Bolsa' },
    ]
  },
  {
    label: 'Outros',
    children: [
      { label: 'Ouvidoria' },
      { label: 'Calendário Acadêmico' },
      { label: 'Contato com Coordenação' },
    ]
  },
]

</script>

<style scoped>
  .container{
    display: flex;
  }

  .content{
    flex: 1;
    padding: 10px;
    background-color: #f5f5f0;
  }
</style>