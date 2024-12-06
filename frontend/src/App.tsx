import {Header, Footer} from "components/layouts";
import Content from "components/Content";
import {AuthProvider} from "./components/services/auth/AuthProvied";

export default function App() {
    return (
        <AuthProvider>
            <div className="App">
                <Header/>
                <Content/>
                <Footer/>
            </div>
        </AuthProvider>
    );
}
