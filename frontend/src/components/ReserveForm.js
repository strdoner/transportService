import {React} from 'react'
import VehicleItem from "./VehicleItem";

const ReserveForm = ({vehicles, parkings}) => {
    // TODO selecting transport and parking values by clicking on item in the list above the form
    // TODO formatting data from form and sending it to backend on address POST localhost:8080/parking/{id}/reserve
    // TODO visualization of free space (progress bar), vehicle number from db
    return (
        <div>
            <form action="POST">
                <div className="row">
                    <div className="col-6">
                        <span className="ps-2">Выберите транспорт</span>
                        <select className="form-select" id="transportSelect" required>
                            {vehicles.map((item, index) => (
                                <option value={item.id}>{item.name}</option>
                            ))}
                        </select>
                    </div>
                    <div className="col-6">
                        <span className="ps-2">Выберите парковку</span>
                        <select className="form-select" id="transportSelect" required>
                            {parkings.map((item, index) => (
                                <option value={item.id}>{item.name}</option>
                            ))}
                        </select>
                    </div>
                    <div className="col-4">
                        <span>Дата бронирования</span>
                        <input type="date" className="form-control" id="reservationDate" required/>
                    </div>
                    <div className="col-4">
                        <span>Время начала</span>
                        <input type="time" className="form-control" id="startTime" required/>
                    </div>
                    <div className="col-4">
                        <span>Время окончания</span>
                        <input type="time" className="form-control" id="endTime" required/>
                    </div>
                    <div className="col-12 pt-3">
                        <div className="d-grid">
                            <button type="submit" className="btn btn-primary btn-lg">
                                Забронировать место
                            </button>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    )
}

export default ReserveForm