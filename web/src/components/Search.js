import './Search.css'
import React, {useEffect, useState} from "react";

export default function Search(props) {
    const {onSearch, value} = props;
    const [text, setText] = useState('');
    useEffect(() => value !== '' && setText(value), [value]);
    const GoButton = React.createRef();

    return (
        <div style={{display: 'flex'}}>
            <input
                type="text"
                className="Search-input"
                onChange={e => setText(e.target.value)}
                value={text}
                onKeyUp={(e) => {
                    if(e.keyCode === 13){
                        GoButton.current.click();
                    }
                }}
            />
            <span className="GO-button" onClick={() => onSearch(text)} ref={GoButton}>
                GO
            </span>
        </div>
    )
}
