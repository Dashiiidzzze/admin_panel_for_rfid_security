import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getData, updateData } from '../components/CheckErrors';

const Detail = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [positions, setPositions] = useState([]);
    const [phoneError, setPhoneError] = useState("");

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

    // Получаем список должностей
    useEffect(() => {
        const fetchPositions = async () => {
            try {
                const data = await getData(`${process.env.REACT_APP_API_URL}/api/positions`);
                const positionNames = data.dolzhnosti.map(d => d.dolzhnost);
                setPositions(positionNames);
            } catch (error) {
                console.error("Ошибка при получении должностей:", error);
            }
        };
        fetchPositions();
    }, []);

    // Загружаем данные сотрудника
    useEffect(() => {
        const loadItem = async () => {
            try {
                const response = await getData(`${process.env.REACT_APP_API_URL}/items/${id}`);
                const staffData = response.staff[0];

                // Убедимся, что номер телефона — строка и без пробелов
                setFormData({
                    ...staffData,
                    phone: staffData.phone?.toString().trim() || "",
                    zone: staffData.zone || {
                        flightZone: false,
                        clearZone: false,
                        runaway: false,
                        baggageZone: false,
                        controlTower: false
                    }
                });
            } catch (error) {
                console.error("Ошибка загрузки данных:", error);
            }
        };
        loadItem();
    }, [id]);

    // Проверка номера телефона: +7 и 10 цифр
    const validatePhone = (phone) => {
        const cleaned = phone.trim().replace(/\s+/g, '');
        const phoneRegex = /^\+7\d{10}$/;
        return phoneRegex.test(cleaned);
    };

    // Обработка изменений в input'ах
    const handleChange = (e) => {
        const { name, value } = e.target;

        if (name === "phone") {
            setPhoneError("");
        }

        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    // Обработка переключения доступа к зонам
    const handleZoneChange = (e, zoneKey) => {
        const value = e.target.value === "true";
        setFormData(prev => ({
            ...prev,
            zone: {
                ...prev.zone,
                [zoneKey]: value
            }
        }));
    };

    // Отправка формы
    const handleSubmit = async (e) => {
        e.preventDefault();

        const cleanedPhone = formData.phone.trim();

        if (!validatePhone(cleanedPhone)) {
            setPhoneError("Формат телефона должен быть: +7XXXXXXXXXX (11 цифр)");
            return;
        }

        try {
            await updateData(
                `${process.env.REACT_APP_API_URL}/items/${id}`,
                JSON.stringify({
                    ...formData,
                    phone: cleanedPhone
                })
            );
            navigate('/');
        } catch (error) {
            console.error("Ошибка при обновлении данных:", error);
        }
    };

    if (!formData || !formData.zone) return <div>Загрузка...</div>;

    return (
        <div>
            <h1>Редактирование карточки сотрудника</h1>
            <form onSubmit={handleSubmit}>
                <div className='inform'>
                    <h3>Общая информация:</h3>

                    <label>
                        ФИО:
                        <input
                            type="text"
                            name="name"
                            value={formData.name}
                            onChange={handleChange}
                            required
                        />
                    </label><br />

                    <label>
                        Должность:
                        <select
                            name="position"
                            value={formData.position}
                            onChange={handleChange}
                            required
                        >
                            <option value="">Выберите должность</option>
                            {positions.map((pos, index) => (
                                <option key={index} value={pos}>{pos}</option>
                            ))}
                        </select>
                    </label><br />

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
                        <input
                            type="text"
                            name="keyNumber"
                            value={formData.keyNumber}
                            onChange={handleChange}
                            required
                        />
                    </label><br />

                    <button className='save' type="submit">Сохранить</button>
                </div>

                <div className='zones'>
                    <h3>Доступ в зоны:</h3>

                    <label>
                        Зона вылета/прилета:
                        <select
                            value={formData.zone.flightZone}
                            onChange={(e) => handleZoneChange(e, "flightZone")}
                        >
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label><br />

                    <label>
                        Чистая зона:
                        <select
                            value={formData.zone.clearZone}
                            onChange={(e) => handleZoneChange(e, "clearZone")}
                        >
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label><br />

                    <label>
                        Взлетно-посадочная полоса:
                        <select
                            value={formData.zone.runaway}
                            onChange={(e) => handleZoneChange(e, "runaway")}
                        >
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label><br />

                    <label>
                        Зона обслуживания багажа:
                        <select
                            value={formData.zone.baggageZone}
                            onChange={(e) => handleZoneChange(e, "baggageZone")}
                        >
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label><br />

                    <label>
                        Диспетчерская:
                        <select
                            value={formData.zone.controlTower}
                            onChange={(e) => handleZoneChange(e, "controlTower")}
                        >
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label><br />
                </div>
            </form>
        </div>
    );
};

export default Detail;
