import React, { useState, useEffect } from 'react';
import {getAllParkingLots, getAllVehicles} from "../api/api.ts";
import ParkingItem from "./ParkingItem";

const ParkingList = ({parkings, setParkings, onParkingSelect, selectedId}) => {
    const [loading, setLoading] = useState(true);
    const [isError, setIsError] = useState(false)

    useEffect (() => {
        const fetchData = async () => {
            setLoading(true)
            try {
                const response = await getAllParkingLots()
                setParkings(response)

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
                ) : (parkings.length === 0 ? (
                    <>Нет данных</>
                ) : (
                    <>
                        {parkings.map((item) => (
                            <ParkingItem key={item.id}
                                         item={item}
                            onClick={() => onParkingSelect(item.id)}
                            isSelected={selectedId === item.id}
                            />

                        ))}
                    </>
                ))}
            </div>
        </>
    )
};

export default ParkingList;