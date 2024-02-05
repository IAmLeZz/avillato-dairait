<script>
    import { onMount } from 'svelte';
    import { history } from './store.js'
    let display = '';
    let result = '';
    let num1 = '';
    let num2 = '';
    let operator = '';
  
    const calculate = async () => {
  if (num1 && num2 && operator) {
    const response = await fetch('http://localhost:8080/operation', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        operator: operator,
        num1: Number(num1),
        num2: Number(num2)
      })
    });
    const res = await response.json();
    result =  Math.round(res.result*100)/100;
    num1 = '';
    num2 = '';
    operator = '';
    console.log(result)
    const newEntry = {
        result: res.result,
        time: res.time,
        operation: res.operation,
    }
    history.update(oldHistory => [...oldHistory, newEntry]);
  }
}
        const handleNumberClick = (number) => {
        if (!operator) {
            num1 = num1 + number;
        } else {
            num2 = num2 + number;
        }
        display = num1 + operator + num2;
    }
  

  const handleOperatorClick = (op) => {
  if (result) {
    num1 = result.toString();
    result = '';
  }
  if (op === '-' && num1 === '') {
    num1 = '-';
    display = num1;
  } else {
    operator = op;
    display = num1 + operator;
  }
}
const clear = () => {
    display = '';
    num1 = '';
    num2 = '';
    operator = '';
    result = '';
}
  
    const clearEntry = () =>{
      if (!operator) {
        num1 = '';
      } else {
        num2 = '';
      }
      display = '';
      result = '';
    }
  
    const changeSign = () => {
      if (!operator) {
        num1 = (-Number(num1)).toString();
        display = num1;
      } else {
        num2 = (-Number(num2)).toString();
        display = num2;
      }
    }

    const handleDecimalClick = () => {
  if (!operator) {
    if (!num1.includes('.')) {
      num1 += '.';
    }
  } else {
    if (!num2.includes('.')) {
      num2 += '.';
    }
  }
  display = num1 + operator + num2;
}
  </script>

  <div class="calculator bg-white border border-gray-300 p-6 rounded-lg min-w-[350px]">
    <div class="display bg-black text-white p-2 rounded mb-4">
      <p class="text-right min-h-6 text-white">{display} {(result) && (" = " + result)}</p>
    </div>
  
    <div class="grid grid-cols-4 gap-3">
      <button class="bg-orange-500 text-white p-4 rounded" on:click={clearEntry}>CE</button>
      <button class="bg-orange-500 text-white p-4 rounded" on:click={clear}>C</button>
      <button class="bg-red-500 text-white p-4 rounded" on:click={changeSign}>±</button>
      <button class="bg-red-500 text-white p-4 rounded" on:click={() => handleOperatorClick('%')}>%</button>
  
      <button class="bg-blue-500 text-white p-4 rounded" on:click={() => handleNumberClick('7')}>7</button>
      <button class="bg-blue-500 text-white p=4 rounded" on:click={() => handleNumberClick('8')}>8</button> 
      <button class="bg-blue-500 text-white p=4 rounded" on:click={() => handleNumberClick('9')}>9</button> 
      <button class="bg-red-800 text-white p=4 rounded" on:click={() => handleOperatorClick('/')}>/</button> 

      <button class="bg-blue-500 text-white p-4 rounded" on:click={() => handleNumberClick('4')}>4</button>
      <button class="bg-blue-500 text-white p=4 rounded" on:click={() => handleNumberClick('5')}>5</button> 
      <button class="bg-blue-500 text-white p=4 rounded" on:click={() => handleNumberClick('6')}>6</button> 
      <button class="bg-red-800 text-white p=4 rounded" on:click={() => handleOperatorClick('*')}>x</button> 

      <button class="bg-blue-500 text-white p-4 rounded" on:click={() => handleNumberClick('3')}>3</button>
      <button class="bg-blue-500 text-white p=4 rounded" on:click={() => handleNumberClick('2')}>2</button> 
      <button class="bg-blue-500 text-white p=4 rounded" on:click={() => handleNumberClick('1')}>1</button> 
      <button class="bg-red-800 text-white p=4 rounded" on:click={() => handleOperatorClick('-')}>-</button> 
       
      <button class="bg-blue-500 text-white p-4 rounded" on:click={() => handleNumberClick('0')}>0</button>
      <button class="bg-blue-500 text-white p=4 rounded" on:click={()=> handleDecimalClick()}>.</button> 
      <button class="bg-green-700 text-white p=4 rounded" on:click={calculate}>=</button>
      <button class="bg-red-800 text-white p=4 rounded" on:click={() => handleOperatorClick('+')}>+</button>
       
    </div>
  </div>