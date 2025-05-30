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
