import './SearchPage.css';
import Search from "../components/Search";
import Results from "../components/Results";
import {useEffect, useState} from "react";
import superagent from "superagent";
import {useNavigate, useSearchParams} from "react-router-dom";

function SearchPage() {
    const navigate = useNavigate();
    const [results, setResults] = useState([]);
    const [searchParams, setSearchParams] = useSearchParams();

    useEffect(() => {
        superagent
            .get('http://localhost:3333/search')
            .query({query: searchParams.get('query')})
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
    }, [searchParams]);

    return (
        <div className="SearchPage">
            <div className="SearchLine">
                <h1
                    style={{
                        fontSize: '40px',
                        cursor: 'pointer',
                    }}
                    onClick={() => navigate('/')}
                >
                    SS
                </h1>
                <Search
                    onSearch={async query => {
                        const newSearchParams = new URLSearchParams({query});
                        setSearchParams(newSearchParams);
                    }}
                    value={searchParams.get('query')}
                />
            </div>
            <Results items={results}/>
        </div>
    );
}

export default SearchPage;