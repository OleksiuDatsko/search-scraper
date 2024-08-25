import { writable } from 'svelte/store';
import type { Link, SearchResult } from './types';

export const blockedIds = writable([] as number[]);
export const findedIds = writable([] as number[]);

export const searchedResults = writable({
    query: '',
    scraped_link: [],
    result_rating: 0
} as SearchResult);


export const tabSet = writable(0);

export const links = writable([] as Link[])