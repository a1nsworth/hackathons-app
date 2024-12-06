import {HackathonCard} from "components/hackathon_card";
import style from "./Main.module.css";
import axios from "axios";
import {useEffect, useState} from "react";
import {HackathonApiUrl} from "config";

interface Hackathon {
    id: string;
    name: string;
    description: string;
    dateBegin: string;
    dateEnd: string;
}

export default function Main() {
    const [hackathons, setHackathons] = useState<Hackathon[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchHackathons = async () => {
            try {
                const response = await axios.get(HackathonApiUrl.GetAll);
                setHackathons(response.data);
                setLoading(false);
            } catch (error) {
                setError("Ошибка загрузки хакатонов");
                setLoading(false);
            }
        };
        fetchHackathons();
    }, []);

    if (loading) {
        return <div>Загрузка...</div>;
    }

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div className={style.container}>
            {hackathons.length === 0 ? (
                <div>Нет хакатонов для отображения</div>
            ) : (
                hackathons.map((hackathon, index) => (
                    <HackathonCard
                        key={index}
                        dateBegin={new Date(2023, 4, 1)}
                        dateEnd={new Date(2023, 4, 10)}
                        name={hackathon.name}
                        discription={hackathon.description}
                    />
                ))
            )}
        </div>
    );
}
