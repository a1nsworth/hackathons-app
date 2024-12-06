import {Link} from "react-router-dom";
import logo from "assets/logo/logo-blue.332bc987.svg";
import style from "./Header.module.css";
import {useAuth} from "components/services/auth/AuthProvied"; // Убедитесь, что путь к вашему AuthContext указан правильно

export default function Header() {
    const {isAuthenticated, logout} = useAuth();
    const menuItems = isAuthenticated
        ? [
            {to: "/", label: "Главная"},
            {to: "/hackathons", label: "Хакатоны"},
            {to: "/about", label: "О нас"},
            {to: "/contacts", label: "Контакты"},
            {to: "/logout", label: "Выйти"},
        ]
        : [
            {to: "/", label: "Главная"},
            {to: "/about", label: "О нас"},
            {to: "/contacts", label: "Контакты"},
            {to: "/login", label: "Вход"},
            {to: "/register", label: "Регистрация"},
        ];

    const handleLogout = () => {
        logout();
    };

    return (
        <header className={style.header}>
            <nav className={style.nav}>
                <div className={style.logo_container}>
                    <Link className={style.Link} to="/">
                        <img src={logo} alt="БГТУ"/>
                    </Link>
                </div>
                <ul>
                    {menuItems.map(({to, label}, index) => (
                        <li key={index}>
                            {to === "/logout" ? (
                                <button className={style.Link} onClick={handleLogout}>
                                    {label}
                                </button>
                            ) : (
                                <Link key={index} className={style.Link} to={to}>
                                    {label}
                                </Link>
                            )}
                        </li>
                    ))}
                </ul>
            </nav>
        </header>
    );
}
