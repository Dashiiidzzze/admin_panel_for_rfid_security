import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { getData, deleteData } from '../components/CheckErrors'; // Импорт API-функций

const Home = () => {
    const [data, setData] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const loadEmployees = async () => {
            try {
                const employees = await getData("http://localhost:5000/items");
                setData(employees);
                setError(null);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        loadEmployees();
    }, []);

    const deleteItem = async (id) => {
        try {
            await deleteData(`http://localhost:5000/items/${id}`);
            setData(prevData => prevData.filter(item => item.id !== id));
        } catch (err) {
            setError(err.message);
        }
    };

    if (loading) return <p>Загрузка...</p>;

    return (
        <div>
            <h1>Список сотрудников</h1>
            <Link className="addition" to="/add">Добавить сотрудника</Link>

            {error && <p className="error">{error}</p>}

            <ul>
                {data.map(item => (
                    <li key={item.id}>
                        <Link to={`/detail/${item.id}`}>{item.name}</Link>
                        <button className="delete" onClick={() => deleteItem(item.id)} style={{ marginLeft: "10px" }}>
                            Удалить
                        </button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Home;
