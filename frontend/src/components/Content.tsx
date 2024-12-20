import {Routes, Route} from "react-router-dom";
import {Main, About, Login, Register} from "./layouts";
import style from "./Content.module.css";
import PrivateRoute from "./services/auth/PrivateRoute";

export default function Content() {
    return (
        <div className={style.content}>
            <Routes>
                <Route path="/" element={<About/>}/>
                <Route path="about" element={<About/>}/>
                <Route path="register" element={<Register/>}/>
                <Route path="login" element={<Login/>}/>
                <Route path="hackathons" element={<PrivateRoute redirectTo={"/login"}><Main/></PrivateRoute>}/>
            </Routes>
        </div>
    );
}
