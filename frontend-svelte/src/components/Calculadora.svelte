<script lang="ts">
  import { GOLANG_SERVER } from "../utils/constants";
  import {
    history,
    currentResultStored,
    currentExpressionStored,
  } from "./store";

  let currentExpression = $currentExpressionStored;
  let currentResult = $currentResultStored;

  currentExpressionStored.subscribe((value) => {
    currentExpression = value;
  });

  currentResultStored.subscribe((value) => {
    currentResult = value;
  });

  // Enviar expresión al backend de Go. Refrescar historial
  const calculate = async () => {
    if (currentExpressionStored) {
      const response = await fetch(`${GOLANG_SERVER}/operation`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          expression: currentExpression,
        }),
      });
      const res = await response.json();
      currentResult = (Math.round(res.result * 100) / 100).toString();
      const newEntry: Operation = {
        result: res.result,
        time: res.time,
        operation: res.operation,
      };
      history.update((oldHistory: Operation[]) => [...oldHistory, newEntry]);
    }
  };

  // Manejar botones ingresados. Si es "=" calcular. Si no, añadir a la operación. Utilizar ultimo resultado después de clickear en otro botón
  const handleButtonClick = (value: string) => {
    if (value === "=") {
      calculate();
    } else if (currentResult) {
      currentExpression = currentResult += value;
      currentResult = "";
    } else {
      // Verificar si el último caracter es un operador
      const lastChar = currentExpression.slice(-1);
      const operators = ["+", "-", "*", "/"];
      if (!operators.includes(lastChar) || !operators.includes(value)) {
        currentExpression += value;
      }
      currentResult = "";
    }
  };

  // Limpiar todo
  const clear = () => {
    currentExpression = "";
    currentResult = "";
  };

  // Limpiar última entrada
  const clearEntry = () => {
    currentExpression = currentExpression.slice(0, -1);
  };
  // Invertir operadores - y +
  const changeSign = () => {
    currentExpression = currentExpression.replace(/[+-]/g, (match) =>
      match === "+" ? "-" : "+",
    );
  };
</script>

<div
  class="calculator bg-white border border-gray-300 p-6 rounded-lg min-w-[350px] max-h-[450px]"
>
  <div
    class="display bg-black text-white p-2 rounded mb-4 border-[5px] border-gray-500 border-double"
  >
    <p class="text-right min-h-6 text-white">
      {currentExpression}
      {currentResult && " = " + currentResult}
    </p>
  </div>

  <div class="grid grid-cols-4 gap-3">
    <button
      class="bg-red-500 text-white p-4 rounded hover:bg-red-600 transition"
      on:click={clearEntry}>CE</button
    >

    <button
      class="bg-red-500 text-white p-4 rounded hover:bg-red-600 transition"
      on:click={clear}>C</button
    >

    <button
      class="bg-orange-500 text-white p-4 rounded hover:bg-orange-600 transition"
      on:click={changeSign}>+/-</button
    >

    <button
      class="bg-orange-500 text-white p-4 rounded hover:bg-orange-600 transition"
      on:click={() => handleButtonClick("%")}>%</button
    >

    {#each ["7", "8", "9"] as buttonValue}
      <button
        class="bg-blue-500 text-white p-4 rounded hover:bg-blue-600 transition"
        on:click={() => handleButtonClick(buttonValue)}>{buttonValue}</button
      >
    {/each}

    <button
      class="bg-orange-500 text-white p-4 rounded hover:bg-orange-600 transition"
      on:click={() => handleButtonClick("/")}>/</button
    >

    {#each ["4", "5", "6"] as buttonValue}
      <button
        class="bg-blue-500 text-white p-4 rounded hover:bg-blue-600 transition"
        on:click={() => handleButtonClick(buttonValue)}>{buttonValue}</button
      >
    {/each}

    <button
      class="bg-orange-500 text-white p-4 rounded hover:bg-orange-600 transition"
      on:click={() => handleButtonClick("*")}>*</button
    >

    {#each ["1", "2", "3"] as buttonValue}
      <button
        class="bg-blue-500 text-white p-4 rounded hover:bg-blue-600 transition"
        on:click={() => handleButtonClick(buttonValue)}>{buttonValue}</button
      >
    {/each}

    <button
      class="bg-orange-500 text-white p-4 rounded hover:bg-orange-600 transition"
      on:click={() => handleButtonClick("-")}>-</button
    >

    {#each ["0", "."] as buttonValue}
      <button
        class="bg-blue-500 text-white p-4 rounded hover:bg-blue-600 transition"
        on:click={() => handleButtonClick(buttonValue)}>{buttonValue}</button
      >
    {/each}

    <button
      class="bg-green-500 text-white p-4 rounded hover:bg-green-600 transition"
      on:click={() => handleButtonClick("=")}>=</button
    >

    <button
      class="bg-orange-500 text-white p-4 rounded hover:bg-orange-600 transition"
      on:click={() => handleButtonClick("+")}>+</button
    >
  </div>
</div>
