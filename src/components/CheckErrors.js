import axios from 'axios';

// Получение данных
export const getData = async (url) => {
    try {
        const response = await axios.get(url);
        return response.data;
    } catch (error) {
        handleApiError(error);
        throw error;
    }
};

// DELETE запрос
export const deleteData = async (url) => {
    try {
        await axios.delete(url);
    } catch (error) {
        handleApiError(error);
        throw error;
    }
};

//POST-запрос
export const postData = async (url, data) => {
    try {
        const response = await axios.post(url, data, {
            headers: { "Content-Type": "application/json" }
        });
        return response.data;
    } catch (error) {
        handleApiError(error);
        throw error;
    }
};

// Обновление данных
export const updateData = async (url, data) => {
    try {
        const response = await axios.put(url, data, {
            headers: { "Content-Type": "application/json" }
        });
        return response.data;
    } catch (error) {
        handleApiError(error);
        throw error;
    }
};

// Функция обработки ошибок API
const handleApiError = (error) => {
    if (error.response) {
        if (error.response.status === 400) {
            alert("Ошибка 400: Неверный запрос");
        } else if (error.response.status === 404) {
            alert("Ошибка 404: Данные не найдены");
        } else if (error.response.status === 500) {
            alert("Ошибка 500: Внутренняя ошибка сервера");
        } else {
            alert(`Ошибка ${error.response.status}: ${error.response.statusText}`);
        }
    } else if (error.request) {
        alert("Ошибка сети: сервер недоступен");
    } else {
        alert("Произошла неизвестная ошибка");
    }
};
