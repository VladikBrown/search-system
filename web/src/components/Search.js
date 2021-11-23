import './Search.css'
import {useEffect, useState} from "react";

export default function Search(props) {
    const {onSearch, value} = props;
    const [text, setText] = useState('');
    useEffect(() => value !== '' && setText(value), [value]);

    return (
        <div style={{display: 'flex'}}>
            <input
                type="text"
                className="Search-input"
                onChange={e => setText(e.target.value)}
                value={text}
            />
            <span className="GO-button" onClick={() => onSearch(text)}>
                GO
            </span>
        </div>
    )
}
