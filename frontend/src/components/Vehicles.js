import React, { useState, useEffect } from 'react';
import { getAllVehicles } from '../api/api.ts';
import VehicleItem from "./VehicleItem";

const Vehicles = ({vehicles, setVehicles, onVehicleSelect, selectedId}) => {
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
                        {vehicles.map((item) => (
                            <VehicleItem key={item.id}
                                         item={item}
                                         onClick={() => onVehicleSelect(item.id)}
                                         isSelected={selectedId === item.id}
                            />

                        ))}
                    </>
                ))}
            </div>
        </>
    )
};

export default Vehicles;