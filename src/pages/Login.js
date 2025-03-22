import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../components/AuthContext';

const Login = () => {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();
    const { login: authenticate } = useAuth();

    const handleSubmit = (e) => {
        e.preventDefault();
        
        if (login === 'admin' && password === '1234') {
            authenticate(); // Обновляем глобальное состояние
            navigate('/'); 
        } else {
            alert('Неверные логин или пароль!');
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
        </div>
    );
}

export default Login;
