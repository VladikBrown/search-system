import './SearchPage.css';
import Search from "../components/Search";
import Results from "../components/Results";
import {useEffect, useState} from "react";
import superagent from "superagent";
import {useSearchParams} from "react-router-dom";
import Header from "../components/Header";
import Metrics from "../components/Metrics";
import {BsFillBarChartLineFill} from 'react-icons/bs';

function SearchPage() {
    const [results, setResults] = useState([]);
    const [metrics, setMetrics] = useState({});
    const [showMetrics, setShowMetrics] = useState(false);
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
                setMetrics(res.body.metricsAggregator)
            })
            .catch(err => {
                alert(err)
            });
    }, [query]);

    return (
        <div className="SearchPage">
            <Header/>
            {showMetrics ?
                <Metrics metrics={JSON.stringify(metrics)}/>
                :
                <Results items={results}/>
            }
            <MetricsButton onClick={() => setShowMetrics(!showMetrics)}/>
        </div>
    );
}

function MetricsButton(props) {
    return (
        <div className="MetricsButton" {...props}>
            <BsFillBarChartLineFill/>
        </div>
    )
}

export default SearchPage;