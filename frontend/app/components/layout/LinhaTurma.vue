<script setup lang="ts">

defineProps<{
  turma: {
    codigoTurma: string
    professorNome: string
    horario: string
    local: string
    vagasTotais: number
  }
}>()
</script>

<template>
    <tr>
        <td class="turma">{{ turma.codigoTurma }}</td>
        <td class="turma">{{ turma.professorNome }}</td>
        <td class="turma">{{ turma.horario }}</td>
        <td class="turma">{{ turma.local }}</td>
        <td class="turma">{{ turma.vagasTotais }}</td>
    </tr>
</template>

<style scoped>

td{
  padding: 14px 12px;
  border-bottom: 1px solid #dcdcdc;
  background-color: #f8fbff;
}

.turmasAbertas{
  color: #003399;
  font-weight: bold;
  font-size: 12px;
}

tr:nth-child(even) td{
  background-color: #d7e7ff;
}

</style>