import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../components/AuthContext';
import { postData } from '../components/CheckErrors'; // путь поправьте под свой проект

const Login = () => {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState(''); // для показа ошибок
    const navigate = useNavigate();
    const { login: authenticate } = useAuth();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setErrorMessage(''); // очищаем предыдущую ошибку

        try {
            const data = await postData(`${process.env.REACT_APP_API_URL}/login`, { login, password });

            // Сохраняем JWT-токен в localStorage
            localStorage.setItem('token', data.token);

            // Обновляем состояние авторизации
            authenticate();
            await new Promise(resolve => setTimeout(resolve, 0)); // микро-задержка

            // Переход на главную страницу
            navigate('/');
        } catch (error) {
            // Показываем ошибку под формой
            if (error.response?.status === 401) {
                setErrorMessage('Неверный логин или пароль');
            } else {
                setErrorMessage('Ошибка сервера: попробуйте позже');
            }
        }
    };

    return (
        <div>
            <h1>Войдите в учетную запись</h1>
            <form className='login' onSubmit={handleSubmit}>
                <label>
                    Логин: <input type="text" value={login} onChange={(e) => setLogin(e.target.value)} required />
                </label>
                <br />
                <label>
                    Пароль: <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
                </label>
                <br />
                <button className='loginsub' type="submit">Войти</button>
            </form>
            {errorMessage && <p style={{ color: 'red',  textAlign: 'center'}}>{errorMessage}</p>}
        </div>
    );
}

export default Login;
