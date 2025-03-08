// import React, { useState, useEffect } from 'react';
// import { useNavigate } from 'react-router-dom';

// const GlobalButtons = () => {
//     const navigate = useNavigate();
//     const [isAuthenticated, setIsAuthenticated] = useState(false);

//     // Проверяем аутентификацию при загрузке
//     useEffect(() => {
//         const authStatus = localStorage.getItem('auth');
//         setIsAuthenticated(!!authStatus); // Приводим к boolean
//     }, []);

//     // Обработчик нажатия кнопки
//     const handleButtonClick = (button) => {
//         if (button === 1) {
//             navigate('/');
//         } else if (button === 2) {
//             navigate('/logs');
//         } else if (button === 3) {
//             localStorage.removeItem('auth'); // Удаляем флаг аутентификации
//             setIsAuthenticated(false); // Обновляем состояние
//             navigate('/login')
//         }
//     };

//     // Если пользователь не авторизован, не показываем кнопки
//     if (!isAuthenticated) {
//         return null;
//     }

//     return (
//         <div className="global-buttons">
//             <button onClick={() => handleButtonClick(1)}>Список сотрудников</button>
//             <button onClick={() => handleButtonClick(2)}>Логи проходов</button>
//             <button onClick={() => handleButtonClick(3)}>Завершить сессию</button>
//         </div>
//     );
// };

// export default GlobalButtons;

import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from './AuthContext';

const GlobalButtons = () => {
    const navigate = useNavigate();
    const { isAuthenticated, logout } = useAuth();

    if (!isAuthenticated) {
        return null; // Не отображаем кнопки, если не авторизован
    }

    const handleButtonClick = (button) => {
        if (button === 1) {
            navigate('/');
        } else if (button === 2) {
            navigate('/logs');
        } else if (button === 3) {
            logout(); // Выход из аккаунта
            navigate('/login');
        }
    };

    return (
        <div className="global-buttons">
            <button onClick={() => handleButtonClick(1)}>Список сотрудников</button>
            <button onClick={() => handleButtonClick(2)}>Логи проходов</button>
            <button onClick={() => handleButtonClick(3)}>Завершить сессию</button>
        </div>
    );
};

export default GlobalButtons;
