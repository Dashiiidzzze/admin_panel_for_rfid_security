// import axios from 'axios';

// // Получение данных
// export const getData = async (url) => {
//     try {
//         const response = await axios.get(url);
//         return response.data;
//     } catch (error) {
//         handleApiError(error);
//         throw error;
//     }
// };

// // DELETE запрос
// export const deleteData = async (url) => {
//     try {
//         await axios.delete(url);
//     } catch (error) {
//         handleApiError(error);
//         throw error;
//     }
// };

// //POST-запрос
// export const postData = async (url, data) => {
//     try {
//         const response = await axios.post(url, data, {
//             headers: { "Content-Type": "application/json" }
//         });
//         return response.data;
//     } catch (error) {
//         handleApiError(error);
//         throw error;
//     }
// };

// // Обновление данных
// export const updateData = async (url, data) => {
//     try {
//         const response = await axios.put(url, data, {
//             headers: { "Content-Type": "application/json" }
//         });
//         return response.data;
//     } catch (error) {
//         handleApiError(error);
//         throw error;
//     }
// };

// // Функция обработки ошибок API
// const handleApiError = (error) => {
//     if (error.response) {
//         if (error.response.status === 400) {
//             alert("Ошибка 400: Неверный запрос");
//         } else if (error.response.status === 404) {
//             alert("Ошибка 404: Данные не найдены");
//         } else if (error.response.status === 500) {
//             alert("Ошибка 500: Внутренняя ошибка сервера");
//         } else {
//             alert(`Ошибка ${error.response.status}: ${error.response.statusText}`);
//         }
//     } else if (error.request) {
//         alert("Ошибка сети: сервер недоступен");
//     } else {
//         alert("Произошла неизвестная ошибка");
//     }
// };


import axios from 'axios';

// Создаём настроенный экземпляр axios
const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL, // Базовый URL
  headers: {
    'Content-Type': 'application/json',
  },
});

// Добавляем interceptor для автоматической вставки токена
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`; // Добавляем токен
  }
  return config;
});

// Функция обработки ошибок API
//const handleApiError = (error) => {
    // if (error.response) {
    //     if (error.response.status === 400) {
    //         alert("Ошибка 400: Неверный запрос");
    //     } else if (error.response.status === 401) {
    //         alert("Ошибка 401: Неверный логин или пароль");
    //     } else if (error.response.status === 404) {
    //         alert("Ошибка 404: Данные не найдены");
    //     } else if (error.response.status === 500) {
    //         alert("Ошибка 500: Внутренняя ошибка сервера");
    //     } else {
    //         alert(`Ошибка ${error.response.status}: ${error.response.statusText}`);
    //     }
    // } else if (error.request) {
    //     alert("Ошибка сети: сервер недоступен");
    // } else {
    //     alert("Произошла неизвестная ошибка");
    // }
const getErrorMessage = (error) => {
    // Если есть ответ от сервера
    if (error.response) {
        const { status, data } = error.response;
        
        // Пытаемся получить сообщение из тела ответа
        const serverMessage = data?.error || data?.message || data?.detail;
        
        switch (status) {
        case 400:
            return serverMessage || "Неверные данные в запросе";
        case 401:
            return serverMessage || "Требуется авторизация";
        case 403:
            return serverMessage || "Доступ запрещён";
        case 404:
            return serverMessage || "Запрашиваемые данные не найдены";
        case 422:
            return serverMessage || "Ошибка валидации данных";
        case 500:
            return serverMessage || "Внутренняя ошибка сервера";
        default:
            return serverMessage || `Ошибка сервера (${status})`;
        }
    }
    
    // Если запрос был сделан, но ответ не получен
    if (error.request) {
        return "Сервер недоступен. Проверьте интернет-соединение";
    }
    
    // Другие ошибки (например, ошибки в коде)
    return error.message || "Произошла неизвестная ошибка";
};

// Функция обработки ошибок API
const handleApiError = (error) => {
    const message = getErrorMessage(error);
    
    alert(message);
};

// GET запрос
export const getData = async (url) => {
    try {
        const response = await api.get(url);
        return response.data;
    } catch (error) {
        handleApiError(error);
        throw error;
    }
};

// POST запрос
export const postData = async (url, data) => {
    try {
        const response = await api.post(url, data);
        return response.data;
    } catch (error) {
        handleApiError(error);
        throw error;
    }
};

// PUT запрос
export const updateData = async (url, data) => {
    try {
        const response = await api.put(url, data);
        return response.data;
    } catch (error) {
        handleApiError(error);
        throw error;
    }
};

// DELETE запрос
export const deleteData = async (url) => {
    try {
        await api.delete(url);
    } catch (error) {
        handleApiError(error);
        throw error;
    }
};