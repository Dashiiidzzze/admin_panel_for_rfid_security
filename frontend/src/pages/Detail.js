import React, { useState, useEffect, useRef } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getData, updateData } from '../components/CheckErrors';

const Detail = () => {
    const { id } = useParams(); // Получаем ID товара из URL
    const navigate = useNavigate();
    
    const [itemData, setItemData] = useState({
        name: "",
        position: "",
        phone: "",
        keyNumber: "",
        zone: {
            flightZone: false,
            cleanZone: false,
            runway: false,
            baggageZone: false,
            controlTower: false
        }
    });
    const nameRef = useRef(null);
    const positionRef = useRef(null);
    const phoneRef = useRef(null);
    const keyNumberRef = useRef(null);
    
    // Загружаем товар при монтировании компонента
    useEffect(() => {
        async function loadItem() {
            try {
                const response = await getData(`${process.env.REACT_APP_API_URL}/items/${id}`);
                setItemData(response); // Обновляем состояние
            } catch (error) {
                console.error("Ошибка загрузки:", error);
            }
        }
        loadItem();
    }, [id]); // Запускать при изменении `id`

    useEffect(() => {
        if (nameRef.current) nameRef.current.value = itemData.name || '';
        if (positionRef.current) positionRef.current.value = itemData.position || '';
        if (phoneRef.current) phoneRef.current.value = itemData.phone || '';
        if (keyNumberRef.current) keyNumberRef.current.value = itemData.keyNumber || '';
    }, [itemData]); // Заполняем input после загрузки данных

    const handleZoneChange = (event, zoneKey) => {
        setItemData(prevData => ({
            ...prevData,
            zone: { ...prevData.zone, [zoneKey]: event.target.value === "true" }
        }));
    };

    // Функция обновления товара
    const handleSubmit = async (e) => {
        e.preventDefault();

        const updatedItem = {
            name: nameRef.current.value,
            position: positionRef.current.value,
            phone: phoneRef.current.value,
            keyNumber: keyNumberRef.current.value,
            zone: itemData.zone
        };

        try {
            await updateData(
                `${process.env.REACT_APP_API_URL}/items/${id}`, JSON.stringify(updatedItem)); // Сериализация объекта в строку JSON
            navigate('/');
        } catch (error) {
            console.error("Ошибка обновления:", error);
        }
    };

    // Проверяем, что itemData и itemData.zone существуют
    if (!itemData || !itemData.zone) {
        return <div>Загрузка...</div>; // Можно вернуть компонент "Загрузка", если данные еще не загружены
    }

    return (
        <div>
            <h1>Редактирование карточки сотрудника</h1>
            <form onSubmit={handleSubmit}>
                <div className='inform'>
                    <h3>Общая информация:</h3>
                    <label>
                        ФИО:
                        <input type="text" ref={nameRef} required />
                    </label>
                    <br />
                    <label>
                        Должность:
                        <input type="text" ref={positionRef} required />
                    </label>
                    <br />
                    <label>
                        Телефон:
                        <input type="text" ref={phoneRef} required />
                    </label>
                    <br />
                    <label>
                        Номер ключа:
                        <input type="text" ref={keyNumberRef} required />
                    </label>
                    <br />
                    <button className='save' type="submit">Сохранить</button>
                </div>

                <div className='zones'>
                    <h3>Доступ в зоны:</h3>
                    <label>
                        Зона вылета/прилета:
                        <select value={itemData.zone.flightZone} onChange={(e) => handleZoneChange(e, "flightZone")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                    <label>
                        Чистая зона:
                        <select value={itemData.zone.cleanZone} onChange={(e) => handleZoneChange(e, "cleanZone")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                    <label>
                        Взлетно-посадочная полоса:
                        <select value={itemData.zone.runway} onChange={(e) => handleZoneChange(e, "runway")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                    <label>
                        Зона обслуживания багажа:
                        <select value={itemData.zone.baggageZone} onChange={(e) => handleZoneChange(e, "baggageZone")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                    <label>
                        Диспетчерская:
                        <select value={itemData.zone.controlTower} onChange={(e) => handleZoneChange(e, "controlTower")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                </div>
            </form>
        </div>
    );
}

export default Detail