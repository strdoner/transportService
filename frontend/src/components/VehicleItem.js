import {React} from "react";

const VehicleItem = ({item}) => {
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
                        <i className="bi bi-car-front me-2"></i>
                        {item.name}
                    </h5>
                    <p className="card-text text-muted mb-1">Тип: {item.type}</p>
                    <p className="coordinates mb-0">
                        <i className="bi bi-geo-alt text-danger"></i>
                        {item.longitude.toFixed(4)}, {item.latitude.toFixed(4)}
                    </p>
                </div>
                <span className="badge bg-success">B456DE</span>
            </div>
            <div className="mt-2 text-end">
                <small className="text-muted">
                    Обновлено: {getDate(item.updated_at)}
                </small>
            </div>
        </div>
    )
}

export default VehicleItem