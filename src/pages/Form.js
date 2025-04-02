// import React, { useState } from 'react';
// import { useNavigate } from 'react-router-dom';
// import { postData } from '../components/CheckErrors';

// const Form = () => {
//     const navigate = useNavigate();

//     // Состояние для всех полей формы
//     const [formData, setFormData] = useState({
//         name: "",
//         position: "",
//         phone: "",
//         keyNumber: "",
//         zone: {
//             flightZone: false,
//             cleanZone: false,
//             runway: false,
//             baggageZone: false,
//             controlTower: false
//         }
//     });

//     //проверка номера телефона
//     const validatePhone = (phone) => {
//         const phoneRegex = /^\+?\d{11}$/;
//         return phoneRegex.test(phone);

//     }

//     // Обработчик изменения полей ввода
//     const handleChange = (e) => {
//         const { name, value } = e.target;
//         setFormData(prev => ({
//             ...prev,
//             [name]: value
//         }));
//     };

//     // Обработчик изменения выбора зоны (true/false)
//     const handleZoneChange = (e, zoneKey) => {
//         setFormData(prev => ({
//             ...prev,
//             zone: {
//                 ...prev.zone,
//                 [zoneKey]: e.target.value === "true"
//             }
//         }));
//     };

//     // Отправка данных на сервер
//     const handleSubmit = async (e) => {
//         e.preventDefault();
//         if (!validatePhone(formData.phone)) {
//             alert("неверный формат номера телефона. Пример +89231335643")
//             return;
//         }
//         try {
//             await postData("http://localhost:5000/items", formData);
//             // await axios.post("http://localhost:5000/items", formData, {
//             //     headers: { "Content-Type": "application/json" }
//             // });
//             navigate("/");
//         } catch (error) {
//             console.error("Ошибка при отправке данных:", error);
//         }
//     };

//     return (
//         <div>
//             <h1>Создание карточки сотрудника</h1>
//             <form onSubmit={handleSubmit}>
//                 <div>
//                     <h3>Общая информация:</h3>
//                     <label>
//                         ФИО:
//                         <input type="text" name="name" value={formData.name} onChange={handleChange} required />
//                     </label>
//                     <br />
//                     <label>
//                         Должность:
//                         <select name="position" value={formData.position} onChange={handleChange} required>
//                             <option value="">Выберите должность</option>
//                             <option value="Пилот">Пилот</option>
//                             <option value="Стюардесса">Стюардесса</option>
//                             <option value="Деспетчер">Деспетчер</option>
//                             <option value="Охранник">Охранник</option>
//                             <option value="Охранник">Грузчик</option>
//                             <option value="Охранник">Уборщик</option>
//                             <option value="Охранник">Досмотрщик</option>
                           
//                         </select>
//                     </label>
//                     <br />
//                     <label>
//                         Телефон:
//                         <input type="text" name="phone" value={formData.phone} onChange={handleChange} required />
//                     </label>
//                     <br />
//                     <label>
//                         Номер ключа:
//                         <input type="text" name="keyNumber" value={formData.keyNumber} onChange={handleChange} required />
//                     </label>
//                     <br />
//                     <button className='save' type="submit">Сохранить</button>
//                 </div>
//                 <div>
//                     <h3>Доступ в зоны:</h3>
//                     <label>
//                         Зона вылета/прилета:
//                         <select value={formData.zone.flightZone} onChange={(e) => handleZoneChange(e, "flightZone")}>
//                             <option value="true">Да</option>
//                             <option value="false">Нет</option>
//                         </select>
//                     </label>
//                     <br />
//                     <label>
//                         Чистая зона:
//                         <select value={formData.zone.cleanZone} onChange={(e) => handleZoneChange(e, "cleanZone")}>
//                             <option value="true">Да</option>
//                             <option value="false">Нет</option>
//                         </select>
//                     </label>
//                     <br />
//                     <label>
//                         Взлетно-посадочная полоса:
//                         <select value={formData.zone.runway} onChange={(e) => handleZoneChange(e, "runway")}>
//                             <option value="true">Да</option>
//                             <option value="false">Нет</option>
//                         </select>
//                     </label>
//                     <br />
//                     <label>
//                         Зона обслуживания багажа:
//                         <select value={formData.zone.baggageZone} onChange={(e) => handleZoneChange(e, "baggageZone")}>
//                             <option value="true">Да</option>
//                             <option value="false">Нет</option>
//                         </select>
//                     </label>
//                     <br />
//                     <label>
//                         Диспетчерская:
//                         <select value={formData.zone.controlTower} onChange={(e) => handleZoneChange(e, "controlTower")}>
//                             <option value="true">Да</option>
//                             <option value="false">Нет</option>
//                         </select>
//                     </label>
//                 </div>
                
//             </form>
//             </div>
//     );
// };

// export default Form;


import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { postData } from '../components/CheckErrors';

const Form = () => {
    const navigate = useNavigate();

    // Состояние для всех полей формы
    const [formData, setFormData] = useState({
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

    // Состояние для ошибки валидации
    const [phoneError, setPhoneError] = useState("");

    //проверка номера телефона
    const validatePhone = (phone) => {
        const phoneRegex = /^\+?\d{11}$/;
        return phoneRegex.test(phone);
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
        
        // Валидация телефона
        if (!validatePhone(formData.phone)) {
            setPhoneError("Неверный формат номера телефона. Пример: +79231335643");
            return;
        }
        
        try {
            await postData("http://localhost:5000/items", formData);
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
                            <option value="Пилот">Пилот</option>
                            <option value="Стюардесса">Стюардесса</option>
                            <option value="Деспетчер">Деспетчер</option>
                            <option value="Охранник">Охранник</option>
                            <option value="Охранник">Грузчик</option>
                            <option value="Охранник">Уборщик</option>
                            <option value="Охранник">Досмотрщик</option>
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
                        <select value={formData.zone.cleanZone} onChange={(e) => handleZoneChange(e, "cleanZone")}>
                            <option value="true">Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </label>
                    <br />
                    <label>
                        Взлетно-посадочная полоса:
                        <select value={formData.zone.runway} onChange={(e) => handleZoneChange(e, "runway")}>
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