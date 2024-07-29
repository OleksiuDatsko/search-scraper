import { writable } from 'svelte/store';

export const blockedIds = writable([] as number[]);
export const findedIds = writable([] as number[]);