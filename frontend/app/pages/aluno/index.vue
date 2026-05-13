<template>
    <div>
        <Header />
        <Menu :items="menuItems" />
        <div class="container">
          <ProfileSidebar />
          <main class="content">
            <TabelaMaterias :materias="materias"/>
          </main>
        </div>
        <Footer/>
    </div>
</template>

<script setup>
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