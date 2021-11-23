import './pages/WelcomePage.css';
import Search from "./components/Search";
import {
    BrowserRouter as Router,
    Route,
    Routes,
} from "react-router-dom";
import WelcomePage from "./pages/WelcomePage";
import SearchPage from "./pages/SearchPage";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/">
                    <Route path="/search" element={<SearchPage/>} />
                    <Route index element={<WelcomePage/>} />
                </Route>
            </Routes>
        </Router>
    );
}

export default App;
