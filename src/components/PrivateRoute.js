import React from 'react';
import { Navigate } from 'react-router-dom';

const PrivateRoute = ({ children }) => {
    const isAuth = localStorage.getItem('auth') === 'true'; // Проверяем, залогинен ли пользователь
    return isAuth ? children : <Navigate to="/login" />;
};

export default PrivateRoute;
