<template>
    <Header />
    <div class="page">
        <div class="top-section">
            <div>
                <h1>Turmas</h1>
                <p>Gerencie todas as turmas cadastradas no sistema</p>
            </div>

            <button class="new-class-btn">
                + Nova Turma
            </button>
        </div>

        <div class="filters">
            <input
                type="text"
                placeholder="Buscar turma, professor ou código..."
            />

            <select>
                <option>Todas as matérias</option>
            </select>
        </div>

        <div v-if="loading">
            Carregando turmas...
        </div>

        <div v-else-if="error">
            {{ error }}
        </div>

        <div v-else>
            <section class="subject-section">
                <div class="subject-header">
                    <h2>Todas as Turmas</h2>
                    <span>{{ turmas.length }} turmas</span>
                </div>

                <table class="classes-table">
                    <thead>
                        <tr>
                            <th>Matéria</th>
                            <th>Código</th>
                            <th>Professor</th>
                            <th>Horário</th>
                            <th>Local</th>
                            <th>Semestre</th>
                            <th>Vagas</th>
                            <th>Status</th>
                        </tr>
                    </thead>

                    <tbody>
                        <tr
                            v-for="turma in turmas"
                            :key="turma.codigoTurma"
                        >
                            <td>{{ turma.nomeMateria }}</td>

                            <td>{{ turma.codigoTurma }}</td>

                            <td>
                                {{
                                    turma.professorNome &&
                                    turma.professorNome.trim()
                                        ? turma.professorNome
                                        : 'Não informado'
                                }}
                            </td>

                            <td>{{ turma.horario }}</td>

                            <td>{{ turma.local }}</td>

                            <td>{{ turma.semestre }}</td>

                            <td>
                                {{ turma.vagasOcupadas }}/{{ turma.vagasTotais }}
                            </td>

                            <td>
                                <span
                                    class="status"
                                    :class="{
                                        active: turma.vagasOcupadas < turma.vagasTotais,
                                        full: turma.vagasOcupadas >= turma.vagasTotais
                                    }"
                                >
                                    {{
                                        turma.vagasOcupadas >= turma.vagasTotais
                                            ? 'Lotada'
                                            : 'Disponível'
                                    }}
                                </span>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </section>
        </div>

        <button
            class="new-class-btn"
            @click="index"
        >
            Voltar
        </button>
    </div>

    <Footer />
</template>

<script setup>
    import { ref, computed, onMounted } from 'vue'

    const turmas = ref([])
    const loading = ref(false)
    const error = ref(null)

    async function carregarTurmas() {
        loading.value = true

        try {
            const response = await $fetch(
                'https://southamerica-east1-matriculas242.cloudfunctions.net/ListarTurmas'
            )

            turmas.value = response
            console.log(response)
            console.log(JSON.stringify(response, null, 2))
        } catch (err) {
            console.error(err)
            error.value = err.message
        } finally {
            loading.value = false
        }
    }

    const turmasDisponiveis = computed(() =>
        turmas.value
    )

    function index() {
        navigateTo('/coordenador')
    }

    onMounted(() => {
        carregarTurmas()
    })
</script>

<style scoped>

.page{
    padding:30px;
    background:#f1f5f9;
    min-height:100vh;
}

.top-section{
    display:flex;
    justify-content:space-between;
    align-items:center;
    margin-bottom:30px;
}

.top-section h1{
    font-size:36px;
    margin-bottom:6px;
}

.top-section p{
    color:#1a5276;
}

.new-class-btn{
    background:#1a5276;
    color:white;
    border:none;
    padding:14px 20px;
    border-radius:10px;
    cursor:pointer;
    font-weight:700;
    transition:0.2s;
}

.new-class-btn:hover{
    opacity:0.9;
}

.filters{
    display:flex;
    gap:16px;
    margin-bottom:30px;
}

.filters input,
.filters select{
    padding:12px;
    border-radius:10px;
    border:1px solid #cbd5e1;
    background:white;
}

.filters input{
    width:320px;
}

.subject-section{
    background:white;
    border-radius:16px;
    padding:24px;
    margin-bottom:30px;
    box-shadow:0 2px 8px rgba(0,0,0,0.06);
}

.subject-header{
    display:flex;
    justify-content:space-between;
    align-items:center;
    margin-bottom:20px;
}

.subject-header h2{
    color:#1a5276;
}

.subject-header span{
    color:#1a5276;
    font-size:14px;
}

.classes-table{
    width:100%;
    border-collapse:collapse;
}

.classes-table th{
    background:#f8fafc;
    padding:14px;
    text-align:left;
    color:#334155;
    font-size:14px;
}

.classes-table td{
    padding:16px 14px;
    border-top:1px solid #e2e8f0;
}

.classes-table tr:hover{
    background:#f8fafc;
}

.actions{
    display:flex;
    gap:8px;
}

.actions button{
    border:none;
    padding:8px 12px;
    border-radius:8px;
    cursor:pointer;
    font-weight:600;
}

.edit{
    background:#dbeafe;
    color:#1d4ed8;
}

.students{
    background:#dcfce7;
    color:#166534;
}

.delete{
    background:#fee2e2;
    color:#b91c1c;
}

.status{
    padding:6px 10px;
    border-radius:999px;
    font-size:12px;
    font-weight:700;
}

.active{
    background:#dcfce7;
    color:#166534;
}

.warning{
    background:#fef3c7;
    color:#92400e;
}

.full{
    background:#fee2e2;
    color:#991b1b;
}

</style>