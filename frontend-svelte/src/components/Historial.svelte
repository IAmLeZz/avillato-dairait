<script lang="ts">
  import { onMount } from "svelte";
  import axios from "axios";
  import {
    history,
    currentExpressionStored,
    currentResultStored,
  } from "./store";
  import { GOLANG_SERVER } from "../utils/constants";

  let historyData: Operation[] = [];

  history.subscribe((value) => {
    historyData = value;
  });

  onMount(async () => {
    const response = await axios.get(`${GOLANG_SERVER}/history`);
    history.set(response.data);
  });
  const handleExpressionClick = (expression: string, result: string) => {
    currentExpressionStored.set(expression);
    currentResultStored.set(result);
  };
</script>

<div
  class="history text-gray-900 bg-white border border-gray-300 p-6 rounded-lg max-h-[460px] overflow-y-scroll"
>
  <h2 class="text-xl mb-2">Historial</h2>
  <table
    class="w-full text-left bg-black text-white min-w-[350px] border-4 border-gray-600"
  >
    <thead>
      <tr class="m-auto p-5 border-b text-center">
        <th>Fecha</th>
        <th>Operación</th>
        <th>Resultado</th>
      </tr>
    </thead>
    <tbody class="mx-2 p-5 space-y-4 text-center">
      {#each historyData as item (item.time)}
        <tr class="mx-2 border-b">
          <td class="border-r">{item.time && item.time}</td>
          <td class="border-r">{item.operation && item.operation}</td>
          <td
            tabindex="0"
            on:click={() => handleExpressionClick(item.operation, item.result)}
            on:keydown={(event) => {
              if (event.key === "Enter") {
                handleExpressionClick(item.operation, item.result);
              }
            }}
            class="text-blue-400 hover:bg-gray-900 hover:cursor-pointer"
          >
            {item.result && item.result}
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
