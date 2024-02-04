<script lang="ts">
  import svelteLogo from "./assets/svelte.svg";
  import viteLogo from "/vite.svg";
  import Calculadora from "./components/Calculadora.svelte"
  import Historial from "./components/Historial.svelte"
  import "./app.css";
</script>

<main class="flex space-x-4">
  <Calculadora />
  <Historial />
</main>


