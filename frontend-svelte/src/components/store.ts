import { writable } from 'svelte/store';

export const history = writable<Operation[]>([]);
export const currentExpressionStored = writable<string>("");
export const currentResultStored = writable<string>("");
