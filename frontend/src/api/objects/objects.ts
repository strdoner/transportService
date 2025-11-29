export interface Timestamp {
    seconds: number;
    nanos: number;
}

export interface Transport {
    id: number;
    name: string;
    type: string;
    latitude: number;
    longitude: number;
    updated_at: Timestamp;
}

export interface Parking {
    id: number;
    name: string;
    capacity: number;
    latitude: number;
    longitude: number;
    created_at: Timestamp;
}

