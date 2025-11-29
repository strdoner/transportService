import {Parking, Transport} from "./objects/objects";

export async function getAllVehicles(): Promise<Transport[]> {
    try {
        const response = await fetch('http://localhost:8080/vehicles')

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`)
        }

        return await response.json();
    } catch (er) {
        console.error('Error via fetching vehicles: ', er)
        throw er
    }
}

export async function getAllParkingLots(): Promise<Parking[]> {
    try {
        const response = await fetch('http://localhost:8080/parking')

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`)
        }

        return await response.json();
    } catch (er) {
        console.error('Error via fetching parkings: ', er)
        throw er
    }
}

