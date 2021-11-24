import './SearchPage.css';
import Search from "../components/Search";
import Results from "../components/Results";
import {useEffect, useState} from "react";
import superagent from "superagent";
import {useSearchParams} from "react-router-dom";
import Header from "../components/Header";

function SearchPage() {
    const [results, setResults] = useState([]);
    const [searchParams] = useSearchParams();

    const query = searchParams.get('query');

    useEffect(() => {
        superagent
            .get('http://localhost:3333/search')
            .query({query})
            .then(res => {
                setResults(
                    res.body.Docs
                        .filter(r => r.similarityRate !== 0)
                        .sort((r1, r2) => r2.similarityRate - r1.similarityRate)
                        .map(r => r.doc)
                );
            })
            .catch(err => {
                alert(err)
            });
    }, [query]);

    return (
        <div className="SearchPage">
            <Header/>
            <Results items={results}/>
        </div>
    );
}

export default SearchPage;