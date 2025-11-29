import {React} from "react";

const ParkingItem = ({item}) => {
    function getDate(timestamp) {
        const secondsInMs = timestamp.seconds * 1000;
        const nanosInMs = timestamp.nanos / 1000000;
        const totalMilliseconds = secondsInMs + nanosInMs;
        return new Intl.DateTimeFormat('ru-RU', {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit',
            hour12: false // Use 24-hour format
        }).format(new Date(totalMilliseconds))
    }
    return (
        <div className="card-item my-2 p-3 card element-hover">
            <div className="d-flex justify-content-between align-items-start">
                <div>
                    <h5 className="card-title mb-1">
                        <i className="bi bi-p-circle me-2"></i>
                        {item.name}
                    </h5>
                    <p className="card-text text-muted mb-1">Вместимость: {item.capacity}</p>
                    <p className="coordinates mb-0">
                        <i className="bi bi-geo-alt text-danger"></i>
                        {item.longitude.toFixed(4)}, {item.latitude.toFixed(4)}
                    </p>
                </div>
            </div>
            <div className="mt-2">
                <div className="progress mt-2">
                    <div className="progress-bar bg-success"></div>
                </div>
                <small className="text-muted">Занято 43 из 50 мест</small>
            </div>
            <div className="mt-2 text-end">
                <small className="text-muted">
                    Добавлено: {getDate(item.created_at)}
                </small>
            </div>
        </div>
    )
}

export default ParkingItem