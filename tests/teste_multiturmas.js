import http from 'k6/http';
import { check, sleep } from 'k6';

/* criado com auxilio do gemini para os testes */

export const options = {
  // simula alunos entrando aos poucos
  stages: [
    { duration: '5s', target: 40 },  // sobe para 40 usuários em 5s
    { duration: '10s', target: 40 }, // mantém 40 usuários por 10s
    { duration: '5s', target: 0 },   // desce para 0 usuários
  ],
};

const listaTurmas = [
  'CC01_0005_01',
  'CIC0099_01_20261',
  'DIR01_0001_01',
  'DIR01_0002_01',
  'CC01_0001_02'
];

export default function () {
  const url = 'https://southamerica-east1-matriculas242.cloudfunctions.net/MatricularExtraordinaria';
  
  const raAleatorio = "262" + Math.floor(100000 + Math.random() * 900000);
  
  const turmaAleatoria = listaTurmas[Math.floor(Math.random() * listaTurmas.length)];

  const payload = JSON.stringify({
    AlunoId: raAleatorio,
    TurmaId: turmaAleatoria,
    Semestre: "20261"
  });

  const params = { headers: { 'Content-Type': 'application/json' } };
  const res = http.post(url, payload, params);

  check(res, {
    'resposta ok (201 ou 409)': (r) => r.status === 201 || r.status === 409,
  });

  sleep(0.05);
}