import './WelcomePage.css';
import Search from "../components/Search";
import {useNavigate} from "react-router-dom";

function WelcomePage() {
    const navigate = useNavigate();
    return (
        <div className="WelcomePage">
            <header className="WelcomePage-header">
                <h1 className="WelcomePage-title">Search system</h1>
                <Search onSearch={query => {
                    navigate(`/search?query=${query}`);
                }}/>
            </header>
        </div>
    );
}

export default WelcomePage;