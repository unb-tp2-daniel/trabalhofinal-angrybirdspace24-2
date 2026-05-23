<template>
  <div class="filtro_container">
    <div class="titulo">BUSCAR TURMAS ABERTAS</div>
    <div class="formulario">
      <div class="campo">
        <input type="checkbox" class="caixinha"/>
        <label>Código do Componente: </label>
        <input v-model="filtro.codigo" type="text"/>
      </div>
      
      <div class="campo">
        <input type="checkbox" class="caixinha"/>
        <label>Nome do Componente: </label>
        <input v-model="filtro.nome" type="text" class="quadradinho"/>
      </div>

      <div class="campo">
        <input type="checkbox" class="caixinha"/>
        <label>Horário: </label>
        <input v-model="filtro.horario" type="text"/>
      </div>

      <div class="campo">
        <input type="checkbox" class="caixinha"/>
        <label>Nome do Docente: </label>
        <input v-model="filtro.docente" type="text" class="quadradinho"/>
      </div>

      <div class="campo">
        <input type="checkbox" class="caixinha"/>
        <label>Unidade Responsável: </label>
        <select v-model="filtro.unidade">
          <option>FACULDADE DE TECNOLOGIA</option>
          <option>FACULDADE DE SAÚDE</option>
          <option>FACULDADE DE AGRONOMIA</option>
          <option>INSTITUTO DE COMPUTAÇÃO</option>
        </select>
      </div>
      
      <button @click="buscarTurmas">Buscar</button>

    </div>
  </div> 
</template>

<script setup lang="ts">
  const filtro = reactive({
    codigo: '',
    nome: '',
    horario: '',
    docente: '',
    unidade: ''
  })

  function buscarTurmas(){
    console.log(filtro)
  }
</script>

<style scoped>

  .titulo{
    background-color: #5b96da;
    color: white;
    font-weight: bold;
    padding: 3px;
    font-size: 17px;
    width: 30% fit-content;
    white-space: nowrap;
    border-radius: 5px;
    margin-bottom: 10px;
  }

  .filtro_container{
    border-radius: 10px 10px 0px 0px;
    background-color: rgb(236, 242, 245);
    padding: 10px;
    border: 1px solid #3f6ea5;
    /* margin: 0 auto;
    max-width: 50%; */
    margin-top: 10px;
    overflow: hidden;
    
  }

  button{
    background-color: #5b96da;
    text-align: center;
    width: 90px;
    padding: 0.5%;
    color: white;
    border: none;
    margin-top: 8px;
    margin-left: 5px;
    cursor: pointer;
    border-radius: 3px;
  }

  input,
  select{
    height: 18px;
    border: 1px solid #bdbdbd;
    font-size: 12px;
    padding: 2px 4px;
  }

  .grande{
    width: 400px;
  }

  .campo{
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 9px;
  }

  .formulario{
    padding: 3px;
  }

  label{
    width: 180px;
    font-size: 15px;
    color: #333;
  }

  button:hover{
    background-color: rgb(65, 102, 202);
  }
</style>