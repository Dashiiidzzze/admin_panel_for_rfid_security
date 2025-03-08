import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { getData } from '../components/CheckErrors';

const Logs = () => {
    // загрузка данных в state для обновления без перезагрузки
    const [data, setData] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        async function loadData() {
            try {
                const employees = await getData("http://localhost:5001/logs");
                setData(employees); // Обновляем состояние
                setError(null);
            } catch (error) {
                setError(error.message);
                console.error("Ошибка запроса:", error);
            } finally {
                setLoading(false);
            }
        }

        loadData();
    }, []);

    if (loading) return <p>Загрузка...</p>;

    return (
        <div>
            <h1>Список проходов</h1>
            {error && <p className="error">{error}</p>}
            <ul>
                {data.map(item => (
                    <li key={item.id}>
                        <div className='info'>{item.time} {item.zone}</div>
                        <Link to={`/detail/${item.id}`}>{item.name}</Link> 
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default Logs;