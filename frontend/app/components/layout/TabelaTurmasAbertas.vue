<script setup lang="ts">

  defineProps<{
  turmasAbertas: any[]
}>()

</script>

<template>
  <div class="turmas_container">
    <h2 class="titulo"></h2>
    
    <table>
      <thead>
        <tr>
          <th>Turma</th>
          <th>Docente</th>
          <th>Horário</th>
          <th>Local</th>
          <th>Capacidade</th>
        </tr>
      </thead>

      <tbody>
        <LinhaTurma v-for="turma in turmasAbertas"
        :key="turma.codigoTurma" 
        :turma="turma"
        />
      </tbody>
    </table>
  </div>
</template>

<style scoped>
  .turmas_container{
    margin-top: 20px;
    background-color: #fff;
  }

  .titulo{
    
    color: #2d2d2d;
    font-size: 20px;
    margin-bottom: 10px;
  }

  table{
    width: 100%;
    border-collapse: collapse;
  }

  thead{
    background-color: #dfe7f3;
  }

  th{
  color: #222;
  text-align: left;
  font-size: 18px;
  padding: 10px;
}
</style>