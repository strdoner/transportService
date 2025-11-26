// src/components/Vehicles.js
import React, { useState, useEffect } from 'react';

const getVehicleIcon = (type) => {
    const icons = {
        car: '🚗',
        truck: '🚛',
        motorcycle: '🏍️',
        bike: '🚲',
        scooter: '🛴'
    };
    return icons[type] || '🚗';
};

const Vehicles = () => {
    const [vehicles, setVehicles] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch('http://localhost:8080/vehicles')
            .then(res => res.json())
            .then(data => {
                setVehicles(data);
                setLoading(false);
            })
            .catch(() => setLoading(false));
    }, []);

    return (
        <div>
            <h2>Транспорт</h2>
            <div className="scrollable">
                {loading ? (
                    <p>Загрузка...</p>
                ) : vehicles.length === 0 ? (
                    <p className="empty">Нет данных</p>
                ) : (
                    vehicles.map(v => (
                        <div key={v.id} className="item">
                            <div className="vehicle-name">{v.name}</div>
                            <div className="vehicle-type">
                                {getVehicleIcon(v.type)} Тип: {v.type}
                            </div>
                            <div className="vehicle-coords">
                                📍 {v.latitude.toFixed(4)}, {v.longitude.toFixed(4)}
                            </div>
                        </div>
                    ))
                )}
            </div>
        </div>
    );
};

export default Vehicles;