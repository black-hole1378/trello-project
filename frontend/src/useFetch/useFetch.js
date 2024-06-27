import useSWR from 'swr';
import axios from 'axios';

const fetcher = (url, options) => axios(url, options).then((res) => res.data);

const useFetch = (url, options) => {
    const { data, error, isLoading } = useSWR(url ? [url, options] : null, fetcher);

    return { data, isLoading, error };
};

export default useFetch;
