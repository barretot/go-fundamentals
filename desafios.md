### **Exercícios: Fundamentos GO**

#### **1. Pacotes**
Crie um programa em Go que tenha dois pacotes: 
- Um pacote chamado `mathutils` que contenha uma função `Soma` para somar dois números.
- Outro pacote `main` que utilize a função `Soma` para calcular a soma de dois números informados pelo usuário.

---

#### **2. Nomes Públicos e Privados**
No pacote `mathutils`, implemente duas funções:
- Uma pública `Multiplicar(a, b int) int` que retorna a multiplicação de dois números.
- Uma privada `subtrair(a, b int) int` que realiza a subtração de dois números.
  
Teste apenas a função pública no pacote `main`.

---

#### **3. Funções**
Crie uma função chamada `Calcular` que receba como parâmetro dois números e uma string (`"soma"`, `"subtracao"`, `"multiplicacao"`, `"divisao"`) para realizar a operação correspondente.

---

#### **4. Variáveis**
Implemente um programa que:
- Declare e inicialize três variáveis (`nome`, `idade`, `saldo`).
- Imprima esses valores no console formatados como: `"Nome: João, Idade: 30, Saldo: 250.75"`

---

#### **5. Arrays**
Crie um array de 5 inteiros. 
- Preencha-o com valores de 1 a 5.
- Imprima cada valor do array em uma nova linha.

---

#### **6. Loops**
Implemente um loop que:
- Conte de 1 até 10.
- Verifique se o número é par ou ímpar e imprima uma mensagem correspondente.

---

#### **7. Condicionais (IFs)**
Crie um programa que:
- Solicite uma nota entre 0 e 10 ao usuário.
- Use uma estrutura `if-else` para exibir mensagens como "Reprovado" (nota < 5), "Recuperação" (nota entre 5 e 6.9) e "Aprovado" (nota >= 7).

---

#### **8. Defer**
Crie um programa que:
- Imprima "Início do programa".
- Use `defer` para imprimir "Fim do programa".
- Antes de finalizar, solicite o nome do usuário e o imprima.

---

#### **9. Ponteiros**
Crie um programa que:
- Declare uma variável `numero` com valor 10.
- Crie um ponteiro que aponte para essa variável.
- Modifique o valor da variável `numero` para 20 utilizando o ponteiro.
- Imprima o valor de `numero` antes e depois da alteração.

---