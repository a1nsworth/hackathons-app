import style from "./HackathonCard.module.css";

interface HackathonCardProps {
  name: string;
  dateBegin: Date;
  dateEnd: Date;
  discription: string;
}

const formatDateRange = (start: Date, end: Date): string => {
  const options: Intl.DateTimeFormatOptions = {
    day: "numeric",
    month: "long",
    year: "numeric",
  };

  const startDate = start.toLocaleDateString("ru-RU", options);
  const endDate = end.toLocaleDateString("ru-RU", options);

  if (startDate === endDate) {
    return startDate;
  }

  return `${startDate} - ${endDate}`;
};

export default function HackathonCard(props: HackathonCardProps) {
  return (
    <div className={style.wrapper}>
      <div className={style.container}>
        <div className={style.name_container}>
          <h1>{props.name}</h1>
        </div>
        <div className={style.content}>
          <p>Дата:</p>
          <p>{formatDateRange(props.dateBegin, props.dateEnd)}</p>
          <hr />
          <p>Описание</p>
          <p style={{ fontSize: "16px" }}>
            Этот хакатон посвящен созданию инновационных решений для улучшения
            образовательного процесса. Команды смогут разрабатывать приложения и
            инструменты, которые помогут в обучении студентов.
          </p>
        </div>
      </div>
    </div>
  );
}
