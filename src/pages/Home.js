// import React, { useState, useEffect } from 'react';
// import { Link } from 'react-router-dom';
// import axios from 'axios';

// const Home = () => {
//     // загрузка данных в state для обновления без перезагрузки
//     const [data, setData] = useState([]);

//     useEffect(() => {
//         async function loadData() {
//             try {
//                 const response = await axios.get("http://localhost:5000/items", {
//                     headers: { "Content-Type": "application/json" } // Указание формата данных
//                 });
//                 setData(response.data); // Обновляем состояние
//                 console.log("Данные загружены:", response.data);
//             } catch (error) {
//                 console.error("Ошибка запроса:", error);
//             }
//         }

//         loadData();
//     }, []);

//     // Функция для удаления товара
//     const deleteItem = async (id) => {
//         try {
//             await axios.delete(`http://localhost:5000/items/${id}`, {
//                 headers: { "Content-Type": "application/json" }
//             });
//             console.log(`Сотрудник ${id} удален`);
//             setData(prevData => prevData.filter(item => item.id !== id)); // Обновляем state
//         } catch (error) {
//             console.error("Ошибка удаления:", error);
//         }
//     };

//     return (
//         <div>
//             <h1>Список сотрудников</h1>
//             <Link className="addition" to="/add">Добавить сотрудника</Link>
//             <ul>
//                 {data.map(item => (
//                     <li key={item.id}>
//                         <Link to={`/detail/${item.id}`}>{item.name}</Link>
//                         <button className="delete" onClick={() => deleteItem(item.id)} style={{marginLeft: "10px" }}>
//                             Удалить
//                         </button>
//                     </li>
//                 ))}
//             </ul>
            
//         </div>
//     );
// };

// export default Home;

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
