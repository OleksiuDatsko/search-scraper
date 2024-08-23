export interface ScrapedLink {
    id: number;
    title: string;
    link: string;
    domain: string;
    snipped: string;
}
export interface SearchResult {
    scraped_link: ScrapedLink[];
    query: string;
    result_rating: number;
}

export interface Link {
    id: number;
    url: string;
    domain: string;
    filter_type: string;
}