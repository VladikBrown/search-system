import './Header.css';

import Search from "./Search";
import {useNavigate, useSearchParams} from "react-router-dom";

function Header(props) {
    const navigate = useNavigate();
    const [searchParams, setSearchParams] = useSearchParams();
    return (
        <div className="Header">
            <h1 className="HeaderLogo" onClick={() => navigate('/')}>
                SS
            </h1>
            <Search
                onSearch={async query => {
                    const newSearchParams = new URLSearchParams({query});
                    setSearchParams(newSearchParams);
                }}
                value={searchParams.get('query') || ''}
            />
        </div>
    );
}

export default Header;