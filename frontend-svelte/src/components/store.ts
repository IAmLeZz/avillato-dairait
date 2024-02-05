import { writable } from 'svelte/store';
import type { Operation } from '../utils/types'

export const history = writable<Operation[]>([]);