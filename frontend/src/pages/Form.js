import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getData, postData } from '../components/CheckErrors';

const Form = () => {
    const navigate = useNavigate();
    // состоние для должностей
    const [positions, setPositions] = useState([]);
    // Состояние для всех полей формы
    const [formData, setFormData] = useState({
        name: "",
        position: "",
        phone: "",
        keyNumber: "",
        zone: {
            flightZone: false,
            clearZone: false,
            runaway: false,
            baggageZone: false,
            controlTower: false
        }
    });

    // Состояние для ошибки валидации
    const [phoneError, setPhoneError] = useState("");

    useEffect(() => {
        // загрузка списка должностей с API при монтировании компонента
        const fetchPositions = async () => {
            try {
                const data = await getData(`${process.env.REACT_APP_API_URL}/api/positions`);
                // Преобразуем массив объектов в массив строк с названиями должностей
                const positionNames = data.dolzhnosti.map(d => d.dolzhnost);
                setPositions(positionNames); 
            } catch (error) {
                console.error("Ошибка при получении списка должностей:", error);
            }
        };

        fetchPositions();
    }, []);

    //проверка номера телефона
    const validatePhone = (phone) => {
        const cleaned = phone.trim().replace(/\s+/g, '');
        const phoneRegex = /^\+7\d{10}$/;
        return phoneRegex.test(cleaned);
    }

    // Обработчик изменения полей ввода
    const handleChange = (e) => {
        const { name, value } = e.target;
        
        if (name === "phone") {
            // Очищаем ошибку при изменении поля
            setPhoneError("");
        }
        
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    // Обработчик изменения выбора зоны (true/false)
    const handleZoneChange = (e, zoneKey) => {
        setFormData(prev => ({
            ...prev,
            zone: {
                ...prev.zone,
                [zoneKey]: e.target.value === "true"
            }
        }));
    };

    // Отправка данных на сервер
    const handleSubmit = async (e) => {
        e.preventDefault();

        const cleanedPhone = formData.phone.trim();
        
        // Валидация телефона
        if (!validatePhone(cleanedPhone)) {
            setPhoneError("Формат телефона должен быть: +7XXXXXXXXXX (11 цифр)");
            return;
        }
        
        try {
            await postData(`${process.env.REACT_APP_API_URL}/items`, formData);
            navigate("/");
        } catch (error) {
            console.error("Ошибка при отправке данных:", error);
        }
    };

    return (
        <div>
            <h1>Создание карточки сотрудника</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <h3>Общая информация:</h3>
                    <label>
                        ФИО:
                        <input type="text" name="name" value={formData.name} onChange={handleChange} required />
                    </label>
                    <br />
                    <label>
                        Должность:
                        <select name="position" value={formData.position} onChange={handleChange} required>
                            <option value="">Выберите должность</option>
                            {positions.map((pos, index) => (
                                <option key={index} value={pos}>{pos}</option>
                            ))}
                        </select>
                    </label>
                    <br />
                    <label>
                        Телефон:
                        <input 
                            type="text" 
                            name="phone" 
                            value={formData.phone} 
                            onChange={handleChange} 
                            required 
                            style={{ 
                                border: phoneError ? '2px solid red' : '1px solid #ccc',
                                outline: 'none'
                            }}

                        />
                    </label>
                    {phoneError && <div style={{color: 'red', fontSize: '0.8em', marginTop: '5px'}}>{phoneError}</div>}
                    <br />
                    <label>
                        Номер ключа:
                        <input type="text" name="keyNumber" value={formData.keyNumber} onChange={handleChange} required />
                    </label>
                    <br />
                    <button className='save' type="submit">Сохранить</button>
                </div>
                <div>
                    <h3>Доступ в зоны:</h3>
                    <label>
                        Зона вылета/прилета:
                        <select value={formData.zone.flightZone} onChange={(e) => handleZoneChange(e, "flightZone")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                    <label>
                        Чистая зона:
                        <select value={formData.zone.clearZone} onChange={(e) => handleZoneChange(e, "clearZone")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                    <label>
                        Взлетно-посадочная полоса:
                        <select value={formData.zone.runaway} onChange={(e) => handleZoneChange(e, "runaway")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                    <label>
                        Зона обслуживания багажа:
                        <select value={formData.zone.baggageZone} onChange={(e) => handleZoneChange(e, "baggageZone")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                    <label>
                        Диспетчерская:
                        <select value={formData.zone.controlTower} onChange={(e) => handleZoneChange(e, "controlTower")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                </div>
            </form>
        </div>
    );
};

export default Form;