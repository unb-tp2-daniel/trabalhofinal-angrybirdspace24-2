import http from 'k6/http';
import { check, sleep } from 'k6';

/* criado com auxilio do gemini para os testes de gargalo */

export const options = {
    // cenario de pico
    scenarios: {
        gargalo_atomico: {
        executor: 'constant-vus',
        vus: 100,           // 60 alunos concorrendo
        duration: '8s',
    },
  },
};

export default function () {
  const url = 'https://southamerica-east1-matriculas242.cloudfunctions.net/MatricularExtraordinaria';
  
  const raAleatorio = "262" + Math.floor(100000 + Math.random() * 900000);

  const payload = JSON.stringify({
    AlunoId: raAleatorio,
    TurmaId: "HIST01_0001_01", // turma aleatoria
    Semestre: "20261"
  });

  const params = { headers: { 'Content-Type': 'application/json' } };
  const res = http.post(url, payload, params);

  check(res, {
    'vaga conquistada (201)': (r) => r.status === 201,
    'vaga esgotada (409)': (r) => r.status === 409,
  });

  sleep(0.01); 
}