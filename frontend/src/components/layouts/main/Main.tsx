import { HackathonCard } from "components/hackathon_card";
import style from "./Main.module.css";

export default function Main() {
  return (
    <div className={style.container}>
      {[...Array(10)].map((_, index) => (
        <HackathonCard
          dateBegin={new Date(2023, 4, 1)}
          dateEnd={new Date(2023, 4, 10)}
          name="Киберспортив"
          discription="Для любителей игр и технологий"
        />
      ))}
    </div>
  );
}
