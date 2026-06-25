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
          <td class="local">{{ materias.local }}</td>
          <td><span class="horario">{{ materias.horario }}</span></td>
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
  text-align: left;
  font-size: 15px;
  letter-spacing: 0.06em;
  padding: 11px 16px;
  font-weight: bold;
  width: 100%;
  border-collapse: collapse;
  border: 1px solid #e2e8f0;
  border-radius: 8px 8px 0px 0px;
  overflow: hidden;
}

  .turmas_container{
    margin-top: 20px;

  }

  .titulo{
    color: #2d2d2d;
    font-size: 18px;
    margin-bottom: 10px;
  }

  .atualizacao{
    color: #555;
    font-size: 12px;
    width: 100%;
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
  }

  .turmas_table thead{
    background-color: #1a3a7a;
  }

  .turmas_table thead th{
    color: rgba(255, 255, 255, 0.85);
    font-size: 12px;
    padding: 11px 16px;
    text-transform: uppercase;
  }

  .turmas_table tbody tr{
  border-bottom: 1px solid #e8eef5;
  transition: background 0.12s;
  }

  td{
    padding: 14px 12px;
    border-bottom: 1px solid #dcdcdc;
    background-color: #fff;
    font-size: 13px;
    color: #333;
  }

  .materia{
    color: #003399;
    font-weight: bold;
    font-size: 12px;
    letter-spacing: 0.03em;
  }

  .local{
    font-size: 13px;
    color: #555;
  }

  .horario{
    font-family: 'DM Mono', 'Courier New', monospace;
    font-size: 12px;
    background: #f0f5fa;
    color: #1a3a7a;
    padding: 3px 8px;
    border-radius: 5px;
    border: 1px solid #c5d8f0;
  }

  tr:nth-child(even) td{
    background-color: #d7e7ff;
  }
</style>