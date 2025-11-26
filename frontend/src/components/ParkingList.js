// src/components/ParkingList.js
import React, { useState, useEffect } from 'react';

const ParkingList = () => {
    const [parkings, setParkings] = useState([]);
    const [vehicles, setVehicles] = useState([]);
    const [parkingId, setParkingId] = useState('');
    const [vehicleId, setVehicleId] = useState('');
    const [message, setMessage] = useState('');

    useEffect(() => {
        fetch('http://localhost:8080/parking')
            .then(res => res.json())
            .then(data => setParkings(data));

        fetch('http://localhost:8080/vehicles')
            .then(res => res.json())
            .then(data => setVehicles(data));
    }, []);

    const handleReserve = async () => {
        if (!parkingId || !vehicleId) {
            setMessage({ type: 'error', text: 'Выберите парковку и транспорт' });
            return;
        }

        const now = new Date();
        const expires = new Date(now.getTime() + 60 * 60 * 1000);

        const reservation = {
            vehicle_id: parseInt(vehicleId),
            starts_at: now.toISOString(),
            expires_at: expires.toISOString()
        };

        try {
            const res = await fetch(`http://localhost:8080/parking/${parkingId}/reserve`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(reservation)
            });

            if (res.ok) {
                setMessage({ type: 'success', text: '✅ Успешно забронировано!' });
                fetch('http://localhost:8080/parking').then(r => r.json()).then(setParkings);
            } else {
                const err = await res.text();
                setMessage({ type: 'error', text: `❌ ${'Парковка заполнена'}` });
            }
        } catch (err) {
            setMessage({ type: 'error', text: '❌ Ошибка сети' });
        }
    };

    return (
        <div>
            <h2>Парковки</h2>
            <div className="scrollable">
                {parkings.length === 0 ? (
                    <p className="empty">Нет данных</p>
                ) : (
                    parkings.map(p => (
                        <div key={p.id} className="item">
                            <div className="vehicle-name">{p.name}</div>
                            <div className="parking-capacity">
                                🅿️ Вместимость: {p.capacity}
                            </div>
                            <div className="parking-coords">
                                📍 {p.latitude.toFixed(4)}, {p.longitude.toFixed(4)}
                            </div>
                        </div>
                    ))
                )}
            </div>

            <div className="controls">
                <select value={parkingId} onChange={e => setParkingId(e.target.value)}>
                    <option value="">Выберите парковку</option>
                    {parkings.map(p => (
                        <option key={p.id} value={p.id}>{p.name}</option>
                    ))}
                </select>

                <select value={vehicleId} onChange={e => setVehicleId(e.target.value)}>
                    <option value="">Выберите транспорт</option>
                    {vehicles.map(v => (
                        <option key={v.id} value={v.id}>{v.name} ({v.type})</option>
                    ))}
                </select>

                <button onClick={handleReserve}>Забронировать</button>
            </div>

            {message.text && (
                <p className={`message ${message.type}`}>{message.text}</p>
            )}
        </div>
    );
};

export default ParkingList;