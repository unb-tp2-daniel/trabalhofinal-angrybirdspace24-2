<script setup lang="ts">

defineProps<{
  turma: {
    codigoTurma: string
    docente: string
    horario: string
    local: string
    capacidade: number
  }
}>()
</script>

<template>
    <tr>
        <td class="turma">{{ turma.codigoTurma }}</td>
        <td class="turma">{{ turma.docente }}</td>
        <td class="turma">{{ turma.horario }}</td>
        <td class="turma">{{ turma.local }}</td>
        <td class="turma">{{ turma.capacidade }}</td>
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