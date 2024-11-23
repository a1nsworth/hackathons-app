import style from "./Footer.module.css";
import ContributionCard, { ContributionCardProps } from "./ContributionCard";
import gosuslugi from "assets/logo/gosyslygi.7cb00466.webp";
import book from "assets/logo/book.29b708c5.webp";
import gerd from "assets/logo/gerd.e480a8e0.webp";
import globus from "assets/logo/globus.6d3d8123.webp";
import logo from "assets/logo/logo.b53dc80c.webp";
import pero from "assets/logo/pero.0043dec6.webp";

const contributors: ContributionCardProps[] = [
  {
    img: book,
    text: "Министерство Науки и Высшего Образования Российской Федерации",
  },
  {
    img: gerd,
    text: "Высшая аттестационная комиссия (ВАК)",
  },
  {
    img: globus,
    text: "Российское образование — федеральный портал www.edu.ru",
  },
  {
    img: pero,
    text: "Федеральный центр информационно-образовательных ресурсов",
  },
  {
    img: logo,
    text: "Информационная витрина достижений и разработок БГТУ им. В.Г. Шухова",
  },
  {
    img: "/images/map.png",
    text: "Интерактивная карта антитеррористической деятельности",
  },
  {
    img: "/images/terrorism.png",
    text: "Экстремизм, пропаганда терроризма в интернете? Сообщите",
  },
  {
    img: "/images/science-against-terror.png",
    text: "Наука и образование против террора",
  },
  {
    img: gosuslugi,
    text: "Госуслуги — это выгодно и удобно",
  },
];

function NavigationTable() {
  return (
    <div className={style.navigation}>
      <table>
        <thead>
          <tr>
            <th colSpan={2} style={{ textAlign: "left" }}>
              Об университете
            </th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Телефонный справочник</td>
            <td>Обучение</td>
          </tr>

          <tr>
            <td>Контактные данные</td>
            <td>Структура университета</td>
          </tr>
          <tr>
            <td>Абитуриенту</td>
            <td>Об университете</td>
          </tr>
          <tr>
            <td>Обратная связь</td>
            <td>Статистика</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}

function SocialTable() {
  return (
    <div className={style.social}>
      <table>
        <thead>
          <th colSpan={2} style={{ textAlign: "left" }}>
            БГТУ в сети
          </th>
        </thead>
        <tbody>
          <tr>
            <td rowSpan={2}>vk</td>
            <td>
              <b>+ 7(4722) 55 05 04</b> - телефон доверия
            </td>
          </tr>
          <tr>
            <td>
              <b>8 800 222 55 71 </b>— телефон горячей линии МинОбрНауки
            </td>
          </tr>
          <tr>
            <td rowSpan={2}>
              <b>Карта сайта</b>
            </td>
            <td>© 2000 — 2024 БГТУ им. В.Г. Шухова</td>
          </tr>
          <tr>
            <td
              colSpan={2}
              style={{ wordWrap: "break-word", maxWidth: "200px" }}
            >
              Политика БГТУ им. В.Г. Шухова в отношении обработки персональных
              данных
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}
export default function Footer() {
  return (
    <footer className={style.footer}>
      <div className={style.footer_container}>
        <div className={style.contribution_container}>
          {contributors.map(({ img, text }, index) => (
            <ContributionCard
              key={index}
              img={img}
              text={text}
              style={{
                flexGrow: 0,
                flexBasis: "170px",
                display: "flex",
                justifyContent: "left",
                alignItems: "center",
              }}
            />
          ))}
        </div>
        <div className={style.info_container}>
          <NavigationTable />
          <SocialTable />
        </div>
      </div>
    </footer>
  );
}
