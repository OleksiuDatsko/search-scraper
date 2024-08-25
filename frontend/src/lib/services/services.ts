import type { Link, SearchResult } from "$lib/utils/types";

const BACKEND_URL = 'http://127.0.0.1:8080';

export async function fetchSearchResults(search_query: string, depth: number): Promise<SearchResult> {
    try {
        const response = await fetch(`${BACKEND_URL}/search?q=${search_query}&d=${depth}`);
        if (response.ok) {
            const data: SearchResult = await response.json();
            data.query = search_query;
            data.bot_detected = response.status === 226;
            return data;
        } else {
            console.log('Error:', response.status);
            throw new Error('Failed to fetch search results');
        }
    } catch (error) {
        console.error(error);
        throw error;
    }
}

export async function fetchLinksList(list_type: string): Promise<Link[]> {
    try {
        const response = await fetch(`${BACKEND_URL}/${list_type}`);
        if (response.ok) {
            const data: Link[] = await response.json();
            return data;
        } else {
            console.log('Error:', response.status);
            throw new Error('Failed to fetch search results');
        }
    } catch (error) {
        console.error(error);
        throw error;
    }
}

export async function updateLink(link: Link, list_type: string) {
    try {
        await fetch(`${BACKEND_URL}/${list_type}/${link.id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ domain: link.domain, url: link.url, filterType: link.filter_type })
        })
    } catch (error) {
        console.error(error);
        throw error;
    }
}

export async function deleteLink(link: Link, list_type: string) {
    try {
        await fetch(`${BACKEND_URL}/${list_type}/${link.id}`, {
            method: 'DELETE'
        })
    } catch (error) {
        console.error(error);
        throw error;
    }
}

export async function addLink(url: string, list_type: string): Promise<Link> {
    try {
        const response = await fetch(`${BACKEND_URL}/${list_type}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ url })
        })
        if (response.ok) {
            const data: Link = await response.json();
            return data;
        } else {
            console.log('Error:', response.status);
            throw new Error('Failed to fetch search results');
        }
    } catch (error) {
        console.error(error);
        throw error;
    }
}