// src/components/Vehicles.js
import React, { useState, useEffect } from 'react';
import { getAllVehicles } from '../api/api.ts';
import VehicleItem from "./VehicleItem";

const getVehicleIcon = (type) => {
    const icons = {
        car: '',
        truck: '🚛',
        motorcycle: '🏍️',
        bike: '🚲',
        scooter: '🛴'
    };
    return icons[type] || '🚗';
};

const Vehicles = ({vehicles, setVehicles}) => {
    const [loading, setLoading] = useState(true);
    const [isError, setIsError] = useState(false)

    useEffect (() => {
        const fetchData = async () => {
            setLoading(true)
            try {
                const response = await getAllVehicles()
                setVehicles(response)

            } catch (er) {
                setIsError(true)
            }
            setLoading(false)
        }
        fetchData()

    }, []);

    return (
        <>
            <div className="list">
                {loading ? (
                    <div className="spinner-border" role="status">
                        <span className="sr-only">Loading...</span>
                    </div>
                ) : (vehicles.length === 0 ? (
                    <>Нет данных</>
                ) : (
                    <>
                        {vehicles.map((item, index) => (
                            <VehicleItem key={index} item={item}/>

                        ))}
                    </>
                ))}
            </div>
        </>
    )
};

export default Vehicles;