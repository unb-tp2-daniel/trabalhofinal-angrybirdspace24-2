<script setup lang="ts">
defineProps<{
  turma: {
    codigoTurma: string
    materiaId: string //??
    nomeMateria: string
    semestre: string
    vagasTotais: number
    vagasOcupadas: number
    ativo: boolean
    local: string
    horario: string

    vagasExclusivas: Record<string, number>

    professorId: string
    professorNome: string

    prioridades: Record<string, string>
  }
}>()

const emite = defineEmits([ 
  'selecionar'
])

/* function capacidadeClass(vagas: number): string {
  if (vagas === 0) return 'cap-zero'
  if (vagas <= 10) return 'cap-full'
  if (vagas <= 25) return 'cap-med'
  return 'cap-ok'
}

function capacidadeLabel(vagas: number): string {
  if (vagas === 0) return 'Sem vagas'
  return `${vagas} vagas`
} */
</script>

<template>
    <tr class="linha_turma" @click="emite('selecionar', turma)">
        <td class="code_badge">{{ turma.codigoTurma }}</td> <!-- para ser o codigo da materia, basta trocar por turma.materiaId -->
        <td class="docente">{{ turma.professorNome }}</td>
        <td class="horario_badge">{{ turma.horario }}</td>
        <td class="local">{{ turma.local }}</td>
        <td class="cap_badge">{{ turma.vagasTotais }}</td>
        <td class="cap_badge">{{ turma.vagasOcupadas }}</td>
    </tr>
</template>

<style scoped>

.linha_turma{
  border-bottom: 1px solid #e2e8f0;
  transition: background 0.12s;
}

.linha_turma :last-child{
  border-bottom: none;
}

.linha_turma:hover td{
  background: #eef4fd;
}

td{
  padding: 13px 16px;
  border-bottom: 1px solid #dcdcdc;
  background-color: #f8fbff;
}

.docente { 
  color: #333; 
}

.local   { 
  font-size: 12px; color: #666;
}

.code_badge {
  font-family: 'DM Mono', 'Courier New', monospace;
  font-size: 12px;
  font-weight: 500;
  background: #e6f1fb;
  color: #0c447c;
  border-radius: 5px;
}

.horario_badge {
  font-family: 'DM Mono', 'Courier New', monospace;
  font-size: 12px;
  background: #f4f4f4;
  color: #555;
  border-radius: 5px;
}

.cap_badge {
  font-size: 12px;
  font-weight: 500;
  border-radius: 20px;
}

/* .cap-ok   { background: #eaf3de; color: #3b6d11; }
.cap-med  { background: #faeeda; color: #854f0b; }
.cap-full { background: #fcebeb; color: #a32d2d; }
.cap-zero { background: #f4f4f4; color: #888; } */

.linha_turma:nth-child(even) td{
  background-color: #f8fafc;
}

</style>