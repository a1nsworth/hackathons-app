import { Link } from "react-router-dom";
import logo from "assets/logo/logo-blue.332bc987.svg";
import styles from "./Header.module.css";

const menuItems = [
  { to: "/", label: "Главная" },
  { to: "/hackathons", label: "Хакатоны" },
  { to: "/about", label: "О нас" },
  { to: "/contacts", label: "Контакты" },
  { to: "/login", label: "Вход" },
  { to: "/register", label: "Регистрация" },
];

export default function Header() {
  return (
    <header className={styles.header}>
      <nav className={styles.nav}>
        <div className={styles.logo_container}>
          <Link className={styles.Link} to="/">
            <img src={logo} alt="БГТУ" />
          </Link>
        </div>
        <ul>
          {menuItems.map(({ to, label }, index) => (
            <li key={index}>
              <Link key={index} className={styles.Link} to={to}>
                {label}
              </Link>
            </li>
          ))}
        </ul>
      </nav>
    </header>
  );
}
