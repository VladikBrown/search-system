import './Metrics.css';

import Header from "./Header";
import {useSearchParams} from "react-router-dom";

function Metrics(props) {
    const {metrics} = props;
    return (
        <div>
            {metrics}
        </div>
    )
}

export default Metrics;