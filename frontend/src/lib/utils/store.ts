import { writable } from 'svelte/store';
import type { SearchResult } from './types';

export const blockedIds = writable([] as number[]);
export const findedIds = writable([] as number[]);

export const searchedResults = writable({
    query: '',
    scraped_link: [],
    result_rating: 0
} as SearchResult);
