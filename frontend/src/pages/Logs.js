import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { getData } from '../components/CheckErrors';

const Logs = () => {
    const [data, setData] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        async function loadData() {
            try {
                const response = await getData(`${process.env.REACT_APP_API_URL}/logs`);
                const logsData = response.logspasses || [];
                // Преобразуем строки времени в объекты Date
                const formattedData = response.logspasses.map(item => ({
                    ...item,
                    time: new Date(item.time)
                }));
                setData(formattedData);
                setError(null);
            } catch (error) {
                setError(error.message);
                console.error("Ошибка запроса:", error);
                setData([]);
            } finally {
                setLoading(false);
            }
        }

        loadData();
    }, []);

    if (loading) {
        return (
            <div className="spinner-container">
                <div className="spinner"></div>
            </div>
        );
    }

    return (
        <div>
            <h1>Список проходов</h1>
            {error && <p className="error">{error}</p>}
            <ul>
                {data.length > 0 ? (
                    data.map((item, index) => (
                        <li key={index}>
                            <div className='info'>
                                {/* Отображаем дату в читаемом формате */}
                                {item.time.toLocaleString('ru-RU', {
                                    day: '2-digit',
                                    month: '2-digit',
                                    year: 'numeric',
                                    hour: '2-digit',
                                    minute: '2-digit',
                                    second: '2-digit',
                                })} — {item.zone}
                            </div>
                            <Link to={`/detail/${index}`}>{item.name}</Link> 
                        </li>
                    ))
                ) : (
                    <p>Логов проходов не найдено.</p>
                )}
            </ul>
        </div>
    );
};

export default Logs;
