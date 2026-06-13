<template>
    <div>
        <Header />
        <Menu :items="menuItems" />
        <div class="container">
          <ProfileSidebarReal :aluno="aluno" :email="email"/>
          <main class="content">
            <TabelaMatriculas :materias="materias"/>
          </main>
        </div>
        <Footer/>
    </div>
</template>

<script setup>

  useHead({ title: 'Aluno - UnB' })
  import { ref } from 'vue';
  import { getFirebaseAuth } from "../../../plugins/firebase.client"
  import { onMounted } from 'vue'

  const auth = getFirebaseAuth()

  const aluno = ref(null);
  const email = ref(null);
  onMounted(async () => {
    
    auth.onAuthStateChanged(async (user) => {
    if (user) {
      email.value = user.email;
      try {
        const res = await $fetch(`https://southamerica-east1-matriculas242.cloudfunctions.net/GetAlunoPorId?id=${user.uid}`)
        aluno.value = res
        console.log("Aluno encontrado:", aluno.value)
      } catch (error) {
        console.error("Erro ao encontrar aluno:", error)
      }
    }
    })
  })
  


  const materias = [
    {
    nome: 'ORGANIZAÇÃO E ARQUITETURA DE COMPUTADORES',
    local: 'PJC BT 085',
    horario: '35M34' 
  },

  {
    nome: 'PROGRAMAÇÃO COMPETITIVA',
    local: 'BSA N B1 41/13',
    horario: '24M34'
  },

  {
    nome: 'TÉCNICAS DE PROGRAMAÇÃO 2',
    local: 'BSA N AT 19/41',
    horario: '24M12'
  }
]

const menuItems = [
  {
    label: 'Ensino',
    children: [
      { label: 'Realizar Matrícula Extraordinária', to: '/aluno/matricula' },
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