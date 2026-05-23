<template>
  <div class="turmas_container">
    <h2 class="titulo">2026.1</h2>
    <div class="atualizacao"><strong>Últimas Atualizações</strong>
      <p>24/05/2026 - 
        <span class="link">TECNICAS DE PROGRAMAÇÃO 2</span>
      </p>
      <p>
        12/04/2026 - 
        <span class="link">ÁLGEBRA 1</span>
      </p>

      <p >Novo tópico: Oportunidade de Estágio</p>
      <p>Aviso: Cancelamento da aula de amanhã 13/05/2026</p>
    </div>
    
    <table class="turmas_table">
      <thead>
        <tr>
          <th>Componente Curricular</th>
          <th>Local</th>
          <th>Horário</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="materias in materias" :key="materias.nome">
          <td class="materia">{{ materias.nome }}</td>
          <td>{{ materias.local }}</td>
          <td>{{ materias.horario }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
  defineProps({
    materias: Array
  })
</script>

<style scoped>

.turmas_table thead th{
  color: rgba(24, 24, 24, 0.85);
  text-align: left;
  font-size: 15px;
  letter-spacing: 0.06em;
  padding: 11px 16px;
  font-weight: 500;
  /* text-transform: uppercase; */
  font-weight: bold;
}

  .turmas container{
    margin-top: 20px;
    background-color: #fff;
  }

  .titulo{
    color: #2d2d2d;
    font-size: 20px;
    margin-bottom: 10px;
  }

  .atualizacao{
    background-color: #eef2f8;
    border: 2px solid #d0d8e4;
    padding: 12px;
    margin-bottom: 20px;
    border-radius: 10px 0px 10px 0px;
  }

  .atualizacao > p{
    color: #2c2c2c;
  }

  .link{
    text-decoration: none;
    font-weight: bold;
    cursor: pointer;
    color: #022a7a;
  }

  .link:hover{
    text-decoration: underline;
  }

  table{
    width: 100%;
    border-collapse: collapse;
    color: #333;
  }

  thead{
    background-color: #dfe7f3;
  }

  th{
    color: #333;
    text-align: left;
    font-size: 18px;
    padding: 10px;
  }

  td{
    padding: 14px 12px;
    border-bottom: 1px solid #dcdcdc;
    background-color: #f8fbff;
  }

  .materia{
    color: #003399;
    font-weight: bold;
    font-size: 12px;
  }

  tr:nth-child(even) td{
    background-color: #d7e7ff;
  }
</style>