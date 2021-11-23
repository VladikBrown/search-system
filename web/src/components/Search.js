import './Search.css'
import {useState} from "react";

export default function Search(props) {
    const {onSearch} = props;
    const [text, setText] = useState('');
    return (
        <div style={{display: 'flex'}}>
            <input
                type="text"
                className="Search-input"
                onChange={e => setText(e.target.value)}/>
            <span className="GO-button" onClick={() => onSearch(text)}>
                GO
            </span>
        </div>
    )
}
